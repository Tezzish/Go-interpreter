package parser

import (
	"Go-interpreter/ast"
	"Go-interpreter/lexer"
	"Go-interpreter/token"
	"fmt"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token // the current token that we're looking at
	peekToken token.Token // the next token that we're looking at
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// checks if the next token is of the type we expect
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		// if it is, we move to the next token
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	// initialise the root node of the AST
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	//until we reach the end of the file
	for p.curToken.Type != token.EOF {
		//the current statement
		stmt := p.parseStatement()
		if stmt != nil {
			//we add the current statements to the list of statements in the program
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	// if the next token is not an identifier, we return nil
	// if it is, this moves the current token to the identifier token
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	// we then assign the identifier value to the current token's value
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	// if the next token is not an assignment operator, we return nil
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// we skip the expression until we reach a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
