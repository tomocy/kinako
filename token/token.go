package token

type Token struct {
	Type    Type
	Literal string
}

type Type string

const (
	Unknown Type = "Unknown"

	EOF = "EOF"

	Plus = "Plus"

	Integer = "Integer"
)

var types = map[string]Type{
	"+": Plus,
}

func LookUpType(s string) Type {
	if t, ok := types[s]; ok {
		return t
	}

	return Unknown
}
