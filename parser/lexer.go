package parser

import (
	"fmt"
	"io"
	"text/scanner"
)

var marks = map[string]int{
	"@":  HEADING_MARKER,
	"\n": CR,
}

type Token struct {
	token   int
	literal string
}

func (t Token) Name() string {
	switch t.token {
	case TEXT:
		return "TEXT"
	case CR:
		return "CR"
	case HEADING_MARKER:
		return "HEADING_MARKER"
	default:
		return "UNKNOWN"
	}
}

type Lexer struct {
	scanner.Scanner
	result Document
	err    *ParseError
}

func NewLexer(in io.Reader) (l *Lexer) {
	l = &Lexer{}
	l.Scanner.Mode = scanner.GoTokens
	l.Init(in)
	return
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	print(l.Column)

	if token == scanner.Ident {
		if mark, ok := marks[l.TokenText()]; ok {
			token = mark
		} else {
			token = TEXT
		}
	} else {
		token = TEXT
	}

	lval.token = Token{token: token, literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	l.err = &ParseError{
		Message: e,
		Line:    l.Line,
		Column:  l.Column,
	}
}

type ParseError struct {
	Message string
	Line    int
	Column  int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Line: %d, Column: %d: %s", e.Line, e.Column, e.Message)
}
