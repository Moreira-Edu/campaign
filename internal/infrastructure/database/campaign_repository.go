package database

import (
	"emailn/internal/domain/campaign"
)

type CampaignRepository struct {
	Campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.Campaigns = append(c.Campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) GetAll() []campaign.Campaign {
	return c.Campaigns
}
