package handler

import (
	"net/http"

	//"github.com/Kantaro0829/go-gin-test/auth"
	//"github.com/Kantaro0829/go-gin-test/infra"
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

	// db := infra.DBInit()
	// timers := []model.Timer{}

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

	c.JSON(http.StatusOK, gin.H{"message": buildingNumAndFloor})
}
