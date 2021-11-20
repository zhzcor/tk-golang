// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"
	"tkserver/internal/store/ent/city"
	"tkserver/internal/store/ent/predicate"
	"tkserver/internal/store/ent/tkquestionbank"
	"tkserver/internal/store/ent/tkquestionbankcity"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkQuestionBankCityUpdate is the builder for updating TkQuestionBankCity entities.
type TkQuestionBankCityUpdate struct {
	config
	hooks    []Hook
	mutation *TkQuestionBankCityMutation
}

// Where adds a new predicate for the TkQuestionBankCityUpdate builder.
func (tqbcu *TkQuestionBankCityUpdate) Where(ps ...predicate.TkQuestionBankCity) *TkQuestionBankCityUpdate {
	tqbcu.mutation.predicates = append(tqbcu.mutation.predicates, ps...)
	return tqbcu
}

// SetUUID sets the "uuid" field.
func (tqbcu *TkQuestionBankCityUpdate) SetUUID(s string) *TkQuestionBankCityUpdate {
	tqbcu.mutation.SetUUID(s)
	return tqbcu
}

// SetUpdatedAt sets the "updated_at" field.
func (tqbcu *TkQuestionBankCityUpdate) SetUpdatedAt(t time.Time) *TkQuestionBankCityUpdate {
	tqbcu.mutation.SetUpdatedAt(t)
	return tqbcu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tqbcu *TkQuestionBankCityUpdate) ClearUpdatedAt() *TkQuestionBankCityUpdate {
	tqbcu.mutation.ClearUpdatedAt()
	return tqbcu
}

// SetDeletedAt sets the "deleted_at" field.
func (tqbcu *TkQuestionBankCityUpdate) SetDeletedAt(t time.Time) *TkQuestionBankCityUpdate {
	tqbcu.mutation.SetDeletedAt(t)
	return tqbcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tqbcu *TkQuestionBankCityUpdate) SetNillableDeletedAt(t *time.Time) *TkQuestionBankCityUpdate {
	if t != nil {
		tqbcu.SetDeletedAt(*t)
	}
	return tqbcu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tqbcu *TkQuestionBankCityUpdate) ClearDeletedAt() *TkQuestionBankCityUpdate {
	tqbcu.mutation.ClearDeletedAt()
	return tqbcu
}

// SetCityID sets the "city_id" field.
func (tqbcu *TkQuestionBankCityUpdate) SetCityID(i int) *TkQuestionBankCityUpdate {
	tqbcu.mutation.ResetCityID()
	tqbcu.mutation.SetCityID(i)
	return tqbcu
}

// SetNillableCityID sets the "city_id" field if the given value is not nil.
func (tqbcu *TkQuestionBankCityUpdate) SetNillableCityID(i *int) *TkQuestionBankCityUpdate {
	if i != nil {
		tqbcu.SetCityID(*i)
	}
	return tqbcu
}

// ClearCityID clears the value of the "city_id" field.
func (tqbcu *TkQuestionBankCityUpdate) ClearCityID() *TkQuestionBankCityUpdate {
	tqbcu.mutation.ClearCityID()
	return tqbcu
}

// SetQuestionBankID sets the "question_bank_id" field.
func (tqbcu *TkQuestionBankCityUpdate) SetQuestionBankID(i int) *TkQuestionBankCityUpdate {
	tqbcu.mutation.ResetQuestionBankID()
	tqbcu.mutation.SetQuestionBankID(i)
	return tqbcu
}

// SetNillableQuestionBankID sets the "question_bank_id" field if the given value is not nil.
func (tqbcu *TkQuestionBankCityUpdate) SetNillableQuestionBankID(i *int) *TkQuestionBankCityUpdate {
	if i != nil {
		tqbcu.SetQuestionBankID(*i)
	}
	return tqbcu
}

// ClearQuestionBankID clears the value of the "question_bank_id" field.
func (tqbcu *TkQuestionBankCityUpdate) ClearQuestionBankID() *TkQuestionBankCityUpdate {
	tqbcu.mutation.ClearQuestionBankID()
	return tqbcu
}

