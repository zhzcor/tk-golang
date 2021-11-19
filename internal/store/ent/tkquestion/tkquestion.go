// Code generated by entc, DO NOT EDIT.

package tkquestion

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the tkquestion type in the database.
	Label = "tk_question"
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
	// FieldDifficulty holds the string denoting the difficulty field in the database.
	FieldDifficulty = "difficulty"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldCreatedAdminID holds the string denoting the created_admin_id field in the database.
	FieldCreatedAdminID = "created_admin_id"
	// FieldQuestionBankID holds the string denoting the question_bank_id field in the database.
	FieldQuestionBankID = "question_bank_id"
	// FieldAnswerCount holds the string denoting the answer_count field in the database.
	FieldAnswerCount = "answer_count"
	// FieldRightCount holds the string denoting the right_count field in the database.
	FieldRightCount = "right_count"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// EdgeQuestionBank holds the string denoting the question_bank edge name in mutations.
	EdgeQuestionBank = "question_bank"
	// EdgeAdmin holds the string denoting the admin edge name in mutations.
	EdgeAdmin = "admin"
	// EdgeUserRandDom holds the string denoting the user_rand_dom edge name in mutations.
	EdgeUserRandDom = "user_rand_dom"
	// EdgeAnswerOptions holds the string denoting the answer_options edge name in mutations.
	EdgeAnswerOptions = "answer_options"
	// EdgeKnowledgePoints holds the string denoting the knowledge_points edge name in mutations.
	EdgeKnowledgePoints = "knowledge_points"
	// EdgeQuestionSectionLinks holds the string denoting the question_section_links edge name in mutations.
	EdgeQuestionSectionLinks = "question_section_links"
	// EdgeExamPartitionQuestions holds the string denoting the exam_partition_questions edge name in mutations.
	EdgeExamPartitionQuestions = "exam_partition_questions"
	// EdgeQuestionErrorFeedbacks holds the string denoting the question_error_feedbacks edge name in mutations.
	EdgeQuestionErrorFeedbacks = "question_error_feedbacks"
	// EdgeUserRecords holds the string denoting the user_records edge name in mutations.
	EdgeUserRecords = "user_records"
	// EdgeSmallCategoryQuestions holds the string denoting the small_category_questions edge name in mutations.
	EdgeSmallCategoryQuestions = "small_category_questions"
	// EdgeUserExamQuestions holds the string denoting the user_exam_questions edge name in mutations.
	EdgeUserExamQuestions = "user_exam_questions"
	// EdgeUserRecodeWrong holds the string denoting the user_recode_wrong edge name in mutations.
	EdgeUserRecodeWrong = "user_recode_wrong"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeCollectionQuestion holds the string denoting the collection_question edge name in mutations.
	EdgeCollectionQuestion = "collection_question"
	// Table holds the table name of the tkquestion in the database.
	Table = "tk_questions"
	// QuestionBankTable is the table the holds the question_bank relation/edge.
	QuestionBankTable = "tk_questions"
	// QuestionBankInverseTable is the table name for the TkQuestionBank entity.
	// It exists in this package in order to avoid circular dependency with the "tkquestionbank" package.
	QuestionBankInverseTable = "tk_question_banks"
	// QuestionBankColumn is the table column denoting the question_bank relation/edge.
	QuestionBankColumn = "question_bank_id"
	// AdminTable is the table the holds the admin relation/edge.
	AdminTable = "tk_questions"
	// AdminInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	AdminInverseTable = "admins"
	// AdminColumn is the table column denoting the admin relation/edge.
	AdminColumn = "created_admin_id"
	// UserRandDomTable is the table the holds the user_rand_dom relation/edge. The primary key declared below.
	UserRandDomTable = "tk_user_random_exam_recode_random_exam_question"
	// UserRandDomInverseTable is the table name for the TkUserRandomExamRecode entity.
	// It exists in this package in order to avoid circular dependency with the "tkuserrandomexamrecode" package.
	UserRandDomInverseTable = "tk_user_random_exam_recodes"
	// AnswerOptionsTable is the table the holds the answer_options relation/edge.
	AnswerOptionsTable = "tk_question_answer_options"
	// AnswerOptionsInverseTable is the table name for the TkQuestionAnswerOption entity.
	// It exists in this package in order to avoid circular dependency with the "tkquestionansweroption" package.
	AnswerOptionsInverseTable = "tk_question_answer_options"
	// AnswerOptionsColumn is the table column denoting the answer_options relation/edge.
	AnswerOptionsColumn = "question_id"
	// KnowledgePointsTable is the table the holds the knowledge_points relation/edge. The primary key declared below.
	KnowledgePointsTable = "tk_question_knowledge_points"
	// KnowledgePointsInverseTable is the table name for the TkKnowledgePoint entity.
	// It exists in this package in order to avoid circular dependency with the "tkknowledgepoint" package.
	KnowledgePointsInverseTable = "tk_knowledge_points"
	// QuestionSectionLinksTable is the table the holds the question_section_links relation/edge.
	QuestionSectionLinksTable = "tk_question_sections"
	// QuestionSectionLinksInverseTable is the table name for the TkQuestionSection entity.
	// It exists in this package in order to avoid circular dependency with the "tkquestionsection" package.
	QuestionSectionLinksInverseTable = "tk_question_sections"
	// QuestionSectionLinksColumn is the table column denoting the question_section_links relation/edge.
	QuestionSectionLinksColumn = "question_id"
	// ExamPartitionQuestionsTable is the table the holds the exam_partition_questions relation/edge.
	ExamPartitionQuestionsTable = "tk_exam_partition_question_links"
	// ExamPartitionQuestionsInverseTable is the table name for the TkExamPartitionQuestionLink entity.
	// It exists in this package in order to avoid circular dependency with the "tkexampartitionquestionlink" package.
	ExamPartitionQuestionsInverseTable = "tk_exam_partition_question_links"
	// ExamPartitionQuestionsColumn is the table column denoting the exam_partition_questions relation/edge.
	ExamPartitionQuestionsColumn = "question_id"
	// QuestionErrorFeedbacksTable is the table the holds the question_error_feedbacks relation/edge.
	QuestionErrorFeedbacksTable = "tk_question_error_feedbacks"
	// QuestionErrorFeedbacksInverseTable is the table name for the TkQuestionErrorFeedback entity.
	// It exists in this package in order to avoid circular dependency with the "tkquestionerrorfeedback" package.
	QuestionErrorFeedbacksInverseTable = "tk_question_error_feedbacks"
	// QuestionErrorFeedbacksColumn is the table column denoting the question_error_feedbacks relation/edge.
	QuestionErrorFeedbacksColumn = "question_id"
	// UserRecordsTable is the table the holds the user_records relation/edge.
	UserRecordsTable = "tk_user_question_records"
	// UserRecordsInverseTable is the table name for the TkUserQuestionRecord entity.
	// It exists in this package in order to avoid circular dependency with the "tkuserquestionrecord" package.
	UserRecordsInverseTable = "tk_user_question_records"
	// UserRecordsColumn is the table column denoting the user_records relation/edge.
	UserRecordsColumn = "question_id"
	// SmallCategoryQuestionsTable is the table the holds the small_category_questions relation/edge.
	SmallCategoryQuestionsTable = "kc_small_category_questions"
	// SmallCategoryQuestionsInverseTable is the table name for the KcSmallCategoryQuestion entity.
	// It exists in this package in order to avoid circular dependency with the "kcsmallcategoryquestion" package.
	SmallCategoryQuestionsInverseTable = "kc_small_category_questions"
	// SmallCategoryQuestionsColumn is the table column denoting the small_category_questions relation/edge.
	SmallCategoryQuestionsColumn = "question_id"
	// UserExamQuestionsTable is the table the holds the user_exam_questions relation/edge.
	UserExamQuestionsTable = "tk_user_simulation_teacher_marks"
	// UserExamQuestionsInverseTable is the table name for the TkUserSimulationTeacherMark entity.
	// It exists in this package in order to avoid circular dependency with the "tkusersimulationteachermark" package.
	UserExamQuestionsInverseTable = "tk_user_simulation_teacher_marks"
	// UserExamQuestionsColumn is the table column denoting the user_exam_questions relation/edge.
	UserExamQuestionsColumn = "question_id"
	// UserRecodeWrongTable is the table the holds the user_recode_wrong relation/edge.
	UserRecodeWrongTable = "tk_user_wrong_question_recodes"
	// UserRecodeWrongInverseTable is the table name for the TkUserWrongQuestionRecode entity.
	// It exists in this package in order to avoid circular dependency with the "tkuserwrongquestionrecode" package.
	UserRecodeWrongInverseTable = "tk_user_wrong_question_recodes"
	// UserRecodeWrongColumn is the table column denoting the user_recode_wrong relation/edge.
	UserRecodeWrongColumn = "question_id"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "tk_questions"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "pid"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "tk_questions"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "pid"
	// CollectionQuestionTable is the table the holds the collection_question relation/edge.
	CollectionQuestionTable = "collections"
	// CollectionQuestionInverseTable is the table name for the Collection entity.
	// It exists in this package in order to avoid circular dependency with the "collection" package.
	CollectionQuestionInverseTable = "collections"
	// CollectionQuestionColumn is the table column denoting the collection_question relation/edge.
	CollectionQuestionColumn = "value_id"
)

// Columns holds all SQL columns for tkquestion fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldDifficulty,
	FieldType,
	FieldDesc,
	FieldCreatedAdminID,
	FieldQuestionBankID,
	FieldAnswerCount,
	FieldRightCount,
	FieldPid,
}

var (
	// UserRandDomPrimaryKey and UserRandDomColumn2 are the table columns denoting the
	// primary key for the user_rand_dom relation (M2M).
	UserRandDomPrimaryKey = []string{"tk_user_random_exam_recode_id", "tk_question_id"}
	// KnowledgePointsPrimaryKey and KnowledgePointsColumn2 are the table columns denoting the
	// primary key for the knowledge_points relation (M2M).
	KnowledgePointsPrimaryKey = []string{"tk_question_id", "tk_knowledge_point_id"}
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
	// DefaultDifficulty holds the default value on creation for the "difficulty" field.
	DefaultDifficulty uint8
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType uint8
	// DefaultAnswerCount holds the default value on creation for the "answer_count" field.
	DefaultAnswerCount int
	// DefaultRightCount holds the default value on creation for the "right_count" field.
	DefaultRightCount int
)
