package lexer

import (
	"Go-interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current character in terms of index
	readPosition int  // current reading position in input after the current character
	character    byte // current character in terms of value
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.character {
	case '=':
		tok = newToken(token.ASSIGN, l.character)
	case ';':
		tok = newToken(token.SEMICOLON, l.character)
	case '(':
		tok = newToken(token.LPAREN, l.character)
	case ')':
		tok = newToken(token.RPAREN, l.character)
	case ',':
		tok = newToken(token.COMMA, l.character)
	case '+':
		tok = newToken(token.PLUS, l.character)
	case '{':
		tok = newToken(token.LBRACE, l.character)
	case '}':
		tok = newToken(token.RBRACE, l.character)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}
