package handler

import (
	"fmt"
	"net/http"

	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/Kantaro0829/go-gin-test/model"
	"github.com/gin-gonic/gin"
)

func ReservationInfo(c *gin.Context) {
	// TODO: /modelにreservationsから取ってくる型宣言をするgoファイルを作る

	// db.goからmysql内のDBにアクセスしてる
	db := infra.DBInit()
	// reservation = model.

	// TODO: 予約テーブルにデータを挿入
	// TODO: 予約テーブルからデータを取得(教室番号、開始時間、終了時間)
	rese := []model.Reservation{}
	result := db.Order("reservations.room_no, reservations.s_time, reservavtions.e_time").
		Table("reservations").
		Select("reservations.room_no, reservations.s_time, reservavtions.e_time").
		Scan(&rese)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": 400})
		return
	}
	json := createReservationInfoJson(rese)
	c.JSON(http.StatusOK, json)
}

type reseClass struct {
	RoomNo string
	Stime  string
	Etime  string
}

func createReservationInfoJson(reseInfos []model.Reservation) map[string][]reseClass {
	//各教室の予約状況を格納するJson配列を作成
	var roomNo string
	reseRoomInfo := make(map[string][]reseClass)
	reseInfo := []reseClass{}

	for i, v := range reseInfos {
		fmt.Printf("%v, %v, %v\n", v.RoomNo, v.STime, v.ETime)

		if i == 0 {
			roomNo = v.RoomNo
		}

		if roomNo != v.RoomNo {
			reseRoomInfo[roomNo] = reseInfo

			reseInfo = []reseClass{}
			roomNo = v.RoomNo
		}

		//各教室の予約状況を配列に格納する
		reseInfo = append(reseInfo, reseClass{
			RoomNo: v.RoomNo,
			Stime:  v.STime,
			Etime:  v.ETime,
		})
	}
	// reseRoomInfo[roomNo] = reseInfo
	fmt.Println("------------------出来上がったJson---------------------")
	fmt.Println(reseRoomInfo)
	return reseRoomInfo

}
