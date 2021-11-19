// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tkserver/internal/store/ent/tkexampaperpartition"
	"tkserver/internal/store/ent/tkexampartitionquestionlink"
	"tkserver/internal/store/ent/tkquestion"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkExamPartitionQuestionLinkCreate is the builder for creating a TkExamPartitionQuestionLink entity.
type TkExamPartitionQuestionLinkCreate struct {
	config
	mutation *TkExamPartitionQuestionLinkMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetUUID(s string) *TkExamPartitionQuestionLinkCreate {
	tepqlc.mutation.SetUUID(s)
	return tepqlc
}

// SetCreatedAt sets the "created_at" field.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetCreatedAt(t time.Time) *TkExamPartitionQuestionLinkCreate {
	tepqlc.mutation.SetCreatedAt(t)
	return tepqlc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetNillableCreatedAt(t *time.Time) *TkExamPartitionQuestionLinkCreate {
	if t != nil {
		tepqlc.SetCreatedAt(*t)
	}
	return tepqlc
}

// SetUpdatedAt sets the "updated_at" field.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetUpdatedAt(t time.Time) *TkExamPartitionQuestionLinkCreate {
	tepqlc.mutation.SetUpdatedAt(t)
	return tepqlc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetNillableUpdatedAt(t *time.Time) *TkExamPartitionQuestionLinkCreate {
	if t != nil {
		tepqlc.SetUpdatedAt(*t)
	}
	return tepqlc
}

// SetDeletedAt sets the "deleted_at" field.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetDeletedAt(t time.Time) *TkExamPartitionQuestionLinkCreate {
	tepqlc.mutation.SetDeletedAt(t)
	return tepqlc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetNillableDeletedAt(t *time.Time) *TkExamPartitionQuestionLinkCreate {
	if t != nil {
		tepqlc.SetDeletedAt(*t)
	}
	return tepqlc
}

// SetQuestionScore sets the "question_score" field.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetQuestionScore(u uint8) *TkExamPartitionQuestionLinkCreate {
	tepqlc.mutation.SetQuestionScore(u)
	return tepqlc
}

// SetNillableQuestionScore sets the "question_score" field if the given value is not nil.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetNillableQuestionScore(u *uint8) *TkExamPartitionQuestionLinkCreate {
	if u != nil {
		tepqlc.SetQuestionScore(*u)
	}
	return tepqlc
}

// SetExamPaperPartitionID sets the "exam_paper_partition_id" field.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetExamPaperPartitionID(i int) *TkExamPartitionQuestionLinkCreate {
	tepqlc.mutation.SetExamPaperPartitionID(i)
	return tepqlc
}

// SetNillableExamPaperPartitionID sets the "exam_paper_partition_id" field if the given value is not nil.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetNillableExamPaperPartitionID(i *int) *TkExamPartitionQuestionLinkCreate {
	if i != nil {
		tepqlc.SetExamPaperPartitionID(*i)
	}
	return tepqlc
}

// SetQuestionID sets the "question_id" field.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetQuestionID(i int) *TkExamPartitionQuestionLinkCreate {
	tepqlc.mutation.SetQuestionID(i)
	return tepqlc
}

// SetNillableQuestionID sets the "question_id" field if the given value is not nil.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetNillableQuestionID(i *int) *TkExamPartitionQuestionLinkCreate {
	if i != nil {
		tepqlc.SetQuestionID(*i)
	}
	return tepqlc
}

// SetExamPaperPartition sets the "exam_paper_partition" edge to the TkExamPaperPartition entity.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetExamPaperPartition(t *TkExamPaperPartition) *TkExamPartitionQuestionLinkCreate {
	return tepqlc.SetExamPaperPartitionID(t.ID)
}

// SetQuestion sets the "question" edge to the TkQuestion entity.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SetQuestion(t *TkQuestion) *TkExamPartitionQuestionLinkCreate {
	return tepqlc.SetQuestionID(t.ID)
}

// Mutation returns the TkExamPartitionQuestionLinkMutation object of the builder.
func (tepqlc *TkExamPartitionQuestionLinkCreate) Mutation() *TkExamPartitionQuestionLinkMutation {
	return tepqlc.mutation
}

