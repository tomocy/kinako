package object

import "strconv"

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
