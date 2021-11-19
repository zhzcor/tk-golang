// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/tkknowledgepoint"
	"gserver/internal/store/ent/tkquestion"
	"gserver/internal/store/ent/tkquestionbank"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkKnowledgePointUpdate is the builder for updating TkKnowledgePoint entities.
type TkKnowledgePointUpdate struct {
	config
	hooks    []Hook
	mutation *TkKnowledgePointMutation
}

// Where adds a new predicate for the TkKnowledgePointUpdate builder.
func (tkpu *TkKnowledgePointUpdate) Where(ps ...predicate.TkKnowledgePoint) *TkKnowledgePointUpdate {
	tkpu.mutation.predicates = append(tkpu.mutation.predicates, ps...)
	return tkpu
}

// SetUUID sets the "uuid" field.
func (tkpu *TkKnowledgePointUpdate) SetUUID(s string) *TkKnowledgePointUpdate {
	tkpu.mutation.SetUUID(s)
	return tkpu
}

// SetUpdatedAt sets the "updated_at" field.
func (tkpu *TkKnowledgePointUpdate) SetUpdatedAt(t time.Time) *TkKnowledgePointUpdate {
	tkpu.mutation.SetUpdatedAt(t)
	return tkpu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tkpu *TkKnowledgePointUpdate) ClearUpdatedAt() *TkKnowledgePointUpdate {
	tkpu.mutation.ClearUpdatedAt()
	return tkpu
}

// SetDeletedAt sets the "deleted_at" field.
func (tkpu *TkKnowledgePointUpdate) SetDeletedAt(t time.Time) *TkKnowledgePointUpdate {
	tkpu.mutation.SetDeletedAt(t)
	return tkpu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tkpu *TkKnowledgePointUpdate) SetNillableDeletedAt(t *time.Time) *TkKnowledgePointUpdate {
	if t != nil {
		tkpu.SetDeletedAt(*t)
	}
	return tkpu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tkpu *TkKnowledgePointUpdate) ClearDeletedAt() *TkKnowledgePointUpdate {
	tkpu.mutation.ClearDeletedAt()
	return tkpu
}

// SetName sets the "name" field.
func (tkpu *TkKnowledgePointUpdate) SetName(s string) *TkKnowledgePointUpdate {
	tkpu.mutation.SetName(s)
	return tkpu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tkpu *TkKnowledgePointUpdate) SetNillableName(s *string) *TkKnowledgePointUpdate {
	if s != nil {
		tkpu.SetName(*s)
	}
	return tkpu
}

// SetQuestionBankID sets the "question_bank_id" field.
func (tkpu *TkKnowledgePointUpdate) SetQuestionBankID(i int) *TkKnowledgePointUpdate {
	tkpu.mutation.ResetQuestionBankID()
	tkpu.mutation.SetQuestionBankID(i)
	return tkpu
}

// SetNillableQuestionBankID sets the "question_bank_id" field if the given value is not nil.
func (tkpu *TkKnowledgePointUpdate) SetNillableQuestionBankID(i *int) *TkKnowledgePointUpdate {
	if i != nil {
		tkpu.SetQuestionBankID(*i)
	}
	return tkpu
}

// ClearQuestionBankID clears the value of the "question_bank_id" field.
func (tkpu *TkKnowledgePointUpdate) ClearQuestionBankID() *TkKnowledgePointUpdate {
	tkpu.mutation.ClearQuestionBankID()
	return tkpu
}

// SetQuestionCount sets the "question_count" field.
func (tkpu *TkKnowledgePointUpdate) SetQuestionCount(i int) *TkKnowledgePointUpdate {
	tkpu.mutation.ResetQuestionCount()
	tkpu.mutation.SetQuestionCount(i)
	return tkpu
}

// SetNillableQuestionCount sets the "question_count" field if the given value is not nil.
func (tkpu *TkKnowledgePointUpdate) SetNillableQuestionCount(i *int) *TkKnowledgePointUpdate {
	if i != nil {
		tkpu.SetQuestionCount(*i)
	}
	return tkpu
}

// AddQuestionCount adds i to the "question_count" field.
func (tkpu *TkKnowledgePointUpdate) AddQuestionCount(i int) *TkKnowledgePointUpdate {
	tkpu.mutation.AddQuestionCount(i)
	return tkpu
}

// SetQuestionBank sets the "question_bank" edge to the TkQuestionBank entity.
func (tkpu *TkKnowledgePointUpdate) SetQuestionBank(t *TkQuestionBank) *TkKnowledgePointUpdate {
	return tkpu.SetQuestionBankID(t.ID)
}

// AddQuestionIDs adds the "questions" edge to the TkQuestion entity by IDs.
func (tkpu *TkKnowledgePointUpdate) AddQuestionIDs(ids ...int) *TkKnowledgePointUpdate {
	tkpu.mutation.AddQuestionIDs(ids...)
	return tkpu
}