// SetTkQuestionBankID sets the "tk_question_bank" edge to the TkQuestionBank entity by ID.
func (tqbcu *TkQuestionBankCityUpdate) SetTkQuestionBankID(id int) *TkQuestionBankCityUpdate {
	tqbcu.mutation.SetTkQuestionBankID(id)
	return tqbcu
}

// SetNillableTkQuestionBankID sets the "tk_question_bank" edge to the TkQuestionBank entity by ID if the given value is not nil.
func (tqbcu *TkQuestionBankCityUpdate) SetNillableTkQuestionBankID(id *int) *TkQuestionBankCityUpdate {
	if id != nil {
		tqbcu = tqbcu.SetTkQuestionBankID(*id)
	}
	return tqbcu
}

// SetTkQuestionBank sets the "tk_question_bank" edge to the TkQuestionBank entity.
func (tqbcu *TkQuestionBankCityUpdate) SetTkQuestionBank(t *TkQuestionBank) *TkQuestionBankCityUpdate {
	return tqbcu.SetTkQuestionBankID(t.ID)
}

// SetCity sets the "city" edge to the City entity.
func (tqbcu *TkQuestionBankCityUpdate) SetCity(c *City) *TkQuestionBankCityUpdate {
	return tqbcu.SetCityID(c.ID)
}

// Mutation returns the TkQuestionBankCityMutation object of the builder.
func (tqbcu *TkQuestionBankCityUpdate) Mutation() *TkQuestionBankCityMutation {
	return tqbcu.mutation
}

// ClearTkQuestionBank clears the "tk_question_bank" edge to the TkQuestionBank entity.
func (tqbcu *TkQuestionBankCityUpdate) ClearTkQuestionBank() *TkQuestionBankCityUpdate {
	tqbcu.mutation.ClearTkQuestionBank()
	return tqbcu
}

// ClearCity clears the "city" edge to the City entity.
func (tqbcu *TkQuestionBankCityUpdate) ClearCity() *TkQuestionBankCityUpdate {
	tqbcu.mutation.ClearCity()
	return tqbcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tqbcu *TkQuestionBankCityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tqbcu.defaults()
	if len(tqbcu.hooks) == 0 {
		affected, err = tqbcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkQuestionBankCityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tqbcu.mutation = mutation
			affected, err = tqbcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tqbcu.hooks) - 1; i >= 0; i-- {
			mut = tqbcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tqbcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tqbcu *TkQuestionBankCityUpdate) SaveX(ctx context.Context) int {
	affected, err := tqbcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tqbcu *TkQuestionBankCityUpdate) Exec(ctx context.Context) error {
	_, err := tqbcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqbcu *TkQuestionBankCityUpdate) ExecX(ctx context.Context) {
	if err := tqbcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tqbcu *TkQuestionBankCityUpdate) defaults() {
	if _, ok := tqbcu.mutation.UpdatedAt(); !ok && !tqbcu.mutation.UpdatedAtCleared() {
		v := tkquestionbankcity.UpdateDefaultUpdatedAt()
		tqbcu.mutation.SetUpdatedAt(v)
	}
}

func (tqbcu *TkQuestionBankCityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tkquestionbankcity.Table,
			Columns: tkquestionbankcity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkquestionbankcity.FieldID,
			},
		},
	}
	if ps := tqbcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tqbcu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkquestionbankcity.FieldUUID,
		})
	}
	if tqbcu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkquestionbankcity.FieldCreatedAt,
		})
	}
	if value, ok := tqbcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkquestionbankcity.FieldUpdatedAt,
		})
	}
	if tqbcu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkquestionbankcity.FieldUpdatedAt,
		})
	}
	if value, ok := tqbcu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkquestionbankcity.FieldDeletedAt,
		})
	}
	if tqbcu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkquestionbankcity.FieldDeletedAt,
		})
	}
	if tqbcu.mutation.TkQuestionBankCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbankcity.TkQuestionBankTable,
			Columns: []string{tkquestionbankcity.TkQuestionBankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestionbank.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tqbcu.mutation.TkQuestionBankIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbankcity.TkQuestionBankTable,
			Columns: []string{tkquestionbankcity.TkQuestionBankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestionbank.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tqbcu.mutation.CityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbankcity.CityTable,
			Columns: []string{tkquestionbankcity.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tqbcu.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbankcity.CityTable,
			Columns: []string{tkquestionbankcity.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tqbcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tkquestionbankcity.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TkQuestionBankCityUpdateOne is the builder for updating a single TkQuestionBankCity entity.
type TkQuestionBankCityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TkQuestionBankCityMutation
}

// SetUUID sets the "uuid" field.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetUUID(s string) *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.SetUUID(s)
	return tqbcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetUpdatedAt(t time.Time) *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.SetUpdatedAt(t)
	return tqbcuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tqbcuo *TkQuestionBankCityUpdateOne) ClearUpdatedAt() *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.ClearUpdatedAt()
	return tqbcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetDeletedAt(t time.Time) *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.SetDeletedAt(t)
	return tqbcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetNillableDeletedAt(t *time.Time) *TkQuestionBankCityUpdateOne {
	if t != nil {
		tqbcuo.SetDeletedAt(*t)
	}
	return tqbcuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tqbcuo *TkQuestionBankCityUpdateOne) ClearDeletedAt() *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.ClearDeletedAt()
	return tqbcuo
}

// SetCityID sets the "city_id" field.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetCityID(i int) *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.ResetCityID()
	tqbcuo.mutation.SetCityID(i)
	return tqbcuo
}

