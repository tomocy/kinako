package object

import (
	"fmt"
)

type Object interface {
	object()
}

type Integer struct {
	Value int64
}

func (o Integer) object() {
}

func (o Integer) String() string {
	return fmt.Sprintf("%d", o.Value)
}

type Boolean struct {
	Value bool
}

func (o Boolean) object() {
}

func (o Boolean) String() string {
	return fmt.Sprintf("%t", o.Value)
}

type Error struct {
	Message string
}

func (o Error) object() {
}

func (o Error) String() string {
	return o.Message
}
