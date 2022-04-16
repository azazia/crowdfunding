package transaction

import (
	"log"
	"testing"
	"website-crowdfunding/campaign"
	"website-crowdfunding/payment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateTransactionService(t *testing.T){
	dsn := "host=localhost user=postgres password=nciruuxz dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	repository := NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	newTransaction := NewService(repository, campaignRepository, payment.NewService())

	input := CreateTransactionInput{}
	input.Amount = 100000
	input.CampaignID = 1
	input.User.ID	= 33
	
	_, err = newTransaction.CreateTransaction(input)
	if err != nil {
		log.Fatal(err)
	}


}