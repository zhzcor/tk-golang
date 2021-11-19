// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/tkuserrandomexamrecode"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkUserRandomExamRecodeDelete is the builder for deleting a TkUserRandomExamRecode entity.
type TkUserRandomExamRecodeDelete struct {
	config
	hooks    []Hook
	mutation *TkUserRandomExamRecodeMutation
}

// Where adds a new predicate to the TkUserRandomExamRecodeDelete builder.
func (turerd *TkUserRandomExamRecodeDelete) Where(ps ...predicate.TkUserRandomExamRecode) *TkUserRandomExamRecodeDelete {
	turerd.mutation.predicates = append(turerd.mutation.predicates, ps...)
	return turerd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (turerd *TkUserRandomExamRecodeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(turerd.hooks) == 0 {
		affected, err = turerd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkUserRandomExamRecodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			turerd.mutation = mutation
			affected, err = turerd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(turerd.hooks) - 1; i >= 0; i-- {
			mut = turerd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, turerd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (turerd *TkUserRandomExamRecodeDelete) ExecX(ctx context.Context) int {
	n, err := turerd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (turerd *TkUserRandomExamRecodeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tkuserrandomexamrecode.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkuserrandomexamrecode.FieldID,
			},
		},
	}
	if ps := turerd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, turerd.driver, _spec)
}

// TkUserRandomExamRecodeDeleteOne is the builder for deleting a single TkUserRandomExamRecode entity.
type TkUserRandomExamRecodeDeleteOne struct {
	turerd *TkUserRandomExamRecodeDelete
}

// Exec executes the deletion query.
func (turerdo *TkUserRandomExamRecodeDeleteOne) Exec(ctx context.Context) error {
	n, err := turerdo.turerd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tkuserrandomexamrecode.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (turerdo *TkUserRandomExamRecodeDeleteOne) ExecX(ctx context.Context) {
	turerdo.turerd.ExecX(ctx)
}