// AddQuestions adds the "questions" edges to the TkQuestion entity.
func (tkpu *TkKnowledgePointUpdate) AddQuestions(t ...*TkQuestion) *TkKnowledgePointUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tkpu.AddQuestionIDs(ids...)
}

// Mutation returns the TkKnowledgePointMutation object of the builder.
func (tkpu *TkKnowledgePointUpdate) Mutation() *TkKnowledgePointMutation {
	return tkpu.mutation
}

// ClearQuestionBank clears the "question_bank" edge to the TkQuestionBank entity.
func (tkpu *TkKnowledgePointUpdate) ClearQuestionBank() *TkKnowledgePointUpdate {
	tkpu.mutation.ClearQuestionBank()
	return tkpu
}

// ClearQuestions clears all "questions" edges to the TkQuestion entity.
func (tkpu *TkKnowledgePointUpdate) ClearQuestions() *TkKnowledgePointUpdate {
	tkpu.mutation.ClearQuestions()
	return tkpu
}

// RemoveQuestionIDs removes the "questions" edge to TkQuestion entities by IDs.
func (tkpu *TkKnowledgePointUpdate) RemoveQuestionIDs(ids ...int) *TkKnowledgePointUpdate {
	tkpu.mutation.RemoveQuestionIDs(ids...)
	return tkpu
}

// RemoveQuestions removes "questions" edges to TkQuestion entities.
func (tkpu *TkKnowledgePointUpdate) RemoveQuestions(t ...*TkQuestion) *TkKnowledgePointUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tkpu.RemoveQuestionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tkpu *TkKnowledgePointUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tkpu.defaults()
	if len(tkpu.hooks) == 0 {
		affected, err = tkpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkKnowledgePointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tkpu.mutation = mutation
			affected, err = tkpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tkpu.hooks) - 1; i >= 0; i-- {
			mut = tkpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tkpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tkpu *TkKnowledgePointUpdate) SaveX(ctx context.Context) int {
	affected, err := tkpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tkpu *TkKnowledgePointUpdate) Exec(ctx context.Context) error {
	_, err := tkpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tkpu *TkKnowledgePointUpdate) ExecX(ctx context.Context) {
	if err := tkpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tkpu *TkKnowledgePointUpdate) defaults() {
	if _, ok := tkpu.mutation.UpdatedAt(); !ok && !tkpu.mutation.UpdatedAtCleared() {
		v := tkknowledgepoint.UpdateDefaultUpdatedAt()
		tkpu.mutation.SetUpdatedAt(v)
	}
}

