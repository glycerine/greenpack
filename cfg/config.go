package cfg

import (
	"flag"
)

type GreenConfig struct {
	Out        string
	GoFile     string
	Encode     bool
	Marshal    bool
	Tests      bool
	Unexported bool

	ReadStringsFast bool

	MethodPrefix string
	SerzEmpty    bool

	AllTuple    bool
	SkipZidClue bool
	Msgpack2    bool // -msgpack2 is an alias for -omit-clue

	ShowVersion           bool
	NoEmbeddedStructNames bool
	NoDedup               bool
}

// call DefineFlags before myflags.Parse()
func (c *GreenConfig) DefineFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.Out, "o", "", "output file (default is {input_file}_gen.go")
	fs.StringVar(&c.GoFile, "file", "", "input file (or directory); default is $GOFILE, which is set by the `go generate` command.")
	fs.BoolVar(&c.Encode, "io", true, "create Encode and Decode methods")
	fs.BoolVar(&c.Marshal, "marshal", true, "create Marshal and Unmarshal methods")
	fs.BoolVar(&c.Tests, "tests", true, "create tests and benchmarks")
	fs.BoolVar(&c.Unexported, "unexported", false, "also process unexported types")
	fs.BoolVar(&c.SerzEmpty, "write-zeros", false, "serialize zero-value fields to the wire, consuming much more space. By default all fields are treated as `omitempty` fields, where they are omitted from the serialization if they contain their zero-value. If -write-zero is given, then only fields specifically marked as `omitempty` are treated as such.")

	fs.BoolVar(&c.ReadStringsFast, "fast-strings", false, "for speed when reading a string in a message that won't be reused, this flag means we'll use unsafe to cast the string header and avoid allocation.")
	fs.StringVar(&c.MethodPrefix, "method-prefix", "", "(optional) prefix that will be pre-prended to the front of generated method names; useful when you need to avoid namespace collisions, but the generated tests will break/the msgp package interfaces won't be satisfied.")
	fs.BoolVar(&c.AllTuple, "alltuple", false, "use tuples for everything. Negates the point of greenpack, but useful in a pinch for performance. Provides no data versioning whatsoever. If you even so much as change the order of your fields, you won't be able to read back your earlier data correctly/without crashing.")
	fs.BoolVar(&c.SkipZidClue, "omit-clue", false, "don't append zid and clue to field name (makes things just like msgpack2 traditional encoding, without version + type clue)")
	fs.BoolVar(&c.Msgpack2, "msgpack2", false, "(alias for -omit-clue) don't append zid and clue to field name (makes things just like msgpack2 traditional encoding, without version + type clue)")
	fs.BoolVar(&c.ShowVersion, "version", false, "print version info and exit")
	fs.BoolVar(&c.NoEmbeddedStructNames, "no-structnames-onwire", false, "don't embed the name of the struct in the serialized greenpack. Skipping the embedded struct names saves time and space and matches what protocol buffers/thrift/capnproto/msgpack do. You must know the type on the wire you expect; or embed a type tag in one universal wrapper struct. Embedded struct names are a feature of GreenPack to help with dynamic language bindings and unmarshalling interface types. Given an interface, without the concrete type, we don't know what to unmarshal.")
	fs.BoolVar(&c.NoDedup, "no-dedup", false, "don't de-duplicate pointers within a struct")
}

// call c.ValidateConfig() after myflags.Parse()
func (c *GreenConfig) ValidateConfig() error {

	if c.Msgpack2 {
		c.SkipZidClue = true
	}

	return nil
}
