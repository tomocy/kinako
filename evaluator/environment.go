package evaluator

import "github.com/tomocy/kinako/object"

type Environment map[string]object.Object

func NewEnvironment() Environment {
	return Environment{
		"true": &object.Boolean{
			Value: true,
		},
		"false": &object.Boolean{
			Value: false,
		},
	}
}
