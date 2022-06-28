package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

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
	fmt.Println(buildingNumAndFloor)
	buildingAndFloor := strconv.FormatInt(buildingNumAndFloor, 10)
	buildingAndFloor = buildingAndFloor + "%"

	db := infra.DBInit()

	today := time.Now()
	fmt.Println(today)
	dayOfWeek := today.Weekday().String() // 曜日の取得
	fmt.Println(dayOfWeek)
	fmt.Println(reflect.TypeOf(dayOfWeek))
	fmt.Println(dayOfWeek[0:3])

	rooms := []model.Room{}
	db.Order("room_no").
		Select("room_no").
		Where("room_no LIKE ?", buildingAndFloor).
		Find(&rooms)

	roomSlice := createRoomsSlice(rooms)
	fmt.Println(roomSlice)

	roomResults := []model.RoomResult{}
	result := db.Order("timetables.room_no, timetables.time_no").Table("timetables").
		Select("timetables.room_no, timetables.time_no, teachers.teacher_name, timetables.subject_name").
		Joins("left join teachers on timetables.teacher_no = teachers.teacher_no").
		Where("timetables.room_no LIKE ? AND timetables.youbi = ?", buildingAndFloor, "Fri").
		Scan(&roomResults)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}
	fmt.Println("結合後のテーブル")
	for _, v := range roomResults {
		fmt.Println("-------------------------------------------------------------------")
		fmt.Printf("%v, %v, %v, %v\n", v.RoomNo, v.TimeNo, v.TeacherName, v.SubjectName)
	}

	roomInfo := createRoomInfoJson(roomResults)
	reservationInfo := createReservationInfoJson()
	response := AllInfo{NormalInfo: roomInfo, ReservationInfo: reservationInfo}
	c.JSON(http.StatusOK, response)
}

type Class struct {
	TimeNo      string
	TeacherName string
	SubjectName string
}

type AllInfo struct {
	NormalInfo      map[string][]Class
	ReservationInfo map[string]string
}

func createReservationInfoJson() map[string]string {
	reservationInfos := make(map[string]string)
	reservationInfos["reservation"] = "予約"
	return reservationInfos
}

func createRoomInfoJson(roomInfos []model.RoomResult) map[string][]Class {
	//各教室の状況を格納するJson配列を作成する
	var currentRoomNo string                  //同じ教室番号を配列に分割するために判断する変数
	eachRoomInfos := make(map[string][]Class) //最終的に出力したいJsonの型宣言
	roomInfo := []Class{}

	for i, v := range roomInfos {
		fmt.Printf("%v, %v, %v, %v\n", v.RoomNo, v.TimeNo, v.TeacherName, v.SubjectName)
		if i == 0 {
			//ループの最初は変数currentRoomに代入
			currentRoomNo = v.RoomNo
		}

		if currentRoomNo != v.RoomNo {

			//以前の教室番号と違う教室番号の場合新しい連想配列を作る
			eachRoomInfos[currentRoomNo] = roomInfo

			roomInfo = []Class{} //各教室1~5限情報をを格納する配列の初期化
			currentRoomNo = v.RoomNo
		}

		//各教室の1〜5限目の情報を格納する配列に値を入れる
		roomInfo = append(roomInfo, Class{
			TimeNo:      v.TimeNo,
			TeacherName: v.TeacherName,
			SubjectName: v.SubjectName,
		})

	}
	//最後だけfor文が回らないので
	eachRoomInfos[currentRoomNo] = roomInfo
	fmt.Println("------------------出来上がったJson---------------------")
	fmt.Println(eachRoomInfos)

	return eachRoomInfos

}

func createRoomsSlice(rooms []model.Room) []string {
	roomSlice := []string{}
	for _, v := range rooms {
		roomSlice = append(roomSlice, v.RoomNo)
	}
	return roomSlice
}
