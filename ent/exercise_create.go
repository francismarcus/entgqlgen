// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/francismarcus/eg/ent/exercise"
)

// ExerciseCreate is the builder for creating a Exercise entity.
type ExerciseCreate struct {
	config
	mutation *ExerciseMutation
	hooks    []Hook
}

// Mutation returns the ExerciseMutation object of the builder.
func (ec *ExerciseCreate) Mutation() *ExerciseMutation {
	return ec.mutation
}

// Save creates the Exercise in the database.
func (ec *ExerciseCreate) Save(ctx context.Context) (*Exercise, error) {
	var (
		err  error
		node *Exercise
	)
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExerciseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *ExerciseCreate) SaveX(ctx context.Context) *Exercise {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (ec *ExerciseCreate) check() error {
	return nil
}

func (ec *ExerciseCreate) sqlSave(ctx context.Context) (*Exercise, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ec *ExerciseCreate) createSpec() (*Exercise, *sqlgraph.CreateSpec) {
	var (
		_node = &Exercise{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: exercise.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: exercise.FieldID,
			},
		}
	)
	return _node, _spec
}

// ExerciseCreateBulk is the builder for creating a bulk of Exercise entities.
type ExerciseCreateBulk struct {
	config
	builders []*ExerciseCreate
}

// Save creates the Exercise entities in the database.
func (ecb *ExerciseCreateBulk) Save(ctx context.Context) ([]*Exercise, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Exercise, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExerciseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ecb *ExerciseCreateBulk) SaveX(ctx context.Context) []*Exercise {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
