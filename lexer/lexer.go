// lexer/lexer.go
package lexer

import "github.com/cloutiy/monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  //current reading position (after current char)
	ch           byte //current char under examination
}

/*
Creates a lexer structure and returns a pointer to its location
*/
func New(input string) *Lexer {
	l := &Lexer{input: input} //instantiate and return its location
	l.readChar()
	return l
}

/*
Reads a character from the input
*/
func (l *Lexer) readChar() {
	//if we have reached end of input, assign EOF to current char value
	if l.readPosition >= len(l.input) {
		l.ch = 0

		//otherwise assign the char to current char
	} else {
		l.ch = l.input[l.readPosition]
	}

	//move the current position to the next char
	l.position = l.readPosition
	l.readPosition += 1
}

/*
Gets the next token from the input
*/
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	//move to next char
	l.readChar()

	//return the token
	return tok
}

/*
create a new token
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
