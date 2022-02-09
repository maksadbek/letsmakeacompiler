package main

import (
	"strings"
)

func main() {
	l := NewLexer(strings.NewReader("1+2*3-4"))
	l.Init()
	l.Expr()
}
