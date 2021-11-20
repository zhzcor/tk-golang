// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"tkserver/internal/store/ent/major"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/tkquestionbank"
	"tkserver/internal/store/ent/tkquestionbankmajor"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkQuestionBankMajorQuery is the builder for querying TkQuestionBankMajor entities.
type TkQuestionBankMajorQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TkQuestionBankMajor
	// eager-loading edges.
	withTkQuestionBank *TkQuestionBankQuery
	withMajor          *MajorQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TkQuestionBankMajorQuery builder.
func (tqbmq *TkQuestionBankMajorQuery) Where(ps ...predicate.TkQuestionBankMajor) *TkQuestionBankMajorQuery {
	tqbmq.predicates = append(tqbmq.predicates, ps...)
	return tqbmq
}

// Limit adds a limit step to the query.
func (tqbmq *TkQuestionBankMajorQuery) Limit(limit int) *TkQuestionBankMajorQuery {
	tqbmq.limit = &limit
	return tqbmq
}

// Offset adds an offset step to the query.
func (tqbmq *TkQuestionBankMajorQuery) Offset(offset int) *TkQuestionBankMajorQuery {
	tqbmq.offset = &offset
	return tqbmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tqbmq *TkQuestionBankMajorQuery) Unique(unique bool) *TkQuestionBankMajorQuery {
	tqbmq.unique = &unique
	return tqbmq
}

// Order adds an order step to the query.
func (tqbmq *TkQuestionBankMajorQuery) Order(o ...OrderFunc) *TkQuestionBankMajorQuery {
	tqbmq.order = append(tqbmq.order, o...)
	return tqbmq
}

