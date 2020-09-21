// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/francismarcus/eg/ent/predicate"
	"github.com/francismarcus/eg/ent/workout"
)

// WorkoutDelete is the builder for deleting a Workout entity.
type WorkoutDelete struct {
	config
	hooks      []Hook
	mutation   *WorkoutMutation
	predicates []predicate.Workout
}

// Where adds a new predicate to the delete builder.
func (wd *WorkoutDelete) Where(ps ...predicate.Workout) *WorkoutDelete {
	wd.predicates = append(wd.predicates, ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WorkoutDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wd.hooks) == 0 {
		affected, err = wd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkoutMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wd.mutation = mutation
			affected, err = wd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wd.hooks) - 1; i >= 0; i-- {
			mut = wd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WorkoutDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WorkoutDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: workout.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workout.FieldID,
			},
		},
	}
	if ps := wd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
}

// WorkoutDeleteOne is the builder for deleting a single Workout entity.
type WorkoutDeleteOne struct {
	wd *WorkoutDelete
}

// Exec executes the deletion query.
func (wdo *WorkoutDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workout.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WorkoutDeleteOne) ExecX(ctx context.Context) {
	wdo.wd.ExecX(ctx)
}
