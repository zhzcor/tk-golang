// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/predicate"
	"gserver/internal/store/ent/user"
	"gserver/internal/store/ent/userloginlog"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserLoginLogUpdate is the builder for updating UserLoginLog entities.
type UserLoginLogUpdate struct {
	config
	hooks    []Hook
	mutation *UserLoginLogMutation
}

// Where adds a new predicate for the UserLoginLogUpdate builder.
func (ullu *UserLoginLogUpdate) Where(ps ...predicate.UserLoginLog) *UserLoginLogUpdate {
	ullu.mutation.predicates = append(ullu.mutation.predicates, ps...)
	return ullu
}

// SetUUID sets the "uuid" field.
func (ullu *UserLoginLogUpdate) SetUUID(s string) *UserLoginLogUpdate {
	ullu.mutation.SetUUID(s)
	return ullu
}

// SetUpdatedAt sets the "updated_at" field.
func (ullu *UserLoginLogUpdate) SetUpdatedAt(t time.Time) *UserLoginLogUpdate {
	ullu.mutation.SetUpdatedAt(t)
	return ullu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ullu *UserLoginLogUpdate) ClearUpdatedAt() *UserLoginLogUpdate {
	ullu.mutation.ClearUpdatedAt()
	return ullu
}

// SetDeletedAt sets the "deleted_at" field.
func (ullu *UserLoginLogUpdate) SetDeletedAt(t time.Time) *UserLoginLogUpdate {
	ullu.mutation.SetDeletedAt(t)
	return ullu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ullu *UserLoginLogUpdate) SetNillableDeletedAt(t *time.Time) *UserLoginLogUpdate {
	if t != nil {
		ullu.SetDeletedAt(*t)
	}
	return ullu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ullu *UserLoginLogUpdate) ClearDeletedAt() *UserLoginLogUpdate {
	ullu.mutation.ClearDeletedAt()
	return ullu
}

// SetCid sets the "cid" field.
func (ullu *UserLoginLogUpdate) SetCid(s string) *UserLoginLogUpdate {
	ullu.mutation.SetCid(s)
	return ullu
}

// SetNillableCid sets the "cid" field if the given value is not nil.
func (ullu *UserLoginLogUpdate) SetNillableCid(s *string) *UserLoginLogUpdate {
	if s != nil {
		ullu.SetCid(*s)
	}
	return ullu
}

// SetPlatform sets the "platform" field.
func (ullu *UserLoginLogUpdate) SetPlatform(s string) *UserLoginLogUpdate {
	ullu.mutation.SetPlatform(s)
	return ullu
}

// SetNillablePlatform sets the "platform" field if the given value is not nil.
func (ullu *UserLoginLogUpdate) SetNillablePlatform(s *string) *UserLoginLogUpdate {
	if s != nil {
		ullu.SetPlatform(*s)
	}
	return ullu
}

