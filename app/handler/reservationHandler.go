package handler

import (
	"fmt"
	"net/http"

	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/Kantaro0829/go-gin-test/model"
	"github.com/gin-gonic/gin"
)

// 予約データを取得
func ReservationInfo(c *gin.Context) {
	// TODO: /modelにreservationsから取ってくる型宣言をするgoファイルを作る

	// db.goからmysql内のDBにアクセス
	db := infra.DBInit()
	// reservation = model.

	// TODO: 予約テーブルからデータを取得(教室番号、開始時間、終了時間)
	rese := []model.Reservation{}
	// state_noが1(承認済み)の予約のみ取り出す
	result := db.Select("room_no, s_time, e_time").
		Where("state_no = 1").
		First(&rese)
		// Scan(&rese)

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

// reservationsテーブルに格納されている予約をJson形式に書き換えている
func createReservationInfoJson(reseInfos []model.Reservation) map[string][]reseClass {
	//各教室の予約状況を格納するJson配列を作成
	var roomNo string
	reseRoomInfo := make(map[string][]reseClass)
	reseInfo := []reseClass{}

	for i, v := range reseInfos {
		fmt.Printf("%v, %v, %v\n", v.RoomNo, v.STime, v.ETime)

		// ループの最初をroomNoに代入
		if i == 0 {
			roomNo = v.RoomNo
		}

		if roomNo != v.RoomNo {
			// 前の教室番号と違う教室番号の場合新しい連想配列を作成
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
	// ↓件数が増えたらいるかも
	reseRoomInfo[roomNo] = reseInfo
	fmt.Println("------------------出来上がったJson---------------------")
	fmt.Println(reseRoomInfo)
	return reseRoomInfo

}
