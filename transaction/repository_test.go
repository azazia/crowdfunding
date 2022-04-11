package transaction

import (
	"log"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCampaignTransactionRepository(t *testing.T) {
	dsn := "host=localhost user=postgres password=nciruuxz dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	campaignTransaction := NewRepository(db)

	campaignTransaction.GetCampaignByID(1)
}

func TestUserTransactionRepository(t *testing.T){
	dsn := "host=localhost user=postgres password=nciruuxz dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	campaignTransaction := NewRepository(db)

	campaignTransaction.GetUserByID(32)
}