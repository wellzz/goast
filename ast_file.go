package goast

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func File(name string) *ast.File {
	return &ast.File{Name: Ident(name)}
}

func ParseFile(fset *token.FileSet, fileName string, src interface{}, mode parser.Mode) *ast.File {
	f, err := parser.ParseFile(fset, fileName, src, mode)
	if err != nil {
		panic(err)
	}
	return f
}