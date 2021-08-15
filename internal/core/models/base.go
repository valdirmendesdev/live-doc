package models

import "github.com/google/uuid"

type ID = uuid.UUID

func NewUUID() ID {
	return uuid.New()
}