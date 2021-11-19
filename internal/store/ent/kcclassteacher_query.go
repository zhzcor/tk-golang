// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/kcclass"
	"gserver/internal/store/ent/kcclassteacher"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/teacher"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KcClassTeacherQuery is the builder for querying KcClassTeacher entities.
type KcClassTeacherQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.KcClassTeacher
	// eager-loading edges.
	withTeacher *TeacherQuery
	withClass   *KcClassQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KcClassTeacherQuery builder.
func (kctq *KcClassTeacherQuery) Where(ps ...predicate.KcClassTeacher) *KcClassTeacherQuery {
	kctq.predicates = append(kctq.predicates, ps...)
	return kctq
}

// Limit adds a limit step to the query.
func (kctq *KcClassTeacherQuery) Limit(limit int) *KcClassTeacherQuery {
	kctq.limit = &limit
	return kctq
}

// Offset adds an offset step to the query.
func (kctq *KcClassTeacherQuery) Offset(offset int) *KcClassTeacherQuery {
	kctq.offset = &offset
	return kctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kctq *KcClassTeacherQuery) Unique(unique bool) *KcClassTeacherQuery {
	kctq.unique = &unique
	return kctq
}

// Order adds an order step to the query.
func (kctq *KcClassTeacherQuery) Order(o ...OrderFunc) *KcClassTeacherQuery {
	kctq.order = append(kctq.order, o...)
	return kctq
}