func (tkpu *TkKnowledgePointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tkknowledgepoint.Table,
			Columns: tkknowledgepoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkknowledgepoint.FieldID,
			},
		},
	}
	if ps := tkpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tkpu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkknowledgepoint.FieldUUID,
		})
	}
	if tkpu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkknowledgepoint.FieldCreatedAt,
		})
	}
	if value, ok := tkpu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkknowledgepoint.FieldUpdatedAt,
		})
	}
	if tkpu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkknowledgepoint.FieldUpdatedAt,
		})
	}
	if value, ok := tkpu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkknowledgepoint.FieldDeletedAt,
		})
	}
	if tkpu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkknowledgepoint.FieldDeletedAt,
		})
	}
	if value, ok := tkpu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkknowledgepoint.FieldName,
		})
	}
	if value, ok := tkpu.mutation.QuestionCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkknowledgepoint.FieldQuestionCount,
		})
	}
	if value, ok := tkpu.mutation.AddedQuestionCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkknowledgepoint.FieldQuestionCount,
		})
	}
	if tkpu.mutation.QuestionBankCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkknowledgepoint.QuestionBankTable,
			Columns: []string{tkknowledgepoint.QuestionBankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestionbank.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tkpu.mutation.QuestionBankIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkknowledgepoint.QuestionBankTable,
			Columns: []string{tkknowledgepoint.QuestionBankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestionbank.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tkpu.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tkknowledgepoint.QuestionsTable,
			Columns: tkknowledgepoint.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestion.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tkpu.mutation.RemovedQuestionsIDs(); len(nodes) > 0 && !tkpu.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tkknowledgepoint.QuestionsTable,
			Columns: tkknowledgepoint.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tkpu.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tkknowledgepoint.QuestionsTable,
			Columns: tkknowledgepoint.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tkpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tkknowledgepoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TkKnowledgePointUpdateOne is the builder for updating a single TkKnowledgePoint entity.
type TkKnowledgePointUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TkKnowledgePointMutation
}

// SetUUID sets the "uuid" field.
func (tkpuo *TkKnowledgePointUpdateOne) SetUUID(s string) *TkKnowledgePointUpdateOne {
	tkpuo.mutation.SetUUID(s)
	return tkpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tkpuo *TkKnowledgePointUpdateOne) SetUpdatedAt(t time.Time) *TkKnowledgePointUpdateOne {
	tkpuo.mutation.SetUpdatedAt(t)
	return tkpuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tkpuo *TkKnowledgePointUpdateOne) ClearUpdatedAt() *TkKnowledgePointUpdateOne {
	tkpuo.mutation.ClearUpdatedAt()
	return tkpuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tkpuo *TkKnowledgePointUpdateOne) SetDeletedAt(t time.Time) *TkKnowledgePointUpdateOne {
	tkpuo.mutation.SetDeletedAt(t)
	return tkpuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tkpuo *TkKnowledgePointUpdateOne) SetNillableDeletedAt(t *time.Time) *TkKnowledgePointUpdateOne {
	if t != nil {
		tkpuo.SetDeletedAt(*t)
	}
	return tkpuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tkpuo *TkKnowledgePointUpdateOne) ClearDeletedAt() *TkKnowledgePointUpdateOne {
	tkpuo.mutation.ClearDeletedAt()
	return tkpuo
}

// SetName sets the "name" field.
func (tkpuo *TkKnowledgePointUpdateOne) SetName(s string) *TkKnowledgePointUpdateOne {
	tkpuo.mutation.SetName(s)
	return tkpuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tkpuo *TkKnowledgePointUpdateOne) SetNillableName(s *string) *TkKnowledgePointUpdateOne {
	if s != nil {
		tkpuo.SetName(*s)
	}
	return tkpuo
}

// SetQuestionBankID sets the "question_bank_id" field.
func (tkpuo *TkKnowledgePointUpdateOne) SetQuestionBankID(i int) *TkKnowledgePointUpdateOne {
	tkpuo.mutation.ResetQuestionBankID()
	tkpuo.mutation.SetQuestionBankID(i)
	return tkpuo
}

// SetNillableQuestionBankID sets the "question_bank_id" field if the given value is not nil.
func (tkpuo *TkKnowledgePointUpdateOne) SetNillableQuestionBankID(i *int) *TkKnowledgePointUpdateOne {
	if i != nil {
		tkpuo.SetQuestionBankID(*i)
	}
	return tkpuo
}

// ClearQuestionBankID clears the value of the "question_bank_id" field.
func (tkpuo *TkKnowledgePointUpdateOne) ClearQuestionBankID() *TkKnowledgePointUpdateOne {
	tkpuo.mutation.ClearQuestionBankID()
	return tkpuo
}

// SetQuestionCount sets the "question_count" field.
func (tkpuo *TkKnowledgePointUpdateOne) SetQuestionCount(i int) *TkKnowledgePointUpdateOne {
	tkpuo.mutation.ResetQuestionCount()
	tkpuo.mutation.SetQuestionCount(i)
	return tkpuo
}

// SetNillableQuestionCount sets the "question_count" field if the given value is not nil.
func (tkpuo *TkKnowledgePointUpdateOne) SetNillableQuestionCount(i *int) *TkKnowledgePointUpdateOne {
	if i != nil {
		tkpuo.SetQuestionCount(*i)
	}
	return tkpuo
}

// AddQuestionCount adds i to the "question_count" field.
func (tkpuo *TkKnowledgePointUpdateOne) AddQuestionCount(i int) *TkKnowledgePointUpdateOne {
	tkpuo.mutation.AddQuestionCount(i)
	return tkpuo
}

// SetQuestionBank sets the "question_bank" edge to the TkQuestionBank entity.
func (tkpuo *TkKnowledgePointUpdateOne) SetQuestionBank(t *TkQuestionBank) *TkKnowledgePointUpdateOne {
	return tkpuo.SetQuestionBankID(t.ID)
}

// AddQuestionIDs adds the "questions" edge to the TkQuestion entity by IDs.
func (tkpuo *TkKnowledgePointUpdateOne) AddQuestionIDs(ids ...int) *TkKnowledgePointUpdateOne {
	tkpuo.mutation.AddQuestionIDs(ids...)
	return tkpuo
}

// AddQuestions adds the "questions" edges to the TkQuestion entity.
func (tkpuo *TkKnowledgePointUpdateOne) AddQuestions(t ...*TkQuestion) *TkKnowledgePointUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tkpuo.AddQuestionIDs(ids...)
}

// Mutation returns the TkKnowledgePointMutation object of the builder.
func (tkpuo *TkKnowledgePointUpdateOne) Mutation() *TkKnowledgePointMutation {
	return tkpuo.mutation
}

