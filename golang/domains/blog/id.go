package blog

import "strconv"

type ID struct {
	value int
}

func NewId(value int) (ID, error) {
	return ID{value: value}, nil
}

func (id ID) String() string {
	return strconv.Itoa(id.value)
}

func (id ID) Value() int {
	return id.value
}