// QueryTeacher chains the current query on the "teacher" edge.
func (kctq *KcClassTeacherQuery) QueryTeacher() *TeacherQuery {
	query := &TeacherQuery{config: kctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := kctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kcclassteacher.Table, kcclassteacher.FieldID, selector),
			sqlgraph.To(teacher.Table, teacher.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, kcclassteacher.TeacherTable, kcclassteacher.TeacherColumn),
		)
		fromU = sqlgraph.SetNeighbors(kctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClass chains the current query on the "class" edge.
func (kctq *KcClassTeacherQuery) QueryClass() *KcClassQuery {
	query := &KcClassQuery{config: kctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := kctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kcclassteacher.Table, kcclassteacher.FieldID, selector),
			sqlgraph.To(kcclass.Table, kcclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, kcclassteacher.ClassTable, kcclassteacher.ClassColumn),
		)
		fromU = sqlgraph.SetNeighbors(kctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first KcClassTeacher entity from the query.
// Returns a *NotFoundError when no KcClassTeacher was found.
func (kctq *KcClassTeacherQuery) First(ctx context.Context) (*KcClassTeacher, error) {
	nodes, err := kctq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{kcclassteacher.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kctq *KcClassTeacherQuery) FirstX(ctx context.Context) *KcClassTeacher {
	node, err := kctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first KcClassTeacher ID from the query.
// Returns a *NotFoundError when no KcClassTeacher ID was found.
func (kctq *KcClassTeacherQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kctq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{kcclassteacher.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kctq *KcClassTeacherQuery) FirstIDX(ctx context.Context) int {
	id, err := kctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single KcClassTeacher entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one KcClassTeacher entity is not found.
// Returns a *NotFoundError when no KcClassTeacher entities are found.
func (kctq *KcClassTeacherQuery) Only(ctx context.Context) (*KcClassTeacher, error) {
	nodes, err := kctq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{kcclassteacher.Label}
	default:
		return nil, &NotSingularError{kcclassteacher.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kctq *KcClassTeacherQuery) OnlyX(ctx context.Context) *KcClassTeacher {
	node, err := kctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only KcClassTeacher ID in the query.
// Returns a *NotSingularError when exactly one KcClassTeacher ID is not found.
// Returns a *NotFoundError when no entities are found.
func (kctq *KcClassTeacherQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kctq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{kcclassteacher.Label}
	default:
		err = &NotSingularError{kcclassteacher.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kctq *KcClassTeacherQuery) OnlyIDX(ctx context.Context) int {
	id, err := kctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KcClassTeachers.
func (kctq *KcClassTeacherQuery) All(ctx context.Context) ([]*KcClassTeacher, error) {
	if err := kctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kctq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kctq *KcClassTeacherQuery) AllX(ctx context.Context) []*KcClassTeacher {
	nodes, err := kctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of KcClassTeacher IDs.
func (kctq *KcClassTeacherQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := kctq.Select(kcclassteacher.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kctq *KcClassTeacherQuery) IDsX(ctx context.Context) []int {
	ids, err := kctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kctq *KcClassTeacherQuery) Count(ctx context.Context) (int, error) {
	if err := kctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kctq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kctq *KcClassTeacherQuery) CountX(ctx context.Context) int {
	count, err := kctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kctq *KcClassTeacherQuery) Exist(ctx context.Context) (bool, error) {
	if err := kctq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kctq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kctq *KcClassTeacherQuery) ExistX(ctx context.Context) bool {
	exist, err := kctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KcClassTeacherQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kctq *KcClassTeacherQuery) Clone() *KcClassTeacherQuery {
	if kctq == nil {
		return nil
	}
	return &KcClassTeacherQuery{
		config:      kctq.config,
		limit:       kctq.limit,
		offset:      kctq.offset,
		order:       append([]OrderFunc{}, kctq.order...),
		predicates:  append([]predicate.KcClassTeacher{}, kctq.predicates...),
		withTeacher: kctq.withTeacher.Clone(),
		withClass:   kctq.withClass.Clone(),
		// clone intermediate query.
		sql:  kctq.sql.Clone(),
		path: kctq.path,
	}
}

// WithTeacher tells the query-builder to eager-load the nodes that are connected to
// the "teacher" edge. The optional arguments are used to configure the query builder of the edge.
func (kctq *KcClassTeacherQuery) WithTeacher(opts ...func(*TeacherQuery)) *KcClassTeacherQuery {
	query := &TeacherQuery{config: kctq.config}
	for _, opt := range opts {
		opt(query)
	}
	kctq.withTeacher = query
	return kctq
}

// WithClass tells the query-builder to eager-load the nodes that are connected to
// the "class" edge. The optional arguments are used to configure the query builder of the edge.
func (kctq *KcClassTeacherQuery) WithClass(opts ...func(*KcClassQuery)) *KcClassTeacherQuery {
	query := &KcClassQuery{config: kctq.config}
	for _, opt := range opts {
		opt(query)
	}
	kctq.withClass = query
	return kctq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ShowStatus uint8 `json:"show_status"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.KcClassTeacher.Query().
//		GroupBy(kcclassteacher.FieldShowStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (kctq *KcClassTeacherQuery) GroupBy(field string, fields ...string) *KcClassTeacherGroupBy {
	group := &KcClassTeacherGroupBy{config: kctq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kctq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ShowStatus uint8 `json:"show_status"`
//	}
//
//	client.KcClassTeacher.Query().
//		Select(kcclassteacher.FieldShowStatus).
//		Scan(ctx, &v)
//
func (kctq *KcClassTeacherQuery) Select(field string, fields ...string) *KcClassTeacherSelect {
	kctq.fields = append([]string{field}, fields...)
	return &KcClassTeacherSelect{KcClassTeacherQuery: kctq}
}

func (kctq *KcClassTeacherQuery) prepareQuery(ctx context.Context) error {
	for _, f := range kctq.fields {
		if !kcclassteacher.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kctq.path != nil {
		prev, err := kctq.path(ctx)
		if err != nil {
			return err
		}
		kctq.sql = prev
	}
	return nil
}

func (kctq *KcClassTeacherQuery) sqlAll(ctx context.Context) ([]*KcClassTeacher, error) {
	var (
		nodes       = []*KcClassTeacher{}
		_spec       = kctq.querySpec()
		loadedTypes = [2]bool{
			kctq.withTeacher != nil,
			kctq.withClass != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &KcClassTeacher{config: kctq.config}
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
	if err := sqlgraph.QueryNodes(ctx, kctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := kctq.withTeacher; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*KcClassTeacher)
		for i := range nodes {
			fk := nodes[i].TeacherID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(teacher.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "teacher_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Teacher = n
			}
		}
	}

	if query := kctq.withClass; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*KcClassTeacher)
		for i := range nodes {
			fk := nodes[i].ClassID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(kcclass.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "class_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Class = n
			}
		}
	}

	return nodes, nil
}

func (kctq *KcClassTeacherQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kctq.querySpec()
	return sqlgraph.CountNodes(ctx, kctq.driver, _spec)
}

func (kctq *KcClassTeacherQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := kctq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (kctq *KcClassTeacherQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kcclassteacher.Table,
			Columns: kcclassteacher.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcclassteacher.FieldID,
			},
		},
		From:   kctq.sql,
		Unique: true,
	}
	if unique := kctq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := kctq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kcclassteacher.FieldID)
		for i := range fields {
			if fields[i] != kcclassteacher.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kctq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kctq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kctq *KcClassTeacherQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kctq.driver.Dialect())
	t1 := builder.Table(kcclassteacher.Table)
	selector := builder.Select(t1.Columns(kcclassteacher.Columns...)...).From(t1)
	if kctq.sql != nil {
		selector = kctq.sql
		selector.Select(selector.Columns(kcclassteacher.Columns...)...)
	}
	for _, p := range kctq.predicates {
		p(selector)
	}
	for _, p := range kctq.order {
		p(selector)
	}
	if offset := kctq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kctq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KcClassTeacherGroupBy is the group-by builder for KcClassTeacher entities.
type KcClassTeacherGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kctgb *KcClassTeacherGroupBy) Aggregate(fns ...AggregateFunc) *KcClassTeacherGroupBy {
	kctgb.fns = append(kctgb.fns, fns...)
	return kctgb
}

// Scan applies the group-by query and scans the result into the given value.
func (kctgb *KcClassTeacherGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := kctgb.path(ctx)
	if err != nil {
		return err
	}
	kctgb.sql = query
	return kctgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kctgb *KcClassTeacherGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := kctgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (kctgb *KcClassTeacherGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(kctgb.fields) > 1 {
		return nil, errors.New("ent: KcClassTeacherGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := kctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kctgb *KcClassTeacherGroupBy) StringsX(ctx context.Context) []string {
	v, err := kctgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kctgb *KcClassTeacherGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kctgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcclassteacher.Label}
	default:
		err = fmt.Errorf("ent: KcClassTeacherGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kctgb *KcClassTeacherGroupBy) StringX(ctx context.Context) string {
	v, err := kctgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (kctgb *KcClassTeacherGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(kctgb.fields) > 1 {
		return nil, errors.New("ent: KcClassTeacherGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := kctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kctgb *KcClassTeacherGroupBy) IntsX(ctx context.Context) []int {
	v, err := kctgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kctgb *KcClassTeacherGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kctgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcclassteacher.Label}
	default:
		err = fmt.Errorf("ent: KcClassTeacherGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kctgb *KcClassTeacherGroupBy) IntX(ctx context.Context) int {
	v, err := kctgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (kctgb *KcClassTeacherGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(kctgb.fields) > 1 {
		return nil, errors.New("ent: KcClassTeacherGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := kctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kctgb *KcClassTeacherGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := kctgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kctgb *KcClassTeacherGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kctgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcclassteacher.Label}
	default:
		err = fmt.Errorf("ent: KcClassTeacherGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kctgb *KcClassTeacherGroupBy) Float64X(ctx context.Context) float64 {
	v, err := kctgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (kctgb *KcClassTeacherGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(kctgb.fields) > 1 {
		return nil, errors.New("ent: KcClassTeacherGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := kctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kctgb *KcClassTeacherGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := kctgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kctgb *KcClassTeacherGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kctgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcclassteacher.Label}
	default:
		err = fmt.Errorf("ent: KcClassTeacherGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kctgb *KcClassTeacherGroupBy) BoolX(ctx context.Context) bool {
	v, err := kctgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kctgb *KcClassTeacherGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range kctgb.fields {
		if !kcclassteacher.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := kctgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kctgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kctgb *KcClassTeacherGroupBy) sqlQuery() *sql.Selector {
	selector := kctgb.sql
	columns := make([]string, 0, len(kctgb.fields)+len(kctgb.fns))
	columns = append(columns, kctgb.fields...)
	for _, fn := range kctgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(kctgb.fields...)
}

// KcClassTeacherSelect is the builder for selecting fields of KcClassTeacher entities.
type KcClassTeacherSelect struct {
	*KcClassTeacherQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (kcts *KcClassTeacherSelect) Scan(ctx context.Context, v interface{}) error {
	if err := kcts.prepareQuery(ctx); err != nil {
		return err
	}
	kcts.sql = kcts.KcClassTeacherQuery.sqlQuery(ctx)
	return kcts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kcts *KcClassTeacherSelect) ScanX(ctx context.Context, v interface{}) {
	if err := kcts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (kcts *KcClassTeacherSelect) Strings(ctx context.Context) ([]string, error) {
	if len(kcts.fields) > 1 {
		return nil, errors.New("ent: KcClassTeacherSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := kcts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kcts *KcClassTeacherSelect) StringsX(ctx context.Context) []string {
	v, err := kcts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (kcts *KcClassTeacherSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kcts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcclassteacher.Label}
	default:
		err = fmt.Errorf("ent: KcClassTeacherSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kcts *KcClassTeacherSelect) StringX(ctx context.Context) string {
	v, err := kcts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (kcts *KcClassTeacherSelect) Ints(ctx context.Context) ([]int, error) {
	if len(kcts.fields) > 1 {
		return nil, errors.New("ent: KcClassTeacherSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := kcts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kcts *KcClassTeacherSelect) IntsX(ctx context.Context) []int {
	v, err := kcts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (kcts *KcClassTeacherSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kcts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcclassteacher.Label}
	default:
		err = fmt.Errorf("ent: KcClassTeacherSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kcts *KcClassTeacherSelect) IntX(ctx context.Context) int {
	v, err := kcts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (kcts *KcClassTeacherSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(kcts.fields) > 1 {
		return nil, errors.New("ent: KcClassTeacherSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := kcts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kcts *KcClassTeacherSelect) Float64sX(ctx context.Context) []float64 {
	v, err := kcts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (kcts *KcClassTeacherSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kcts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcclassteacher.Label}
	default:
		err = fmt.Errorf("ent: KcClassTeacherSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kcts *KcClassTeacherSelect) Float64X(ctx context.Context) float64 {
	v, err := kcts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (kcts *KcClassTeacherSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(kcts.fields) > 1 {
		return nil, errors.New("ent: KcClassTeacherSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := kcts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kcts *KcClassTeacherSelect) BoolsX(ctx context.Context) []bool {
	v, err := kcts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (kcts *KcClassTeacherSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kcts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kcclassteacher.Label}
	default:
		err = fmt.Errorf("ent: KcClassTeacherSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kcts *KcClassTeacherSelect) BoolX(ctx context.Context) bool {
	v, err := kcts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kcts *KcClassTeacherSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := kcts.sqlQuery().Query()
	if err := kcts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kcts *KcClassTeacherSelect) sqlQuery() sql.Querier {
	selector := kcts.sql
	selector.Select(selector.Columns(kcts.fields...)...)
	return selector
}
