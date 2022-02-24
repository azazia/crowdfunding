package campaign

import "strings"

type CampaignFormatter struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ShortDesc     string `json:"short_description"`
	FileName      string `json:"image_url"`
	GoalAmount    int    `json:"goal_amount"`
	CurrentAmount int    `json:"current_amount"`
	Slug          string `json:"slug"`
	UserID        int    `json:"user_id"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{}
	formatter.ID = campaign.ID
	formatter.Name = campaign.Name
	formatter.ShortDesc = campaign.ShortDescription
	formatter.GoalAmount = campaign.GoalAmount
	formatter.CurrentAmount = campaign.CurrentAmount
	formatter.Slug = campaign.Slug
	formatter.UserID = campaign.UserID
	formatter.FileName = ""

	if len(campaign.CampaignImages) > 0 {
		formatter.FileName = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	ShortDesc     string   `json:"short_description"`
	FileName      string   `json:"image_url"`
	GoalAmount    int      `json:"goal_amount"`
	CurrentAmount int      `json:"current_amount"`
	Slug          string   `json:"slug"`
	UserID        int      `json:"user_id"`
	Desc          string   `json:"description"`
	Perks         []string `json:"perks"`
	User		  CampaignUserFormatter	`json:"user"`
	Images		  []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name		string		`json:"name"`
	ImageURL	string		`json:"image_url"`	
}

type CampaignImageFormatter struct {
	ImageURL	string	`json:"image_url"`
	IsPrimary	bool	`json:"is_primary"`
}

func FormatDetailCampaign(campaign Campaign) CampaignDetailFormatter {
	formatter := CampaignDetailFormatter{}
	formatter.ID = campaign.ID
	formatter.Name = campaign.Name
	formatter.ShortDesc = campaign.ShortDescription
	formatter.GoalAmount = campaign.GoalAmount
	formatter.CurrentAmount = campaign.CurrentAmount
	formatter.UserID = campaign.UserID
	formatter.Desc = campaign.Description
	formatter.Slug = campaign.Slug

	formatter.FileName = ""
	if len(formatter.FileName) > 0 {
		formatter.FileName = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ","){
		perks = append(perks, strings.TrimSpace(perk))
	}
	formatter.Perks = perks

	formatter.User = CampaignUserFormatter{
		Name: campaign.User.Name,
		ImageURL: campaign.User.AvatarFileName,
	}

	images := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages{
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName
		campaignImageFormatter.IsPrimary = image.IsPrimary

		images = append(images, campaignImageFormatter)
	}
	formatter.Images = images

	return formatter
}