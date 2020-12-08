package testdoc_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Eun/go-testdoc"

	"github.com/traefik/yaegi/interp"
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
		t.Run("public", test("./tests/package", false, []string{
			"Package",
		}))
		t.Run("private", test("./tests/package", true, []string{
			"Package",
		}))
	})

	t.Run("func", func(t *testing.T) {
		t.Run("public", test("./tests/func", false, []string{
			"PublicFunc",
		}))
		t.Run("private", test("./tests/func", true, []string{
			"PublicFunc",
			"privateFunc",
		}))
	})

	t.Run("variable", func(t *testing.T) {
		t.Run("public", test("./tests/variable", false, []string{
			"PublicVariable",
		}))
		t.Run("private", test("./tests/variable", true, []string{
			"PublicVariable",
			"privateVariable",
		}))
	})

	t.Run("const", func(t *testing.T) {
		t.Run("public", test("./tests/const", false, []string{
			"PublicConst",
		}))
		t.Run("private", test("./tests/const", true, []string{
			"PublicConst",
			"privateConst",
		}))
	})

	t.Run("struct", func(t *testing.T) {
		t.Run("public", test("./tests/struct", false, []string{
			"PublicStruct",
			"PublicStruct.PublicStructField",
			"privateStruct.PublicStructField",
		}))
		t.Run("private", test("./tests/struct", true, []string{
			"PublicStruct",
			"PublicStruct.PublicStructField",
			"PublicStruct.privateStructField",
			"privateStruct",
			"privateStruct.PublicStructField",
			"privateStruct.privateStructField",
		}))
	})

	t.Run("interface", func(t *testing.T) {
		t.Run("public", test("./tests/interface", false, []string{
			"PublicInterface",
			"PublicInterface.PublicInterfaceMethod",
			"privateInterface.PublicInterfaceMethod",
		}))
		t.Run("private", test("./tests/interface", true, []string{
			"PublicInterface",
			"PublicInterface.PublicInterfaceMethod",
			"PublicInterface.privateInterfaceMethod",
			"privateInterface",
			"privateInterface.PublicInterfaceMethod",
			"privateInterface.privateInterfaceMethod",
		}))
	})

	t.Run("reference", func(t *testing.T) {
		t.Run("public", test("./tests/reference", false, []string{
			"Reference",
		}))
		t.Run("private", test("./tests/reference", true, []string{
			"Reference",
		}))
	})
}
