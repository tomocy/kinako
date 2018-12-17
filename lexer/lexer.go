package lexer

import (
	"strings"

	"github.com/tomocy/kinako/token"
)

const whitespaces = " \t\r\n"

type Lexer struct {
	input            string
	currentCharacter rune
	currentPosition  int
	readingPosition  int
}

func New(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

func (l *Lexer) ReadNextToken() token.Token {
	l.readCharacter()
	l.skipWhitespaces()

	return l.readToken()
}

func (l *Lexer) readToken() token.Token {
	switch l.currentCharacter {
	case '+', '-', '*':
		return l.readSingleToken()
	case 0:
		return l.readEOF()
	default:
		if l.hasDigit() {
			return l.raedInteger()
		}

		return l.readUnknown()
	}
}

func (l *Lexer) readSingleToken() token.Token {
	literal := string(l.currentCharacter)
	return token.Token{
		Type:    token.LookUpType(literal),
		Literal: literal,
	}
}

func (l *Lexer) readEOF() token.Token {
	return token.Token{
		Type:    token.EOF,
		Literal: "",
	}
}

func (l *Lexer) hasDigit() bool {
	return isDigit(l.currentCharacter)
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func (l *Lexer) raedInteger() token.Token {
	return token.Token{
		Type:    token.Integer,
		Literal: l.readNumber(),
	}
}

func (l *Lexer) readNumber() string {
	begin := l.currentPosition
	for l.willHaveDigit() {
		l.readCharacter()
	}

	return l.input[begin:l.readingPosition]
}

func (l *Lexer) willHaveDigit() bool {
	return isDigit(l.peekCharacter())
}

func (l *Lexer) readUnknown() token.Token {
	return token.Token{
		Type:    token.Unknown,
		Literal: l.readWord(),
	}
}

func (l *Lexer) readWord() string {
	begin := l.currentPosition
	for l.willHaveWhitespace() {
		l.readCharacter()
	}

	return l.input[begin:l.readingPosition]
}

func (l *Lexer) skipWhitespaces() {
	for l.hasWhitespace() {
		l.readCharacter()
	}
}

func (l *Lexer) hasWhitespace() bool {
	return isWhitespace(l.currentCharacter)
}

func (l *Lexer) willHaveWhitespace() bool {
	return isWhitespace(l.peekCharacter())
}

func isWhitespace(r rune) bool {
	return strings.ContainsRune(whitespaces, r)
}

func (l *Lexer) readCharacter() {
	if len(l.input) <= l.readingPosition {
		l.currentCharacter = 0
	} else {
		l.currentCharacter = rune(l.input[l.readingPosition])
	}

	l.currentPosition = l.readingPosition
	l.readingPosition++
}

func (l *Lexer) peekCharacter() rune {
	if len(l.input) <= l.readingPosition {
		return 0
	}

	return rune(l.input[l.readingPosition])
}
