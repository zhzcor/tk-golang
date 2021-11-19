// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/teacher"
	"tkserver/internal/store/ent/teachertag"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeacherTagUpdate is the builder for updating TeacherTag entities.
type TeacherTagUpdate struct {
	config
	hooks    []Hook
	mutation *TeacherTagMutation
}

// Where adds a new predicate for the TeacherTagUpdate builder.
func (ttu *TeacherTagUpdate) Where(ps ...predicate.TeacherTag) *TeacherTagUpdate {
	ttu.mutation.predicates = append(ttu.mutation.predicates, ps...)
	return ttu
}

// SetUUID sets the "uuid" field.
func (ttu *TeacherTagUpdate) SetUUID(s string) *TeacherTagUpdate {
	ttu.mutation.SetUUID(s)
	return ttu
}

// SetUpdatedAt sets the "updated_at" field.
func (ttu *TeacherTagUpdate) SetUpdatedAt(t time.Time) *TeacherTagUpdate {
	ttu.mutation.SetUpdatedAt(t)
	return ttu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ttu *TeacherTagUpdate) ClearUpdatedAt() *TeacherTagUpdate {
	ttu.mutation.ClearUpdatedAt()
	return ttu
}

// SetDeletedAt sets the "deleted_at" field.
func (ttu *TeacherTagUpdate) SetDeletedAt(t time.Time) *TeacherTagUpdate {
	ttu.mutation.SetDeletedAt(t)
	return ttu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ttu *TeacherTagUpdate) SetNillableDeletedAt(t *time.Time) *TeacherTagUpdate {
	if t != nil {
		ttu.SetDeletedAt(*t)
	}
	return ttu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ttu *TeacherTagUpdate) ClearDeletedAt() *TeacherTagUpdate {
	ttu.mutation.ClearDeletedAt()
	return ttu
}

// SetName sets the "name" field.
func (ttu *TeacherTagUpdate) SetName(s string) *TeacherTagUpdate {
	ttu.mutation.SetName(s)
	return ttu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ttu *TeacherTagUpdate) SetNillableName(s *string) *TeacherTagUpdate {
	if s != nil {
		ttu.SetName(*s)
	}
	return ttu
}

// SetTeacherID sets the "teacher_id" field.
func (ttu *TeacherTagUpdate) SetTeacherID(i int) *TeacherTagUpdate {
	ttu.mutation.ResetTeacherID()
	ttu.mutation.SetTeacherID(i)
	return ttu
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (ttu *TeacherTagUpdate) SetNillableTeacherID(i *int) *TeacherTagUpdate {
	if i != nil {
		ttu.SetTeacherID(*i)
	}
	return ttu
}

// ClearTeacherID clears the value of the "teacher_id" field.
func (ttu *TeacherTagUpdate) ClearTeacherID() *TeacherTagUpdate {
	ttu.mutation.ClearTeacherID()
	return ttu
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (ttu *TeacherTagUpdate) SetTeacher(t *Teacher) *TeacherTagUpdate {
	return ttu.SetTeacherID(t.ID)
}

// Mutation returns the TeacherTagMutation object of the builder.
func (ttu *TeacherTagUpdate) Mutation() *TeacherTagMutation {
	return ttu.mutation
}

// ClearTeacher clears the "teacher" edge to the Teacher entity.
func (ttu *TeacherTagUpdate) ClearTeacher() *TeacherTagUpdate {
	ttu.mutation.ClearTeacher()
	return ttu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ttu *TeacherTagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ttu.defaults()
	if len(ttu.hooks) == 0 {
		affected, err = ttu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeacherTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ttu.mutation = mutation
			affected, err = ttu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttu.hooks) - 1; i >= 0; i-- {
			mut = ttu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttu *TeacherTagUpdate) SaveX(ctx context.Context) int {
	affected, err := ttu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ttu *TeacherTagUpdate) Exec(ctx context.Context) error {
	_, err := ttu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttu *TeacherTagUpdate) ExecX(ctx context.Context) {
	if err := ttu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttu *TeacherTagUpdate) defaults() {
	if _, ok := ttu.mutation.UpdatedAt(); !ok && !ttu.mutation.UpdatedAtCleared() {
		v := teachertag.UpdateDefaultUpdatedAt()
		ttu.mutation.SetUpdatedAt(v)
	}
}

func (ttu *TeacherTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teachertag.Table,
			Columns: teachertag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: teachertag.FieldID,
			},
		},
	}
	if ps := ttu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teachertag.FieldUUID,
		})
	}
	if ttu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: teachertag.FieldCreatedAt,
		})
	}
	if value, ok := ttu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teachertag.FieldUpdatedAt,
		})
	}
	if ttu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: teachertag.FieldUpdatedAt,
		})
	}
	if value, ok := ttu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teachertag.FieldDeletedAt,
		})
	}
	if ttu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: teachertag.FieldDeletedAt,
		})
	}
	if value, ok := ttu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teachertag.FieldName,
		})
	}
	if ttu.mutation.TeacherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teachertag.TeacherTable,
			Columns: []string{teachertag.TeacherColumn},
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
	if nodes := ttu.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teachertag.TeacherTable,
			Columns: []string{teachertag.TeacherColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ttu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teachertag.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TeacherTagUpdateOne is the builder for updating a single TeacherTag entity.
type TeacherTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeacherTagMutation
}

