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