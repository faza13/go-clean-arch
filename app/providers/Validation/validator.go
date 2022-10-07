package Validation

import (
	"base/app/common"
	"base/app/providers/translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"strings"
)

type IValidator interface {
	Make(StructToBevalidate interface{}) map[string]string
}

type Validator struct {
	translator translator.ITranslator
	validate   *validator.Validate
	err        error
}

func NewValidatroProvider(translator translator.ITranslator) IValidator {

	validate := validator.New()
	id_translations.RegisterDefaultTranslations(validate, translator.GetTrans())

	return &Validator{
		validate:   validate,
		translator: translator,
	}
}

func (v Validator) Make(StructToBevalidate interface{}) map[string]string {
	err := v.validate.Struct(StructToBevalidate)
	if err != nil {
		return v.getErrors(err)
	}

	return nil
}

func (v Validator) getErrors(err error) map[string]string {
	validatorErrs := err.(validator.ValidationErrors)
	errorMap := make(map[string]string)
	for _, e := range validatorErrs {
		tranlate := v.translator.T(e.Field())
		if tranlate == "" {
			tranlate = e.Field()
		}
		errorMap[common.ToSnakeCase(e.Field())] =
			strings.ReplaceAll(e.Translate(v.translator.GetTrans()), e.Field(), tranlate)

	}

	return errorMap
}
