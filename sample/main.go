package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "my.go", nil, 0)
	if err != nil {
		fmt.Println(err)
	}
	ast.Print(nil, f)
}
