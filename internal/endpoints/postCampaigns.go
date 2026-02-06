package endpoints

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var Request contract.NewCampaign

	err := render.DecodeJSON(r.Body, &Request)

	id, err := h.CampaignService.Create(Request)

	if err != nil {

		if errors.Is(err, internalerrors.ErrInernal) {
			render.Status(r, 500)
		} else {
			render.Status(r, 400)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})

}
