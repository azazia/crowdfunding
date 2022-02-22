package campaign

type CampaignFormatter struct {
	ID					int		`json:"id"`
	Name				string	`json:"name"`
	ShortDesc			string	`json:"short_description"`
	FileName			string	`json:"image_url"`
	GoalAmount			int		`json:"goal_amount"`
	CurrentAmount		int		`json:"current_amount"`
	UserID				int		`json:"user_id"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter{
	formatter := CampaignFormatter{}
	formatter.ID = campaign.ID
	formatter.Name = campaign.Name
	formatter.ShortDesc = campaign.ShortDescription
	formatter.GoalAmount = campaign.GoalAmount
	formatter.CurrentAmount = campaign.CurrentAmount
	formatter.UserID = campaign.UserID
	formatter.FileName = ""

	if len(campaign.CampaignImages) > 0{
		formatter.FileName = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter{
	// jika tidak ada data yg ditemukan akan mengembalikan array kosong
	if len(campaigns) == 0{
		return []CampaignFormatter{}
	}

	var campaignsFormatter []CampaignFormatter

	for _, campaign := range campaigns{
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}