package campaign

import (
	"testing"
	"gorm.io/gorm"
	"log"

	"gorm.io/driver/postgres"
)

func TestRepository(t *testing.T){
	dsn := "host=localhost user=postgres password=nciruuxz dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}
	campaignRepository := NewRepository(db)
	campaign := Campaign{
		UserID: 33,
		Name: "donasi gempa",
		ShortDescription: "donasi",
		GoalAmount: 500000000,
	}

	campaignRepository.Save(campaign)
}