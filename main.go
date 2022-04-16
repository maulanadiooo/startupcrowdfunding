package main

import (
	"log"
	"startup/handler"
	"startup/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("api/v1")

	api.POST("/jancuk", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()

	/////// alur kerja golang
	// input dari user
	// handler, mapping input dari user menjadi struct input
	// service : melakukan mapping dari struct input ke struct User
	// repository
	// db
	/////// alur kerja golang untuk register user
}
