// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/francismarcus/eg/ent/diet"
	"github.com/francismarcus/eg/ent/predicate"
	"github.com/francismarcus/eg/ent/user"
)

// DietUpdate is the builder for updating Diet entities.
type DietUpdate struct {
	config
	hooks      []Hook
	mutation   *DietMutation
	predicates []predicate.Diet
}

// Where adds a new predicate for the builder.
func (du *DietUpdate) Where(ps ...predicate.Diet) *DietUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetUpdatedAt sets the updated_at field.
func (du *DietUpdate) SetUpdatedAt(t time.Time) *DietUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetName sets the name field.
func (du *DietUpdate) SetName(s string) *DietUpdate {
	du.mutation.SetName(s)
	return du
}

// SetGoalWeight sets the goal_weight field.
func (du *DietUpdate) SetGoalWeight(i int) *DietUpdate {
	du.mutation.ResetGoalWeight()
	du.mutation.SetGoalWeight(i)
	return du
}

// AddGoalWeight adds i to goal_weight.
func (du *DietUpdate) AddGoalWeight(i int) *DietUpdate {
	du.mutation.AddGoalWeight(i)
	return du
}

// SetLength sets the length field.
func (du *DietUpdate) SetLength(i int) *DietUpdate {
	du.mutation.ResetLength()
	du.mutation.SetLength(i)
	return du
}

// AddLength adds i to length.
func (du *DietUpdate) AddLength(i int) *DietUpdate {
	du.mutation.AddLength(i)
	return du
}

// SetAuthorID sets the author edge to User by id.
func (du *DietUpdate) SetAuthorID(id int) *DietUpdate {
	du.mutation.SetAuthorID(id)
	return du
}

// SetAuthor sets the author edge to User.
func (du *DietUpdate) SetAuthor(u *User) *DietUpdate {
	return du.SetAuthorID(u.ID)
}

// Mutation returns the DietMutation object of the builder.
func (du *DietUpdate) Mutation() *DietMutation {
	return du.mutation
}

// ClearAuthor clears the "author" edge to type User.
func (du *DietUpdate) ClearAuthor() *DietUpdate {
	du.mutation.ClearAuthor()
	return du
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DietUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	du.defaults()
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DietMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DietUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DietUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DietUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DietUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := diet.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DietUpdate) check() error {
	if _, ok := du.mutation.AuthorID(); du.mutation.AuthorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"author\"")
	}
	return nil
}

func (du *DietUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   diet.Table,
			Columns: diet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: diet.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: diet.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diet.FieldName,
		})
	}
	if value, ok := du.mutation.GoalWeight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diet.FieldGoalWeight,
		})
	}
	if value, ok := du.mutation.AddedGoalWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diet.FieldGoalWeight,
		})
	}
	if value, ok := du.mutation.Length(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diet.FieldLength,
		})
	}
	if value, ok := du.mutation.AddedLength(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diet.FieldLength,
		})
	}
	if du.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   diet.AuthorTable,
			Columns: []string{diet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   diet.AuthorTable,
			Columns: []string{diet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{diet.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DietUpdateOne is the builder for updating a single Diet entity.
type DietUpdateOne struct {
	config
	hooks    []Hook
	mutation *DietMutation
}

// SetUpdatedAt sets the updated_at field.
func (duo *DietUpdateOne) SetUpdatedAt(t time.Time) *DietUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetName sets the name field.
func (duo *DietUpdateOne) SetName(s string) *DietUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetGoalWeight sets the goal_weight field.
func (duo *DietUpdateOne) SetGoalWeight(i int) *DietUpdateOne {
	duo.mutation.ResetGoalWeight()
	duo.mutation.SetGoalWeight(i)
	return duo
}

// AddGoalWeight adds i to goal_weight.
func (duo *DietUpdateOne) AddGoalWeight(i int) *DietUpdateOne {
	duo.mutation.AddGoalWeight(i)
	return duo
}

// SetLength sets the length field.
func (duo *DietUpdateOne) SetLength(i int) *DietUpdateOne {
	duo.mutation.ResetLength()
	duo.mutation.SetLength(i)
	return duo
}

// AddLength adds i to length.
func (duo *DietUpdateOne) AddLength(i int) *DietUpdateOne {
	duo.mutation.AddLength(i)
	return duo
}

// SetAuthorID sets the author edge to User by id.
func (duo *DietUpdateOne) SetAuthorID(id int) *DietUpdateOne {
	duo.mutation.SetAuthorID(id)
	return duo
}

// SetAuthor sets the author edge to User.
func (duo *DietUpdateOne) SetAuthor(u *User) *DietUpdateOne {
	return duo.SetAuthorID(u.ID)
}

// Mutation returns the DietMutation object of the builder.
func (duo *DietUpdateOne) Mutation() *DietMutation {
	return duo.mutation
}

// ClearAuthor clears the "author" edge to type User.
func (duo *DietUpdateOne) ClearAuthor() *DietUpdateOne {
	duo.mutation.ClearAuthor()
	return duo
}

// Save executes the query and returns the updated entity.
func (duo *DietUpdateOne) Save(ctx context.Context) (*Diet, error) {
	var (
		err  error
		node *Diet
	)
	duo.defaults()
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DietMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DietUpdateOne) SaveX(ctx context.Context) *Diet {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DietUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DietUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DietUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := diet.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DietUpdateOne) check() error {
	if _, ok := duo.mutation.AuthorID(); duo.mutation.AuthorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"author\"")
	}
	return nil
}

func (duo *DietUpdateOne) sqlSave(ctx context.Context) (_node *Diet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   diet.Table,
			Columns: diet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: diet.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Diet.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: diet.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diet.FieldName,
		})
	}
	if value, ok := duo.mutation.GoalWeight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diet.FieldGoalWeight,
		})
	}
	if value, ok := duo.mutation.AddedGoalWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diet.FieldGoalWeight,
		})
	}
	if value, ok := duo.mutation.Length(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diet.FieldLength,
		})
	}
	if value, ok := duo.mutation.AddedLength(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: diet.FieldLength,
		})
	}
	if duo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   diet.AuthorTable,
			Columns: []string{diet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   diet.AuthorTable,
			Columns: []string{diet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Diet{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{diet.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}