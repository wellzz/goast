package goast

import "go/ast"

func Field(t string, n ...string) *ast.Field {
	if t == "" {
		return nil
	}
	var e ast.Expr
	if t[0] == '*' {
		e = StarExpr(t[1:])
	} else {
		e = Ident(t)
	}
	return &ast.Field{
		Names: Idents(n...),
		Type:  e,
	}
}

func Fields(fields ...*ast.Field) *ast.FieldList {
	if len(fields) == 0 || fields[0] == nil {
		return nil
	}
	return &ast.FieldList{List: fields}
}
