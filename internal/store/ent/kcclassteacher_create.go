// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/kcclass"
	"gserver/internal/store/ent/kcclassteacher"
	"gserver/internal/store/ent/teacher"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KcClassTeacherCreate is the builder for creating a KcClassTeacher entity.
type KcClassTeacherCreate struct {
	config
	mutation *KcClassTeacherMutation
	hooks    []Hook
}

// SetShowStatus sets the "show_status" field.
func (kctc *KcClassTeacherCreate) SetShowStatus(u uint8) *KcClassTeacherCreate {
	kctc.mutation.SetShowStatus(u)
	return kctc
}

// SetNillableShowStatus sets the "show_status" field if the given value is not nil.
func (kctc *KcClassTeacherCreate) SetNillableShowStatus(u *uint8) *KcClassTeacherCreate {
	if u != nil {
		kctc.SetShowStatus(*u)
	}
	return kctc
}

// SetSortOrder sets the "sort_order" field.
func (kctc *KcClassTeacherCreate) SetSortOrder(i int) *KcClassTeacherCreate {
	kctc.mutation.SetSortOrder(i)
	return kctc
}

// SetNillableSortOrder sets the "sort_order" field if the given value is not nil.
func (kctc *KcClassTeacherCreate) SetNillableSortOrder(i *int) *KcClassTeacherCreate {
	if i != nil {
		kctc.SetSortOrder(*i)
	}
	return kctc
}

// SetClassID sets the "class_id" field.
func (kctc *KcClassTeacherCreate) SetClassID(i int) *KcClassTeacherCreate {
	kctc.mutation.SetClassID(i)
	return kctc
}

// SetNillableClassID sets the "class_id" field if the given value is not nil.
func (kctc *KcClassTeacherCreate) SetNillableClassID(i *int) *KcClassTeacherCreate {
	if i != nil {
		kctc.SetClassID(*i)
	}
	return kctc
}

// SetTeacherID sets the "teacher_id" field.
func (kctc *KcClassTeacherCreate) SetTeacherID(i int) *KcClassTeacherCreate {
	kctc.mutation.SetTeacherID(i)
	return kctc
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (kctc *KcClassTeacherCreate) SetNillableTeacherID(i *int) *KcClassTeacherCreate {
	if i != nil {
		kctc.SetTeacherID(*i)
	}
	return kctc
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (kctc *KcClassTeacherCreate) SetTeacher(t *Teacher) *KcClassTeacherCreate {
	return kctc.SetTeacherID(t.ID)
}

// SetClass sets the "class" edge to the KcClass entity.
func (kctc *KcClassTeacherCreate) SetClass(k *KcClass) *KcClassTeacherCreate {
	return kctc.SetClassID(k.ID)
}

// Mutation returns the KcClassTeacherMutation object of the builder.
func (kctc *KcClassTeacherCreate) Mutation() *KcClassTeacherMutation {
	return kctc.mutation
}

// Save creates the KcClassTeacher in the database.
func (kctc *KcClassTeacherCreate) Save(ctx context.Context) (*KcClassTeacher, error) {
	var (
		err  error
		node *KcClassTeacher
	)
	kctc.defaults()
	if len(kctc.hooks) == 0 {
		if err = kctc.check(); err != nil {
			return nil, err
		}
		node, err = kctc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KcClassTeacherMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kctc.check(); err != nil {
				return nil, err
			}
			kctc.mutation = mutation
			node, err = kctc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kctc.hooks) - 1; i >= 0; i-- {
			mut = kctc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kctc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kctc *KcClassTeacherCreate) SaveX(ctx context.Context) *KcClassTeacher {
	v, err := kctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (kctc *KcClassTeacherCreate) defaults() {
	if _, ok := kctc.mutation.ShowStatus(); !ok {
		v := kcclassteacher.DefaultShowStatus
		kctc.mutation.SetShowStatus(v)
	}
	if _, ok := kctc.mutation.SortOrder(); !ok {
		v := kcclassteacher.DefaultSortOrder
		kctc.mutation.SetSortOrder(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kctc *KcClassTeacherCreate) check() error {
	if _, ok := kctc.mutation.ShowStatus(); !ok {
		return &ValidationError{Name: "show_status", err: errors.New("ent: missing required field \"show_status\"")}
	}
	if _, ok := kctc.mutation.SortOrder(); !ok {
		return &ValidationError{Name: "sort_order", err: errors.New("ent: missing required field \"sort_order\"")}
	}
	return nil
}

func (kctc *KcClassTeacherCreate) sqlSave(ctx context.Context) (*KcClassTeacher, error) {
	_node, _spec := kctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kctc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (kctc *KcClassTeacherCreate) createSpec() (*KcClassTeacher, *sqlgraph.CreateSpec) {
	var (
		_node = &KcClassTeacher{config: kctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kcclassteacher.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcclassteacher.FieldID,
			},
		}
	)
	if value, ok := kctc.mutation.ShowStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcclassteacher.FieldShowStatus,
		})
		_node.ShowStatus = value
	}
	if value, ok := kctc.mutation.SortOrder(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: kcclassteacher.FieldSortOrder,
		})
		_node.SortOrder = value
	}
	if nodes := kctc.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcclassteacher.TeacherTable,
			Columns: []string{kcclassteacher.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TeacherID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kctc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcclassteacher.ClassTable,
			Columns: []string{kcclassteacher.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kcclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ClassID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KcClassTeacherCreateBulk is the builder for creating many KcClassTeacher entities in bulk.
type KcClassTeacherCreateBulk struct {
	config
	builders []*KcClassTeacherCreate
}

// Save creates the KcClassTeacher entities in the database.
func (kctcb *KcClassTeacherCreateBulk) Save(ctx context.Context) ([]*KcClassTeacher, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kctcb.builders))
	nodes := make([]*KcClassTeacher, len(kctcb.builders))
	mutators := make([]Mutator, len(kctcb.builders))
	for i := range kctcb.builders {
		func(i int, root context.Context) {
			builder := kctcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KcClassTeacherMutation)
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
					_, err = mutators[i+1].Mutate(root, kctcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kctcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kctcb *KcClassTeacherCreateBulk) SaveX(ctx context.Context) []*KcClassTeacher {
	v, err := kctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
