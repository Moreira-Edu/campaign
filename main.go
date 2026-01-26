package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	contacts := []campaign.Contact{{Email: ""}, {Email: "testgo@email.com"}}
	campaign := campaign.Campaign{Contacts: contacts}
	validate := validator.New()
	err := validate.Struct(campaign)

	if err == nil {
		println("nenhum erro")
	} else {
		validationErr := err.(validator.ValidationErrors)

		for _, v := range validationErr {
			switch v.Tag() {
			case "required":
				println(v.StructField() + " is required")
			case "min":
				println(v.StructField() + " is required with min " + v.Param())
			case "max":
				println(v.StructField() + " is required with max " + v.Param())
			case "email":
				println(v.StructField() + " is invalid")
			}
		}
	}

}
