package parse

import (
	"go/ast"
	//"go/parser"
	"fmt"
	"go/token"
	"go/types"

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
			_ = expr
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

// TypeInstantiationInfo holds information about a generic type instantiation
type TypeInstantiationInfo struct {
	TypeName     string
	TypeArgs     []types.Type
	TypeArgNames []string
	Position     token.Position
}

func AnalyzeGenericTypes(filepath string) (generics map[string]*TypeInstantiationInfo, err error) {
	// Configure package loading
	cfg := &packages.Config{
		Mode:  packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		Tests: false,
	}
	vv("about to load filepath '%v'", filepath)
	pkgs, err := packages.Load(cfg, filepath)
	if err != nil {
		return nil, err
	}
	vv("back from Load okay.")
	// parent type key
	generics = make(map[string]*TypeInstantiationInfo)

	// Analyze each package
	for _, pkg := range pkgs {
		info := pkg.TypesInfo

		// Visit all AST nodes
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
						typeArgs := make([]types.Type, n)
						typeArgNames := make([]string, n)
						for i := 0; i < inst.TypeArgs().Len(); i++ {
							typeArgs[i] = inst.TypeArgs().At(i)
							typeArgNames[i] = typeArgs[i].String()
						}

						info := &TypeInstantiationInfo{
							TypeName:     inst.Obj().Name(),
							TypeArgs:     typeArgs,
							TypeArgNames: typeArgNames,
							Position:     pkg.Fset.Position(indexExpr.Pos()),
						}
						generics[info.TypeName] = info
					}
				}
				return true
			})
		}
	}

	return
}
