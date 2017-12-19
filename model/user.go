package model

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
