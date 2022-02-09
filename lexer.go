package main

import (
	"io"
)

const (
	tokenEOF    int = 1
	tokenPlus   int = '+'
	tokenMinus  int = '-'
	tokenTimes  int = '*'
	tokenDiv    int = '/'
	tokenLParen int = '('
	tokenRParen int = ')'
)

type Lexer struct {
	src       io.ByteReader
	lookahead int
}

func NewLexer(src io.ByteReader) *Lexer {
	return &Lexer{
		src: src,
	}
}

func (l *Lexer) Init() {
	l.getChar()
}

func (l *Lexer) Expr() {
	if l.isAddop() {
		println("addop")
	} else {
		l.term()
	}

	for l.lookahead == tokenPlus || l.lookahead == tokenMinus {
		if l.lookahead == '+' {
			l.add()
			println("+")
		} else if l.lookahead == '-' {
			l.subtract()
			println("-")
		} else {
			l.expected("addop")
		}
	}
}

func (l *Lexer) term() {
	l.factor()

	for l.lookahead == tokenDiv || l.lookahead == tokenTimes {
		if l.lookahead == tokenTimes {
			l.mul()
			println("*")
		} else if l.lookahead == tokenDiv {
			l.div()
			println("/")
		} else {
			l.expected("Mulop")
		}
	}
}

func (l *Lexer) ident() {
	name := l.getName()

	if l.lookahead == tokenLParen {
		l.match(tokenLParen)
		l.match(tokenRParen)
		print("funcCall", name)
	} else {
		println(name)
	}
}

func (l *Lexer) assign() {
	name := l.getName()

	l.match('=')
	l.Expr()
	println(name)
}

func (l *Lexer) isAlnum(ch int) bool {
	return l.isAlpha(ch) || l.isDigit(ch)
}

func (l *Lexer) factor() {
	if l.lookahead == tokenLParen {
		l.match(tokenLParen)
		l.Expr()
		l.match(tokenRParen)
	} else if l.isAlpha(l.lookahead) {
		println(l.getNum())
	} else {
		println(l.getNum())
	}
}

func (l *Lexer) add() {
	l.match(tokenPlus)
	l.term()
}

func (l *Lexer) subtract() {
	l.match(tokenMinus)
	l.term()
}

func (l *Lexer) mul() {
	l.match(tokenTimes)
	l.factor()
}

func (l *Lexer) div() {
	l.match(tokenDiv)
	l.factor()
}

func (l *Lexer) isAddop() bool {
	return l.lookahead == tokenPlus || l.lookahead == tokenMinus
}

func (l *Lexer) getChar() {
	lookahead, err := l.src.ReadByte()
	if err != nil {
		l.lookahead = tokenEOF
		return
	}

	l.lookahead = int(lookahead)
}

func (l *Lexer) error(msg string) {
	print("error:", msg)
}

func (l *Lexer) abort(msg string) {
	l.error(msg)
	panic("abort")
}

func (l *Lexer) expected(msg string) {
	l.abort("expected " + msg)
}

func (l *Lexer) match(t int) {
	if l.lookahead == t {
		l.getChar()
		l.skipWhite()
	} else {
		l.expected(string(t))
	}
}

func (l *Lexer) isAlpha(ch int) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func (l *Lexer) isDigit(ch int) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) getName() string {
	if !l.isAlpha(l.lookahead) {
		l.expected("Name")
	}

	name := ""

	for l.isAlnum(l.lookahead) {
		name += string(rune(l.lookahead))
		l.getChar()
	}

	l.skipWhite()
	return name
}

func (l *Lexer) getNum() string {
	if !l.isDigit(l.lookahead) {
		l.expected("Integer")
	}

	num := ""

	for l.isDigit(l.lookahead) {
		num += string(rune(l.lookahead))
		l.getChar()
	}

	l.skipWhite()

	return num
}

func (l *Lexer) isWhite(ch int) bool {
	return ch == ' ' || ch == '\t'
}

func (l *Lexer) skipWhite() {
	for l.isWhite(l.lookahead) {
		l.getChar()
	}
}
