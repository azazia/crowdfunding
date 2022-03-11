package handler

import (
	"net/http"
	"strconv"
	"website-crowdfunding/campaign"
	"website-crowdfunding/helper"
	"website-crowdfunding/user"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler
// handler ke service
// service menentukan repository mana yg di-call
// repository: GetAll, GetByID
// db
type campaignHandler struct{
	campaignService campaign.Service
} 

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler{
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context){
	// konversi string ke int dari hasil query user_id
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatCampaign := campaign.FormatCampaigns(campaigns)

	response := helper.APIResponse("List of campaign", http.StatusOK, "success", formatCampaign)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context){
	// handler : mapping id yang di url ke struct input => service, call formatter
	// service: inputnya struct input => menangkap id di url, memanggil repo
	// repository : get campaign id
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	detailCampaign, err := h.campaignService.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	response := helper.APIResponse("Success to get detail of campaign", http.StatusOK, "error", campaign.FormatDetailCampaign(detailCampaign))
	c.JSON(http.StatusOK, response)
}

// tangkap parameter dari user ke input struct
// ambil current user dari jwt/handler
// panggil service, parameternya input struct(dan buat slug)
// panggil repository,untuk simpan data campaign baru

func(h *campaignHandler) CreateCampaign(c *gin.Context){
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create new campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil{
		response := helper.APIResponse("Failed to create new campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create new campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))	
	c.JSON(http.StatusOK, response)
}

// update
// user memasukan input
// handler
// mapping dari input ke input struct
// input dari user, dan input dari uri(passing dari service)
// service
// repository
