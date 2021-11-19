// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/tkexampaperpartition"
	"gserver/internal/store/ent/tkexampartitionquestionlink"
	"gserver/internal/store/ent/tkquestion"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkExamPartitionQuestionLinkQuery is the builder for querying TkExamPartitionQuestionLink entities.
type TkExamPartitionQuestionLinkQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TkExamPartitionQuestionLink
	// eager-loading edges.
	withExamPaperPartition *TkExamPaperPartitionQuery
	withQuestion           *TkQuestionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TkExamPartitionQuestionLinkQuery builder.
func (tepqlq *TkExamPartitionQuestionLinkQuery) Where(ps ...predicate.TkExamPartitionQuestionLink) *TkExamPartitionQuestionLinkQuery {
	tepqlq.predicates = append(tepqlq.predicates, ps...)
	return tepqlq
}

// Limit adds a limit step to the query.
func (tepqlq *TkExamPartitionQuestionLinkQuery) Limit(limit int) *TkExamPartitionQuestionLinkQuery {
	tepqlq.limit = &limit
	return tepqlq
}

// Offset adds an offset step to the query.
func (tepqlq *TkExamPartitionQuestionLinkQuery) Offset(offset int) *TkExamPartitionQuestionLinkQuery {
	tepqlq.offset = &offset
	return tepqlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tepqlq *TkExamPartitionQuestionLinkQuery) Unique(unique bool) *TkExamPartitionQuestionLinkQuery {
	tepqlq.unique = &unique
	return tepqlq
}

// Order adds an order step to the query.
func (tepqlq *TkExamPartitionQuestionLinkQuery) Order(o ...OrderFunc) *TkExamPartitionQuestionLinkQuery {
	tepqlq.order = append(tepqlq.order, o...)
	return tepqlq
}

