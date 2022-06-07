package model

import (
	"gorm.io/gorm"
)

type Timetable struct {
	gorm.Model
	No        uint16 `gorm:"primaryKey"`
	Class     string
	RoomNo    string
	Name      string
	Youbi     string
	TeacherNo uint16
	TimeNo    string
}
