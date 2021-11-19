// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"
	"tkserver/internal/store/ent/attachment"
	"tkserver/internal/store/ent/kccoursesmallcategory"
	"tkserver/internal/store/ent/kcsmallcategoryattachment"
	"tkserver/internal/store/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KcSmallCategoryAttachmentUpdate is the builder for updating KcSmallCategoryAttachment entities.
type KcSmallCategoryAttachmentUpdate struct {
	config
	hooks    []Hook
	mutation *KcSmallCategoryAttachmentMutation
}

// Where adds a new predicate for the KcSmallCategoryAttachmentUpdate builder.
func (kscau *KcSmallCategoryAttachmentUpdate) Where(ps ...predicate.KcSmallCategoryAttachment) *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.predicates = append(kscau.mutation.predicates, ps...)
	return kscau
}

// SetUUID sets the "uuid" field.
func (kscau *KcSmallCategoryAttachmentUpdate) SetUUID(s string) *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.SetUUID(s)
	return kscau
}

// SetUpdatedAt sets the "updated_at" field.
func (kscau *KcSmallCategoryAttachmentUpdate) SetUpdatedAt(t time.Time) *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.SetUpdatedAt(t)
	return kscau
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (kscau *KcSmallCategoryAttachmentUpdate) ClearUpdatedAt() *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.ClearUpdatedAt()
	return kscau
}

// SetDeletedAt sets the "deleted_at" field.
func (kscau *KcSmallCategoryAttachmentUpdate) SetDeletedAt(t time.Time) *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.SetDeletedAt(t)
	return kscau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kscau *KcSmallCategoryAttachmentUpdate) SetNillableDeletedAt(t *time.Time) *KcSmallCategoryAttachmentUpdate {
	if t != nil {
		kscau.SetDeletedAt(*t)
	}
	return kscau
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (kscau *KcSmallCategoryAttachmentUpdate) ClearDeletedAt() *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.ClearDeletedAt()
	return kscau
}

// SetAttachmentID sets the "attachment_id" field.
func (kscau *KcSmallCategoryAttachmentUpdate) SetAttachmentID(i int) *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.ResetAttachmentID()
	kscau.mutation.SetAttachmentID(i)
	return kscau
}

// SetNillableAttachmentID sets the "attachment_id" field if the given value is not nil.
func (kscau *KcSmallCategoryAttachmentUpdate) SetNillableAttachmentID(i *int) *KcSmallCategoryAttachmentUpdate {
	if i != nil {
		kscau.SetAttachmentID(*i)
	}
	return kscau
}

// ClearAttachmentID clears the value of the "attachment_id" field.
func (kscau *KcSmallCategoryAttachmentUpdate) ClearAttachmentID() *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.ClearAttachmentID()
	return kscau
}

// SetAttachmentName sets the "attachment_name" field.
func (kscau *KcSmallCategoryAttachmentUpdate) SetAttachmentName(s string) *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.SetAttachmentName(s)
	return kscau
}

// SetNillableAttachmentName sets the "attachment_name" field if the given value is not nil.
func (kscau *KcSmallCategoryAttachmentUpdate) SetNillableAttachmentName(s *string) *KcSmallCategoryAttachmentUpdate {
	if s != nil {
		kscau.SetAttachmentName(*s)
	}
	return kscau
}

// SetSmallCategoryID sets the "small_category_id" field.
func (kscau *KcSmallCategoryAttachmentUpdate) SetSmallCategoryID(i int) *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.ResetSmallCategoryID()
	kscau.mutation.SetSmallCategoryID(i)
	return kscau
}

// SetNillableSmallCategoryID sets the "small_category_id" field if the given value is not nil.
func (kscau *KcSmallCategoryAttachmentUpdate) SetNillableSmallCategoryID(i *int) *KcSmallCategoryAttachmentUpdate {
	if i != nil {
		kscau.SetSmallCategoryID(*i)
	}
	return kscau
}

// ClearSmallCategoryID clears the value of the "small_category_id" field.
func (kscau *KcSmallCategoryAttachmentUpdate) ClearSmallCategoryID() *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.ClearSmallCategoryID()
	return kscau
}

// SetAttachment sets the "attachment" edge to the Attachment entity.
func (kscau *KcSmallCategoryAttachmentUpdate) SetAttachment(a *Attachment) *KcSmallCategoryAttachmentUpdate {
	return kscau.SetAttachmentID(a.ID)
}

// SetSmallCategory sets the "small_category" edge to the KcCourseSmallCategory entity.
func (kscau *KcSmallCategoryAttachmentUpdate) SetSmallCategory(k *KcCourseSmallCategory) *KcSmallCategoryAttachmentUpdate {
	return kscau.SetSmallCategoryID(k.ID)
}

// Mutation returns the KcSmallCategoryAttachmentMutation object of the builder.
func (kscau *KcSmallCategoryAttachmentUpdate) Mutation() *KcSmallCategoryAttachmentMutation {
	return kscau.mutation
}

