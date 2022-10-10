package graph

import (
	"github.com/Sntree2mi8/gqlgen-validator-sample/graph/customerr"
	"github.com/Sntree2mi8/gqlgen-validator-sample/graph/validation"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func validateInputModel(m any) *gqlerror.Error {
	validationErrors, err := validation.ValidateModel(m)
	if err != nil {
		return &gqlerror.Error{
			Message:    customerr.ErrorMessage(customerr.InternalServerError),
			Extensions: customerr.InternalServerErrorExtension(),
		}
	}
	if len(validationErrors) > 0 {
		return &gqlerror.Error{
			Message:    customerr.ErrorMessage(customerr.BadUserInput),
			Extensions: customerr.BadUserInputExtension(validationErrors),
		}
	}
	return nil
}
