// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"
	"tkserver/internal/store/ent/appversion"
	"tkserver/internal/store/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AppVersionUpdate is the builder for updating AppVersion entities.
type AppVersionUpdate struct {
	config
	hooks    []Hook
	mutation *AppVersionMutation
}

// Where adds a new predicate for the AppVersionUpdate builder.
func (avu *AppVersionUpdate) Where(ps ...predicate.AppVersion) *AppVersionUpdate {
	avu.mutation.predicates = append(avu.mutation.predicates, ps...)
	return avu
}

// SetUUID sets the "uuid" field.
func (avu *AppVersionUpdate) SetUUID(s string) *AppVersionUpdate {
	avu.mutation.SetUUID(s)
	return avu
}

// SetUpdatedAt sets the "updated_at" field.
func (avu *AppVersionUpdate) SetUpdatedAt(t time.Time) *AppVersionUpdate {
	avu.mutation.SetUpdatedAt(t)
	return avu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (avu *AppVersionUpdate) ClearUpdatedAt() *AppVersionUpdate {
	avu.mutation.ClearUpdatedAt()
	return avu
}

// SetDeletedAt sets the "deleted_at" field.
func (avu *AppVersionUpdate) SetDeletedAt(t time.Time) *AppVersionUpdate {
	avu.mutation.SetDeletedAt(t)
	return avu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (avu *AppVersionUpdate) SetNillableDeletedAt(t *time.Time) *AppVersionUpdate {
	if t != nil {
		avu.SetDeletedAt(*t)
	}
	return avu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (avu *AppVersionUpdate) ClearDeletedAt() *AppVersionUpdate {
	avu.mutation.ClearDeletedAt()
	return avu
}

// SetIP sets the "ip" field.
func (avu *AppVersionUpdate) SetIP(u uint16) *AppVersionUpdate {
	avu.mutation.ResetIP()
	avu.mutation.SetIP(u)
	return avu
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (avu *AppVersionUpdate) SetNillableIP(u *uint16) *AppVersionUpdate {
	if u != nil {
		avu.SetIP(*u)
	}
	return avu
}

// AddIP adds u to the "ip" field.
func (avu *AppVersionUpdate) AddIP(u uint16) *AppVersionUpdate {
	avu.mutation.AddIP(u)
	return avu
}

// SetName sets the "name" field.
func (avu *AppVersionUpdate) SetName(s string) *AppVersionUpdate {
	avu.mutation.SetName(s)
	return avu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (avu *AppVersionUpdate) SetNillableName(s *string) *AppVersionUpdate {
	if s != nil {
		avu.SetName(*s)
	}
	return avu
}

// SetSn sets the "sn" field.
func (avu *AppVersionUpdate) SetSn(s string) *AppVersionUpdate {
	avu.mutation.SetSn(s)
	return avu
}

// SetNillableSn sets the "sn" field if the given value is not nil.
func (avu *AppVersionUpdate) SetNillableSn(s *string) *AppVersionUpdate {
	if s != nil {
		avu.SetSn(*s)
	}
	return avu
}

// SetRemark sets the "remark" field.
func (avu *AppVersionUpdate) SetRemark(s string) *AppVersionUpdate {
	avu.mutation.SetRemark(s)
	return avu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (avu *AppVersionUpdate) SetNillableRemark(s *string) *AppVersionUpdate {
	if s != nil {
		avu.SetRemark(*s)
	}
	return avu
}

// ClearRemark clears the value of the "remark" field.
func (avu *AppVersionUpdate) ClearRemark() *AppVersionUpdate {
	avu.mutation.ClearRemark()
	return avu
}

// SetURL sets the "url" field.
func (avu *AppVersionUpdate) SetURL(s string) *AppVersionUpdate {
	avu.mutation.SetURL(s)
	return avu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (avu *AppVersionUpdate) SetNillableURL(s *string) *AppVersionUpdate {
	if s != nil {
		avu.SetURL(*s)
	}
	return avu
}

// SetPhoneType sets the "phone_type" field.
func (avu *AppVersionUpdate) SetPhoneType(u uint8) *AppVersionUpdate {
	avu.mutation.ResetPhoneType()
	avu.mutation.SetPhoneType(u)
	return avu
}

// SetNillablePhoneType sets the "phone_type" field if the given value is not nil.
func (avu *AppVersionUpdate) SetNillablePhoneType(u *uint8) *AppVersionUpdate {
	if u != nil {
		avu.SetPhoneType(*u)
	}
	return avu
}

// AddPhoneType adds u to the "phone_type" field.
func (avu *AppVersionUpdate) AddPhoneType(u uint8) *AppVersionUpdate {
	avu.mutation.AddPhoneType(u)
	return avu
}

// SetIsForceUpdate sets the "is_force_update" field.
func (avu *AppVersionUpdate) SetIsForceUpdate(u uint8) *AppVersionUpdate {
	avu.mutation.ResetIsForceUpdate()
	avu.mutation.SetIsForceUpdate(u)
	return avu
}

// SetNillableIsForceUpdate sets the "is_force_update" field if the given value is not nil.
func (avu *AppVersionUpdate) SetNillableIsForceUpdate(u *uint8) *AppVersionUpdate {
	if u != nil {
		avu.SetIsForceUpdate(*u)
	}
	return avu
}

// AddIsForceUpdate adds u to the "is_force_update" field.
func (avu *AppVersionUpdate) AddIsForceUpdate(u uint8) *AppVersionUpdate {
	avu.mutation.AddIsForceUpdate(u)
	return avu
}

// Mutation returns the AppVersionMutation object of the builder.
func (avu *AppVersionUpdate) Mutation() *AppVersionMutation {
	return avu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (avu *AppVersionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	avu.defaults()
	if len(avu.hooks) == 0 {
		affected, err = avu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppVersionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			avu.mutation = mutation
			affected, err = avu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(avu.hooks) - 1; i >= 0; i-- {
			mut = avu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, avu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (avu *AppVersionUpdate) SaveX(ctx context.Context) int {
	affected, err := avu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (avu *AppVersionUpdate) Exec(ctx context.Context) error {
	_, err := avu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (avu *AppVersionUpdate) ExecX(ctx context.Context) {
	if err := avu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (avu *AppVersionUpdate) defaults() {
	if _, ok := avu.mutation.UpdatedAt(); !ok && !avu.mutation.UpdatedAtCleared() {
		v := appversion.UpdateDefaultUpdatedAt()
		avu.mutation.SetUpdatedAt(v)
	}
}

func (avu *AppVersionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appversion.Table,
			Columns: appversion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: appversion.FieldID,
			},
		},
	}
	if ps := avu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := avu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appversion.FieldUUID,
		})
	}
	if avu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: appversion.FieldCreatedAt,
		})
	}
	if value, ok := avu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appversion.FieldUpdatedAt,
		})
	}
	if avu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: appversion.FieldUpdatedAt,
		})
	}
	if value, ok := avu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appversion.FieldDeletedAt,
		})
	}
	if avu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: appversion.FieldDeletedAt,
		})
	}
	if value, ok := avu.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: appversion.FieldIP,
		})
	}
	if value, ok := avu.mutation.AddedIP(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: appversion.FieldIP,
		})
	}
	if value, ok := avu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appversion.FieldName,
		})
	}
	if value, ok := avu.mutation.Sn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appversion.FieldSn,
		})
	}
	if value, ok := avu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appversion.FieldRemark,
		})
	}
	if avu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appversion.FieldRemark,
		})
	}
	if value, ok := avu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appversion.FieldURL,
		})
	}
	if value, ok := avu.mutation.PhoneType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appversion.FieldPhoneType,
		})
	}
	if value, ok := avu.mutation.AddedPhoneType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appversion.FieldPhoneType,
		})
	}
	if value, ok := avu.mutation.IsForceUpdate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appversion.FieldIsForceUpdate,
		})
	}
	if value, ok := avu.mutation.AddedIsForceUpdate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appversion.FieldIsForceUpdate,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, avu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appversion.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AppVersionUpdateOne is the builder for updating a single AppVersion entity.
type AppVersionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppVersionMutation
}

// SetUUID sets the "uuid" field.
func (avuo *AppVersionUpdateOne) SetUUID(s string) *AppVersionUpdateOne {
	avuo.mutation.SetUUID(s)
	return avuo
}

// SetUpdatedAt sets the "updated_at" field.
func (avuo *AppVersionUpdateOne) SetUpdatedAt(t time.Time) *AppVersionUpdateOne {
	avuo.mutation.SetUpdatedAt(t)
	return avuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (avuo *AppVersionUpdateOne) ClearUpdatedAt() *AppVersionUpdateOne {
	avuo.mutation.ClearUpdatedAt()
	return avuo
}

// SetDeletedAt sets the "deleted_at" field.
func (avuo *AppVersionUpdateOne) SetDeletedAt(t time.Time) *AppVersionUpdateOne {
	avuo.mutation.SetDeletedAt(t)
	return avuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (avuo *AppVersionUpdateOne) SetNillableDeletedAt(t *time.Time) *AppVersionUpdateOne {
	if t != nil {
		avuo.SetDeletedAt(*t)
	}
	return avuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (avuo *AppVersionUpdateOne) ClearDeletedAt() *AppVersionUpdateOne {
	avuo.mutation.ClearDeletedAt()
	return avuo
}

// SetIP sets the "ip" field.
func (avuo *AppVersionUpdateOne) SetIP(u uint16) *AppVersionUpdateOne {
	avuo.mutation.ResetIP()
	avuo.mutation.SetIP(u)
	return avuo
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (avuo *AppVersionUpdateOne) SetNillableIP(u *uint16) *AppVersionUpdateOne {
	if u != nil {
		avuo.SetIP(*u)
	}
	return avuo
}

// AddIP adds u to the "ip" field.
func (avuo *AppVersionUpdateOne) AddIP(u uint16) *AppVersionUpdateOne {
	avuo.mutation.AddIP(u)
	return avuo
}

// SetName sets the "name" field.
func (avuo *AppVersionUpdateOne) SetName(s string) *AppVersionUpdateOne {
	avuo.mutation.SetName(s)
	return avuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (avuo *AppVersionUpdateOne) SetNillableName(s *string) *AppVersionUpdateOne {
	if s != nil {
		avuo.SetName(*s)
	}
	return avuo
}

// SetSn sets the "sn" field.
func (avuo *AppVersionUpdateOne) SetSn(s string) *AppVersionUpdateOne {
	avuo.mutation.SetSn(s)
	return avuo
}

// SetNillableSn sets the "sn" field if the given value is not nil.
func (avuo *AppVersionUpdateOne) SetNillableSn(s *string) *AppVersionUpdateOne {
	if s != nil {
		avuo.SetSn(*s)
	}
	return avuo
}

// SetRemark sets the "remark" field.
func (avuo *AppVersionUpdateOne) SetRemark(s string) *AppVersionUpdateOne {
	avuo.mutation.SetRemark(s)
	return avuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (avuo *AppVersionUpdateOne) SetNillableRemark(s *string) *AppVersionUpdateOne {
	if s != nil {
		avuo.SetRemark(*s)
	}
	return avuo
}

// ClearRemark clears the value of the "remark" field.
func (avuo *AppVersionUpdateOne) ClearRemark() *AppVersionUpdateOne {
	avuo.mutation.ClearRemark()
	return avuo
}

// SetURL sets the "url" field.
func (avuo *AppVersionUpdateOne) SetURL(s string) *AppVersionUpdateOne {
	avuo.mutation.SetURL(s)
	return avuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (avuo *AppVersionUpdateOne) SetNillableURL(s *string) *AppVersionUpdateOne {
	if s != nil {
		avuo.SetURL(*s)
	}
	return avuo
}

// SetPhoneType sets the "phone_type" field.
func (avuo *AppVersionUpdateOne) SetPhoneType(u uint8) *AppVersionUpdateOne {
	avuo.mutation.ResetPhoneType()
	avuo.mutation.SetPhoneType(u)
	return avuo
}

// SetNillablePhoneType sets the "phone_type" field if the given value is not nil.
func (avuo *AppVersionUpdateOne) SetNillablePhoneType(u *uint8) *AppVersionUpdateOne {
	if u != nil {
		avuo.SetPhoneType(*u)
	}
	return avuo
}

// AddPhoneType adds u to the "phone_type" field.
func (avuo *AppVersionUpdateOne) AddPhoneType(u uint8) *AppVersionUpdateOne {
	avuo.mutation.AddPhoneType(u)
	return avuo
}

// SetIsForceUpdate sets the "is_force_update" field.
func (avuo *AppVersionUpdateOne) SetIsForceUpdate(u uint8) *AppVersionUpdateOne {
	avuo.mutation.ResetIsForceUpdate()
	avuo.mutation.SetIsForceUpdate(u)
	return avuo
}

// SetNillableIsForceUpdate sets the "is_force_update" field if the given value is not nil.
func (avuo *AppVersionUpdateOne) SetNillableIsForceUpdate(u *uint8) *AppVersionUpdateOne {
	if u != nil {
		avuo.SetIsForceUpdate(*u)
	}
	return avuo
}

// AddIsForceUpdate adds u to the "is_force_update" field.
func (avuo *AppVersionUpdateOne) AddIsForceUpdate(u uint8) *AppVersionUpdateOne {
	avuo.mutation.AddIsForceUpdate(u)
	return avuo
}

// Mutation returns the AppVersionMutation object of the builder.
func (avuo *AppVersionUpdateOne) Mutation() *AppVersionMutation {
	return avuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (avuo *AppVersionUpdateOne) Select(field string, fields ...string) *AppVersionUpdateOne {
	avuo.fields = append([]string{field}, fields...)
	return avuo
}

// Save executes the query and returns the updated AppVersion entity.
func (avuo *AppVersionUpdateOne) Save(ctx context.Context) (*AppVersion, error) {
	var (
		err  error
		node *AppVersion
	)
	avuo.defaults()
	if len(avuo.hooks) == 0 {
		node, err = avuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppVersionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			avuo.mutation = mutation
			node, err = avuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(avuo.hooks) - 1; i >= 0; i-- {
			mut = avuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, avuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (avuo *AppVersionUpdateOne) SaveX(ctx context.Context) *AppVersion {
	node, err := avuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (avuo *AppVersionUpdateOne) Exec(ctx context.Context) error {
	_, err := avuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (avuo *AppVersionUpdateOne) ExecX(ctx context.Context) {
	if err := avuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (avuo *AppVersionUpdateOne) defaults() {
	if _, ok := avuo.mutation.UpdatedAt(); !ok && !avuo.mutation.UpdatedAtCleared() {
		v := appversion.UpdateDefaultUpdatedAt()
		avuo.mutation.SetUpdatedAt(v)
	}
}

func (avuo *AppVersionUpdateOne) sqlSave(ctx context.Context) (_node *AppVersion, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appversion.Table,
			Columns: appversion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: appversion.FieldID,
			},
		},
	}
	id, ok := avuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AppVersion.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := avuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appversion.FieldID)
		for _, f := range fields {
			if !appversion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := avuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := avuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appversion.FieldUUID,
		})
	}
	if avuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: appversion.FieldCreatedAt,
		})
	}
	if value, ok := avuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appversion.FieldUpdatedAt,
		})
	}
	if avuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: appversion.FieldUpdatedAt,
		})
	}
	if value, ok := avuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: appversion.FieldDeletedAt,
		})
	}
	if avuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: appversion.FieldDeletedAt,
		})
	}
	if value, ok := avuo.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: appversion.FieldIP,
		})
	}
	if value, ok := avuo.mutation.AddedIP(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: appversion.FieldIP,
		})
	}
	if value, ok := avuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appversion.FieldName,
		})
	}
	if value, ok := avuo.mutation.Sn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appversion.FieldSn,
		})
	}
	if value, ok := avuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appversion.FieldRemark,
		})
	}
	if avuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appversion.FieldRemark,
		})
	}
	if value, ok := avuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appversion.FieldURL,
		})
	}
	if value, ok := avuo.mutation.PhoneType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appversion.FieldPhoneType,
		})
	}
	if value, ok := avuo.mutation.AddedPhoneType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appversion.FieldPhoneType,
		})
	}
	if value, ok := avuo.mutation.IsForceUpdate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appversion.FieldIsForceUpdate,
		})
	}
	if value, ok := avuo.mutation.AddedIsForceUpdate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appversion.FieldIsForceUpdate,
		})
	}
	_node = &AppVersion{config: avuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, avuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appversion.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
