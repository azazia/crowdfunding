package main

import (
	// "fmt"
	"fmt"
	"log"
	"website-crowdfunding/handler"
	"website-crowdfunding/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	dsn := "host=localhost user=postgres password=nciruuxz dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	findID, err := userRepository.FindByID(1)
	if err != nil{
		fmt.Println(err.Error())
	}

	fmt.Println(findID.Name)

	userHandler := handler.NewUserHandler(userService)

	// membuat router
	router := gin.Default()
	// membuat grup
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.EmailAvaliability)

	router.Run()

}

