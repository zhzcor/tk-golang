// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"tkserver/internal/store/ent/attachment"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/useraskanswer"
	"tkserver/internal/store/ent/useraskanswerattachment"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserAskAnswerAttachmentQuery is the builder for querying UserAskAnswerAttachment entities.
type UserAskAnswerAttachmentQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserAskAnswerAttachment
	// eager-loading edges.
	withAttachment *AttachmentQuery
	withAsk        *UserAskAnswerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserAskAnswerAttachmentQuery builder.
func (uaaaq *UserAskAnswerAttachmentQuery) Where(ps ...predicate.UserAskAnswerAttachment) *UserAskAnswerAttachmentQuery {
	uaaaq.predicates = append(uaaaq.predicates, ps...)
	return uaaaq
}

// Limit adds a limit step to the query.
func (uaaaq *UserAskAnswerAttachmentQuery) Limit(limit int) *UserAskAnswerAttachmentQuery {
	uaaaq.limit = &limit
	return uaaaq
}

// Offset adds an offset step to the query.
func (uaaaq *UserAskAnswerAttachmentQuery) Offset(offset int) *UserAskAnswerAttachmentQuery {
	uaaaq.offset = &offset
	return uaaaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uaaaq *UserAskAnswerAttachmentQuery) Unique(unique bool) *UserAskAnswerAttachmentQuery {
	uaaaq.unique = &unique
	return uaaaq
}

// Order adds an order step to the query.
func (uaaaq *UserAskAnswerAttachmentQuery) Order(o ...OrderFunc) *UserAskAnswerAttachmentQuery {
	uaaaq.order = append(uaaaq.order, o...)
	return uaaaq
}

