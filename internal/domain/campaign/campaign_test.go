package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign x"
	content  = "body"
	contacts = []string{"emailum@email.com", "emaildois@email.com"}
)

func Test_createNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))

}

func Test_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaing.Id)

}

func Test_createdOnIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaing.CreatedOn)

}

func Test_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	campaing, _ := NewCampaign(name, content, contacts)

	now := time.Now().Add(-time.Minute)

	assert.Greater(campaing.CreatedOn, now)

}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateEmails(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("emails are required", err.Error())
}
