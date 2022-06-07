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
	for i, v := range teachers {
		fmt.Printf("%v回目\n", i)
		fmt.Printf("%v, %v\n", v.TeacherNo, v.Name)
	}

	fmt.Println("teachersテーブル")
	for i, v := range timetables {
		fmt.Printf("%v回目\n", i)
		fmt.Printf("%v, %v\n", v.No, v.TimeNo)
	}

	roomResult := []model.RoomResult{}
	result := db.Model(timetables).
		Select("timetables.room_no, timetables.class, timetables.time_no, teachers.name, timetables.name").
		Joins("left join teachers on teachers.teacher_no = timetables.teacher_no").Scan(&model.RoomResult{})

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}

	for i, v := range roomResult {
		fmt.Printf("%v回目\n", i)
		fmt.Printf("%v, %v, %v, %v, %v\n", v.RoomNum, v.Class, v.TimeNo, v.Teacher, v.Subject)
	}
	// if result := db.Find(&timers); result.Error != nil {
	// 	fmt.Println("データ取得失敗")
	// }
	// // fmt.Println(users)
	// for i, v := range timers {
	// 	fmt.Printf("%v回目", i)
	// 	fmt.Println(v.TimeNo)
	// 	fmt.Println(v.STime)
	// 	fmt.Println(v.ETime)
	// 	fmt.Println(v.ID)
	// }

	//

	c.JSON(http.StatusOK, gin.H{"message": buildingNumAndFloor})
}
