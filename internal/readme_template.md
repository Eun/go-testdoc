<$-
package main

import (
	"io"
	"os"
)

func printFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := io.WriteString(os.Stdout, "```go\n"); err != nil {
		panic(err)
    }
	if _, err := io.Copy(os.Stdout, f); err != nil {
		panic(err)
    }
	if _, err := io.WriteString(os.Stdout, "\n```"); err != nil {
		panic(err)
	}
}
-$>
# go-testdoc [![Actions Status](https://github.com/Eun/go-testdoc/workflows/CI/badge.svg)](https://github.com/Eun/go-testdoc/actions) [![Coverage Status](https://coveralls.io/repos/github/Eun/go-hit/badge.svg?branch=master)](https://coveralls.io/github/Eun/go-hit?branch=master) [![PkgGoDev](https://img.shields.io/badge/pkg.go.dev-reference-blue)](https://pkg.go.dev/github.com/Eun/go-hit) [![GoDoc](https://godoc.org/github.com/Eun/go-hit?status.svg)](https://godoc.org/github.com/Eun/go-hit) [![go-report](https://goreportcard.com/badge/github.com/Eun/go-testdoc)](https://goreportcard.com/report/github.com/Eun/go-testdoc)
*go-testdoc* runs your code documentation examples during normal test time.
> go get -u github.com/Eun/go-testdoc

### Example
This example will run the go code inside the `Example` and `Examples` section using [yaegi](https://github.com/traefik/yaegi).
<$ printFile("./example/example.go") $>
<$ printFile("./example/example_test.go") $>


## See also
[Testing your `README.md` file](https://github.com/Eun/yaegi-template/tree/master/examples/evaluate_readme)