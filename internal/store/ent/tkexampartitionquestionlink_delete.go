// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/tkexampartitionquestionlink"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkExamPartitionQuestionLinkDelete is the builder for deleting a TkExamPartitionQuestionLink entity.
type TkExamPartitionQuestionLinkDelete struct {
	config
	hooks    []Hook
	mutation *TkExamPartitionQuestionLinkMutation
}

// Where adds a new predicate to the TkExamPartitionQuestionLinkDelete builder.
func (tepqld *TkExamPartitionQuestionLinkDelete) Where(ps ...predicate.TkExamPartitionQuestionLink) *TkExamPartitionQuestionLinkDelete {
	tepqld.mutation.predicates = append(tepqld.mutation.predicates, ps...)
	return tepqld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tepqld *TkExamPartitionQuestionLinkDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tepqld.hooks) == 0 {
		affected, err = tepqld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkExamPartitionQuestionLinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tepqld.mutation = mutation
			affected, err = tepqld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tepqld.hooks) - 1; i >= 0; i-- {
			mut = tepqld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tepqld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tepqld *TkExamPartitionQuestionLinkDelete) ExecX(ctx context.Context) int {
	n, err := tepqld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tepqld *TkExamPartitionQuestionLinkDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tkexampartitionquestionlink.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkexampartitionquestionlink.FieldID,
			},
		},
	}
	if ps := tepqld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tepqld.driver, _spec)
}

// TkExamPartitionQuestionLinkDeleteOne is the builder for deleting a single TkExamPartitionQuestionLink entity.
type TkExamPartitionQuestionLinkDeleteOne struct {
	tepqld *TkExamPartitionQuestionLinkDelete
}

// Exec executes the deletion query.
func (tepqldo *TkExamPartitionQuestionLinkDeleteOne) Exec(ctx context.Context) error {
	n, err := tepqldo.tepqld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tkexampartitionquestionlink.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tepqldo *TkExamPartitionQuestionLinkDeleteOne) ExecX(ctx context.Context) {
	tepqldo.tepqld.ExecX(ctx)
}
