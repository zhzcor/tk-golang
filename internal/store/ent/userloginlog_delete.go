// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/userloginlog"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserLoginLogDelete is the builder for deleting a UserLoginLog entity.
type UserLoginLogDelete struct {
	config
	hooks    []Hook
	mutation *UserLoginLogMutation
}

// Where adds a new predicate to the UserLoginLogDelete builder.
func (ulld *UserLoginLogDelete) Where(ps ...predicate.UserLoginLog) *UserLoginLogDelete {
	ulld.mutation.predicates = append(ulld.mutation.predicates, ps...)
	return ulld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ulld *UserLoginLogDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ulld.hooks) == 0 {
		affected, err = ulld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserLoginLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ulld.mutation = mutation
			affected, err = ulld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ulld.hooks) - 1; i >= 0; i-- {
			mut = ulld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ulld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulld *UserLoginLogDelete) ExecX(ctx context.Context) int {
	n, err := ulld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ulld *UserLoginLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userloginlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userloginlog.FieldID,
			},
		},
	}
	if ps := ulld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ulld.driver, _spec)
}

// UserLoginLogDeleteOne is the builder for deleting a single UserLoginLog entity.
type UserLoginLogDeleteOne struct {
	ulld *UserLoginLogDelete
}

// Exec executes the deletion query.
func (ulldo *UserLoginLogDeleteOne) Exec(ctx context.Context) error {
	n, err := ulldo.ulld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userloginlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ulldo *UserLoginLogDeleteOne) ExecX(ctx context.Context) {
	ulldo.ulld.ExecX(ctx)
}
