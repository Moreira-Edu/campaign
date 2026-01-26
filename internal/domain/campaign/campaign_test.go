package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign x"
	content  = "body hi :D"
	contacts = []string{"emailum@email.com", "emaildois@email.com"}
	fake     = faker.New()
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

func Test_NewCampaign_MustValidateMinName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required with min 5", err.Error())
}
func Test_NewCampaign_MustValidateMaxName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_MustValidateMinContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required with min 5", err.Error())
}
func Test_NewCampaign_MustValidateMinContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required with min 1", err.Error())
}
func Test_NewCampaign_MustValidateMaxContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(2000), contacts)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateEmails(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"gg"})

	assert.Equal("email is invalid", err.Error())
}
