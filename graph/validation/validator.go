package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
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

func ValidateModel(model any) *gqlerror.Error {
	if err := validate.Struct(model); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Printf("failed to validate model %+v", err)
			return &gqlerror.Error{
				Message: "internal server error",
			}
		}

		errs := err.(validator.ValidationErrors)
		validationErrors := make(map[string]any, len(errs))
		for _, ve := range errs {
			validationErrors[ve.Field()] = msgForTag(ve)
		}

		return &gqlerror.Error{
			Message:    "validation error",
			Extensions: validationErrors,
		}
	}
	return nil
}
