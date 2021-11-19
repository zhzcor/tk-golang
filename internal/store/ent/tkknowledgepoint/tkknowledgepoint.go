// Code generated by entc, DO NOT EDIT.

package tkknowledgepoint

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the tkknowledgepoint type in the database.
	Label = "tk_knowledge_point"
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
	// EdgeQuestions holds the string denoting the questions edge name in mutations.
	EdgeQuestions = "questions"
	// Table holds the table name of the tkknowledgepoint in the database.
	Table = "tk_knowledge_points"
	// QuestionBankTable is the table the holds the question_bank relation/edge.
	QuestionBankTable = "tk_knowledge_points"
	// QuestionBankInverseTable is the table name for the TkQuestionBank entity.
	// It exists in this package in order to avoid circular dependency with the "tkquestionbank" package.
	QuestionBankInverseTable = "tk_question_banks"
	// QuestionBankColumn is the table column denoting the question_bank relation/edge.
	QuestionBankColumn = "question_bank_id"
	// QuestionsTable is the table the holds the questions relation/edge. The primary key declared below.
	QuestionsTable = "tk_question_knowledge_points"
	// QuestionsInverseTable is the table name for the TkQuestion entity.
	// It exists in this package in order to avoid circular dependency with the "tkquestion" package.
	QuestionsInverseTable = "tk_questions"
)

// Columns holds all SQL columns for tkknowledgepoint fields.
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

var (
	// QuestionsPrimaryKey and QuestionsColumn2 are the table columns denoting the
	// primary key for the questions relation (M2M).
	QuestionsPrimaryKey = []string{"tk_question_id", "tk_knowledge_point_id"}
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultQuestionBankID holds the default value on creation for the "question_bank_id" field.
	DefaultQuestionBankID int
	// DefaultQuestionCount holds the default value on creation for the "question_count" field.
	DefaultQuestionCount int
)
