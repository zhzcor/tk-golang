// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/videorecord"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoRecordDelete is the builder for deleting a VideoRecord entity.
type VideoRecordDelete struct {
	config
	hooks    []Hook
	mutation *VideoRecordMutation
}

// Where adds a new predicate to the VideoRecordDelete builder.
func (vrd *VideoRecordDelete) Where(ps ...predicate.VideoRecord) *VideoRecordDelete {
	vrd.mutation.predicates = append(vrd.mutation.predicates, ps...)
	return vrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vrd *VideoRecordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vrd.hooks) == 0 {
		affected, err = vrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VideoRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vrd.mutation = mutation
			affected, err = vrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vrd.hooks) - 1; i >= 0; i-- {
			mut = vrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vrd *VideoRecordDelete) ExecX(ctx context.Context) int {
	n, err := vrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vrd *VideoRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: videorecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: videorecord.FieldID,
			},
		},
	}
	if ps := vrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vrd.driver, _spec)
}

// VideoRecordDeleteOne is the builder for deleting a single VideoRecord entity.
type VideoRecordDeleteOne struct {
	vrd *VideoRecordDelete
}

// Exec executes the deletion query.
func (vrdo *VideoRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := vrdo.vrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{videorecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vrdo *VideoRecordDeleteOne) ExecX(ctx context.Context) {
	vrdo.vrd.ExecX(ctx)
}
