package token

type Token struct {
	Type    Type
	Literal string
}

type Type string

const (
	Unknown Type = "Unknown"

	EOF = "EOF"

	Integer = "Integer"
)
