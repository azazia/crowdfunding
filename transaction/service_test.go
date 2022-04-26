package transaction

import (
	"log"
	"testing"
	"website-crowdfunding/campaign"
	"website-crowdfunding/payment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPaymentProcess(t *testing.T){
	dsn := "host=localhost user=postgres password=nciruuxz dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}

	transactionRepository := NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	paymentService := payment.NewService()

	transactionService := NewService(transactionRepository, campaignRepository, paymentService)
	input := TransactionNotificationInput{}
	input.OrderID = "17"
	input.TransactionStatus = "settlement"

	err = transactionService.PaymentProcess(input)
	if err != nil{
		log.Fatal(err)
	}

}