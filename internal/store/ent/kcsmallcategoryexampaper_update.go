// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"
	"tkserver/internal/store/ent/kccoursesmallcategory"
	"tkserver/internal/store/ent/kcsmallcategoryexampaper"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/tkexampaper"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KcSmallCategoryExamPaperUpdate is the builder for updating KcSmallCategoryExamPaper entities.
type KcSmallCategoryExamPaperUpdate struct {
	config
	hooks    []Hook
	mutation *KcSmallCategoryExamPaperMutation
}

// Where adds a new predicate for the KcSmallCategoryExamPaperUpdate builder.
func (kscepu *KcSmallCategoryExamPaperUpdate) Where(ps ...predicate.KcSmallCategoryExamPaper) *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.predicates = append(kscepu.mutation.predicates, ps...)
	return kscepu
}

// SetUUID sets the "uuid" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetUUID(s string) *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.SetUUID(s)
	return kscepu
}

// SetUpdatedAt sets the "updated_at" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetUpdatedAt(t time.Time) *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.SetUpdatedAt(t)
	return kscepu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) ClearUpdatedAt() *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.ClearUpdatedAt()
	return kscepu
}

// SetDeletedAt sets the "deleted_at" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetDeletedAt(t time.Time) *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.SetDeletedAt(t)
	return kscepu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetNillableDeletedAt(t *time.Time) *KcSmallCategoryExamPaperUpdate {
	if t != nil {
		kscepu.SetDeletedAt(*t)
	}
	return kscepu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) ClearDeletedAt() *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.ClearDeletedAt()
	return kscepu
}

// SetType sets the "type" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetType(u uint8) *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.ResetType()
	kscepu.mutation.SetType(u)
	return kscepu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetNillableType(u *uint8) *KcSmallCategoryExamPaperUpdate {
	if u != nil {
		kscepu.SetType(*u)
	}
	return kscepu
}

// AddType adds u to the "type" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) AddType(u uint8) *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.AddType(u)
	return kscepu
}

// SetExamPaperID sets the "exam_paper_id" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetExamPaperID(i int) *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.ResetExamPaperID()
	kscepu.mutation.SetExamPaperID(i)
	return kscepu
}

// SetNillableExamPaperID sets the "exam_paper_id" field if the given value is not nil.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetNillableExamPaperID(i *int) *KcSmallCategoryExamPaperUpdate {
	if i != nil {
		kscepu.SetExamPaperID(*i)
	}
	return kscepu
}

// ClearExamPaperID clears the value of the "exam_paper_id" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) ClearExamPaperID() *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.ClearExamPaperID()
	return kscepu
}

// SetSmallCategoryID sets the "small_category_id" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetSmallCategoryID(i int) *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.ResetSmallCategoryID()
	kscepu.mutation.SetSmallCategoryID(i)
	return kscepu
}

// SetNillableSmallCategoryID sets the "small_category_id" field if the given value is not nil.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetNillableSmallCategoryID(i *int) *KcSmallCategoryExamPaperUpdate {
	if i != nil {
		kscepu.SetSmallCategoryID(*i)
	}
	return kscepu
}

// ClearSmallCategoryID clears the value of the "small_category_id" field.
func (kscepu *KcSmallCategoryExamPaperUpdate) ClearSmallCategoryID() *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.ClearSmallCategoryID()
	return kscepu
}

// SetExamPaper sets the "exam_paper" edge to the TkExamPaper entity.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetExamPaper(t *TkExamPaper) *KcSmallCategoryExamPaperUpdate {
	return kscepu.SetExamPaperID(t.ID)
}

// SetCourseSmallCategoryID sets the "course_small_category" edge to the KcCourseSmallCategory entity by ID.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetCourseSmallCategoryID(id int) *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.SetCourseSmallCategoryID(id)
	return kscepu
}

