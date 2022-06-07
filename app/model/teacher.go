package model

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	TeacherNo uint16
	Name      string
	PerNo     uint8
}
