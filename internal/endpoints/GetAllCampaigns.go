package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGetAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaigns, err := h.CampaignService.Repository.GetAll()
	return campaigns, 200, err
}
