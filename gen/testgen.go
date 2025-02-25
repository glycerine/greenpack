package gen

import (
	"github.com/glycerine/greenpack/cfg"
	"io"
	"text/template"
)

var (
	marshalTestTempl = template.New("MarshalTest")
	encodeTestTempl  = template.New("EncodeTest")
)

// TODO(philhofer):
// for simplicity's sake, right now
// we can only generate tests for types
// that can be initialized with the
// "Type{}" syntax.
// we should support all the types.

func mtest(w io.Writer, cfg *cfg.GreenConfig) *mtestGen {
	return &mtestGen{w: w, cfg: cfg}
}

type mtestGen struct {
	passes
	w   io.Writer
	cfg *cfg.GreenConfig
}

func (m *mtestGen) MethodPrefix() string {
	return m.cfg.MethodPrefix
}

func (m *mtestGen) Execute(p Elem) error {
	p = m.applyall(p)
	if p != nil && IsPrintable(p) {
		switch p.(type) {
		case *Struct, *Array, *Slice, *Map:
			p.SetHasMethodPrefix(m)
			return marshalTestTempl.Execute(m.w, p)
		}
	}
	return nil
}

func (m *mtestGen) Method() Method { return marshaltest }

type etestGen struct {
	passes
	w   io.Writer
	cfg *cfg.GreenConfig
}

func etest(w io.Writer, cfg *cfg.GreenConfig) *etestGen {
	return &etestGen{w: w, cfg: cfg}
}

func (e *etestGen) MethodPrefix() string {
	return e.cfg.MethodPrefix
}

func (e *etestGen) Execute(p Elem) error {
	p = e.applyall(p)
	if p != nil && IsPrintable(p) {
		switch p.(type) {
		case *Struct, *Array, *Slice, *Map:
			p.SetHasMethodPrefix(e)
			return encodeTestTempl.Execute(e.w, p)
		}
	}
	return nil
}

func (e *etestGen) Method() Method { return encodetest }

func init() {
	template.Must(marshalTestTempl.Parse(`func TestMarshalUnmarshal{{.TypeName}}{{.MethodPrefix}}(t *testing.T) {
	v := {{.TypeName}}{{.GenericInstantiation}}{}
	bts, err := v.{{.MethodPrefix}}MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}
	left, err := v.{{.MethodPrefix}}UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after {{.MethodPrefix}}UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func BenchmarkMarshalMsg{{.TypeName}}{{.MethodPrefix}}(b *testing.B) {
	v := {{.TypeName}}{{.GenericInstantiation}}{}
	b.ReportAllocs()
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		v.{{.MethodPrefix}}MarshalMsg(nil)
	}
}

func BenchmarkAppendMsg{{.TypeName}}{{.MethodPrefix}}(b *testing.B) {
	v := {{.TypeName}}{{.GenericInstantiation}}{}
	bts := make([]byte, 0, v.{{.MethodPrefix}}Msgsize())
	bts, _ = v.{{.MethodPrefix}}MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		bts, _ = v.{{.MethodPrefix}}MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshal{{.TypeName}}{{.MethodPrefix}}(b *testing.B) {
	v := {{.TypeName}}{{.GenericInstantiation}}{}
	bts, _ := v.{{.MethodPrefix}}MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_, err := v.{{.MethodPrefix}}UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

`))

	template.Must(encodeTestTempl.Parse(`func TestEncodeDecode{{.TypeName}}{{.MethodPrefix}}(t *testing.T) {
	v := {{.TypeName}}{{.GenericInstantiation}}{}
	var buf bytes.Buffer
	msgp.Encode(&buf, &v)

	m := v.{{.MethodPrefix}}Msgsize()
	if buf.Len() > m {
		t.Logf("WARNING: {{.MethodPrefix}}Msgsize() for %v is inaccurate", v)
	}

	vn := {{.TypeName}}{{.GenericInstantiation}}{}
	err := msgp.Decode(&buf, &vn)
	if err != nil {
		t.Error(err)
	}

	buf.Reset()
	msgp.Encode(&buf, &v)
	err = msgp.NewReader(&buf).Skip()
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkEncode{{.TypeName}}{{.MethodPrefix}}(b *testing.B) {
	v := {{.TypeName}}{{.GenericInstantiation}}{}
	var buf bytes.Buffer 
	msgp.Encode(&buf, &v)
	b.SetBytes(int64(buf.Len()))
	en := msgp.NewWriter(msgp.Nowhere)
	b.ReportAllocs()
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		v.{{.MethodPrefix}}EncodeMsg(en)
	}
	en.Flush()
}

func BenchmarkDecode{{.TypeName}}{{.MethodPrefix}}(b *testing.B) {
	v := {{.TypeName}}{{.GenericInstantiation}}{}
	var buf bytes.Buffer
	msgp.Encode(&buf, &v)
	b.SetBytes(int64(buf.Len()))
	rd := msgp.NewEndlessReader(buf.Bytes(), b)
	dc := msgp.NewReader(rd)
	b.ReportAllocs()
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		err := v.{{.MethodPrefix}}DecodeMsg(dc)
		if  err != nil {
			b.Fatal(err)
		}
	}
}

`))

}
