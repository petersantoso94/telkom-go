package user

import (
	"time"
)

type TelinUser struct {
	ID               uint64    `gorm:"primary_key"`
	Email            string  `gorm:"type:varchar(50);not null"`
	Password         string    `gorm:"type:varchar(100);not null"`
	Position         string    `gorm:"type:varchar(50)"`
	LockIP         string    `gorm:"type:varchar(50)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
