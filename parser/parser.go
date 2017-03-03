//parser/parser.go

package parser

import {
	"github.com/cloutiy/ast"
	"github.com/cloutiy/lexer"
	"github.com/cloutiy/token"
}

type Parser struct {
	l *lexer.Lexer
	
	curToken token.Token
	peekToken token.Token
}

func New() *lexer.Lexer {
	p := &Parser{l: l}
	
	//Read 2 tokens so that current and peek tokens are set
	p.nextToken()
	p.nextToken()
	
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.nextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}