// ClearQuestionBank clears the "question_bank" edge to the TkQuestionBank entity.
func (tkpuo *TkKnowledgePointUpdateOne) ClearQuestionBank() *TkKnowledgePointUpdateOne {
	tkpuo.mutation.ClearQuestionBank()
	return tkpuo
}

// ClearQuestions clears all "questions" edges to the TkQuestion entity.
func (tkpuo *TkKnowledgePointUpdateOne) ClearQuestions() *TkKnowledgePointUpdateOne {
	tkpuo.mutation.ClearQuestions()
	return tkpuo
}

// RemoveQuestionIDs removes the "questions" edge to TkQuestion entities by IDs.
func (tkpuo *TkKnowledgePointUpdateOne) RemoveQuestionIDs(ids ...int) *TkKnowledgePointUpdateOne {
	tkpuo.mutation.RemoveQuestionIDs(ids...)
	return tkpuo
}

// RemoveQuestions removes "questions" edges to TkQuestion entities.
func (tkpuo *TkKnowledgePointUpdateOne) RemoveQuestions(t ...*TkQuestion) *TkKnowledgePointUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tkpuo.RemoveQuestionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tkpuo *TkKnowledgePointUpdateOne) Select(field string, fields ...string) *TkKnowledgePointUpdateOne {
	tkpuo.fields = append([]string{field}, fields...)
	return tkpuo
}

// Save executes the query and returns the updated TkKnowledgePoint entity.
func (tkpuo *TkKnowledgePointUpdateOne) Save(ctx context.Context) (*TkKnowledgePoint, error) {
	var (
		err  error
		node *TkKnowledgePoint
	)
	tkpuo.defaults()
	if len(tkpuo.hooks) == 0 {
		node, err = tkpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkKnowledgePointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tkpuo.mutation = mutation
			node, err = tkpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tkpuo.hooks) - 1; i >= 0; i-- {
			mut = tkpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tkpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tkpuo *TkKnowledgePointUpdateOne) SaveX(ctx context.Context) *TkKnowledgePoint {
	node, err := tkpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tkpuo *TkKnowledgePointUpdateOne) Exec(ctx context.Context) error {
	_, err := tkpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tkpuo *TkKnowledgePointUpdateOne) ExecX(ctx context.Context) {
	if err := tkpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tkpuo *TkKnowledgePointUpdateOne) defaults() {
	if _, ok := tkpuo.mutation.UpdatedAt(); !ok && !tkpuo.mutation.UpdatedAtCleared() {
		v := tkknowledgepoint.UpdateDefaultUpdatedAt()
		tkpuo.mutation.SetUpdatedAt(v)
	}
}

func (tkpuo *TkKnowledgePointUpdateOne) sqlSave(ctx context.Context) (_node *TkKnowledgePoint, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tkknowledgepoint.Table,
			Columns: tkknowledgepoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkknowledgepoint.FieldID,
			},
		},
	}
	id, ok := tkpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TkKnowledgePoint.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tkpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tkknowledgepoint.FieldID)
		for _, f := range fields {
			if !tkknowledgepoint.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tkknowledgepoint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tkpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tkpuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkknowledgepoint.FieldUUID,
		})
	}
	if tkpuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkknowledgepoint.FieldCreatedAt,
		})
	}
	if value, ok := tkpuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkknowledgepoint.FieldUpdatedAt,
		})
	}
	if tkpuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkknowledgepoint.FieldUpdatedAt,
		})
	}
	if value, ok := tkpuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkknowledgepoint.FieldDeletedAt,
		})
	}
	if tkpuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkknowledgepoint.FieldDeletedAt,
		})
	}
	if value, ok := tkpuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkknowledgepoint.FieldName,
		})
	}
	if value, ok := tkpuo.mutation.QuestionCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkknowledgepoint.FieldQuestionCount,
		})
	}
	if value, ok := tkpuo.mutation.AddedQuestionCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkknowledgepoint.FieldQuestionCount,
		})
	}
	if tkpuo.mutation.QuestionBankCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkknowledgepoint.QuestionBankTable,
			Columns: []string{tkknowledgepoint.QuestionBankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestionbank.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tkpuo.mutation.QuestionBankIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkknowledgepoint.QuestionBankTable,
			Columns: []string{tkknowledgepoint.QuestionBankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestionbank.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tkpuo.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tkknowledgepoint.QuestionsTable,
			Columns: tkknowledgepoint.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestion.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tkpuo.mutation.RemovedQuestionsIDs(); len(nodes) > 0 && !tkpuo.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tkknowledgepoint.QuestionsTable,
			Columns: tkknowledgepoint.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tkpuo.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tkknowledgepoint.QuestionsTable,
			Columns: tkknowledgepoint.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TkKnowledgePoint{config: tkpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tkpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tkknowledgepoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
