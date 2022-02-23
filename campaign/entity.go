package campaign

import (
	"time"
	"website-crowdfunding/user"
)

type Campaign struct {
	ID            		int
	UserID        		int
	Name          		string
	ShortDescription	string
	GoalAmount    		int
	CurrentAmount 		int
	Description         string
	Perks         		string
	BackerCount   		int
	Slug          		string
	CreatedAt	  		time.Time
	UpdatedAt	  		time.Time
	CampaignImages		[]CampaignImage
	User				user.User
}

type CampaignImage struct{
	ID				int
	CampaignID		int
	FileName		string
	IsPrimary		bool
	CreatedAt		time.Time
	UpdatedAt		time.Time
}