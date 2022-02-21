package handler

import (
	"net/http"
	"strconv"
	"website-crowdfunding/campaign"
	"website-crowdfunding/helper"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler
// handler ke service
// service menentukan repository mana yg di-call
// repository: GetAll, GetByID
// db
type campaignHandler struct{
	service campaign.Service
} 

func NewCampaignHandler(service campaign.Service) *campaignHandler{
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context){
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaign", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}