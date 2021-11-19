// Code generated by entc, DO NOT EDIT.

package tkexampaperpartition

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the tkexampaperpartition type in the database.
	Label = "tk_exam_paper_partition"
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
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldQuestionCount holds the string denoting the question_count field in the database.
	FieldQuestionCount = "question_count"
	// FieldExamPaperID holds the string denoting the exam_paper_id field in the database.
	FieldExamPaperID = "exam_paper_id"
	// EdgeExamPaper holds the string denoting the exam_paper edge name in mutations.
	EdgeExamPaper = "exam_paper"
	// EdgeExamPartitionLinks holds the string denoting the exam_partition_links edge name in mutations.
	EdgeExamPartitionLinks = "exam_partition_links"
	// EdgeExamPartitionScores holds the string denoting the exam_partition_scores edge name in mutations.
	EdgeExamPartitionScores = "exam_partition_scores"
	// Table holds the table name of the tkexampaperpartition in the database.
	Table = "tk_exam_paper_partitions"
	// ExamPaperTable is the table the holds the exam_paper relation/edge.
	ExamPaperTable = "tk_exam_paper_partitions"
	// ExamPaperInverseTable is the table name for the TkExamPaper entity.
	// It exists in this package in order to avoid circular dependency with the "tkexampaper" package.
	ExamPaperInverseTable = "tk_exam_papers"
	// ExamPaperColumn is the table column denoting the exam_paper relation/edge.
	ExamPaperColumn = "exam_paper_id"
	// ExamPartitionLinksTable is the table the holds the exam_partition_links relation/edge.
	ExamPartitionLinksTable = "tk_exam_partition_question_links"
	// ExamPartitionLinksInverseTable is the table name for the TkExamPartitionQuestionLink entity.
	// It exists in this package in order to avoid circular dependency with the "tkexampartitionquestionlink" package.
	ExamPartitionLinksInverseTable = "tk_exam_partition_question_links"
	// ExamPartitionLinksColumn is the table column denoting the exam_partition_links relation/edge.
	ExamPartitionLinksColumn = "exam_paper_partition_id"
	// ExamPartitionScoresTable is the table the holds the exam_partition_scores relation/edge.
	ExamPartitionScoresTable = "tk_exam_paper_partition_scores"
	// ExamPartitionScoresInverseTable is the table name for the TkExamPaperPartitionScore entity.
	// It exists in this package in order to avoid circular dependency with the "tkexampaperpartitionscore" package.
	ExamPartitionScoresInverseTable = "tk_exam_paper_partition_scores"
	// ExamPartitionScoresColumn is the table column denoting the exam_partition_scores relation/edge.
	ExamPartitionScoresColumn = "exam_paper_partition_id"
)

// Columns holds all SQL columns for tkexampaperpartition fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldDesc,
	FieldDuration,
	FieldType,
	FieldQuestionCount,
	FieldExamPaperID,
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultDesc holds the default value on creation for the "desc" field.
	DefaultDesc string
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration int
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int
	// DefaultQuestionCount holds the default value on creation for the "question_count" field.
	DefaultQuestionCount uint8
)
