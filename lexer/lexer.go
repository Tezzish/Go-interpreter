package lexer

import (
	"Go-interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current character in terms of index
	readPosition int  // position of next character in terms of index
	character    byte // current character in terms of value
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// function for creating a new token
func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

// function for reading the character
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

	l.skipWhitespace()

	switch l.character {
	case '=':
		if l.peekChar() == '=' {
			ch := l.character
			tok = newToken(token.ASSIGN, l.character)
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.character)}
		} else {
			tok = newToken(token.ASSIGN, l.character)
		}
	case ',':
		tok = newToken(token.COMMA, l.character)
	case '!':
		if l.peekChar() == '=' {
			ch := l.character
			l.readChar()
			tok = token.Token{Type: token.NEQ, Literal: string(ch) + string(l.character)}
		} else {
			tok = newToken(token.BANG, l.character)
		}
	case '+':
		tok = newToken(token.PLUS, l.character)
	case '-':
		tok = newToken(token.MINUS, l.character)
	case '*':
		tok = newToken(token.ASTERISK, l.character)
	case '/':
		tok = newToken(token.SLASH, l.character)
	case '<':
		tok = newToken(token.LT, l.character)
	case '>':
		tok = newToken(token.GT, l.character)
	case '(':
		tok = newToken(token.LPAREN, l.character)
	case ')':
		tok = newToken(token.RPAREN, l.character)
	case '{':
		tok = newToken(token.LBRACE, l.character)
	case '}':
		tok = newToken(token.RBRACE, l.character)
	case ';':
		tok = newToken(token.SEMICOLON, l.character)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.character) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.character) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.character)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.character) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.character) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readChar()
	}
}
