package parse

import (
	"go/ast"
	//"go/parser"
	"fmt"
	//"go/token"
	"go/types"
	"path/filepath"

	"github.com/glycerine/greenpack/gen"
	"golang.org/x/tools/go/packages"
)

func genericTypeParam(path string) (generics map[string]string, err error) {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
	}

	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		return nil, fmt.Errorf("error loading "+
			"packages for path '%v': %v\n", path, err)
	}

	generics = make(map[string]string) // parent type, param names
	for _, pkg := range pkgs {
		info := pkg.TypesInfo

		// Look through type information for generic instantiations
		for expr, tv := range info.Types {
			// tv is go/types.TypeAndValue
			if named, ok := tv.Type.(*types.Named); ok {

				ta := named.TypeArgs()
				if ta != nil {
					fmt.Printf("Found generic instantiation at %v:\n", pkg.Fset.Position(expr.Pos()))
					//fmt.Printf("  Type: %v\n", tv.Type)
					//fmt.Printf("  Type arguments: %#v\n", ta)
					types := ta.Types()
					for t := range types {
						// command-line-arguments.Matrix[float64] types has param float64
						fmt.Printf("%v types has param %v\n", tv.Type.String(), t)
						generics[tv.Type.String()] = t.String()
					}
				}
			}
		}
	}
	return
}

func analyzeGenericTypes(path string) (generics map[string][]*gen.Instan, err error) {
	vv("top analyzeGenericTypes")

	// Get the absolute directory path of the file
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	dir := filepath.Dir(absPath)

	// Configure package loading
	cfg := &packages.Config{
		Mode: packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedSyntax |
			packages.NeedFiles |
			packages.NeedDeps |
			packages.NeedImports,
		Tests: false,
		Dir:   absPath, // Set working directory to the module directory
	}
	vv("about to load path '%v' from dir '%v'", path, dir)

	// Load using the package pattern
	pkgs, err := packages.Load(cfg, ".") // Use "." to load package in current dir
	if err != nil {
		return nil, err
	}
	vv("back from Load okay. len(pkgs)=%v", len(pkgs))
	for _, pkg := range pkgs {
		vv("package %v has %v syntax files and %v types",
			pkg.ID,
			len(pkg.Syntax),
			len(pkg.TypesInfo.Types))
		if len(pkg.Errors) > 0 {
			vv("package had errors: %v", pkg.Errors)
		}
	}
	// parent type key -> slice of instantiations
	generics = make(map[string][]*gen.Instan)

	// Analyze each package
	for _, pkg := range pkgs {
		info := pkg.TypesInfo

		// Look through type information for generic instantiations

		// more of these than Syntax, so use Syntax below
		if true {
			//vv("looking through %v of info.Types", len(info.Types))
			for expr, tv := range info.Types {
				if inst, ok := tv.Type.(*types.Named); ok {
					// Get the type arguments
					n := inst.TypeArgs().Len()
					if n == 0 {
						continue
					}
					typeArgs := make([]types.Type, n)
					typeArgNames := make([]string, n)
					for i := 0; i < inst.TypeArgs().Len(); i++ {
						typeArgs[i] = inst.TypeArgs().At(i)
						typeArgNames[i] = typeArgs[i].String()
					}
					nm := inst.Obj().Name()
					info := &gen.Instan{
						TypeName:     nm,
						TypeArgs:     typeArgs,
						TypeArgNames: typeArgNames,
						Position:     pkg.Fset.Position(expr.Pos()),
					}
					//vv("attempt 1 instan-> %v with %v", nm, info.TypeArgNames)
					generics[nm] = append(generics[nm], info)
				}
			}
		}

		if false {
			// They both work, this has fewer pkg.Syntax (1).
			// Visit all AST nodes
			vv("looking through %v of pkg.Syntax", len(pkg.Syntax))
			for _, file := range pkg.Syntax {
				ast.Inspect(file, func(n ast.Node) bool {
					// Look for IndexExpr which is used
					// for generic type instantiation

					indexExpr, ok := n.(*ast.IndexExpr)
					if !ok {
						return true
					}

					// Get type information for the expression
					if t, ok := info.Types[indexExpr]; ok {
						if inst, ok := t.Type.(*types.Named); ok {
							// Get the type arguments
							n := inst.TypeArgs().Len()
							if n == 0 {
								return true
							}
							typeArgs := make([]types.Type, n)
							typeArgNames := make([]string, n)
							for i := 0; i < inst.TypeArgs().Len(); i++ {
								typeArgs[i] = inst.TypeArgs().At(i)
								typeArgNames[i] = typeArgs[i].String()
							}
							nm := inst.Obj().Name()
							info := &gen.Instan{
								TypeName:     nm,
								TypeArgs:     typeArgs,
								TypeArgNames: typeArgNames,
								Position:     pkg.Fset.Position(indexExpr.Pos()),
							}
							vv("attempt 2 instan-> %v with %v", nm, info.TypeArgNames)
							generics[nm] = append(generics[nm], info)
						}
					}
					return true
				})
			}
		}
	}

	return
}