// SetNillableCourseSmallCategoryID sets the "course_small_category" edge to the KcCourseSmallCategory entity by ID if the given value is not nil.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetNillableCourseSmallCategoryID(id *int) *KcSmallCategoryExamPaperUpdate {
	if id != nil {
		kscepu = kscepu.SetCourseSmallCategoryID(*id)
	}
	return kscepu
}

// SetCourseSmallCategory sets the "course_small_category" edge to the KcCourseSmallCategory entity.
func (kscepu *KcSmallCategoryExamPaperUpdate) SetCourseSmallCategory(k *KcCourseSmallCategory) *KcSmallCategoryExamPaperUpdate {
	return kscepu.SetCourseSmallCategoryID(k.ID)
}

// Mutation returns the KcSmallCategoryExamPaperMutation object of the builder.
func (kscepu *KcSmallCategoryExamPaperUpdate) Mutation() *KcSmallCategoryExamPaperMutation {
	return kscepu.mutation
}

// ClearExamPaper clears the "exam_paper" edge to the TkExamPaper entity.
func (kscepu *KcSmallCategoryExamPaperUpdate) ClearExamPaper() *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.ClearExamPaper()
	return kscepu
}

// ClearCourseSmallCategory clears the "course_small_category" edge to the KcCourseSmallCategory entity.
func (kscepu *KcSmallCategoryExamPaperUpdate) ClearCourseSmallCategory() *KcSmallCategoryExamPaperUpdate {
	kscepu.mutation.ClearCourseSmallCategory()
	return kscepu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (kscepu *KcSmallCategoryExamPaperUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	kscepu.defaults()
	if len(kscepu.hooks) == 0 {
		affected, err = kscepu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KcSmallCategoryExamPaperMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kscepu.mutation = mutation
			affected, err = kscepu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kscepu.hooks) - 1; i >= 0; i-- {
			mut = kscepu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kscepu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (kscepu *KcSmallCategoryExamPaperUpdate) SaveX(ctx context.Context) int {
	affected, err := kscepu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (kscepu *KcSmallCategoryExamPaperUpdate) Exec(ctx context.Context) error {
	_, err := kscepu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kscepu *KcSmallCategoryExamPaperUpdate) ExecX(ctx context.Context) {
	if err := kscepu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kscepu *KcSmallCategoryExamPaperUpdate) defaults() {
	if _, ok := kscepu.mutation.UpdatedAt(); !ok && !kscepu.mutation.UpdatedAtCleared() {
		v := kcsmallcategoryexampaper.UpdateDefaultUpdatedAt()
		kscepu.mutation.SetUpdatedAt(v)
	}
}

func (kscepu *KcSmallCategoryExamPaperUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kcsmallcategoryexampaper.Table,
			Columns: kcsmallcategoryexampaper.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcsmallcategoryexampaper.FieldID,
			},
		},
	}
	if ps := kscepu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kscepu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldUUID,
		})
	}
	if kscepu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryexampaper.FieldCreatedAt,
		})
	}
	if value, ok := kscepu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldUpdatedAt,
		})
	}
	if kscepu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryexampaper.FieldUpdatedAt,
		})
	}
	if value, ok := kscepu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldDeletedAt,
		})
	}
	if kscepu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryexampaper.FieldDeletedAt,
		})
	}
	if value, ok := kscepu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldType,
		})
	}
	if value, ok := kscepu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldType,
		})
	}
	if kscepu.mutation.ExamPaperCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryexampaper.ExamPaperTable,
			Columns: []string{kcsmallcategoryexampaper.ExamPaperColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkexampaper.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kscepu.mutation.ExamPaperIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryexampaper.ExamPaperTable,
			Columns: []string{kcsmallcategoryexampaper.ExamPaperColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkexampaper.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if kscepu.mutation.CourseSmallCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryexampaper.CourseSmallCategoryTable,
			Columns: []string{kcsmallcategoryexampaper.CourseSmallCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccoursesmallcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kscepu.mutation.CourseSmallCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryexampaper.CourseSmallCategoryTable,
			Columns: []string{kcsmallcategoryexampaper.CourseSmallCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccoursesmallcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, kscepu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kcsmallcategoryexampaper.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// KcSmallCategoryExamPaperUpdateOne is the builder for updating a single KcSmallCategoryExamPaper entity.
type KcSmallCategoryExamPaperUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KcSmallCategoryExamPaperMutation
}

// SetUUID sets the "uuid" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetUUID(s string) *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.SetUUID(s)
	return kscepuo
}

// SetUpdatedAt sets the "updated_at" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetUpdatedAt(t time.Time) *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.SetUpdatedAt(t)
	return kscepuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) ClearUpdatedAt() *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.ClearUpdatedAt()
	return kscepuo
}

