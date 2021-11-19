// Code generated by entc, DO NOT EDIT.

package tkexampaperpartitionscore

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the tkexampaperpartitionscore type in the database.
	Label = "tk_exam_paper_partition_score"
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
	// FieldSingeSelect holds the string denoting the singe_select field in the database.
	FieldSingeSelect = "singe_select"
	// FieldJudgeQuestion holds the string denoting the judge_question field in the database.
	FieldJudgeQuestion = "judge_question"
	// FieldMultiSelect holds the string denoting the multi_select field in the database.
	FieldMultiSelect = "multi_select"
	// FieldShorterAnswer holds the string denoting the shorter_answer field in the database.
	FieldShorterAnswer = "shorter_answer"
	// FieldMaterialQuestion holds the string denoting the material_question field in the database.
	FieldMaterialQuestion = "material_question"
	// FieldSingeSelectCount holds the string denoting the singe_select_count field in the database.
	FieldSingeSelectCount = "singe_select_count"
	// FieldJudgeQuestionCount holds the string denoting the judge_question_count field in the database.
	FieldJudgeQuestionCount = "judge_question_count"
	// FieldMultiSelectCount holds the string denoting the multi_select_count field in the database.
	FieldMultiSelectCount = "multi_select_count"
	// FieldShorterAnswerCount holds the string denoting the shorter_answer_count field in the database.
	FieldShorterAnswerCount = "shorter_answer_count"
	// FieldMaterialQuestionCount holds the string denoting the material_question_count field in the database.
	FieldMaterialQuestionCount = "material_question_count"
	// FieldExamPaperPartitionID holds the string denoting the exam_paper_partition_id field in the database.
	FieldExamPaperPartitionID = "exam_paper_partition_id"
	// EdgeExamPaperPartition holds the string denoting the exam_paper_partition edge name in mutations.
	EdgeExamPaperPartition = "exam_paper_partition"
	// Table holds the table name of the tkexampaperpartitionscore in the database.
	Table = "tk_exam_paper_partition_scores"
	// ExamPaperPartitionTable is the table the holds the exam_paper_partition relation/edge.
	ExamPaperPartitionTable = "tk_exam_paper_partition_scores"
	// ExamPaperPartitionInverseTable is the table name for the TkExamPaperPartition entity.
	// It exists in this package in order to avoid circular dependency with the "tkexampaperpartition" package.
	ExamPaperPartitionInverseTable = "tk_exam_paper_partitions"
	// ExamPaperPartitionColumn is the table column denoting the exam_paper_partition relation/edge.
	ExamPaperPartitionColumn = "exam_paper_partition_id"
)

// Columns holds all SQL columns for tkexampaperpartitionscore fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldSingeSelect,
	FieldJudgeQuestion,
	FieldMultiSelect,
	FieldShorterAnswer,
	FieldMaterialQuestion,
	FieldSingeSelectCount,
	FieldJudgeQuestionCount,
	FieldMultiSelectCount,
	FieldShorterAnswerCount,
	FieldMaterialQuestionCount,
	FieldExamPaperPartitionID,
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
	// DefaultSingeSelect holds the default value on creation for the "singe_select" field.
	DefaultSingeSelect uint8
	// DefaultJudgeQuestion holds the default value on creation for the "judge_question" field.
	DefaultJudgeQuestion uint8
	// DefaultMultiSelect holds the default value on creation for the "multi_select" field.
	DefaultMultiSelect uint8
	// DefaultShorterAnswer holds the default value on creation for the "shorter_answer" field.
	DefaultShorterAnswer uint8
	// DefaultMaterialQuestion holds the default value on creation for the "material_question" field.
	DefaultMaterialQuestion uint8
	// DefaultSingeSelectCount holds the default value on creation for the "singe_select_count" field.
	DefaultSingeSelectCount uint8
	// DefaultJudgeQuestionCount holds the default value on creation for the "judge_question_count" field.
	DefaultJudgeQuestionCount uint8
	// DefaultMultiSelectCount holds the default value on creation for the "multi_select_count" field.
	DefaultMultiSelectCount uint8
	// DefaultShorterAnswerCount holds the default value on creation for the "shorter_answer_count" field.
	DefaultShorterAnswerCount uint8
	// DefaultMaterialQuestionCount holds the default value on creation for the "material_question_count" field.
	DefaultMaterialQuestionCount uint8
)
