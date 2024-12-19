// greenpack is a code generation tool for
// creating methods to serialize and de-serialize
// Go data structures to and from Greenpack (a
// schema-based serialization format that is derived
// from MessagePack2).
//
// This package is targeted at the `go generate` tool.
// To use it, include the following directive in a
// go source file with types requiring source generation:
//
//	//go:generate greenpack
//
// The go generate tool should set the proper environment variables for
// the generator to execute without any command-line flags. However, the
// following options are supported, if you need them (See greenpack -h):
//
//	  $ greenpack -h
//
//	  Usage of greenpack:
//
//	  -fast-strings
//	    	for speed when reading a string in a message that won't be
//	     reused, this flag means we'll use unsafe to cast the string
//	     header and avoid allocation.
//
//	  -file go generate
//	    	input file (or directory); default is $GOFILE, which
//	     is set by the go generate command.
//
//	  -genid
//	    	generate a fresh random greenSchemaId64 value to
//	     include in your Go source schema
//
//	  -io
//	    	create Encode and Decode methods (default true)
//
//	  -marshal
//	    	create Marshal and Unmarshal methods (default true)
//
//	 -method-prefix string
//	   	(optional) prefix that will be pre-prended to
//	     the front of generated method names; useful when
//	     you need to avoid namespace collisions, but the
//	     generated tests will break/the msgp package
//	     interfaces won't be satisfied.
//
//	 -no-embedded-schema
//	     don't embed the schema in the generated files
//
//	 -no-structnames-onwire
//	   	don't embed the name of the struct in the
//	     serialized greenpack. Skipping the embedded
//	     struct names saves time and space and matches
//	     what protocol buffers/thrift/capnproto/msgpack do.
//	     You must know the type on the wire you expect;
//	     or embed a type tag in one universal wrapper
//	     struct. Embedded struct names are a feature
//	     of Greenpack to help with dynamic language
//	     bindings.
//
//	  -o string
//	    	output file (default is {input_file}_gen.go
//
//	  -schema-to-go string
//	    	(standalone functionality) path to schema in msgpack2
//	     format; we will convert it to Go, write the Go on stdout,
//	     and exit immediately
//
//	  -tests
//	    	create tests and benchmarks (default true)
//
//	  -unexported
//	    	also process unexported types
//
//	  -write-schema string
//			write schema header to this file; - for stdout
//
// For more information, please read README.md, and the wiki at github.com/glycerine/greenpack
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/glycerine/greenpack/cfg"
	"github.com/glycerine/greenpack/gen"
	"github.com/glycerine/greenpack/parse"
	"github.com/glycerine/greenpack/printer"
)

func main() {
	myflags := flag.NewFlagSet("greenpack", flag.ExitOnError)
	c := &cfg.GreenConfig{}
	c.DefineFlags(myflags)

	err := myflags.Parse(os.Args[1:])
	err = c.ValidateConfig()
	if err != nil {
		fmt.Printf("greenpack command line flag error: '%s'\n", err)
		os.Exit(1)
	}

	if c.ShowVersion {
		fmt.Println(GetCodeVersion(os.Args[0]))
		os.Exit(0)
	}

	// GOFILE is set by go generate
	if c.GoFile == "" {
		c.GoFile = os.Getenv("GOFILE")
		if c.GoFile == "" {
			fmt.Println("No file to parse.")
			os.Exit(1)
		}
	}
	gen.SetFilename(c.GoFile)

	var mode gen.Method
	if c.Encode {
		mode |= (gen.Encode | gen.Decode | gen.Size | gen.FieldsEmpty)
	}
	if c.Marshal {
		mode |= (gen.Marshal | gen.Unmarshal | gen.Size | gen.FieldsEmpty)
	}
	if c.Tests {
		mode |= gen.Test
	}
	if c.StoreToSQL != "" {
		mode |= gen.StoreToSQL
		mode |= gen.GetFromSQL
	}

	if mode&^gen.Test == 0 {
		fmt.Println("No methods to generate; -io=false && -marshal=false")
		os.Exit(1)
	}

	if err := Run(mode, c); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// Run writes all methods using the associated file or path, e.g.
//
//	err := msgp.Run("path/to/myfile.go", gen.Size|gen.Marshal|gen.Unmarshal|gen.Test, false)
func Run(mode gen.Method, c *cfg.GreenConfig) error {
	if mode&^gen.Test == 0 {
		return nil
	}
	fmt.Println("======== Greenpack Code Generator  =======")
	fmt.Printf(">>> Input: \"%s\"\n", c.GoFile)
	var fs *parse.FileSet
	var err error
	//if c.NoLoad {
	fs, err = parse.FileNoLoad(c)
	//} else {
	//	fs, err = parse.File(c)
	//}
	if err != nil {
		return err
	}

	if len(fs.Identities) == 0 {
		fmt.Println("No types requiring code generation were found!")
		return nil
	}

	return printer.PrintFile(newFilename(c.Out, c.GoFile, fs.Package), fs, mode, c, c.GoFile)
}

// picks a new file name based on input flags and input filename(s).
func newFilename(out, old, pkg string) string {
	if out != "" {
		if pre := strings.TrimPrefix(out, old); len(pre) > 0 &&
			!strings.HasSuffix(out, ".go") {
			return filepath.Join(old, out)
		}
		return out
	}

	if fi, err := os.Stat(old); err == nil && fi.IsDir() {
		old = filepath.Join(old, pkg)
	}
	// new file name is old file name + _gen.go
	return strings.TrimSuffix(old, ".go") + "_gen.go"
}

func fileExists(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	}
	return true
}
