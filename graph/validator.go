package graph

import (
	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func validateModel(model any) *gqlerror.Error {
	if err := validate.Struct(model); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Printf("failed to validate model %+v", err)
			return &gqlerror.Error{
				Message: "internal server error",
			}
		}

		validationErrors := make(map[string]any, len(err.(validator.ValidationErrors)))
		for _, ve := range err.(validator.ValidationErrors) {
			validationErrors[ve.Field()] = ve.Error()
		}

		return &gqlerror.Error{
			Message:    "validation error",
			Extensions: validationErrors,
		}
	}
	return nil
}
