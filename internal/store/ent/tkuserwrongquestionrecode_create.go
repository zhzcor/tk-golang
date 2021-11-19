// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tkserver/internal/store/ent/tkquestion"
	"tkserver/internal/store/ent/tkuserwrongquestionrecode"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkUserWrongQuestionRecodeCreate is the builder for creating a TkUserWrongQuestionRecode entity.
type TkUserWrongQuestionRecodeCreate struct {
	config
	mutation *TkUserWrongQuestionRecodeMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetUUID(s string) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetUUID(s)
	return tuwqrc
}

// SetCreatedAt sets the "created_at" field.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetCreatedAt(t time.Time) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetCreatedAt(t)
	return tuwqrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetNillableCreatedAt(t *time.Time) *TkUserWrongQuestionRecodeCreate {
	if t != nil {
		tuwqrc.SetCreatedAt(*t)
	}
	return tuwqrc
}

// SetUpdatedAt sets the "updated_at" field.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetUpdatedAt(t time.Time) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetUpdatedAt(t)
	return tuwqrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetNillableUpdatedAt(t *time.Time) *TkUserWrongQuestionRecodeCreate {
	if t != nil {
		tuwqrc.SetUpdatedAt(*t)
	}
	return tuwqrc
}

// SetDeletedAt sets the "deleted_at" field.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetDeletedAt(t time.Time) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetDeletedAt(t)
	return tuwqrc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetNillableDeletedAt(t *time.Time) *TkUserWrongQuestionRecodeCreate {
	if t != nil {
		tuwqrc.SetDeletedAt(*t)
	}
	return tuwqrc
}

// SetUserID sets the "user_id" field.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetUserID(i int) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetUserID(i)
	return tuwqrc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetNillableUserID(i *int) *TkUserWrongQuestionRecodeCreate {
	if i != nil {
		tuwqrc.SetUserID(*i)
	}
	return tuwqrc
}

// SetQuestionID sets the "question_id" field.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetQuestionID(i int) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetQuestionID(i)
	return tuwqrc
}

// SetNillableQuestionID sets the "question_id" field if the given value is not nil.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetNillableQuestionID(i *int) *TkUserWrongQuestionRecodeCreate {
	if i != nil {
		tuwqrc.SetQuestionID(*i)
	}
	return tuwqrc
}

// SetQuestionBankID sets the "question_bank_id" field.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetQuestionBankID(i int) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetQuestionBankID(i)
	return tuwqrc
}

// SetNillableQuestionBankID sets the "question_bank_id" field if the given value is not nil.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetNillableQuestionBankID(i *int) *TkUserWrongQuestionRecodeCreate {
	if i != nil {
		tuwqrc.SetQuestionBankID(*i)
	}
	return tuwqrc
}

// SetAnswer sets the "answer" field.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetAnswer(s string) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetAnswer(s)
	return tuwqrc
}

// SetNillableAnswer sets the "answer" field if the given value is not nil.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetNillableAnswer(s *string) *TkUserWrongQuestionRecodeCreate {
	if s != nil {
		tuwqrc.SetAnswer(*s)
	}
	return tuwqrc
}

// SetWrongExamType sets the "wrong_exam_type" field.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetWrongExamType(i int) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetWrongExamType(i)
	return tuwqrc
}

// SetNillableWrongExamType sets the "wrong_exam_type" field if the given value is not nil.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetNillableWrongExamType(i *int) *TkUserWrongQuestionRecodeCreate {
	if i != nil {
		tuwqrc.SetWrongExamType(*i)
	}
	return tuwqrc
}

// SetWrongQuestionType sets the "wrong_question_type" field.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetWrongQuestionType(i int) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetWrongQuestionType(i)
	return tuwqrc
}

// SetNillableWrongQuestionType sets the "wrong_question_type" field if the given value is not nil.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetNillableWrongQuestionType(i *int) *TkUserWrongQuestionRecodeCreate {
	if i != nil {
		tuwqrc.SetWrongQuestionType(*i)
	}
	return tuwqrc
}

// SetQuestionWrongID sets the "question_wrong" edge to the TkQuestion entity by ID.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetQuestionWrongID(id int) *TkUserWrongQuestionRecodeCreate {
	tuwqrc.mutation.SetQuestionWrongID(id)
	return tuwqrc
}

// SetNillableQuestionWrongID sets the "question_wrong" edge to the TkQuestion entity by ID if the given value is not nil.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetNillableQuestionWrongID(id *int) *TkUserWrongQuestionRecodeCreate {
	if id != nil {
		tuwqrc = tuwqrc.SetQuestionWrongID(*id)
	}
	return tuwqrc
}

// SetQuestionWrong sets the "question_wrong" edge to the TkQuestion entity.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SetQuestionWrong(t *TkQuestion) *TkUserWrongQuestionRecodeCreate {
	return tuwqrc.SetQuestionWrongID(t.ID)
}

// Mutation returns the TkUserWrongQuestionRecodeMutation object of the builder.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) Mutation() *TkUserWrongQuestionRecodeMutation {
	return tuwqrc.mutation
}

