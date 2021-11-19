// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"tkserver/internal/store/ent/message"
	"tkserver/internal/store/ent/messagetype"
	"tkserver/internal/store/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageTypeQuery is the builder for querying MessageType entities.
type MessageTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MessageType
	// eager-loading edges.
	withMessages *MessageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageTypeQuery builder.
func (mtq *MessageTypeQuery) Where(ps ...predicate.MessageType) *MessageTypeQuery {
	mtq.predicates = append(mtq.predicates, ps...)
	return mtq
}

// Limit adds a limit step to the query.
func (mtq *MessageTypeQuery) Limit(limit int) *MessageTypeQuery {
	mtq.limit = &limit
	return mtq
}

// Offset adds an offset step to the query.
func (mtq *MessageTypeQuery) Offset(offset int) *MessageTypeQuery {
	mtq.offset = &offset
	return mtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mtq *MessageTypeQuery) Unique(unique bool) *MessageTypeQuery {
	mtq.unique = &unique
	return mtq
}

// Order adds an order step to the query.
func (mtq *MessageTypeQuery) Order(o ...OrderFunc) *MessageTypeQuery {
	mtq.order = append(mtq.order, o...)
	return mtq
}

// QueryMessages chains the current query on the "messages" edge.
func (mtq *MessageTypeQuery) QueryMessages() *MessageQuery {
	query := &MessageQuery{config: mtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(messagetype.Table, messagetype.FieldID, selector),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, messagetype.MessagesTable, messagetype.MessagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MessageType entity from the query.
// Returns a *NotFoundError when no MessageType was found.
func (mtq *MessageTypeQuery) First(ctx context.Context) (*MessageType, error) {
	nodes, err := mtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{messagetype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mtq *MessageTypeQuery) FirstX(ctx context.Context) *MessageType {
	node, err := mtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MessageType ID from the query.
// Returns a *NotFoundError when no MessageType ID was found.
func (mtq *MessageTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{messagetype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mtq *MessageTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := mtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MessageType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one MessageType entity is not found.
// Returns a *NotFoundError when no MessageType entities are found.
func (mtq *MessageTypeQuery) Only(ctx context.Context) (*MessageType, error) {
	nodes, err := mtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{messagetype.Label}
	default:
		return nil, &NotSingularError{messagetype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mtq *MessageTypeQuery) OnlyX(ctx context.Context) *MessageType {
	node, err := mtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MessageType ID in the query.
// Returns a *NotSingularError when exactly one MessageType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mtq *MessageTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{messagetype.Label}
	default:
		err = &NotSingularError{messagetype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mtq *MessageTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := mtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MessageTypes.
func (mtq *MessageTypeQuery) All(ctx context.Context) ([]*MessageType, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mtq *MessageTypeQuery) AllX(ctx context.Context) []*MessageType {
	nodes, err := mtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MessageType IDs.
func (mtq *MessageTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mtq.Select(messagetype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mtq *MessageTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := mtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mtq *MessageTypeQuery) Count(ctx context.Context) (int, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mtq *MessageTypeQuery) CountX(ctx context.Context) int {
	count, err := mtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mtq *MessageTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mtq *MessageTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := mtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mtq *MessageTypeQuery) Clone() *MessageTypeQuery {
	if mtq == nil {
		return nil
	}
	return &MessageTypeQuery{
		config:       mtq.config,
		limit:        mtq.limit,
		offset:       mtq.offset,
		order:        append([]OrderFunc{}, mtq.order...),
		predicates:   append([]predicate.MessageType{}, mtq.predicates...),
		withMessages: mtq.withMessages.Clone(),
		// clone intermediate query.
		sql:  mtq.sql.Clone(),
		path: mtq.path,
	}
}

// WithMessages tells the query-builder to eager-load the nodes that are connected to
// the "messages" edge. The optional arguments are used to configure the query builder of the edge.
func (mtq *MessageTypeQuery) WithMessages(opts ...func(*MessageQuery)) *MessageTypeQuery {
	query := &MessageQuery{config: mtq.config}
	for _, opt := range opts {
		opt(query)
	}
	mtq.withMessages = query
	return mtq
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
//	client.MessageType.Query().
//		GroupBy(messagetype.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mtq *MessageTypeQuery) GroupBy(field string, fields ...string) *MessageTypeGroupBy {
	group := &MessageTypeGroupBy{config: mtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mtq.sqlQuery(ctx), nil
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
//	client.MessageType.Query().
//		Select(messagetype.FieldUUID).
//		Scan(ctx, &v)
//
func (mtq *MessageTypeQuery) Select(field string, fields ...string) *MessageTypeSelect {
	mtq.fields = append([]string{field}, fields...)
	return &MessageTypeSelect{MessageTypeQuery: mtq}
}

func (mtq *MessageTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mtq.fields {
		if !messagetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mtq.path != nil {
		prev, err := mtq.path(ctx)
		if err != nil {
			return err
		}
		mtq.sql = prev
	}
	return nil
}

func (mtq *MessageTypeQuery) sqlAll(ctx context.Context) ([]*MessageType, error) {
	var (
		nodes       = []*MessageType{}
		_spec       = mtq.querySpec()
		loadedTypes = [1]bool{
			mtq.withMessages != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MessageType{config: mtq.config}
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
	if err := sqlgraph.QueryNodes(ctx, mtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mtq.withMessages; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*MessageType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Messages = []*Message{}
		}
		query.Where(predicate.Message(func(s *sql.Selector) {
			s.Where(sql.InValues(messagetype.MessagesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.MessageTypeID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "message_type_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Messages = append(node.Edges.Messages, n)
		}
	}

	return nodes, nil
}

func (mtq *MessageTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mtq.querySpec()
	return sqlgraph.CountNodes(ctx, mtq.driver, _spec)
}

func (mtq *MessageTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mtq *MessageTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagetype.Table,
			Columns: messagetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagetype.FieldID,
			},
		},
		From:   mtq.sql,
		Unique: true,
	}
	if unique := mtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagetype.FieldID)
		for i := range fields {
			if fields[i] != messagetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mtq *MessageTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mtq.driver.Dialect())
	t1 := builder.Table(messagetype.Table)
	selector := builder.Select(t1.Columns(messagetype.Columns...)...).From(t1)
	if mtq.sql != nil {
		selector = mtq.sql
		selector.Select(selector.Columns(messagetype.Columns...)...)
	}
	for _, p := range mtq.predicates {
		p(selector)
	}
	for _, p := range mtq.order {
		p(selector)
	}
	if offset := mtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MessageTypeGroupBy is the group-by builder for MessageType entities.
type MessageTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mtgb *MessageTypeGroupBy) Aggregate(fns ...AggregateFunc) *MessageTypeGroupBy {
	mtgb.fns = append(mtgb.fns, fns...)
	return mtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mtgb *MessageTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mtgb.path(ctx)
	if err != nil {
		return err
	}
	mtgb.sql = query
	return mtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mtgb *MessageTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MessageTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MessageTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mtgb *MessageTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := mtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MessageTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagetype.Label}
	default:
		err = fmt.Errorf("ent: MessageTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mtgb *MessageTypeGroupBy) StringX(ctx context.Context) string {
	v, err := mtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MessageTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MessageTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mtgb *MessageTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := mtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MessageTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagetype.Label}
	default:
		err = fmt.Errorf("ent: MessageTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mtgb *MessageTypeGroupBy) IntX(ctx context.Context) int {
	v, err := mtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MessageTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MessageTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mtgb *MessageTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MessageTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagetype.Label}
	default:
		err = fmt.Errorf("ent: MessageTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mtgb *MessageTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MessageTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MessageTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mtgb *MessageTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mtgb *MessageTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagetype.Label}
	default:
		err = fmt.Errorf("ent: MessageTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mtgb *MessageTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := mtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mtgb *MessageTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mtgb.fields {
		if !messagetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mtgb *MessageTypeGroupBy) sqlQuery() *sql.Selector {
	selector := mtgb.sql
	columns := make([]string, 0, len(mtgb.fields)+len(mtgb.fns))
	columns = append(columns, mtgb.fields...)
	for _, fn := range mtgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(mtgb.fields...)
}

// MessageTypeSelect is the builder for selecting fields of MessageType entities.
type MessageTypeSelect struct {
	*MessageTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mts *MessageTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mts.prepareQuery(ctx); err != nil {
		return err
	}
	mts.sql = mts.MessageTypeQuery.sqlQuery(ctx)
	return mts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mts *MessageTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mts *MessageTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MessageTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mts *MessageTypeSelect) StringsX(ctx context.Context) []string {
	v, err := mts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mts *MessageTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagetype.Label}
	default:
		err = fmt.Errorf("ent: MessageTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mts *MessageTypeSelect) StringX(ctx context.Context) string {
	v, err := mts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mts *MessageTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MessageTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mts *MessageTypeSelect) IntsX(ctx context.Context) []int {
	v, err := mts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mts *MessageTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagetype.Label}
	default:
		err = fmt.Errorf("ent: MessageTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mts *MessageTypeSelect) IntX(ctx context.Context) int {
	v, err := mts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mts *MessageTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MessageTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mts *MessageTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mts *MessageTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagetype.Label}
	default:
		err = fmt.Errorf("ent: MessageTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mts *MessageTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := mts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mts *MessageTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MessageTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mts *MessageTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := mts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mts *MessageTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagetype.Label}
	default:
		err = fmt.Errorf("ent: MessageTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mts *MessageTypeSelect) BoolX(ctx context.Context) bool {
	v, err := mts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mts *MessageTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mts.sqlQuery().Query()
	if err := mts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mts *MessageTypeSelect) sqlQuery() sql.Querier {
	selector := mts.sql
	selector.Select(selector.Columns(mts.fields...)...)
	return selector
}
