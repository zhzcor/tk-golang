// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/tkquestionbank"
	"gserver/internal/store/ent/tkuserquestionbankrecord"
	"gserver/internal/store/ent/user"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkUserQuestionBankRecordQuery is the builder for querying TkUserQuestionBankRecord entities.
type TkUserQuestionBankRecordQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TkUserQuestionBankRecord
	// eager-loading edges.
	withQuestionBank *TkQuestionBankQuery
	withUser         *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TkUserQuestionBankRecordQuery builder.
func (tuqbrq *TkUserQuestionBankRecordQuery) Where(ps ...predicate.TkUserQuestionBankRecord) *TkUserQuestionBankRecordQuery {
	tuqbrq.predicates = append(tuqbrq.predicates, ps...)
	return tuqbrq
}

// Limit adds a limit step to the query.
func (tuqbrq *TkUserQuestionBankRecordQuery) Limit(limit int) *TkUserQuestionBankRecordQuery {
	tuqbrq.limit = &limit
	return tuqbrq
}

// Offset adds an offset step to the query.
func (tuqbrq *TkUserQuestionBankRecordQuery) Offset(offset int) *TkUserQuestionBankRecordQuery {
	tuqbrq.offset = &offset
	return tuqbrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tuqbrq *TkUserQuestionBankRecordQuery) Unique(unique bool) *TkUserQuestionBankRecordQuery {
	tuqbrq.unique = &unique
	return tuqbrq
}

// Order adds an order step to the query.
func (tuqbrq *TkUserQuestionBankRecordQuery) Order(o ...OrderFunc) *TkUserQuestionBankRecordQuery {
	tuqbrq.order = append(tuqbrq.order, o...)
	return tuqbrq
}

