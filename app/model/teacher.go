package model

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	TeacherNo   uint16
	TeacherName string
	PerNo       uint8
}
