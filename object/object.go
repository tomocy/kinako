package object

import (
	"fmt"
	"strconv"
)

type Object interface {
	object()
}

type Integer struct {
	Value int64
}

func (i Integer) object() {
}

func (i Integer) String() string {
	return strconv.FormatInt(i.Value, 10)
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

func (e Error) object() {
}

func (e Error) String() string {
	return e.Message
}