// ClearAttachment clears the "attachment" edge to the Attachment entity.
func (kscau *KcSmallCategoryAttachmentUpdate) ClearAttachment() *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.ClearAttachment()
	return kscau
}

// ClearSmallCategory clears the "small_category" edge to the KcCourseSmallCategory entity.
func (kscau *KcSmallCategoryAttachmentUpdate) ClearSmallCategory() *KcSmallCategoryAttachmentUpdate {
	kscau.mutation.ClearSmallCategory()
	return kscau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (kscau *KcSmallCategoryAttachmentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	kscau.defaults()
	if len(kscau.hooks) == 0 {
		affected, err = kscau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KcSmallCategoryAttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kscau.mutation = mutation
			affected, err = kscau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kscau.hooks) - 1; i >= 0; i-- {
			mut = kscau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kscau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (kscau *KcSmallCategoryAttachmentUpdate) SaveX(ctx context.Context) int {
	affected, err := kscau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (kscau *KcSmallCategoryAttachmentUpdate) Exec(ctx context.Context) error {
	_, err := kscau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kscau *KcSmallCategoryAttachmentUpdate) ExecX(ctx context.Context) {
	if err := kscau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kscau *KcSmallCategoryAttachmentUpdate) defaults() {
	if _, ok := kscau.mutation.UpdatedAt(); !ok && !kscau.mutation.UpdatedAtCleared() {
		v := kcsmallcategoryattachment.UpdateDefaultUpdatedAt()
		kscau.mutation.SetUpdatedAt(v)
	}
}

func (kscau *KcSmallCategoryAttachmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kcsmallcategoryattachment.Table,
			Columns: kcsmallcategoryattachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcsmallcategoryattachment.FieldID,
			},
		},
	}
	if ps := kscau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kscau.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldUUID,
		})
	}
	if kscau.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryattachment.FieldCreatedAt,
		})
	}
	if value, ok := kscau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldUpdatedAt,
		})
	}
	if kscau.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryattachment.FieldUpdatedAt,
		})
	}
	if value, ok := kscau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldDeletedAt,
		})
	}
	if kscau.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryattachment.FieldDeletedAt,
		})
	}
	if value, ok := kscau.mutation.AttachmentName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldAttachmentName,
		})
	}
	if kscau.mutation.AttachmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryattachment.AttachmentTable,
			Columns: []string{kcsmallcategoryattachment.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kscau.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryattachment.AttachmentTable,
			Columns: []string{kcsmallcategoryattachment.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if kscau.mutation.SmallCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryattachment.SmallCategoryTable,
			Columns: []string{kcsmallcategoryattachment.SmallCategoryColumn},
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
	if nodes := kscau.mutation.SmallCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryattachment.SmallCategoryTable,
			Columns: []string{kcsmallcategoryattachment.SmallCategoryColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, kscau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kcsmallcategoryattachment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// KcSmallCategoryAttachmentUpdateOne is the builder for updating a single KcSmallCategoryAttachment entity.
type KcSmallCategoryAttachmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KcSmallCategoryAttachmentMutation
}

// SetUUID sets the "uuid" field.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetUUID(s string) *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.SetUUID(s)
	return kscauo
}

// SetUpdatedAt sets the "updated_at" field.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetUpdatedAt(t time.Time) *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.SetUpdatedAt(t)
	return kscauo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) ClearUpdatedAt() *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.ClearUpdatedAt()
	return kscauo
}

// SetDeletedAt sets the "deleted_at" field.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetDeletedAt(t time.Time) *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.SetDeletedAt(t)
	return kscauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetNillableDeletedAt(t *time.Time) *KcSmallCategoryAttachmentUpdateOne {
	if t != nil {
		kscauo.SetDeletedAt(*t)
	}
	return kscauo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) ClearDeletedAt() *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.ClearDeletedAt()
	return kscauo
}

// SetAttachmentID sets the "attachment_id" field.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetAttachmentID(i int) *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.ResetAttachmentID()
	kscauo.mutation.SetAttachmentID(i)
	return kscauo
}

// SetNillableAttachmentID sets the "attachment_id" field if the given value is not nil.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetNillableAttachmentID(i *int) *KcSmallCategoryAttachmentUpdateOne {
	if i != nil {
		kscauo.SetAttachmentID(*i)
	}
	return kscauo
}

// ClearAttachmentID clears the value of the "attachment_id" field.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) ClearAttachmentID() *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.ClearAttachmentID()
	return kscauo
}

// SetAttachmentName sets the "attachment_name" field.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetAttachmentName(s string) *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.SetAttachmentName(s)
	return kscauo
}

// SetNillableAttachmentName sets the "attachment_name" field if the given value is not nil.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetNillableAttachmentName(s *string) *KcSmallCategoryAttachmentUpdateOne {
	if s != nil {
		kscauo.SetAttachmentName(*s)
	}
	return kscauo
}

