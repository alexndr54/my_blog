package repository

import (
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
)

func SetLanguageID() (*validator.Validate, *ut.Translator) {
	Indonesia := id.New()
	uni := ut.New(Indonesia, Indonesia)
	trans, _ := uni.GetTranslator("id")
	validate := validator.New()

	err := idTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic("register translation error")
	}

	return validate, &trans
}
