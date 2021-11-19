// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tkserver/internal/store/ent/message"
	"tkserver/internal/store/ent/messagetype"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageTypeCreate is the builder for creating a MessageType entity.
type MessageTypeCreate struct {
	config
	mutation *MessageTypeMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (mtc *MessageTypeCreate) SetUUID(s string) *MessageTypeCreate {
	mtc.mutation.SetUUID(s)
	return mtc
}

// SetCreatedAt sets the "created_at" field.
func (mtc *MessageTypeCreate) SetCreatedAt(t time.Time) *MessageTypeCreate {
	mtc.mutation.SetCreatedAt(t)
	return mtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mtc *MessageTypeCreate) SetNillableCreatedAt(t *time.Time) *MessageTypeCreate {
	if t != nil {
		mtc.SetCreatedAt(*t)
	}
	return mtc
}

// SetUpdatedAt sets the "updated_at" field.
func (mtc *MessageTypeCreate) SetUpdatedAt(t time.Time) *MessageTypeCreate {
	mtc.mutation.SetUpdatedAt(t)
	return mtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mtc *MessageTypeCreate) SetNillableUpdatedAt(t *time.Time) *MessageTypeCreate {
	if t != nil {
		mtc.SetUpdatedAt(*t)
	}
	return mtc
}

// SetDeletedAt sets the "deleted_at" field.
func (mtc *MessageTypeCreate) SetDeletedAt(t time.Time) *MessageTypeCreate {
	mtc.mutation.SetDeletedAt(t)
	return mtc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mtc *MessageTypeCreate) SetNillableDeletedAt(t *time.Time) *MessageTypeCreate {
	if t != nil {
		mtc.SetDeletedAt(*t)
	}
	return mtc
}

// SetName sets the "name" field.
func (mtc *MessageTypeCreate) SetName(s string) *MessageTypeCreate {
	mtc.mutation.SetName(s)
	return mtc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mtc *MessageTypeCreate) SetNillableName(s *string) *MessageTypeCreate {
	if s != nil {
		mtc.SetName(*s)
	}
	return mtc
}

// SetStatus sets the "status" field.
func (mtc *MessageTypeCreate) SetStatus(u uint8) *MessageTypeCreate {
	mtc.mutation.SetStatus(u)
	return mtc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mtc *MessageTypeCreate) SetNillableStatus(u *uint8) *MessageTypeCreate {
	if u != nil {
		mtc.SetStatus(*u)
	}
	return mtc
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (mtc *MessageTypeCreate) AddMessageIDs(ids ...int) *MessageTypeCreate {
	mtc.mutation.AddMessageIDs(ids...)
	return mtc
}

// AddMessages adds the "messages" edges to the Message entity.
func (mtc *MessageTypeCreate) AddMessages(m ...*Message) *MessageTypeCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mtc.AddMessageIDs(ids...)
}

// Mutation returns the MessageTypeMutation object of the builder.
func (mtc *MessageTypeCreate) Mutation() *MessageTypeMutation {
	return mtc.mutation
}

// Save creates the MessageType in the database.
func (mtc *MessageTypeCreate) Save(ctx context.Context) (*MessageType, error) {
	var (
		err  error
		node *MessageType
	)
	mtc.defaults()
	if len(mtc.hooks) == 0 {
		if err = mtc.check(); err != nil {
			return nil, err
		}
		node, err = mtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mtc.check(); err != nil {
				return nil, err
			}
			mtc.mutation = mutation
			node, err = mtc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mtc.hooks) - 1; i >= 0; i-- {
			mut = mtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mtc *MessageTypeCreate) SaveX(ctx context.Context) *MessageType {
	v, err := mtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (mtc *MessageTypeCreate) defaults() {
	if _, ok := mtc.mutation.CreatedAt(); !ok {
		v := messagetype.DefaultCreatedAt()
		mtc.mutation.SetCreatedAt(v)
	}
	if _, ok := mtc.mutation.UpdatedAt(); !ok {
		v := messagetype.DefaultUpdatedAt()
		mtc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mtc.mutation.Name(); !ok {
		v := messagetype.DefaultName
		mtc.mutation.SetName(v)
	}
	if _, ok := mtc.mutation.Status(); !ok {
		v := messagetype.DefaultStatus
		mtc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mtc *MessageTypeCreate) check() error {
	if _, ok := mtc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := mtc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := mtc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	return nil
}

func (mtc *MessageTypeCreate) sqlSave(ctx context.Context) (*MessageType, error) {
	_node, _spec := mtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mtc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mtc *MessageTypeCreate) createSpec() (*MessageType, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageType{config: mtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: messagetype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagetype.FieldID,
			},
		}
	)
	if value, ok := mtc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messagetype.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := mtc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messagetype.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := mtc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messagetype.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := mtc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messagetype.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := mtc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messagetype.FieldName,
		})
		_node.Name = value
	}
	if value, ok := mtc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: messagetype.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := mtc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   messagetype.MessagesTable,
			Columns: []string{messagetype.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MessageTypeCreateBulk is the builder for creating many MessageType entities in bulk.
type MessageTypeCreateBulk struct {
	config
	builders []*MessageTypeCreate
}

// Save creates the MessageType entities in the database.
func (mtcb *MessageTypeCreateBulk) Save(ctx context.Context) ([]*MessageType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mtcb.builders))
	nodes := make([]*MessageType, len(mtcb.builders))
	mutators := make([]Mutator, len(mtcb.builders))
	for i := range mtcb.builders {
		func(i int, root context.Context) {
			builder := mtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, mtcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mtcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mtcb *MessageTypeCreateBulk) SaveX(ctx context.Context) []*MessageType {
	v, err := mtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
