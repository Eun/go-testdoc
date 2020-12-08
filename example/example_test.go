// example_test.go file
package example_test

import (
	"testing"

	"github.com/Eun/go-testdoc"
)

func TestDocumentation(t *testing.T) {
	testdoc.TestCodeDocumentation(t, &testdoc.Options{
		// Test for this folder
		Path: ".",

		// Test the `example` package
		PkgName: "example",

		// Execute code inside the `Example` and `Examples` sections
		Sections: []string{"Example", "Examples"},

		Imports: []testdoc.Import{
			// Import some standard packages we need
			{Name: "", Path: "fmt"},

			// Import the current package so we can call the functions.
			{Name: ".", Path: "./"},
		},
	})
}