// QueryQuestionBank chains the current query on the "question_bank" edge.
func (tuqbrq *TkUserQuestionBankRecordQuery) QueryQuestionBank() *TkQuestionBankQuery {
	query := &TkQuestionBankQuery{config: tuqbrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tuqbrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tuqbrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkuserquestionbankrecord.Table, tkuserquestionbankrecord.FieldID, selector),
			sqlgraph.To(tkquestionbank.Table, tkquestionbank.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tkuserquestionbankrecord.QuestionBankTable, tkuserquestionbankrecord.QuestionBankColumn),
		)
		fromU = sqlgraph.SetNeighbors(tuqbrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (tuqbrq *TkUserQuestionBankRecordQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: tuqbrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tuqbrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tuqbrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkuserquestionbankrecord.Table, tkuserquestionbankrecord.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tkuserquestionbankrecord.UserTable, tkuserquestionbankrecord.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tuqbrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TkUserQuestionBankRecord entity from the query.
// Returns a *NotFoundError when no TkUserQuestionBankRecord was found.
func (tuqbrq *TkUserQuestionBankRecordQuery) First(ctx context.Context) (*TkUserQuestionBankRecord, error) {
	nodes, err := tuqbrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tkuserquestionbankrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tuqbrq *TkUserQuestionBankRecordQuery) FirstX(ctx context.Context) *TkUserQuestionBankRecord {
	node, err := tuqbrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TkUserQuestionBankRecord ID from the query.
// Returns a *NotFoundError when no TkUserQuestionBankRecord ID was found.
func (tuqbrq *TkUserQuestionBankRecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tuqbrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tkuserquestionbankrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tuqbrq *TkUserQuestionBankRecordQuery) FirstIDX(ctx context.Context) int {
	id, err := tuqbrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TkUserQuestionBankRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TkUserQuestionBankRecord entity is not found.
// Returns a *NotFoundError when no TkUserQuestionBankRecord entities are found.
func (tuqbrq *TkUserQuestionBankRecordQuery) Only(ctx context.Context) (*TkUserQuestionBankRecord, error) {
	nodes, err := tuqbrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		return nil, &NotSingularError{tkuserquestionbankrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tuqbrq *TkUserQuestionBankRecordQuery) OnlyX(ctx context.Context) *TkUserQuestionBankRecord {
	node, err := tuqbrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TkUserQuestionBankRecord ID in the query.
// Returns a *NotSingularError when exactly one TkUserQuestionBankRecord ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tuqbrq *TkUserQuestionBankRecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tuqbrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		err = &NotSingularError{tkuserquestionbankrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tuqbrq *TkUserQuestionBankRecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := tuqbrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TkUserQuestionBankRecords.
func (tuqbrq *TkUserQuestionBankRecordQuery) All(ctx context.Context) ([]*TkUserQuestionBankRecord, error) {
	if err := tuqbrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tuqbrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tuqbrq *TkUserQuestionBankRecordQuery) AllX(ctx context.Context) []*TkUserQuestionBankRecord {
	nodes, err := tuqbrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TkUserQuestionBankRecord IDs.
func (tuqbrq *TkUserQuestionBankRecordQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tuqbrq.Select(tkuserquestionbankrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tuqbrq *TkUserQuestionBankRecordQuery) IDsX(ctx context.Context) []int {
	ids, err := tuqbrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tuqbrq *TkUserQuestionBankRecordQuery) Count(ctx context.Context) (int, error) {
	if err := tuqbrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tuqbrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tuqbrq *TkUserQuestionBankRecordQuery) CountX(ctx context.Context) int {
	count, err := tuqbrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tuqbrq *TkUserQuestionBankRecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := tuqbrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tuqbrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tuqbrq *TkUserQuestionBankRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := tuqbrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TkUserQuestionBankRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tuqbrq *TkUserQuestionBankRecordQuery) Clone() *TkUserQuestionBankRecordQuery {
	if tuqbrq == nil {
		return nil
	}
	return &TkUserQuestionBankRecordQuery{
		config:           tuqbrq.config,
		limit:            tuqbrq.limit,
		offset:           tuqbrq.offset,
		order:            append([]OrderFunc{}, tuqbrq.order...),
		predicates:       append([]predicate.TkUserQuestionBankRecord{}, tuqbrq.predicates...),
		withQuestionBank: tuqbrq.withQuestionBank.Clone(),
		withUser:         tuqbrq.withUser.Clone(),
		// clone intermediate query.
		sql:  tuqbrq.sql.Clone(),
		path: tuqbrq.path,
	}
}

// WithQuestionBank tells the query-builder to eager-load the nodes that are connected to
// the "question_bank" edge. The optional arguments are used to configure the query builder of the edge.
func (tuqbrq *TkUserQuestionBankRecordQuery) WithQuestionBank(opts ...func(*TkQuestionBankQuery)) *TkUserQuestionBankRecordQuery {
	query := &TkQuestionBankQuery{config: tuqbrq.config}
	for _, opt := range opts {
		opt(query)
	}
	tuqbrq.withQuestionBank = query
	return tuqbrq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tuqbrq *TkUserQuestionBankRecordQuery) WithUser(opts ...func(*UserQuery)) *TkUserQuestionBankRecordQuery {
	query := &UserQuery{config: tuqbrq.config}
	for _, opt := range opts {
		opt(query)
	}
	tuqbrq.withUser = query
	return tuqbrq
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
//	client.TkUserQuestionBankRecord.Query().
//		GroupBy(tkuserquestionbankrecord.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tuqbrq *TkUserQuestionBankRecordQuery) GroupBy(field string, fields ...string) *TkUserQuestionBankRecordGroupBy {
	group := &TkUserQuestionBankRecordGroupBy{config: tuqbrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tuqbrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tuqbrq.sqlQuery(ctx), nil
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
//	client.TkUserQuestionBankRecord.Query().
//		Select(tkuserquestionbankrecord.FieldUUID).
//		Scan(ctx, &v)
//
func (tuqbrq *TkUserQuestionBankRecordQuery) Select(field string, fields ...string) *TkUserQuestionBankRecordSelect {
	tuqbrq.fields = append([]string{field}, fields...)
	return &TkUserQuestionBankRecordSelect{TkUserQuestionBankRecordQuery: tuqbrq}
}

func (tuqbrq *TkUserQuestionBankRecordQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tuqbrq.fields {
		if !tkuserquestionbankrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tuqbrq.path != nil {
		prev, err := tuqbrq.path(ctx)
		if err != nil {
			return err
		}
		tuqbrq.sql = prev
	}
	return nil
}

func (tuqbrq *TkUserQuestionBankRecordQuery) sqlAll(ctx context.Context) ([]*TkUserQuestionBankRecord, error) {
	var (
		nodes       = []*TkUserQuestionBankRecord{}
		_spec       = tuqbrq.querySpec()
		loadedTypes = [2]bool{
			tuqbrq.withQuestionBank != nil,
			tuqbrq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TkUserQuestionBankRecord{config: tuqbrq.config}
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
	if err := sqlgraph.QueryNodes(ctx, tuqbrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tuqbrq.withQuestionBank; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TkUserQuestionBankRecord)
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
				nodes[i].Edges.QuestionBank = n
			}
		}
	}

	if query := tuqbrq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TkUserQuestionBankRecord)
		for i := range nodes {
			fk := nodes[i].UserID
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (tuqbrq *TkUserQuestionBankRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tuqbrq.querySpec()
	return sqlgraph.CountNodes(ctx, tuqbrq.driver, _spec)
}

func (tuqbrq *TkUserQuestionBankRecordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tuqbrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tuqbrq *TkUserQuestionBankRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tkuserquestionbankrecord.Table,
			Columns: tkuserquestionbankrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkuserquestionbankrecord.FieldID,
			},
		},
		From:   tuqbrq.sql,
		Unique: true,
	}
	if unique := tuqbrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tuqbrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tkuserquestionbankrecord.FieldID)
		for i := range fields {
			if fields[i] != tkuserquestionbankrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tuqbrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tuqbrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tuqbrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tuqbrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tuqbrq *TkUserQuestionBankRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tuqbrq.driver.Dialect())
	t1 := builder.Table(tkuserquestionbankrecord.Table)
	selector := builder.Select(t1.Columns(tkuserquestionbankrecord.Columns...)...).From(t1)
	if tuqbrq.sql != nil {
		selector = tuqbrq.sql
		selector.Select(selector.Columns(tkuserquestionbankrecord.Columns...)...)
	}
	for _, p := range tuqbrq.predicates {
		p(selector)
	}
	for _, p := range tuqbrq.order {
		p(selector)
	}
	if offset := tuqbrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tuqbrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TkUserQuestionBankRecordGroupBy is the group-by builder for TkUserQuestionBankRecord entities.
type TkUserQuestionBankRecordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Aggregate(fns ...AggregateFunc) *TkUserQuestionBankRecordGroupBy {
	tuqbrgb.fns = append(tuqbrgb.fns, fns...)
	return tuqbrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tuqbrgb.path(ctx)
	if err != nil {
		return err
	}
	tuqbrgb.sql = query
	return tuqbrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tuqbrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tuqbrgb.fields) > 1 {
		return nil, errors.New("ent: TkUserQuestionBankRecordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tuqbrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) StringsX(ctx context.Context) []string {
	v, err := tuqbrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tuqbrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		err = fmt.Errorf("ent: TkUserQuestionBankRecordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) StringX(ctx context.Context) string {
	v, err := tuqbrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tuqbrgb.fields) > 1 {
		return nil, errors.New("ent: TkUserQuestionBankRecordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tuqbrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) IntsX(ctx context.Context) []int {
	v, err := tuqbrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tuqbrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		err = fmt.Errorf("ent: TkUserQuestionBankRecordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) IntX(ctx context.Context) int {
	v, err := tuqbrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tuqbrgb.fields) > 1 {
		return nil, errors.New("ent: TkUserQuestionBankRecordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tuqbrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tuqbrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tuqbrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		err = fmt.Errorf("ent: TkUserQuestionBankRecordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tuqbrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tuqbrgb.fields) > 1 {
		return nil, errors.New("ent: TkUserQuestionBankRecordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tuqbrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tuqbrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tuqbrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		err = fmt.Errorf("ent: TkUserQuestionBankRecordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tuqbrgb *TkUserQuestionBankRecordGroupBy) BoolX(ctx context.Context) bool {
	v, err := tuqbrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tuqbrgb *TkUserQuestionBankRecordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tuqbrgb.fields {
		if !tkuserquestionbankrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tuqbrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tuqbrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tuqbrgb *TkUserQuestionBankRecordGroupBy) sqlQuery() *sql.Selector {
	selector := tuqbrgb.sql
	columns := make([]string, 0, len(tuqbrgb.fields)+len(tuqbrgb.fns))
	columns = append(columns, tuqbrgb.fields...)
	for _, fn := range tuqbrgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tuqbrgb.fields...)
}

// TkUserQuestionBankRecordSelect is the builder for selecting fields of TkUserQuestionBankRecord entities.
type TkUserQuestionBankRecordSelect struct {
	*TkUserQuestionBankRecordQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tuqbrs *TkUserQuestionBankRecordSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tuqbrs.prepareQuery(ctx); err != nil {
		return err
	}
	tuqbrs.sql = tuqbrs.TkUserQuestionBankRecordQuery.sqlQuery(ctx)
	return tuqbrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tuqbrs *TkUserQuestionBankRecordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tuqbrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tuqbrs *TkUserQuestionBankRecordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tuqbrs.fields) > 1 {
		return nil, errors.New("ent: TkUserQuestionBankRecordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tuqbrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tuqbrs *TkUserQuestionBankRecordSelect) StringsX(ctx context.Context) []string {
	v, err := tuqbrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tuqbrs *TkUserQuestionBankRecordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tuqbrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		err = fmt.Errorf("ent: TkUserQuestionBankRecordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tuqbrs *TkUserQuestionBankRecordSelect) StringX(ctx context.Context) string {
	v, err := tuqbrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tuqbrs *TkUserQuestionBankRecordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tuqbrs.fields) > 1 {
		return nil, errors.New("ent: TkUserQuestionBankRecordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tuqbrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tuqbrs *TkUserQuestionBankRecordSelect) IntsX(ctx context.Context) []int {
	v, err := tuqbrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tuqbrs *TkUserQuestionBankRecordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tuqbrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		err = fmt.Errorf("ent: TkUserQuestionBankRecordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tuqbrs *TkUserQuestionBankRecordSelect) IntX(ctx context.Context) int {
	v, err := tuqbrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tuqbrs *TkUserQuestionBankRecordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tuqbrs.fields) > 1 {
		return nil, errors.New("ent: TkUserQuestionBankRecordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tuqbrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tuqbrs *TkUserQuestionBankRecordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tuqbrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tuqbrs *TkUserQuestionBankRecordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tuqbrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		err = fmt.Errorf("ent: TkUserQuestionBankRecordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tuqbrs *TkUserQuestionBankRecordSelect) Float64X(ctx context.Context) float64 {
	v, err := tuqbrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tuqbrs *TkUserQuestionBankRecordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tuqbrs.fields) > 1 {
		return nil, errors.New("ent: TkUserQuestionBankRecordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tuqbrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tuqbrs *TkUserQuestionBankRecordSelect) BoolsX(ctx context.Context) []bool {
	v, err := tuqbrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tuqbrs *TkUserQuestionBankRecordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tuqbrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		err = fmt.Errorf("ent: TkUserQuestionBankRecordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tuqbrs *TkUserQuestionBankRecordSelect) BoolX(ctx context.Context) bool {
	v, err := tuqbrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tuqbrs *TkUserQuestionBankRecordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tuqbrs.sqlQuery().Query()
	if err := tuqbrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tuqbrs *TkUserQuestionBankRecordSelect) sqlQuery() sql.Querier {
	selector := tuqbrs.sql
	selector.Select(selector.Columns(tuqbrs.fields...)...)
	return selector
}
