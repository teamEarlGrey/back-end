package model

import (
	"gorm.io/gorm"
)

type Timetable struct {
	gorm.Model
	No          uint16 `gorm:"primaryKey"`
	RoomNo      string
	SubjectName string
	Youbi       string
	TeacherNo   uint16
	TimeNo      string
}
