package gen

import (
	//"fmt"
	"io"

	"github.com/glycerine/greenpack/cfg"
)

func getFromSQL(w io.Writer, cfg *cfg.GreenConfig) *getFromSqlGen {
	return &getFromSqlGen{
		p:   printer{w: w, cfg: cfg},
		cfg: cfg,
	}
}

type getFromSqlGen struct {
	passes
	p    printer
	fuse []byte
	cfg  *cfg.GreenConfig
}

func (e *getFromSqlGen) MethodPrefix() string {
	return e.cfg.MethodPrefix
}

func (e *getFromSqlGen) Method() Method { return GetFromSQL }

func (e *getFromSqlGen) Apply(dirs []string) error {
	return nil
}

func (e *getFromSqlGen) writeAndCheck(typ string, argfmt string, arg interface{}) {
	//e.p.printf("\nerr = en.Write%s(%s)", typ, fmt.Sprintf(argfmt, arg))
	//e.p.printf("%s %s", arg, typ)
}

func (e *getFromSqlGen) fuseHook() {
	if len(e.fuse) > 0 {
		e.appendraw(e.fuse)
		e.fuse = e.fuse[:0]
	}
}

func (e *getFromSqlGen) Fuse(b []byte) {
	if len(e.fuse) > 0 {
		e.fuse = append(e.fuse, b...)
	} else {
		e.fuse = b
	}
}

func (e *getFromSqlGen) Execute(p Elem) error {
	if !e.p.ok() {
		return e.p.err
	}
	p = e.applyall(p)
	if p == nil {
		return nil
	}
	if !IsPrintable(p) {
		return nil
	}

	// only gen SQL for structs
	_, isStruct := p.(*Struct)
	if !isStruct {
		return nil
	}

	e.p.printf(`
// Note that the github.com/go-sql-driver/mysql driver will need parseTime=true
// added to the DSN string in order to pull DATETIME from sql into time.Time in Go.
// For example:
// db, err := sql.Open("mysql", "user:pw@tcp(localhost:3306)/databasename?parseTime=true")
//
// Also storing time.Time from Go into MariaDB will be lossy because Go keeps
// nanoseconds but MariaDB only stores microseconds.
`)

	e.p.printf("\nfunc (%s %s) %sGetFromSQL(ctx context.Context, db *sql.DB, dbName, tableName, where string) (res []%s, sqlSel string, err error) {\n", p.Varname(), imutMethodReceiver(p), e.cfg.MethodPrefix, imutMethodReceiver(p))
	e.p.printf(`
		if strings.HasPrefix(where, "select") {
			sqlSel = where
		} else {
	`)

	next(e, p, nil)
	e.p.nakedReturn()
	return e.p.err
}

func (e *getFromSqlGen) gStruct(s *Struct, x *extra) {
	if !e.p.ok() {
		return
	}

	e.structmap(s)

	return
}

func (e *getFromSqlGen) appendraw(bts []byte) {
	return
	e.p.print("\nerr = en.Append(")
	for i, b := range bts {
		if i != 0 {
			e.p.print(", ")
		}
		e.p.printf("0x%x", b)
	}
	e.p.print(")\nif err != nil { return err }")
}

func (e *getFromSqlGen) structmap(s *Struct) {
	//nfields := len(s.Fields) - s.SkipCount

	recv := s.TypeName() // imutMethodReceiver(s)
	e.p.printf(`

// get from sql and write into '%s'
`, recv)

	sel := "sqlSel = \"select "
	var actuals string

	first := true
	columns := make(map[string]bool) // insure uniqueness of column names
	columns["rowid"] = true
	columns["updatetm"] = true

	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		if !e.p.ok() {
			return
		}

		// or .FieldName ?
		fld := s.Fields[i].FieldTag // the string inside the `msg:""` tag

		// handle collisions with rowid and updatetm, which are always
		// the first two columns.
		if fld == "rowid" {
			fld = "rowid_"
		} else if fld == "updatetm" {
			fld = "updatetm_"
		}

		// make sure all columns are uniquely named
		for {
			_, isDup := columns[fld]
			if isDup {
				fld += "x"
			} else {
				columns[fld] = true
				break
			}
		}

		if first {
			first = false
		} else {
			sel += ", "
			actuals += ","
		}
		fld = "”" + fld + "”"
		sel += fld
		actuals += "&v." + s.Fields[i].FieldName

		// type-switch dispatch to the correct method
		//next(e, s.Fields[i].FieldElem, &extra{pointerOrIface: s.Fields[i].IsPointer || s.Fields[i].IsIface})

	}

	sel += " from \" + dbName + \".\" + tableName "
	e.p.printf(sel)
	e.p.printf("\nsqlSel = strings.ReplaceAll(sqlSel, \"”\", \"`\") + \" \" + where")
	e.p.printf(`
}

if db == nil {
    return // just providing the sqlSel statement.
}

var rows *sql.Rows
rows, err = db.QueryContext(ctx, sqlSel)
if err != nil {
   err = fmt.Errorf("error on select: '%%v'; sql was '%%v'", err, sqlSel)
   return
}
defer rows.Close()

    var v *%v
	for i := 0; rows.Next(); i++ {
		if i == 0 {
			v = z
		} else {
			v = &%v{}
		}
   err = rows.Scan(%v)
   if err != nil {
       return
   }
   res = append(res, v)
}

`, recv, recv, actuals)

}

func (e *getFromSqlGen) gMap(m *Map, x *extra) {
	return
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	vname := m.Varname()
	e.writeAndCheck(mapHeader, lenAsUint32, vname)

	e.p.printf("\nfor %s, %s := range %s {", m.Keyidx, m.Validx, vname)
	e.writeAndCheck(m.KeyTyp, literalFmt, m.Keyidx)
	next(e, m.Value, nil)
	e.p.closeblock()
}

func (e *getFromSqlGen) gPtr(s *Ptr, x *extra) {
	return
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	e.p.print("\n // gPtr.getFromSqlGen():\n")
	e.p.printf("\nif %s == nil { err = en.WriteNil(); if err != nil { return; } } else {", s.Varname())

	next(e, s.Value, nil)
	e.p.closeblock()
}

func (e *getFromSqlGen) gSlice(s *Slice, x *extra) {
	return
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	e.writeAndCheck(arrayHeader, lenAsUint32, s.Varname())
	e.p.rangeBlock(s.Index, s.Varname(), e, s.Els)
}

func (e *getFromSqlGen) gArray(a *Array, x *extra) {
	return
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	// shortcut for [const]byte
	if be, ok := a.Els.(*BaseElem); ok && (be.Value == Byte || be.Value == Uint8) {
		e.p.printf("\nerr = en.WriteBytes(%s[:])", a.Varname())
		e.p.print(errcheck)
		return
	}

	e.writeAndCheck(arrayHeader, literalFmt, a.SizeResolved)
	e.p.rangeBlock(a.Index, a.Varname(), e, a.Els)
}

func (e *getFromSqlGen) gBase(b *BaseElem, x *extra) {
	return
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	vname := b.Varname()
	//if b.Convert {
	//	vname = tobaseConvert(b)
	//}

	if b.Value == IDENT { // unknown identity
		e.p.printf("\n // getFromSqlGen.gBase unknown IDENT '%v': ignore for now.\n", vname)
		e.p.printf("\n // err = %s.%sGetFromSQL(db, tableName)", vname, e.cfg.MethodPrefix)
	} else { // typical case
		e.writeAndCheck(b.BaseName(), literalFmt, vname)
	}
}