// SetDeletedAt sets the "deleted_at" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetDeletedAt(t time.Time) *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.SetDeletedAt(t)
	return kscepuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetNillableDeletedAt(t *time.Time) *KcSmallCategoryExamPaperUpdateOne {
	if t != nil {
		kscepuo.SetDeletedAt(*t)
	}
	return kscepuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) ClearDeletedAt() *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.ClearDeletedAt()
	return kscepuo
}

// SetType sets the "type" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetType(u uint8) *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.ResetType()
	kscepuo.mutation.SetType(u)
	return kscepuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetNillableType(u *uint8) *KcSmallCategoryExamPaperUpdateOne {
	if u != nil {
		kscepuo.SetType(*u)
	}
	return kscepuo
}

// AddType adds u to the "type" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) AddType(u uint8) *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.AddType(u)
	return kscepuo
}

// SetExamPaperID sets the "exam_paper_id" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetExamPaperID(i int) *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.ResetExamPaperID()
	kscepuo.mutation.SetExamPaperID(i)
	return kscepuo
}

// SetNillableExamPaperID sets the "exam_paper_id" field if the given value is not nil.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetNillableExamPaperID(i *int) *KcSmallCategoryExamPaperUpdateOne {
	if i != nil {
		kscepuo.SetExamPaperID(*i)
	}
	return kscepuo
}

// ClearExamPaperID clears the value of the "exam_paper_id" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) ClearExamPaperID() *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.ClearExamPaperID()
	return kscepuo
}

// SetSmallCategoryID sets the "small_category_id" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetSmallCategoryID(i int) *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.ResetSmallCategoryID()
	kscepuo.mutation.SetSmallCategoryID(i)
	return kscepuo
}

// SetNillableSmallCategoryID sets the "small_category_id" field if the given value is not nil.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetNillableSmallCategoryID(i *int) *KcSmallCategoryExamPaperUpdateOne {
	if i != nil {
		kscepuo.SetSmallCategoryID(*i)
	}
	return kscepuo
}

// ClearSmallCategoryID clears the value of the "small_category_id" field.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) ClearSmallCategoryID() *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.ClearSmallCategoryID()
	return kscepuo
}

// SetExamPaper sets the "exam_paper" edge to the TkExamPaper entity.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetExamPaper(t *TkExamPaper) *KcSmallCategoryExamPaperUpdateOne {
	return kscepuo.SetExamPaperID(t.ID)
}

// SetCourseSmallCategoryID sets the "course_small_category" edge to the KcCourseSmallCategory entity by ID.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetCourseSmallCategoryID(id int) *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.SetCourseSmallCategoryID(id)
	return kscepuo
}

// SetNillableCourseSmallCategoryID sets the "course_small_category" edge to the KcCourseSmallCategory entity by ID if the given value is not nil.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetNillableCourseSmallCategoryID(id *int) *KcSmallCategoryExamPaperUpdateOne {
	if id != nil {
		kscepuo = kscepuo.SetCourseSmallCategoryID(*id)
	}
	return kscepuo
}

// SetCourseSmallCategory sets the "course_small_category" edge to the KcCourseSmallCategory entity.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SetCourseSmallCategory(k *KcCourseSmallCategory) *KcSmallCategoryExamPaperUpdateOne {
	return kscepuo.SetCourseSmallCategoryID(k.ID)
}

