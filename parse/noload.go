package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/glycerine/greenpack/cfg"
	"github.com/glycerine/greenpack/gen"
)

// FileNoLoad parses a file at the relative path
// provided and produces a new *FileSet.
// If you pass in a path to a directory, the entire
// directory will be parsed.
// If unexport is false, only exported identifiers are included in the FileSet.
// If the resulting FileSet would be empty, an error is returned.
//
// FileNoLoad(), in noload.go, is
// the original msgp version of File()
// that doesn't require full
// compilability/avialability
// of all dependencies. Although this
// doesn't support resolution of
// named constants like the loader
// version does, this can be
// useful when reading a partial
// completed source file or otherwise
// in a situation where it is
// inconvient to have to meet
// the compiler's demands just yet.
func FileNoLoad(c *cfg.GreenConfig) (*FileSet, error) {
	//vv("FileNoLoad top")
	ok, isDir := fileOrDir(c.GoFile)
	if !ok {
		return nil, fmt.Errorf("error: path '%s' does not exist", c.GoFile)
	}

	name := c.GoFile
	pushstate(name)
	defer popstate()
	fs := &FileSet{
		Specs:              make(map[string]*ast.TypeSpec),
		Identities:         make(map[string]gen.Elem),
		Cfg:                c,
		InterfaceTypeNames: make(map[string]bool),
		GenericTypeParams:  make(map[string]*gen.Genric),
		Instan:             make(map[string][]*gen.Instan),
	}
	fs.InterfaceTypeNames["any"] = true

	fset := token.NewFileSet()
	if isDir {
		pkgs, err := parser.ParseDir(fset, name, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		if len(pkgs) != 1 {
			return nil, fmt.Errorf("multiple packages in directory: %s", name)
		}
		var one *ast.Package
		for _, nm := range pkgs {
			one = nm
			break
		}
		fs.Package = one.Name
		for _, fl := range one.Files {
			pushstate(fl.Name.Name)
			fs.Directives = append(fs.Directives, yieldComments(fl.Comments)...)
			fs.getInterfaceNames(fl)
			fs.getZebraSchemaId(fl)
			if !c.Unexported {
				ast.FileExports(fl)
			}
			fs.getTypeSpecs(fl)
			fs.getTemplateInstantiations(fl)
			popstate()
		}
	} else {
		f, err := parser.ParseFile(fset, name, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		fs.Package = f.Name.Name
		fs.Directives = yieldComments(f.Comments)
		fs.getInterfaceNames(f)
		fs.getZebraSchemaId(f)
		if !c.Unexported {
			ast.FileExports(f)
		}
		fs.getTypeSpecs(f)
		fs.getTemplateInstantiations(f)
	}

	if len(fs.Specs) == 0 {
		return nil, fmt.Errorf("no definitions in %s", name)
	}

	err := fs.process()
	if err != nil {
		return nil, err
	}
	fs.applyDirectives()
	fs.propInline()

	return fs, nil
}
