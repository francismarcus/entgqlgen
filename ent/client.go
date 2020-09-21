// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/francismarcus/eg/ent/migrate"

	"github.com/francismarcus/eg/ent/exercise"
	"github.com/francismarcus/eg/ent/program"
	"github.com/francismarcus/eg/ent/shout"
	"github.com/francismarcus/eg/ent/user"
	"github.com/francismarcus/eg/ent/usersettings"
	"github.com/francismarcus/eg/ent/workout"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Exercise is the client for interacting with the Exercise builders.
	Exercise *ExerciseClient
	// Program is the client for interacting with the Program builders.
	Program *ProgramClient
	// Shout is the client for interacting with the Shout builders.
	Shout *ShoutClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserSettings is the client for interacting with the UserSettings builders.
	UserSettings *UserSettingsClient
	// Workout is the client for interacting with the Workout builders.
	Workout *WorkoutClient

	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Exercise = NewExerciseClient(c.config)
	c.Program = NewProgramClient(c.config)
	c.Shout = NewShoutClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserSettings = NewUserSettingsClient(c.config)
	c.Workout = NewWorkoutClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Exercise:     NewExerciseClient(cfg),
		Program:      NewProgramClient(cfg),
		Shout:        NewShoutClient(cfg),
		User:         NewUserClient(cfg),
		UserSettings: NewUserSettingsClient(cfg),
		Workout:      NewWorkoutClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:       cfg,
		Exercise:     NewExerciseClient(cfg),
		Program:      NewProgramClient(cfg),
		Shout:        NewShoutClient(cfg),
		User:         NewUserClient(cfg),
		UserSettings: NewUserSettingsClient(cfg),
		Workout:      NewWorkoutClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Exercise.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Exercise.Use(hooks...)
	c.Program.Use(hooks...)
	c.Shout.Use(hooks...)
	c.User.Use(hooks...)
	c.UserSettings.Use(hooks...)
	c.Workout.Use(hooks...)
}

// ExerciseClient is a client for the Exercise schema.
type ExerciseClient struct {
	config
}

