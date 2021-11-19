// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/attachment"
	"gserver/internal/store/ent/useraskanswer"
	"gserver/internal/store/ent/useraskanswerattachment"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserAskAnswerAttachmentCreate is the builder for creating a UserAskAnswerAttachment entity.
type UserAskAnswerAttachmentCreate struct {
	config
	mutation *UserAskAnswerAttachmentMutation
	hooks    []Hook
}

// SetAttachmentID sets the "attachment_id" field.
func (uaaac *UserAskAnswerAttachmentCreate) SetAttachmentID(i int) *UserAskAnswerAttachmentCreate {
	uaaac.mutation.SetAttachmentID(i)
	return uaaac
}

// SetNillableAttachmentID sets the "attachment_id" field if the given value is not nil.
func (uaaac *UserAskAnswerAttachmentCreate) SetNillableAttachmentID(i *int) *UserAskAnswerAttachmentCreate {
	if i != nil {
		uaaac.SetAttachmentID(*i)
	}
	return uaaac
}

// SetAskID sets the "ask_id" field.
func (uaaac *UserAskAnswerAttachmentCreate) SetAskID(i int) *UserAskAnswerAttachmentCreate {
	uaaac.mutation.SetAskID(i)
	return uaaac
}

// SetNillableAskID sets the "ask_id" field if the given value is not nil.
func (uaaac *UserAskAnswerAttachmentCreate) SetNillableAskID(i *int) *UserAskAnswerAttachmentCreate {
	if i != nil {
		uaaac.SetAskID(*i)
	}
	return uaaac
}

// SetType sets the "type" field.
func (uaaac *UserAskAnswerAttachmentCreate) SetType(u uint8) *UserAskAnswerAttachmentCreate {
	uaaac.mutation.SetType(u)
	return uaaac
}

// SetNillableType sets the "type" field if the given value is not nil.
func (uaaac *UserAskAnswerAttachmentCreate) SetNillableType(u *uint8) *UserAskAnswerAttachmentCreate {
	if u != nil {
		uaaac.SetType(*u)
	}
	return uaaac
}

// SetAttachment sets the "attachment" edge to the Attachment entity.
func (uaaac *UserAskAnswerAttachmentCreate) SetAttachment(a *Attachment) *UserAskAnswerAttachmentCreate {
	return uaaac.SetAttachmentID(a.ID)
}

// SetAsk sets the "ask" edge to the UserAskAnswer entity.
func (uaaac *UserAskAnswerAttachmentCreate) SetAsk(u *UserAskAnswer) *UserAskAnswerAttachmentCreate {
	return uaaac.SetAskID(u.ID)
}

// Mutation returns the UserAskAnswerAttachmentMutation object of the builder.
func (uaaac *UserAskAnswerAttachmentCreate) Mutation() *UserAskAnswerAttachmentMutation {
	return uaaac.mutation
}

// Save creates the UserAskAnswerAttachment in the database.
func (uaaac *UserAskAnswerAttachmentCreate) Save(ctx context.Context) (*UserAskAnswerAttachment, error) {
	var (
		err  error
		node *UserAskAnswerAttachment
	)
	uaaac.defaults()
	if len(uaaac.hooks) == 0 {
		if err = uaaac.check(); err != nil {
			return nil, err
		}
		node, err = uaaac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserAskAnswerAttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uaaac.check(); err != nil {
				return nil, err
			}
			uaaac.mutation = mutation
			node, err = uaaac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uaaac.hooks) - 1; i >= 0; i-- {
			mut = uaaac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uaaac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uaaac *UserAskAnswerAttachmentCreate) SaveX(ctx context.Context) *UserAskAnswerAttachment {
	v, err := uaaac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uaaac *UserAskAnswerAttachmentCreate) defaults() {
	if _, ok := uaaac.mutation.GetType(); !ok {
		v := useraskanswerattachment.DefaultType
		uaaac.mutation.SetType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uaaac *UserAskAnswerAttachmentCreate) check() error {
	if _, ok := uaaac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	return nil
}

func (uaaac *UserAskAnswerAttachmentCreate) sqlSave(ctx context.Context) (*UserAskAnswerAttachment, error) {
	_node, _spec := uaaac.createSpec()
	if err := sqlgraph.CreateNode(ctx, uaaac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uaaac *UserAskAnswerAttachmentCreate) createSpec() (*UserAskAnswerAttachment, *sqlgraph.CreateSpec) {
	var (
		_node = &UserAskAnswerAttachment{config: uaaac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: useraskanswerattachment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: useraskanswerattachment.FieldID,
			},
		}
	)
	if value, ok := uaaac.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: useraskanswerattachment.FieldType,
		})
		_node.Type = value
	}
	if nodes := uaaac.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useraskanswerattachment.AttachmentTable,
			Columns: []string{useraskanswerattachment.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AttachmentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uaaac.mutation.AskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useraskanswerattachment.AskTable,
			Columns: []string{useraskanswerattachment.AskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useraskanswer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AskID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserAskAnswerAttachmentCreateBulk is the builder for creating many UserAskAnswerAttachment entities in bulk.
type UserAskAnswerAttachmentCreateBulk struct {
	config
	builders []*UserAskAnswerAttachmentCreate
}

// Save creates the UserAskAnswerAttachment entities in the database.
func (uaaacb *UserAskAnswerAttachmentCreateBulk) Save(ctx context.Context) ([]*UserAskAnswerAttachment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uaaacb.builders))
	nodes := make([]*UserAskAnswerAttachment, len(uaaacb.builders))
	mutators := make([]Mutator, len(uaaacb.builders))
	for i := range uaaacb.builders {
		func(i int, root context.Context) {
			builder := uaaacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserAskAnswerAttachmentMutation)
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
					_, err = mutators[i+1].Mutate(root, uaaacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uaaacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uaaacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uaaacb *UserAskAnswerAttachmentCreateBulk) SaveX(ctx context.Context) []*UserAskAnswerAttachment {
	v, err := uaaacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
