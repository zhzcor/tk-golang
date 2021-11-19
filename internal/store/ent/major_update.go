// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gserver/internal/store/ent/kcclass"
	"gserver/internal/store/ent/kccourse"
	"gserver/internal/store/ent/major"
	"gserver/internal/store/ent/majordetail"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/teacher"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MajorUpdate is the builder for updating Major entities.
type MajorUpdate struct {
	config
	hooks    []Hook
	mutation *MajorMutation
}

// Where adds a new predicate for the MajorUpdate builder.
func (mu *MajorUpdate) Where(ps ...predicate.Major) *MajorUpdate {
	mu.mutation.predicates = append(mu.mutation.predicates, ps...)
	return mu
}

// SetUUID sets the "uuid" field.
func (mu *MajorUpdate) SetUUID(s string) *MajorUpdate {
	mu.mutation.SetUUID(s)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MajorUpdate) SetUpdatedAt(t time.Time) *MajorUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mu *MajorUpdate) ClearUpdatedAt() *MajorUpdate {
	mu.mutation.ClearUpdatedAt()
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *MajorUpdate) SetDeletedAt(t time.Time) *MajorUpdate {
	mu.mutation.SetDeletedAt(t)
	return mu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mu *MajorUpdate) SetNillableDeletedAt(t *time.Time) *MajorUpdate {
	if t != nil {
		mu.SetDeletedAt(*t)
	}
	return mu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (mu *MajorUpdate) ClearDeletedAt() *MajorUpdate {
	mu.mutation.ClearDeletedAt()
	return mu
}

// SetName sets the "name" field.
func (mu *MajorUpdate) SetName(s string) *MajorUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MajorUpdate) SetNillableName(s *string) *MajorUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetStatus sets the "status" field.
func (mu *MajorUpdate) SetStatus(u uint8) *MajorUpdate {
	mu.mutation.ResetStatus()
	mu.mutation.SetStatus(u)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MajorUpdate) SetNillableStatus(u *uint8) *MajorUpdate {
	if u != nil {
		mu.SetStatus(*u)
	}
	return mu
}

// AddStatus adds u to the "status" field.
func (mu *MajorUpdate) AddStatus(u uint8) *MajorUpdate {
	mu.mutation.AddStatus(u)
	return mu
}

// SetCode sets the "code" field.
func (mu *MajorUpdate) SetCode(s string) *MajorUpdate {
	mu.mutation.SetCode(s)
	return mu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (mu *MajorUpdate) SetNillableCode(s *string) *MajorUpdate {
	if s != nil {
		mu.SetCode(*s)
	}
	return mu
}

// SetDesc sets the "desc" field.
func (mu *MajorUpdate) SetDesc(s string) *MajorUpdate {
	mu.mutation.SetDesc(s)
	return mu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (mu *MajorUpdate) SetNillableDesc(s *string) *MajorUpdate {
	if s != nil {
		mu.SetDesc(*s)
	}
	return mu
}

// SetSortOrder sets the "sort_order" field.
func (mu *MajorUpdate) SetSortOrder(i int) *MajorUpdate {
	mu.mutation.ResetSortOrder()
	mu.mutation.SetSortOrder(i)
	return mu
}

// SetNillableSortOrder sets the "sort_order" field if the given value is not nil.
func (mu *MajorUpdate) SetNillableSortOrder(i *int) *MajorUpdate {
	if i != nil {
		mu.SetSortOrder(*i)
	}
	return mu
}

// AddSortOrder adds i to the "sort_order" field.
func (mu *MajorUpdate) AddSortOrder(i int) *MajorUpdate {
	mu.mutation.AddSortOrder(i)
	return mu
}

// AddTeacherIDs adds the "teachers" edge to the Teacher entity by IDs.
func (mu *MajorUpdate) AddTeacherIDs(ids ...int) *MajorUpdate {
	mu.mutation.AddTeacherIDs(ids...)
	return mu
}

// AddTeachers adds the "teachers" edges to the Teacher entity.
func (mu *MajorUpdate) AddTeachers(t ...*Teacher) *MajorUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.AddTeacherIDs(ids...)
}

// SetMajorDetailID sets the "major_detail" edge to the MajorDetail entity by ID.
func (mu *MajorUpdate) SetMajorDetailID(id int) *MajorUpdate {
	mu.mutation.SetMajorDetailID(id)
	return mu
}

// SetNillableMajorDetailID sets the "major_detail" edge to the MajorDetail entity by ID if the given value is not nil.
func (mu *MajorUpdate) SetNillableMajorDetailID(id *int) *MajorUpdate {
	if id != nil {
		mu = mu.SetMajorDetailID(*id)
	}
	return mu
}

// SetMajorDetail sets the "major_detail" edge to the MajorDetail entity.
func (mu *MajorUpdate) SetMajorDetail(m *MajorDetail) *MajorUpdate {
	return mu.SetMajorDetailID(m.ID)
}

// AddKcClassIDs adds the "kc_classes" edge to the KcClass entity by IDs.
func (mu *MajorUpdate) AddKcClassIDs(ids ...int) *MajorUpdate {
	mu.mutation.AddKcClassIDs(ids...)
	return mu
}

// AddKcClasses adds the "kc_classes" edges to the KcClass entity.
func (mu *MajorUpdate) AddKcClasses(k ...*KcClass) *MajorUpdate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return mu.AddKcClassIDs(ids...)
}

