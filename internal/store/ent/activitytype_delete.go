// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"tkserver/internal/store/ent/activitytype"
	"tkserver/internal/store/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ActivityTypeDelete is the builder for deleting a ActivityType entity.
type ActivityTypeDelete struct {
	config
	hooks    []Hook
	mutation *ActivityTypeMutation
}

// Where adds a new predicate to the ActivityTypeDelete builder.
func (atd *ActivityTypeDelete) Where(ps ...predicate.ActivityType) *ActivityTypeDelete {
	atd.mutation.predicates = append(atd.mutation.predicates, ps...)
	return atd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atd *ActivityTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atd.hooks) == 0 {
		affected, err = atd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atd.mutation = mutation
			affected, err = atd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atd.hooks) - 1; i >= 0; i-- {
			mut = atd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (atd *ActivityTypeDelete) ExecX(ctx context.Context) int {
	n, err := atd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atd *ActivityTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: activitytype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: activitytype.FieldID,
			},
		},
	}
	if ps := atd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, atd.driver, _spec)
}

// ActivityTypeDeleteOne is the builder for deleting a single ActivityType entity.
type ActivityTypeDeleteOne struct {
	atd *ActivityTypeDelete
}

// Exec executes the deletion query.
func (atdo *ActivityTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := atdo.atd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{activitytype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atdo *ActivityTypeDeleteOne) ExecX(ctx context.Context) {
	atdo.atd.ExecX(ctx)
}
