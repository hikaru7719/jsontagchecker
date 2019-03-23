package jsontagchecker

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "jsontagchecker",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "jsontagchecker is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	inspect.Preorder(nodeFilter, func(node ast.Node) {
		if structType, ok := node.(*ast.StructType); ok {
			for _, field := range structType.Fields.List {
				fieldName := field.Names[0].Name
				tag := field.Tag.Value
				if !Checker(fieldName, tag) {
					pass.Reportf(field.Tag.Pos(), "invalid snake case json tag")
				}
			}
		}
	})
	return nil, nil
}
