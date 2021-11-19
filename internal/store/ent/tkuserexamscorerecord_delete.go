// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/tkuserexamscorerecord"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkUserExamScoreRecordDelete is the builder for deleting a TkUserExamScoreRecord entity.
type TkUserExamScoreRecordDelete struct {
	config
	hooks    []Hook
	mutation *TkUserExamScoreRecordMutation
}

// Where adds a new predicate to the TkUserExamScoreRecordDelete builder.
func (tuesrd *TkUserExamScoreRecordDelete) Where(ps ...predicate.TkUserExamScoreRecord) *TkUserExamScoreRecordDelete {
	tuesrd.mutation.predicates = append(tuesrd.mutation.predicates, ps...)
	return tuesrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tuesrd *TkUserExamScoreRecordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tuesrd.hooks) == 0 {
		affected, err = tuesrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkUserExamScoreRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuesrd.mutation = mutation
			affected, err = tuesrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tuesrd.hooks) - 1; i >= 0; i-- {
			mut = tuesrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuesrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuesrd *TkUserExamScoreRecordDelete) ExecX(ctx context.Context) int {
	n, err := tuesrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tuesrd *TkUserExamScoreRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tkuserexamscorerecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkuserexamscorerecord.FieldID,
			},
		},
	}
	if ps := tuesrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tuesrd.driver, _spec)
}

// TkUserExamScoreRecordDeleteOne is the builder for deleting a single TkUserExamScoreRecord entity.
type TkUserExamScoreRecordDeleteOne struct {
	tuesrd *TkUserExamScoreRecordDelete
}

// Exec executes the deletion query.
func (tuesrdo *TkUserExamScoreRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := tuesrdo.tuesrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tkuserexamscorerecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tuesrdo *TkUserExamScoreRecordDeleteOne) ExecX(ctx context.Context) {
	tuesrdo.tuesrd.ExecX(ctx)
}
