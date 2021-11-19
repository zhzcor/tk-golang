// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/tkuserquestionbankrecord"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkUserQuestionBankRecordDelete is the builder for deleting a TkUserQuestionBankRecord entity.
type TkUserQuestionBankRecordDelete struct {
	config
	hooks    []Hook
	mutation *TkUserQuestionBankRecordMutation
}

// Where adds a new predicate to the TkUserQuestionBankRecordDelete builder.
func (tuqbrd *TkUserQuestionBankRecordDelete) Where(ps ...predicate.TkUserQuestionBankRecord) *TkUserQuestionBankRecordDelete {
	tuqbrd.mutation.predicates = append(tuqbrd.mutation.predicates, ps...)
	return tuqbrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tuqbrd *TkUserQuestionBankRecordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tuqbrd.hooks) == 0 {
		affected, err = tuqbrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkUserQuestionBankRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuqbrd.mutation = mutation
			affected, err = tuqbrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tuqbrd.hooks) - 1; i >= 0; i-- {
			mut = tuqbrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuqbrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuqbrd *TkUserQuestionBankRecordDelete) ExecX(ctx context.Context) int {
	n, err := tuqbrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tuqbrd *TkUserQuestionBankRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tkuserquestionbankrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkuserquestionbankrecord.FieldID,
			},
		},
	}
	if ps := tuqbrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tuqbrd.driver, _spec)
}

// TkUserQuestionBankRecordDeleteOne is the builder for deleting a single TkUserQuestionBankRecord entity.
type TkUserQuestionBankRecordDeleteOne struct {
	tuqbrd *TkUserQuestionBankRecordDelete
}

// Exec executes the deletion query.
func (tuqbrdo *TkUserQuestionBankRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := tuqbrdo.tuqbrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tkuserquestionbankrecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tuqbrdo *TkUserQuestionBankRecordDeleteOne) ExecX(ctx context.Context) {
	tuqbrdo.tuqbrd.ExecX(ctx)
}
