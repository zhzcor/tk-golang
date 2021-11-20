// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"tkserver/internal/store/ent/groupcard"
	"tkserver/internal/store/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupCardDelete is the builder for deleting a GroupCard entity.
type GroupCardDelete struct {
	config
	hooks    []Hook
	mutation *GroupCardMutation
}

// Where adds a new predicate to the GroupCardDelete builder.
func (gcd *GroupCardDelete) Where(ps ...predicate.GroupCard) *GroupCardDelete {
	gcd.mutation.predicates = append(gcd.mutation.predicates, ps...)
	return gcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gcd *GroupCardDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gcd.hooks) == 0 {
		affected, err = gcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupCardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcd.mutation = mutation
			affected, err = gcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gcd.hooks) - 1; i >= 0; i-- {
			mut = gcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcd *GroupCardDelete) ExecX(ctx context.Context) int {
	n, err := gcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gcd *GroupCardDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: groupcard.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupcard.FieldID,
			},
		},
	}
	if ps := gcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gcd.driver, _spec)
}

// GroupCardDeleteOne is the builder for deleting a single GroupCard entity.
type GroupCardDeleteOne struct {
	gcd *GroupCardDelete
}

// Exec executes the deletion query.
func (gcdo *GroupCardDeleteOne) Exec(ctx context.Context) error {
	n, err := gcdo.gcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{groupcard.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gcdo *GroupCardDeleteOne) ExecX(ctx context.Context) {
	gcdo.gcd.ExecX(ctx)
}
