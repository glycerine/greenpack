// Code generated by GREENPACK (github.com/glycerine/greenpack). DO NOT EDIT.

package testdata

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/glycerine/greenpack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *StarShipFireAnt) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields0zgensym_4d6c949d38038a8c_1 = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zgensym_4d6c949d38038a8c_1 uint32
	totalEncodedFields0zgensym_4d6c949d38038a8c_1, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zgensym_4d6c949d38038a8c_1 := totalEncodedFields0zgensym_4d6c949d38038a8c_1
	missingFieldsLeft0zgensym_4d6c949d38038a8c_1 := maxFields0zgensym_4d6c949d38038a8c_1 - totalEncodedFields0zgensym_4d6c949d38038a8c_1

	var nextMiss0zgensym_4d6c949d38038a8c_1 int32 = -1
	var found0zgensym_4d6c949d38038a8c_1 [maxFields0zgensym_4d6c949d38038a8c_1]bool
	var curField0zgensym_4d6c949d38038a8c_1 string

doneWithStruct0zgensym_4d6c949d38038a8c_1:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zgensym_4d6c949d38038a8c_1 > 0 || missingFieldsLeft0zgensym_4d6c949d38038a8c_1 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zgensym_4d6c949d38038a8c_1, missingFieldsLeft0zgensym_4d6c949d38038a8c_1, msgp.ShowFound(found0zgensym_4d6c949d38038a8c_1[:]), decodeMsgFieldOrder0zgensym_4d6c949d38038a8c_1)
		if encodedFieldsLeft0zgensym_4d6c949d38038a8c_1 > 0 {
			encodedFieldsLeft0zgensym_4d6c949d38038a8c_1--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zgensym_4d6c949d38038a8c_1 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zgensym_4d6c949d38038a8c_1 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zgensym_4d6c949d38038a8c_1 = 0
			}
			for nextMiss0zgensym_4d6c949d38038a8c_1 < maxFields0zgensym_4d6c949d38038a8c_1 && (found0zgensym_4d6c949d38038a8c_1[nextMiss0zgensym_4d6c949d38038a8c_1] || decodeMsgFieldSkip0zgensym_4d6c949d38038a8c_1[nextMiss0zgensym_4d6c949d38038a8c_1]) {
				nextMiss0zgensym_4d6c949d38038a8c_1++
			}
			if nextMiss0zgensym_4d6c949d38038a8c_1 == maxFields0zgensym_4d6c949d38038a8c_1 {
				// filled all the empty fields!
				break doneWithStruct0zgensym_4d6c949d38038a8c_1
			}
			missingFieldsLeft0zgensym_4d6c949d38038a8c_1--
			curField0zgensym_4d6c949d38038a8c_1 = decodeMsgFieldOrder0zgensym_4d6c949d38038a8c_1[nextMiss0zgensym_4d6c949d38038a8c_1]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zgensym_4d6c949d38038a8c_1)
		switch curField0zgensym_4d6c949d38038a8c_1 {
		// -- templateDecodeMsg ends here --

		case "captain_zid00_str":
			found0zgensym_4d6c949d38038a8c_1[0] = true
			z.Captain, err = dc.ReadString()
			if err != nil {
				return
			}
		case "cargo_zid01_f64":
			found0zgensym_4d6c949d38038a8c_1[1] = true
			z.CargoAreaMetersSquared, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "shuttles_zid02_int":
			found0zgensym_4d6c949d38038a8c_1[2] = true
			z.Shuttles, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "rawBytesData_zid03_bin":
			found0zgensym_4d6c949d38038a8c_1[3] = true
			z.RawBytesData, err = dc.ReadBytes(z.RawBytesData)
			if err != nil {
				return
			}
		case "lastMessageTime_zid04_tim":
			found0zgensym_4d6c949d38038a8c_1[4] = true
			z.LastMessageTime, err = dc.ReadTime()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zgensym_4d6c949d38038a8c_1 != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of StarShipFireAnt
var decodeMsgFieldOrder0zgensym_4d6c949d38038a8c_1 = []string{"captain_zid00_str", "cargo_zid01_f64", "shuttles_zid02_int", "rawBytesData_zid03_bin", "lastMessageTime_zid04_tim"}

var decodeMsgFieldSkip0zgensym_4d6c949d38038a8c_1 = []bool{false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *StarShipFireAnt) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[0] = (len(z.Captain) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.CargoAreaMetersSquared == 0) // number, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.Shuttles == 0) // number, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (len(z.RawBytesData) == 0) // string, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (z.LastMessageTime.IsZero()) // time.Time, omitempty
	if isempty[4] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *StarShipFireAnt) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zgensym_4d6c949d38038a8c_2 [5]bool
	fieldsInUse_zgensym_4d6c949d38038a8c_3 := z.fieldsNotEmpty(empty_zgensym_4d6c949d38038a8c_2[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_4d6c949d38038a8c_3 + 1)
	if err != nil {
		return err
	}

	// runtime struct type identification for 'StarShipFireAnt'
	err = en.Append(0xa1, 0x40)
	if err != nil {
		return err
	}
	err = en.WriteStringFromBytes([]byte{0x53, 0x74, 0x61, 0x72, 0x53, 0x68, 0x69, 0x70, 0x46, 0x69, 0x72, 0x65, 0x41, 0x6e, 0x74})
	if err != nil {
		return err
	}

	if !empty_zgensym_4d6c949d38038a8c_2[0] {
		// write "captain_zid00_str"
		err = en.Append(0xb1, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Captain)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_4d6c949d38038a8c_2[1] {
		// write "cargo_zid01_f64"
		err = en.Append(0xaf, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x66, 0x36, 0x34)
		if err != nil {
			return err
		}
		err = en.WriteFloat64(z.CargoAreaMetersSquared)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_4d6c949d38038a8c_2[2] {
		// write "shuttles_zid02_int"
		err = en.Append(0xb2, 0x73, 0x68, 0x75, 0x74, 0x74, 0x6c, 0x65, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x69, 0x6e, 0x74)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Shuttles)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_4d6c949d38038a8c_2[3] {
		// write "rawBytesData_zid03_bin"
		err = en.Append(0xb6, 0x72, 0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x62, 0x69, 0x6e)
		if err != nil {
			return err
		}
		err = en.WriteBytes(z.RawBytesData)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_4d6c949d38038a8c_2[4] {
		// write "lastMessageTime_zid04_tim"
		err = en.Append(0xb9, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x74, 0x69, 0x6d)
		if err != nil {
			return err
		}
		err = en.WriteTime(z.LastMessageTime)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *StarShipFireAnt) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [5]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "captain_zid00_str"
		o = append(o, 0xb1, 0x63, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.Captain)
	}

	if !empty[1] {
		// string "cargo_zid01_f64"
		o = append(o, 0xaf, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x66, 0x36, 0x34)
		o = msgp.AppendFloat64(o, z.CargoAreaMetersSquared)
	}

	if !empty[2] {
		// string "shuttles_zid02_int"
		o = append(o, 0xb2, 0x73, 0x68, 0x75, 0x74, 0x74, 0x6c, 0x65, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x69, 0x6e, 0x74)
		o = msgp.AppendInt(o, z.Shuttles)
	}

	if !empty[3] {
		// string "rawBytesData_zid03_bin"
		o = append(o, 0xb6, 0x72, 0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x62, 0x69, 0x6e)
		o = msgp.AppendBytes(o, z.RawBytesData)
	}

	if !empty[4] {
		// string "lastMessageTime_zid04_tim"
		o = append(o, 0xb9, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x74, 0x69, 0x6d)
		o = msgp.AppendTime(o, z.LastMessageTime)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StarShipFireAnt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *StarShipFireAnt) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields4zgensym_4d6c949d38038a8c_5 = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zgensym_4d6c949d38038a8c_5 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zgensym_4d6c949d38038a8c_5, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zgensym_4d6c949d38038a8c_5 := totalEncodedFields4zgensym_4d6c949d38038a8c_5
	missingFieldsLeft4zgensym_4d6c949d38038a8c_5 := maxFields4zgensym_4d6c949d38038a8c_5 - totalEncodedFields4zgensym_4d6c949d38038a8c_5

	var nextMiss4zgensym_4d6c949d38038a8c_5 int32 = -1
	var found4zgensym_4d6c949d38038a8c_5 [maxFields4zgensym_4d6c949d38038a8c_5]bool
	var curField4zgensym_4d6c949d38038a8c_5 string

doneWithStruct4zgensym_4d6c949d38038a8c_5:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zgensym_4d6c949d38038a8c_5 > 0 || missingFieldsLeft4zgensym_4d6c949d38038a8c_5 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zgensym_4d6c949d38038a8c_5, missingFieldsLeft4zgensym_4d6c949d38038a8c_5, msgp.ShowFound(found4zgensym_4d6c949d38038a8c_5[:]), unmarshalMsgFieldOrder4zgensym_4d6c949d38038a8c_5)
		if encodedFieldsLeft4zgensym_4d6c949d38038a8c_5 > 0 {
			encodedFieldsLeft4zgensym_4d6c949d38038a8c_5--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zgensym_4d6c949d38038a8c_5 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zgensym_4d6c949d38038a8c_5 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zgensym_4d6c949d38038a8c_5 = 0
			}
			for nextMiss4zgensym_4d6c949d38038a8c_5 < maxFields4zgensym_4d6c949d38038a8c_5 && (found4zgensym_4d6c949d38038a8c_5[nextMiss4zgensym_4d6c949d38038a8c_5] || unmarshalMsgFieldSkip4zgensym_4d6c949d38038a8c_5[nextMiss4zgensym_4d6c949d38038a8c_5]) {
				nextMiss4zgensym_4d6c949d38038a8c_5++
			}
			if nextMiss4zgensym_4d6c949d38038a8c_5 == maxFields4zgensym_4d6c949d38038a8c_5 {
				// filled all the empty fields!
				break doneWithStruct4zgensym_4d6c949d38038a8c_5
			}
			missingFieldsLeft4zgensym_4d6c949d38038a8c_5--
			curField4zgensym_4d6c949d38038a8c_5 = unmarshalMsgFieldOrder4zgensym_4d6c949d38038a8c_5[nextMiss4zgensym_4d6c949d38038a8c_5]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zgensym_4d6c949d38038a8c_5)
		switch curField4zgensym_4d6c949d38038a8c_5 {
		// -- templateUnmarshalMsg ends here --

		case "captain_zid00_str":
			found4zgensym_4d6c949d38038a8c_5[0] = true
			z.Captain, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "cargo_zid01_f64":
			found4zgensym_4d6c949d38038a8c_5[1] = true
			z.CargoAreaMetersSquared, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "shuttles_zid02_int":
			found4zgensym_4d6c949d38038a8c_5[2] = true
			z.Shuttles, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "rawBytesData_zid03_bin":
			found4zgensym_4d6c949d38038a8c_5[3] = true
			if nbs.AlwaysNil || msgp.IsNil(bts) {
				if !nbs.AlwaysNil {
					bts = bts[1:]
				}
				z.RawBytesData = z.RawBytesData[:0]
			} else {
				z.RawBytesData, bts, err = nbs.ReadBytesBytes(bts, z.RawBytesData)

				if err != nil {
					return
				}
			}
			if err != nil {
				return
			}
		case "lastMessageTime_zid04_tim":
			found4zgensym_4d6c949d38038a8c_5[4] = true
			z.LastMessageTime, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss4zgensym_4d6c949d38038a8c_5 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of StarShipFireAnt
var unmarshalMsgFieldOrder4zgensym_4d6c949d38038a8c_5 = []string{"captain_zid00_str", "cargo_zid01_f64", "shuttles_zid02_int", "rawBytesData_zid03_bin", "lastMessageTime_zid04_tim"}

var unmarshalMsgFieldSkip4zgensym_4d6c949d38038a8c_5 = []bool{false, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StarShipFireAnt) Msgsize() (s int) {
	s = 1 + 18 + msgp.StringPrefixSize + len(z.Captain) + 16 + msgp.Float64Size + 19 + msgp.IntSize + 23 + msgp.BytesPrefixSize + len(z.RawBytesData) + 26 + msgp.TimeSize
	return
}

// StoreTogSQL stores z into the table specified by tableName,
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
// if create was true.
func (z *StarShipFireAnt) StoreToSQL(db *sql.DB, dbName, tableName string, create bool, reuseStmt *sql.Stmt) (stmt *sql.Stmt, injectedRowID int64, sqlIns, sqlCreate string, err error) {
	stmt = reuseStmt

	// create table to store type 'StarShipFireAnt'
	if create {
		sqlCreate = "CREATE TABLE IF NOT EXISTS " + dbName + "." + tableName + ` (
  rowid bigint AUTO_INCREMENT not null primary key
, updatetm TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
, ”captain” LONGTEXT
, ”cargo” DOUBLE
, ”shuttles” BIGINT
, ”rawBytesData” LONGBLOB
, ”lastMessageTime” DATETIME(6)

)`
		// mariaDB needs backtick quoted strings
		sqlCreate = strings.ReplaceAll(sqlCreate, "”", "`")
		if db != nil {
			_, err = db.Exec(sqlCreate)
			if err != nil {
				err = fmt.Errorf("error creating table: '%v'; sql was: '%v'", err, sqlCreate)
				return
			}
		}
	} // end if create

	if stmt == nil {
		sqlIns = "insert into " + dbName + "." + tableName + "(”captain”, ”cargo”, ”shuttles”, ”rawBytesData”, ”lastMessageTime”) values (?,?,?,?,?)"
		sqlIns = strings.ReplaceAll(sqlIns, "”", "`")
		if db != nil {
			stmt, err = db.Prepare(sqlIns)
			if err != nil {
				err = fmt.Errorf("error preparing insert: '%v'; sql was: '%v'", err, sqlIns)
				return
			}
		}
	}
	if db != nil {
		var res sql.Result
		res, err = stmt.Exec(z.Captain, z.CargoAreaMetersSquared, z.Shuttles, z.RawBytesData, z.LastMessageTime)
		if err != nil {
			return
		}
		injectedRowID, err = res.LastInsertId()
	}

	return
}

// Note that the github.com/go-sql-driver/mysql driver will need parseTime=true
// added to the DSN string in order to pull DATETIME from sql into time.Time in Go.
// For example:
// db, err := sql.Open("mysql", "user:pw@tcp(localhost:3306)/databasename?parseTime=true")
//
// Also storing time.Time from Go into MariaDB will be lossy because Go keeps
// nanoseconds but MariaDB only stores microseconds.
//
// Beware that ctx cancellation will close the db connection(!) This is
// very unexpected. Likely this is the only way it can be implemented.
//
// From the github.com/go-sql-driver/mysql README docs:
//
// "The QueryContext, ExecContext, etc. variants provided by database/sql will
//  cause the connection to be closed if the provided context is
//  cancelled or timed out before the result is received by the driver"

// GetFromSQL fetches from the database and fills in z (this method's receiver)
// and returns a variable sized slice in res. res[0] will always be z.
//
// The where clause can be the empty string. It can also be
// used to get a subset of the contents of tableName back.
//
// With an empty where, all records in tableName
// are selected and returned. Otherwise, the where string
// is appended to a select for all records.
//
// Advanced use:
// If the db is nil, the sqlSel will be returned but the
// database will not be contacted. This allows you get the
// query that has all the fields, and then
// modify the select query and re-submit it with complex
// joins, etc.
//
// Hence if the where clause starts with "select",
// then it will be used directly; it will not be appended
// to a select of all the columns. Be careful with this,
// however, since the rows.Scan() call generated below
// expects to get all the fields, so you must keep
// the select statement in sync with those expectations.
func (z *StarShipFireAnt) GetFromSQL(ctx context.Context, db *sql.DB, dbName, tableName, where string) (res []*StarShipFireAnt, sqlSel string, err error) {

	if strings.HasPrefix(where, "select") {
		sqlSel = where
	} else {

		// get from sql and write into 'StarShipFireAnt'
		sqlSel = "select ”captain”, ”cargo”, ”shuttles”, ”rawBytesData”, ”lastMessageTime” from " + dbName + "." + tableName
		sqlSel = strings.ReplaceAll(sqlSel, "”", "`") + " " + where
	}

	if db == nil {
		return // just providing the sqlSel statement.
	}

	var rows *sql.Rows
	rows, err = db.QueryContext(ctx, sqlSel)
	if err != nil {
		err = fmt.Errorf("error on select: '%v'; sql was '%v'", err, sqlSel)
		return
	}
	defer rows.Close()

	var v *StarShipFireAnt
	for i := 0; rows.Next(); i++ {
		if i == 0 {
			v = z
		} else {
			v = &StarShipFireAnt{}
		}
		err = rows.Scan(&v.Captain, &v.CargoAreaMetersSquared, &v.Shuttles, &v.RawBytesData, &v.LastMessageTime)
		if err != nil {
			return
		}
		res = append(res, v)
	}

	return
}
