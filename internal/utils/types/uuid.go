package types

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func ParseID(id string) (ID, error) {
	return uuid.Parse(id)
}