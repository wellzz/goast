package goast

import "go/ast"

func MethodDecl(name string, recv *ast.Field, funcType *ast.FuncType, body []ast.Stmt) *ast.FuncDecl {
	f := FuncDecl(name, funcType, body)
	f.Recv = Fields(recv)
	return f
}

func FuncDecl(n string, t *ast.FuncType, body []ast.Stmt) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: Ident(n),
		Body: &ast.BlockStmt{List: body},
		Type: t,
	}
}

func FuncType(params, results *ast.FieldList) *ast.FuncType {
	return &ast.FuncType{
		Params:  params,
		Results: results,
	}
}