// SetDevice sets the "device" field.
func (ullu *UserLoginLogUpdate) SetDevice(s string) *UserLoginLogUpdate {
	ullu.mutation.SetDevice(s)
	return ullu
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (ullu *UserLoginLogUpdate) SetNillableDevice(s *string) *UserLoginLogUpdate {
	if s != nil {
		ullu.SetDevice(*s)
	}
	return ullu
}

// SetVersion sets the "version" field.
func (ullu *UserLoginLogUpdate) SetVersion(s string) *UserLoginLogUpdate {
	ullu.mutation.SetVersion(s)
	return ullu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ullu *UserLoginLogUpdate) SetNillableVersion(s *string) *UserLoginLogUpdate {
	if s != nil {
		ullu.SetVersion(*s)
	}
	return ullu
}

// SetIP sets the "ip" field.
func (ullu *UserLoginLogUpdate) SetIP(s string) *UserLoginLogUpdate {
	ullu.mutation.SetIP(s)
	return ullu
}

// SetProvince sets the "province" field.
func (ullu *UserLoginLogUpdate) SetProvince(s string) *UserLoginLogUpdate {
	ullu.mutation.SetProvince(s)
	return ullu
}

// SetCity sets the "city" field.
func (ullu *UserLoginLogUpdate) SetCity(s string) *UserLoginLogUpdate {
	ullu.mutation.SetCity(s)
	return ullu
}

// SetLatestLoginAt sets the "latest_login_at" field.
func (ullu *UserLoginLogUpdate) SetLatestLoginAt(t time.Time) *UserLoginLogUpdate {
	ullu.mutation.SetLatestLoginAt(t)
	return ullu
}

// SetNillableLatestLoginAt sets the "latest_login_at" field if the given value is not nil.
func (ullu *UserLoginLogUpdate) SetNillableLatestLoginAt(t *time.Time) *UserLoginLogUpdate {
	if t != nil {
		ullu.SetLatestLoginAt(*t)
	}
	return ullu
}

// ClearLatestLoginAt clears the value of the "latest_login_at" field.
func (ullu *UserLoginLogUpdate) ClearLatestLoginAt() *UserLoginLogUpdate {
	ullu.mutation.ClearLatestLoginAt()
	return ullu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ullu *UserLoginLogUpdate) SetUserID(id int) *UserLoginLogUpdate {
	ullu.mutation.SetUserID(id)
	return ullu
}

// SetUser sets the "user" edge to the User entity.
func (ullu *UserLoginLogUpdate) SetUser(u *User) *UserLoginLogUpdate {
	return ullu.SetUserID(u.ID)
}

// Mutation returns the UserLoginLogMutation object of the builder.
func (ullu *UserLoginLogUpdate) Mutation() *UserLoginLogMutation {
	return ullu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ullu *UserLoginLogUpdate) ClearUser() *UserLoginLogUpdate {
	ullu.mutation.ClearUser()
	return ullu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ullu *UserLoginLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ullu.defaults()
	if len(ullu.hooks) == 0 {
		if err = ullu.check(); err != nil {
			return 0, err
		}
		affected, err = ullu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserLoginLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ullu.check(); err != nil {
				return 0, err
			}
			ullu.mutation = mutation
			affected, err = ullu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ullu.hooks) - 1; i >= 0; i-- {
			mut = ullu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ullu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ullu *UserLoginLogUpdate) SaveX(ctx context.Context) int {
	affected, err := ullu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ullu *UserLoginLogUpdate) Exec(ctx context.Context) error {
	_, err := ullu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ullu *UserLoginLogUpdate) ExecX(ctx context.Context) {
	if err := ullu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ullu *UserLoginLogUpdate) defaults() {
	if _, ok := ullu.mutation.UpdatedAt(); !ok && !ullu.mutation.UpdatedAtCleared() {
		v := userloginlog.UpdateDefaultUpdatedAt()
		ullu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ullu *UserLoginLogUpdate) check() error {
	if _, ok := ullu.mutation.UserID(); ullu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (ullu *UserLoginLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userloginlog.Table,
			Columns: userloginlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userloginlog.FieldID,
			},
		},
	}
	if ps := ullu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ullu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldUUID,
		})
	}
	if ullu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userloginlog.FieldCreatedAt,
		})
	}
	if value, ok := ullu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userloginlog.FieldUpdatedAt,
		})
	}
	if ullu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userloginlog.FieldUpdatedAt,
		})
	}
	if value, ok := ullu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userloginlog.FieldDeletedAt,
		})
	}
	if ullu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userloginlog.FieldDeletedAt,
		})
	}
	if value, ok := ullu.mutation.Cid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldCid,
		})
	}
	if value, ok := ullu.mutation.Platform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldPlatform,
		})
	}
	if value, ok := ullu.mutation.Device(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldDevice,
		})
	}
	if value, ok := ullu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldVersion,
		})
	}
	if value, ok := ullu.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldIP,
		})
	}
	if value, ok := ullu.mutation.Province(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldProvince,
		})
	}
	if value, ok := ullu.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldCity,
		})
	}
	if value, ok := ullu.mutation.LatestLoginAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userloginlog.FieldLatestLoginAt,
		})
	}
	if ullu.mutation.LatestLoginAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userloginlog.FieldLatestLoginAt,
		})
	}
	if ullu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userloginlog.UserTable,
			Columns: []string{userloginlog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ullu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userloginlog.UserTable,
			Columns: []string{userloginlog.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ullu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userloginlog.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserLoginLogUpdateOne is the builder for updating a single UserLoginLog entity.
type UserLoginLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserLoginLogMutation
}

// SetUUID sets the "uuid" field.
func (ulluo *UserLoginLogUpdateOne) SetUUID(s string) *UserLoginLogUpdateOne {
	ulluo.mutation.SetUUID(s)
	return ulluo
}

// SetUpdatedAt sets the "updated_at" field.
func (ulluo *UserLoginLogUpdateOne) SetUpdatedAt(t time.Time) *UserLoginLogUpdateOne {
	ulluo.mutation.SetUpdatedAt(t)
	return ulluo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ulluo *UserLoginLogUpdateOne) ClearUpdatedAt() *UserLoginLogUpdateOne {
	ulluo.mutation.ClearUpdatedAt()
	return ulluo
}

// SetDeletedAt sets the "deleted_at" field.
func (ulluo *UserLoginLogUpdateOne) SetDeletedAt(t time.Time) *UserLoginLogUpdateOne {
	ulluo.mutation.SetDeletedAt(t)
	return ulluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ulluo *UserLoginLogUpdateOne) SetNillableDeletedAt(t *time.Time) *UserLoginLogUpdateOne {
	if t != nil {
		ulluo.SetDeletedAt(*t)
	}
	return ulluo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ulluo *UserLoginLogUpdateOne) ClearDeletedAt() *UserLoginLogUpdateOne {
	ulluo.mutation.ClearDeletedAt()
	return ulluo
}

// SetCid sets the "cid" field.
func (ulluo *UserLoginLogUpdateOne) SetCid(s string) *UserLoginLogUpdateOne {
	ulluo.mutation.SetCid(s)
	return ulluo
}

// SetNillableCid sets the "cid" field if the given value is not nil.
func (ulluo *UserLoginLogUpdateOne) SetNillableCid(s *string) *UserLoginLogUpdateOne {
	if s != nil {
		ulluo.SetCid(*s)
	}
	return ulluo
}

// SetPlatform sets the "platform" field.
func (ulluo *UserLoginLogUpdateOne) SetPlatform(s string) *UserLoginLogUpdateOne {
	ulluo.mutation.SetPlatform(s)
	return ulluo
}

// SetNillablePlatform sets the "platform" field if the given value is not nil.
func (ulluo *UserLoginLogUpdateOne) SetNillablePlatform(s *string) *UserLoginLogUpdateOne {
	if s != nil {
		ulluo.SetPlatform(*s)
	}
	return ulluo
}

// SetDevice sets the "device" field.
func (ulluo *UserLoginLogUpdateOne) SetDevice(s string) *UserLoginLogUpdateOne {
	ulluo.mutation.SetDevice(s)
	return ulluo
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (ulluo *UserLoginLogUpdateOne) SetNillableDevice(s *string) *UserLoginLogUpdateOne {
	if s != nil {
		ulluo.SetDevice(*s)
	}
	return ulluo
}

// SetVersion sets the "version" field.
func (ulluo *UserLoginLogUpdateOne) SetVersion(s string) *UserLoginLogUpdateOne {
	ulluo.mutation.SetVersion(s)
	return ulluo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ulluo *UserLoginLogUpdateOne) SetNillableVersion(s *string) *UserLoginLogUpdateOne {
	if s != nil {
		ulluo.SetVersion(*s)
	}
	return ulluo
}

// SetIP sets the "ip" field.
func (ulluo *UserLoginLogUpdateOne) SetIP(s string) *UserLoginLogUpdateOne {
	ulluo.mutation.SetIP(s)
	return ulluo
}

// SetProvince sets the "province" field.
func (ulluo *UserLoginLogUpdateOne) SetProvince(s string) *UserLoginLogUpdateOne {
	ulluo.mutation.SetProvince(s)
	return ulluo
}

// SetCity sets the "city" field.
func (ulluo *UserLoginLogUpdateOne) SetCity(s string) *UserLoginLogUpdateOne {
	ulluo.mutation.SetCity(s)
	return ulluo
}

// SetLatestLoginAt sets the "latest_login_at" field.
func (ulluo *UserLoginLogUpdateOne) SetLatestLoginAt(t time.Time) *UserLoginLogUpdateOne {
	ulluo.mutation.SetLatestLoginAt(t)
	return ulluo
}

// SetNillableLatestLoginAt sets the "latest_login_at" field if the given value is not nil.
func (ulluo *UserLoginLogUpdateOne) SetNillableLatestLoginAt(t *time.Time) *UserLoginLogUpdateOne {
	if t != nil {
		ulluo.SetLatestLoginAt(*t)
	}
	return ulluo
}

// ClearLatestLoginAt clears the value of the "latest_login_at" field.
func (ulluo *UserLoginLogUpdateOne) ClearLatestLoginAt() *UserLoginLogUpdateOne {
	ulluo.mutation.ClearLatestLoginAt()
	return ulluo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ulluo *UserLoginLogUpdateOne) SetUserID(id int) *UserLoginLogUpdateOne {
	ulluo.mutation.SetUserID(id)
	return ulluo
}

// SetUser sets the "user" edge to the User entity.
func (ulluo *UserLoginLogUpdateOne) SetUser(u *User) *UserLoginLogUpdateOne {
	return ulluo.SetUserID(u.ID)
}

// Mutation returns the UserLoginLogMutation object of the builder.
func (ulluo *UserLoginLogUpdateOne) Mutation() *UserLoginLogMutation {
	return ulluo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ulluo *UserLoginLogUpdateOne) ClearUser() *UserLoginLogUpdateOne {
	ulluo.mutation.ClearUser()
	return ulluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ulluo *UserLoginLogUpdateOne) Select(field string, fields ...string) *UserLoginLogUpdateOne {
	ulluo.fields = append([]string{field}, fields...)
	return ulluo
}

// Save executes the query and returns the updated UserLoginLog entity.
func (ulluo *UserLoginLogUpdateOne) Save(ctx context.Context) (*UserLoginLog, error) {
	var (
		err  error
		node *UserLoginLog
	)
	ulluo.defaults()
	if len(ulluo.hooks) == 0 {
		if err = ulluo.check(); err != nil {
			return nil, err
		}
		node, err = ulluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserLoginLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ulluo.check(); err != nil {
				return nil, err
			}
			ulluo.mutation = mutation
			node, err = ulluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ulluo.hooks) - 1; i >= 0; i-- {
			mut = ulluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ulluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ulluo *UserLoginLogUpdateOne) SaveX(ctx context.Context) *UserLoginLog {
	node, err := ulluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ulluo *UserLoginLogUpdateOne) Exec(ctx context.Context) error {
	_, err := ulluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulluo *UserLoginLogUpdateOne) ExecX(ctx context.Context) {
	if err := ulluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ulluo *UserLoginLogUpdateOne) defaults() {
	if _, ok := ulluo.mutation.UpdatedAt(); !ok && !ulluo.mutation.UpdatedAtCleared() {
		v := userloginlog.UpdateDefaultUpdatedAt()
		ulluo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ulluo *UserLoginLogUpdateOne) check() error {
	if _, ok := ulluo.mutation.UserID(); ulluo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (ulluo *UserLoginLogUpdateOne) sqlSave(ctx context.Context) (_node *UserLoginLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userloginlog.Table,
			Columns: userloginlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userloginlog.FieldID,
			},
		},
	}
	id, ok := ulluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserLoginLog.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ulluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userloginlog.FieldID)
		for _, f := range fields {
			if !userloginlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userloginlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ulluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ulluo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldUUID,
		})
	}
	if ulluo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userloginlog.FieldCreatedAt,
		})
	}
	if value, ok := ulluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userloginlog.FieldUpdatedAt,
		})
	}
	if ulluo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userloginlog.FieldUpdatedAt,
		})
	}
	if value, ok := ulluo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userloginlog.FieldDeletedAt,
		})
	}
	if ulluo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userloginlog.FieldDeletedAt,
		})
	}
	if value, ok := ulluo.mutation.Cid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldCid,
		})
	}
	if value, ok := ulluo.mutation.Platform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldPlatform,
		})
	}
	if value, ok := ulluo.mutation.Device(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldDevice,
		})
	}
	if value, ok := ulluo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldVersion,
		})
	}
	if value, ok := ulluo.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldIP,
		})
	}
	if value, ok := ulluo.mutation.Province(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldProvince,
		})
	}
	if value, ok := ulluo.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userloginlog.FieldCity,
		})
	}
	if value, ok := ulluo.mutation.LatestLoginAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userloginlog.FieldLatestLoginAt,
		})
	}
	if ulluo.mutation.LatestLoginAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userloginlog.FieldLatestLoginAt,
		})
	}
	if ulluo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userloginlog.UserTable,
			Columns: []string{userloginlog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ulluo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userloginlog.UserTable,
			Columns: []string{userloginlog.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserLoginLog{config: ulluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ulluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userloginlog.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
