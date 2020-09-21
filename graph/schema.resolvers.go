package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/francismarcus/eg/ent"
	"github.com/francismarcus/eg/ent/program"
	"github.com/francismarcus/eg/ent/shout"
	"github.com/francismarcus/eg/ent/user"
	"github.com/francismarcus/eg/graph/generated"
	"github.com/francismarcus/eg/graph/models"
	"github.com/francismarcus/eg/pkg/auth"
	"github.com/francismarcus/eg/pkg/middlewares"
)

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

func (r *mutationResolver) CreateProgram(ctx context.Context, input models.CreateProgramInput) (*ent.Program, error) {
	return r.client.Program.Create().SetName(input.Name).
		SetAuthorID(input.UserID).
		SaveX(ctx), nil
}

func (r *mutationResolver) CreateWorkout(ctx context.Context, input models.CreateWorkoutInput) (*ent.Workout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateShout(ctx context.Context, input models.CreateShoutInput) (*ent.Shout, error) {
	return r.client.Shout.Create().SetMessage(input.Message).
		SetAuthorID(input.UserID).
		SaveX(ctx), nil
}

func (r *mutationResolver) LikeShout(ctx context.Context, input models.LikeShoutInput) (*ent.Shout, error) {
	s := r.client.Shout.Query().Where(shout.ID(input.ShoutID)).OnlyX(ctx)
	a := s.QueryAuthor().OnlyX(ctx)
	users := s.QueryLikedBy().AllX(ctx)

	for _, u := range users {
		if u.ID == input.UserID {
			return nil, fmt.Errorf("You already like this")
		}
	}

	if a.ID == input.UserID {
		return nil, fmt.Errorf("Cant like your own shouts")
	}

	return s.Update().AddLikedByIDs(input.UserID).AddLikes(1).Save(ctx)
}

func (r *mutationResolver) UnlikeShout(ctx context.Context, input models.UnlikeShoutInput) (*ent.Shout, error) {
	s := r.client.Shout.Query().Where(shout.ID(input.ShoutID)).OnlyX(ctx)
	a := s.QueryLikedBy().AllX(ctx)

	var b bool

	for _, u := range a {
		if u.ID == input.UserID {
			b = true
		}
	}

	if b == true {
		return s.Update().RemoveLikedByIDs(input.UserID).AddLikes(-1).Save(ctx)
	}

	return nil, fmt.Errorf("Cant unlike a shout you dont like")
}

func (r *mutationResolver) FollowUser(ctx context.Context, input models.FollowUserInput) (*ent.User, error) {
	if input.UserID == input.FollowID {
		return nil, fmt.Errorf("Cant follow yourself")
	}

	f := r.client.User.UpdateOneID(input.FollowID).AddFollowerIDs(input.UserID).AddFollowersCount(1).SaveX(ctx)
	r.client.User.UpdateOneID(input.UserID).AddFollowsCount(1).SaveX(ctx)

	return f, nil
}

func (r *mutationResolver) UnFollowUser(ctx context.Context, input models.UnFollowUserInput) (*ent.User, error) {
	if input.UserID == input.FollowID {
		return nil, fmt.Errorf("Cant unfollow yourself")
	}

	f := r.client.User.UpdateOneID(input.FollowID).RemoveFollowerIDs(input.UserID).AddFollowersCount(-1).SaveX(ctx)
	r.client.User.UpdateOneID(input.UserID).AddFollowsCount(-1).SaveX(ctx)

	return f, nil
}

func (r *mutationResolver) AddUserSettings(ctx context.Context, input models.AddUserSettingsInput) (*ent.UserSettings, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *programResolver) Author(ctx context.Context, obj *ent.Program) (*ent.User, error) {
	return obj.QueryAuthor().Only(ctx)
}

func (r *programResolver) Workouts(ctx context.Context, obj *ent.Program, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.WorkoutOrder) (*ent.WorkoutConnection, error) {
	panic(fmt.Errorf("not implemented"))
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

func (r *queryResolver) Programs(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProgramOrder) (*ent.ProgramConnection, error) {
	return r.client.Program.Query().
		Paginate(ctx, after, first, before, last, ent.WithProgramOrder(orderBy))
}

func (r *queryResolver) Feed(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ShoutOrder, id int) (*ent.ShoutConnection, error) {
	return r.client.Shout.Query().Where(
		shout.HasAuthorWith(
			user.HasFollowersWith(
				user.ID(id),
			),
		),
	).Paginate(ctx, after, first, before, last, ent.WithShoutOrder(orderBy))
}

func (r *queryResolver) Shouts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ShoutOrder) (*ent.ShoutConnection, error) {
	return r.client.Shout.Query().Paginate(ctx, after, first, before, last, ent.WithShoutOrder(orderBy))
}

func (r *queryResolver) ProgramByID(ctx context.Context, id int) (*ent.Program, error) {
	return r.client.Program.Query().Where(
		program.ID(id),
	).Only(ctx)
}

func (r *queryResolver) MyPrograms(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProgramOrder, id int) (*ent.ProgramConnection, error) {
	return r.client.Program.Query().Where(
		program.HasAuthorWith(user.ID(id)),
	).Paginate(ctx, after, first, before, last, ent.WithProgramOrder(orderBy))
}

func (r *shoutResolver) Author(ctx context.Context, obj *ent.Shout) (*ent.User, error) {
	return obj.QueryAuthor().Only(ctx)
}

func (r *shoutResolver) LikedBy(ctx context.Context, obj *ent.Shout) ([]*ent.User, error) {
	return obj.QueryLikedBy().All(ctx)
}

func (r *userResolver) Settings(ctx context.Context, obj *ent.User) (*ent.UserSettings, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Programs(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProgramOrder) (*ent.ProgramConnection, error) {
	return obj.QueryPrograms().Paginate(ctx, after, first, before, last, ent.WithProgramOrder(orderBy))
}

func (r *userResolver) Followers(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return obj.QueryFollowers().Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy))
}

func (r *userResolver) Follows(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return obj.QueryFollowing().Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy))
}

func (r *userResolver) Shouts(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ShoutOrder) (*ent.ShoutConnection, error) {
	return obj.QueryShouts().Paginate(ctx, after, first, before, last, ent.WithShoutOrder(orderBy))
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

func (r *workoutResolver) Program(ctx context.Context, obj *ent.Workout) (*ent.Program, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Program returns generated.ProgramResolver implementation.
func (r *Resolver) Program() generated.ProgramResolver { return &programResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Shout returns generated.ShoutResolver implementation.
func (r *Resolver) Shout() generated.ShoutResolver { return &shoutResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserSettings returns generated.UserSettingsResolver implementation.
func (r *Resolver) UserSettings() generated.UserSettingsResolver { return &userSettingsResolver{r} }

// Workout returns generated.WorkoutResolver implementation.
func (r *Resolver) Workout() generated.WorkoutResolver { return &workoutResolver{r} }

type mutationResolver struct{ *Resolver }
type programResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type shoutResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userSettingsResolver struct{ *Resolver }
type workoutResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *shoutResolver) User(ctx context.Context, obj *ent.Shout) (*ent.User, error) {
	return obj.QueryAuthor().Only(ctx)
}
func (r *programResolver) BatchedAuthor(ctx context.Context, obj *ent.Program) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}
