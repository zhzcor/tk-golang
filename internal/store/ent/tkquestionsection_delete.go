// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/tkquestionsection"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkQuestionSectionDelete is the builder for deleting a TkQuestionSection entity.
type TkQuestionSectionDelete struct {
	config
	hooks    []Hook
	mutation *TkQuestionSectionMutation
}

// Where adds a new predicate to the TkQuestionSectionDelete builder.
func (tqsd *TkQuestionSectionDelete) Where(ps ...predicate.TkQuestionSection) *TkQuestionSectionDelete {
	tqsd.mutation.predicates = append(tqsd.mutation.predicates, ps...)
	return tqsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tqsd *TkQuestionSectionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tqsd.hooks) == 0 {
		affected, err = tqsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkQuestionSectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tqsd.mutation = mutation
			affected, err = tqsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tqsd.hooks) - 1; i >= 0; i-- {
			mut = tqsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tqsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqsd *TkQuestionSectionDelete) ExecX(ctx context.Context) int {
	n, err := tqsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tqsd *TkQuestionSectionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tkquestionsection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkquestionsection.FieldID,
			},
		},
	}
	if ps := tqsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tqsd.driver, _spec)
}

// TkQuestionSectionDeleteOne is the builder for deleting a single TkQuestionSection entity.
type TkQuestionSectionDeleteOne struct {
	tqsd *TkQuestionSectionDelete
}

// Exec executes the deletion query.
func (tqsdo *TkQuestionSectionDeleteOne) Exec(ctx context.Context) error {
	n, err := tqsdo.tqsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tkquestionsection.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tqsdo *TkQuestionSectionDeleteOne) ExecX(ctx context.Context) {
	tqsdo.tqsd.ExecX(ctx)
}
