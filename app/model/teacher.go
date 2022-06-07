package model

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	TeacherNo uint16 `gorm:"primaryKey"`
	Name      string
	PerNo     uint8
}
