// Code generated by entc, DO NOT EDIT.

package tkchapter

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the tkchapter type in the database.
	Label = "tk_chapter"
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
	// FieldQuestionBankID holds the string denoting the question_bank_id field in the database.
	FieldQuestionBankID = "question_bank_id"
	// FieldQuestionCount holds the string denoting the question_count field in the database.
	FieldQuestionCount = "question_count"
	// EdgeQuestionBank holds the string denoting the question_bank edge name in mutations.
	EdgeQuestionBank = "question_bank"
	// EdgeSections holds the string denoting the sections edge name in mutations.
	EdgeSections = "sections"
	// Table holds the table name of the tkchapter in the database.
	Table = "tk_chapters"
	// QuestionBankTable is the table the holds the question_bank relation/edge.
	QuestionBankTable = "tk_chapters"
	// QuestionBankInverseTable is the table name for the TkQuestionBank entity.
	// It exists in this package in order to avoid circular dependency with the "tkquestionbank" package.
	QuestionBankInverseTable = "tk_question_banks"
	// QuestionBankColumn is the table column denoting the question_bank relation/edge.
	QuestionBankColumn = "question_bank_id"
	// SectionsTable is the table the holds the sections relation/edge.
	SectionsTable = "tk_sections"
	// SectionsInverseTable is the table name for the TkSection entity.
	// It exists in this package in order to avoid circular dependency with the "tksection" package.
	SectionsInverseTable = "tk_sections"
	// SectionsColumn is the table column denoting the sections relation/edge.
	SectionsColumn = "chapter_id"
)

// Columns holds all SQL columns for tkchapter fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldQuestionBankID,
	FieldQuestionCount,
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
	// DefaultQuestionCount holds the default value on creation for the "question_count" field.
	DefaultQuestionCount int
)
