package goast

import (
	"go/ast"
	"go/token"
)

func ImportDecl() *ast.GenDecl {
	return &ast.GenDecl{
		Tok:    token.IMPORT,
		Lparen: 1,
		Rparen: 1,
		Specs:  make([]ast.Spec, 0),
	}
}


