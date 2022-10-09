package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Sntree2mi8/gqlgen-validator-sample/graph/customerr"
	"github.com/Sntree2mi8/gqlgen-validator-sample/graph/generated"
	"github.com/Sntree2mi8/gqlgen-validator-sample/graph/model"
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

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	if err := validateInputModel(input); err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:   "NewTodoID",
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "UserName",
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
