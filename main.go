package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)


type Lexer struct {
	src io.ByteReader
	lookahead int
}

func (l *Lexer) getChar() {
	lookahead, err := l.src.ReadByte()
	if err != nil {
		panic(err)
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

	name := l.lookahead
	l.getChar()

	return string(name)
}

func (l *Lexer) getNum() int {
	if !l.isDigit(l.lookahead) {
		l.expected("Integer")
	}

	return l.lookahead
}

func (l *Lexer) init() {
	l.getChar()
}
