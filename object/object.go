package object

type Object interface {
	object()
}

type Integer struct {
	Value int64
}

func (i Integer) object() {
}
