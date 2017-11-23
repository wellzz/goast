package goast

import (
	"go/ast"
	"bytes"
	"go/printer"
	"go/token"
)


func Bytes(node ast.Node) []byte {
	fset := token.NewFileSet()
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, node)
	return buf.Bytes()
}

func String(node ast.Node) string {
	fset := token.NewFileSet()
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, node)
	return buf.String()
}

func Dump(node ast.Node) {
	fset := token.NewFileSet()
	ast.Print(fset, node)
}
