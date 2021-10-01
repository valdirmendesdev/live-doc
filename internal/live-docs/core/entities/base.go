package entities

import (
	"time"

	"github.com/valdirmendesdev/live-doc/internal/utils/types"
	"gorm.io/gorm"
)

type Model struct {
	// ID        types.ID `gorm:"column:id;primaryKey;type:uuid;default:default:uuid_generate_v4()"`
	ID        types.ID `gorm:"column:id;primaryKey;type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
