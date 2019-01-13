package evaluator

import (
	"fmt"
	"log"

	"github.com/tomocy/kinako/object"
)

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
	env := make(Environment)
	for name, obj := range builtins {
		env[name] = obj
	}

	return env
}

func (e Environment) Set(name string, obj object.Object) error {
	if _, ok := builtins[name]; ok {
		log.Println(builtins)
		return fmt.Errorf("cannot assign to %s", name)
	}

	e[name] = obj
	return nil
}
