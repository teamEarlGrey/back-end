package main

import (
	"github.com/Kantaro0829/go-gin-test/handler"
	"github.com/Kantaro0829/go-gin-test/infra"
	"github.com/gin-gonic/gin"
)

func main() {

	infra.DBInit()
	router := gin.Default()
	user := router.Group("/user")
	{
		user.GET("/get", handler.Getting)
		user.PUT("/reg", handler.UserReg)
		user.POST("/login", handler.UserLogin)
	}
	//add to main
	//commit from feature/kantaro branch
	//2nd aaa
	//aaa
	//test branch

	router.Run(":3000")

}