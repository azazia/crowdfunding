package transaction

import (
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int    	`json:"id"`
	Name      string 	`json:"name"`
	Amount    int    	`json:"amount"`
	CreatedAt time.Time	`json:"created_at"`
}

type UserTransactionFormatter struct{
	ID			int			`json:"id"`
	Amount		int			`json:"amount"`
	Status		string		`json:"status"`
	CreatedAt	time.Time	`json:"created_at"`
	Campaign	UserCampaignFormatter `json:"campaign"`
}

type UserCampaignFormatter struct{
	Name		string					`json:"name"`
	FileName	string	`json:"image_url"`
}


func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter{
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter{
	transactionsFormatter := []CampaignTransactionFormatter{}

	for _, transaction := range transactions{
		transactionFormatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, transactionFormatter)
	}

	return transactionsFormatter
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter{
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	formatter.Campaign = UserCampaignFormatter{}
	formatter.Campaign.Name = transaction.Campaign.Name
	formatter.Campaign.FileName = ""
	if len(transaction.Campaign.CampaignImages)>0 {
		formatter.Campaign.FileName = transaction.Campaign.CampaignImages[0].FileName
	}
	
	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter{
	userTransactionsFormatter := []UserTransactionFormatter{}

	for _, transaction := range transactions{
		userTransactionFormatter := FormatUserTransaction(transaction)
		userTransactionsFormatter = append(userTransactionsFormatter, userTransactionFormatter)
	}

	return userTransactionsFormatter
}