// QueryTkQuestionBank chains the current query on the "tk_question_bank" edge.
func (tqbmq *TkQuestionBankMajorQuery) QueryTkQuestionBank() *TkQuestionBankQuery {
	query := &TkQuestionBankQuery{config: tqbmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tqbmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tqbmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkquestionbankmajor.Table, tkquestionbankmajor.FieldID, selector),
			sqlgraph.To(tkquestionbank.Table, tkquestionbank.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tkquestionbankmajor.TkQuestionBankTable, tkquestionbankmajor.TkQuestionBankColumn),
		)
		fromU = sqlgraph.SetNeighbors(tqbmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMajor chains the current query on the "major" edge.
func (tqbmq *TkQuestionBankMajorQuery) QueryMajor() *MajorQuery {
	query := &MajorQuery{config: tqbmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tqbmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tqbmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkquestionbankmajor.Table, tkquestionbankmajor.FieldID, selector),
			sqlgraph.To(major.Table, major.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tkquestionbankmajor.MajorTable, tkquestionbankmajor.MajorColumn),
		)
		fromU = sqlgraph.SetNeighbors(tqbmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TkQuestionBankMajor entity from the query.
// Returns a *NotFoundError when no TkQuestionBankMajor was found.
func (tqbmq *TkQuestionBankMajorQuery) First(ctx context.Context) (*TkQuestionBankMajor, error) {
	nodes, err := tqbmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tkquestionbankmajor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tqbmq *TkQuestionBankMajorQuery) FirstX(ctx context.Context) *TkQuestionBankMajor {
	node, err := tqbmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TkQuestionBankMajor ID from the query.
// Returns a *NotFoundError when no TkQuestionBankMajor ID was found.
func (tqbmq *TkQuestionBankMajorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tqbmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tkquestionbankmajor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tqbmq *TkQuestionBankMajorQuery) FirstIDX(ctx context.Context) int {
	id, err := tqbmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TkQuestionBankMajor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TkQuestionBankMajor entity is not found.
// Returns a *NotFoundError when no TkQuestionBankMajor entities are found.
func (tqbmq *TkQuestionBankMajorQuery) Only(ctx context.Context) (*TkQuestionBankMajor, error) {
	nodes, err := tqbmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tkquestionbankmajor.Label}
	default:
		return nil, &NotSingularError{tkquestionbankmajor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tqbmq *TkQuestionBankMajorQuery) OnlyX(ctx context.Context) *TkQuestionBankMajor {
	node, err := tqbmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TkQuestionBankMajor ID in the query.
// Returns a *NotSingularError when exactly one TkQuestionBankMajor ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tqbmq *TkQuestionBankMajorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tqbmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tkquestionbankmajor.Label}
	default:
		err = &NotSingularError{tkquestionbankmajor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tqbmq *TkQuestionBankMajorQuery) OnlyIDX(ctx context.Context) int {
	id, err := tqbmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TkQuestionBankMajors.
func (tqbmq *TkQuestionBankMajorQuery) All(ctx context.Context) ([]*TkQuestionBankMajor, error) {
	if err := tqbmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tqbmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tqbmq *TkQuestionBankMajorQuery) AllX(ctx context.Context) []*TkQuestionBankMajor {
	nodes, err := tqbmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TkQuestionBankMajor IDs.
func (tqbmq *TkQuestionBankMajorQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tqbmq.Select(tkquestionbankmajor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tqbmq *TkQuestionBankMajorQuery) IDsX(ctx context.Context) []int {
	ids, err := tqbmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tqbmq *TkQuestionBankMajorQuery) Count(ctx context.Context) (int, error) {
	if err := tqbmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tqbmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tqbmq *TkQuestionBankMajorQuery) CountX(ctx context.Context) int {
	count, err := tqbmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tqbmq *TkQuestionBankMajorQuery) Exist(ctx context.Context) (bool, error) {
	if err := tqbmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tqbmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tqbmq *TkQuestionBankMajorQuery) ExistX(ctx context.Context) bool {
	exist, err := tqbmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TkQuestionBankMajorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tqbmq *TkQuestionBankMajorQuery) Clone() *TkQuestionBankMajorQuery {
	if tqbmq == nil {
		return nil
	}
	return &TkQuestionBankMajorQuery{
		config:             tqbmq.config,
		limit:              tqbmq.limit,
		offset:             tqbmq.offset,
		order:              append([]OrderFunc{}, tqbmq.order...),
		predicates:         append([]predicate.TkQuestionBankMajor{}, tqbmq.predicates...),
		withTkQuestionBank: tqbmq.withTkQuestionBank.Clone(),
		withMajor:          tqbmq.withMajor.Clone(),
		// clone intermediate query.
		sql:  tqbmq.sql.Clone(),
		path: tqbmq.path,
	}
}

// WithTkQuestionBank tells the query-builder to eager-load the nodes that are connected to
// the "tk_question_bank" edge. The optional arguments are used to configure the query builder of the edge.
func (tqbmq *TkQuestionBankMajorQuery) WithTkQuestionBank(opts ...func(*TkQuestionBankQuery)) *TkQuestionBankMajorQuery {
	query := &TkQuestionBankQuery{config: tqbmq.config}
	for _, opt := range opts {
		opt(query)
	}
	tqbmq.withTkQuestionBank = query
	return tqbmq
}

// WithMajor tells the query-builder to eager-load the nodes that are connected to
// the "major" edge. The optional arguments are used to configure the query builder of the edge.
func (tqbmq *TkQuestionBankMajorQuery) WithMajor(opts ...func(*MajorQuery)) *TkQuestionBankMajorQuery {
	query := &MajorQuery{config: tqbmq.config}
	for _, opt := range opts {
		opt(query)
	}
	tqbmq.withMajor = query
	return tqbmq
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
//	client.TkQuestionBankMajor.Query().
//		GroupBy(tkquestionbankmajor.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tqbmq *TkQuestionBankMajorQuery) GroupBy(field string, fields ...string) *TkQuestionBankMajorGroupBy {
	group := &TkQuestionBankMajorGroupBy{config: tqbmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tqbmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tqbmq.sqlQuery(ctx), nil
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
//	client.TkQuestionBankMajor.Query().
//		Select(tkquestionbankmajor.FieldUUID).
//		Scan(ctx, &v)
//
func (tqbmq *TkQuestionBankMajorQuery) Select(field string, fields ...string) *TkQuestionBankMajorSelect {
	tqbmq.fields = append([]string{field}, fields...)
	return &TkQuestionBankMajorSelect{TkQuestionBankMajorQuery: tqbmq}
}

func (tqbmq *TkQuestionBankMajorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tqbmq.fields {
		if !tkquestionbankmajor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tqbmq.path != nil {
		prev, err := tqbmq.path(ctx)
		if err != nil {
			return err
		}
		tqbmq.sql = prev
	}
	return nil
}

func (tqbmq *TkQuestionBankMajorQuery) sqlAll(ctx context.Context) ([]*TkQuestionBankMajor, error) {
	var (
		nodes       = []*TkQuestionBankMajor{}
		_spec       = tqbmq.querySpec()
		loadedTypes = [2]bool{
			tqbmq.withTkQuestionBank != nil,
			tqbmq.withMajor != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TkQuestionBankMajor{config: tqbmq.config}
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
	if err := sqlgraph.QueryNodes(ctx, tqbmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tqbmq.withTkQuestionBank; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TkQuestionBankMajor)
		for i := range nodes {
			fk := nodes[i].QuestionBankID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(tkquestionbank.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "question_bank_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.TkQuestionBank = n
			}
		}
	}

	if query := tqbmq.withMajor; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TkQuestionBankMajor)
		for i := range nodes {
			fk := nodes[i].MajorID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(major.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "major_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Major = n
			}
		}
	}

	return nodes, nil
}

func (tqbmq *TkQuestionBankMajorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tqbmq.querySpec()
	return sqlgraph.CountNodes(ctx, tqbmq.driver, _spec)
}

func (tqbmq *TkQuestionBankMajorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tqbmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tqbmq *TkQuestionBankMajorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tkquestionbankmajor.Table,
			Columns: tkquestionbankmajor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkquestionbankmajor.FieldID,
			},
		},
		From:   tqbmq.sql,
		Unique: true,
	}
	if unique := tqbmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tqbmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tkquestionbankmajor.FieldID)
		for i := range fields {
			if fields[i] != tkquestionbankmajor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tqbmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tqbmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tqbmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tqbmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tqbmq *TkQuestionBankMajorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tqbmq.driver.Dialect())
	t1 := builder.Table(tkquestionbankmajor.Table)
	selector := builder.Select(t1.Columns(tkquestionbankmajor.Columns...)...).From(t1)
	if tqbmq.sql != nil {
		selector = tqbmq.sql
		selector.Select(selector.Columns(tkquestionbankmajor.Columns...)...)
	}
	for _, p := range tqbmq.predicates {
		p(selector)
	}
	for _, p := range tqbmq.order {
		p(selector)
	}
	if offset := tqbmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tqbmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TkQuestionBankMajorGroupBy is the group-by builder for TkQuestionBankMajor entities.
type TkQuestionBankMajorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tqbmgb *TkQuestionBankMajorGroupBy) Aggregate(fns ...AggregateFunc) *TkQuestionBankMajorGroupBy {
	tqbmgb.fns = append(tqbmgb.fns, fns...)
	return tqbmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tqbmgb *TkQuestionBankMajorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tqbmgb.path(ctx)
	if err != nil {
		return err
	}
	tqbmgb.sql = query
	return tqbmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tqbmgb *TkQuestionBankMajorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tqbmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tqbmgb *TkQuestionBankMajorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tqbmgb.fields) > 1 {
		return nil, errors.New("ent: TkQuestionBankMajorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tqbmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tqbmgb *TkQuestionBankMajorGroupBy) StringsX(ctx context.Context) []string {
	v, err := tqbmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tqbmgb *TkQuestionBankMajorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tqbmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionbankmajor.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionBankMajorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tqbmgb *TkQuestionBankMajorGroupBy) StringX(ctx context.Context) string {
	v, err := tqbmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tqbmgb *TkQuestionBankMajorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tqbmgb.fields) > 1 {
		return nil, errors.New("ent: TkQuestionBankMajorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tqbmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tqbmgb *TkQuestionBankMajorGroupBy) IntsX(ctx context.Context) []int {
	v, err := tqbmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tqbmgb *TkQuestionBankMajorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tqbmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionbankmajor.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionBankMajorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tqbmgb *TkQuestionBankMajorGroupBy) IntX(ctx context.Context) int {
	v, err := tqbmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tqbmgb *TkQuestionBankMajorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tqbmgb.fields) > 1 {
		return nil, errors.New("ent: TkQuestionBankMajorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tqbmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tqbmgb *TkQuestionBankMajorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tqbmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tqbmgb *TkQuestionBankMajorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tqbmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionbankmajor.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionBankMajorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tqbmgb *TkQuestionBankMajorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tqbmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tqbmgb *TkQuestionBankMajorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tqbmgb.fields) > 1 {
		return nil, errors.New("ent: TkQuestionBankMajorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tqbmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tqbmgb *TkQuestionBankMajorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tqbmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tqbmgb *TkQuestionBankMajorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tqbmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionbankmajor.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionBankMajorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tqbmgb *TkQuestionBankMajorGroupBy) BoolX(ctx context.Context) bool {
	v, err := tqbmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tqbmgb *TkQuestionBankMajorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tqbmgb.fields {
		if !tkquestionbankmajor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tqbmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tqbmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tqbmgb *TkQuestionBankMajorGroupBy) sqlQuery() *sql.Selector {
	selector := tqbmgb.sql
	columns := make([]string, 0, len(tqbmgb.fields)+len(tqbmgb.fns))
	columns = append(columns, tqbmgb.fields...)
	for _, fn := range tqbmgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tqbmgb.fields...)
}

// TkQuestionBankMajorSelect is the builder for selecting fields of TkQuestionBankMajor entities.
type TkQuestionBankMajorSelect struct {
	*TkQuestionBankMajorQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tqbms *TkQuestionBankMajorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tqbms.prepareQuery(ctx); err != nil {
		return err
	}
	tqbms.sql = tqbms.TkQuestionBankMajorQuery.sqlQuery(ctx)
	return tqbms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tqbms *TkQuestionBankMajorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tqbms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tqbms *TkQuestionBankMajorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tqbms.fields) > 1 {
		return nil, errors.New("ent: TkQuestionBankMajorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tqbms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tqbms *TkQuestionBankMajorSelect) StringsX(ctx context.Context) []string {
	v, err := tqbms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tqbms *TkQuestionBankMajorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tqbms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionbankmajor.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionBankMajorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tqbms *TkQuestionBankMajorSelect) StringX(ctx context.Context) string {
	v, err := tqbms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tqbms *TkQuestionBankMajorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tqbms.fields) > 1 {
		return nil, errors.New("ent: TkQuestionBankMajorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tqbms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tqbms *TkQuestionBankMajorSelect) IntsX(ctx context.Context) []int {
	v, err := tqbms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tqbms *TkQuestionBankMajorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tqbms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionbankmajor.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionBankMajorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tqbms *TkQuestionBankMajorSelect) IntX(ctx context.Context) int {
	v, err := tqbms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tqbms *TkQuestionBankMajorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tqbms.fields) > 1 {
		return nil, errors.New("ent: TkQuestionBankMajorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tqbms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tqbms *TkQuestionBankMajorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tqbms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tqbms *TkQuestionBankMajorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tqbms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionbankmajor.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionBankMajorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tqbms *TkQuestionBankMajorSelect) Float64X(ctx context.Context) float64 {
	v, err := tqbms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tqbms *TkQuestionBankMajorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tqbms.fields) > 1 {
		return nil, errors.New("ent: TkQuestionBankMajorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tqbms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tqbms *TkQuestionBankMajorSelect) BoolsX(ctx context.Context) []bool {
	v, err := tqbms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tqbms *TkQuestionBankMajorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tqbms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionbankmajor.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionBankMajorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tqbms *TkQuestionBankMajorSelect) BoolX(ctx context.Context) bool {
	v, err := tqbms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tqbms *TkQuestionBankMajorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tqbms.sqlQuery().Query()
	if err := tqbms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tqbms *TkQuestionBankMajorSelect) sqlQuery() sql.Querier {
	selector := tqbms.sql
	selector.Select(selector.Columns(tqbms.fields...)...)
	return selector
}
