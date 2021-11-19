// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"tkserver/internal/store/ent/majordetail"
	"tkserver/internal/store/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MajorDetailDelete is the builder for deleting a MajorDetail entity.
type MajorDetailDelete struct {
	config
	hooks    []Hook
	mutation *MajorDetailMutation
}

// Where adds a new predicate to the MajorDetailDelete builder.
func (mdd *MajorDetailDelete) Where(ps ...predicate.MajorDetail) *MajorDetailDelete {
	mdd.mutation.predicates = append(mdd.mutation.predicates, ps...)
	return mdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mdd *MajorDetailDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mdd.hooks) == 0 {
		affected, err = mdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MajorDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mdd.mutation = mutation
			affected, err = mdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mdd.hooks) - 1; i >= 0; i-- {
			mut = mdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mdd *MajorDetailDelete) ExecX(ctx context.Context) int {
	n, err := mdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mdd *MajorDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: majordetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: majordetail.FieldID,
			},
		},
	}
	if ps := mdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mdd.driver, _spec)
}

// MajorDetailDeleteOne is the builder for deleting a single MajorDetail entity.
type MajorDetailDeleteOne struct {
	mdd *MajorDetailDelete
}

// Exec executes the deletion query.
func (mddo *MajorDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := mddo.mdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{majordetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mddo *MajorDetailDeleteOne) ExecX(ctx context.Context) {
	mddo.mdd.ExecX(ctx)
}
