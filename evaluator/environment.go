package evaluator

import "github.com/tomocy/kinako/object"

type Environment map[string]object.Object

var builtins = Environment{
	"true": &object.Boolean{
		Value: true,
	},
	"false": &object.Boolean{
		Value: false,
	},
}

func NewEnvironment() Environment {
	return builtins
}