// QueryAttachment chains the current query on the "attachment" edge.
func (uaaaq *UserAskAnswerAttachmentQuery) QueryAttachment() *AttachmentQuery {
	query := &AttachmentQuery{config: uaaaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaaaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaaaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraskanswerattachment.Table, useraskanswerattachment.FieldID, selector),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useraskanswerattachment.AttachmentTable, useraskanswerattachment.AttachmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaaaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAsk chains the current query on the "ask" edge.
func (uaaaq *UserAskAnswerAttachmentQuery) QueryAsk() *UserAskAnswerQuery {
	query := &UserAskAnswerQuery{config: uaaaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaaaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaaaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraskanswerattachment.Table, useraskanswerattachment.FieldID, selector),
			sqlgraph.To(useraskanswer.Table, useraskanswer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useraskanswerattachment.AskTable, useraskanswerattachment.AskColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaaaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserAskAnswerAttachment entity from the query.
// Returns a *NotFoundError when no UserAskAnswerAttachment was found.
func (uaaaq *UserAskAnswerAttachmentQuery) First(ctx context.Context) (*UserAskAnswerAttachment, error) {
	nodes, err := uaaaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{useraskanswerattachment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uaaaq *UserAskAnswerAttachmentQuery) FirstX(ctx context.Context) *UserAskAnswerAttachment {
	node, err := uaaaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserAskAnswerAttachment ID from the query.
// Returns a *NotFoundError when no UserAskAnswerAttachment ID was found.
func (uaaaq *UserAskAnswerAttachmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaaaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{useraskanswerattachment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uaaaq *UserAskAnswerAttachmentQuery) FirstIDX(ctx context.Context) int {
	id, err := uaaaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserAskAnswerAttachment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UserAskAnswerAttachment entity is not found.
// Returns a *NotFoundError when no UserAskAnswerAttachment entities are found.
func (uaaaq *UserAskAnswerAttachmentQuery) Only(ctx context.Context) (*UserAskAnswerAttachment, error) {
	nodes, err := uaaaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{useraskanswerattachment.Label}
	default:
		return nil, &NotSingularError{useraskanswerattachment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uaaaq *UserAskAnswerAttachmentQuery) OnlyX(ctx context.Context) *UserAskAnswerAttachment {
	node, err := uaaaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserAskAnswerAttachment ID in the query.
// Returns a *NotSingularError when exactly one UserAskAnswerAttachment ID is not found.
// Returns a *NotFoundError when no entities are found.
func (uaaaq *UserAskAnswerAttachmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaaaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{useraskanswerattachment.Label}
	default:
		err = &NotSingularError{useraskanswerattachment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uaaaq *UserAskAnswerAttachmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := uaaaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserAskAnswerAttachments.
func (uaaaq *UserAskAnswerAttachmentQuery) All(ctx context.Context) ([]*UserAskAnswerAttachment, error) {
	if err := uaaaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uaaaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uaaaq *UserAskAnswerAttachmentQuery) AllX(ctx context.Context) []*UserAskAnswerAttachment {
	nodes, err := uaaaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserAskAnswerAttachment IDs.
func (uaaaq *UserAskAnswerAttachmentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uaaaq.Select(useraskanswerattachment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uaaaq *UserAskAnswerAttachmentQuery) IDsX(ctx context.Context) []int {
	ids, err := uaaaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uaaaq *UserAskAnswerAttachmentQuery) Count(ctx context.Context) (int, error) {
	if err := uaaaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uaaaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uaaaq *UserAskAnswerAttachmentQuery) CountX(ctx context.Context) int {
	count, err := uaaaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uaaaq *UserAskAnswerAttachmentQuery) Exist(ctx context.Context) (bool, error) {
	if err := uaaaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uaaaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uaaaq *UserAskAnswerAttachmentQuery) ExistX(ctx context.Context) bool {
	exist, err := uaaaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserAskAnswerAttachmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uaaaq *UserAskAnswerAttachmentQuery) Clone() *UserAskAnswerAttachmentQuery {
	if uaaaq == nil {
		return nil
	}
	return &UserAskAnswerAttachmentQuery{
		config:         uaaaq.config,
		limit:          uaaaq.limit,
		offset:         uaaaq.offset,
		order:          append([]OrderFunc{}, uaaaq.order...),
		predicates:     append([]predicate.UserAskAnswerAttachment{}, uaaaq.predicates...),
		withAttachment: uaaaq.withAttachment.Clone(),
		withAsk:        uaaaq.withAsk.Clone(),
		// clone intermediate query.
		sql:  uaaaq.sql.Clone(),
		path: uaaaq.path,
	}
}

// WithAttachment tells the query-builder to eager-load the nodes that are connected to
// the "attachment" edge. The optional arguments are used to configure the query builder of the edge.
func (uaaaq *UserAskAnswerAttachmentQuery) WithAttachment(opts ...func(*AttachmentQuery)) *UserAskAnswerAttachmentQuery {
	query := &AttachmentQuery{config: uaaaq.config}
	for _, opt := range opts {
		opt(query)
	}
	uaaaq.withAttachment = query
	return uaaaq
}

// WithAsk tells the query-builder to eager-load the nodes that are connected to
// the "ask" edge. The optional arguments are used to configure the query builder of the edge.
func (uaaaq *UserAskAnswerAttachmentQuery) WithAsk(opts ...func(*UserAskAnswerQuery)) *UserAskAnswerAttachmentQuery {
	query := &UserAskAnswerQuery{config: uaaaq.config}
	for _, opt := range opts {
		opt(query)
	}
	uaaaq.withAsk = query
	return uaaaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AttachmentID int `json:"attachment_id"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserAskAnswerAttachment.Query().
//		GroupBy(useraskanswerattachment.FieldAttachmentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uaaaq *UserAskAnswerAttachmentQuery) GroupBy(field string, fields ...string) *UserAskAnswerAttachmentGroupBy {
	group := &UserAskAnswerAttachmentGroupBy{config: uaaaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uaaaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uaaaq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AttachmentID int `json:"attachment_id"`
//	}
//
//	client.UserAskAnswerAttachment.Query().
//		Select(useraskanswerattachment.FieldAttachmentID).
//		Scan(ctx, &v)
//
func (uaaaq *UserAskAnswerAttachmentQuery) Select(field string, fields ...string) *UserAskAnswerAttachmentSelect {
	uaaaq.fields = append([]string{field}, fields...)
	return &UserAskAnswerAttachmentSelect{UserAskAnswerAttachmentQuery: uaaaq}
}

func (uaaaq *UserAskAnswerAttachmentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uaaaq.fields {
		if !useraskanswerattachment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uaaaq.path != nil {
		prev, err := uaaaq.path(ctx)
		if err != nil {
			return err
		}
		uaaaq.sql = prev
	}
	return nil
}

func (uaaaq *UserAskAnswerAttachmentQuery) sqlAll(ctx context.Context) ([]*UserAskAnswerAttachment, error) {
	var (
		nodes       = []*UserAskAnswerAttachment{}
		_spec       = uaaaq.querySpec()
		loadedTypes = [2]bool{
			uaaaq.withAttachment != nil,
			uaaaq.withAsk != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UserAskAnswerAttachment{config: uaaaq.config}
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
	if err := sqlgraph.QueryNodes(ctx, uaaaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uaaaq.withAttachment; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UserAskAnswerAttachment)
		for i := range nodes {
			fk := nodes[i].AttachmentID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(attachment.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "attachment_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Attachment = n
			}
		}
	}

	if query := uaaaq.withAsk; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UserAskAnswerAttachment)
		for i := range nodes {
			fk := nodes[i].AskID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(useraskanswer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "ask_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Ask = n
			}
		}
	}

	return nodes, nil
}

func (uaaaq *UserAskAnswerAttachmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uaaaq.querySpec()
	return sqlgraph.CountNodes(ctx, uaaaq.driver, _spec)
}

func (uaaaq *UserAskAnswerAttachmentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uaaaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uaaaq *UserAskAnswerAttachmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   useraskanswerattachment.Table,
			Columns: useraskanswerattachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: useraskanswerattachment.FieldID,
			},
		},
		From:   uaaaq.sql,
		Unique: true,
	}
	if unique := uaaaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uaaaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, useraskanswerattachment.FieldID)
		for i := range fields {
			if fields[i] != useraskanswerattachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uaaaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uaaaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uaaaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uaaaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uaaaq *UserAskAnswerAttachmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uaaaq.driver.Dialect())
	t1 := builder.Table(useraskanswerattachment.Table)
	selector := builder.Select(t1.Columns(useraskanswerattachment.Columns...)...).From(t1)
	if uaaaq.sql != nil {
		selector = uaaaq.sql
		selector.Select(selector.Columns(useraskanswerattachment.Columns...)...)
	}
	for _, p := range uaaaq.predicates {
		p(selector)
	}
	for _, p := range uaaaq.order {
		p(selector)
	}
	if offset := uaaaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uaaaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserAskAnswerAttachmentGroupBy is the group-by builder for UserAskAnswerAttachment entities.
type UserAskAnswerAttachmentGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Aggregate(fns ...AggregateFunc) *UserAskAnswerAttachmentGroupBy {
	uaaagb.fns = append(uaaagb.fns, fns...)
	return uaaagb
}

// Scan applies the group-by query and scans the result into the given value.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uaaagb.path(ctx)
	if err != nil {
		return err
	}
	uaaagb.sql = query
	return uaaagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uaaagb *UserAskAnswerAttachmentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uaaagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uaaagb.fields) > 1 {
		return nil, errors.New("ent: UserAskAnswerAttachmentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uaaagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uaaagb *UserAskAnswerAttachmentGroupBy) StringsX(ctx context.Context) []string {
	v, err := uaaagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uaaagb *UserAskAnswerAttachmentGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uaaagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{useraskanswerattachment.Label}
	default:
		err = fmt.Errorf("ent: UserAskAnswerAttachmentGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uaaagb *UserAskAnswerAttachmentGroupBy) StringX(ctx context.Context) string {
	v, err := uaaagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uaaagb.fields) > 1 {
		return nil, errors.New("ent: UserAskAnswerAttachmentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uaaagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uaaagb *UserAskAnswerAttachmentGroupBy) IntsX(ctx context.Context) []int {
	v, err := uaaagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uaaagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{useraskanswerattachment.Label}
	default:
		err = fmt.Errorf("ent: UserAskAnswerAttachmentGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uaaagb *UserAskAnswerAttachmentGroupBy) IntX(ctx context.Context) int {
	v, err := uaaagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uaaagb.fields) > 1 {
		return nil, errors.New("ent: UserAskAnswerAttachmentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uaaagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uaaagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uaaagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{useraskanswerattachment.Label}
	default:
		err = fmt.Errorf("ent: UserAskAnswerAttachmentGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uaaagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uaaagb.fields) > 1 {
		return nil, errors.New("ent: UserAskAnswerAttachmentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uaaagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uaaagb *UserAskAnswerAttachmentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uaaagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uaaagb *UserAskAnswerAttachmentGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uaaagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{useraskanswerattachment.Label}
	default:
		err = fmt.Errorf("ent: UserAskAnswerAttachmentGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uaaagb *UserAskAnswerAttachmentGroupBy) BoolX(ctx context.Context) bool {
	v, err := uaaagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uaaagb *UserAskAnswerAttachmentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uaaagb.fields {
		if !useraskanswerattachment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uaaagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uaaagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uaaagb *UserAskAnswerAttachmentGroupBy) sqlQuery() *sql.Selector {
	selector := uaaagb.sql
	columns := make([]string, 0, len(uaaagb.fields)+len(uaaagb.fns))
	columns = append(columns, uaaagb.fields...)
	for _, fn := range uaaagb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(uaaagb.fields...)
}

// UserAskAnswerAttachmentSelect is the builder for selecting fields of UserAskAnswerAttachment entities.
type UserAskAnswerAttachmentSelect struct {
	*UserAskAnswerAttachmentQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uaaas *UserAskAnswerAttachmentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uaaas.prepareQuery(ctx); err != nil {
		return err
	}
	uaaas.sql = uaaas.UserAskAnswerAttachmentQuery.sqlQuery(ctx)
	return uaaas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uaaas *UserAskAnswerAttachmentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uaaas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (uaaas *UserAskAnswerAttachmentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uaaas.fields) > 1 {
		return nil, errors.New("ent: UserAskAnswerAttachmentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uaaas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uaaas *UserAskAnswerAttachmentSelect) StringsX(ctx context.Context) []string {
	v, err := uaaas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (uaaas *UserAskAnswerAttachmentSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uaaas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{useraskanswerattachment.Label}
	default:
		err = fmt.Errorf("ent: UserAskAnswerAttachmentSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uaaas *UserAskAnswerAttachmentSelect) StringX(ctx context.Context) string {
	v, err := uaaas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (uaaas *UserAskAnswerAttachmentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uaaas.fields) > 1 {
		return nil, errors.New("ent: UserAskAnswerAttachmentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uaaas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uaaas *UserAskAnswerAttachmentSelect) IntsX(ctx context.Context) []int {
	v, err := uaaas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (uaaas *UserAskAnswerAttachmentSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uaaas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{useraskanswerattachment.Label}
	default:
		err = fmt.Errorf("ent: UserAskAnswerAttachmentSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uaaas *UserAskAnswerAttachmentSelect) IntX(ctx context.Context) int {
	v, err := uaaas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (uaaas *UserAskAnswerAttachmentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uaaas.fields) > 1 {
		return nil, errors.New("ent: UserAskAnswerAttachmentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uaaas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uaaas *UserAskAnswerAttachmentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uaaas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (uaaas *UserAskAnswerAttachmentSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uaaas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{useraskanswerattachment.Label}
	default:
		err = fmt.Errorf("ent: UserAskAnswerAttachmentSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uaaas *UserAskAnswerAttachmentSelect) Float64X(ctx context.Context) float64 {
	v, err := uaaas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (uaaas *UserAskAnswerAttachmentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uaaas.fields) > 1 {
		return nil, errors.New("ent: UserAskAnswerAttachmentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uaaas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uaaas *UserAskAnswerAttachmentSelect) BoolsX(ctx context.Context) []bool {
	v, err := uaaas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (uaaas *UserAskAnswerAttachmentSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uaaas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{useraskanswerattachment.Label}
	default:
		err = fmt.Errorf("ent: UserAskAnswerAttachmentSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uaaas *UserAskAnswerAttachmentSelect) BoolX(ctx context.Context) bool {
	v, err := uaaas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uaaas *UserAskAnswerAttachmentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uaaas.sqlQuery().Query()
	if err := uaaas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uaaas *UserAskAnswerAttachmentSelect) sqlQuery() sql.Querier {
	selector := uaaas.sql
	selector.Select(selector.Columns(uaaas.fields...)...)
	return selector
}
