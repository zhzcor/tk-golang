// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"tkserver/internal/store/ent/kcclass"
	"tkserver/internal/store/ent/kccourse"
	"tkserver/internal/store/ent/major"
	"tkserver/internal/store/ent/majordetail"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/teacher"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MajorQuery is the builder for querying Major entities.
type MajorQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Major
	// eager-loading edges.
	withTeachers    *TeacherQuery
	withMajorDetail *MajorDetailQuery
	withKcClasses   *KcClassQuery
	withCourses     *KcCourseQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MajorQuery builder.
func (mq *MajorQuery) Where(ps ...predicate.Major) *MajorQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MajorQuery) Limit(limit int) *MajorQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MajorQuery) Offset(offset int) *MajorQuery {
	mq.offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MajorQuery) Unique(unique bool) *MajorQuery {
	mq.unique = &unique
	return mq
}

// Order adds an order step to the query.
func (mq *MajorQuery) Order(o ...OrderFunc) *MajorQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryTeachers chains the current query on the "teachers" edge.
func (mq *MajorQuery) QueryTeachers() *TeacherQuery {
	query := &TeacherQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(major.Table, major.FieldID, selector),
			sqlgraph.To(teacher.Table, teacher.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, major.TeachersTable, major.TeachersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMajorDetail chains the current query on the "major_detail" edge.
func (mq *MajorQuery) QueryMajorDetail() *MajorDetailQuery {
	query := &MajorDetailQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(major.Table, major.FieldID, selector),
			sqlgraph.To(majordetail.Table, majordetail.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, major.MajorDetailTable, major.MajorDetailColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryKcClasses chains the current query on the "kc_classes" edge.
func (mq *MajorQuery) QueryKcClasses() *KcClassQuery {
	query := &KcClassQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(major.Table, major.FieldID, selector),
			sqlgraph.To(kcclass.Table, kcclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, major.KcClassesTable, major.KcClassesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCourses chains the current query on the "courses" edge.
func (mq *MajorQuery) QueryCourses() *KcCourseQuery {
	query := &KcCourseQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(major.Table, major.FieldID, selector),
			sqlgraph.To(kccourse.Table, kccourse.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, major.CoursesTable, major.CoursesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Major entity from the query.
// Returns a *NotFoundError when no Major was found.
func (mq *MajorQuery) First(ctx context.Context) (*Major, error) {
	nodes, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{major.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MajorQuery) FirstX(ctx context.Context) *Major {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Major ID from the query.
// Returns a *NotFoundError when no Major ID was found.
func (mq *MajorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{major.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MajorQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Major entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Major entity is not found.
// Returns a *NotFoundError when no Major entities are found.
func (mq *MajorQuery) Only(ctx context.Context) (*Major, error) {
	nodes, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{major.Label}
	default:
		return nil, &NotSingularError{major.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MajorQuery) OnlyX(ctx context.Context) *Major {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Major ID in the query.
// Returns a *NotSingularError when exactly one Major ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mq *MajorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{major.Label}
	default:
		err = &NotSingularError{major.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MajorQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Majors.
func (mq *MajorQuery) All(ctx context.Context) ([]*Major, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MajorQuery) AllX(ctx context.Context) []*Major {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Major IDs.
func (mq *MajorQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(major.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MajorQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MajorQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MajorQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MajorQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MajorQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MajorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MajorQuery) Clone() *MajorQuery {
	if mq == nil {
		return nil
	}
	return &MajorQuery{
		config:          mq.config,
		limit:           mq.limit,
		offset:          mq.offset,
		order:           append([]OrderFunc{}, mq.order...),
		predicates:      append([]predicate.Major{}, mq.predicates...),
		withTeachers:    mq.withTeachers.Clone(),
		withMajorDetail: mq.withMajorDetail.Clone(),
		withKcClasses:   mq.withKcClasses.Clone(),
		withCourses:     mq.withCourses.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithTeachers tells the query-builder to eager-load the nodes that are connected to
// the "teachers" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MajorQuery) WithTeachers(opts ...func(*TeacherQuery)) *MajorQuery {
	query := &TeacherQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withTeachers = query
	return mq
}

// WithMajorDetail tells the query-builder to eager-load the nodes that are connected to
// the "major_detail" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MajorQuery) WithMajorDetail(opts ...func(*MajorDetailQuery)) *MajorQuery {
	query := &MajorDetailQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMajorDetail = query
	return mq
}

// WithKcClasses tells the query-builder to eager-load the nodes that are connected to
// the "kc_classes" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MajorQuery) WithKcClasses(opts ...func(*KcClassQuery)) *MajorQuery {
	query := &KcClassQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withKcClasses = query
	return mq
}

// WithCourses tells the query-builder to eager-load the nodes that are connected to
// the "courses" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MajorQuery) WithCourses(opts ...func(*KcCourseQuery)) *MajorQuery {
	query := &KcCourseQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withCourses = query
	return mq
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
//	client.Major.Query().
//		GroupBy(major.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mq *MajorQuery) GroupBy(field string, fields ...string) *MajorGroupBy {
	group := &MajorGroupBy{config: mq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(ctx), nil
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
//	client.Major.Query().
//		Select(major.FieldUUID).
//		Scan(ctx, &v)
//
func (mq *MajorQuery) Select(field string, fields ...string) *MajorSelect {
	mq.fields = append([]string{field}, fields...)
	return &MajorSelect{MajorQuery: mq}
}

func (mq *MajorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mq.fields {
		if !major.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MajorQuery) sqlAll(ctx context.Context) ([]*Major, error) {
	var (
		nodes       = []*Major{}
		_spec       = mq.querySpec()
		loadedTypes = [4]bool{
			mq.withTeachers != nil,
			mq.withMajorDetail != nil,
			mq.withKcClasses != nil,
			mq.withCourses != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Major{config: mq.config}
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
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mq.withTeachers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Major, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Teachers = []*Teacher{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Major)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   major.TeachersTable,
				Columns: major.TeachersPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(major.TeachersPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, mq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "teachers": %w`, err)
		}
		query.Where(teacher.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "teachers" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Teachers = append(nodes[i].Edges.Teachers, n)
			}
		}
	}

	if query := mq.withMajorDetail; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Major)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.Where(predicate.MajorDetail(func(s *sql.Selector) {
			s.Where(sql.InValues(major.MajorDetailColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.MajorID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "major_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.MajorDetail = n
		}
	}

	if query := mq.withKcClasses; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Major, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.KcClasses = []*KcClass{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Major)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   major.KcClassesTable,
				Columns: major.KcClassesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(major.KcClassesPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, mq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "kc_classes": %w`, err)
		}
		query.Where(kcclass.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "kc_classes" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.KcClasses = append(nodes[i].Edges.KcClasses, n)
			}
		}
	}

	if query := mq.withCourses; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Major, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Courses = []*KcCourse{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Major)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   major.CoursesTable,
				Columns: major.CoursesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(major.CoursesPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, mq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "courses": %w`, err)
		}
		query.Where(kccourse.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "courses" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Courses = append(nodes[i].Edges.Courses, n)
			}
		}
	}

	return nodes, nil
}

func (mq *MajorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MajorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mq *MajorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   major.Table,
			Columns: major.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: major.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if unique := mq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, major.FieldID)
		for i := range fields {
			if fields[i] != major.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MajorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(major.Table)
	selector := builder.Select(t1.Columns(major.Columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(major.Columns...)...)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MajorGroupBy is the group-by builder for Major entities.
type MajorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MajorGroupBy) Aggregate(fns ...AggregateFunc) *MajorGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mgb *MajorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mgb *MajorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MajorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MajorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mgb *MajorGroupBy) StringsX(ctx context.Context) []string {
	v, err := mgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MajorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{major.Label}
	default:
		err = fmt.Errorf("ent: MajorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mgb *MajorGroupBy) StringX(ctx context.Context) string {
	v, err := mgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MajorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MajorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mgb *MajorGroupBy) IntsX(ctx context.Context) []int {
	v, err := mgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MajorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{major.Label}
	default:
		err = fmt.Errorf("ent: MajorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mgb *MajorGroupBy) IntX(ctx context.Context) int {
	v, err := mgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MajorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MajorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mgb *MajorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MajorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{major.Label}
	default:
		err = fmt.Errorf("ent: MajorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mgb *MajorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MajorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MajorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mgb *MajorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MajorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{major.Label}
	default:
		err = fmt.Errorf("ent: MajorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mgb *MajorGroupBy) BoolX(ctx context.Context) bool {
	v, err := mgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mgb *MajorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mgb.fields {
		if !major.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MajorGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql
	columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
	columns = append(columns, mgb.fields...)
	for _, fn := range mgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(mgb.fields...)
}

// MajorSelect is the builder for selecting fields of Major entities.
type MajorSelect struct {
	*MajorQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MajorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	ms.sql = ms.MajorQuery.sqlQuery(ctx)
	return ms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ms *MajorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ms *MajorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MajorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ms *MajorSelect) StringsX(ctx context.Context) []string {
	v, err := ms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ms *MajorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{major.Label}
	default:
		err = fmt.Errorf("ent: MajorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ms *MajorSelect) StringX(ctx context.Context) string {
	v, err := ms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ms *MajorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MajorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ms *MajorSelect) IntsX(ctx context.Context) []int {
	v, err := ms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ms *MajorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{major.Label}
	default:
		err = fmt.Errorf("ent: MajorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ms *MajorSelect) IntX(ctx context.Context) int {
	v, err := ms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ms *MajorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MajorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ms *MajorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ms *MajorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{major.Label}
	default:
		err = fmt.Errorf("ent: MajorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ms *MajorSelect) Float64X(ctx context.Context) float64 {
	v, err := ms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ms *MajorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MajorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ms *MajorSelect) BoolsX(ctx context.Context) []bool {
	v, err := ms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ms *MajorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{major.Label}
	default:
		err = fmt.Errorf("ent: MajorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ms *MajorSelect) BoolX(ctx context.Context) bool {
	v, err := ms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ms *MajorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ms.sqlQuery().Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ms *MajorSelect) sqlQuery() sql.Querier {
	selector := ms.sql
	selector.Select(selector.Columns(ms.fields...)...)
	return selector
}
