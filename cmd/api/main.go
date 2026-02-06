package main

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	internalerrors "emailn/internal/internalErrors"
	"errors"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var Request contract.NewCampaign

		err := render.DecodeJSON(r.Body, &Request)

		id, err := service.Create(Request)

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

	})

	http.ListenAndServe(":3000", r)
}
