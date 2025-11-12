package campaign

import (
	"emailn/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}

	NewCampaign := contract.NewCampaign{
		Name:    "Test",
		Content: "body",
		Emails:  []string{"teste@email.com"},
	}

	id, err := service.Create(NewCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}
