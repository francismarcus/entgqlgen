// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/francismarcus/eg/ent/predicate"
	"github.com/francismarcus/eg/ent/shout"
	"github.com/francismarcus/eg/ent/user"
)

// ShoutUpdate is the builder for updating Shout entities.
type ShoutUpdate struct {
	config
	hooks      []Hook
	mutation   *ShoutMutation
	predicates []predicate.Shout
}

// Where adds a new predicate for the builder.
func (su *ShoutUpdate) Where(ps ...predicate.Shout) *ShoutUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetUpdatedAt sets the updated_at field.
func (su *ShoutUpdate) SetUpdatedAt(t time.Time) *ShoutUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetMessage sets the message field.
func (su *ShoutUpdate) SetMessage(s string) *ShoutUpdate {
	su.mutation.SetMessage(s)
	return su
}

// SetLikes sets the likes field.
func (su *ShoutUpdate) SetLikes(i int) *ShoutUpdate {
	su.mutation.ResetLikes()
	su.mutation.SetLikes(i)
	return su
}

// SetNillableLikes sets the likes field if the given value is not nil.
func (su *ShoutUpdate) SetNillableLikes(i *int) *ShoutUpdate {
	if i != nil {
		su.SetLikes(*i)
	}
	return su
}

// AddLikes adds i to likes.
func (su *ShoutUpdate) AddLikes(i int) *ShoutUpdate {
	su.mutation.AddLikes(i)
	return su
}

// SetAuthorID sets the author edge to User by id.
func (su *ShoutUpdate) SetAuthorID(id int) *ShoutUpdate {
	su.mutation.SetAuthorID(id)
	return su
}

// SetNillableAuthorID sets the author edge to User by id if the given value is not nil.
func (su *ShoutUpdate) SetNillableAuthorID(id *int) *ShoutUpdate {
	if id != nil {
		su = su.SetAuthorID(*id)
	}
	return su
}

// SetAuthor sets the author edge to User.
func (su *ShoutUpdate) SetAuthor(u *User) *ShoutUpdate {
	return su.SetAuthorID(u.ID)
}

// AddLikedByIDs adds the liked_by edge to User by ids.
func (su *ShoutUpdate) AddLikedByIDs(ids ...int) *ShoutUpdate {
	su.mutation.AddLikedByIDs(ids...)
	return su
}

// AddLikedBy adds the liked_by edges to User.
func (su *ShoutUpdate) AddLikedBy(u ...*User) *ShoutUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.AddLikedByIDs(ids...)
}

// Mutation returns the ShoutMutation object of the builder.
func (su *ShoutUpdate) Mutation() *ShoutMutation {
	return su.mutation
}

// ClearAuthor clears the "author" edge to type User.
func (su *ShoutUpdate) ClearAuthor() *ShoutUpdate {
	su.mutation.ClearAuthor()
	return su
}

// ClearLikedBy clears all "liked_by" edges to type User.
func (su *ShoutUpdate) ClearLikedBy() *ShoutUpdate {
	su.mutation.ClearLikedBy()
	return su
}

// RemoveLikedByIDs removes the liked_by edge to User by ids.
func (su *ShoutUpdate) RemoveLikedByIDs(ids ...int) *ShoutUpdate {
	su.mutation.RemoveLikedByIDs(ids...)
	return su
}

// RemoveLikedBy removes liked_by edges to User.
func (su *ShoutUpdate) RemoveLikedBy(u ...*User) *ShoutUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.RemoveLikedByIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *ShoutUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShoutMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ShoutUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ShoutUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ShoutUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *ShoutUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := shout.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *ShoutUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shout.Table,
			Columns: shout.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shout.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shout.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shout.FieldMessage,
		})
	}
	if value, ok := su.mutation.Likes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shout.FieldLikes,
		})
	}
	if value, ok := su.mutation.AddedLikes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shout.FieldLikes,
		})
	}
	if su.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shout.AuthorTable,
			Columns: []string{shout.AuthorColumn},
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
	if nodes := su.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shout.AuthorTable,
			Columns: []string{shout.AuthorColumn},
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
	if su.mutation.LikedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   shout.LikedByTable,
			Columns: shout.LikedByPrimaryKey,
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
	if nodes := su.mutation.RemovedLikedByIDs(); len(nodes) > 0 && !su.mutation.LikedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   shout.LikedByTable,
			Columns: shout.LikedByPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.LikedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   shout.LikedByTable,
			Columns: shout.LikedByPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shout.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ShoutUpdateOne is the builder for updating a single Shout entity.
