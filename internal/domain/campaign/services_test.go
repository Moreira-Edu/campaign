package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)

	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test de save",
		Content: "body hi :D",
		Emails:  []string{"teste@email.com"},
	}
	repository = new(repositoryMock)
	service    = Service{Repository: repository}
)

func Test_CreateCampaign(t *testing.T) {

	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		}
		if campaign.Content != newCampaign.Content {
			return false
		}
		if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	}),
	).Return(nil)

	service.Create(newCampaign)

	repository.AssertExpectations(t)
}

func Test_Validate_Error_Domain(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)
	assert.NotNil(err)
	assert.Equal("name is required with min 5", err.Error())

}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	newCampaign := contract.NewCampaign{
		Name:    "Test V",
		Content: "Body",
		Emails:  []string{"test@email.com"},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("Internal server error"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.False(errors.Is(internalerrors.ErrInernal, err))
}
