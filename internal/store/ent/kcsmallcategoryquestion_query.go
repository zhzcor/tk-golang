// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/kccoursesmallcategory"
	"gserver/internal/store/ent/kcsmallcategoryquestion"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/tkquestion"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KcSmallCategoryQuestionQuery is the builder for querying KcSmallCategoryQuestion entities.
type KcSmallCategoryQuestionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.KcSmallCategoryQuestion
	// eager-loading edges.
	withQuestion            *TkQuestionQuery
	withCourseSmallCategory *KcCourseSmallCategoryQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KcSmallCategoryQuestionQuery builder.
func (kscqq *KcSmallCategoryQuestionQuery) Where(ps ...predicate.KcSmallCategoryQuestion) *KcSmallCategoryQuestionQuery {
	kscqq.predicates = append(kscqq.predicates, ps...)
	return kscqq
}

// Limit adds a limit step to the query.
func (kscqq *KcSmallCategoryQuestionQuery) Limit(limit int) *KcSmallCategoryQuestionQuery {
	kscqq.limit = &limit
	return kscqq
}

// Offset adds an offset step to the query.
func (kscqq *KcSmallCategoryQuestionQuery) Offset(offset int) *KcSmallCategoryQuestionQuery {
	kscqq.offset = &offset
	return kscqq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kscqq *KcSmallCategoryQuestionQuery) Unique(unique bool) *KcSmallCategoryQuestionQuery {
	kscqq.unique = &unique
	return kscqq
}

// Order adds an order step to the query.
func (kscqq *KcSmallCategoryQuestionQuery) Order(o ...OrderFunc) *KcSmallCategoryQuestionQuery {
	kscqq.order = append(kscqq.order, o...)
	return kscqq
}

