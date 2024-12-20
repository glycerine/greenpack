package gen

import (
	//"fmt"
	"io"

	"github.com/glycerine/greenpack/cfg"
	"github.com/glycerine/greenpack/green"
)

func storeToSQL(w io.Writer, cfg *cfg.GreenConfig) *storeToSqlGen {
	return &storeToSqlGen{
		p:   printer{w: w, cfg: cfg},
		cfg: cfg,
	}
}

type storeToSqlGen struct {
	passes
	p    printer
	fuse []byte
	cfg  *cfg.GreenConfig
}

func (e *storeToSqlGen) MethodPrefix() string {
	return e.cfg.MethodPrefix
}

func (e *storeToSqlGen) Method() Method { return StoreToSQL }

func (e *storeToSqlGen) Apply(dirs []string) error {
	return nil
}

func (e *storeToSqlGen) writeAndCheck(typ string, argfmt string, arg interface{}) {
	//e.p.printf("\nerr = en.Write%s(%s)", typ, fmt.Sprintf(argfmt, arg))
	//e.p.printf("%s %s", arg, typ)
}

func (e *storeToSqlGen) fuseHook() {
	if len(e.fuse) > 0 {
		e.appendraw(e.fuse)
		e.fuse = e.fuse[:0]
	}
}

func (e *storeToSqlGen) Fuse(b []byte) {
	if len(e.fuse) > 0 {
		e.fuse = append(e.fuse, b...)
	} else {
		e.fuse = b
	}
}

func (e *storeToSqlGen) Execute(p Elem) error {
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

	//e.p.comment(fmt.Sprintf("%sStoreToSQL implements msgp.ToSQL", e.cfg.MethodPrefix))

	e.p.printf(`

// %vStoreTogSQL stores %v into the table specified by tableName,
// that resides inside the database dbName. If create is true,
// then we will attempt to create the table if it does not
// already exist. The resuseStmt can be nil. If it is nil,
// then a new insert statement will be prepared and
// returned in stmt. On subsequent calls, this stmt
// should be passed in as the reuseStmt parameter, in
// order to allow efficient re-use of the prepared statement.
// After the insert, the newly added rowid will be
// returned in injectedRowID. The insert SQL is returned in
// in sqlIns. If db is nil, only the insert SQL string
// will be returned, and the database will not be contacted.
// Similarly, sqlCreate will return the table creation SQL;
// if create was true. The rest will be inserted too, in
// one big batch, or rest can be nil if there is
// nothing more to insert.
`, e.cfg.MethodPrefix, p.Varname())

	e.p.printf("func (%s %s) %sStoreToSQL(db *sql.DB, dbName, tableName string, create bool, reuseStmt *sql.Stmt, rest []%v) (stmt *sql.Stmt, injectedRowID int64, sqlIns, sqlCreate string, err error) {\n stmt = reuseStmt\n", p.Varname(), imutMethodReceiver(p), e.cfg.MethodPrefix, imutMethodReceiver(p))

	e.p.printf(`
      var tx *sql.Tx
      if db != nil {
          tx, err = db.BeginTx(context.Background(), nil)
          if err != nil {
              return
          }
          defer tx.Rollback()
      }
`)

	next(e, p, nil)
	e.p.nakedReturn()
	return e.p.err
}

func (e *storeToSqlGen) gStruct(s *Struct, x *extra) {
	if !e.p.ok() {
		return
	}

	e.structmap(s)

	return
}

