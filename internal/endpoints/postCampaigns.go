package endpoints

import (
	"emailn/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var Request contract.NewCampaign

	err := render.DecodeJSON(r.Body, &Request)

	id, err := h.CampaignService.Create(Request)

	return map[string]string{"id": id}, 201, err
}
