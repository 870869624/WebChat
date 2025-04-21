package models

import (
	"time"

	"gorm.io/gorm"
)

type Friends struct {
	gorm.Model
	ID        uint
	UserId    int
	ToUserId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