// Mutation returns the KcSmallCategoryExamPaperMutation object of the builder.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) Mutation() *KcSmallCategoryExamPaperMutation {
	return kscepuo.mutation
}

// ClearExamPaper clears the "exam_paper" edge to the TkExamPaper entity.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) ClearExamPaper() *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.ClearExamPaper()
	return kscepuo
}

// ClearCourseSmallCategory clears the "course_small_category" edge to the KcCourseSmallCategory entity.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) ClearCourseSmallCategory() *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.mutation.ClearCourseSmallCategory()
	return kscepuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) Select(field string, fields ...string) *KcSmallCategoryExamPaperUpdateOne {
	kscepuo.fields = append([]string{field}, fields...)
	return kscepuo
}

// Save executes the query and returns the updated KcSmallCategoryExamPaper entity.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) Save(ctx context.Context) (*KcSmallCategoryExamPaper, error) {
	var (
		err  error
		node *KcSmallCategoryExamPaper
	)
	kscepuo.defaults()
	if len(kscepuo.hooks) == 0 {
		node, err = kscepuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KcSmallCategoryExamPaperMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kscepuo.mutation = mutation
			node, err = kscepuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kscepuo.hooks) - 1; i >= 0; i-- {
			mut = kscepuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kscepuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) SaveX(ctx context.Context) *KcSmallCategoryExamPaper {
	node, err := kscepuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) Exec(ctx context.Context) error {
	_, err := kscepuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) ExecX(ctx context.Context) {
	if err := kscepuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kscepuo *KcSmallCategoryExamPaperUpdateOne) defaults() {
	if _, ok := kscepuo.mutation.UpdatedAt(); !ok && !kscepuo.mutation.UpdatedAtCleared() {
		v := kcsmallcategoryexampaper.UpdateDefaultUpdatedAt()
		kscepuo.mutation.SetUpdatedAt(v)
	}
}

func (kscepuo *KcSmallCategoryExamPaperUpdateOne) sqlSave(ctx context.Context) (_node *KcSmallCategoryExamPaper, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kcsmallcategoryexampaper.Table,
			Columns: kcsmallcategoryexampaper.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcsmallcategoryexampaper.FieldID,
			},
		},
	}
	id, ok := kscepuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing KcSmallCategoryExamPaper.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := kscepuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kcsmallcategoryexampaper.FieldID)
		for _, f := range fields {
			if !kcsmallcategoryexampaper.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != kcsmallcategoryexampaper.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kscepuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kscepuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldUUID,
		})
	}
	if kscepuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryexampaper.FieldCreatedAt,
		})
	}
	if value, ok := kscepuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldUpdatedAt,
		})
	}
	if kscepuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryexampaper.FieldUpdatedAt,
		})
	}
	if value, ok := kscepuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldDeletedAt,
		})
	}
	if kscepuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryexampaper.FieldDeletedAt,
		})
	}
	if value, ok := kscepuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldType,
		})
	}
	if value, ok := kscepuo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldType,
		})
	}
	if kscepuo.mutation.ExamPaperCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryexampaper.ExamPaperTable,
			Columns: []string{kcsmallcategoryexampaper.ExamPaperColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkexampaper.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kscepuo.mutation.ExamPaperIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryexampaper.ExamPaperTable,
			Columns: []string{kcsmallcategoryexampaper.ExamPaperColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkexampaper.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if kscepuo.mutation.CourseSmallCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryexampaper.CourseSmallCategoryTable,
			Columns: []string{kcsmallcategoryexampaper.CourseSmallCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccoursesmallcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kscepuo.mutation.CourseSmallCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryexampaper.CourseSmallCategoryTable,
			Columns: []string{kcsmallcategoryexampaper.CourseSmallCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccoursesmallcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &KcSmallCategoryExamPaper{config: kscepuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kscepuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kcsmallcategoryexampaper.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
