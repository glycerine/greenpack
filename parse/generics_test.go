package parse

import (
	"fmt"
	//"io/ioutil"
	//"log"
	//"os"
	//"sort"

	cv "github.com/glycerine/goconvey/convey"
	//"github.com/glycerine/greenpack/cfg"
	//"github.com/glycerine/greenpack/gen"
	"testing"
)

func Test999_how_are_generics_instantiated_for_gen_tests(t *testing.T) {
	cv.Convey("If we know how generics are instantiated we can generate tests for them", t, func() {

		path := "."
		instantiations, err := analyzeGenericTypes(path)
		if err != nil {
			t.Fatalf("AnalyzeGenericTypes on path '%v' gave err='%v'", path, err)
		}

		for nm, inst := range instantiations {
			// there are type sets too, skip these.
			// nm=Addable, Generic type Addable instantiated at - with args:
			for _, ins := range inst {
				fmt.Printf("nm=%v, Generic type %s instantiated at %v with args: \n", nm, ins.TypeName, ins.Position)
				for _, arg := range ins.TypeArgs {
					fmt.Printf("  - %v\n", arg)
				}
			}
		}
	})
}
