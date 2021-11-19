// Code generated by entc, DO NOT EDIT.

package role

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeAdmins holds the string denoting the admins edge name in mutations.
	EdgeAdmins = "admins"
	// EdgeRolePermission holds the string denoting the role_permission edge name in mutations.
	EdgeRolePermission = "role_permission"
	// Table holds the table name of the role in the database.
	Table = "roles"
	// AdminsTable is the table the holds the admins relation/edge. The primary key declared below.
	AdminsTable = "admin_roles"
	// AdminsInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	AdminsInverseTable = "admins"
	// RolePermissionTable is the table the holds the role_permission relation/edge.
	RolePermissionTable = "role_permissions"
	// RolePermissionInverseTable is the table name for the RolePermission entity.
	// It exists in this package in order to avoid circular dependency with the "rolepermission" package.
	RolePermissionInverseTable = "role_permissions"
	// RolePermissionColumn is the table column denoting the role_permission relation/edge.
	RolePermissionColumn = "role_id"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldDesc,
	FieldStatus,
}

var (
	// AdminsPrimaryKey and AdminsColumn2 are the table columns denoting the
	// primary key for the admins relation (M2M).
	AdminsPrimaryKey = []string{"admin_id", "role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "gserver/internal/store/ent/runtime"
//
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultDesc holds the default value on creation for the "desc" field.
	DefaultDesc string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
)
