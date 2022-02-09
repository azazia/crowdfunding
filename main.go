package main

import (
	// "fmt"
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
	
	input := user.CheckEmailInput{
		Email: "",
	}
	IsAvailable, err := userService.IsEmailAvailable(input)
	if err != nil{
		println("error")
	}


	if IsAvailable == true{
		println("email available")
	}else{
		println("email not available")
	}

	userHandler := handler.NewUserHandler(userService)

	// membuat router
	router := gin.Default()
	// membuat grup
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	// api.GET("/users/fetch", userHandler.FetchUser)

	router.Run()

}

