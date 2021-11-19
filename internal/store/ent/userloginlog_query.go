// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/user"
	"tkserver/internal/store/ent/userloginlog"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserLoginLogQuery is the builder for querying UserLoginLog entities.
type UserLoginLogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserLoginLog
	// eager-loading edges.
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserLoginLogQuery builder.
func (ullq *UserLoginLogQuery) Where(ps ...predicate.UserLoginLog) *UserLoginLogQuery {
	ullq.predicates = append(ullq.predicates, ps...)
	return ullq
}

// Limit adds a limit step to the query.
func (ullq *UserLoginLogQuery) Limit(limit int) *UserLoginLogQuery {
	ullq.limit = &limit
	return ullq
}

// Offset adds an offset step to the query.
func (ullq *UserLoginLogQuery) Offset(offset int) *UserLoginLogQuery {
	ullq.offset = &offset
	return ullq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ullq *UserLoginLogQuery) Unique(unique bool) *UserLoginLogQuery {
	ullq.unique = &unique
	return ullq
}

// Order adds an order step to the query.
func (ullq *UserLoginLogQuery) Order(o ...OrderFunc) *UserLoginLogQuery {
	ullq.order = append(ullq.order, o...)
	return ullq
}

// QueryUser chains the current query on the "user" edge.
func (ullq *UserLoginLogQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: ullq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ullq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ullq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userloginlog.Table, userloginlog.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userloginlog.UserTable, userloginlog.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ullq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserLoginLog entity from the query.
// Returns a *NotFoundError when no UserLoginLog was found.
func (ullq *UserLoginLogQuery) First(ctx context.Context) (*UserLoginLog, error) {
	nodes, err := ullq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userloginlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ullq *UserLoginLogQuery) FirstX(ctx context.Context) *UserLoginLog {
	node, err := ullq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserLoginLog ID from the query.
// Returns a *NotFoundError when no UserLoginLog ID was found.
func (ullq *UserLoginLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ullq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userloginlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ullq *UserLoginLogQuery) FirstIDX(ctx context.Context) int {
	id, err := ullq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserLoginLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UserLoginLog entity is not found.
// Returns a *NotFoundError when no UserLoginLog entities are found.
func (ullq *UserLoginLogQuery) Only(ctx context.Context) (*UserLoginLog, error) {
	nodes, err := ullq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userloginlog.Label}
	default:
		return nil, &NotSingularError{userloginlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ullq *UserLoginLogQuery) OnlyX(ctx context.Context) *UserLoginLog {
	node, err := ullq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserLoginLog ID in the query.
// Returns a *NotSingularError when exactly one UserLoginLog ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ullq *UserLoginLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ullq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userloginlog.Label}
	default:
		err = &NotSingularError{userloginlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ullq *UserLoginLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := ullq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserLoginLogs.
func (ullq *UserLoginLogQuery) All(ctx context.Context) ([]*UserLoginLog, error) {
	if err := ullq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ullq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ullq *UserLoginLogQuery) AllX(ctx context.Context) []*UserLoginLog {
	nodes, err := ullq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserLoginLog IDs.
func (ullq *UserLoginLogQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ullq.Select(userloginlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ullq *UserLoginLogQuery) IDsX(ctx context.Context) []int {
	ids, err := ullq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ullq *UserLoginLogQuery) Count(ctx context.Context) (int, error) {
	if err := ullq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ullq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ullq *UserLoginLogQuery) CountX(ctx context.Context) int {
	count, err := ullq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ullq *UserLoginLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := ullq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ullq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ullq *UserLoginLogQuery) ExistX(ctx context.Context) bool {
	exist, err := ullq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserLoginLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ullq *UserLoginLogQuery) Clone() *UserLoginLogQuery {
	if ullq == nil {
		return nil
	}
	return &UserLoginLogQuery{
		config:     ullq.config,
		limit:      ullq.limit,
		offset:     ullq.offset,
		order:      append([]OrderFunc{}, ullq.order...),
		predicates: append([]predicate.UserLoginLog{}, ullq.predicates...),
		withUser:   ullq.withUser.Clone(),
		// clone intermediate query.
		sql:  ullq.sql.Clone(),
		path: ullq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ullq *UserLoginLogQuery) WithUser(opts ...func(*UserQuery)) *UserLoginLogQuery {
	query := &UserQuery{config: ullq.config}
	for _, opt := range opts {
		opt(query)
	}
	ullq.withUser = query
	return ullq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserLoginLog.Query().
//		GroupBy(userloginlog.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ullq *UserLoginLogQuery) GroupBy(field string, fields ...string) *UserLoginLogGroupBy {
	group := &UserLoginLogGroupBy{config: ullq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ullq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ullq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid"`
//	}
//
//	client.UserLoginLog.Query().
//		Select(userloginlog.FieldUUID).
//		Scan(ctx, &v)
//
func (ullq *UserLoginLogQuery) Select(field string, fields ...string) *UserLoginLogSelect {
	ullq.fields = append([]string{field}, fields...)
	return &UserLoginLogSelect{UserLoginLogQuery: ullq}
}

func (ullq *UserLoginLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ullq.fields {
		if !userloginlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ullq.path != nil {
		prev, err := ullq.path(ctx)
		if err != nil {
			return err
		}
		ullq.sql = prev
	}
	return nil
}

func (ullq *UserLoginLogQuery) sqlAll(ctx context.Context) ([]*UserLoginLog, error) {
	var (
		nodes       = []*UserLoginLog{}
		withFKs     = ullq.withFKs
		_spec       = ullq.querySpec()
		loadedTypes = [1]bool{
			ullq.withUser != nil,
		}
	)
	if ullq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userloginlog.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UserLoginLog{config: ullq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ullq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ullq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UserLoginLog)
		for i := range nodes {
			if nodes[i].user_login_log == nil {
				continue
			}
			fk := *nodes[i].user_login_log
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_login_log" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (ullq *UserLoginLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ullq.querySpec()
	return sqlgraph.CountNodes(ctx, ullq.driver, _spec)
}

func (ullq *UserLoginLogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ullq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ullq *UserLoginLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userloginlog.Table,
			Columns: userloginlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userloginlog.FieldID,
			},
		},
		From:   ullq.sql,
		Unique: true,
	}
	if unique := ullq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ullq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userloginlog.FieldID)
		for i := range fields {
			if fields[i] != userloginlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ullq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ullq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ullq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ullq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ullq *UserLoginLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ullq.driver.Dialect())
	t1 := builder.Table(userloginlog.Table)
	selector := builder.Select(t1.Columns(userloginlog.Columns...)...).From(t1)
	if ullq.sql != nil {
		selector = ullq.sql
		selector.Select(selector.Columns(userloginlog.Columns...)...)
	}
	for _, p := range ullq.predicates {
		p(selector)
	}
	for _, p := range ullq.order {
		p(selector)
	}
	if offset := ullq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ullq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserLoginLogGroupBy is the group-by builder for UserLoginLog entities.
type UserLoginLogGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ullgb *UserLoginLogGroupBy) Aggregate(fns ...AggregateFunc) *UserLoginLogGroupBy {
	ullgb.fns = append(ullgb.fns, fns...)
	return ullgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ullgb *UserLoginLogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ullgb.path(ctx)
	if err != nil {
		return err
	}
	ullgb.sql = query
	return ullgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ullgb *UserLoginLogGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ullgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ullgb *UserLoginLogGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ullgb.fields) > 1 {
		return nil, errors.New("ent: UserLoginLogGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ullgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ullgb *UserLoginLogGroupBy) StringsX(ctx context.Context) []string {
	v, err := ullgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ullgb *UserLoginLogGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ullgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginlog.Label}
	default:
		err = fmt.Errorf("ent: UserLoginLogGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ullgb *UserLoginLogGroupBy) StringX(ctx context.Context) string {
	v, err := ullgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ullgb *UserLoginLogGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ullgb.fields) > 1 {
		return nil, errors.New("ent: UserLoginLogGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ullgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ullgb *UserLoginLogGroupBy) IntsX(ctx context.Context) []int {
	v, err := ullgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ullgb *UserLoginLogGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ullgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginlog.Label}
	default:
		err = fmt.Errorf("ent: UserLoginLogGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ullgb *UserLoginLogGroupBy) IntX(ctx context.Context) int {
	v, err := ullgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ullgb *UserLoginLogGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ullgb.fields) > 1 {
		return nil, errors.New("ent: UserLoginLogGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ullgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ullgb *UserLoginLogGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ullgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ullgb *UserLoginLogGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ullgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginlog.Label}
	default:
		err = fmt.Errorf("ent: UserLoginLogGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ullgb *UserLoginLogGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ullgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ullgb *UserLoginLogGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ullgb.fields) > 1 {
		return nil, errors.New("ent: UserLoginLogGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ullgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ullgb *UserLoginLogGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ullgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ullgb *UserLoginLogGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ullgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginlog.Label}
	default:
		err = fmt.Errorf("ent: UserLoginLogGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ullgb *UserLoginLogGroupBy) BoolX(ctx context.Context) bool {
	v, err := ullgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ullgb *UserLoginLogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ullgb.fields {
		if !userloginlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ullgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ullgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ullgb *UserLoginLogGroupBy) sqlQuery() *sql.Selector {
	selector := ullgb.sql
	columns := make([]string, 0, len(ullgb.fields)+len(ullgb.fns))
	columns = append(columns, ullgb.fields...)
	for _, fn := range ullgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ullgb.fields...)
}

// UserLoginLogSelect is the builder for selecting fields of UserLoginLog entities.
type UserLoginLogSelect struct {
	*UserLoginLogQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ulls *UserLoginLogSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ulls.prepareQuery(ctx); err != nil {
		return err
	}
	ulls.sql = ulls.UserLoginLogQuery.sqlQuery(ctx)
	return ulls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ulls *UserLoginLogSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ulls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ulls *UserLoginLogSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ulls.fields) > 1 {
		return nil, errors.New("ent: UserLoginLogSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ulls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ulls *UserLoginLogSelect) StringsX(ctx context.Context) []string {
	v, err := ulls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ulls *UserLoginLogSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ulls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginlog.Label}
	default:
		err = fmt.Errorf("ent: UserLoginLogSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ulls *UserLoginLogSelect) StringX(ctx context.Context) string {
	v, err := ulls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ulls *UserLoginLogSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ulls.fields) > 1 {
		return nil, errors.New("ent: UserLoginLogSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ulls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ulls *UserLoginLogSelect) IntsX(ctx context.Context) []int {
	v, err := ulls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ulls *UserLoginLogSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ulls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginlog.Label}
	default:
		err = fmt.Errorf("ent: UserLoginLogSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ulls *UserLoginLogSelect) IntX(ctx context.Context) int {
	v, err := ulls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ulls *UserLoginLogSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ulls.fields) > 1 {
		return nil, errors.New("ent: UserLoginLogSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ulls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ulls *UserLoginLogSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ulls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ulls *UserLoginLogSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ulls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginlog.Label}
	default:
		err = fmt.Errorf("ent: UserLoginLogSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ulls *UserLoginLogSelect) Float64X(ctx context.Context) float64 {
	v, err := ulls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ulls *UserLoginLogSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ulls.fields) > 1 {
		return nil, errors.New("ent: UserLoginLogSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ulls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ulls *UserLoginLogSelect) BoolsX(ctx context.Context) []bool {
	v, err := ulls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ulls *UserLoginLogSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ulls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginlog.Label}
	default:
		err = fmt.Errorf("ent: UserLoginLogSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ulls *UserLoginLogSelect) BoolX(ctx context.Context) bool {
	v, err := ulls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ulls *UserLoginLogSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ulls.sqlQuery().Query()
	if err := ulls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ulls *UserLoginLogSelect) sqlQuery() sql.Querier {
	selector := ulls.sql
	selector.Select(selector.Columns(ulls.fields...)...)
	return selector
}
