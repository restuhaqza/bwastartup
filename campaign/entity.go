package campaign

import "time"

type Campaign struct {
	ID               string
	UserID           string
	Name             string
	ShortDescription string
	Perks            string
	BackerCount      int
	GoalAmount       int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