type ShoutUpdateOne struct {
	config
	hooks    []Hook
	mutation *ShoutMutation
}

// SetUpdatedAt sets the updated_at field.
func (suo *ShoutUpdateOne) SetUpdatedAt(t time.Time) *ShoutUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetMessage sets the message field.
func (suo *ShoutUpdateOne) SetMessage(s string) *ShoutUpdateOne {
	suo.mutation.SetMessage(s)
	return suo
}

// SetLikes sets the likes field.
func (suo *ShoutUpdateOne) SetLikes(i int) *ShoutUpdateOne {
	suo.mutation.ResetLikes()
	suo.mutation.SetLikes(i)
	return suo
}

// SetNillableLikes sets the likes field if the given value is not nil.
func (suo *ShoutUpdateOne) SetNillableLikes(i *int) *ShoutUpdateOne {
	if i != nil {
		suo.SetLikes(*i)
	}
	return suo
}

// AddLikes adds i to likes.
func (suo *ShoutUpdateOne) AddLikes(i int) *ShoutUpdateOne {
	suo.mutation.AddLikes(i)
	return suo
}

// SetAuthorID sets the author edge to User by id.
func (suo *ShoutUpdateOne) SetAuthorID(id int) *ShoutUpdateOne {
	suo.mutation.SetAuthorID(id)
	return suo
}

// SetNillableAuthorID sets the author edge to User by id if the given value is not nil.
func (suo *ShoutUpdateOne) SetNillableAuthorID(id *int) *ShoutUpdateOne {
	if id != nil {
		suo = suo.SetAuthorID(*id)
	}
	return suo
}

// SetAuthor sets the author edge to User.
func (suo *ShoutUpdateOne) SetAuthor(u *User) *ShoutUpdateOne {
	return suo.SetAuthorID(u.ID)
}

// AddLikedByIDs adds the liked_by edge to User by ids.
func (suo *ShoutUpdateOne) AddLikedByIDs(ids ...int) *ShoutUpdateOne {
	suo.mutation.AddLikedByIDs(ids...)
	return suo
}

// AddLikedBy adds the liked_by edges to User.
func (suo *ShoutUpdateOne) AddLikedBy(u ...*User) *ShoutUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.AddLikedByIDs(ids...)
}

// Mutation returns the ShoutMutation object of the builder.
func (suo *ShoutUpdateOne) Mutation() *ShoutMutation {
	return suo.mutation
}

// ClearAuthor clears the "author" edge to type User.
func (suo *ShoutUpdateOne) ClearAuthor() *ShoutUpdateOne {
	suo.mutation.ClearAuthor()
	return suo
}

// ClearLikedBy clears all "liked_by" edges to type User.
func (suo *ShoutUpdateOne) ClearLikedBy() *ShoutUpdateOne {
	suo.mutation.ClearLikedBy()
	return suo
}

// RemoveLikedByIDs removes the liked_by edge to User by ids.
func (suo *ShoutUpdateOne) RemoveLikedByIDs(ids ...int) *ShoutUpdateOne {
	suo.mutation.RemoveLikedByIDs(ids...)
	return suo
}

// RemoveLikedBy removes liked_by edges to User.
func (suo *ShoutUpdateOne) RemoveLikedBy(u ...*User) *ShoutUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.RemoveLikedByIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *ShoutUpdateOne) Save(ctx context.Context) (*Shout, error) {
	var (
		err  error
		node *Shout
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShoutMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ShoutUpdateOne) SaveX(ctx context.Context) *Shout {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ShoutUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ShoutUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *ShoutUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := shout.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *ShoutUpdateOne) sqlSave(ctx context.Context) (_node *Shout, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shout.Table,
			Columns: shout.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shout.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Shout.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shout.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shout.FieldMessage,
		})
	}
	if value, ok := suo.mutation.Likes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shout.FieldLikes,
		})
	}
	if value, ok := suo.mutation.AddedLikes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shout.FieldLikes,
		})
	}
	if suo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shout.AuthorTable,
			Columns: []string{shout.AuthorColumn},
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
	if nodes := suo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shout.AuthorTable,
			Columns: []string{shout.AuthorColumn},
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
	if suo.mutation.LikedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   shout.LikedByTable,
			Columns: shout.LikedByPrimaryKey,
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
	if nodes := suo.mutation.RemovedLikedByIDs(); len(nodes) > 0 && !suo.mutation.LikedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   shout.LikedByTable,
			Columns: shout.LikedByPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.LikedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   shout.LikedByTable,
			Columns: shout.LikedByPrimaryKey,
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
	_node = &Shout{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shout.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}