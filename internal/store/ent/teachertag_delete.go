// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/teachertag"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeacherTagDelete is the builder for deleting a TeacherTag entity.
type TeacherTagDelete struct {
	config
	hooks    []Hook
	mutation *TeacherTagMutation
}

// Where adds a new predicate to the TeacherTagDelete builder.
func (ttd *TeacherTagDelete) Where(ps ...predicate.TeacherTag) *TeacherTagDelete {
	ttd.mutation.predicates = append(ttd.mutation.predicates, ps...)
	return ttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttd *TeacherTagDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttd.hooks) == 0 {
		affected, err = ttd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeacherTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ttd.mutation = mutation
			affected, err = ttd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttd.hooks) - 1; i >= 0; i-- {
			mut = ttd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttd *TeacherTagDelete) ExecX(ctx context.Context) int {
	n, err := ttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttd *TeacherTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: teachertag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: teachertag.FieldID,
			},
		},
	}
	if ps := ttd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ttd.driver, _spec)
}

// TeacherTagDeleteOne is the builder for deleting a single TeacherTag entity.
type TeacherTagDeleteOne struct {
	ttd *TeacherTagDelete
}

// Exec executes the deletion query.
func (ttdo *TeacherTagDeleteOne) Exec(ctx context.Context) error {
	n, err := ttdo.ttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{teachertag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttdo *TeacherTagDeleteOne) ExecX(ctx context.Context) {
	ttdo.ttd.ExecX(ctx)
}
