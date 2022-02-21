package campaign

import "time"

type Campaign struct {
	ID            int
	UserID        int
	Name          string
	ShortDesc     string
	GoalAmount    int
	CurrentAmount int
	Desc          string
	Perks         string
	BackerCount   int
	Slug          string
	CreatedAt	  time.Time
	UpdatedAt	  time.Time
	CampaignImages	[]CampaignImage
}

type CampaignImage struct{
	ID				int
	CampaignID		int
	FileName		string
	IsPrimary		bool
	CreatedAt		time.Time
	UpdatedAt		time.Time
}