func (e *storeToSqlGen) appendraw(bts []byte) {
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

func (e *storeToSqlGen) structmap(s *Struct) {
	//nfields := len(s.Fields) - s.SkipCount

	recv := s.TypeName() // imutMethodReceiver(s)
	e.p.printf(`

// create table to store type '%s'
`, recv)
	e.p.printf("if create { sqlCreate = \"CREATE TABLE IF NOT EXISTS \" + dbName + \".\" + tableName + ` (\n")
	e.p.printf(`  rowid bigint AUTO_INCREMENT not null primary key
, updatetm TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
`)

	ins := "sqlIns = \"insert into \" + dbName + \".\" + tableName + \"("
	values := "" // the right number of ?,?,?,... question mark place-holders.
	var actuals, actuals2 string

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
			actuals2 += ","
		}
		fld = "”" + fld + "”"
		ins += fld
		values += "?"
		actuals += s.Varname() + "." + s.Fields[i].FieldName
		actuals2 += "r." + s.Fields[i].FieldName

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
			/*
				comments here https://mariadb.com/kb/en/timestamp/

				from 2018 by Dean Trower
				tell us how messed up TIMESTAMP handling is, and why we want to avoid it:
				"When using the TIMESTAMP type in timezone settings with daylight savings,
				you must be aware that: Storing data in the column will be ambiguous
				at the end of daylight savings, because the local-time hour gets repeated twice.

				At the start of daylight savings, when clocks are turned forwards,
				there's a "missing hour". Technically times in this hour aren't valid,
				but if you do try to store a time from this missing hour in a TIMESTAMP
				column and then read it back, you WON'T get out what you put in!
				(Contrary to what the MySQL docs state)."

				from 2016, Andy Walker says:
				"It's important to note that MariaDB (and MySQL) both have the notion
				of a server timezone - when writing an explicit value to a timestamp column,
				time data is converted from the server timezone into UTC for storage, then
				converted back to server time when reading.

				If you're planning on using the timestamp data type to hold timestamps
				which already in UTC, it's probably worth setting your server timezone
				to UTC. Alternatively, pay the size penalty and use the DATETIME type,
				which isn't affected by timezones."
			*/
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
  if db != nil {
    _, err = db.Exec(sqlCreate)
    if err != nil {
      err = fmt.Errorf("error creating table: '%%v'; sql was: '%%v'", err, sqlCreate)
      return
    }
  }
} // end if create

`)

	values = "(" + values + ")"
	ins += ") values " + values + "\"\n"
	e.p.printf(ins)
	e.p.printf("sqlIns = strings.ReplaceAll(sqlIns, \"”\", \"`\")")

	e.p.printf(`
    if db == nil {
        return
    }
`)

	e.p.printf(`

	var vals []any
	znil := 1 // allow z to be nil
	if z != nil {
		znil = 0
		vals = []any{%v}
	}

	var res sql.Result

	if len(rest) > 0 {
		// do multi-statement instead of prepared stmt
		sqlIns += strings.Repeat(",%v", len(rest)-znil)
		for _, r := range rest {
			vals = append(vals, %v)
		}
		// use a tmpStmt because the length of the batch will
		// rarely be repeated; this Stmt is not re-usable.
		var tmpStmt *sql.Stmt
		tmpStmt, err = tx.Prepare(sqlIns)
		if err != nil {
			err = fmt.Errorf("error preparing multi-insert: '%%v'; sql was: '%%v'", err, sqlIns)
			return
		}
		defer tmpStmt.Close()
		res, err = tmpStmt.Exec(vals...)
		if err != nil {
			err = fmt.Errorf("error on tmpStmt.Exec insert: '%%v'; sql was: '%%v'", err, sqlIns)
			return
		}

	} else {
		if stmt == nil {
			stmt, err = db.Prepare(sqlIns)
			if err != nil {
				err = fmt.Errorf("error preparing insert: '%%v'; sql was: '%%v'", err, sqlIns)
				return
			}
		}
		res, err = stmt.Exec(vals...)
		if err != nil {
			err = fmt.Errorf("error on stmt.Exec(): '%%v'; sql was: '%%v'", err, sqlIns)
			return
		}
	}

	injectedRowID, err = res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("error getting res.LastInsertId(): '%%v'", err)
		return
	}
	err = tx.Commit()

`, actuals, values, actuals2)

}

func (e *storeToSqlGen) gMap(m *Map, x *extra) {
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

func (e *storeToSqlGen) gPtr(s *Ptr, x *extra) {
	return
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	e.p.print("\n // gPtr.storeToSqlGen():\n")
	e.p.printf("\nif %s == nil { err = en.WriteNil(); if err != nil { return; } } else {", s.Varname())

	next(e, s.Value, nil)
	e.p.closeblock()
}

func (e *storeToSqlGen) gSlice(s *Slice, x *extra) {
	return
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	e.writeAndCheck(arrayHeader, lenAsUint32, s.Varname())
	e.p.rangeBlock(s.Index, s.Varname(), e, s.Els)
}

func (e *storeToSqlGen) gArray(a *Array, x *extra) {
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

func (e *storeToSqlGen) gBase(b *BaseElem, x *extra) {
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
		e.p.printf("\n // storeToSqlGen.gBase unknown IDENT '%v': ignore for now.\n", vname)
		e.p.printf("\n // err = %s.%sStoreToSQL(db, tableName)", vname, e.cfg.MethodPrefix)
	} else { // typical case
		e.writeAndCheck(b.BaseName(), literalFmt, vname)
	}
}