// SetUUID sets the "uuid" field.
func (ttuo *TeacherTagUpdateOne) SetUUID(s string) *TeacherTagUpdateOne {
	ttuo.mutation.SetUUID(s)
	return ttuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ttuo *TeacherTagUpdateOne) SetUpdatedAt(t time.Time) *TeacherTagUpdateOne {
	ttuo.mutation.SetUpdatedAt(t)
	return ttuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ttuo *TeacherTagUpdateOne) ClearUpdatedAt() *TeacherTagUpdateOne {
	ttuo.mutation.ClearUpdatedAt()
	return ttuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ttuo *TeacherTagUpdateOne) SetDeletedAt(t time.Time) *TeacherTagUpdateOne {
	ttuo.mutation.SetDeletedAt(t)
	return ttuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ttuo *TeacherTagUpdateOne) SetNillableDeletedAt(t *time.Time) *TeacherTagUpdateOne {
	if t != nil {
		ttuo.SetDeletedAt(*t)
	}
	return ttuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ttuo *TeacherTagUpdateOne) ClearDeletedAt() *TeacherTagUpdateOne {
	ttuo.mutation.ClearDeletedAt()
	return ttuo
}

// SetName sets the "name" field.
func (ttuo *TeacherTagUpdateOne) SetName(s string) *TeacherTagUpdateOne {
	ttuo.mutation.SetName(s)
	return ttuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ttuo *TeacherTagUpdateOne) SetNillableName(s *string) *TeacherTagUpdateOne {
	if s != nil {
		ttuo.SetName(*s)
	}
	return ttuo
}

// SetTeacherID sets the "teacher_id" field.
func (ttuo *TeacherTagUpdateOne) SetTeacherID(i int) *TeacherTagUpdateOne {
	ttuo.mutation.ResetTeacherID()
	ttuo.mutation.SetTeacherID(i)
	return ttuo
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (ttuo *TeacherTagUpdateOne) SetNillableTeacherID(i *int) *TeacherTagUpdateOne {
	if i != nil {
		ttuo.SetTeacherID(*i)
	}
	return ttuo
}

// ClearTeacherID clears the value of the "teacher_id" field.
func (ttuo *TeacherTagUpdateOne) ClearTeacherID() *TeacherTagUpdateOne {
	ttuo.mutation.ClearTeacherID()
	return ttuo
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (ttuo *TeacherTagUpdateOne) SetTeacher(t *Teacher) *TeacherTagUpdateOne {
	return ttuo.SetTeacherID(t.ID)
}

// Mutation returns the TeacherTagMutation object of the builder.
func (ttuo *TeacherTagUpdateOne) Mutation() *TeacherTagMutation {
	return ttuo.mutation
}

// ClearTeacher clears the "teacher" edge to the Teacher entity.
func (ttuo *TeacherTagUpdateOne) ClearTeacher() *TeacherTagUpdateOne {
	ttuo.mutation.ClearTeacher()
	return ttuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ttuo *TeacherTagUpdateOne) Select(field string, fields ...string) *TeacherTagUpdateOne {
	ttuo.fields = append([]string{field}, fields...)
	return ttuo
}

// Save executes the query and returns the updated TeacherTag entity.
func (ttuo *TeacherTagUpdateOne) Save(ctx context.Context) (*TeacherTag, error) {
	var (
		err  error
		node *TeacherTag
	)
	ttuo.defaults()
	if len(ttuo.hooks) == 0 {
		node, err = ttuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeacherTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ttuo.mutation = mutation
			node, err = ttuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ttuo.hooks) - 1; i >= 0; i-- {
			mut = ttuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttuo *TeacherTagUpdateOne) SaveX(ctx context.Context) *TeacherTag {
	node, err := ttuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ttuo *TeacherTagUpdateOne) Exec(ctx context.Context) error {
	_, err := ttuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttuo *TeacherTagUpdateOne) ExecX(ctx context.Context) {
	if err := ttuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttuo *TeacherTagUpdateOne) defaults() {
	if _, ok := ttuo.mutation.UpdatedAt(); !ok && !ttuo.mutation.UpdatedAtCleared() {
		v := teachertag.UpdateDefaultUpdatedAt()
		ttuo.mutation.SetUpdatedAt(v)
	}
}

func (ttuo *TeacherTagUpdateOne) sqlSave(ctx context.Context) (_node *TeacherTag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teachertag.Table,
			Columns: teachertag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: teachertag.FieldID,
			},
		},
	}
	id, ok := ttuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TeacherTag.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ttuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teachertag.FieldID)
		for _, f := range fields {
			if !teachertag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != teachertag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ttuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teachertag.FieldUUID,
		})
	}
	if ttuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: teachertag.FieldCreatedAt,
		})
	}
	if value, ok := ttuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teachertag.FieldUpdatedAt,
		})
	}
	if ttuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: teachertag.FieldUpdatedAt,
		})
	}
	if value, ok := ttuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teachertag.FieldDeletedAt,
		})
	}
	if ttuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: teachertag.FieldDeletedAt,
		})
	}
	if value, ok := ttuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teachertag.FieldName,
		})
	}
	if ttuo.mutation.TeacherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teachertag.TeacherTable,
			Columns: []string{teachertag.TeacherColumn},
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
	if nodes := ttuo.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teachertag.TeacherTable,
			Columns: []string{teachertag.TeacherColumn},
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
	_node = &TeacherTag{config: ttuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ttuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teachertag.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
