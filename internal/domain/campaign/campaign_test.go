package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign x"
	content := "body"
	contacts := []string{"emailum@email.com", "emaildois@email.com"}

	campaing := NewCampaign(name, content, contacts)

	assert.Equal(campaing.Id, "2")
	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))

}
