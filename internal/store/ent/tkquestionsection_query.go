// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/tkquestion"
	"tkserver/internal/store/ent/tkquestionsection"
	"tkserver/internal/store/ent/tksection"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkQuestionSectionQuery is the builder for querying TkQuestionSection entities.
type TkQuestionSectionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TkQuestionSection
	// eager-loading edges.
	withQuestionSection *TkSectionQuery
	withSectionQuestion *TkQuestionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TkQuestionSectionQuery builder.
func (tqsq *TkQuestionSectionQuery) Where(ps ...predicate.TkQuestionSection) *TkQuestionSectionQuery {
	tqsq.predicates = append(tqsq.predicates, ps...)
	return tqsq
}

// Limit adds a limit step to the query.
func (tqsq *TkQuestionSectionQuery) Limit(limit int) *TkQuestionSectionQuery {
	tqsq.limit = &limit
	return tqsq
}

// Offset adds an offset step to the query.
func (tqsq *TkQuestionSectionQuery) Offset(offset int) *TkQuestionSectionQuery {
	tqsq.offset = &offset
	return tqsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tqsq *TkQuestionSectionQuery) Unique(unique bool) *TkQuestionSectionQuery {
	tqsq.unique = &unique
	return tqsq
}

// Order adds an order step to the query.
func (tqsq *TkQuestionSectionQuery) Order(o ...OrderFunc) *TkQuestionSectionQuery {
	tqsq.order = append(tqsq.order, o...)
	return tqsq
}

