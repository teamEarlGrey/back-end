package handler

import (
	"fmt"
	"net/http"

	//"github.com/Kantaro0829/go-gin-test/auth"
	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/Kantaro0829/go-gin-test/model"

	//"github.com/Kantaro0829/go-gin-test/json"
	//"github.com/Kantaro0829/go-gin-test/model"
	//"github.com/gin-gonic/gin"
	//"golang.org/x/crypto/bcrypt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRoomInfo(c *gin.Context) {

	roomNumStr := c.Param("roomNo")
	//fmt.Println(roomNumStr)
	roomNum, _ := strconv.ParseInt(roomNumStr, 10, 16)
	//roomNumberの上二桁だけ切り取り
	buildingNumAndFloor := roomNum / 100

	db := infra.DBInit()
	teachers := []model.Teacher{}
	timetables := []model.Timetable{}

	if result := db.Find(&teachers); result.Error != nil {
		fmt.Println("データ取得失敗")
	}

	if result := db.Find(&timetables); result.Error != nil {
		fmt.Println("データ取得失敗")
	}

	fmt.Println("teachersテーブル")
	for _, v := range teachers {
		fmt.Printf("-----------------------------")
		fmt.Printf("%v, %v, %v\n", v.TeacherNo, v.TeacherName, v.PerNo)
	}

	fmt.Println("timetableテーブル")
	for _, v := range timetables {
		fmt.Println("-------------------------")
		fmt.Printf("%v, %v, %v, %v, %v, %v\n", v.No, v.RoomNo, v.SubjectName, v.Youbi, v.TeacherNo, v.TimeNo)
	}

	roomResults := []model.RoomResult{}
	result := db.Table("timetables").
		Select("timetables.room_no, timetables.time_no, teachers.teacher_name, timetables.subject_name").
		Joins("left join teachers on timetables.teacher_no = teachers.teacher_no").Scan(&roomResults)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}
	fmt.Println("結合後のテーブル")
	for _, v := range roomResults {
		fmt.Println("-------------------------------------------------------------------")
		fmt.Printf("%v, %v, %v, %v\n", v.RoomNo, v.TimeNo, v.TeacherName, v.SubjectName)
	}

	c.JSON(http.StatusOK, gin.H{"message": buildingNumAndFloor})
}
