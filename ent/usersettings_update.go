// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/francismarcus/eg/ent/predicate"
	"github.com/francismarcus/eg/ent/user"
	"github.com/francismarcus/eg/ent/usersettings"
)

// UserSettingsUpdate is the builder for updating UserSettings entities.
type UserSettingsUpdate struct {
	config
	hooks      []Hook
	mutation   *UserSettingsMutation
	predicates []predicate.UserSettings
}

// Where adds a new predicate for the builder.
func (usu *UserSettingsUpdate) Where(ps ...predicate.UserSettings) *UserSettingsUpdate {
	usu.predicates = append(usu.predicates, ps...)
	return usu
}

// SetGender sets the gender field.
func (usu *UserSettingsUpdate) SetGender(u usersettings.Gender) *UserSettingsUpdate {
	usu.mutation.SetGender(u)
	return usu
}

// SetNillableGender sets the gender field if the given value is not nil.
func (usu *UserSettingsUpdate) SetNillableGender(u *usersettings.Gender) *UserSettingsUpdate {
	if u != nil {
		usu.SetGender(*u)
	}
	return usu
}

// ClearGender clears the value of gender.
func (usu *UserSettingsUpdate) ClearGender() *UserSettingsUpdate {
	usu.mutation.ClearGender()
	return usu
}

// SetAge sets the age field.
func (usu *UserSettingsUpdate) SetAge(i int) *UserSettingsUpdate {
	usu.mutation.ResetAge()
	usu.mutation.SetAge(i)
	return usu
}

// SetNillableAge sets the age field if the given value is not nil.
func (usu *UserSettingsUpdate) SetNillableAge(i *int) *UserSettingsUpdate {
	if i != nil {
		usu.SetAge(*i)
	}
	return usu
}

// AddAge adds i to age.
func (usu *UserSettingsUpdate) AddAge(i int) *UserSettingsUpdate {
	usu.mutation.AddAge(i)
	return usu
}

// ClearAge clears the value of age.
func (usu *UserSettingsUpdate) ClearAge() *UserSettingsUpdate {
	usu.mutation.ClearAge()
	return usu
}

// SetWeight sets the weight field.
func (usu *UserSettingsUpdate) SetWeight(i int) *UserSettingsUpdate {
	usu.mutation.ResetWeight()
	usu.mutation.SetWeight(i)
	return usu
}

// SetNillableWeight sets the weight field if the given value is not nil.
func (usu *UserSettingsUpdate) SetNillableWeight(i *int) *UserSettingsUpdate {
	if i != nil {
		usu.SetWeight(*i)
	}
	return usu
}

// AddWeight adds i to weight.
func (usu *UserSettingsUpdate) AddWeight(i int) *UserSettingsUpdate {
	usu.mutation.AddWeight(i)
	return usu
}

// ClearWeight clears the value of weight.
func (usu *UserSettingsUpdate) ClearWeight() *UserSettingsUpdate {
	usu.mutation.ClearWeight()
	return usu
}

// SetHeight sets the height field.
func (usu *UserSettingsUpdate) SetHeight(i int) *UserSettingsUpdate {
	usu.mutation.ResetHeight()
	usu.mutation.SetHeight(i)
	return usu
}

// SetNillableHeight sets the height field if the given value is not nil.
func (usu *UserSettingsUpdate) SetNillableHeight(i *int) *UserSettingsUpdate {
	if i != nil {
		usu.SetHeight(*i)
	}
	return usu
}

// AddHeight adds i to height.
func (usu *UserSettingsUpdate) AddHeight(i int) *UserSettingsUpdate {
	usu.mutation.AddHeight(i)
	return usu
}

// ClearHeight clears the value of height.
func (usu *UserSettingsUpdate) ClearHeight() *UserSettingsUpdate {
	usu.mutation.ClearHeight()
	return usu
}

// SetLevel sets the level field.
func (usu *UserSettingsUpdate) SetLevel(u usersettings.Level) *UserSettingsUpdate {
	usu.mutation.SetLevel(u)
	return usu
}

// SetBelongsToID sets the belongs_to edge to User by id.
func (usu *UserSettingsUpdate) SetBelongsToID(id int) *UserSettingsUpdate {
	usu.mutation.SetBelongsToID(id)
	return usu
}

// SetBelongsTo sets the belongs_to edge to User.
func (usu *UserSettingsUpdate) SetBelongsTo(u *User) *UserSettingsUpdate {
	return usu.SetBelongsToID(u.ID)
}

// Mutation returns the UserSettingsMutation object of the builder.
func (usu *UserSettingsUpdate) Mutation() *UserSettingsMutation {
	return usu.mutation
}

