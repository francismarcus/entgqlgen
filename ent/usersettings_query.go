// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/francismarcus/eg/ent/predicate"
	"github.com/francismarcus/eg/ent/user"
	"github.com/francismarcus/eg/ent/usersettings"
)

// UserSettingsQuery is the builder for querying UserSettings entities.
type UserSettingsQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.UserSettings
	// eager-loading edges.
	withBelongsTo *UserQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (usq *UserSettingsQuery) Where(ps ...predicate.UserSettings) *UserSettingsQuery {
	usq.predicates = append(usq.predicates, ps...)
	return usq
}

// Limit adds a limit step to the query.
func (usq *UserSettingsQuery) Limit(limit int) *UserSettingsQuery {
	usq.limit = &limit
	return usq
}

// Offset adds an offset step to the query.
func (usq *UserSettingsQuery) Offset(offset int) *UserSettingsQuery {
	usq.offset = &offset
	return usq
}

// Order adds an order step to the query.
func (usq *UserSettingsQuery) Order(o ...OrderFunc) *UserSettingsQuery {
	usq.order = append(usq.order, o...)
	return usq
}

// QueryBelongsTo chains the current query on the belongs_to edge.
func (usq *UserSettingsQuery) QueryBelongsTo() *UserQuery {
	query := &UserQuery{config: usq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := usq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usersettings.Table, usersettings.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, usersettings.BelongsToTable, usersettings.BelongsToColumn),
		)
		fromU = sqlgraph.SetNeighbors(usq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserSettings entity in the query. Returns *NotFoundError when no usersettings was found.
func (usq *UserSettingsQuery) First(ctx context.Context) (*UserSettings, error) {
	nodes, err := usq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usersettings.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (usq *UserSettingsQuery) FirstX(ctx context.Context) *UserSettings {
	node, err := usq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserSettings id in the query. Returns *NotFoundError when no id was found.
func (usq *UserSettingsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = usq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usersettings.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (usq *UserSettingsQuery) FirstXID(ctx context.Context) int {
	id, err := usq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only UserSettings entity in the query, returns an error if not exactly one entity was returned.
func (usq *UserSettingsQuery) Only(ctx context.Context) (*UserSettings, error) {
	nodes, err := usq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usersettings.Label}
	default:
		return nil, &NotSingularError{usersettings.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (usq *UserSettingsQuery) OnlyX(ctx context.Context) *UserSettings {
	node, err := usq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only UserSettings id in the query, returns an error if not exactly one id was returned.
func (usq *UserSettingsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = usq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usersettings.Label}
	default:
		err = &NotSingularError{usersettings.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (usq *UserSettingsQuery) OnlyIDX(ctx context.Context) int {
	id, err := usq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserSettingsSlice.
func (usq *UserSettingsQuery) All(ctx context.Context) ([]*UserSettings, error) {
	if err := usq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return usq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (usq *UserSettingsQuery) AllX(ctx context.Context) []*UserSettings {
	nodes, err := usq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserSettings ids.
func (usq *UserSettingsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := usq.Select(usersettings.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (usq *UserSettingsQuery) IDsX(ctx context.Context) []int {
	ids, err := usq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (usq *UserSettingsQuery) Count(ctx context.Context) (int, error) {
	if err := usq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return usq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (usq *UserSettingsQuery) CountX(ctx context.Context) int {
	count, err := usq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (usq *UserSettingsQuery) Exist(ctx context.Context) (bool, error) {
	if err := usq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return usq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (usq *UserSettingsQuery) ExistX(ctx context.Context) bool {
	exist, err := usq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (usq *UserSettingsQuery) Clone() *UserSettingsQuery {
	return &UserSettingsQuery{
		config:     usq.config,
		limit:      usq.limit,
		offset:     usq.offset,
		order:      append([]OrderFunc{}, usq.order...),
		unique:     append([]string{}, usq.unique...),
		predicates: append([]predicate.UserSettings{}, usq.predicates...),
		// clone intermediate query.
		sql:  usq.sql.Clone(),
		path: usq.path,
	}
}

//  WithBelongsTo tells the query-builder to eager-loads the nodes that are connected to
// the "belongs_to" edge. The optional arguments used to configure the query builder of the edge.
func (usq *UserSettingsQuery) WithBelongsTo(opts ...func(*UserQuery)) *UserSettingsQuery {
	query := &UserQuery{config: usq.config}
	for _, opt := range opts {
		opt(query)
	}
	usq.withBelongsTo = query
	return usq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Gender usersettings.Gender `json:"gender,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserSettings.Query().
//		GroupBy(usersettings.FieldGender).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (usq *UserSettingsQuery) GroupBy(field string, fields ...string) *UserSettingsGroupBy {
	group := &UserSettingsGroupBy{config: usq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return usq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Gender usersettings.Gender `json:"gender,omitempty"`
//	}
//
//	client.UserSettings.Query().
//		Select(usersettings.FieldGender).
//		Scan(ctx, &v)
//
func (usq *UserSettingsQuery) Select(field string, fields ...string) *UserSettingsSelect {
	selector := &UserSettingsSelect{config: usq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return usq.sqlQuery(), nil
	}
	return selector
}

func (usq *UserSettingsQuery) prepareQuery(ctx context.Context) error {
	if usq.path != nil {
		prev, err := usq.path(ctx)
		if err != nil {
			return err
		}
		usq.sql = prev
	}
	return nil
}

func (usq *UserSettingsQuery) sqlAll(ctx context.Context) ([]*UserSettings, error) {
	var (
		nodes       = []*UserSettings{}
		withFKs     = usq.withFKs
		_spec       = usq.querySpec()
		loadedTypes = [1]bool{
			usq.withBelongsTo != nil,
		}
	)
	if usq.withBelongsTo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, usersettings.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &UserSettings{config: usq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, usq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := usq.withBelongsTo; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UserSettings)
		for i := range nodes {
			if fk := nodes[i].user_settings; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_settings" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.BelongsTo = n
			}
		}
	}

	return nodes, nil
}

func (usq *UserSettingsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := usq.querySpec()
	return sqlgraph.CountNodes(ctx, usq.driver, _spec)
}

func (usq *UserSettingsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := usq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (usq *UserSettingsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usersettings.Table,
			Columns: usersettings.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usersettings.FieldID,
			},
		},
		From:   usq.sql,
		Unique: true,
	}
	if ps := usq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := usq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := usq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := usq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, usersettings.ValidColumn)
			}
		}
	}
	return _spec
}

func (usq *UserSettingsQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(usq.driver.Dialect())
	t1 := builder.Table(usersettings.Table)
	selector := builder.Select(t1.Columns(usersettings.Columns...)...).From(t1)
	if usq.sql != nil {
		selector = usq.sql
		selector.Select(selector.Columns(usersettings.Columns...)...)
	}
	for _, p := range usq.predicates {
		p(selector)
	}
	for _, p := range usq.order {
		p(selector, usersettings.ValidColumn)
	}
	if offset := usq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := usq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserSettingsGroupBy is the builder for group-by UserSettings entities.
type UserSettingsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (usgb *UserSettingsGroupBy) Aggregate(fns ...AggregateFunc) *UserSettingsGroupBy {
	usgb.fns = append(usgb.fns, fns...)
	return usgb
}

// Scan applies the group-by query and scan the result into the given value.
func (usgb *UserSettingsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := usgb.path(ctx)
	if err != nil {
		return err
	}
	usgb.sql = query
	return usgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (usgb *UserSettingsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := usgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserSettingsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(usgb.fields) > 1 {
		return nil, errors.New("ent: UserSettingsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := usgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (usgb *UserSettingsGroupBy) StringsX(ctx context.Context) []string {
	v, err := usgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserSettingsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = usgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usersettings.Label}
	default:
		err = fmt.Errorf("ent: UserSettingsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (usgb *UserSettingsGroupBy) StringX(ctx context.Context) string {
	v, err := usgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserSettingsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(usgb.fields) > 1 {
		return nil, errors.New("ent: UserSettingsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := usgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (usgb *UserSettingsGroupBy) IntsX(ctx context.Context) []int {
	v, err := usgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserSettingsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = usgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usersettings.Label}
	default:
		err = fmt.Errorf("ent: UserSettingsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (usgb *UserSettingsGroupBy) IntX(ctx context.Context) int {
	v, err := usgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserSettingsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(usgb.fields) > 1 {
		return nil, errors.New("ent: UserSettingsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := usgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (usgb *UserSettingsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := usgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserSettingsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = usgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usersettings.Label}
	default:
		err = fmt.Errorf("ent: UserSettingsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (usgb *UserSettingsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := usgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserSettingsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(usgb.fields) > 1 {
		return nil, errors.New("ent: UserSettingsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := usgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (usgb *UserSettingsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := usgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserSettingsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = usgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usersettings.Label}
	default:
		err = fmt.Errorf("ent: UserSettingsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (usgb *UserSettingsGroupBy) BoolX(ctx context.Context) bool {
	v, err := usgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (usgb *UserSettingsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range usgb.fields {
		if !usersettings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := usgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := usgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (usgb *UserSettingsGroupBy) sqlQuery() *sql.Selector {
	selector := usgb.sql
	columns := make([]string, 0, len(usgb.fields)+len(usgb.fns))
	columns = append(columns, usgb.fields...)
	for _, fn := range usgb.fns {
		columns = append(columns, fn(selector, usersettings.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(usgb.fields...)
}

// UserSettingsSelect is the builder for select fields of UserSettings entities.
type UserSettingsSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (uss *UserSettingsSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := uss.path(ctx)
	if err != nil {
		return err
	}
	uss.sql = query
	return uss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uss *UserSettingsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (uss *UserSettingsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uss.fields) > 1 {
		return nil, errors.New("ent: UserSettingsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uss *UserSettingsSelect) StringsX(ctx context.Context) []string {
	v, err := uss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (uss *UserSettingsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usersettings.Label}
	default:
		err = fmt.Errorf("ent: UserSettingsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uss *UserSettingsSelect) StringX(ctx context.Context) string {
	v, err := uss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (uss *UserSettingsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uss.fields) > 1 {
		return nil, errors.New("ent: UserSettingsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uss *UserSettingsSelect) IntsX(ctx context.Context) []int {
	v, err := uss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (uss *UserSettingsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usersettings.Label}
	default:
		err = fmt.Errorf("ent: UserSettingsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uss *UserSettingsSelect) IntX(ctx context.Context) int {
	v, err := uss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (uss *UserSettingsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uss.fields) > 1 {
		return nil, errors.New("ent: UserSettingsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uss *UserSettingsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (uss *UserSettingsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usersettings.Label}
	default:
		err = fmt.Errorf("ent: UserSettingsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uss *UserSettingsSelect) Float64X(ctx context.Context) float64 {
	v, err := uss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (uss *UserSettingsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uss.fields) > 1 {
		return nil, errors.New("ent: UserSettingsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uss *UserSettingsSelect) BoolsX(ctx context.Context) []bool {
	v, err := uss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (uss *UserSettingsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usersettings.Label}
	default:
		err = fmt.Errorf("ent: UserSettingsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uss *UserSettingsSelect) BoolX(ctx context.Context) bool {
	v, err := uss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uss *UserSettingsSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uss.fields {
		if !usersettings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := uss.sqlQuery().Query()
	if err := uss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uss *UserSettingsSelect) sqlQuery() sql.Querier {
	selector := uss.sql
	selector.Select(selector.Columns(uss.fields...)...)
	return selector
}