// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gserver/internal/store/ent/role"
	"gserver/internal/store/ent/rolepermission"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// RolePermission is the model entity for the RolePermission schema.
type RolePermission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RoleID holds the value of the "role_id" field.
	// 角色id
	RoleID int `json:"role_id"`
	// PermissionID holds the value of the "permission_id" field.
	// 权限ids
	PermissionID string `json:"permission_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RolePermissionQuery when eager-loading is set.
	Edges RolePermissionEdges `json:"edges"`
}

// RolePermissionEdges holds the relations/edges for other nodes in the graph.
type RolePermissionEdges struct {
	// Role holds the value of the role edge.
	Role *Role `json:"role,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RolePermissionEdges) RoleOrErr() (*Role, error) {
	if e.loadedTypes[0] {
		if e.Role == nil {
			// The edge role was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RolePermission) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case rolepermission.FieldID, rolepermission.FieldRoleID:
			values[i] = new(sql.NullInt64)
		case rolepermission.FieldPermissionID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RolePermission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RolePermission fields.
func (rp *RolePermission) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rolepermission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rp.ID = int(value.Int64)
		case rolepermission.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				rp.RoleID = int(value.Int64)
			}
		case rolepermission.FieldPermissionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field permission_id", values[i])
			} else if value.Valid {
				rp.PermissionID = value.String
			}
		}
	}
	return nil
}

// QueryRole queries the "role" edge of the RolePermission entity.
func (rp *RolePermission) QueryRole() *RoleQuery {
	return (&RolePermissionClient{config: rp.config}).QueryRole(rp)
}

// Update returns a builder for updating this RolePermission.
// Note that you need to call RolePermission.Unwrap() before calling this method if this RolePermission
// was returned from a transaction, and the transaction was committed or rolled back.
func (rp *RolePermission) Update() *RolePermissionUpdateOne {
	return (&RolePermissionClient{config: rp.config}).UpdateOne(rp)
}

// Unwrap unwraps the RolePermission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rp *RolePermission) Unwrap() *RolePermission {
	tx, ok := rp.config.driver.(*txDriver)
	if !ok {
		panic("ent: RolePermission is not a transactional entity")
	}
	rp.config.driver = tx.drv
	return rp
}

// String implements the fmt.Stringer.
func (rp *RolePermission) String() string {
	var builder strings.Builder
	builder.WriteString("RolePermission(")
	builder.WriteString(fmt.Sprintf("id=%v", rp.ID))
	builder.WriteString(", role_id=")
	builder.WriteString(fmt.Sprintf("%v", rp.RoleID))
	builder.WriteString(", permission_id=")
	builder.WriteString(rp.PermissionID)
	builder.WriteByte(')')
	return builder.String()
}

// RolePermissions is a parsable slice of RolePermission.
type RolePermissions []*RolePermission

func (rp RolePermissions) config(cfg config) {
	for _i := range rp {
		rp[_i].config = cfg
	}
}
