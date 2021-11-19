// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/tkquestionansweroption"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkQuestionAnswerOptionDelete is the builder for deleting a TkQuestionAnswerOption entity.
type TkQuestionAnswerOptionDelete struct {
	config
	hooks    []Hook
	mutation *TkQuestionAnswerOptionMutation
}

// Where adds a new predicate to the TkQuestionAnswerOptionDelete builder.
func (tqaod *TkQuestionAnswerOptionDelete) Where(ps ...predicate.TkQuestionAnswerOption) *TkQuestionAnswerOptionDelete {
	tqaod.mutation.predicates = append(tqaod.mutation.predicates, ps...)
	return tqaod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tqaod *TkQuestionAnswerOptionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tqaod.hooks) == 0 {
		affected, err = tqaod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkQuestionAnswerOptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tqaod.mutation = mutation
			affected, err = tqaod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tqaod.hooks) - 1; i >= 0; i-- {
			mut = tqaod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tqaod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqaod *TkQuestionAnswerOptionDelete) ExecX(ctx context.Context) int {
	n, err := tqaod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tqaod *TkQuestionAnswerOptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tkquestionansweroption.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkquestionansweroption.FieldID,
			},
		},
	}
	if ps := tqaod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tqaod.driver, _spec)
}

// TkQuestionAnswerOptionDeleteOne is the builder for deleting a single TkQuestionAnswerOption entity.
type TkQuestionAnswerOptionDeleteOne struct {
	tqaod *TkQuestionAnswerOptionDelete
}

// Exec executes the deletion query.
func (tqaodo *TkQuestionAnswerOptionDeleteOne) Exec(ctx context.Context) error {
	n, err := tqaodo.tqaod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tkquestionansweroption.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tqaodo *TkQuestionAnswerOptionDeleteOne) ExecX(ctx context.Context) {
	tqaodo.tqaod.ExecX(ctx)
}
