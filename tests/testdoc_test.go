package tests_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/traefik/yaegi/interp"

	"github.com/Eun/go-testdoc"
)

func test(dir string, includePrivate bool, expectedCalls []string) func(t *testing.T) {
	return func(t *testing.T) {
		var gotCalls []string
		testdoc.TestCodeDocumentation(t, &testdoc.Options{
			Path:     dir,
			PkgName:  "test",
			Sections: []string{"Example"},
			Symbols: []interp.Exports{
				map[string]map[string]reflect.Value{
					"test": {
						"test": reflect.ValueOf(func(caller string) {
							gotCalls = append(gotCalls, caller)
						}),
					},
				},
			},
			Imports: []testdoc.Import{
				{
					Name: ".",
					Path: "test",
				},
				{
					Name: ".",
					Path: dir,
				},
			},
			IncludePrivate: includePrivate,
		})
		require.Equal(t, expectedCalls, gotCalls)
	}
}

func TestTestCodeDocumentation(t *testing.T) {
	t.Run("package", func(t *testing.T) {
		t.Run("public", test("./package", false, []string{
			"Package",
		}))
		t.Run("private", test("./package", true, []string{
			"Package",
		}))
	})

	t.Run("func", func(t *testing.T) {
		t.Run("public", test("./func", false, []string{
			"PublicFunc",
		}))
		t.Run("private", test("./func", true, []string{
			"PublicFunc",
			"privateFunc",
		}))
	})

	t.Run("variable", func(t *testing.T) {
		t.Run("public", test("./variable", false, []string{
			"PublicVariable",
		}))
		t.Run("private", test("./variable", true, []string{
			"PublicVariable",
			"privateVariable",
		}))
	})

	t.Run("const", func(t *testing.T) {
		t.Run("public", test("./const", false, []string{
			"PublicConst",
		}))
		t.Run("private", test("./const", true, []string{
			"PublicConst",
			"privateConst",
		}))
	})

	t.Run("struct", func(t *testing.T) {
		t.Run("public", test("./struct", false, []string{
			"PublicStruct",
			"PublicStruct.PublicStructField",
			"privateStruct.PublicStructField",
		}))
		t.Run("private", test("./struct", true, []string{
			"PublicStruct",
			"PublicStruct.PublicStructField",
			"PublicStruct.privateStructField",
			"privateStruct",
			"privateStruct.PublicStructField",
			"privateStruct.privateStructField",
		}))
	})

	t.Run("interface", func(t *testing.T) {
		t.Run("public", test("./interface", false, []string{
			"PublicInterface",
			"PublicInterface.PublicInterfaceMethod",
			"privateInterface.PublicInterfaceMethod",
		}))
		t.Run("private", test("./interface", true, []string{
			"PublicInterface",
			"PublicInterface.PublicInterfaceMethod",
			"PublicInterface.privateInterfaceMethod",
			"privateInterface",
			"privateInterface.PublicInterfaceMethod",
			"privateInterface.privateInterfaceMethod",
		}))
	})

	t.Run("reference", func(t *testing.T) {
		t.Run("public", test("./reference", false, []string{
			"Reference",
		}))
		t.Run("private", test("./reference", true, []string{
			"Reference",
		}))
	})
}