// ClearBelongsTo clears the "belongs_to" edge to type User.
func (usu *UserSettingsUpdate) ClearBelongsTo() *UserSettingsUpdate {
	usu.mutation.ClearBelongsTo()
	return usu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (usu *UserSettingsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(usu.hooks) == 0 {
		if err = usu.check(); err != nil {
			return 0, err
		}
		affected, err = usu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserSettingsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usu.check(); err != nil {
				return 0, err
			}
			usu.mutation = mutation
			affected, err = usu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(usu.hooks) - 1; i >= 0; i-- {
			mut = usu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UserSettingsUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UserSettingsUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UserSettingsUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usu *UserSettingsUpdate) check() error {
	if v, ok := usu.mutation.Gender(); ok {
		if err := usersettings.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	if v, ok := usu.mutation.Level(); ok {
		if err := usersettings.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf("ent: validator failed for field \"level\": %w", err)}
		}
	}
	if _, ok := usu.mutation.BelongsToID(); usu.mutation.BelongsToCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"belongs_to\"")
	}
	return nil
}

func (usu *UserSettingsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usersettings.Table,
			Columns: usersettings.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usersettings.FieldID,
			},
		},
	}
	if ps := usu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: usersettings.FieldGender,
		})
	}
	if usu.mutation.GenderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: usersettings.FieldGender,
		})
	}
	if value, ok := usu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldAge,
		})
	}
	if value, ok := usu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldAge,
		})
	}
	if usu.mutation.AgeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: usersettings.FieldAge,
		})
	}
	if value, ok := usu.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldWeight,
		})
	}
	if value, ok := usu.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldWeight,
		})
	}
	if usu.mutation.WeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: usersettings.FieldWeight,
		})
	}
	if value, ok := usu.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldHeight,
		})
	}
	if value, ok := usu.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldHeight,
		})
	}
	if usu.mutation.HeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: usersettings.FieldHeight,
		})
	}
	if value, ok := usu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: usersettings.FieldLevel,
		})
	}
	if usu.mutation.BelongsToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usersettings.BelongsToTable,
			Columns: []string{usersettings.BelongsToColumn},
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
	if nodes := usu.mutation.BelongsToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usersettings.BelongsToTable,
			Columns: []string{usersettings.BelongsToColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersettings.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserSettingsUpdateOne is the builder for updating a single UserSettings entity.
type UserSettingsUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserSettingsMutation
}

// SetGender sets the gender field.
func (usuo *UserSettingsUpdateOne) SetGender(u usersettings.Gender) *UserSettingsUpdateOne {
	usuo.mutation.SetGender(u)
	return usuo
}

// SetNillableGender sets the gender field if the given value is not nil.
func (usuo *UserSettingsUpdateOne) SetNillableGender(u *usersettings.Gender) *UserSettingsUpdateOne {
	if u != nil {
		usuo.SetGender(*u)
	}
	return usuo
}

// ClearGender clears the value of gender.
func (usuo *UserSettingsUpdateOne) ClearGender() *UserSettingsUpdateOne {
	usuo.mutation.ClearGender()
	return usuo
}

// SetAge sets the age field.
func (usuo *UserSettingsUpdateOne) SetAge(i int) *UserSettingsUpdateOne {
	usuo.mutation.ResetAge()
	usuo.mutation.SetAge(i)
	return usuo
}

// SetNillableAge sets the age field if the given value is not nil.
func (usuo *UserSettingsUpdateOne) SetNillableAge(i *int) *UserSettingsUpdateOne {
	if i != nil {
		usuo.SetAge(*i)
	}
	return usuo
}

// AddAge adds i to age.
func (usuo *UserSettingsUpdateOne) AddAge(i int) *UserSettingsUpdateOne {
	usuo.mutation.AddAge(i)
	return usuo
}

// ClearAge clears the value of age.
func (usuo *UserSettingsUpdateOne) ClearAge() *UserSettingsUpdateOne {
	usuo.mutation.ClearAge()
	return usuo
}

// SetWeight sets the weight field.
func (usuo *UserSettingsUpdateOne) SetWeight(i int) *UserSettingsUpdateOne {
	usuo.mutation.ResetWeight()
	usuo.mutation.SetWeight(i)
	return usuo
}

// SetNillableWeight sets the weight field if the given value is not nil.
func (usuo *UserSettingsUpdateOne) SetNillableWeight(i *int) *UserSettingsUpdateOne {
	if i != nil {
		usuo.SetWeight(*i)
	}
	return usuo
}

// AddWeight adds i to weight.
func (usuo *UserSettingsUpdateOne) AddWeight(i int) *UserSettingsUpdateOne {
	usuo.mutation.AddWeight(i)
	return usuo
}

