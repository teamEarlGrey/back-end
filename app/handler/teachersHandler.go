package handler

import (
	"fmt"
	"net/http"

	"github.com/Kantaro0829/go-gin-test/model"
	"gorm.io/gorm"

	"github.com/Kantaro0829/go-gin-test/auth"
	"github.com/Kantaro0829/go-gin-test/json"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func TeacherReg(c *gin.Context) {

	var teacherJson json.RegTeacher

	//上で宣言した構造体にJsonをバインド。エラーならエラー処理を返す
	if err := c.ShouldBindJSON(&teacherJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacherName := teacherJson.TeacherName
	password := teacherJson.Password
	mail := teacherJson.TeacherMail
	//パスワードハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic("failed to hash password")
	}
	fmt.Println("以下ハッシュ化されたパスワード")
	fmt.Println(hashedPassword)
	fmt.Println(string(hashedPassword))

	teacher := model.Teacher{TeacherName: teacherName, PerNo: 1, Password: string(hashedPassword), Mail: mail}
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})

	//ユーザ情報登録
	if result := tx.Create(&teacher); result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusConflict, gin.H{"status": 400, "message": "データの登録に失敗しました"})
		return
	}
	tx.Commit()
	fmt.Println("登録されたパスワード")
	fmt.Println(teacher.Password)
	fmt.Println(teacher.PerNo)
	fmt.Println(teacher.TeacherName)
	fmt.Println(teacher.ID)

	//IDと権限番号で元にJWTを発行
	token := auth.CreateTokenString(
		teacher.ID,
		teacher.Mail, //teacher.PerNo,
	)

	c.JSON(http.StatusOK, gin.H{"message": "data was inserted", "token": token})

}
