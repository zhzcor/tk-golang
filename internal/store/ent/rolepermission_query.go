// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/role"
	"tkserver/internal/store/ent/rolepermission"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RolePermissionQuery is the builder for querying RolePermission entities.
type RolePermissionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.RolePermission
	// eager-loading edges.
	withRole *RoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RolePermissionQuery builder.
func (rpq *RolePermissionQuery) Where(ps ...predicate.RolePermission) *RolePermissionQuery {
	rpq.predicates = append(rpq.predicates, ps...)
	return rpq
}

// Limit adds a limit step to the query.
func (rpq *RolePermissionQuery) Limit(limit int) *RolePermissionQuery {
	rpq.limit = &limit
	return rpq
}

// Offset adds an offset step to the query.
func (rpq *RolePermissionQuery) Offset(offset int) *RolePermissionQuery {
	rpq.offset = &offset
	return rpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rpq *RolePermissionQuery) Unique(unique bool) *RolePermissionQuery {
	rpq.unique = &unique
	return rpq
}

// Order adds an order step to the query.
func (rpq *RolePermissionQuery) Order(o ...OrderFunc) *RolePermissionQuery {
	rpq.order = append(rpq.order, o...)
	return rpq
}

// QueryRole chains the current query on the "role" edge.
func (rpq *RolePermissionQuery) QueryRole() *RoleQuery {
	query := &RoleQuery{config: rpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rolepermission.Table, rolepermission.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, rolepermission.RoleTable, rolepermission.RoleColumn),
		)
		fromU = sqlgraph.SetNeighbors(rpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RolePermission entity from the query.
// Returns a *NotFoundError when no RolePermission was found.
func (rpq *RolePermissionQuery) First(ctx context.Context) (*RolePermission, error) {
	nodes, err := rpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rolepermission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rpq *RolePermissionQuery) FirstX(ctx context.Context) *RolePermission {
	node, err := rpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RolePermission ID from the query.
// Returns a *NotFoundError when no RolePermission ID was found.
func (rpq *RolePermissionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rolepermission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rpq *RolePermissionQuery) FirstIDX(ctx context.Context) int {
	id, err := rpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RolePermission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one RolePermission entity is not found.
// Returns a *NotFoundError when no RolePermission entities are found.
func (rpq *RolePermissionQuery) Only(ctx context.Context) (*RolePermission, error) {
	nodes, err := rpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rolepermission.Label}
	default:
		return nil, &NotSingularError{rolepermission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rpq *RolePermissionQuery) OnlyX(ctx context.Context) *RolePermission {
	node, err := rpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RolePermission ID in the query.
// Returns a *NotSingularError when exactly one RolePermission ID is not found.
// Returns a *NotFoundError when no entities are found.
func (rpq *RolePermissionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rolepermission.Label}
	default:
		err = &NotSingularError{rolepermission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rpq *RolePermissionQuery) OnlyIDX(ctx context.Context) int {
	id, err := rpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RolePermissions.
func (rpq *RolePermissionQuery) All(ctx context.Context) ([]*RolePermission, error) {
	if err := rpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rpq *RolePermissionQuery) AllX(ctx context.Context) []*RolePermission {
	nodes, err := rpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RolePermission IDs.
func (rpq *RolePermissionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rpq.Select(rolepermission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rpq *RolePermissionQuery) IDsX(ctx context.Context) []int {
	ids, err := rpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rpq *RolePermissionQuery) Count(ctx context.Context) (int, error) {
	if err := rpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rpq *RolePermissionQuery) CountX(ctx context.Context) int {
	count, err := rpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rpq *RolePermissionQuery) Exist(ctx context.Context) (bool, error) {
	if err := rpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rpq *RolePermissionQuery) ExistX(ctx context.Context) bool {
	exist, err := rpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RolePermissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rpq *RolePermissionQuery) Clone() *RolePermissionQuery {
	if rpq == nil {
		return nil
	}
	return &RolePermissionQuery{
		config:     rpq.config,
		limit:      rpq.limit,
		offset:     rpq.offset,
		order:      append([]OrderFunc{}, rpq.order...),
		predicates: append([]predicate.RolePermission{}, rpq.predicates...),
		withRole:   rpq.withRole.Clone(),
		// clone intermediate query.
		sql:  rpq.sql.Clone(),
		path: rpq.path,
	}
}

// WithRole tells the query-builder to eager-load the nodes that are connected to
// the "role" edge. The optional arguments are used to configure the query builder of the edge.
func (rpq *RolePermissionQuery) WithRole(opts ...func(*RoleQuery)) *RolePermissionQuery {
	query := &RoleQuery{config: rpq.config}
	for _, opt := range opts {
		opt(query)
	}
	rpq.withRole = query
	return rpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RoleID int `json:"role_id"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RolePermission.Query().
//		GroupBy(rolepermission.FieldRoleID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rpq *RolePermissionQuery) GroupBy(field string, fields ...string) *RolePermissionGroupBy {
	group := &RolePermissionGroupBy{config: rpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RoleID int `json:"role_id"`
//	}
//
//	client.RolePermission.Query().
//		Select(rolepermission.FieldRoleID).
//		Scan(ctx, &v)
//
func (rpq *RolePermissionQuery) Select(field string, fields ...string) *RolePermissionSelect {
	rpq.fields = append([]string{field}, fields...)
	return &RolePermissionSelect{RolePermissionQuery: rpq}
}

func (rpq *RolePermissionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rpq.fields {
		if !rolepermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rpq.path != nil {
		prev, err := rpq.path(ctx)
		if err != nil {
			return err
		}
		rpq.sql = prev
	}
	return nil
}

func (rpq *RolePermissionQuery) sqlAll(ctx context.Context) ([]*RolePermission, error) {
	var (
		nodes       = []*RolePermission{}
		_spec       = rpq.querySpec()
		loadedTypes = [1]bool{
			rpq.withRole != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &RolePermission{config: rpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, rpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rpq.withRole; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*RolePermission)
		for i := range nodes {
			fk := nodes[i].RoleID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(role.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "role_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Role = n
			}
		}
	}

	return nodes, nil
}

func (rpq *RolePermissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rpq.querySpec()
	return sqlgraph.CountNodes(ctx, rpq.driver, _spec)
}

func (rpq *RolePermissionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (rpq *RolePermissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rolepermission.Table,
			Columns: rolepermission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rolepermission.FieldID,
			},
		},
		From:   rpq.sql,
		Unique: true,
	}
	if unique := rpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rolepermission.FieldID)
		for i := range fields {
			if fields[i] != rolepermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rpq *RolePermissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rpq.driver.Dialect())
	t1 := builder.Table(rolepermission.Table)
	selector := builder.Select(t1.Columns(rolepermission.Columns...)...).From(t1)
	if rpq.sql != nil {
		selector = rpq.sql
		selector.Select(selector.Columns(rolepermission.Columns...)...)
	}
	for _, p := range rpq.predicates {
		p(selector)
	}
	for _, p := range rpq.order {
		p(selector)
	}
	if offset := rpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RolePermissionGroupBy is the group-by builder for RolePermission entities.
type RolePermissionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rpgb *RolePermissionGroupBy) Aggregate(fns ...AggregateFunc) *RolePermissionGroupBy {
	rpgb.fns = append(rpgb.fns, fns...)
	return rpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rpgb *RolePermissionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rpgb.path(ctx)
	if err != nil {
		return err
	}
	rpgb.sql = query
	return rpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rpgb *RolePermissionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RolePermissionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rpgb.fields) > 1 {
		return nil, errors.New("ent: RolePermissionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rpgb *RolePermissionGroupBy) StringsX(ctx context.Context) []string {
	v, err := rpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RolePermissionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolepermission.Label}
	default:
		err = fmt.Errorf("ent: RolePermissionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rpgb *RolePermissionGroupBy) StringX(ctx context.Context) string {
	v, err := rpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RolePermissionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rpgb.fields) > 1 {
		return nil, errors.New("ent: RolePermissionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rpgb *RolePermissionGroupBy) IntsX(ctx context.Context) []int {
	v, err := rpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RolePermissionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolepermission.Label}
	default:
		err = fmt.Errorf("ent: RolePermissionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rpgb *RolePermissionGroupBy) IntX(ctx context.Context) int {
	v, err := rpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RolePermissionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rpgb.fields) > 1 {
		return nil, errors.New("ent: RolePermissionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rpgb *RolePermissionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RolePermissionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolepermission.Label}
	default:
		err = fmt.Errorf("ent: RolePermissionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rpgb *RolePermissionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RolePermissionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rpgb.fields) > 1 {
		return nil, errors.New("ent: RolePermissionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rpgb *RolePermissionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RolePermissionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolepermission.Label}
	default:
		err = fmt.Errorf("ent: RolePermissionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rpgb *RolePermissionGroupBy) BoolX(ctx context.Context) bool {
	v, err := rpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rpgb *RolePermissionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rpgb.fields {
		if !rolepermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rpgb *RolePermissionGroupBy) sqlQuery() *sql.Selector {
	selector := rpgb.sql
	columns := make([]string, 0, len(rpgb.fields)+len(rpgb.fns))
	columns = append(columns, rpgb.fields...)
	for _, fn := range rpgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rpgb.fields...)
}

// RolePermissionSelect is the builder for selecting fields of RolePermission entities.
type RolePermissionSelect struct {
	*RolePermissionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rps *RolePermissionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rps.prepareQuery(ctx); err != nil {
		return err
	}
	rps.sql = rps.RolePermissionQuery.sqlQuery(ctx)
	return rps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rps *RolePermissionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rps *RolePermissionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rps.fields) > 1 {
		return nil, errors.New("ent: RolePermissionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rps *RolePermissionSelect) StringsX(ctx context.Context) []string {
	v, err := rps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rps *RolePermissionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolepermission.Label}
	default:
		err = fmt.Errorf("ent: RolePermissionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rps *RolePermissionSelect) StringX(ctx context.Context) string {
	v, err := rps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rps *RolePermissionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rps.fields) > 1 {
		return nil, errors.New("ent: RolePermissionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rps *RolePermissionSelect) IntsX(ctx context.Context) []int {
	v, err := rps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rps *RolePermissionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolepermission.Label}
	default:
		err = fmt.Errorf("ent: RolePermissionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rps *RolePermissionSelect) IntX(ctx context.Context) int {
	v, err := rps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rps *RolePermissionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rps.fields) > 1 {
		return nil, errors.New("ent: RolePermissionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rps *RolePermissionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rps *RolePermissionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolepermission.Label}
	default:
		err = fmt.Errorf("ent: RolePermissionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rps *RolePermissionSelect) Float64X(ctx context.Context) float64 {
	v, err := rps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rps *RolePermissionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rps.fields) > 1 {
		return nil, errors.New("ent: RolePermissionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rps *RolePermissionSelect) BoolsX(ctx context.Context) []bool {
	v, err := rps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rps *RolePermissionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolepermission.Label}
	default:
		err = fmt.Errorf("ent: RolePermissionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rps *RolePermissionSelect) BoolX(ctx context.Context) bool {
	v, err := rps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rps *RolePermissionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rps.sqlQuery().Query()
	if err := rps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rps *RolePermissionSelect) sqlQuery() sql.Querier {
	selector := rps.sql
	selector.Select(selector.Columns(rps.fields...)...)
	return selector
}
