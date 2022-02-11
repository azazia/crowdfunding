package handler

import (
	"net/http"
	"website-crowdfunding/helper"
	"website-crowdfunding/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct diatas kita passing sebagai parameter service

	var input user.RegisterUserInput

	// mengubah json ke struct RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidationError(err)

		// map untuk menampung list error
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// mengisi nilai input ke RegisterUser
	newUser, err := h.userService.RegisterUser(input)
	if err != nil{
		response := helper.APIResponse("Register account failed", 400, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatUser := user.FormatUser(newUser, "initoken")

	response := helper.APIResponse("Account has been registered", 200, "success", formatUser)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context){
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidationError(err)
	
		errorMessage := gin.H{
			"error": errors,
		}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}	

	loggedInUser, err := h.userService.LoginUser(input)
	if err != nil{
		errorMessage := gin.H{
			"error": err.Error(),
		}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return	
	}

	formatUser := user.FormatUser(loggedInUser, "initoken")

	response := helper.APIResponse("Login success", http.StatusOK, "success", formatUser)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) EmailAvaliability(c *gin.Context){
	// input dari user
	// input email dimapping ke struct input
	// struct input dipassing keservice
	// service akan memanggil repository - email sudah ada atau belum
	// repository - db

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": errors,
		}
		response := helper.APIResponse("email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// check apakah email sudah ada
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{
			"error": "server error",
		}
		response := helper.APIResponse("email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "email has been registered"
	if isEmailAvailable{
		metaMessage = "email is available"
	}
	response := helper.APIResponse(metaMessage, http.StatusUnprocessableEntity, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context){
	//  input dari user
	// simpan gambar ke folder "/images"
	// di service panggil repo
	// JWT
	// repo ambil data user berdasarkan ID
	// repo update data user simpan ke  folder
}