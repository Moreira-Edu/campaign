package campaign

import (
	"emailn/internal/contract"
	"errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(NewCampaign contract.NewCampaign) (string, error) {

	return "", errors.New("teste")

}
