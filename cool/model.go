package cool

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primaryKey"`
	CreatedAt time.Time  `gorm:"index"`
	UpdatedAt time.Time  `gorm:"index"`
	DeletedAt *time.Time `gorm:"index"`
}
