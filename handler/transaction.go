package handler

import (
	"net/http"
	"website-crowdfunding/helper"
	"website-crowdfunding/transaction"
	"website-crowdfunding/user"

	"github.com/gin-gonic/gin"
)

// parameter di uri
// tangkap parameter mapping input struct
// panggil service, parameter input struct
// service, panggil repo dengan campaign ID
// repo mencari data transaction suatu campaign

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler{
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context){
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed to get campaign's transactions.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.transactionService.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatTransaction := transaction.FormatCampaignTransactions(transactions)

	response := helper.APIResponse("campaign's transactions", http.StatusOK, "success", formatTransaction)
	c.JSON(http.StatusOK, response)
}

// GetUserTransaction
// handler
// ambil nilai user dari jwt/middleware
// service
// repo => ambil data transaksi (preload campaign)
func (h *transactionHandler) GetUserTransaction(c *gin.Context){
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions , err := h.transactionService.GetTransactionByUserID(userID)
	if err != nil {
		response := helper.APIResponse("error to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of user's transactions", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)
}