package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "len":
		return fmt.Sprintf("The length of this field must be %s", fe.Param())
	}
	return fe.Error()
}

func ValidateModel(model any) (map[string]string, error) {
	if err := validate.Struct(model); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err
		}

		errs := err.(validator.ValidationErrors)
		validationErrors := make(map[string]string, len(errs))
		for _, ve := range errs {
			validationErrors[ve.Field()] = msgForTag(ve)
		}

		return validationErrors, nil
	}
	return nil, nil
}
