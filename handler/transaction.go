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
	transactionService 	transaction.Service
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
	
	formatTransaction := transaction.FormatUserTransactions(transactions)

	response := helper.APIResponse("List of user's transactions", http.StatusOK, "success", formatTransaction)
	c.JSON(http.StatusOK, response)
}

// input dari user
// tangkap input dan dimapping ke struct input
// panggil service untuk transaksi, panggil sistem midtrans
// panggil repository create new transaction data
func (h *transactionHandler) CreateTransaction(c *gin.Context){
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create new transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.transactionService.CreateTransaction(input)
	if err != nil{
		response := helper.APIResponse("Failed to create new transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatTransaction(newTransaction)

	response := helper.APIResponse("success to create new transaction", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// yang mengkonsumsi adalah midtrans bukan frontend
func (h *transactionHandler) GetNotification(c *gin.Context){
	var input transaction.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		response := helper.APIResponse("Failed to get notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.transactionService.PaymentProcess(input)
	if err != nil{
		response := helper.APIResponse("Failed to get notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	c.JSON(http.StatusOK, input)
}