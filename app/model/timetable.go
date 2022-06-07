package model

import (
	"gorm.io/gorm"
)

type Timetable struct {
	gorm.Model
	No        uint16 `gorm:"primaryKey"`
	class     string
	RoomNo    uint16
	Name      string
	Youbi     string
	TeacherNo uint16
	TimeNo    string
}
