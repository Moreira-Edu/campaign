package database

import (
	"emailn/internal/domain/campaign"
	"errors"
)

type CampaignRepository struct {
	Campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.Campaigns = append(c.Campaigns, *campaign)
	return errors.New("an error")
}
