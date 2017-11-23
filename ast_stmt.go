package goast

import (
	"go/ast"
	"go/token"
)

func ExprStmt(x ast.Expr) *ast.ExprStmt{
	return &ast.ExprStmt{X:x}
}

//func AssignStmt(l, r []ast.Expr, tok token.Token) *ast.AssignStmt {
//	return &ast.AssignStmt{
//		Lhs: l,
//		Tok: tok,
//		Rhs: r,
//	}
//}
func AssignStmt(l, r ast.Expr, tok token.Token) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{l},
		Tok: tok,
		Rhs: []ast.Expr{r},
	}
}

func ReturnStmt(n ...string) *ast.ReturnStmt{
	if len(n) == 0{
		return nil
	}
	return &ast.ReturnStmt{Results:Exprs(n...)}
}
