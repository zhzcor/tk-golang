// Code generated by entc, DO NOT EDIT.

package kcuserclass

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the kcuserclass type in the database.
	Label = "kc_user_class"
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
	// FieldPeriodType holds the string denoting the period_type field in the database.
	FieldPeriodType = "period_type"
	// FieldClosingDate holds the string denoting the closing_date field in the database.
	FieldClosingDate = "closing_date"
	// FieldStudyRate holds the string denoting the study_rate field in the database.
	FieldStudyRate = "study_rate"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldClassID holds the string denoting the class_id field in the database.
	FieldClassID = "class_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// Table holds the table name of the kcuserclass in the database.
	Table = "kc_user_classes"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "kc_user_classes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// ClassTable is the table the holds the class relation/edge.
	ClassTable = "kc_user_classes"
	// ClassInverseTable is the table name for the KcClass entity.
	// It exists in this package in order to avoid circular dependency with the "kcclass" package.
	ClassInverseTable = "kc_classes"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "class_id"
)

// Columns holds all SQL columns for kcuserclass fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldPeriodType,
	FieldClosingDate,
	FieldStudyRate,
	FieldRemark,
	FieldPrice,
	FieldUserID,
	FieldClassID,
}

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
//	import _ "tkserver/internal/store/ent/runtime"
//
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultPeriodType holds the default value on creation for the "period_type" field.
	DefaultPeriodType uint8
	// DefaultStudyRate holds the default value on creation for the "study_rate" field.
	DefaultStudyRate float64
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
	// DefaultPrice holds the default value on creation for the "price" field.
	DefaultPrice float64
)