// SetNillableCityID sets the "city_id" field if the given value is not nil.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetNillableCityID(i *int) *TkQuestionBankCityUpdateOne {
	if i != nil {
		tqbcuo.SetCityID(*i)
	}
	return tqbcuo
}

// ClearCityID clears the value of the "city_id" field.
func (tqbcuo *TkQuestionBankCityUpdateOne) ClearCityID() *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.ClearCityID()
	return tqbcuo
}

// SetQuestionBankID sets the "question_bank_id" field.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetQuestionBankID(i int) *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.ResetQuestionBankID()
	tqbcuo.mutation.SetQuestionBankID(i)
	return tqbcuo
}

// SetNillableQuestionBankID sets the "question_bank_id" field if the given value is not nil.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetNillableQuestionBankID(i *int) *TkQuestionBankCityUpdateOne {
	if i != nil {
		tqbcuo.SetQuestionBankID(*i)
	}
	return tqbcuo
}

// ClearQuestionBankID clears the value of the "question_bank_id" field.
func (tqbcuo *TkQuestionBankCityUpdateOne) ClearQuestionBankID() *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.ClearQuestionBankID()
	return tqbcuo
}

// SetTkQuestionBankID sets the "tk_question_bank" edge to the TkQuestionBank entity by ID.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetTkQuestionBankID(id int) *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.SetTkQuestionBankID(id)
	return tqbcuo
}

// SetNillableTkQuestionBankID sets the "tk_question_bank" edge to the TkQuestionBank entity by ID if the given value is not nil.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetNillableTkQuestionBankID(id *int) *TkQuestionBankCityUpdateOne {
	if id != nil {
		tqbcuo = tqbcuo.SetTkQuestionBankID(*id)
	}
	return tqbcuo
}

// SetTkQuestionBank sets the "tk_question_bank" edge to the TkQuestionBank entity.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetTkQuestionBank(t *TkQuestionBank) *TkQuestionBankCityUpdateOne {
	return tqbcuo.SetTkQuestionBankID(t.ID)
}

// SetCity sets the "city" edge to the City entity.
func (tqbcuo *TkQuestionBankCityUpdateOne) SetCity(c *City) *TkQuestionBankCityUpdateOne {
	return tqbcuo.SetCityID(c.ID)
}

