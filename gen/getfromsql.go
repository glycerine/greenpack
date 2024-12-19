package gen

import (
	//"fmt"
	"io"

	"github.com/glycerine/greenpack/cfg"
	"github.com/glycerine/greenpack/green"
)

func getFromSQL(w io.Writer, cfg *cfg.GreenConfig) *storeToSqlGen {
	return &storeToSqlGen{
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

func (e *getFromSqlGen) Method() Method { return StoreToSQL }

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

	e.p.printf("\nfunc (%s %s) %sGetFromSQL(db *sql.DB, dbName, tableName string, reuseStmt *sql.Stmt) (stmt *sql.Stmt, err error) {\n stmt = reuseStmt\n", p.Varname(), imutMethodReceiver(p), e.cfg.MethodPrefix)
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

// create table to store type '%s'
`, recv)
	e.p.printf("if create { sqlCreate := \"CREATE TABLE IF NOT EXISTS \" + dbName + \".\" + tableName + ` (\n")
	e.p.printf(`  rowid bigint AUTO_INCREMENT not null primary key
, updatetm TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
`)

	ins := "sqlIns := \"insert into \" + dbName + \".\" + tableName + \"("
	values := "" // the right number of ?,?,?,... question mark place-holders.
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
			ins += ", "
			values += ","
			actuals += ","
		}
		fld = "”" + fld + "”"
		ins += fld
		values += "?"
		actuals += "z." + s.Fields[i].FieldName

		typ := s.Fields[i].FieldElem // field type, Elem
		ztyp := typ.GetZtype()       // GetZtype provides type info in a uniform way.

		kind := ""
		switch ztyp.Kind {

		case green.Bytes: // []byte:
			kind = "LONGBLOB" // variable len, up to 4GB
		case green.String:
			kind = "LONGTEXT" // variable len, up to 4GB
		case green.Float64:
			kind = "DOUBLE"
		case green.Uint16:
			kind = "UNSIGNED SMALLINT"
		case green.Uint32:
			kind = "UNSIGNED INT"
		case green.Uint64:
			kind = "UNSIGNED BIGINT"
		case green.Uint:
			kind = "UNSIGNED BIGINT"
		case green.Byte, green.Uint8:
			kind = "UNSIGNED TINYINT"
		case green.Int: // as in Go, matches native word size.
			kind = "BIGINT"
		case green.Int8:
			kind = "TINYINT"
		case green.Int16:
			kind = "SMALLINT"
		case green.Int32:
			kind = "int"
		case green.Int64:
			kind = "BIGINT"
		case green.Bool:
			kind = "BOOL"
		case green.Time: // time.Time:
			kind = "DATETIME(6)"
		case green.Duration: // time.Duration
			kind = "BIGINT"
		case green.Float32:
			kind = "FLOAT"
		case green.Complex64:
		case green.Complex128:
		case green.Intf: // interface{}:
		case green.Ext: // extension

			// IDENT means an unrecognized identifier;
			// it typically means a named struct type.
			// The Str field in the Ztype will hold the
			// name of the struct.:
		case green.IDENT:
			// compound types
			// implementation note: should correspond to gen/Elem.:
		case green.BaseElemCat:
		case green.MapCat:
		case green.StructCat:
		case green.SliceCat:
		case green.ArrayCat:
		case green.PointerCat:

		}

		e.p.printf(", %v %v\n", fld, kind)

		// type-switch dispatch to the correct method
		//next(e, s.Fields[i].FieldElem, &extra{pointerOrIface: s.Fields[i].IsPointer || s.Fields[i].IsIface})

	}
	e.p.printf("\n)`\n")
	e.p.printf("// mariaDB needs backtick quoted strings\n")
	e.p.printf("sqlCreate = strings.ReplaceAll(sqlCreate, \"”\", \"`\")")
	e.p.printf(`
  _, err = db.Exec(sqlCreate)
  if err != nil {
     err = fmt.Errorf("error creating table: '%%v'; sql was: '%%v'", err, sqlCreate)
     return
  }
} // end if create

`)

	ins += ") values (" + values + ")\""
	e.p.printf(`
    if stmt == nil {
`)
	e.p.printf(ins)
	e.p.printf("\nsqlIns = strings.ReplaceAll(sqlIns, \"”\", \"`\")")
	e.p.printf(`
	    stmt, err = db.Prepare(sqlIns)
        if err != nil {
            err = fmt.Errorf("error preparing insert: '%%v'; sql was: '%%v'", err, sqlIns)
            return
        }
    }
    var res sql.Result
    res, err = stmt.Exec(%v)
    if err != nil {
        return
    }
    injectedRowID, err = res.LastInsertId()

`, actuals)

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
		e.p.printf("\n // err = %s.%sStoreToSQL(db, tableName)", vname, e.cfg.MethodPrefix)
	} else { // typical case
		e.writeAndCheck(b.BaseName(), literalFmt, vname)
	}
}
