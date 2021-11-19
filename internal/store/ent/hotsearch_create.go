// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tkserver/internal/store/ent/hotsearch"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HotSearchCreate is the builder for creating a HotSearch entity.
type HotSearchCreate struct {
	config
	mutation *HotSearchMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (hsc *HotSearchCreate) SetUUID(s string) *HotSearchCreate {
	hsc.mutation.SetUUID(s)
	return hsc
}

// SetCreatedAt sets the "created_at" field.
func (hsc *HotSearchCreate) SetCreatedAt(t time.Time) *HotSearchCreate {
	hsc.mutation.SetCreatedAt(t)
	return hsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hsc *HotSearchCreate) SetNillableCreatedAt(t *time.Time) *HotSearchCreate {
	if t != nil {
		hsc.SetCreatedAt(*t)
	}
	return hsc
}

// SetUpdatedAt sets the "updated_at" field.
func (hsc *HotSearchCreate) SetUpdatedAt(t time.Time) *HotSearchCreate {
	hsc.mutation.SetUpdatedAt(t)
	return hsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hsc *HotSearchCreate) SetNillableUpdatedAt(t *time.Time) *HotSearchCreate {
	if t != nil {
		hsc.SetUpdatedAt(*t)
	}
	return hsc
}

// SetDeletedAt sets the "deleted_at" field.
func (hsc *HotSearchCreate) SetDeletedAt(t time.Time) *HotSearchCreate {
	hsc.mutation.SetDeletedAt(t)
	return hsc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (hsc *HotSearchCreate) SetNillableDeletedAt(t *time.Time) *HotSearchCreate {
	if t != nil {
		hsc.SetDeletedAt(*t)
	}
	return hsc
}

// SetName sets the "name" field.
func (hsc *HotSearchCreate) SetName(s string) *HotSearchCreate {
	hsc.mutation.SetName(s)
	return hsc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (hsc *HotSearchCreate) SetNillableName(s *string) *HotSearchCreate {
	if s != nil {
		hsc.SetName(*s)
	}
	return hsc
}

// SetStatus sets the "status" field.
func (hsc *HotSearchCreate) SetStatus(u uint8) *HotSearchCreate {
	hsc.mutation.SetStatus(u)
	return hsc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hsc *HotSearchCreate) SetNillableStatus(u *uint8) *HotSearchCreate {
	if u != nil {
		hsc.SetStatus(*u)
	}
	return hsc
}

// SetSearchCount sets the "search_count" field.
func (hsc *HotSearchCreate) SetSearchCount(i int) *HotSearchCreate {
	hsc.mutation.SetSearchCount(i)
	return hsc
}

// SetNillableSearchCount sets the "search_count" field if the given value is not nil.
func (hsc *HotSearchCreate) SetNillableSearchCount(i *int) *HotSearchCreate {
	if i != nil {
		hsc.SetSearchCount(*i)
	}
	return hsc
}

// Mutation returns the HotSearchMutation object of the builder.
func (hsc *HotSearchCreate) Mutation() *HotSearchMutation {
	return hsc.mutation
}

// Save creates the HotSearch in the database.
func (hsc *HotSearchCreate) Save(ctx context.Context) (*HotSearch, error) {
	var (
		err  error
		node *HotSearch
	)
	hsc.defaults()
	if len(hsc.hooks) == 0 {
		if err = hsc.check(); err != nil {
			return nil, err
		}
		node, err = hsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HotSearchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hsc.check(); err != nil {
				return nil, err
			}
			hsc.mutation = mutation
			node, err = hsc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hsc.hooks) - 1; i >= 0; i-- {
			mut = hsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hsc *HotSearchCreate) SaveX(ctx context.Context) *HotSearch {
	v, err := hsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (hsc *HotSearchCreate) defaults() {
	if _, ok := hsc.mutation.CreatedAt(); !ok {
		v := hotsearch.DefaultCreatedAt()
		hsc.mutation.SetCreatedAt(v)
	}
	if _, ok := hsc.mutation.UpdatedAt(); !ok {
		v := hotsearch.DefaultUpdatedAt()
		hsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := hsc.mutation.Name(); !ok {
		v := hotsearch.DefaultName
		hsc.mutation.SetName(v)
	}
	if _, ok := hsc.mutation.Status(); !ok {
		v := hotsearch.DefaultStatus
		hsc.mutation.SetStatus(v)
	}
	if _, ok := hsc.mutation.SearchCount(); !ok {
		v := hotsearch.DefaultSearchCount
		hsc.mutation.SetSearchCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hsc *HotSearchCreate) check() error {
	if _, ok := hsc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := hsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := hsc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := hsc.mutation.SearchCount(); !ok {
		return &ValidationError{Name: "search_count", err: errors.New("ent: missing required field \"search_count\"")}
	}
	return nil
}

func (hsc *HotSearchCreate) sqlSave(ctx context.Context) (*HotSearch, error) {
	_node, _spec := hsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hsc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hsc *HotSearchCreate) createSpec() (*HotSearch, *sqlgraph.CreateSpec) {
	var (
		_node = &HotSearch{config: hsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hotsearch.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hotsearch.FieldID,
			},
		}
	)
	if value, ok := hsc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hotsearch.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := hsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: hotsearch.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := hsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: hotsearch.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := hsc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: hotsearch.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := hsc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hotsearch.FieldName,
		})
		_node.Name = value
	}
	if value, ok := hsc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: hotsearch.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := hsc.mutation.SearchCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hotsearch.FieldSearchCount,
		})
		_node.SearchCount = value
	}
	return _node, _spec
}

// HotSearchCreateBulk is the builder for creating many HotSearch entities in bulk.
type HotSearchCreateBulk struct {
	config
	builders []*HotSearchCreate
}

// Save creates the HotSearch entities in the database.
func (hscb *HotSearchCreateBulk) Save(ctx context.Context) ([]*HotSearch, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hscb.builders))
	nodes := make([]*HotSearch, len(hscb.builders))
	mutators := make([]Mutator, len(hscb.builders))
	for i := range hscb.builders {
		func(i int, root context.Context) {
			builder := hscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HotSearchMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hscb *HotSearchCreateBulk) SaveX(ctx context.Context) []*HotSearch {
	v, err := hscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
