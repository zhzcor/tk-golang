// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gserver/internal/store/ent/role"
	"gserver/internal/store/ent/rolepermission"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at"`
	// Name holds the value of the "name" field.
	// 角色名称
	Name string `json:"name"`
	// Desc holds the value of the "desc" field.
	// 描述
	Desc string `json:"desc"`
	// Status holds the value of the "status" field.
	// 状态 1：开启 2：关闭
	Status uint8 `json:"status"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges RoleEdges `json:"edges"`
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// Admins holds the value of the admins edge.
	Admins []*Admin `json:"admins,omitempty"`
	// RolePermission holds the value of the role_permission edge.
	RolePermission *RolePermission `json:"role_permission,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AdminsOrErr returns the Admins value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) AdminsOrErr() ([]*Admin, error) {
	if e.loadedTypes[0] {
		return e.Admins, nil
	}
	return nil, &NotLoadedError{edge: "admins"}
}

// RolePermissionOrErr returns the RolePermission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleEdges) RolePermissionOrErr() (*RolePermission, error) {
	if e.loadedTypes[1] {
		if e.RolePermission == nil {
			// The edge role_permission was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: rolepermission.Label}
		}
		return e.RolePermission, nil
	}
	return nil, &NotLoadedError{edge: "role_permission"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldID, role.FieldStatus:
			values[i] = new(sql.NullInt64)
		case role.FieldUUID, role.FieldName, role.FieldDesc:
			values[i] = new(sql.NullString)
		case role.FieldCreatedAt, role.FieldUpdatedAt, role.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Role", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case role.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				r.UUID = value.String
			}
		case role.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = new(time.Time)
				*r.CreatedAt = value.Time
			}
		case role.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = new(time.Time)
				*r.UpdatedAt = value.Time
			}
		case role.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = new(time.Time)
				*r.DeletedAt = value.Time
			}
		case role.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case role.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				r.Desc = value.String
			}
		case role.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = uint8(value.Int64)
			}
		}
	}
	return nil
}

// QueryAdmins queries the "admins" edge of the Role entity.
func (r *Role) QueryAdmins() *AdminQuery {
	return (&RoleClient{config: r.config}).QueryAdmins(r)
}

// QueryRolePermission queries the "role_permission" edge of the Role entity.
func (r *Role) QueryRolePermission() *RolePermissionQuery {
	return (&RoleClient{config: r.config}).QueryRolePermission(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return (&RoleClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(r.UUID)
	if v := r.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := r.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := r.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", desc=")
	builder.WriteString(r.Desc)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role

func (r Roles) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
