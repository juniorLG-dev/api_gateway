package vo

import (
	"github.com/google/uuid"
)

type ID struct {
	Value string
}

func NewID() *ID {
	return &ID{
		Value: uuid.New().String(),
	}
}