// Save creates the TkUserWrongQuestionRecode in the database.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) Save(ctx context.Context) (*TkUserWrongQuestionRecode, error) {
	var (
		err  error
		node *TkUserWrongQuestionRecode
	)
	tuwqrc.defaults()
	if len(tuwqrc.hooks) == 0 {
		if err = tuwqrc.check(); err != nil {
			return nil, err
		}
		node, err = tuwqrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkUserWrongQuestionRecodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuwqrc.check(); err != nil {
				return nil, err
			}
			tuwqrc.mutation = mutation
			node, err = tuwqrc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuwqrc.hooks) - 1; i >= 0; i-- {
			mut = tuwqrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuwqrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) SaveX(ctx context.Context) *TkUserWrongQuestionRecode {
	v, err := tuwqrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) defaults() {
	if _, ok := tuwqrc.mutation.CreatedAt(); !ok {
		v := tkuserwrongquestionrecode.DefaultCreatedAt()
		tuwqrc.mutation.SetCreatedAt(v)
	}
	if _, ok := tuwqrc.mutation.UpdatedAt(); !ok {
		v := tkuserwrongquestionrecode.DefaultUpdatedAt()
		tuwqrc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tuwqrc.mutation.WrongExamType(); !ok {
		v := tkuserwrongquestionrecode.DefaultWrongExamType
		tuwqrc.mutation.SetWrongExamType(v)
	}
	if _, ok := tuwqrc.mutation.WrongQuestionType(); !ok {
		v := tkuserwrongquestionrecode.DefaultWrongQuestionType
		tuwqrc.mutation.SetWrongQuestionType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuwqrc *TkUserWrongQuestionRecodeCreate) check() error {
	if _, ok := tuwqrc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := tuwqrc.mutation.WrongExamType(); !ok {
		return &ValidationError{Name: "wrong_exam_type", err: errors.New("ent: missing required field \"wrong_exam_type\"")}
	}
	if _, ok := tuwqrc.mutation.WrongQuestionType(); !ok {
		return &ValidationError{Name: "wrong_question_type", err: errors.New("ent: missing required field \"wrong_question_type\"")}
	}
	return nil
}

func (tuwqrc *TkUserWrongQuestionRecodeCreate) sqlSave(ctx context.Context) (*TkUserWrongQuestionRecode, error) {
	_node, _spec := tuwqrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tuwqrc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tuwqrc *TkUserWrongQuestionRecodeCreate) createSpec() (*TkUserWrongQuestionRecode, *sqlgraph.CreateSpec) {
	var (
		_node = &TkUserWrongQuestionRecode{config: tuwqrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tkuserwrongquestionrecode.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkuserwrongquestionrecode.FieldID,
			},
		}
	)
	if value, ok := tuwqrc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkuserwrongquestionrecode.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := tuwqrc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkuserwrongquestionrecode.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := tuwqrc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkuserwrongquestionrecode.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := tuwqrc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkuserwrongquestionrecode.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := tuwqrc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkuserwrongquestionrecode.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := tuwqrc.mutation.QuestionBankID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkuserwrongquestionrecode.FieldQuestionBankID,
		})
		_node.QuestionBankID = value
	}
	if value, ok := tuwqrc.mutation.Answer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkuserwrongquestionrecode.FieldAnswer,
		})
		_node.Answer = value
	}
	if value, ok := tuwqrc.mutation.WrongExamType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkuserwrongquestionrecode.FieldWrongExamType,
		})
		_node.WrongExamType = value
	}
	if value, ok := tuwqrc.mutation.WrongQuestionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkuserwrongquestionrecode.FieldWrongQuestionType,
		})
		_node.WrongQuestionType = value
	}
	if nodes := tuwqrc.mutation.QuestionWrongIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkuserwrongquestionrecode.QuestionWrongTable,
			Columns: []string{tkuserwrongquestionrecode.QuestionWrongColumn},
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
		_node.QuestionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TkUserWrongQuestionRecodeCreateBulk is the builder for creating many TkUserWrongQuestionRecode entities in bulk.
type TkUserWrongQuestionRecodeCreateBulk struct {
	config
	builders []*TkUserWrongQuestionRecodeCreate
}

// Save creates the TkUserWrongQuestionRecode entities in the database.
func (tuwqrcb *TkUserWrongQuestionRecodeCreateBulk) Save(ctx context.Context) ([]*TkUserWrongQuestionRecode, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tuwqrcb.builders))
	nodes := make([]*TkUserWrongQuestionRecode, len(tuwqrcb.builders))
	mutators := make([]Mutator, len(tuwqrcb.builders))
	for i := range tuwqrcb.builders {
		func(i int, root context.Context) {
			builder := tuwqrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TkUserWrongQuestionRecodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tuwqrcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tuwqrcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tuwqrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tuwqrcb *TkUserWrongQuestionRecodeCreateBulk) SaveX(ctx context.Context) []*TkUserWrongQuestionRecode {
	v, err := tuwqrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