// QueryExamPaperPartition chains the current query on the "exam_paper_partition" edge.
func (tepqlq *TkExamPartitionQuestionLinkQuery) QueryExamPaperPartition() *TkExamPaperPartitionQuery {
	query := &TkExamPaperPartitionQuery{config: tepqlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tepqlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tepqlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkexampartitionquestionlink.Table, tkexampartitionquestionlink.FieldID, selector),
			sqlgraph.To(tkexampaperpartition.Table, tkexampaperpartition.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tkexampartitionquestionlink.ExamPaperPartitionTable, tkexampartitionquestionlink.ExamPaperPartitionColumn),
		)
		fromU = sqlgraph.SetNeighbors(tepqlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryQuestion chains the current query on the "question" edge.
func (tepqlq *TkExamPartitionQuestionLinkQuery) QueryQuestion() *TkQuestionQuery {
	query := &TkQuestionQuery{config: tepqlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tepqlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tepqlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkexampartitionquestionlink.Table, tkexampartitionquestionlink.FieldID, selector),
			sqlgraph.To(tkquestion.Table, tkquestion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tkexampartitionquestionlink.QuestionTable, tkexampartitionquestionlink.QuestionColumn),
		)
		fromU = sqlgraph.SetNeighbors(tepqlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TkExamPartitionQuestionLink entity from the query.
// Returns a *NotFoundError when no TkExamPartitionQuestionLink was found.
func (tepqlq *TkExamPartitionQuestionLinkQuery) First(ctx context.Context) (*TkExamPartitionQuestionLink, error) {
	nodes, err := tepqlq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tkexampartitionquestionlink.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tepqlq *TkExamPartitionQuestionLinkQuery) FirstX(ctx context.Context) *TkExamPartitionQuestionLink {
	node, err := tepqlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TkExamPartitionQuestionLink ID from the query.
// Returns a *NotFoundError when no TkExamPartitionQuestionLink ID was found.
func (tepqlq *TkExamPartitionQuestionLinkQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tepqlq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tkexampartitionquestionlink.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tepqlq *TkExamPartitionQuestionLinkQuery) FirstIDX(ctx context.Context) int {
	id, err := tepqlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TkExamPartitionQuestionLink entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TkExamPartitionQuestionLink entity is not found.
// Returns a *NotFoundError when no TkExamPartitionQuestionLink entities are found.
func (tepqlq *TkExamPartitionQuestionLinkQuery) Only(ctx context.Context) (*TkExamPartitionQuestionLink, error) {
	nodes, err := tepqlq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		return nil, &NotSingularError{tkexampartitionquestionlink.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tepqlq *TkExamPartitionQuestionLinkQuery) OnlyX(ctx context.Context) *TkExamPartitionQuestionLink {
	node, err := tepqlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TkExamPartitionQuestionLink ID in the query.
// Returns a *NotSingularError when exactly one TkExamPartitionQuestionLink ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tepqlq *TkExamPartitionQuestionLinkQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tepqlq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		err = &NotSingularError{tkexampartitionquestionlink.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tepqlq *TkExamPartitionQuestionLinkQuery) OnlyIDX(ctx context.Context) int {
	id, err := tepqlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TkExamPartitionQuestionLinks.
func (tepqlq *TkExamPartitionQuestionLinkQuery) All(ctx context.Context) ([]*TkExamPartitionQuestionLink, error) {
	if err := tepqlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tepqlq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tepqlq *TkExamPartitionQuestionLinkQuery) AllX(ctx context.Context) []*TkExamPartitionQuestionLink {
	nodes, err := tepqlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TkExamPartitionQuestionLink IDs.
func (tepqlq *TkExamPartitionQuestionLinkQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tepqlq.Select(tkexampartitionquestionlink.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tepqlq *TkExamPartitionQuestionLinkQuery) IDsX(ctx context.Context) []int {
	ids, err := tepqlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tepqlq *TkExamPartitionQuestionLinkQuery) Count(ctx context.Context) (int, error) {
	if err := tepqlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tepqlq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tepqlq *TkExamPartitionQuestionLinkQuery) CountX(ctx context.Context) int {
	count, err := tepqlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tepqlq *TkExamPartitionQuestionLinkQuery) Exist(ctx context.Context) (bool, error) {
	if err := tepqlq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tepqlq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tepqlq *TkExamPartitionQuestionLinkQuery) ExistX(ctx context.Context) bool {
	exist, err := tepqlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TkExamPartitionQuestionLinkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tepqlq *TkExamPartitionQuestionLinkQuery) Clone() *TkExamPartitionQuestionLinkQuery {
	if tepqlq == nil {
		return nil
	}
	return &TkExamPartitionQuestionLinkQuery{
		config:                 tepqlq.config,
		limit:                  tepqlq.limit,
		offset:                 tepqlq.offset,
		order:                  append([]OrderFunc{}, tepqlq.order...),
		predicates:             append([]predicate.TkExamPartitionQuestionLink{}, tepqlq.predicates...),
		withExamPaperPartition: tepqlq.withExamPaperPartition.Clone(),
		withQuestion:           tepqlq.withQuestion.Clone(),
		// clone intermediate query.
		sql:  tepqlq.sql.Clone(),
		path: tepqlq.path,
	}
}

// WithExamPaperPartition tells the query-builder to eager-load the nodes that are connected to
// the "exam_paper_partition" edge. The optional arguments are used to configure the query builder of the edge.
func (tepqlq *TkExamPartitionQuestionLinkQuery) WithExamPaperPartition(opts ...func(*TkExamPaperPartitionQuery)) *TkExamPartitionQuestionLinkQuery {
	query := &TkExamPaperPartitionQuery{config: tepqlq.config}
	for _, opt := range opts {
		opt(query)
	}
	tepqlq.withExamPaperPartition = query
	return tepqlq
}

// WithQuestion tells the query-builder to eager-load the nodes that are connected to
// the "question" edge. The optional arguments are used to configure the query builder of the edge.
func (tepqlq *TkExamPartitionQuestionLinkQuery) WithQuestion(opts ...func(*TkQuestionQuery)) *TkExamPartitionQuestionLinkQuery {
	query := &TkQuestionQuery{config: tepqlq.config}
	for _, opt := range opts {
		opt(query)
	}
	tepqlq.withQuestion = query
	return tepqlq
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
//	client.TkExamPartitionQuestionLink.Query().
//		GroupBy(tkexampartitionquestionlink.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tepqlq *TkExamPartitionQuestionLinkQuery) GroupBy(field string, fields ...string) *TkExamPartitionQuestionLinkGroupBy {
	group := &TkExamPartitionQuestionLinkGroupBy{config: tepqlq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tepqlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tepqlq.sqlQuery(ctx), nil
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
//	client.TkExamPartitionQuestionLink.Query().
//		Select(tkexampartitionquestionlink.FieldUUID).
//		Scan(ctx, &v)
//
func (tepqlq *TkExamPartitionQuestionLinkQuery) Select(field string, fields ...string) *TkExamPartitionQuestionLinkSelect {
	tepqlq.fields = append([]string{field}, fields...)
	return &TkExamPartitionQuestionLinkSelect{TkExamPartitionQuestionLinkQuery: tepqlq}
}

func (tepqlq *TkExamPartitionQuestionLinkQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tepqlq.fields {
		if !tkexampartitionquestionlink.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tepqlq.path != nil {
		prev, err := tepqlq.path(ctx)
		if err != nil {
			return err
		}
		tepqlq.sql = prev
	}
	return nil
}

func (tepqlq *TkExamPartitionQuestionLinkQuery) sqlAll(ctx context.Context) ([]*TkExamPartitionQuestionLink, error) {
	var (
		nodes       = []*TkExamPartitionQuestionLink{}
		_spec       = tepqlq.querySpec()
		loadedTypes = [2]bool{
			tepqlq.withExamPaperPartition != nil,
			tepqlq.withQuestion != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TkExamPartitionQuestionLink{config: tepqlq.config}
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
	if err := sqlgraph.QueryNodes(ctx, tepqlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tepqlq.withExamPaperPartition; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TkExamPartitionQuestionLink)
		for i := range nodes {
			fk := nodes[i].ExamPaperPartitionID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(tkexampaperpartition.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "exam_paper_partition_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ExamPaperPartition = n
			}
		}
	}

	if query := tepqlq.withQuestion; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TkExamPartitionQuestionLink)
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
				nodes[i].Edges.Question = n
			}
		}
	}

	return nodes, nil
}

func (tepqlq *TkExamPartitionQuestionLinkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tepqlq.querySpec()
	return sqlgraph.CountNodes(ctx, tepqlq.driver, _spec)
}

func (tepqlq *TkExamPartitionQuestionLinkQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tepqlq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tepqlq *TkExamPartitionQuestionLinkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tkexampartitionquestionlink.Table,
			Columns: tkexampartitionquestionlink.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkexampartitionquestionlink.FieldID,
			},
		},
		From:   tepqlq.sql,
		Unique: true,
	}
	if unique := tepqlq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tepqlq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tkexampartitionquestionlink.FieldID)
		for i := range fields {
			if fields[i] != tkexampartitionquestionlink.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tepqlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tepqlq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tepqlq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tepqlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tepqlq *TkExamPartitionQuestionLinkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tepqlq.driver.Dialect())
	t1 := builder.Table(tkexampartitionquestionlink.Table)
	selector := builder.Select(t1.Columns(tkexampartitionquestionlink.Columns...)...).From(t1)
	if tepqlq.sql != nil {
		selector = tepqlq.sql
		selector.Select(selector.Columns(tkexampartitionquestionlink.Columns...)...)
	}
	for _, p := range tepqlq.predicates {
		p(selector)
	}
	for _, p := range tepqlq.order {
		p(selector)
	}
	if offset := tepqlq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tepqlq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TkExamPartitionQuestionLinkGroupBy is the group-by builder for TkExamPartitionQuestionLink entities.
type TkExamPartitionQuestionLinkGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Aggregate(fns ...AggregateFunc) *TkExamPartitionQuestionLinkGroupBy {
	tepqlgb.fns = append(tepqlgb.fns, fns...)
	return tepqlgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tepqlgb.path(ctx)
	if err != nil {
		return err
	}
	tepqlgb.sql = query
	return tepqlgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tepqlgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tepqlgb.fields) > 1 {
		return nil, errors.New("ent: TkExamPartitionQuestionLinkGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tepqlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) StringsX(ctx context.Context) []string {
	v, err := tepqlgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tepqlgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		err = fmt.Errorf("ent: TkExamPartitionQuestionLinkGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) StringX(ctx context.Context) string {
	v, err := tepqlgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tepqlgb.fields) > 1 {
		return nil, errors.New("ent: TkExamPartitionQuestionLinkGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tepqlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) IntsX(ctx context.Context) []int {
	v, err := tepqlgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tepqlgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		err = fmt.Errorf("ent: TkExamPartitionQuestionLinkGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) IntX(ctx context.Context) int {
	v, err := tepqlgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tepqlgb.fields) > 1 {
		return nil, errors.New("ent: TkExamPartitionQuestionLinkGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tepqlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tepqlgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tepqlgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		err = fmt.Errorf("ent: TkExamPartitionQuestionLinkGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tepqlgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tepqlgb.fields) > 1 {
		return nil, errors.New("ent: TkExamPartitionQuestionLinkGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tepqlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tepqlgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tepqlgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		err = fmt.Errorf("ent: TkExamPartitionQuestionLinkGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) BoolX(ctx context.Context) bool {
	v, err := tepqlgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tepqlgb.fields {
		if !tkexampartitionquestionlink.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tepqlgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tepqlgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tepqlgb *TkExamPartitionQuestionLinkGroupBy) sqlQuery() *sql.Selector {
	selector := tepqlgb.sql
	columns := make([]string, 0, len(tepqlgb.fields)+len(tepqlgb.fns))
	columns = append(columns, tepqlgb.fields...)
	for _, fn := range tepqlgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tepqlgb.fields...)
}

// TkExamPartitionQuestionLinkSelect is the builder for selecting fields of TkExamPartitionQuestionLink entities.
type TkExamPartitionQuestionLinkSelect struct {
	*TkExamPartitionQuestionLinkQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tepqls *TkExamPartitionQuestionLinkSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tepqls.prepareQuery(ctx); err != nil {
		return err
	}
	tepqls.sql = tepqls.TkExamPartitionQuestionLinkQuery.sqlQuery(ctx)
	return tepqls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tepqls *TkExamPartitionQuestionLinkSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tepqls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tepqls *TkExamPartitionQuestionLinkSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tepqls.fields) > 1 {
		return nil, errors.New("ent: TkExamPartitionQuestionLinkSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tepqls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tepqls *TkExamPartitionQuestionLinkSelect) StringsX(ctx context.Context) []string {
	v, err := tepqls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tepqls *TkExamPartitionQuestionLinkSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tepqls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		err = fmt.Errorf("ent: TkExamPartitionQuestionLinkSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tepqls *TkExamPartitionQuestionLinkSelect) StringX(ctx context.Context) string {
	v, err := tepqls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tepqls *TkExamPartitionQuestionLinkSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tepqls.fields) > 1 {
		return nil, errors.New("ent: TkExamPartitionQuestionLinkSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tepqls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tepqls *TkExamPartitionQuestionLinkSelect) IntsX(ctx context.Context) []int {
	v, err := tepqls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tepqls *TkExamPartitionQuestionLinkSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tepqls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		err = fmt.Errorf("ent: TkExamPartitionQuestionLinkSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tepqls *TkExamPartitionQuestionLinkSelect) IntX(ctx context.Context) int {
	v, err := tepqls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tepqls *TkExamPartitionQuestionLinkSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tepqls.fields) > 1 {
		return nil, errors.New("ent: TkExamPartitionQuestionLinkSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tepqls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tepqls *TkExamPartitionQuestionLinkSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tepqls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tepqls *TkExamPartitionQuestionLinkSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tepqls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		err = fmt.Errorf("ent: TkExamPartitionQuestionLinkSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tepqls *TkExamPartitionQuestionLinkSelect) Float64X(ctx context.Context) float64 {
	v, err := tepqls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tepqls *TkExamPartitionQuestionLinkSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tepqls.fields) > 1 {
		return nil, errors.New("ent: TkExamPartitionQuestionLinkSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tepqls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tepqls *TkExamPartitionQuestionLinkSelect) BoolsX(ctx context.Context) []bool {
	v, err := tepqls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tepqls *TkExamPartitionQuestionLinkSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tepqls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		err = fmt.Errorf("ent: TkExamPartitionQuestionLinkSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tepqls *TkExamPartitionQuestionLinkSelect) BoolX(ctx context.Context) bool {
	v, err := tepqls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tepqls *TkExamPartitionQuestionLinkSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tepqls.sqlQuery().Query()
	if err := tepqls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tepqls *TkExamPartitionQuestionLinkSelect) sqlQuery() sql.Querier {
	selector := tepqls.sql
	selector.Select(selector.Columns(tepqls.fields...)...)
	return selector
}
