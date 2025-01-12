package repository

import (
	"blog/entity"
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const (
	Danger  = "danger"
	Success = "success"
)

func SetStatusResponse(Label, Message string) map[string]string {
	return map[string]string{
		"Label":   Label,
		"Message": Message,
	}
}
func ValidasiResponse(c *fiber.Ctx, Trans *ut.Translator, Validation error) (error, *entity.ResponseDataRoute) {
	response := entity.ResponseDataRoute{}
	err := errors.New(FailedValidasiForm)

	if Validation != nil {
		validationErrors := Validation.(validator.ValidationErrors)

		Set := entity.ValidationResponse{ValidationListError: map[string]entity.ValidationList{}}
		for _, e := range validationErrors {
			Set.ValidationListError[e.Field()] = entity.ValidationList{
				Message: e.Translate(*Trans),
			}
		}

		response.Validation = entity.ValidationResponse{
			Status: entity.StatusResponse{
				Label:   Danger,
				Message: FailedValidasiForm,
			},
			ValidationListError: Set.ValidationListError,
		}

	} else {
		err = nil
	}

	return err, &response

}