// AddCourseIDs adds the "courses" edge to the KcCourse entity by IDs.
func (mu *MajorUpdate) AddCourseIDs(ids ...int) *MajorUpdate {
	mu.mutation.AddCourseIDs(ids...)
	return mu
}

// AddCourses adds the "courses" edges to the KcCourse entity.
func (mu *MajorUpdate) AddCourses(k ...*KcCourse) *MajorUpdate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return mu.AddCourseIDs(ids...)
}

// Mutation returns the MajorMutation object of the builder.
func (mu *MajorUpdate) Mutation() *MajorMutation {
	return mu.mutation
}

// ClearTeachers clears all "teachers" edges to the Teacher entity.
func (mu *MajorUpdate) ClearTeachers() *MajorUpdate {
	mu.mutation.ClearTeachers()
	return mu
}

// RemoveTeacherIDs removes the "teachers" edge to Teacher entities by IDs.
func (mu *MajorUpdate) RemoveTeacherIDs(ids ...int) *MajorUpdate {
	mu.mutation.RemoveTeacherIDs(ids...)
	return mu
}

// RemoveTeachers removes "teachers" edges to Teacher entities.
func (mu *MajorUpdate) RemoveTeachers(t ...*Teacher) *MajorUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.RemoveTeacherIDs(ids...)
}

// ClearMajorDetail clears the "major_detail" edge to the MajorDetail entity.
func (mu *MajorUpdate) ClearMajorDetail() *MajorUpdate {
	mu.mutation.ClearMajorDetail()
	return mu
}

// ClearKcClasses clears all "kc_classes" edges to the KcClass entity.
func (mu *MajorUpdate) ClearKcClasses() *MajorUpdate {
	mu.mutation.ClearKcClasses()
	return mu
}

// RemoveKcClassIDs removes the "kc_classes" edge to KcClass entities by IDs.
func (mu *MajorUpdate) RemoveKcClassIDs(ids ...int) *MajorUpdate {
	mu.mutation.RemoveKcClassIDs(ids...)
	return mu
}

// RemoveKcClasses removes "kc_classes" edges to KcClass entities.
func (mu *MajorUpdate) RemoveKcClasses(k ...*KcClass) *MajorUpdate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return mu.RemoveKcClassIDs(ids...)
}

// ClearCourses clears all "courses" edges to the KcCourse entity.
func (mu *MajorUpdate) ClearCourses() *MajorUpdate {
	mu.mutation.ClearCourses()
	return mu
}

// RemoveCourseIDs removes the "courses" edge to KcCourse entities by IDs.
func (mu *MajorUpdate) RemoveCourseIDs(ids ...int) *MajorUpdate {
	mu.mutation.RemoveCourseIDs(ids...)
	return mu
}

