// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/tkexampaper"
	"gserver/internal/store/ent/tkexampaperpartition"
	"gserver/internal/store/ent/tkexampaperpartitionscore"
	"gserver/internal/store/ent/tkexampartitionquestionlink"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkExamPaperPartitionQuery is the builder for querying TkExamPaperPartition entities.
type TkExamPaperPartitionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TkExamPaperPartition
	// eager-loading edges.
	withExamPaper           *TkExamPaperQuery
	withExamPartitionLinks  *TkExamPartitionQuestionLinkQuery
	withExamPartitionScores *TkExamPaperPartitionScoreQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TkExamPaperPartitionQuery builder.
func (teppq *TkExamPaperPartitionQuery) Where(ps ...predicate.TkExamPaperPartition) *TkExamPaperPartitionQuery {
	teppq.predicates = append(teppq.predicates, ps...)
	return teppq
}

// Limit adds a limit step to the query.
func (teppq *TkExamPaperPartitionQuery) Limit(limit int) *TkExamPaperPartitionQuery {
	teppq.limit = &limit
	return teppq
}

// Offset adds an offset step to the query.
func (teppq *TkExamPaperPartitionQuery) Offset(offset int) *TkExamPaperPartitionQuery {
	teppq.offset = &offset
	return teppq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (teppq *TkExamPaperPartitionQuery) Unique(unique bool) *TkExamPaperPartitionQuery {
	teppq.unique = &unique
	return teppq
}

// Order adds an order step to the query.
func (teppq *TkExamPaperPartitionQuery) Order(o ...OrderFunc) *TkExamPaperPartitionQuery {
	teppq.order = append(teppq.order, o...)
	return teppq
}

// QueryExamPaper chains the current query on the "exam_paper" edge.
func (teppq *TkExamPaperPartitionQuery) QueryExamPaper() *TkExamPaperQuery {
	query := &TkExamPaperQuery{config: teppq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := teppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := teppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkexampaperpartition.Table, tkexampaperpartition.FieldID, selector),
			sqlgraph.To(tkexampaper.Table, tkexampaper.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tkexampaperpartition.ExamPaperTable, tkexampaperpartition.ExamPaperColumn),
		)
		fromU = sqlgraph.SetNeighbors(teppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExamPartitionLinks chains the current query on the "exam_partition_links" edge.
func (teppq *TkExamPaperPartitionQuery) QueryExamPartitionLinks() *TkExamPartitionQuestionLinkQuery {
	query := &TkExamPartitionQuestionLinkQuery{config: teppq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := teppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := teppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkexampaperpartition.Table, tkexampaperpartition.FieldID, selector),
			sqlgraph.To(tkexampartitionquestionlink.Table, tkexampartitionquestionlink.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, tkexampaperpartition.ExamPartitionLinksTable, tkexampaperpartition.ExamPartitionLinksColumn),
		)
		fromU = sqlgraph.SetNeighbors(teppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExamPartitionScores chains the current query on the "exam_partition_scores" edge.
func (teppq *TkExamPaperPartitionQuery) QueryExamPartitionScores() *TkExamPaperPartitionScoreQuery {
	query := &TkExamPaperPartitionScoreQuery{config: teppq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := teppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := teppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tkexampaperpartition.Table, tkexampaperpartition.FieldID, selector),
			sqlgraph.To(tkexampaperpartitionscore.Table, tkexampaperpartitionscore.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, tkexampaperpartition.ExamPartitionScoresTable, tkexampaperpartition.ExamPartitionScoresColumn),
		)
		fromU = sqlgraph.SetNeighbors(teppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TkExamPaperPartition entity from the query.
// Returns a *NotFoundError when no TkExamPaperPartition was found.
func (teppq *TkExamPaperPartitionQuery) First(ctx context.Context) (*TkExamPaperPartition, error) {
	nodes, err := teppq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tkexampaperpartition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (teppq *TkExamPaperPartitionQuery) FirstX(ctx context.Context) *TkExamPaperPartition {
	node, err := teppq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TkExamPaperPartition ID from the query.
// Returns a *NotFoundError when no TkExamPaperPartition ID was found.
func (teppq *TkExamPaperPartitionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = teppq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tkexampaperpartition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (teppq *TkExamPaperPartitionQuery) FirstIDX(ctx context.Context) int {
	id, err := teppq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TkExamPaperPartition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TkExamPaperPartition entity is not found.
// Returns a *NotFoundError when no TkExamPaperPartition entities are found.
func (teppq *TkExamPaperPartitionQuery) Only(ctx context.Context) (*TkExamPaperPartition, error) {
	nodes, err := teppq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tkexampaperpartition.Label}
	default:
		return nil, &NotSingularError{tkexampaperpartition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (teppq *TkExamPaperPartitionQuery) OnlyX(ctx context.Context) *TkExamPaperPartition {
	node, err := teppq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TkExamPaperPartition ID in the query.
// Returns a *NotSingularError when exactly one TkExamPaperPartition ID is not found.
// Returns a *NotFoundError when no entities are found.
func (teppq *TkExamPaperPartitionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = teppq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tkexampaperpartition.Label}
	default:
		err = &NotSingularError{tkexampaperpartition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (teppq *TkExamPaperPartitionQuery) OnlyIDX(ctx context.Context) int {
	id, err := teppq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TkExamPaperPartitions.
func (teppq *TkExamPaperPartitionQuery) All(ctx context.Context) ([]*TkExamPaperPartition, error) {
	if err := teppq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return teppq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (teppq *TkExamPaperPartitionQuery) AllX(ctx context.Context) []*TkExamPaperPartition {
	nodes, err := teppq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TkExamPaperPartition IDs.
func (teppq *TkExamPaperPartitionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := teppq.Select(tkexampaperpartition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (teppq *TkExamPaperPartitionQuery) IDsX(ctx context.Context) []int {
	ids, err := teppq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (teppq *TkExamPaperPartitionQuery) Count(ctx context.Context) (int, error) {
	if err := teppq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return teppq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (teppq *TkExamPaperPartitionQuery) CountX(ctx context.Context) int {
	count, err := teppq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (teppq *TkExamPaperPartitionQuery) Exist(ctx context.Context) (bool, error) {
	if err := teppq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return teppq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (teppq *TkExamPaperPartitionQuery) ExistX(ctx context.Context) bool {
	exist, err := teppq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TkExamPaperPartitionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (teppq *TkExamPaperPartitionQuery) Clone() *TkExamPaperPartitionQuery {
	if teppq == nil {
		return nil
	}
	return &TkExamPaperPartitionQuery{
		config:                  teppq.config,
		limit:                   teppq.limit,
		offset:                  teppq.offset,
		order:                   append([]OrderFunc{}, teppq.order...),
		predicates:              append([]predicate.TkExamPaperPartition{}, teppq.predicates...),
		withExamPaper:           teppq.withExamPaper.Clone(),
		withExamPartitionLinks:  teppq.withExamPartitionLinks.Clone(),
		withExamPartitionScores: teppq.withExamPartitionScores.Clone(),
		// clone intermediate query.
		sql:  teppq.sql.Clone(),
		path: teppq.path,
	}
}

// WithExamPaper tells the query-builder to eager-load the nodes that are connected to
// the "exam_paper" edge. The optional arguments are used to configure the query builder of the edge.
func (teppq *TkExamPaperPartitionQuery) WithExamPaper(opts ...func(*TkExamPaperQuery)) *TkExamPaperPartitionQuery {
	query := &TkExamPaperQuery{config: teppq.config}
	for _, opt := range opts {
		opt(query)
	}
	teppq.withExamPaper = query
	return teppq
}

// WithExamPartitionLinks tells the query-builder to eager-load the nodes that are connected to
// the "exam_partition_links" edge. The optional arguments are used to configure the query builder of the edge.
func (teppq *TkExamPaperPartitionQuery) WithExamPartitionLinks(opts ...func(*TkExamPartitionQuestionLinkQuery)) *TkExamPaperPartitionQuery {
	query := &TkExamPartitionQuestionLinkQuery{config: teppq.config}
	for _, opt := range opts {
		opt(query)
	}
	teppq.withExamPartitionLinks = query
	return teppq
}

// WithExamPartitionScores tells the query-builder to eager-load the nodes that are connected to
// the "exam_partition_scores" edge. The optional arguments are used to configure the query builder of the edge.
func (teppq *TkExamPaperPartitionQuery) WithExamPartitionScores(opts ...func(*TkExamPaperPartitionScoreQuery)) *TkExamPaperPartitionQuery {
	query := &TkExamPaperPartitionScoreQuery{config: teppq.config}
	for _, opt := range opts {
		opt(query)
	}
	teppq.withExamPartitionScores = query
	return teppq
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
//	client.TkExamPaperPartition.Query().
//		GroupBy(tkexampaperpartition.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (teppq *TkExamPaperPartitionQuery) GroupBy(field string, fields ...string) *TkExamPaperPartitionGroupBy {
	group := &TkExamPaperPartitionGroupBy{config: teppq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := teppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return teppq.sqlQuery(ctx), nil
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
//	client.TkExamPaperPartition.Query().
//		Select(tkexampaperpartition.FieldUUID).
//		Scan(ctx, &v)
//
func (teppq *TkExamPaperPartitionQuery) Select(field string, fields ...string) *TkExamPaperPartitionSelect {
	teppq.fields = append([]string{field}, fields...)
	return &TkExamPaperPartitionSelect{TkExamPaperPartitionQuery: teppq}
}

func (teppq *TkExamPaperPartitionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range teppq.fields {
		if !tkexampaperpartition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if teppq.path != nil {
		prev, err := teppq.path(ctx)
		if err != nil {
			return err
		}
		teppq.sql = prev
	}
	return nil
}

func (teppq *TkExamPaperPartitionQuery) sqlAll(ctx context.Context) ([]*TkExamPaperPartition, error) {
	var (
		nodes       = []*TkExamPaperPartition{}
		_spec       = teppq.querySpec()
		loadedTypes = [3]bool{
			teppq.withExamPaper != nil,
			teppq.withExamPartitionLinks != nil,
			teppq.withExamPartitionScores != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TkExamPaperPartition{config: teppq.config}
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
	if err := sqlgraph.QueryNodes(ctx, teppq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := teppq.withExamPaper; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TkExamPaperPartition)
		for i := range nodes {
			fk := nodes[i].ExamPaperID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(tkexampaper.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "exam_paper_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ExamPaper = n
			}
		}
	}

	if query := teppq.withExamPartitionLinks; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*TkExamPaperPartition)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ExamPartitionLinks = []*TkExamPartitionQuestionLink{}
		}
		query.Where(predicate.TkExamPartitionQuestionLink(func(s *sql.Selector) {
			s.Where(sql.InValues(tkexampaperpartition.ExamPartitionLinksColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ExamPaperPartitionID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "exam_paper_partition_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.ExamPartitionLinks = append(node.Edges.ExamPartitionLinks, n)
		}
	}

	if query := teppq.withExamPartitionScores; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*TkExamPaperPartition)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ExamPartitionScores = []*TkExamPaperPartitionScore{}
		}
		query.Where(predicate.TkExamPaperPartitionScore(func(s *sql.Selector) {
			s.Where(sql.InValues(tkexampaperpartition.ExamPartitionScoresColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ExamPaperPartitionID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "exam_paper_partition_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.ExamPartitionScores = append(node.Edges.ExamPartitionScores, n)
		}
	}

	return nodes, nil
}

func (teppq *TkExamPaperPartitionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := teppq.querySpec()
	return sqlgraph.CountNodes(ctx, teppq.driver, _spec)
}

func (teppq *TkExamPaperPartitionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := teppq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (teppq *TkExamPaperPartitionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tkexampaperpartition.Table,
			Columns: tkexampaperpartition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkexampaperpartition.FieldID,
			},
		},
		From:   teppq.sql,
		Unique: true,
	}
	if unique := teppq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := teppq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tkexampaperpartition.FieldID)
		for i := range fields {
			if fields[i] != tkexampaperpartition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := teppq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := teppq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := teppq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := teppq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (teppq *TkExamPaperPartitionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(teppq.driver.Dialect())
	t1 := builder.Table(tkexampaperpartition.Table)
	selector := builder.Select(t1.Columns(tkexampaperpartition.Columns...)...).From(t1)
	if teppq.sql != nil {
		selector = teppq.sql
		selector.Select(selector.Columns(tkexampaperpartition.Columns...)...)
	}
	for _, p := range teppq.predicates {
		p(selector)
	}
	for _, p := range teppq.order {
		p(selector)
	}
	if offset := teppq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := teppq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TkExamPaperPartitionGroupBy is the group-by builder for TkExamPaperPartition entities.
type TkExamPaperPartitionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (teppgb *TkExamPaperPartitionGroupBy) Aggregate(fns ...AggregateFunc) *TkExamPaperPartitionGroupBy {
	teppgb.fns = append(teppgb.fns, fns...)
	return teppgb
}

// Scan applies the group-by query and scans the result into the given value.
func (teppgb *TkExamPaperPartitionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := teppgb.path(ctx)
	if err != nil {
		return err
	}
	teppgb.sql = query
	return teppgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (teppgb *TkExamPaperPartitionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := teppgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (teppgb *TkExamPaperPartitionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(teppgb.fields) > 1 {
		return nil, errors.New("ent: TkExamPaperPartitionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := teppgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (teppgb *TkExamPaperPartitionGroupBy) StringsX(ctx context.Context) []string {
	v, err := teppgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (teppgb *TkExamPaperPartitionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = teppgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampaperpartition.Label}
	default:
		err = fmt.Errorf("ent: TkExamPaperPartitionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (teppgb *TkExamPaperPartitionGroupBy) StringX(ctx context.Context) string {
	v, err := teppgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (teppgb *TkExamPaperPartitionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(teppgb.fields) > 1 {
		return nil, errors.New("ent: TkExamPaperPartitionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := teppgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (teppgb *TkExamPaperPartitionGroupBy) IntsX(ctx context.Context) []int {
	v, err := teppgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (teppgb *TkExamPaperPartitionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = teppgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampaperpartition.Label}
	default:
		err = fmt.Errorf("ent: TkExamPaperPartitionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (teppgb *TkExamPaperPartitionGroupBy) IntX(ctx context.Context) int {
	v, err := teppgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (teppgb *TkExamPaperPartitionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(teppgb.fields) > 1 {
		return nil, errors.New("ent: TkExamPaperPartitionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := teppgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (teppgb *TkExamPaperPartitionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := teppgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (teppgb *TkExamPaperPartitionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = teppgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampaperpartition.Label}
	default:
		err = fmt.Errorf("ent: TkExamPaperPartitionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (teppgb *TkExamPaperPartitionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := teppgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (teppgb *TkExamPaperPartitionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(teppgb.fields) > 1 {
		return nil, errors.New("ent: TkExamPaperPartitionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := teppgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (teppgb *TkExamPaperPartitionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := teppgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (teppgb *TkExamPaperPartitionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = teppgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampaperpartition.Label}
	default:
		err = fmt.Errorf("ent: TkExamPaperPartitionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (teppgb *TkExamPaperPartitionGroupBy) BoolX(ctx context.Context) bool {
	v, err := teppgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (teppgb *TkExamPaperPartitionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range teppgb.fields {
		if !tkexampaperpartition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := teppgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := teppgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (teppgb *TkExamPaperPartitionGroupBy) sqlQuery() *sql.Selector {
	selector := teppgb.sql
	columns := make([]string, 0, len(teppgb.fields)+len(teppgb.fns))
	columns = append(columns, teppgb.fields...)
	for _, fn := range teppgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(teppgb.fields...)
}

// TkExamPaperPartitionSelect is the builder for selecting fields of TkExamPaperPartition entities.
type TkExamPaperPartitionSelect struct {
	*TkExamPaperPartitionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tepps *TkExamPaperPartitionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tepps.prepareQuery(ctx); err != nil {
		return err
	}
	tepps.sql = tepps.TkExamPaperPartitionQuery.sqlQuery(ctx)
	return tepps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tepps *TkExamPaperPartitionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tepps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tepps *TkExamPaperPartitionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tepps.fields) > 1 {
		return nil, errors.New("ent: TkExamPaperPartitionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tepps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tepps *TkExamPaperPartitionSelect) StringsX(ctx context.Context) []string {
	v, err := tepps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tepps *TkExamPaperPartitionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tepps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampaperpartition.Label}
	default:
		err = fmt.Errorf("ent: TkExamPaperPartitionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tepps *TkExamPaperPartitionSelect) StringX(ctx context.Context) string {
	v, err := tepps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tepps *TkExamPaperPartitionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tepps.fields) > 1 {
		return nil, errors.New("ent: TkExamPaperPartitionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tepps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tepps *TkExamPaperPartitionSelect) IntsX(ctx context.Context) []int {
	v, err := tepps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tepps *TkExamPaperPartitionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tepps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampaperpartition.Label}
	default:
		err = fmt.Errorf("ent: TkExamPaperPartitionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tepps *TkExamPaperPartitionSelect) IntX(ctx context.Context) int {
	v, err := tepps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tepps *TkExamPaperPartitionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tepps.fields) > 1 {
		return nil, errors.New("ent: TkExamPaperPartitionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tepps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tepps *TkExamPaperPartitionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tepps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tepps *TkExamPaperPartitionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tepps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampaperpartition.Label}
	default:
		err = fmt.Errorf("ent: TkExamPaperPartitionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tepps *TkExamPaperPartitionSelect) Float64X(ctx context.Context) float64 {
	v, err := tepps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tepps *TkExamPaperPartitionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tepps.fields) > 1 {
		return nil, errors.New("ent: TkExamPaperPartitionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tepps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tepps *TkExamPaperPartitionSelect) BoolsX(ctx context.Context) []bool {
	v, err := tepps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tepps *TkExamPaperPartitionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tepps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tkexampaperpartition.Label}
	default:
		err = fmt.Errorf("ent: TkExamPaperPartitionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tepps *TkExamPaperPartitionSelect) BoolX(ctx context.Context) bool {
	v, err := tepps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tepps *TkExamPaperPartitionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tepps.sqlQuery().Query()
	if err := tepps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tepps *TkExamPaperPartitionSelect) sqlQuery() sql.Querier {
	selector := tepps.sql
	selector.Select(selector.Columns(tepps.fields...)...)
	return selector
}
