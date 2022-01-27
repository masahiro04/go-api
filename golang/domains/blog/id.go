package blog

import "strconv"

type ID struct {
	Value int
}

func NewId(value int) (ID, error) {
	return ID{Value: value}, nil
}

func (id ID) String() string {
	return strconv.Itoa(id.Value)
}
