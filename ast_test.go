package goast

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestAst(t *testing.T) {

	fset := token.NewFileSet()
	f := ParseFile(fset, "foo.go", fileTpl, parser.ParseComments)
	Dump(fset, f)
}

var fileTpl = `
package main

import (
	"net/url"
	"time"
)

func (c *Client) Action(a int) *Out {
	in := In{
		A : a,
		B : a,
	}
	params := url.Values{}
	c.Send("Func", params, in)
	return a,b
}
`