// QueryQuestion chains the current query on the "question" edge.
func (kscqq *KcSmallCategoryQuestionQuery) QueryQuestion() *TkQuestionQuery {
	query := &TkQuestionQuery{config: kscqq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kscqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := kscqq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kcsmallcategoryquestion.Table, kcsmallcategoryquestion.FieldID, selector),
			sqlgraph.To(tkquestion.Table, tkquestion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, kcsmallcategoryquestion.QuestionTable, kcsmallcategoryquestion.QuestionColumn),
		)
		fromU = sqlgraph.SetNeighbors(kscqq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCourseSmallCategory chains the current query on the "course_small_category" edge.
func (kscqq *KcSmallCategoryQuestionQuery) QueryCourseSmallCategory() *KcCourseSmallCategoryQuery {
	query := &KcCourseSmallCategoryQuery{config: kscqq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kscqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := kscqq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kcsmallcategoryquestion.Table, kcsmallcategoryquestion.FieldID, selector),
			sqlgraph.To(kccoursesmallcategory.Table, kccoursesmallcategory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, kcsmallcategoryquestion.CourseSmallCategoryTable, kcsmallcategoryquestion.CourseSmallCategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(kscqq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first KcSmallCategoryQuestion entity from the query.
// Returns a *NotFoundError when no KcSmallCategoryQuestion was found.
func (kscqq *KcSmallCategoryQuestionQuery) First(ctx context.Context) (*KcSmallCategoryQuestion, error) {
	nodes, err := kscqq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{kcsmallcategoryquestion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kscqq *KcSmallCategoryQuestionQuery) FirstX(ctx context.Context) *KcSmallCategoryQuestion {
	node, err := kscqq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first KcSmallCategoryQuestion ID from the query.
// Returns a *NotFoundError when no KcSmallCategoryQuestion ID was found.
func (kscqq *KcSmallCategoryQuestionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kscqq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{kcsmallcategoryquestion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kscqq *KcSmallCategoryQuestionQuery) FirstIDX(ctx context.Context) int {
	id, err := kscqq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single KcSmallCategoryQuestion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one KcSmallCategoryQuestion entity is not found.
// Returns a *NotFoundError when no KcSmallCategoryQuestion entities are found.
func (kscqq *KcSmallCategoryQuestionQuery) Only(ctx context.Context) (*KcSmallCategoryQuestion, error) {
	nodes, err := kscqq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{kcsmallcategoryquestion.Label}
	default:
		return nil, &NotSingularError{kcsmallcategoryquestion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kscqq *KcSmallCategoryQuestionQuery) OnlyX(ctx context.Context) *KcSmallCategoryQuestion {
	node, err := kscqq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only KcSmallCategoryQuestion ID in the query.
// Returns a *NotSingularError when exactly one KcSmallCategoryQuestion ID is not found.
// Returns a *NotFoundError when no entities are found.
func (kscqq *KcSmallCategoryQuestionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kscqq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{kcsmallcategoryquestion.Label}
	default:
		err = &NotSingularError{kcsmallcategoryquestion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kscqq *KcSmallCategoryQuestionQuery) OnlyIDX(ctx context.Context) int {
	id, err := kscqq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KcSmallCategoryQuestions.
func (kscqq *KcSmallCategoryQuestionQuery) All(ctx context.Context) ([]*KcSmallCategoryQuestion, error) {
	if err := kscqq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kscqq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kscqq *KcSmallCategoryQuestionQuery) AllX(ctx context.Context) []*KcSmallCategoryQuestion {
	nodes, err := kscqq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of KcSmallCategoryQuestion IDs.
func (kscqq *KcSmallCategoryQuestionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := kscqq.Select(kcsmallcategoryquestion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kscqq *KcSmallCategoryQuestionQuery) IDsX(ctx context.Context) []int {
	ids, err := kscqq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kscqq *KcSmallCategoryQuestionQuery) Count(ctx context.Context) (int, error) {
	if err := kscqq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kscqq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kscqq *KcSmallCategoryQuestionQuery) CountX(ctx context.Context) int {
	count, err := kscqq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kscqq *KcSmallCategoryQuestionQuery) Exist(ctx context.Context) (bool, error) {
	if err := kscqq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kscqq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kscqq *KcSmallCategoryQuestionQuery) ExistX(ctx context.Context) bool {
	exist, err := kscqq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KcSmallCategoryQuestionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kscqq *KcSmallCategoryQuestionQuery) Clone() *KcSmallCategoryQuestionQuery {
	if kscqq == nil {
		return nil
	}
	return &KcSmallCategoryQuestionQuery{
		config:                  kscqq.config,
		limit:                   kscqq.limit,
		offset:                  kscqq.offset,
		order:                   append([]OrderFunc{}, kscqq.order...),
		predicates:              append([]predicate.KcSmallCategoryQuestion{}, kscqq.predicates...),
		withQuestion:            kscqq.withQuestion.Clone(),
		withCourseSmallCategory: kscqq.withCourseSmallCategory.Clone(),
		// clone intermediate query.
		sql:  kscqq.sql.Clone(),
		path: kscqq.path,
	}
}

// WithQuestion tells the query-builder to eager-load the nodes that are connected to
// the "question" edge. The optional arguments are used to configure the query builder of the edge.
func (kscqq *KcSmallCategoryQuestionQuery) WithQuestion(opts ...func(*TkQuestionQuery)) *KcSmallCategoryQuestionQuery {
	query := &TkQuestionQuery{config: kscqq.config}
	for _, opt := range opts {
		opt(query)
	}
	kscqq.withQuestion = query
	return kscqq
}

// WithCourseSmallCategory tells the query-builder to eager-load the nodes that are connected to
// the "course_small_category" edge. The optional arguments are used to configure the query builder of the edge.
func (kscqq *KcSmallCategoryQuestionQuery) WithCourseSmallCategory(opts ...func(*KcCourseSmallCategoryQuery)) *KcSmallCategoryQuestionQuery {
	query := &KcCourseSmallCategoryQuery{config: kscqq.config}
	for _, opt := range opts {
		opt(query)
	}
	kscqq.withCourseSmallCategory = query
	return kscqq
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
//	client.KcSmallCategoryQuestion.Query().
//		GroupBy(kcsmallcategoryquestion.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (kscqq *KcSmallCategoryQuestionQuery) GroupBy(field string, fields ...string) *KcSmallCategoryQuestionGroupBy {
	group := &KcSmallCategoryQuestionGroupBy{config: kscqq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kscqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kscqq.sqlQuery(ctx), nil
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
//	client.KcSmallCategoryQuestion.Query().
//		Select(kcsmallcategoryquestion.FieldUUID).
//		Scan(ctx, &v)
//
func (kscqq *KcSmallCategoryQuestionQuery) Select(field string, fields ...string) *KcSmallCategoryQuestionSelect {
	kscqq.fields = append([]string{field}, fields...)
	return &KcSmallCategoryQuestionSelect{KcSmallCategoryQuestionQuery: kscqq}
}

func (kscqq *KcSmallCategoryQuestionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range kscqq.fields {
		if !kcsmallcategoryquestion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kscqq.path != nil {
		prev, err := kscqq.path(ctx)
		if err != nil {
			return err
		}
		kscqq.sql = prev
	}
	return nil
}

func (kscqq *KcSmallCategoryQuestionQuery) sqlAll(ctx context.Context) ([]*KcSmallCategoryQuestion, error) {
	var (
		nodes       = []*KcSmallCategoryQuestion{}
		_spec       = kscqq.querySpec()
		loadedTypes = [2]bool{
			kscqq.withQuestion != nil,
			kscqq.withCourseSmallCategory != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &KcSmallCategoryQuestion{config: kscqq.config}
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
	if err := sqlgraph.QueryNodes(ctx, kscqq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := kscqq.withQuestion; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*KcSmallCategoryQuestion)
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

	if query := kscqq.withCourseSmallCategory; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*KcSmallCategoryQuestion)
		for i := range nodes {
			fk := nodes[i].SmallCategoryID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(kccoursesmallcategory.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "small_category_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.CourseSmallCategory = n
			}
		}
	}

	return nodes, nil
}

func (kscqq *KcSmallCategoryQuestionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kscqq.querySpec()
	return sqlgraph.CountNodes(ctx, kscqq.driver, _spec)
}

func (kscqq *KcSmallCategoryQuestionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := kscqq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (kscqq *KcSmallCategoryQuestionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kcsmallcategoryquestion.Table,
			Columns: kcsmallcategoryquestion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcsmallcategoryquestion.FieldID,
			},
		},
		From:   kscqq.sql,
		Unique: true,
	}
	if unique := kscqq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := kscqq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kcsmallcategoryquestion.FieldID)
		for i := range fields {
			if fields[i] != kcsmallcategoryquestion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kscqq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kscqq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kscqq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kscqq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kscqq *KcSmallCategoryQuestionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kscqq.driver.Dialect())
	t1 := builder.Table(kcsmallcategoryquestion.Table)
	selector := builder.Select(t1.Columns(kcsmallcategoryquestion.Columns...)...).From(t1)
	if kscqq.sql != nil {
		selector = kscqq.sql
		selector.Select(selector.Columns(kcsmallcategoryquestion.Columns...)...)
	}
	for _, p := range kscqq.predicates {
		p(selector)
	}
	for _, p := range kscqq.order {
		p(selector)
	}
	if offset := kscqq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kscqq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KcSmallCategoryQuestionGroupBy is the group-by builder for KcSmallCategoryQuestion entities.
type KcSmallCategoryQuestionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Aggregate(fns ...AggregateFunc) *KcSmallCategoryQuestionGroupBy {
	kscqgb.fns = append(kscqgb.fns, fns...)
	return kscqgb
}

// Scan applies the group-by query and scans the result into the given value.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := kscqgb.path(ctx)
	if err != nil {
		return err
	}
	kscqgb.sql = query
	return kscqgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kscqgb *KcSmallCategoryQuestionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := kscqgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(kscqgb.fields) > 1 {
		return nil, errors.New("ent: KcSmallCategoryQuestionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := kscqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kscqgb *KcSmallCategoryQuestionGroupBy) StringsX(ctx context.Context) []string {
	v, err := kscqgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kscqgb *KcSmallCategoryQuestionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kscqgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcsmallcategoryquestion.Label}
	default:
		err = fmt.Errorf("ent: KcSmallCategoryQuestionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kscqgb *KcSmallCategoryQuestionGroupBy) StringX(ctx context.Context) string {
	v, err := kscqgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(kscqgb.fields) > 1 {
		return nil, errors.New("ent: KcSmallCategoryQuestionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := kscqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kscqgb *KcSmallCategoryQuestionGroupBy) IntsX(ctx context.Context) []int {
	v, err := kscqgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kscqgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcsmallcategoryquestion.Label}
	default:
		err = fmt.Errorf("ent: KcSmallCategoryQuestionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kscqgb *KcSmallCategoryQuestionGroupBy) IntX(ctx context.Context) int {
	v, err := kscqgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(kscqgb.fields) > 1 {
		return nil, errors.New("ent: KcSmallCategoryQuestionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := kscqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := kscqgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kscqgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcsmallcategoryquestion.Label}
	default:
		err = fmt.Errorf("ent: KcSmallCategoryQuestionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := kscqgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(kscqgb.fields) > 1 {
		return nil, errors.New("ent: KcSmallCategoryQuestionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := kscqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kscqgb *KcSmallCategoryQuestionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := kscqgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kscqgb *KcSmallCategoryQuestionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kscqgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcsmallcategoryquestion.Label}
	default:
		err = fmt.Errorf("ent: KcSmallCategoryQuestionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kscqgb *KcSmallCategoryQuestionGroupBy) BoolX(ctx context.Context) bool {
	v, err := kscqgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kscqgb *KcSmallCategoryQuestionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range kscqgb.fields {
		if !kcsmallcategoryquestion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := kscqgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kscqgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kscqgb *KcSmallCategoryQuestionGroupBy) sqlQuery() *sql.Selector {
	selector := kscqgb.sql
	columns := make([]string, 0, len(kscqgb.fields)+len(kscqgb.fns))
	columns = append(columns, kscqgb.fields...)
	for _, fn := range kscqgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(kscqgb.fields...)
}

// KcSmallCategoryQuestionSelect is the builder for selecting fields of KcSmallCategoryQuestion entities.
type KcSmallCategoryQuestionSelect struct {
	*KcSmallCategoryQuestionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (kscqs *KcSmallCategoryQuestionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := kscqs.prepareQuery(ctx); err != nil {
		return err
	}
	kscqs.sql = kscqs.KcSmallCategoryQuestionQuery.sqlQuery(ctx)
	return kscqs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kscqs *KcSmallCategoryQuestionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := kscqs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (kscqs *KcSmallCategoryQuestionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(kscqs.fields) > 1 {
		return nil, errors.New("ent: KcSmallCategoryQuestionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := kscqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kscqs *KcSmallCategoryQuestionSelect) StringsX(ctx context.Context) []string {
	v, err := kscqs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (kscqs *KcSmallCategoryQuestionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kscqs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcsmallcategoryquestion.Label}
	default:
		err = fmt.Errorf("ent: KcSmallCategoryQuestionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kscqs *KcSmallCategoryQuestionSelect) StringX(ctx context.Context) string {
	v, err := kscqs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (kscqs *KcSmallCategoryQuestionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(kscqs.fields) > 1 {
		return nil, errors.New("ent: KcSmallCategoryQuestionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := kscqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kscqs *KcSmallCategoryQuestionSelect) IntsX(ctx context.Context) []int {
	v, err := kscqs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (kscqs *KcSmallCategoryQuestionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kscqs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcsmallcategoryquestion.Label}
	default:
		err = fmt.Errorf("ent: KcSmallCategoryQuestionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kscqs *KcSmallCategoryQuestionSelect) IntX(ctx context.Context) int {
	v, err := kscqs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (kscqs *KcSmallCategoryQuestionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(kscqs.fields) > 1 {
		return nil, errors.New("ent: KcSmallCategoryQuestionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := kscqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kscqs *KcSmallCategoryQuestionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := kscqs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (kscqs *KcSmallCategoryQuestionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kscqs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcsmallcategoryquestion.Label}
	default:
		err = fmt.Errorf("ent: KcSmallCategoryQuestionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kscqs *KcSmallCategoryQuestionSelect) Float64X(ctx context.Context) float64 {
	v, err := kscqs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (kscqs *KcSmallCategoryQuestionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(kscqs.fields) > 1 {
		return nil, errors.New("ent: KcSmallCategoryQuestionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := kscqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kscqs *KcSmallCategoryQuestionSelect) BoolsX(ctx context.Context) []bool {
	v, err := kscqs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (kscqs *KcSmallCategoryQuestionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kscqs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcsmallcategoryquestion.Label}
	default:
		err = fmt.Errorf("ent: KcSmallCategoryQuestionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kscqs *KcSmallCategoryQuestionSelect) BoolX(ctx context.Context) bool {
	v, err := kscqs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kscqs *KcSmallCategoryQuestionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := kscqs.sqlQuery().Query()
	if err := kscqs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kscqs *KcSmallCategoryQuestionSelect) sqlQuery() sql.Querier {
	selector := kscqs.sql
	selector.Select(selector.Columns(kscqs.fields...)...)
	return selector
}