// Save creates the TkExamPartitionQuestionLink in the database.
func (tepqlc *TkExamPartitionQuestionLinkCreate) Save(ctx context.Context) (*TkExamPartitionQuestionLink, error) {
	var (
		err  error
		node *TkExamPartitionQuestionLink
	)
	tepqlc.defaults()
	if len(tepqlc.hooks) == 0 {
		if err = tepqlc.check(); err != nil {
			return nil, err
		}
		node, err = tepqlc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkExamPartitionQuestionLinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tepqlc.check(); err != nil {
				return nil, err
			}
			tepqlc.mutation = mutation
			node, err = tepqlc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tepqlc.hooks) - 1; i >= 0; i-- {
			mut = tepqlc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tepqlc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tepqlc *TkExamPartitionQuestionLinkCreate) SaveX(ctx context.Context) *TkExamPartitionQuestionLink {
	v, err := tepqlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tepqlc *TkExamPartitionQuestionLinkCreate) defaults() {
	if _, ok := tepqlc.mutation.CreatedAt(); !ok {
		v := tkexampartitionquestionlink.DefaultCreatedAt()
		tepqlc.mutation.SetCreatedAt(v)
	}
	if _, ok := tepqlc.mutation.UpdatedAt(); !ok {
		v := tkexampartitionquestionlink.DefaultUpdatedAt()
		tepqlc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tepqlc.mutation.QuestionScore(); !ok {
		v := tkexampartitionquestionlink.DefaultQuestionScore
		tepqlc.mutation.SetQuestionScore(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tepqlc *TkExamPartitionQuestionLinkCreate) check() error {
	if _, ok := tepqlc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := tepqlc.mutation.QuestionScore(); !ok {
		return &ValidationError{Name: "question_score", err: errors.New("ent: missing required field \"question_score\"")}
	}
	return nil
}

func (tepqlc *TkExamPartitionQuestionLinkCreate) sqlSave(ctx context.Context) (*TkExamPartitionQuestionLink, error) {
	_node, _spec := tepqlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tepqlc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tepqlc *TkExamPartitionQuestionLinkCreate) createSpec() (*TkExamPartitionQuestionLink, *sqlgraph.CreateSpec) {
	var (
		_node = &TkExamPartitionQuestionLink{config: tepqlc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tkexampartitionquestionlink.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkexampartitionquestionlink.FieldID,
			},
		}
	)
	if value, ok := tepqlc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkexampartitionquestionlink.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := tepqlc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkexampartitionquestionlink.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := tepqlc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkexampartitionquestionlink.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := tepqlc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkexampartitionquestionlink.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := tepqlc.mutation.QuestionScore(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: tkexampartitionquestionlink.FieldQuestionScore,
		})
		_node.QuestionScore = value
	}
	if nodes := tepqlc.mutation.ExamPaperPartitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkexampartitionquestionlink.ExamPaperPartitionTable,
			Columns: []string{tkexampartitionquestionlink.ExamPaperPartitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkexampaperpartition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ExamPaperPartitionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tepqlc.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkexampartitionquestionlink.QuestionTable,
			Columns: []string{tkexampartitionquestionlink.QuestionColumn},
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

// TkExamPartitionQuestionLinkCreateBulk is the builder for creating many TkExamPartitionQuestionLink entities in bulk.
type TkExamPartitionQuestionLinkCreateBulk struct {
	config
	builders []*TkExamPartitionQuestionLinkCreate
}

// Save creates the TkExamPartitionQuestionLink entities in the database.
func (tepqlcb *TkExamPartitionQuestionLinkCreateBulk) Save(ctx context.Context) ([]*TkExamPartitionQuestionLink, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tepqlcb.builders))
	nodes := make([]*TkExamPartitionQuestionLink, len(tepqlcb.builders))
	mutators := make([]Mutator, len(tepqlcb.builders))
	for i := range tepqlcb.builders {
		func(i int, root context.Context) {
			builder := tepqlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TkExamPartitionQuestionLinkMutation)
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
					_, err = mutators[i+1].Mutate(root, tepqlcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tepqlcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tepqlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tepqlcb *TkExamPartitionQuestionLinkCreateBulk) SaveX(ctx context.Context) []*TkExamPartitionQuestionLink {
	v, err := tepqlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
