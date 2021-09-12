package entities

import (
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        types.ID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4();"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
