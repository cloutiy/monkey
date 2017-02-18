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

	//Skip whitespace
	l.skipWisespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
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

/*
Reads an identifier
*/
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/*
Checks if a character is a letter
*/
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '-'
}

/*
Skips whitespace
Note to self: This function skips over newlines. To include position
and line number in errors, it would be necessary to handle newline
differently (if newline => currentline++ and resetCursorPosition()
*/
func (l *Lexer) skipWisespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/*
Reads a number
*/
func (l *Lexer) readNumber() string {
	position := l.readPosition

	//Keep advancing the cursor position while the char is a digit
	for isDigit(l.ch) {
		l.readChar()
	}

	//Return the string representing the number
	return l.input[position:l.readPosition]
}

/*
Checks if a char is a digit
*/
func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

/*
Looks ahead onechar
*/
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
