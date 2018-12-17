package token

type Token struct {
	Type    Type
	Literal string
}

type Type string

const (
	Unknown Type = "Unknown"

	EOF = "EOF"

	Plus  = "Plus"
	Minus = "Minus"

	Integer = "Integer"
)

var types = map[string]Type{
	"+": Plus,
	"-": Minus,
}

func LookUpType(s string) Type {
	if t, ok := types[s]; ok {
		return t
	}

	return Unknown
}