// ClearWeight clears the value of weight.
func (usuo *UserSettingsUpdateOne) ClearWeight() *UserSettingsUpdateOne {
	usuo.mutation.ClearWeight()
	return usuo
}

// SetHeight sets the height field.
func (usuo *UserSettingsUpdateOne) SetHeight(i int) *UserSettingsUpdateOne {
	usuo.mutation.ResetHeight()
	usuo.mutation.SetHeight(i)
	return usuo
}

// SetNillableHeight sets the height field if the given value is not nil.
func (usuo *UserSettingsUpdateOne) SetNillableHeight(i *int) *UserSettingsUpdateOne {
	if i != nil {
		usuo.SetHeight(*i)
	}
	return usuo
}

// AddHeight adds i to height.
func (usuo *UserSettingsUpdateOne) AddHeight(i int) *UserSettingsUpdateOne {
	usuo.mutation.AddHeight(i)
	return usuo
}

// ClearHeight clears the value of height.
func (usuo *UserSettingsUpdateOne) ClearHeight() *UserSettingsUpdateOne {
	usuo.mutation.ClearHeight()
	return usuo
}

// SetLevel sets the level field.
func (usuo *UserSettingsUpdateOne) SetLevel(u usersettings.Level) *UserSettingsUpdateOne {
	usuo.mutation.SetLevel(u)
	return usuo
}

// SetBelongsToID sets the belongs_to edge to User by id.
func (usuo *UserSettingsUpdateOne) SetBelongsToID(id int) *UserSettingsUpdateOne {
	usuo.mutation.SetBelongsToID(id)
	return usuo
}

// SetBelongsTo sets the belongs_to edge to User.
func (usuo *UserSettingsUpdateOne) SetBelongsTo(u *User) *UserSettingsUpdateOne {
	return usuo.SetBelongsToID(u.ID)
}

// Mutation returns the UserSettingsMutation object of the builder.
func (usuo *UserSettingsUpdateOne) Mutation() *UserSettingsMutation {
	return usuo.mutation
}

// ClearBelongsTo clears the "belongs_to" edge to type User.
func (usuo *UserSettingsUpdateOne) ClearBelongsTo() *UserSettingsUpdateOne {
	usuo.mutation.ClearBelongsTo()
	return usuo
}

// Save executes the query and returns the updated entity.
func (usuo *UserSettingsUpdateOne) Save(ctx context.Context) (*UserSettings, error) {
	var (
		err  error
		node *UserSettings
	)
	if len(usuo.hooks) == 0 {
		if err = usuo.check(); err != nil {
			return nil, err
		}
		node, err = usuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserSettingsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usuo.check(); err != nil {
				return nil, err
			}
			usuo.mutation = mutation
			node, err = usuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(usuo.hooks) - 1; i >= 0; i-- {
			mut = usuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UserSettingsUpdateOne) SaveX(ctx context.Context) *UserSettings {
	node, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usuo *UserSettingsUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UserSettingsUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usuo *UserSettingsUpdateOne) check() error {
	if v, ok := usuo.mutation.Gender(); ok {
		if err := usersettings.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	if v, ok := usuo.mutation.Level(); ok {
		if err := usersettings.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf("ent: validator failed for field \"level\": %w", err)}
		}
	}
	if _, ok := usuo.mutation.BelongsToID(); usuo.mutation.BelongsToCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"belongs_to\"")
	}
	return nil
}

func (usuo *UserSettingsUpdateOne) sqlSave(ctx context.Context) (_node *UserSettings, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usersettings.Table,
			Columns: usersettings.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usersettings.FieldID,
			},
		},
	}
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserSettings.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := usuo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: usersettings.FieldGender,
		})
	}
	if usuo.mutation.GenderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: usersettings.FieldGender,
		})
	}
	if value, ok := usuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldAge,
		})
	}
	if value, ok := usuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldAge,
		})
	}
	if usuo.mutation.AgeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: usersettings.FieldAge,
		})
	}
	if value, ok := usuo.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldWeight,
		})
	}
	if value, ok := usuo.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldWeight,
		})
	}
	if usuo.mutation.WeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: usersettings.FieldWeight,
		})
	}
	if value, ok := usuo.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldHeight,
		})
	}
	if value, ok := usuo.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldHeight,
		})
	}
	if usuo.mutation.HeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: usersettings.FieldHeight,
		})
	}
	if value, ok := usuo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: usersettings.FieldLevel,
		})
	}
	if usuo.mutation.BelongsToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usersettings.BelongsToTable,
			Columns: []string{usersettings.BelongsToColumn},
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
	if nodes := usuo.mutation.BelongsToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usersettings.BelongsToTable,
			Columns: []string{usersettings.BelongsToColumn},
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
	_node = &UserSettings{config: usuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersettings.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
