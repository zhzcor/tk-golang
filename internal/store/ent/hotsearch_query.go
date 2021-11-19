// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/hotsearch"
	"gserver/internal/store/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HotSearchQuery is the builder for querying HotSearch entities.
type HotSearchQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.HotSearch
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HotSearchQuery builder.
func (hsq *HotSearchQuery) Where(ps ...predicate.HotSearch) *HotSearchQuery {
	hsq.predicates = append(hsq.predicates, ps...)
	return hsq
}

// Limit adds a limit step to the query.
func (hsq *HotSearchQuery) Limit(limit int) *HotSearchQuery {
	hsq.limit = &limit
	return hsq
}

// Offset adds an offset step to the query.
func (hsq *HotSearchQuery) Offset(offset int) *HotSearchQuery {
	hsq.offset = &offset
	return hsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hsq *HotSearchQuery) Unique(unique bool) *HotSearchQuery {
	hsq.unique = &unique
	return hsq
}

// Order adds an order step to the query.
func (hsq *HotSearchQuery) Order(o ...OrderFunc) *HotSearchQuery {
	hsq.order = append(hsq.order, o...)
	return hsq
}

// First returns the first HotSearch entity from the query.
// Returns a *NotFoundError when no HotSearch was found.
func (hsq *HotSearchQuery) First(ctx context.Context) (*HotSearch, error) {
	nodes, err := hsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hotsearch.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hsq *HotSearchQuery) FirstX(ctx context.Context) *HotSearch {
	node, err := hsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HotSearch ID from the query.
// Returns a *NotFoundError when no HotSearch ID was found.
func (hsq *HotSearchQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hotsearch.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hsq *HotSearchQuery) FirstIDX(ctx context.Context) int {
	id, err := hsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HotSearch entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one HotSearch entity is not found.
// Returns a *NotFoundError when no HotSearch entities are found.
func (hsq *HotSearchQuery) Only(ctx context.Context) (*HotSearch, error) {
	nodes, err := hsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hotsearch.Label}
	default:
		return nil, &NotSingularError{hotsearch.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hsq *HotSearchQuery) OnlyX(ctx context.Context) *HotSearch {
	node, err := hsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HotSearch ID in the query.
// Returns a *NotSingularError when exactly one HotSearch ID is not found.
// Returns a *NotFoundError when no entities are found.
func (hsq *HotSearchQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hotsearch.Label}
	default:
		err = &NotSingularError{hotsearch.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hsq *HotSearchQuery) OnlyIDX(ctx context.Context) int {
	id, err := hsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HotSearches.
func (hsq *HotSearchQuery) All(ctx context.Context) ([]*HotSearch, error) {
	if err := hsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return hsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (hsq *HotSearchQuery) AllX(ctx context.Context) []*HotSearch {
	nodes, err := hsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HotSearch IDs.
func (hsq *HotSearchQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := hsq.Select(hotsearch.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hsq *HotSearchQuery) IDsX(ctx context.Context) []int {
	ids, err := hsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hsq *HotSearchQuery) Count(ctx context.Context) (int, error) {
	if err := hsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return hsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (hsq *HotSearchQuery) CountX(ctx context.Context) int {
	count, err := hsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hsq *HotSearchQuery) Exist(ctx context.Context) (bool, error) {
	if err := hsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return hsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (hsq *HotSearchQuery) ExistX(ctx context.Context) bool {
	exist, err := hsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HotSearchQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hsq *HotSearchQuery) Clone() *HotSearchQuery {
	if hsq == nil {
		return nil
	}
	return &HotSearchQuery{
		config:     hsq.config,
		limit:      hsq.limit,
		offset:     hsq.offset,
		order:      append([]OrderFunc{}, hsq.order...),
		predicates: append([]predicate.HotSearch{}, hsq.predicates...),
		// clone intermediate query.
		sql:  hsq.sql.Clone(),
		path: hsq.path,
	}
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
//	client.HotSearch.Query().
//		GroupBy(hotsearch.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (hsq *HotSearchQuery) GroupBy(field string, fields ...string) *HotSearchGroupBy {
	group := &HotSearchGroupBy{config: hsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := hsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return hsq.sqlQuery(ctx), nil
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
//	client.HotSearch.Query().
//		Select(hotsearch.FieldUUID).
//		Scan(ctx, &v)
//
func (hsq *HotSearchQuery) Select(field string, fields ...string) *HotSearchSelect {
	hsq.fields = append([]string{field}, fields...)
	return &HotSearchSelect{HotSearchQuery: hsq}
}

func (hsq *HotSearchQuery) prepareQuery(ctx context.Context) error {
	for _, f := range hsq.fields {
		if !hotsearch.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hsq.path != nil {
		prev, err := hsq.path(ctx)
		if err != nil {
			return err
		}
		hsq.sql = prev
	}
	return nil
}

func (hsq *HotSearchQuery) sqlAll(ctx context.Context) ([]*HotSearch, error) {
	var (
		nodes = []*HotSearch{}
		_spec = hsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &HotSearch{config: hsq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, hsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (hsq *HotSearchQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hsq.querySpec()
	return sqlgraph.CountNodes(ctx, hsq.driver, _spec)
}

func (hsq *HotSearchQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := hsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (hsq *HotSearchQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hotsearch.Table,
			Columns: hotsearch.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hotsearch.FieldID,
			},
		},
		From:   hsq.sql,
		Unique: true,
	}
	if unique := hsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := hsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hotsearch.FieldID)
		for i := range fields {
			if fields[i] != hotsearch.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hsq *HotSearchQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hsq.driver.Dialect())
	t1 := builder.Table(hotsearch.Table)
	selector := builder.Select(t1.Columns(hotsearch.Columns...)...).From(t1)
	if hsq.sql != nil {
		selector = hsq.sql
		selector.Select(selector.Columns(hotsearch.Columns...)...)
	}
	for _, p := range hsq.predicates {
		p(selector)
	}
	for _, p := range hsq.order {
		p(selector)
	}
	if offset := hsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HotSearchGroupBy is the group-by builder for HotSearch entities.
type HotSearchGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hsgb *HotSearchGroupBy) Aggregate(fns ...AggregateFunc) *HotSearchGroupBy {
	hsgb.fns = append(hsgb.fns, fns...)
	return hsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (hsgb *HotSearchGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := hsgb.path(ctx)
	if err != nil {
		return err
	}
	hsgb.sql = query
	return hsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hsgb *HotSearchGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := hsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (hsgb *HotSearchGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(hsgb.fields) > 1 {
		return nil, errors.New("ent: HotSearchGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := hsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hsgb *HotSearchGroupBy) StringsX(ctx context.Context) []string {
	v, err := hsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hsgb *HotSearchGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hotsearch.Label}
	default:
		err = fmt.Errorf("ent: HotSearchGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hsgb *HotSearchGroupBy) StringX(ctx context.Context) string {
	v, err := hsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (hsgb *HotSearchGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(hsgb.fields) > 1 {
		return nil, errors.New("ent: HotSearchGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := hsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hsgb *HotSearchGroupBy) IntsX(ctx context.Context) []int {
	v, err := hsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hsgb *HotSearchGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hotsearch.Label}
	default:
		err = fmt.Errorf("ent: HotSearchGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hsgb *HotSearchGroupBy) IntX(ctx context.Context) int {
	v, err := hsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (hsgb *HotSearchGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(hsgb.fields) > 1 {
		return nil, errors.New("ent: HotSearchGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := hsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hsgb *HotSearchGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := hsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hsgb *HotSearchGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hotsearch.Label}
	default:
		err = fmt.Errorf("ent: HotSearchGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hsgb *HotSearchGroupBy) Float64X(ctx context.Context) float64 {
	v, err := hsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (hsgb *HotSearchGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(hsgb.fields) > 1 {
		return nil, errors.New("ent: HotSearchGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := hsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hsgb *HotSearchGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := hsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hsgb *HotSearchGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hotsearch.Label}
	default:
		err = fmt.Errorf("ent: HotSearchGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hsgb *HotSearchGroupBy) BoolX(ctx context.Context) bool {
	v, err := hsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hsgb *HotSearchGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range hsgb.fields {
		if !hotsearch.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := hsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hsgb *HotSearchGroupBy) sqlQuery() *sql.Selector {
	selector := hsgb.sql
	columns := make([]string, 0, len(hsgb.fields)+len(hsgb.fns))
	columns = append(columns, hsgb.fields...)
	for _, fn := range hsgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(hsgb.fields...)
}

// HotSearchSelect is the builder for selecting fields of HotSearch entities.
type HotSearchSelect struct {
	*HotSearchQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (hss *HotSearchSelect) Scan(ctx context.Context, v interface{}) error {
	if err := hss.prepareQuery(ctx); err != nil {
		return err
	}
	hss.sql = hss.HotSearchQuery.sqlQuery(ctx)
	return hss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hss *HotSearchSelect) ScanX(ctx context.Context, v interface{}) {
	if err := hss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (hss *HotSearchSelect) Strings(ctx context.Context) ([]string, error) {
	if len(hss.fields) > 1 {
		return nil, errors.New("ent: HotSearchSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := hss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hss *HotSearchSelect) StringsX(ctx context.Context) []string {
	v, err := hss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (hss *HotSearchSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hotsearch.Label}
	default:
		err = fmt.Errorf("ent: HotSearchSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hss *HotSearchSelect) StringX(ctx context.Context) string {
	v, err := hss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (hss *HotSearchSelect) Ints(ctx context.Context) ([]int, error) {
	if len(hss.fields) > 1 {
		return nil, errors.New("ent: HotSearchSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := hss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hss *HotSearchSelect) IntsX(ctx context.Context) []int {
	v, err := hss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (hss *HotSearchSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hotsearch.Label}
	default:
		err = fmt.Errorf("ent: HotSearchSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hss *HotSearchSelect) IntX(ctx context.Context) int {
	v, err := hss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (hss *HotSearchSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(hss.fields) > 1 {
		return nil, errors.New("ent: HotSearchSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := hss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hss *HotSearchSelect) Float64sX(ctx context.Context) []float64 {
	v, err := hss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (hss *HotSearchSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hotsearch.Label}
	default:
		err = fmt.Errorf("ent: HotSearchSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hss *HotSearchSelect) Float64X(ctx context.Context) float64 {
	v, err := hss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (hss *HotSearchSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(hss.fields) > 1 {
		return nil, errors.New("ent: HotSearchSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := hss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hss *HotSearchSelect) BoolsX(ctx context.Context) []bool {
	v, err := hss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (hss *HotSearchSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hotsearch.Label}
	default:
		err = fmt.Errorf("ent: HotSearchSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hss *HotSearchSelect) BoolX(ctx context.Context) bool {
	v, err := hss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hss *HotSearchSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hss.sqlQuery().Query()
	if err := hss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hss *HotSearchSelect) sqlQuery() sql.Querier {
	selector := hss.sql
	selector.Select(selector.Columns(hss.fields...)...)
	return selector
}
