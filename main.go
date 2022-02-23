package main

import (
	"log"
	"net/http"
	"strings"
	"website-crowdfunding/auth"
	"website-crowdfunding/campaign"
	"website-crowdfunding/handler"
	"website-crowdfunding/helper"
	"website-crowdfunding/user"

	"github.com/dgrijalva/jwt-go"
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
	campaignRepository := campaign.NewRepository(db)


	userService := user.NewService(userRepository) //parsing userRepository agar punya akses ke repository
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	// membuat router
	router := gin.Default()
	router.Static("/images", "./images")
	// membuat grup

	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.EmailAvaliability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaign", campaignHandler.GetCampaigns)
	api.GET("/campaign/:id", campaignHandler.GetCampaign)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc{
	return func (c *gin.Context){
		authHeader := c.GetHeader("Authorization")
	
		if !strings.Contains(authHeader, "Bearer"){
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			// batalkan jika gagal autentikasi
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	
		var tokenString string
		// memecah string menjadi slice
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		// validasi token
		token, err := authService.ValidateToken(tokenString)
		if err != nil{
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ambil data dari token
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid{
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ambil ID dari claim
		userID := int(claim["user_id"].(float64))

		// ambil user dari db berdasarkan ID
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// set ccontext dengan key currentuser isinya user
		c.Set("currentUser", user)
	}
}

