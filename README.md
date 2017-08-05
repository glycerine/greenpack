greenpack: serialization that extends msgpack2 with field versioning and strong typing
==========

`greenpack` keeps your data completely self-contained. No coordination with a (long lost?) schema is needed to decode.

`greenpack` is a data definition language and serialization format, a layer of convention on top of msgpack2 that provides field versioning and additional type safety.

`msgpack2` [https://github.com/msgpack/msgpack/blob/master/spec.md] [http://msgpack.org] enjoys wide cross-language support, and provides efficient and self-contained data serialization. We find only two problems with msgpack2: weak support for data evolution, and insufficiently strong typing of integers.

The greenpack format addresses these problems while keeping serialized data fully self-describing. Greenpack is independent of any external schema, but as an optimization uses the Go source file itself as a schema to maintain current versioning and type information. Dynamic languages still have an easy time reading greenpack--it is just msgpack2. There's no need to worry about coordinating the schema under which data was written, as data is self-contained.


# the main idea

```
//given this definition, defined in Go:
type A struct {
  Name     string      `zid:"0"`
  Bday     time.Time   `zid:"1"`
  Phone    string      `zid:"2"`
  Sibs     int         `zid:"3"`
  GPA      float64     `zid:"4"`
  Friend   bool        `zid:"5"`
}

then when greenpack serializes, the it looks like msgpack2 on the wire with extended field names:
         
greenpack
--------              
a := A{               
  "Name_zid00_str"  :  "Atlanta",
  "Bday_zid01_tim"  :  tm("1990-12-20"),
  "Phone_zid02_str" :  "650-555-1212",  
  "Sibs_zid03_i64"  :  3,               
  "GPA_zid04_f64"   :  3.95,            
  "Friend_zid05_boo":  true,            
}

```

Notice the only thing that changed with respect to the msgpack2 encoding is that the the fieldnames have been extended to contain a version and a type clue.

The central idea of greenpack: start with msgpack2, and append version numbers and type clues to the end of the field names when stored on the wire. We say type "clues" because the type information clarifies the original size and signed-ness of the type, which adds the missing detail to integers needed to fully reconstruct the original data from the serialization. This address the problem that commonly msgpack2 implementations ignore the spec and encode numbers using the smallest unsigned type possible, which corrupts the original type information and can induce decoding errors for large and negative numbers.

If you've ever had your msgpack crash your server because you tried to change the type of a field but keep the same name, then you know how fragile msgpack can be. The type clue fixes that.

The version `zid` number gives us the ability to evolve our data without crashes. The moniker `zid` reveals `greenpacks` evolution from `zebrapack`, where it stood for "zebrapack version id". Rather than rework all the tooling to expect `gid`, which might be confused with a `GUID`, we simply keep the convention. `zid` indicates the field version.

The second easy idea: use the Go language struct definition syntax as our serialization schema. There is no need to invent a completely different format. Serialization for Go developers should be almost trivially easy. While we are focused on a serialization format for Go, because other language can read msgpack2, they can also readily parse the schema. While the schema is optional, greenpack (this repo) provides code generation tools based on the schema (Go file) that generates extremely fast serialization code.

# the need for stronger integer typing

Starting point: [msgpack2](http://msgpack.org) is great.
It is has an easy to read spec, it defines a compact
serialization format, and it has wide language support from
both dynamic and compiled languages.

Nonetheless, data update
conflicts still happen and can be hard to
resolve. Encoders could use the guidance from
type clues to avoid signed versus unsigned integer
encodings.

For instance, sadly the widely emulated C-encoder
for msgpack chooses to encode signed positive integers
as unsigned integers. This causes crashes in readers
who were expected a signed integer, which they may
have originated themselves in the original struct.

Astonishing, but true: the existing practice for msgpack2
language bindings allows the data types to change as
they are read and re-serialized. Simple copying of
a serialized struct can change the types of data
from signed to unsigned. This is horrible. Now we have to guess
whether an unsigned integer was really intended because
of the integer's range, or if data will be silently
truncated or lost when coercing a 64-bit integer to
a 63-bit signed integer--assuming such coercing ever
makes logical sense, which it may not.

This kind of tragedy happens because of a lack of
shared communication across time and space between
readers and writers. It is easily addressed with
type clues, small extra information about the
originally defined type.

# field version info, using the `zid` tag.

* Conflict resolution: the Cap'nProto numbering and
update conflict resolution method is used here.
This method originated in the ProtocolBuffers
scheme, and was enhanced after experience in
Cap'nProto. How it works: Additions are always
made by incrementing by one the largest number available
prior to the addition. No gaps in numbering are
allowed, and no numbers are ever deleted.
To get the effect of deletion, add the `deprecated` value
in `msg` tag. This is an effective tombstone.
It allows the tools (the `go` compiler and the
`greenpack` code generator) to help detect
merge conflicts as soon as possible. If
two people try to merge schemas where the same
struct or field number is re-used, then
when `greenpack` is run to regenerate the
serialization code (under `go generate`),
it will automatically detect the conflict,
and flag the human to resolve the conflict
before proceeding.

* All fields optional. Just as in msgpack2,
Cap'nProto, Gobs, and Flatbuffers, all fields
are optional. Most everyone, after experience
and time with ProtocolBuffers, has come to the
conclusion that required fields are a misfeature
that hurt the ability to evolve data gracefully
and maintain efficiency.

Design:

* Schema language: the schema language for
defining structs is identical to the Go
language. Go is expressive and yet easily parsed
by the standard library packages included
with Go itself.

* Requirement: greenpack requires that the msgpack2 standard
be adhered to. Strings and raw binary byte arrays
are distinct, and must be marked distinctly; msgpack1 encoding is
not allowed.

* All language bindings must respect the declared type in
the type clue when writing data. For example,
this means that signed and unsigned declarations
must be respected. Even if another language uses
a msgpack2 implimentation that converts signed to
unsigned, as long as the field name is preserved
we can still acurately reconstruct what the
data's type was originally.

performance and comparison
=========================

`greenpack -fast-strings` is zero-allocation, and one
of the fastest serialization formats avaiable for Go.[1]

[1] https://github.com/glycerine/go_serialization_benchmarks

For write speed, only Zebrapack is faster. For
reads, only CapnProto and Gencode are slightly faster.
Gencode isn't zero alloc, and has no versioning support.
CapnProto isn't very portable to dynamic languages
like R or Javascript; Java support was never
finished. It requires keeping duplicate
mirror structs in your code. I like CapnProto and
maintained Go bindings for CapnProto for quite a
while. However the convenience of msgpack2 won
me over. Moreover CapnProto's layout format
is undocumented, it requires a C++ build chain to
build the IDL compiler, and unused fields always
take space on the wire. `greenpack` is pure Go,
and there are over 50 msgpack libraries -- one for every
language imaginable -- cited at http://msgpack.org.

Compared to (Gogoprotobuf) ProtcolBuffers, greenpack reads
are 6% faster on these microbenchmarks. Writes
are 15% faster and do no allocation; GogoprotobufMarshal
appears to allocate on write.


deprecating fields
------------------

to actually deprecate a field, you start by adding the `,deprecated` value to the `msg` tag key:
```
type A struct {
  Name     string      `zid:"0"`
  Bday     time.Time   `zid:"1"`
  Phone    string      `zid:"2"`
  Sibs     int         `zid:"3"`
  GPA      float64     `zid:"4" msg:",deprecated"` // a deprecated field.
  Friend   bool        `zid:"5"`
}
```
*In addition,* you'll want to change the type of the deprecated field, substituting `struct{}` for the old type. By converting the type of the deprecated field to struct{}, it will no longer takes up any space in the Go struct. This saves space. Even if a struct evolves heavily in time (rare), the changes will cause no extra overhead in terms of memory. It also allows the compiler to detect and reject any new writes to the field that are using the old type. 
```
// best practice for deprecation of fields, to save space + get compiler support for deprecation
type A struct {
  Name     string      `zid:"0"`
  Bday     time.Time   `zid:"1"`
  Phone    string      `zid:"2"`
  Sibs     int         `zid:"3"`
  GPA      struct{}    `zid:"4" msg:",deprecated"` // a deprecated field should have its type changed to struct{}, as well as being marked msg:",deprecated"
  Friend   bool        `zid:"5"`
}
```

Rules for safe data changes: To preserve forwards/backwards compatible changes, you must *never remove a field* from a struct, once that field has been defined and used. In the example above, the `zid:"4"` tag must stay in place, to prevent someone else from ever using 4 again. This allows sane data forward evolution, without tears, fears, or crashing of servers. The fact that `struct{}` fields take up no space also means that there is no need to worry about loss of performance when deprecating. We retain all fields ever used for their zebra ids, and the compiled Go code wastes no extra space for the deprecated fields.

NB: There is one exception to this `struct{}` consumes no space rule: if the newly deprecated `struct{}` field happens to be *the very last field* in a struct, it will take up one pointer worth of space. If you want to deprecate the last field in a struct, if possible you should move it up in the field order (e.g. make it the first field in the Go struct), so it doesn't still consume space; reference https://github.com/golang/go/issues/17450.


command line flags
------------------

~~~
  $ greenpack -h

Usage of greenpack:

  -alltuple
    	use tuples for everything. Negates the point
        of greenpack, but useful in a pinch for
        performance. Provides no data versioning
        whatsoever. If you even so much as change
        the order of your fields, you won't be
        able to read back your earlier data
        correctly/without crashing.
        
  -fast-strings
    	for speed when reading a string in
        a message that won't be reused, this
        flag means we'll use unsafe to cast
        the string header and avoid allocation.
        
  -file go generate
    	input file (or directory); default
        is $GOFILE, which is set by the
        go generate command.
        
  -io
    	create Encode and Decode methods (default true)
        
  -marshal
    	create Marshal and Unmarshal methods
        (default true)
        
  -method-prefix string
    	(optional) prefix that will be pre-prended
        to the front of generated method names;
        useful when you need to avoid namespace
        collisions, but the generated tests will
        break/the msgp package interfaces won't be satisfied.
        
  -o string
    	output file (default is {input_file}_gen.go
        
  -omit-clue
    	don't append zid and clue to field name
        (makes things just like msgpack2 traditional
        encoding, without version + type clue)
        
  -tests
    	create tests and benchmarks (default true)
        
  -unexported
    	also process unexported types
        
  -write-zeros
    	serialize zero-value fields to the wire,
        consuming much more space. By default
        all fields are treated as omitempty fields,
        where they are omitted from the
        serialization if they contain their zero-value.
        If -write-zero is given, then only fields
        specifically marked as `omitempty` are
        treated as such.
        
~~~

### `msg:",omitempty"` tags on struct fields

By default, all fields are treated as `omitempty`. If the
field contains its zero-value (see the Go spec), then it
is not serialized on the wire.

If you wish to consume space unnecessarily, you can
use the `greenpack -write-zeros` flag. Then only
fields specifically marked with the struct tag
`omitempty` will be treated as such.


For example, in the following example,
```
type Hedgehog struct {
   Furriness string `msg:",omitempty"`
}
```

If Furriness is the empty string, the field will not be serialized, thus saving the space of the field name on the wire. If the `-write-zeros` flags was given and the `omitempty` tag removed, then Furriness would be serialized no matter what value it contained.

It is safe to re-use structs by default, and with `omitempty`. For reference:

from https://github.com/tinylib/msgp/issues/154:
> The only special feature of UnmarshalMsg and DecodeMsg (from a zero-alloc standpoint) is that they will use pre-existing fields in an object rather than allocating new ones. So, if you decode into the same object repeatedly, things like slices and maps won't be re-allocated on each decode; instead, they will be re-sized appropriately. In other words, mutable fields are simply mutated in-place.

This continues to hold true, and missing fields on the wire will zero the field in any re-used struct.

NB: Under tuple encoding (https://github.com/tinylib/msgp/wiki/Preprocessor-Directives), for example `//msgp:tuple Hedgehog`, then all fields are always serialized and the omitempty tag is ignored.

## `addzid` utility

The `addzid` utility (in the cmd/addzid subdir) can help you
get started. Running `addzid mysource.go` on a .go source file
will add the `zid:"0"`... fields automatically. This makes adding greenpack
serialization to existing Go projects easy.
See https://github.com/glycerine/greenpack/blob/master/cmd/addzid/README.md
for more detail.

notices
-------

~~~
Portions Copyright (c) 2016, 2017 Jason E. Aten, Ph.D.
Portions Copyright (c) 2014 Philip Hofer
Portions Copyright (c) 2009 The Go Authors (license at http://golang.org) where indicated
~~~

LICENSE: MIT. See https://github.com/glycerine/greenpack/blob/master/LICENSE

ancestor codebase: tinylib/msgp
------------------

`greenpack` gets most of its speed by descending from the
fantastic and highly tuned https://github.com/tinylib/msgp library by
Philip Hofer. The special tag and shim handling is best documented
in the `msgp` writeup and wiki [https://github.com/tinylib/msgp/wiki].


# from the original https://github.com/tinylib/msgp README

MessagePack Code Generator [![Build Status](https://travis-ci.org/tinylib/msgp.svg?branch=master)](https://travis-ci.org/tinylib/msgp)
=======

This is a code generation tool and serialization library for [MessagePack](http://msgpack.org). You can read more about MessagePack [in the wiki](http://github.com/tinylib/msgp/wiki), or at [msgpack.org](http://msgpack.org).

### Why?

- Use Go as your schema language
- Performance
- [JSON interop](http://godoc.org/github.com/tinylib/msgp/msgp#CopyToJSON)
- [User-defined extensions](http://github.com/tinylib/msgp/wiki/Using-Extensions)
- Type safety
- Encoding flexibility

### Quickstart

In a source file, include the following directive:

```go
//go:generate greenpack
```

The `greenpack` command will generate serialization methods for all exported type declarations in the file. If you add the flag `-msgp`, it will generate msgpack2 rather than greenpack format.

For other language's use, schemas can can be written to a separate file using `greenpack -file my.go -write-schema` at the shell. (By default schemas are not written to the wire, just as in protobufs/CapnProto/Thrift.)

You can [read more about the code generation options here](http://github.com/tinylib/msgp/wiki/Using-the-Code-Generator).

### Use

Field names can be set in much the same way as the `encoding/json` package. For example:

```go
type Person struct {
	Name       string `msg:"name"`
	Address    string `msg:"address"`
	Age        int    `msg:"age"`
	Hidden     string `msg:"-"` // this field is ignored
	unexported bool             // this field is also ignored
}
```

By default, the code generator will satisfy `msgp.Sizer`, `msgp.Encodable`, `msgp.Decodable`, 
`msgp.Marshaler`, and `msgp.Unmarshaler`. Carefully-designed applications can use these methods to do
marshalling/unmarshalling with zero heap allocations.

While `msgp.Marshaler` and `msgp.Unmarshaler` are quite similar to the standard library's
`json.Marshaler` and `json.Unmarshaler`, `msgp.Encodable` and `msgp.Decodable` are useful for 
stream serialization. (`*msgp.Writer` and `*msgp.Reader` are essentially protocol-aware versions
of `*bufio.Writer` and `*bufio.Reader`, respectively.)

### Features

 - Extremely fast generated code
 - Test and benchmark generation
 - JSON interoperability (see `msgp.CopyToJSON() and msgp.UnmarshalAsJSON()`)
 - Support for complex type declarations
 - Native support for Go's `time.Time`, `complex64`, and `complex128` types 
 - Generation of both `[]byte`-oriented and `io.Reader/io.Writer`-oriented methods
 - Support for arbitrary type system extensions
 - [Preprocessor directives](http://github.com/tinylib/msgp/wiki/Preprocessor-Directives)
 - File-based dependency model means fast codegen regardless of source tree size.

Consider the following:
```go
const Eight = 8
type MyInt int
type Data []byte

type Struct struct {
	Which  map[string]*MyInt `msg:"which"`
	Other  Data              `msg:"other"`
	Nums   [Eight]float64    `msg:"nums"`
}
```
As long as the declarations of `MyInt` and `Data` are in the same file as `Struct`, the parser will determine that the type information for `MyInt` and `Data` can be passed into the definition of `Struct` before its methods are generated.

#### Extensions

MessagePack supports defining your own types through "extensions," which are just a tuple of
the data "type" (`int8`) and the raw binary. You [can see a worked example in the wiki.](http://github.com/tinylib/msgp/wiki/Using-Extensions)

### Status

Mostly stable, in that no breaking changes have been made to the `/msgp` library in more than a year. Newer versions
of the code may generate different code than older versions for performance reasons. I (@philhofer) am aware of a
number of stability-critical commercial applications that use this code with good results. But, caveat emptor.

You can read more about how `msgp` maps MessagePack types onto Go types [in the wiki](http://github.com/tinylib/msgp/wiki).

Here some of the known limitations/restrictions:

- Identifiers from outside the processed source file are assumed (optimistically) to satisfy the generator's interfaces. If this isn't the case, your code will fail to compile.
- Like most serializers, `chan` and `func` fields are ignored, as well as non-exported fields.
- Encoding of `interface{}` is limited to built-ins or types that have explicit encoding methods.


If the output compiles, then there's a pretty good chance things are fine. (Plus, we generate tests for you.) *Please, please, please* file an issue if you think the generator is writing broken code.

### Performance

If you like benchmarks, see [here](http://bravenewgeek.com/so-you-wanna-go-fast/) and above in the greenpack benchmarks section; [see here for the benchmark source code](https://github.com/glycerine/go_serialization_benchmarks).

As one might expect, the generated methods that deal with `[]byte` are faster for small objects, but the `io.Reader/Writer` methods are generally more memory-efficient (and, at some point, faster) for large (> 2KB) objects.