// RemoveCourses removes "courses" edges to KcCourse entities.
func (mu *MajorUpdate) RemoveCourses(k ...*KcCourse) *MajorUpdate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return mu.RemoveCourseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MajorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MajorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MajorUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MajorUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MajorUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MajorUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok && !mu.mutation.UpdatedAtCleared() {
		v := major.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

func (mu *MajorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   major.Table,
			Columns: major.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: major.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: major.FieldUUID,
		})
	}
	if mu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: major.FieldCreatedAt,
		})
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: major.FieldUpdatedAt,
		})
	}
	if mu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: major.FieldUpdatedAt,
		})
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: major.FieldDeletedAt,
		})
	}
	if mu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: major.FieldDeletedAt,
		})
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: major.FieldName,
		})
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: major.FieldStatus,
		})
	}
	if value, ok := mu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: major.FieldStatus,
		})
	}
	if value, ok := mu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: major.FieldCode,
		})
	}
	if value, ok := mu.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: major.FieldDesc,
		})
	}
	if value, ok := mu.mutation.SortOrder(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: major.FieldSortOrder,
		})
	}
	if value, ok := mu.mutation.AddedSortOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: major.FieldSortOrder,
		})
	}
	if mu.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   major.TeachersTable,
			Columns: major.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedTeachersIDs(); len(nodes) > 0 && !mu.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   major.TeachersTable,
			Columns: major.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.TeachersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   major.TeachersTable,
			Columns: major.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MajorDetailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   major.MajorDetailTable,
			Columns: []string{major.MajorDetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: majordetail.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MajorDetailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   major.MajorDetailTable,
			Columns: []string{major.MajorDetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: majordetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.KcClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.KcClassesTable,
			Columns: major.KcClassesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kcclass.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedKcClassesIDs(); len(nodes) > 0 && !mu.mutation.KcClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.KcClassesTable,
			Columns: major.KcClassesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kcclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.KcClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.KcClassesTable,
			Columns: major.KcClassesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kcclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.CoursesTable,
			Columns: major.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccourse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedCoursesIDs(); len(nodes) > 0 && !mu.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.CoursesTable,
			Columns: major.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.CoursesTable,
			Columns: major.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{major.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MajorUpdateOne is the builder for updating a single Major entity.
type MajorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MajorMutation
}

// SetUUID sets the "uuid" field.
func (muo *MajorUpdateOne) SetUUID(s string) *MajorUpdateOne {
	muo.mutation.SetUUID(s)
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MajorUpdateOne) SetUpdatedAt(t time.Time) *MajorUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (muo *MajorUpdateOne) ClearUpdatedAt() *MajorUpdateOne {
	muo.mutation.ClearUpdatedAt()
	return muo
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *MajorUpdateOne) SetDeletedAt(t time.Time) *MajorUpdateOne {
	muo.mutation.SetDeletedAt(t)
	return muo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (muo *MajorUpdateOne) SetNillableDeletedAt(t *time.Time) *MajorUpdateOne {
	if t != nil {
		muo.SetDeletedAt(*t)
	}
	return muo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (muo *MajorUpdateOne) ClearDeletedAt() *MajorUpdateOne {
	muo.mutation.ClearDeletedAt()
	return muo
}

// SetName sets the "name" field.
func (muo *MajorUpdateOne) SetName(s string) *MajorUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MajorUpdateOne) SetNillableName(s *string) *MajorUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetStatus sets the "status" field.
func (muo *MajorUpdateOne) SetStatus(u uint8) *MajorUpdateOne {
	muo.mutation.ResetStatus()
	muo.mutation.SetStatus(u)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MajorUpdateOne) SetNillableStatus(u *uint8) *MajorUpdateOne {
	if u != nil {
		muo.SetStatus(*u)
	}
	return muo
}

// AddStatus adds u to the "status" field.
func (muo *MajorUpdateOne) AddStatus(u uint8) *MajorUpdateOne {
	muo.mutation.AddStatus(u)
	return muo
}

// SetCode sets the "code" field.
func (muo *MajorUpdateOne) SetCode(s string) *MajorUpdateOne {
	muo.mutation.SetCode(s)
	return muo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (muo *MajorUpdateOne) SetNillableCode(s *string) *MajorUpdateOne {
	if s != nil {
		muo.SetCode(*s)
	}
	return muo
}

// SetDesc sets the "desc" field.
func (muo *MajorUpdateOne) SetDesc(s string) *MajorUpdateOne {
	muo.mutation.SetDesc(s)
	return muo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (muo *MajorUpdateOne) SetNillableDesc(s *string) *MajorUpdateOne {
	if s != nil {
		muo.SetDesc(*s)
	}
	return muo
}

// SetSortOrder sets the "sort_order" field.
func (muo *MajorUpdateOne) SetSortOrder(i int) *MajorUpdateOne {
	muo.mutation.ResetSortOrder()
	muo.mutation.SetSortOrder(i)
	return muo
}

// SetNillableSortOrder sets the "sort_order" field if the given value is not nil.
func (muo *MajorUpdateOne) SetNillableSortOrder(i *int) *MajorUpdateOne {
	if i != nil {
		muo.SetSortOrder(*i)
	}
	return muo
}

// AddSortOrder adds i to the "sort_order" field.
func (muo *MajorUpdateOne) AddSortOrder(i int) *MajorUpdateOne {
	muo.mutation.AddSortOrder(i)
	return muo
}

// AddTeacherIDs adds the "teachers" edge to the Teacher entity by IDs.
func (muo *MajorUpdateOne) AddTeacherIDs(ids ...int) *MajorUpdateOne {
	muo.mutation.AddTeacherIDs(ids...)
	return muo
}

// AddTeachers adds the "teachers" edges to the Teacher entity.
func (muo *MajorUpdateOne) AddTeachers(t ...*Teacher) *MajorUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.AddTeacherIDs(ids...)
}

// SetMajorDetailID sets the "major_detail" edge to the MajorDetail entity by ID.
func (muo *MajorUpdateOne) SetMajorDetailID(id int) *MajorUpdateOne {
	muo.mutation.SetMajorDetailID(id)
	return muo
}

// SetNillableMajorDetailID sets the "major_detail" edge to the MajorDetail entity by ID if the given value is not nil.
func (muo *MajorUpdateOne) SetNillableMajorDetailID(id *int) *MajorUpdateOne {
	if id != nil {
		muo = muo.SetMajorDetailID(*id)
	}
	return muo
}

// SetMajorDetail sets the "major_detail" edge to the MajorDetail entity.
func (muo *MajorUpdateOne) SetMajorDetail(m *MajorDetail) *MajorUpdateOne {
	return muo.SetMajorDetailID(m.ID)
}

// AddKcClassIDs adds the "kc_classes" edge to the KcClass entity by IDs.
func (muo *MajorUpdateOne) AddKcClassIDs(ids ...int) *MajorUpdateOne {
	muo.mutation.AddKcClassIDs(ids...)
	return muo
}

// AddKcClasses adds the "kc_classes" edges to the KcClass entity.
func (muo *MajorUpdateOne) AddKcClasses(k ...*KcClass) *MajorUpdateOne {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return muo.AddKcClassIDs(ids...)
}

// AddCourseIDs adds the "courses" edge to the KcCourse entity by IDs.
func (muo *MajorUpdateOne) AddCourseIDs(ids ...int) *MajorUpdateOne {
	muo.mutation.AddCourseIDs(ids...)
	return muo
}

// AddCourses adds the "courses" edges to the KcCourse entity.
func (muo *MajorUpdateOne) AddCourses(k ...*KcCourse) *MajorUpdateOne {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return muo.AddCourseIDs(ids...)
}

// Mutation returns the MajorMutation object of the builder.
func (muo *MajorUpdateOne) Mutation() *MajorMutation {
	return muo.mutation
}

// ClearTeachers clears all "teachers" edges to the Teacher entity.
func (muo *MajorUpdateOne) ClearTeachers() *MajorUpdateOne {
	muo.mutation.ClearTeachers()
	return muo
}

// RemoveTeacherIDs removes the "teachers" edge to Teacher entities by IDs.
func (muo *MajorUpdateOne) RemoveTeacherIDs(ids ...int) *MajorUpdateOne {
	muo.mutation.RemoveTeacherIDs(ids...)
	return muo
}

// RemoveTeachers removes "teachers" edges to Teacher entities.
func (muo *MajorUpdateOne) RemoveTeachers(t ...*Teacher) *MajorUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.RemoveTeacherIDs(ids...)
}

// ClearMajorDetail clears the "major_detail" edge to the MajorDetail entity.
func (muo *MajorUpdateOne) ClearMajorDetail() *MajorUpdateOne {
	muo.mutation.ClearMajorDetail()
	return muo
}

// ClearKcClasses clears all "kc_classes" edges to the KcClass entity.
func (muo *MajorUpdateOne) ClearKcClasses() *MajorUpdateOne {
	muo.mutation.ClearKcClasses()
	return muo
}

// RemoveKcClassIDs removes the "kc_classes" edge to KcClass entities by IDs.
func (muo *MajorUpdateOne) RemoveKcClassIDs(ids ...int) *MajorUpdateOne {
	muo.mutation.RemoveKcClassIDs(ids...)
	return muo
}

// RemoveKcClasses removes "kc_classes" edges to KcClass entities.
func (muo *MajorUpdateOne) RemoveKcClasses(k ...*KcClass) *MajorUpdateOne {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return muo.RemoveKcClassIDs(ids...)
}

// ClearCourses clears all "courses" edges to the KcCourse entity.
func (muo *MajorUpdateOne) ClearCourses() *MajorUpdateOne {
	muo.mutation.ClearCourses()
	return muo
}

// RemoveCourseIDs removes the "courses" edge to KcCourse entities by IDs.
func (muo *MajorUpdateOne) RemoveCourseIDs(ids ...int) *MajorUpdateOne {
	muo.mutation.RemoveCourseIDs(ids...)
	return muo
}

// RemoveCourses removes "courses" edges to KcCourse entities.
func (muo *MajorUpdateOne) RemoveCourses(k ...*KcCourse) *MajorUpdateOne {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return muo.RemoveCourseIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MajorUpdateOne) Select(field string, fields ...string) *MajorUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Major entity.
func (muo *MajorUpdateOne) Save(ctx context.Context) (*Major, error) {
	var (
		err  error
		node *Major
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MajorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MajorUpdateOne) SaveX(ctx context.Context) *Major {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MajorUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MajorUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MajorUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok && !muo.mutation.UpdatedAtCleared() {
		v := major.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

func (muo *MajorUpdateOne) sqlSave(ctx context.Context) (_node *Major, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   major.Table,
			Columns: major.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: major.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Major.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, major.FieldID)
		for _, f := range fields {
			if !major.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != major.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: major.FieldUUID,
		})
	}
	if muo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: major.FieldCreatedAt,
		})
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: major.FieldUpdatedAt,
		})
	}
	if muo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: major.FieldUpdatedAt,
		})
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: major.FieldDeletedAt,
		})
	}
	if muo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: major.FieldDeletedAt,
		})
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: major.FieldName,
		})
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: major.FieldStatus,
		})
	}
	if value, ok := muo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: major.FieldStatus,
		})
	}
	if value, ok := muo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: major.FieldCode,
		})
	}
	if value, ok := muo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: major.FieldDesc,
		})
	}
	if value, ok := muo.mutation.SortOrder(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: major.FieldSortOrder,
		})
	}
	if value, ok := muo.mutation.AddedSortOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: major.FieldSortOrder,
		})
	}
	if muo.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   major.TeachersTable,
			Columns: major.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedTeachersIDs(); len(nodes) > 0 && !muo.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   major.TeachersTable,
			Columns: major.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.TeachersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   major.TeachersTable,
			Columns: major.TeachersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MajorDetailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   major.MajorDetailTable,
			Columns: []string{major.MajorDetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: majordetail.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MajorDetailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   major.MajorDetailTable,
			Columns: []string{major.MajorDetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: majordetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.KcClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.KcClassesTable,
			Columns: major.KcClassesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kcclass.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedKcClassesIDs(); len(nodes) > 0 && !muo.mutation.KcClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.KcClassesTable,
			Columns: major.KcClassesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kcclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.KcClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.KcClassesTable,
			Columns: major.KcClassesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kcclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.CoursesTable,
			Columns: major.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccourse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedCoursesIDs(); len(nodes) > 0 && !muo.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.CoursesTable,
			Columns: major.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   major.CoursesTable,
			Columns: major.CoursesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Major{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{major.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
