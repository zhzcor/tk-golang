// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tkserver/internal/store/ent/attachment"
	"tkserver/internal/store/ent/kccoursesmallcategory"
	"tkserver/internal/store/ent/kcsmallcategoryattachment"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KcSmallCategoryAttachmentCreate is the builder for creating a KcSmallCategoryAttachment entity.
type KcSmallCategoryAttachmentCreate struct {
	config
	mutation *KcSmallCategoryAttachmentMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (kscac *KcSmallCategoryAttachmentCreate) SetUUID(s string) *KcSmallCategoryAttachmentCreate {
	kscac.mutation.SetUUID(s)
	return kscac
}

// SetCreatedAt sets the "created_at" field.
func (kscac *KcSmallCategoryAttachmentCreate) SetCreatedAt(t time.Time) *KcSmallCategoryAttachmentCreate {
	kscac.mutation.SetCreatedAt(t)
	return kscac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kscac *KcSmallCategoryAttachmentCreate) SetNillableCreatedAt(t *time.Time) *KcSmallCategoryAttachmentCreate {
	if t != nil {
		kscac.SetCreatedAt(*t)
	}
	return kscac
}

// SetUpdatedAt sets the "updated_at" field.
func (kscac *KcSmallCategoryAttachmentCreate) SetUpdatedAt(t time.Time) *KcSmallCategoryAttachmentCreate {
	kscac.mutation.SetUpdatedAt(t)
	return kscac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kscac *KcSmallCategoryAttachmentCreate) SetNillableUpdatedAt(t *time.Time) *KcSmallCategoryAttachmentCreate {
	if t != nil {
		kscac.SetUpdatedAt(*t)
	}
	return kscac
}

// SetDeletedAt sets the "deleted_at" field.
func (kscac *KcSmallCategoryAttachmentCreate) SetDeletedAt(t time.Time) *KcSmallCategoryAttachmentCreate {
	kscac.mutation.SetDeletedAt(t)
	return kscac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kscac *KcSmallCategoryAttachmentCreate) SetNillableDeletedAt(t *time.Time) *KcSmallCategoryAttachmentCreate {
	if t != nil {
		kscac.SetDeletedAt(*t)
	}
	return kscac
}

// SetAttachmentID sets the "attachment_id" field.
func (kscac *KcSmallCategoryAttachmentCreate) SetAttachmentID(i int) *KcSmallCategoryAttachmentCreate {
	kscac.mutation.SetAttachmentID(i)
	return kscac
}

// SetNillableAttachmentID sets the "attachment_id" field if the given value is not nil.
func (kscac *KcSmallCategoryAttachmentCreate) SetNillableAttachmentID(i *int) *KcSmallCategoryAttachmentCreate {
	if i != nil {
		kscac.SetAttachmentID(*i)
	}
	return kscac
}

// SetAttachmentName sets the "attachment_name" field.
func (kscac *KcSmallCategoryAttachmentCreate) SetAttachmentName(s string) *KcSmallCategoryAttachmentCreate {
	kscac.mutation.SetAttachmentName(s)
	return kscac
}

// SetNillableAttachmentName sets the "attachment_name" field if the given value is not nil.
func (kscac *KcSmallCategoryAttachmentCreate) SetNillableAttachmentName(s *string) *KcSmallCategoryAttachmentCreate {
	if s != nil {
		kscac.SetAttachmentName(*s)
	}
	return kscac
}

// SetSmallCategoryID sets the "small_category_id" field.
func (kscac *KcSmallCategoryAttachmentCreate) SetSmallCategoryID(i int) *KcSmallCategoryAttachmentCreate {
	kscac.mutation.SetSmallCategoryID(i)
	return kscac
}

// SetNillableSmallCategoryID sets the "small_category_id" field if the given value is not nil.
func (kscac *KcSmallCategoryAttachmentCreate) SetNillableSmallCategoryID(i *int) *KcSmallCategoryAttachmentCreate {
	if i != nil {
		kscac.SetSmallCategoryID(*i)
	}
	return kscac
}

// SetAttachment sets the "attachment" edge to the Attachment entity.
func (kscac *KcSmallCategoryAttachmentCreate) SetAttachment(a *Attachment) *KcSmallCategoryAttachmentCreate {
	return kscac.SetAttachmentID(a.ID)
}

// SetSmallCategory sets the "small_category" edge to the KcCourseSmallCategory entity.
func (kscac *KcSmallCategoryAttachmentCreate) SetSmallCategory(k *KcCourseSmallCategory) *KcSmallCategoryAttachmentCreate {
	return kscac.SetSmallCategoryID(k.ID)
}

// Mutation returns the KcSmallCategoryAttachmentMutation object of the builder.
func (kscac *KcSmallCategoryAttachmentCreate) Mutation() *KcSmallCategoryAttachmentMutation {
	return kscac.mutation
}

// Save creates the KcSmallCategoryAttachment in the database.
func (kscac *KcSmallCategoryAttachmentCreate) Save(ctx context.Context) (*KcSmallCategoryAttachment, error) {
	var (
		err  error
		node *KcSmallCategoryAttachment
	)
	kscac.defaults()
	if len(kscac.hooks) == 0 {
		if err = kscac.check(); err != nil {
			return nil, err
		}
		node, err = kscac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KcSmallCategoryAttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kscac.check(); err != nil {
				return nil, err
			}
			kscac.mutation = mutation
			node, err = kscac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kscac.hooks) - 1; i >= 0; i-- {
			mut = kscac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kscac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kscac *KcSmallCategoryAttachmentCreate) SaveX(ctx context.Context) *KcSmallCategoryAttachment {
	v, err := kscac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (kscac *KcSmallCategoryAttachmentCreate) defaults() {
	if _, ok := kscac.mutation.CreatedAt(); !ok {
		v := kcsmallcategoryattachment.DefaultCreatedAt()
		kscac.mutation.SetCreatedAt(v)
	}
	if _, ok := kscac.mutation.UpdatedAt(); !ok {
		v := kcsmallcategoryattachment.DefaultUpdatedAt()
		kscac.mutation.SetUpdatedAt(v)
	}
	if _, ok := kscac.mutation.AttachmentName(); !ok {
		v := kcsmallcategoryattachment.DefaultAttachmentName
		kscac.mutation.SetAttachmentName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kscac *KcSmallCategoryAttachmentCreate) check() error {
	if _, ok := kscac.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := kscac.mutation.AttachmentName(); !ok {
		return &ValidationError{Name: "attachment_name", err: errors.New("ent: missing required field \"attachment_name\"")}
	}
	return nil
}

func (kscac *KcSmallCategoryAttachmentCreate) sqlSave(ctx context.Context) (*KcSmallCategoryAttachment, error) {
	_node, _spec := kscac.createSpec()
	if err := sqlgraph.CreateNode(ctx, kscac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (kscac *KcSmallCategoryAttachmentCreate) createSpec() (*KcSmallCategoryAttachment, *sqlgraph.CreateSpec) {
	var (
		_node = &KcSmallCategoryAttachment{config: kscac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kcsmallcategoryattachment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcsmallcategoryattachment.FieldID,
			},
		}
	)
	if value, ok := kscac.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := kscac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := kscac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := kscac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := kscac.mutation.AttachmentName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcsmallcategoryattachment.FieldAttachmentName,
		})
		_node.AttachmentName = value
	}
	if nodes := kscac.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryattachment.AttachmentTable,
			Columns: []string{kcsmallcategoryattachment.AttachmentColumn},
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
	if nodes := kscac.mutation.SmallCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryattachment.SmallCategoryTable,
			Columns: []string{kcsmallcategoryattachment.SmallCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccoursesmallcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SmallCategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KcSmallCategoryAttachmentCreateBulk is the builder for creating many KcSmallCategoryAttachment entities in bulk.
type KcSmallCategoryAttachmentCreateBulk struct {
	config
	builders []*KcSmallCategoryAttachmentCreate
}

// Save creates the KcSmallCategoryAttachment entities in the database.
func (kscacb *KcSmallCategoryAttachmentCreateBulk) Save(ctx context.Context) ([]*KcSmallCategoryAttachment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kscacb.builders))
	nodes := make([]*KcSmallCategoryAttachment, len(kscacb.builders))
	mutators := make([]Mutator, len(kscacb.builders))
	for i := range kscacb.builders {
		func(i int, root context.Context) {
			builder := kscacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KcSmallCategoryAttachmentMutation)
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
					_, err = mutators[i+1].Mutate(root, kscacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kscacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kscacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kscacb *KcSmallCategoryAttachmentCreateBulk) SaveX(ctx context.Context) []*KcSmallCategoryAttachment {
	v, err := kscacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
