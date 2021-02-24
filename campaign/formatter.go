package campaign

type CampaignFormatter struct {
	ID               int    `json="id"`
	UserID           int    `json="user_id"`
	Name             string `json="name"`
	ShortDescription string `json="short_description"`
	ImageUrl         string `json="image_url"`
	GoalAmount       int    `json="goal_amount"`
	CurrentAmount    int    `json="current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormat := []CampaignFormatter{}
	for _, campaign := range campaigns {
		campaignFormat := FormatCampaign(campaign)
		campaignsFormat = append(campaignsFormat, campaignFormat)
	}

	return campaignsFormat
}