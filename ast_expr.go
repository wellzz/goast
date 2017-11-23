package goast

import "go/ast"

func Ident(n string) *ast.Ident {
	if n == "" {
		return nil
	}
	return &ast.Ident{Name: n}
}

func Idents(n ...string) []*ast.Ident {
	if len(n) == 0 {
		return nil
	}
	ret := make([]*ast.Ident, 0, len(n))
	for _, n := range n {
		ret = append(ret, Ident(n))
	}
	return ret
}

func Expr(n string) ast.Expr {
	if n == "" {
		return nil
	}
	if n[0] == '*' {
		return StarExpr(n[1:])
	}
	return Ident(n)
}

func Exprs(n ...string) []ast.Expr {
	if len(n) == 0 {
		return nil
	}
	ret := make([]ast.Expr, 0, len(n))
	for _, n := range n {
		ret = append(ret, Expr(n))
	}
	return ret
}

func TypeAssertExpr(x ast.Expr, t string) *ast.TypeAssertExpr {
	return &ast.TypeAssertExpr{X: x, Type: Ident(t)}
}

func KeyValueExpr(key, value string) *ast.KeyValueExpr {
	return &ast.KeyValueExpr{Key: Ident(key), Value: Ident(value)}
}

func StarExpr(n string) *ast.StarExpr {
	if n == "" {
		return nil
	}
	return &ast.StarExpr{X: Ident(n)}
}

func CallExpr(f string, args ...string) *ast.CallExpr {
	return &ast.CallExpr{Fun: Ident(f), Args: Exprs(args...)}
}

func CompositeLit(t string, elt ...ast.Expr) *ast.CompositeLit {
	return &ast.CompositeLit{
		Type: Expr(t),
		Elts: elt,
	}
}