// QueryQuestionSection chains the current query on the "question_section" edge.
func (tqsq *TkQuestionSectionQuery) QueryQuestionSection() *TkSectionQuery {
	query := &TkSectionQuery{config: tqsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tqsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tqsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkquestionsection.Table, tkquestionsection.FieldID, selector),
			sqlgraph.To(tksection.Table, tksection.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tkquestionsection.QuestionSectionTable, tkquestionsection.QuestionSectionColumn),
		)
		fromU = sqlgraph.SetNeighbors(tqsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySectionQuestion chains the current query on the "section_question" edge.
func (tqsq *TkQuestionSectionQuery) QuerySectionQuestion() *TkQuestionQuery {
	query := &TkQuestionQuery{config: tqsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tqsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tqsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkquestionsection.Table, tkquestionsection.FieldID, selector),
			sqlgraph.To(tkquestion.Table, tkquestion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tkquestionsection.SectionQuestionTable, tkquestionsection.SectionQuestionColumn),
		)
		fromU = sqlgraph.SetNeighbors(tqsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TkQuestionSection entity from the query.
// Returns a *NotFoundError when no TkQuestionSection was found.
func (tqsq *TkQuestionSectionQuery) First(ctx context.Context) (*TkQuestionSection, error) {
	nodes, err := tqsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tkquestionsection.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tqsq *TkQuestionSectionQuery) FirstX(ctx context.Context) *TkQuestionSection {
	node, err := tqsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TkQuestionSection ID from the query.
// Returns a *NotFoundError when no TkQuestionSection ID was found.
func (tqsq *TkQuestionSectionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tqsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tkquestionsection.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tqsq *TkQuestionSectionQuery) FirstIDX(ctx context.Context) int {
	id, err := tqsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TkQuestionSection entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TkQuestionSection entity is not found.
// Returns a *NotFoundError when no TkQuestionSection entities are found.
func (tqsq *TkQuestionSectionQuery) Only(ctx context.Context) (*TkQuestionSection, error) {
	nodes, err := tqsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tkquestionsection.Label}
	default:
		return nil, &NotSingularError{tkquestionsection.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tqsq *TkQuestionSectionQuery) OnlyX(ctx context.Context) *TkQuestionSection {
	node, err := tqsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TkQuestionSection ID in the query.
// Returns a *NotSingularError when exactly one TkQuestionSection ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tqsq *TkQuestionSectionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tqsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tkquestionsection.Label}
	default:
		err = &NotSingularError{tkquestionsection.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tqsq *TkQuestionSectionQuery) OnlyIDX(ctx context.Context) int {
	id, err := tqsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TkQuestionSections.
func (tqsq *TkQuestionSectionQuery) All(ctx context.Context) ([]*TkQuestionSection, error) {
	if err := tqsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tqsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tqsq *TkQuestionSectionQuery) AllX(ctx context.Context) []*TkQuestionSection {
	nodes, err := tqsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TkQuestionSection IDs.
func (tqsq *TkQuestionSectionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tqsq.Select(tkquestionsection.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tqsq *TkQuestionSectionQuery) IDsX(ctx context.Context) []int {
	ids, err := tqsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tqsq *TkQuestionSectionQuery) Count(ctx context.Context) (int, error) {
	if err := tqsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tqsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tqsq *TkQuestionSectionQuery) CountX(ctx context.Context) int {
	count, err := tqsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tqsq *TkQuestionSectionQuery) Exist(ctx context.Context) (bool, error) {
	if err := tqsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tqsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tqsq *TkQuestionSectionQuery) ExistX(ctx context.Context) bool {
	exist, err := tqsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TkQuestionSectionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tqsq *TkQuestionSectionQuery) Clone() *TkQuestionSectionQuery {
	if tqsq == nil {
		return nil
	}
	return &TkQuestionSectionQuery{
		config:              tqsq.config,
		limit:               tqsq.limit,
		offset:              tqsq.offset,
		order:               append([]OrderFunc{}, tqsq.order...),
		predicates:          append([]predicate.TkQuestionSection{}, tqsq.predicates...),
		withQuestionSection: tqsq.withQuestionSection.Clone(),
		withSectionQuestion: tqsq.withSectionQuestion.Clone(),
		// clone intermediate query.
		sql:  tqsq.sql.Clone(),
		path: tqsq.path,
	}
}

// WithQuestionSection tells the query-builder to eager-load the nodes that are connected to
// the "question_section" edge. The optional arguments are used to configure the query builder of the edge.
func (tqsq *TkQuestionSectionQuery) WithQuestionSection(opts ...func(*TkSectionQuery)) *TkQuestionSectionQuery {
	query := &TkSectionQuery{config: tqsq.config}
	for _, opt := range opts {
		opt(query)
	}
	tqsq.withQuestionSection = query
	return tqsq
}

// WithSectionQuestion tells the query-builder to eager-load the nodes that are connected to
// the "section_question" edge. The optional arguments are used to configure the query builder of the edge.
func (tqsq *TkQuestionSectionQuery) WithSectionQuestion(opts ...func(*TkQuestionQuery)) *TkQuestionSectionQuery {
	query := &TkQuestionQuery{config: tqsq.config}
	for _, opt := range opts {
		opt(query)
	}
	tqsq.withSectionQuestion = query
	return tqsq
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
//	client.TkQuestionSection.Query().
//		GroupBy(tkquestionsection.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tqsq *TkQuestionSectionQuery) GroupBy(field string, fields ...string) *TkQuestionSectionGroupBy {
	group := &TkQuestionSectionGroupBy{config: tqsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tqsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tqsq.sqlQuery(ctx), nil
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
//	client.TkQuestionSection.Query().
//		Select(tkquestionsection.FieldUUID).
//		Scan(ctx, &v)
//
func (tqsq *TkQuestionSectionQuery) Select(field string, fields ...string) *TkQuestionSectionSelect {
	tqsq.fields = append([]string{field}, fields...)
	return &TkQuestionSectionSelect{TkQuestionSectionQuery: tqsq}
}

func (tqsq *TkQuestionSectionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tqsq.fields {
		if !tkquestionsection.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tqsq.path != nil {
		prev, err := tqsq.path(ctx)
		if err != nil {
			return err
		}
		tqsq.sql = prev
	}
	return nil
}

func (tqsq *TkQuestionSectionQuery) sqlAll(ctx context.Context) ([]*TkQuestionSection, error) {
	var (
		nodes       = []*TkQuestionSection{}
		_spec       = tqsq.querySpec()
		loadedTypes = [2]bool{
			tqsq.withQuestionSection != nil,
			tqsq.withSectionQuestion != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TkQuestionSection{config: tqsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, tqsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tqsq.withQuestionSection; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TkQuestionSection)
		for i := range nodes {
			fk := nodes[i].SectionID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(tksection.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "section_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.QuestionSection = n
			}
		}
	}

	if query := tqsq.withSectionQuestion; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TkQuestionSection)
		for i := range nodes {
			fk := nodes[i].QuestionID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(tkquestion.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "question_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.SectionQuestion = n
			}
		}
	}

	return nodes, nil
}

func (tqsq *TkQuestionSectionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tqsq.querySpec()
	return sqlgraph.CountNodes(ctx, tqsq.driver, _spec)
}

func (tqsq *TkQuestionSectionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tqsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tqsq *TkQuestionSectionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tkquestionsection.Table,
			Columns: tkquestionsection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkquestionsection.FieldID,
			},
		},
		From:   tqsq.sql,
		Unique: true,
	}
	if unique := tqsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tqsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tkquestionsection.FieldID)
		for i := range fields {
			if fields[i] != tkquestionsection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tqsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tqsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tqsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tqsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tqsq *TkQuestionSectionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tqsq.driver.Dialect())
	t1 := builder.Table(tkquestionsection.Table)
	selector := builder.Select(t1.Columns(tkquestionsection.Columns...)...).From(t1)
	if tqsq.sql != nil {
		selector = tqsq.sql
		selector.Select(selector.Columns(tkquestionsection.Columns...)...)
	}
	for _, p := range tqsq.predicates {
		p(selector)
	}
	for _, p := range tqsq.order {
		p(selector)
	}
	if offset := tqsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tqsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TkQuestionSectionGroupBy is the group-by builder for TkQuestionSection entities.
type TkQuestionSectionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tqsgb *TkQuestionSectionGroupBy) Aggregate(fns ...AggregateFunc) *TkQuestionSectionGroupBy {
	tqsgb.fns = append(tqsgb.fns, fns...)
	return tqsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tqsgb *TkQuestionSectionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tqsgb.path(ctx)
	if err != nil {
		return err
	}
	tqsgb.sql = query
	return tqsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tqsgb *TkQuestionSectionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tqsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tqsgb *TkQuestionSectionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tqsgb.fields) > 1 {
		return nil, errors.New("ent: TkQuestionSectionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tqsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tqsgb *TkQuestionSectionGroupBy) StringsX(ctx context.Context) []string {
	v, err := tqsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tqsgb *TkQuestionSectionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tqsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionsection.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionSectionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tqsgb *TkQuestionSectionGroupBy) StringX(ctx context.Context) string {
	v, err := tqsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tqsgb *TkQuestionSectionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tqsgb.fields) > 1 {
		return nil, errors.New("ent: TkQuestionSectionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tqsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tqsgb *TkQuestionSectionGroupBy) IntsX(ctx context.Context) []int {
	v, err := tqsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tqsgb *TkQuestionSectionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tqsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionsection.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionSectionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tqsgb *TkQuestionSectionGroupBy) IntX(ctx context.Context) int {
	v, err := tqsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tqsgb *TkQuestionSectionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tqsgb.fields) > 1 {
		return nil, errors.New("ent: TkQuestionSectionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tqsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tqsgb *TkQuestionSectionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tqsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tqsgb *TkQuestionSectionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tqsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionsection.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionSectionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tqsgb *TkQuestionSectionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tqsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tqsgb *TkQuestionSectionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tqsgb.fields) > 1 {
		return nil, errors.New("ent: TkQuestionSectionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tqsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tqsgb *TkQuestionSectionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tqsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tqsgb *TkQuestionSectionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tqsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionsection.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionSectionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tqsgb *TkQuestionSectionGroupBy) BoolX(ctx context.Context) bool {
	v, err := tqsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tqsgb *TkQuestionSectionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tqsgb.fields {
		if !tkquestionsection.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tqsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tqsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tqsgb *TkQuestionSectionGroupBy) sqlQuery() *sql.Selector {
	selector := tqsgb.sql
	columns := make([]string, 0, len(tqsgb.fields)+len(tqsgb.fns))
	columns = append(columns, tqsgb.fields...)
	for _, fn := range tqsgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tqsgb.fields...)
}

// TkQuestionSectionSelect is the builder for selecting fields of TkQuestionSection entities.
type TkQuestionSectionSelect struct {
	*TkQuestionSectionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tqss *TkQuestionSectionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tqss.prepareQuery(ctx); err != nil {
		return err
	}
	tqss.sql = tqss.TkQuestionSectionQuery.sqlQuery(ctx)
	return tqss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tqss *TkQuestionSectionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tqss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tqss *TkQuestionSectionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tqss.fields) > 1 {
		return nil, errors.New("ent: TkQuestionSectionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tqss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tqss *TkQuestionSectionSelect) StringsX(ctx context.Context) []string {
	v, err := tqss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tqss *TkQuestionSectionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tqss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionsection.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionSectionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tqss *TkQuestionSectionSelect) StringX(ctx context.Context) string {
	v, err := tqss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tqss *TkQuestionSectionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tqss.fields) > 1 {
		return nil, errors.New("ent: TkQuestionSectionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tqss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tqss *TkQuestionSectionSelect) IntsX(ctx context.Context) []int {
	v, err := tqss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tqss *TkQuestionSectionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tqss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionsection.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionSectionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tqss *TkQuestionSectionSelect) IntX(ctx context.Context) int {
	v, err := tqss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tqss *TkQuestionSectionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tqss.fields) > 1 {
		return nil, errors.New("ent: TkQuestionSectionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tqss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tqss *TkQuestionSectionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tqss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tqss *TkQuestionSectionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tqss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionsection.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionSectionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tqss *TkQuestionSectionSelect) Float64X(ctx context.Context) float64 {
	v, err := tqss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tqss *TkQuestionSectionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tqss.fields) > 1 {
		return nil, errors.New("ent: TkQuestionSectionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tqss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tqss *TkQuestionSectionSelect) BoolsX(ctx context.Context) []bool {
	v, err := tqss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tqss *TkQuestionSectionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tqss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkquestionsection.Label}
	default:
		err = fmt.Errorf("ent: TkQuestionSectionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tqss *TkQuestionSectionSelect) BoolX(ctx context.Context) bool {
	v, err := tqss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tqss *TkQuestionSectionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tqss.sqlQuery().Query()
	if err := tqss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tqss *TkQuestionSectionSelect) sqlQuery() sql.Querier {
	selector := tqss.sql
	selector.Select(selector.Columns(tqss.fields...)...)
	return selector
}