// SetSmallCategoryID sets the "small_category_id" field.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetSmallCategoryID(i int) *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.ResetSmallCategoryID()
	kscauo.mutation.SetSmallCategoryID(i)
	return kscauo
}

// SetNillableSmallCategoryID sets the "small_category_id" field if the given value is not nil.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetNillableSmallCategoryID(i *int) *KcSmallCategoryAttachmentUpdateOne {
	if i != nil {
		kscauo.SetSmallCategoryID(*i)
	}
	return kscauo
}

// ClearSmallCategoryID clears the value of the "small_category_id" field.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) ClearSmallCategoryID() *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.ClearSmallCategoryID()
	return kscauo
}

// SetAttachment sets the "attachment" edge to the Attachment entity.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetAttachment(a *Attachment) *KcSmallCategoryAttachmentUpdateOne {
	return kscauo.SetAttachmentID(a.ID)
}

// SetSmallCategory sets the "small_category" edge to the KcCourseSmallCategory entity.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SetSmallCategory(k *KcCourseSmallCategory) *KcSmallCategoryAttachmentUpdateOne {
	return kscauo.SetSmallCategoryID(k.ID)
}

// Mutation returns the KcSmallCategoryAttachmentMutation object of the builder.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) Mutation() *KcSmallCategoryAttachmentMutation {
	return kscauo.mutation
}

// ClearAttachment clears the "attachment" edge to the Attachment entity.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) ClearAttachment() *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.ClearAttachment()
	return kscauo
}

// ClearSmallCategory clears the "small_category" edge to the KcCourseSmallCategory entity.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) ClearSmallCategory() *KcSmallCategoryAttachmentUpdateOne {
	kscauo.mutation.ClearSmallCategory()
	return kscauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) Select(field string, fields ...string) *KcSmallCategoryAttachmentUpdateOne {
	kscauo.fields = append([]string{field}, fields...)
	return kscauo
}

// Save executes the query and returns the updated KcSmallCategoryAttachment entity.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) Save(ctx context.Context) (*KcSmallCategoryAttachment, error) {
	var (
		err  error
		node *KcSmallCategoryAttachment
	)
	kscauo.defaults()
	if len(kscauo.hooks) == 0 {
		node, err = kscauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KcSmallCategoryAttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kscauo.mutation = mutation
			node, err = kscauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kscauo.hooks) - 1; i >= 0; i-- {
			mut = kscauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kscauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) SaveX(ctx context.Context) *KcSmallCategoryAttachment {
	node, err := kscauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) Exec(ctx context.Context) error {
	_, err := kscauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) ExecX(ctx context.Context) {
	if err := kscauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kscauo *KcSmallCategoryAttachmentUpdateOne) defaults() {
	if _, ok := kscauo.mutation.UpdatedAt(); !ok && !kscauo.mutation.UpdatedAtCleared() {
		v := kcsmallcategoryattachment.UpdateDefaultUpdatedAt()
		kscauo.mutation.SetUpdatedAt(v)
	}
}

func (kscauo *KcSmallCategoryAttachmentUpdateOne) sqlSave(ctx context.Context) (_node *KcSmallCategoryAttachment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kcsmallcategoryattachment.Table,
			Columns: kcsmallcategoryattachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcsmallcategoryattachment.FieldID,
			},
		},
	}
	id, ok := kscauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing KcSmallCategoryAttachment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := kscauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kcsmallcategoryattachment.FieldID)
		for _, f := range fields {
			if !kcsmallcategoryattachment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != kcsmallcategoryattachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kscauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kscauo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldUUID,
		})
	}
	if kscauo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryattachment.FieldCreatedAt,
		})
	}
	if value, ok := kscauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldUpdatedAt,
		})
	}
	if kscauo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryattachment.FieldUpdatedAt,
		})
	}
	if value, ok := kscauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldDeletedAt,
		})
	}
	if kscauo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: kcsmallcategoryattachment.FieldDeletedAt,
		})
	}
	if value, ok := kscauo.mutation.AttachmentName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldAttachmentName,
		})
	}
	if kscauo.mutation.AttachmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryattachment.AttachmentTable,
			Columns: []string{kcsmallcategoryattachment.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kscauo.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryattachment.AttachmentTable,
			Columns: []string{kcsmallcategoryattachment.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if kscauo.mutation.SmallCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryattachment.SmallCategoryTable,
			Columns: []string{kcsmallcategoryattachment.SmallCategoryColumn},
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
	if nodes := kscauo.mutation.SmallCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryattachment.SmallCategoryTable,
			Columns: []string{kcsmallcategoryattachment.SmallCategoryColumn},
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
	_node = &KcSmallCategoryAttachment{config: kscauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kscauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kcsmallcategoryattachment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