// Mutation returns the TkQuestionBankCityMutation object of the builder.
func (tqbcuo *TkQuestionBankCityUpdateOne) Mutation() *TkQuestionBankCityMutation {
	return tqbcuo.mutation
}

// ClearTkQuestionBank clears the "tk_question_bank" edge to the TkQuestionBank entity.
func (tqbcuo *TkQuestionBankCityUpdateOne) ClearTkQuestionBank() *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.ClearTkQuestionBank()
	return tqbcuo
}

// ClearCity clears the "city" edge to the City entity.
func (tqbcuo *TkQuestionBankCityUpdateOne) ClearCity() *TkQuestionBankCityUpdateOne {
	tqbcuo.mutation.ClearCity()
	return tqbcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tqbcuo *TkQuestionBankCityUpdateOne) Select(field string, fields ...string) *TkQuestionBankCityUpdateOne {
	tqbcuo.fields = append([]string{field}, fields...)
	return tqbcuo
}

// Save executes the query and returns the updated TkQuestionBankCity entity.
func (tqbcuo *TkQuestionBankCityUpdateOne) Save(ctx context.Context) (*TkQuestionBankCity, error) {
	var (
		err  error
		node *TkQuestionBankCity
	)
	tqbcuo.defaults()
	if len(tqbcuo.hooks) == 0 {
		node, err = tqbcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkQuestionBankCityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tqbcuo.mutation = mutation
			node, err = tqbcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tqbcuo.hooks) - 1; i >= 0; i-- {
			mut = tqbcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tqbcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tqbcuo *TkQuestionBankCityUpdateOne) SaveX(ctx context.Context) *TkQuestionBankCity {
	node, err := tqbcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tqbcuo *TkQuestionBankCityUpdateOne) Exec(ctx context.Context) error {
	_, err := tqbcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tqbcuo *TkQuestionBankCityUpdateOne) ExecX(ctx context.Context) {
	if err := tqbcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tqbcuo *TkQuestionBankCityUpdateOne) defaults() {
	if _, ok := tqbcuo.mutation.UpdatedAt(); !ok && !tqbcuo.mutation.UpdatedAtCleared() {
		v := tkquestionbankcity.UpdateDefaultUpdatedAt()
		tqbcuo.mutation.SetUpdatedAt(v)
	}
}

func (tqbcuo *TkQuestionBankCityUpdateOne) sqlSave(ctx context.Context) (_node *TkQuestionBankCity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tkquestionbankcity.Table,
			Columns: tkquestionbankcity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkquestionbankcity.FieldID,
			},
		},
	}
	id, ok := tqbcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TkQuestionBankCity.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tqbcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tkquestionbankcity.FieldID)
		for _, f := range fields {
			if !tkquestionbankcity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tkquestionbankcity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tqbcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tqbcuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkquestionbankcity.FieldUUID,
		})
	}
	if tqbcuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkquestionbankcity.FieldCreatedAt,
		})
	}
	if value, ok := tqbcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkquestionbankcity.FieldUpdatedAt,
		})
	}
	if tqbcuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkquestionbankcity.FieldUpdatedAt,
		})
	}
	if value, ok := tqbcuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkquestionbankcity.FieldDeletedAt,
		})
	}
	if tqbcuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tkquestionbankcity.FieldDeletedAt,
		})
	}
	if tqbcuo.mutation.TkQuestionBankCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbankcity.TkQuestionBankTable,
			Columns: []string{tkquestionbankcity.TkQuestionBankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestionbank.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tqbcuo.mutation.TkQuestionBankIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbankcity.TkQuestionBankTable,
			Columns: []string{tkquestionbankcity.TkQuestionBankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestionbank.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tqbcuo.mutation.CityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbankcity.CityTable,
			Columns: []string{tkquestionbankcity.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tqbcuo.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbankcity.CityTable,
			Columns: []string{tkquestionbankcity.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TkQuestionBankCity{config: tqbcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tqbcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tkquestionbankcity.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
