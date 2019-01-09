package token

type Token struct {
	Type    Type
	Literal string
}

type Type string

const (
	Unknown Type = "Unknown"

	EOF = "EOF"

	Plus     = "Plus"
	Minus    = "Minus"
	Asterisk = "Asterisk"
	Slash    = "Slash"

	Assign = "Assign"

	LParen = "LParen"
	RParen = "RParen"

	Ident   = "Ident"
	Integer = "Integer"

	Var = "var"
)

var types = map[string]Type{
	"+": Plus,
	"-": Minus,
	"*": Asterisk,
	"/": Slash,
	"=": Assign,
	"(": LParen,
	")": RParen,
}

func LookUpType(s string) Type {
	if t, ok := types[s]; ok {
		return t
	}

	return Unknown
}

var keywords = map[string]Type{
	"var": Var,
}

func LookUpKeywordOrIdentifier(s string) Type {
	if t, ok := keywords[s]; ok {
		return t
	}

	return Ident
}
