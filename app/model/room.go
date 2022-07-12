package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomNo     string `gorm:"primaryKey"`
	Memo       string
	IsDetected bool
}
