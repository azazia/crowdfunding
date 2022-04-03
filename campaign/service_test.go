package campaign

import (
	"testing"
	"gorm.io/gorm"
	"log"

	"gorm.io/driver/postgres"
)

func TestService(T *testing.T){
	dsn := "host=localhost user=postgres password=nciruuxz dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}

	campaignRepository := NewRepository(db)
	campaignService := NewService(campaignRepository)

	input := CreateCampaignInput{}
	input.Name = "penggalangan dana startup"
	input.ShortDescription = "short"
	input.Description = "loooooong"
	input.GoalAmount = 100000000
	input.Perks = "satu, dua, tiga"
	

	_, err = campaignService.CreateCampaign(input)
	if err != nil{
		log.Fatal(err)
	}

}

// func TestServiceUpdateCampaign(T *testing.T){
// 	dsn := "host=localhost user=postgres password=nciruuxz dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil{
// 		log.Fatal(err)
// 	}

	
// 	campaignRepository := NewRepository(db)
// 	campaignService := NewService(campaignRepository)
// 	input := CreateCampaignInput{}
// 	input.Name = "penggalangan dana startup"
// 	input.ShortDescription = "short"
// 	input.Description = "loooooong"
// 	input.GoalAmount = 100000000
// 	input.Perks = "satu, dua, tiga"
	

// 	// _, err = campaignService.Update(12, input)
// 	// if err != nil{
// 	// 	log.Fatal(err)
// 	// }
// }