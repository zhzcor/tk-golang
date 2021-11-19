// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tkserver/internal/store/ent/kcclass"
	"tkserver/internal/store/ent/kcuserclass"
	"tkserver/internal/store/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KcUserClassCreate is the builder for creating a KcUserClass entity.
type KcUserClassCreate struct {
	config
	mutation *KcUserClassMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (kucc *KcUserClassCreate) SetUUID(s string) *KcUserClassCreate {
	kucc.mutation.SetUUID(s)
	return kucc
}

// SetCreatedAt sets the "created_at" field.
func (kucc *KcUserClassCreate) SetCreatedAt(t time.Time) *KcUserClassCreate {
	kucc.mutation.SetCreatedAt(t)
	return kucc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kucc *KcUserClassCreate) SetNillableCreatedAt(t *time.Time) *KcUserClassCreate {
	if t != nil {
		kucc.SetCreatedAt(*t)
	}
	return kucc
}

// SetUpdatedAt sets the "updated_at" field.
func (kucc *KcUserClassCreate) SetUpdatedAt(t time.Time) *KcUserClassCreate {
	kucc.mutation.SetUpdatedAt(t)
	return kucc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kucc *KcUserClassCreate) SetNillableUpdatedAt(t *time.Time) *KcUserClassCreate {
	if t != nil {
		kucc.SetUpdatedAt(*t)
	}
	return kucc
}

// SetDeletedAt sets the "deleted_at" field.
func (kucc *KcUserClassCreate) SetDeletedAt(t time.Time) *KcUserClassCreate {
	kucc.mutation.SetDeletedAt(t)
	return kucc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kucc *KcUserClassCreate) SetNillableDeletedAt(t *time.Time) *KcUserClassCreate {
	if t != nil {
		kucc.SetDeletedAt(*t)
	}
	return kucc
}

// SetPeriodType sets the "period_type" field.
func (kucc *KcUserClassCreate) SetPeriodType(u uint8) *KcUserClassCreate {
	kucc.mutation.SetPeriodType(u)
	return kucc
}

// SetNillablePeriodType sets the "period_type" field if the given value is not nil.
func (kucc *KcUserClassCreate) SetNillablePeriodType(u *uint8) *KcUserClassCreate {
	if u != nil {
		kucc.SetPeriodType(*u)
	}
	return kucc
}

// SetClosingDate sets the "closing_date" field.
func (kucc *KcUserClassCreate) SetClosingDate(t time.Time) *KcUserClassCreate {
	kucc.mutation.SetClosingDate(t)
	return kucc
}

// SetNillableClosingDate sets the "closing_date" field if the given value is not nil.
func (kucc *KcUserClassCreate) SetNillableClosingDate(t *time.Time) *KcUserClassCreate {
	if t != nil {
		kucc.SetClosingDate(*t)
	}
	return kucc
}

// SetStudyRate sets the "study_rate" field.
func (kucc *KcUserClassCreate) SetStudyRate(f float64) *KcUserClassCreate {
	kucc.mutation.SetStudyRate(f)
	return kucc
}

// SetNillableStudyRate sets the "study_rate" field if the given value is not nil.
func (kucc *KcUserClassCreate) SetNillableStudyRate(f *float64) *KcUserClassCreate {
	if f != nil {
		kucc.SetStudyRate(*f)
	}
	return kucc
}

// SetRemark sets the "remark" field.
func (kucc *KcUserClassCreate) SetRemark(s string) *KcUserClassCreate {
	kucc.mutation.SetRemark(s)
	return kucc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (kucc *KcUserClassCreate) SetNillableRemark(s *string) *KcUserClassCreate {
	if s != nil {
		kucc.SetRemark(*s)
	}
	return kucc
}

// SetPrice sets the "price" field.
func (kucc *KcUserClassCreate) SetPrice(f float64) *KcUserClassCreate {
	kucc.mutation.SetPrice(f)
	return kucc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (kucc *KcUserClassCreate) SetNillablePrice(f *float64) *KcUserClassCreate {
	if f != nil {
		kucc.SetPrice(*f)
	}
	return kucc
}

// SetUserID sets the "user_id" field.
func (kucc *KcUserClassCreate) SetUserID(i int) *KcUserClassCreate {
	kucc.mutation.SetUserID(i)
	return kucc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (kucc *KcUserClassCreate) SetNillableUserID(i *int) *KcUserClassCreate {
	if i != nil {
		kucc.SetUserID(*i)
	}
	return kucc
}

// SetClassID sets the "class_id" field.
func (kucc *KcUserClassCreate) SetClassID(i int) *KcUserClassCreate {
	kucc.mutation.SetClassID(i)
	return kucc
}

// SetNillableClassID sets the "class_id" field if the given value is not nil.
func (kucc *KcUserClassCreate) SetNillableClassID(i *int) *KcUserClassCreate {
	if i != nil {
		kucc.SetClassID(*i)
	}
	return kucc
}

// SetUser sets the "user" edge to the User entity.
func (kucc *KcUserClassCreate) SetUser(u *User) *KcUserClassCreate {
	return kucc.SetUserID(u.ID)
}

// SetClass sets the "class" edge to the KcClass entity.
func (kucc *KcUserClassCreate) SetClass(k *KcClass) *KcUserClassCreate {
	return kucc.SetClassID(k.ID)
}

// Mutation returns the KcUserClassMutation object of the builder.
func (kucc *KcUserClassCreate) Mutation() *KcUserClassMutation {
	return kucc.mutation
}

// Save creates the KcUserClass in the database.
func (kucc *KcUserClassCreate) Save(ctx context.Context) (*KcUserClass, error) {
	var (
		err  error
		node *KcUserClass
	)
	kucc.defaults()
	if len(kucc.hooks) == 0 {
		if err = kucc.check(); err != nil {
			return nil, err
		}
		node, err = kucc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KcUserClassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kucc.check(); err != nil {
				return nil, err
			}
			kucc.mutation = mutation
			node, err = kucc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kucc.hooks) - 1; i >= 0; i-- {
			mut = kucc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kucc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kucc *KcUserClassCreate) SaveX(ctx context.Context) *KcUserClass {
	v, err := kucc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (kucc *KcUserClassCreate) defaults() {
	if _, ok := kucc.mutation.CreatedAt(); !ok {
		v := kcuserclass.DefaultCreatedAt()
		kucc.mutation.SetCreatedAt(v)
	}
	if _, ok := kucc.mutation.UpdatedAt(); !ok {
		v := kcuserclass.DefaultUpdatedAt()
		kucc.mutation.SetUpdatedAt(v)
	}
	if _, ok := kucc.mutation.PeriodType(); !ok {
		v := kcuserclass.DefaultPeriodType
		kucc.mutation.SetPeriodType(v)
	}
	if _, ok := kucc.mutation.StudyRate(); !ok {
		v := kcuserclass.DefaultStudyRate
		kucc.mutation.SetStudyRate(v)
	}
	if _, ok := kucc.mutation.Remark(); !ok {
		v := kcuserclass.DefaultRemark
		kucc.mutation.SetRemark(v)
	}
	if _, ok := kucc.mutation.Price(); !ok {
		v := kcuserclass.DefaultPrice
		kucc.mutation.SetPrice(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kucc *KcUserClassCreate) check() error {
	if _, ok := kucc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := kucc.mutation.PeriodType(); !ok {
		return &ValidationError{Name: "period_type", err: errors.New("ent: missing required field \"period_type\"")}
	}
	if _, ok := kucc.mutation.StudyRate(); !ok {
		return &ValidationError{Name: "study_rate", err: errors.New("ent: missing required field \"study_rate\"")}
	}
	if _, ok := kucc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New("ent: missing required field \"remark\"")}
	}
	if _, ok := kucc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	return nil
}

func (kucc *KcUserClassCreate) sqlSave(ctx context.Context) (*KcUserClass, error) {
	_node, _spec := kucc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kucc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (kucc *KcUserClassCreate) createSpec() (*KcUserClass, *sqlgraph.CreateSpec) {
	var (
		_node = &KcUserClass{config: kucc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kcuserclass.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcuserclass.FieldID,
			},
		}
	)
	if value, ok := kucc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcuserclass.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := kucc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcuserclass.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := kucc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcuserclass.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := kucc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcuserclass.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := kucc.mutation.PeriodType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcuserclass.FieldPeriodType,
		})
		_node.PeriodType = value
	}
	if value, ok := kucc.mutation.ClosingDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcuserclass.FieldClosingDate,
		})
		_node.ClosingDate = value
	}
	if value, ok := kucc.mutation.StudyRate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: kcuserclass.FieldStudyRate,
		})
		_node.StudyRate = value
	}
	if value, ok := kucc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcuserclass.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := kucc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: kcuserclass.FieldPrice,
		})
		_node.Price = value
	}
	if nodes := kucc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcuserclass.UserTable,
			Columns: []string{kcuserclass.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kucc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcuserclass.ClassTable,
			Columns: []string{kcuserclass.ClassColumn},
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

// KcUserClassCreateBulk is the builder for creating many KcUserClass entities in bulk.
type KcUserClassCreateBulk struct {
	config
	builders []*KcUserClassCreate
}

// Save creates the KcUserClass entities in the database.
func (kuccb *KcUserClassCreateBulk) Save(ctx context.Context) ([]*KcUserClass, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kuccb.builders))
	nodes := make([]*KcUserClass, len(kuccb.builders))
	mutators := make([]Mutator, len(kuccb.builders))
	for i := range kuccb.builders {
		func(i int, root context.Context) {
			builder := kuccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KcUserClassMutation)
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
					_, err = mutators[i+1].Mutate(root, kuccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kuccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kuccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kuccb *KcUserClassCreateBulk) SaveX(ctx context.Context) []*KcUserClass {
	v, err := kuccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
