// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/kccoursesmallcategory"
	"gserver/internal/store/ent/kcsmallcategoryexampaper"
	"gserver/internal/store/ent/tkexampaper"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KcSmallCategoryExamPaperCreate is the builder for creating a KcSmallCategoryExamPaper entity.
type KcSmallCategoryExamPaperCreate struct {
	config
	mutation *KcSmallCategoryExamPaperMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (kscepc *KcSmallCategoryExamPaperCreate) SetUUID(s string) *KcSmallCategoryExamPaperCreate {
	kscepc.mutation.SetUUID(s)
	return kscepc
}

// SetCreatedAt sets the "created_at" field.
func (kscepc *KcSmallCategoryExamPaperCreate) SetCreatedAt(t time.Time) *KcSmallCategoryExamPaperCreate {
	kscepc.mutation.SetCreatedAt(t)
	return kscepc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kscepc *KcSmallCategoryExamPaperCreate) SetNillableCreatedAt(t *time.Time) *KcSmallCategoryExamPaperCreate {
	if t != nil {
		kscepc.SetCreatedAt(*t)
	}
	return kscepc
}

// SetUpdatedAt sets the "updated_at" field.
func (kscepc *KcSmallCategoryExamPaperCreate) SetUpdatedAt(t time.Time) *KcSmallCategoryExamPaperCreate {
	kscepc.mutation.SetUpdatedAt(t)
	return kscepc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kscepc *KcSmallCategoryExamPaperCreate) SetNillableUpdatedAt(t *time.Time) *KcSmallCategoryExamPaperCreate {
	if t != nil {
		kscepc.SetUpdatedAt(*t)
	}
	return kscepc
}

// SetDeletedAt sets the "deleted_at" field.
func (kscepc *KcSmallCategoryExamPaperCreate) SetDeletedAt(t time.Time) *KcSmallCategoryExamPaperCreate {
	kscepc.mutation.SetDeletedAt(t)
	return kscepc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kscepc *KcSmallCategoryExamPaperCreate) SetNillableDeletedAt(t *time.Time) *KcSmallCategoryExamPaperCreate {
	if t != nil {
		kscepc.SetDeletedAt(*t)
	}
	return kscepc
}

// SetType sets the "type" field.
func (kscepc *KcSmallCategoryExamPaperCreate) SetType(u uint8) *KcSmallCategoryExamPaperCreate {
	kscepc.mutation.SetType(u)
	return kscepc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (kscepc *KcSmallCategoryExamPaperCreate) SetNillableType(u *uint8) *KcSmallCategoryExamPaperCreate {
	if u != nil {
		kscepc.SetType(*u)
	}
	return kscepc
}

// SetExamPaperID sets the "exam_paper_id" field.
func (kscepc *KcSmallCategoryExamPaperCreate) SetExamPaperID(i int) *KcSmallCategoryExamPaperCreate {
	kscepc.mutation.SetExamPaperID(i)
	return kscepc
}

// SetNillableExamPaperID sets the "exam_paper_id" field if the given value is not nil.
func (kscepc *KcSmallCategoryExamPaperCreate) SetNillableExamPaperID(i *int) *KcSmallCategoryExamPaperCreate {
	if i != nil {
		kscepc.SetExamPaperID(*i)
	}
	return kscepc
}

// SetSmallCategoryID sets the "small_category_id" field.
func (kscepc *KcSmallCategoryExamPaperCreate) SetSmallCategoryID(i int) *KcSmallCategoryExamPaperCreate {
	kscepc.mutation.SetSmallCategoryID(i)
	return kscepc
}

// SetNillableSmallCategoryID sets the "small_category_id" field if the given value is not nil.
func (kscepc *KcSmallCategoryExamPaperCreate) SetNillableSmallCategoryID(i *int) *KcSmallCategoryExamPaperCreate {
	if i != nil {
		kscepc.SetSmallCategoryID(*i)
	}
	return kscepc
}

// SetExamPaper sets the "exam_paper" edge to the TkExamPaper entity.
func (kscepc *KcSmallCategoryExamPaperCreate) SetExamPaper(t *TkExamPaper) *KcSmallCategoryExamPaperCreate {
	return kscepc.SetExamPaperID(t.ID)
}

// SetCourseSmallCategoryID sets the "course_small_category" edge to the KcCourseSmallCategory entity by ID.
func (kscepc *KcSmallCategoryExamPaperCreate) SetCourseSmallCategoryID(id int) *KcSmallCategoryExamPaperCreate {
	kscepc.mutation.SetCourseSmallCategoryID(id)
	return kscepc
}

// SetNillableCourseSmallCategoryID sets the "course_small_category" edge to the KcCourseSmallCategory entity by ID if the given value is not nil.
func (kscepc *KcSmallCategoryExamPaperCreate) SetNillableCourseSmallCategoryID(id *int) *KcSmallCategoryExamPaperCreate {
	if id != nil {
		kscepc = kscepc.SetCourseSmallCategoryID(*id)
	}
	return kscepc
}

// SetCourseSmallCategory sets the "course_small_category" edge to the KcCourseSmallCategory entity.
func (kscepc *KcSmallCategoryExamPaperCreate) SetCourseSmallCategory(k *KcCourseSmallCategory) *KcSmallCategoryExamPaperCreate {
	return kscepc.SetCourseSmallCategoryID(k.ID)
}

// Mutation returns the KcSmallCategoryExamPaperMutation object of the builder.
func (kscepc *KcSmallCategoryExamPaperCreate) Mutation() *KcSmallCategoryExamPaperMutation {
	return kscepc.mutation
}

// Save creates the KcSmallCategoryExamPaper in the database.
func (kscepc *KcSmallCategoryExamPaperCreate) Save(ctx context.Context) (*KcSmallCategoryExamPaper, error) {
	var (
		err  error
		node *KcSmallCategoryExamPaper
	)
	kscepc.defaults()
	if len(kscepc.hooks) == 0 {
		if err = kscepc.check(); err != nil {
			return nil, err
		}
		node, err = kscepc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KcSmallCategoryExamPaperMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kscepc.check(); err != nil {
				return nil, err
			}
			kscepc.mutation = mutation
			node, err = kscepc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kscepc.hooks) - 1; i >= 0; i-- {
			mut = kscepc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kscepc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kscepc *KcSmallCategoryExamPaperCreate) SaveX(ctx context.Context) *KcSmallCategoryExamPaper {
	v, err := kscepc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (kscepc *KcSmallCategoryExamPaperCreate) defaults() {
	if _, ok := kscepc.mutation.CreatedAt(); !ok {
		v := kcsmallcategoryexampaper.DefaultCreatedAt()
		kscepc.mutation.SetCreatedAt(v)
	}
	if _, ok := kscepc.mutation.UpdatedAt(); !ok {
		v := kcsmallcategoryexampaper.DefaultUpdatedAt()
		kscepc.mutation.SetUpdatedAt(v)
	}
	if _, ok := kscepc.mutation.GetType(); !ok {
		v := kcsmallcategoryexampaper.DefaultType
		kscepc.mutation.SetType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kscepc *KcSmallCategoryExamPaperCreate) check() error {
	if _, ok := kscepc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := kscepc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	return nil
}

func (kscepc *KcSmallCategoryExamPaperCreate) sqlSave(ctx context.Context) (*KcSmallCategoryExamPaper, error) {
	_node, _spec := kscepc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kscepc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (kscepc *KcSmallCategoryExamPaperCreate) createSpec() (*KcSmallCategoryExamPaper, *sqlgraph.CreateSpec) {
	var (
		_node = &KcSmallCategoryExamPaper{config: kscepc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kcsmallcategoryexampaper.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcsmallcategoryexampaper.FieldID,
			},
		}
	)
	if value, ok := kscepc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := kscepc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := kscepc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := kscepc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := kscepc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcsmallcategoryexampaper.FieldType,
		})
		_node.Type = value
	}
	if nodes := kscepc.mutation.ExamPaperIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryexampaper.ExamPaperTable,
			Columns: []string{kcsmallcategoryexampaper.ExamPaperColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkexampaper.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ExamPaperID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kscepc.mutation.CourseSmallCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcsmallcategoryexampaper.CourseSmallCategoryTable,
			Columns: []string{kcsmallcategoryexampaper.CourseSmallCategoryColumn},
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

// KcSmallCategoryExamPaperCreateBulk is the builder for creating many KcSmallCategoryExamPaper entities in bulk.
type KcSmallCategoryExamPaperCreateBulk struct {
	config
	builders []*KcSmallCategoryExamPaperCreate
}

// Save creates the KcSmallCategoryExamPaper entities in the database.
func (kscepcb *KcSmallCategoryExamPaperCreateBulk) Save(ctx context.Context) ([]*KcSmallCategoryExamPaper, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kscepcb.builders))
	nodes := make([]*KcSmallCategoryExamPaper, len(kscepcb.builders))
	mutators := make([]Mutator, len(kscepcb.builders))
	for i := range kscepcb.builders {
		func(i int, root context.Context) {
			builder := kscepcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KcSmallCategoryExamPaperMutation)
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
					_, err = mutators[i+1].Mutate(root, kscepcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kscepcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kscepcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kscepcb *KcSmallCategoryExamPaperCreateBulk) SaveX(ctx context.Context) []*KcSmallCategoryExamPaper {
	v, err := kscepcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
