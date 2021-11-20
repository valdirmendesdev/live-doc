package survey

import (
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type Repository interface {
	Add(s *entities.Survey) (*entities.Survey, error)
	GetById(id types.ID) (*entities.Survey, error)
	All() ([]entities.Survey, error)
}
