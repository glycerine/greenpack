package zebra

//go:generate msgp

// Zprimitive describes the basic type category of the field
// It should match gen/Primitive
type Zprimitive uint8

const (
	Invalid Zprimitive = iota
	Bytes
	String
	Float32
	Float64
	Complex64
	Complex128
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Byte
	Int
	Int8
	Int16
	Int32
	Int64
	Bool
	Intf // interface{}
	Time // time.Time
	Ext  // extension

	IDENT // IDENT means an unrecognized identifier
)

// ZebraSchema is the top level container
type Schema struct {
	PkgPath string // reflect.TypeOf().PkgPath()

	Structs []Struct
}

// Struct represents a single message or struct.
type Struct struct {
	StructName string  // name of struct
	Fld        []Field // fields
}

// Field represents fields within a struct.
type Field struct {

	// Zid locates update collisions and ease resolution.
	//
	// Both follow Cap'nProto semantics: start at numbering at 0,
	// and strictly/always increase numbers monotically.
	//
	// No gaps and no duplicate Zid are allowed.
	//
	// Duplicate numbers are how collisions (between two
	// developers adding two distinct new fields independently
	// but at the same time) are detected.
	//
	// Therefore this ironclad rule: never delete a field or Zid number,
	// just mark it as deprecated with the `deprecated:"true"`
	// tag, and change its Go type to struct{}.
	//
	Zid       int64
	Nam       string
	TypStr    string
	OmitEmpty bool
	Skip      bool
	Tag       map[string]string `msg:",omitempty"`
}