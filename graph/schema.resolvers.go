package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/francismarcus/eg/ent"
	"github.com/francismarcus/eg/ent/user"
	"github.com/francismarcus/eg/graph/generated"
	"github.com/francismarcus/eg/graph/models"
	"github.com/francismarcus/eg/pkg/auth"
	"github.com/francismarcus/eg/pkg/middlewares"
)

func (r *dietResolver) User(ctx context.Context, obj *ent.Diet) (*ent.User, error) {
	return obj.QueryAuthor().Only(ctx)
}

func (r *mutationResolver) UsernameAvailable(ctx context.Context, username string) (*bool, error) {
	var b bool
	u, err := r.client.User.Query().Where(user.Username(username)).Only(ctx)

	if err != nil {
		b = true
	}

	if u != nil {
		b = false
		return &b, fmt.Errorf("Username taken: %v", u)
	}

	return &b, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthPayload, error) {
	u, err := r.client.User.Query().Where(user.Username(input.Username)).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed loginUser: %v", err)
	}

	if (auth.CompareHash(input.Password, u.Password)) != true {
		return nil, fmt.Errorf("wrong password: %v", err)
	}

	token, _ := auth.GenerateToken(u.ID)

	return &models.AuthPayload{
		User:  u,
		Token: &token,
	}, nil
}

func (r *mutationResolver) Signup(ctx context.Context, input models.SignupInput) (*models.AuthPayload, error) {
	client := ent.FromContext(ctx)
	password, _ := auth.HashPassword(input.Password)
	username := strings.ToLower(input.Username)

	u, err := client.User.Create().SetUsername(username).SetEmail(input.Email).SetPassword(password).Save(ctx)
	token, _ := auth.GenerateToken(u.ID)

	if err != nil {
		return nil, fmt.Errorf("failed signupUser: %v", err)
	}

	return &models.AuthPayload{
		User:  u,
		Token: &token,
	}, nil
}

func (r *mutationResolver) AddUserSettings(ctx context.Context, input models.AddUserSettingsInput) (*ent.UserSettings, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateDiet(ctx context.Context, input models.CreateDietInput) (*ent.Diet, error) {
	return r.client.Diet.Create().SetAuthorID(input.UserID).SetName(input.Name).SetLength(4).SetGoalWeight(90).Save(ctx)
}

func (r *queryResolver) Whoami(ctx context.Context) (*ent.User, error) {
	u := middlewares.UserContext(ctx)
	user := r.client.User.Query().Where(user.ID(u.ID)).OnlyX(ctx)
	return user, nil
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Me(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Query().Where(user.ID(id)).Only(ctx)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy))
}

func (r *userResolver) Settings(ctx context.Context, obj *ent.User) (*ent.UserSettings, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Diets(ctx context.Context, obj *ent.User) ([]*ent.Diet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userSettingsResolver) Gender(ctx context.Context, obj *ent.UserSettings) (*models.Gender, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userSettingsResolver) Level(ctx context.Context, obj *ent.UserSettings) (*models.Level, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userSettingsResolver) User(ctx context.Context, obj *ent.UserSettings) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Diet returns generated.DietResolver implementation.
func (r *Resolver) Diet() generated.DietResolver { return &dietResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserSettings returns generated.UserSettingsResolver implementation.
func (r *Resolver) UserSettings() generated.UserSettingsResolver { return &userSettingsResolver{r} }

type dietResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userSettingsResolver struct{ *Resolver }