// NewExerciseClient returns a client for the Exercise from the given config.
func NewExerciseClient(c config) *ExerciseClient {
	return &ExerciseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `exercise.Hooks(f(g(h())))`.
func (c *ExerciseClient) Use(hooks ...Hook) {
	c.hooks.Exercise = append(c.hooks.Exercise, hooks...)
}

// Create returns a create builder for Exercise.
func (c *ExerciseClient) Create() *ExerciseCreate {
	mutation := newExerciseMutation(c.config, OpCreate)
	return &ExerciseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Exercise entities.
func (c *ExerciseClient) CreateBulk(builders ...*ExerciseCreate) *ExerciseCreateBulk {
	return &ExerciseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Exercise.
func (c *ExerciseClient) Update() *ExerciseUpdate {
	mutation := newExerciseMutation(c.config, OpUpdate)
	return &ExerciseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ExerciseClient) UpdateOne(e *Exercise) *ExerciseUpdateOne {
	mutation := newExerciseMutation(c.config, OpUpdateOne, withExercise(e))
	return &ExerciseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ExerciseClient) UpdateOneID(id int) *ExerciseUpdateOne {
	mutation := newExerciseMutation(c.config, OpUpdateOne, withExerciseID(id))
	return &ExerciseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Exercise.
func (c *ExerciseClient) Delete() *ExerciseDelete {
	mutation := newExerciseMutation(c.config, OpDelete)
	return &ExerciseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ExerciseClient) DeleteOne(e *Exercise) *ExerciseDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ExerciseClient) DeleteOneID(id int) *ExerciseDeleteOne {
	builder := c.Delete().Where(exercise.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ExerciseDeleteOne{builder}
}

// Query returns a query builder for Exercise.
func (c *ExerciseClient) Query() *ExerciseQuery {
	return &ExerciseQuery{config: c.config}
}

// Get returns a Exercise entity by its id.
func (c *ExerciseClient) Get(ctx context.Context, id int) (*Exercise, error) {
	return c.Query().Where(exercise.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ExerciseClient) GetX(ctx context.Context, id int) *Exercise {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ExerciseClient) Hooks() []Hook {
	return c.hooks.Exercise
}

// ProgramClient is a client for the Program schema.
type ProgramClient struct {
	config
}

// NewProgramClient returns a client for the Program from the given config.
func NewProgramClient(c config) *ProgramClient {
	return &ProgramClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `program.Hooks(f(g(h())))`.
func (c *ProgramClient) Use(hooks ...Hook) {
	c.hooks.Program = append(c.hooks.Program, hooks...)
}

// Create returns a create builder for Program.
func (c *ProgramClient) Create() *ProgramCreate {
	mutation := newProgramMutation(c.config, OpCreate)
	return &ProgramCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Program entities.
func (c *ProgramClient) CreateBulk(builders ...*ProgramCreate) *ProgramCreateBulk {
	return &ProgramCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Program.
func (c *ProgramClient) Update() *ProgramUpdate {
	mutation := newProgramMutation(c.config, OpUpdate)
	return &ProgramUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProgramClient) UpdateOne(pr *Program) *ProgramUpdateOne {
	mutation := newProgramMutation(c.config, OpUpdateOne, withProgram(pr))
	return &ProgramUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProgramClient) UpdateOneID(id int) *ProgramUpdateOne {
	mutation := newProgramMutation(c.config, OpUpdateOne, withProgramID(id))
	return &ProgramUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Program.
func (c *ProgramClient) Delete() *ProgramDelete {
	mutation := newProgramMutation(c.config, OpDelete)
	return &ProgramDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProgramClient) DeleteOne(pr *Program) *ProgramDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProgramClient) DeleteOneID(id int) *ProgramDeleteOne {
	builder := c.Delete().Where(program.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProgramDeleteOne{builder}
}

// Query returns a query builder for Program.
func (c *ProgramClient) Query() *ProgramQuery {
	return &ProgramQuery{config: c.config}
}

// Get returns a Program entity by its id.
func (c *ProgramClient) Get(ctx context.Context, id int) (*Program, error) {
	return c.Query().Where(program.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProgramClient) GetX(ctx context.Context, id int) *Program {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a Program.
func (c *ProgramClient) QueryAuthor(pr *Program) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(program.Table, program.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, program.AuthorTable, program.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWorkouts queries the workouts edge of a Program.
func (c *ProgramClient) QueryWorkouts(pr *Program) *WorkoutQuery {
	query := &WorkoutQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(program.Table, program.FieldID, id),
			sqlgraph.To(workout.Table, workout.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, program.WorkoutsTable, program.WorkoutsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProgramClient) Hooks() []Hook {
	return c.hooks.Program
}

// ShoutClient is a client for the Shout schema.
type ShoutClient struct {
	config
}

// NewShoutClient returns a client for the Shout from the given config.
func NewShoutClient(c config) *ShoutClient {
	return &ShoutClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `shout.Hooks(f(g(h())))`.
func (c *ShoutClient) Use(hooks ...Hook) {
	c.hooks.Shout = append(c.hooks.Shout, hooks...)
}

// Create returns a create builder for Shout.
func (c *ShoutClient) Create() *ShoutCreate {
	mutation := newShoutMutation(c.config, OpCreate)
	return &ShoutCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Shout entities.
func (c *ShoutClient) CreateBulk(builders ...*ShoutCreate) *ShoutCreateBulk {
	return &ShoutCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Shout.
func (c *ShoutClient) Update() *ShoutUpdate {
	mutation := newShoutMutation(c.config, OpUpdate)
	return &ShoutUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ShoutClient) UpdateOne(s *Shout) *ShoutUpdateOne {
	mutation := newShoutMutation(c.config, OpUpdateOne, withShout(s))
	return &ShoutUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ShoutClient) UpdateOneID(id int) *ShoutUpdateOne {
	mutation := newShoutMutation(c.config, OpUpdateOne, withShoutID(id))
	return &ShoutUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Shout.
func (c *ShoutClient) Delete() *ShoutDelete {
	mutation := newShoutMutation(c.config, OpDelete)
	return &ShoutDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ShoutClient) DeleteOne(s *Shout) *ShoutDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ShoutClient) DeleteOneID(id int) *ShoutDeleteOne {
	builder := c.Delete().Where(shout.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ShoutDeleteOne{builder}
}

// Query returns a query builder for Shout.
func (c *ShoutClient) Query() *ShoutQuery {
	return &ShoutQuery{config: c.config}
}

// Get returns a Shout entity by its id.
func (c *ShoutClient) Get(ctx context.Context, id int) (*Shout, error) {
	return c.Query().Where(shout.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ShoutClient) GetX(ctx context.Context, id int) *Shout {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a Shout.
func (c *ShoutClient) QueryAuthor(s *Shout) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(shout.Table, shout.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, shout.AuthorTable, shout.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikedBy queries the liked_by edge of a Shout.
func (c *ShoutClient) QueryLikedBy(s *Shout) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(shout.Table, shout.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, shout.LikedByTable, shout.LikedByPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ShoutClient) Hooks() []Hook {
	return c.hooks.Shout
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFollowers queries the followers edge of a User.
func (c *UserClient) QueryFollowers(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.FollowersTable, user.FollowersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFollowing queries the following edge of a User.
func (c *UserClient) QueryFollowing(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.FollowingTable, user.FollowingPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPrograms queries the programs edge of a User.
func (c *UserClient) QueryPrograms(u *User) *ProgramQuery {
	query := &ProgramQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(program.Table, program.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ProgramsTable, user.ProgramsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryShouts queries the shouts edge of a User.
func (c *UserClient) QueryShouts(u *User) *ShoutQuery {
	query := &ShoutQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(shout.Table, shout.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ShoutsTable, user.ShoutsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikedShouts queries the liked_shouts edge of a User.
func (c *UserClient) QueryLikedShouts(u *User) *ShoutQuery {
	query := &ShoutQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(shout.Table, shout.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.LikedShoutsTable, user.LikedShoutsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySettings queries the settings edge of a User.
func (c *UserClient) QuerySettings(u *User) *UserSettingsQuery {
	query := &UserSettingsQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(usersettings.Table, usersettings.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.SettingsTable, user.SettingsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// UserSettingsClient is a client for the UserSettings schema.
type UserSettingsClient struct {
	config
}

// NewUserSettingsClient returns a client for the UserSettings from the given config.
func NewUserSettingsClient(c config) *UserSettingsClient {
	return &UserSettingsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usersettings.Hooks(f(g(h())))`.
func (c *UserSettingsClient) Use(hooks ...Hook) {
	c.hooks.UserSettings = append(c.hooks.UserSettings, hooks...)
}

// Create returns a create builder for UserSettings.
func (c *UserSettingsClient) Create() *UserSettingsCreate {
	mutation := newUserSettingsMutation(c.config, OpCreate)
	return &UserSettingsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of UserSettings entities.
func (c *UserSettingsClient) CreateBulk(builders ...*UserSettingsCreate) *UserSettingsCreateBulk {
	return &UserSettingsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserSettings.
func (c *UserSettingsClient) Update() *UserSettingsUpdate {
	mutation := newUserSettingsMutation(c.config, OpUpdate)
	return &UserSettingsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserSettingsClient) UpdateOne(us *UserSettings) *UserSettingsUpdateOne {
	mutation := newUserSettingsMutation(c.config, OpUpdateOne, withUserSettings(us))
	return &UserSettingsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserSettingsClient) UpdateOneID(id int) *UserSettingsUpdateOne {
	mutation := newUserSettingsMutation(c.config, OpUpdateOne, withUserSettingsID(id))
	return &UserSettingsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserSettings.
func (c *UserSettingsClient) Delete() *UserSettingsDelete {
	mutation := newUserSettingsMutation(c.config, OpDelete)
	return &UserSettingsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserSettingsClient) DeleteOne(us *UserSettings) *UserSettingsDeleteOne {
	return c.DeleteOneID(us.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserSettingsClient) DeleteOneID(id int) *UserSettingsDeleteOne {
	builder := c.Delete().Where(usersettings.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserSettingsDeleteOne{builder}
}

// Query returns a query builder for UserSettings.
func (c *UserSettingsClient) Query() *UserSettingsQuery {
	return &UserSettingsQuery{config: c.config}
}

// Get returns a UserSettings entity by its id.
func (c *UserSettingsClient) Get(ctx context.Context, id int) (*UserSettings, error) {
	return c.Query().Where(usersettings.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserSettingsClient) GetX(ctx context.Context, id int) *UserSettings {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBelongsTo queries the belongs_to edge of a UserSettings.
func (c *UserSettingsClient) QueryBelongsTo(us *UserSettings) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := us.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usersettings.Table, usersettings.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, usersettings.BelongsToTable, usersettings.BelongsToColumn),
		)
		fromV = sqlgraph.Neighbors(us.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserSettingsClient) Hooks() []Hook {
	return c.hooks.UserSettings
}

// WorkoutClient is a client for the Workout schema.
type WorkoutClient struct {
	config
}

// NewWorkoutClient returns a client for the Workout from the given config.
func NewWorkoutClient(c config) *WorkoutClient {
	return &WorkoutClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `workout.Hooks(f(g(h())))`.
func (c *WorkoutClient) Use(hooks ...Hook) {
	c.hooks.Workout = append(c.hooks.Workout, hooks...)
}

// Create returns a create builder for Workout.
func (c *WorkoutClient) Create() *WorkoutCreate {
	mutation := newWorkoutMutation(c.config, OpCreate)
	return &WorkoutCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Workout entities.
func (c *WorkoutClient) CreateBulk(builders ...*WorkoutCreate) *WorkoutCreateBulk {
	return &WorkoutCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Workout.
func (c *WorkoutClient) Update() *WorkoutUpdate {
	mutation := newWorkoutMutation(c.config, OpUpdate)
	return &WorkoutUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkoutClient) UpdateOne(w *Workout) *WorkoutUpdateOne {
	mutation := newWorkoutMutation(c.config, OpUpdateOne, withWorkout(w))
	return &WorkoutUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkoutClient) UpdateOneID(id int) *WorkoutUpdateOne {
	mutation := newWorkoutMutation(c.config, OpUpdateOne, withWorkoutID(id))
	return &WorkoutUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Workout.
func (c *WorkoutClient) Delete() *WorkoutDelete {
	mutation := newWorkoutMutation(c.config, OpDelete)
	return &WorkoutDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WorkoutClient) DeleteOne(w *Workout) *WorkoutDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WorkoutClient) DeleteOneID(id int) *WorkoutDeleteOne {
	builder := c.Delete().Where(workout.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkoutDeleteOne{builder}
}

// Query returns a query builder for Workout.
func (c *WorkoutClient) Query() *WorkoutQuery {
	return &WorkoutQuery{config: c.config}
}

// Get returns a Workout entity by its id.
func (c *WorkoutClient) Get(ctx context.Context, id int) (*Workout, error) {
	return c.Query().Where(workout.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkoutClient) GetX(ctx context.Context, id int) *Workout {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProgram queries the program edge of a Workout.
func (c *WorkoutClient) QueryProgram(w *Workout) *ProgramQuery {
	query := &ProgramQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workout.Table, workout.FieldID, id),
			sqlgraph.To(program.Table, program.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workout.ProgramTable, workout.ProgramColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WorkoutClient) Hooks() []Hook {
	return c.hooks.Workout
}