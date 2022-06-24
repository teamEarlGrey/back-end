package handler

import (
	"net/http"

	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/Kantaro0829/go-gin-test/model"
	"github.com/gin-gonic/gin"
)

func ReservationInfo(c *gin.Context){
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

func createReservationInfoJson(reseInfo []model.Reservation) map[string][]Class{
	return
}


