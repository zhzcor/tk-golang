// Code generated by entc, DO NOT EDIT.

package useraskanswer

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the useraskanswer type in the database.
	Label = "user_ask_answer"
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
	// FieldAskDesc holds the string denoting the ask_desc field in the database.
	FieldAskDesc = "ask_desc"
	// FieldAnswerStatus holds the string denoting the answer_status field in the database.
	FieldAnswerStatus = "answer_status"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTeacherID holds the string denoting the teacher_id field in the database.
	FieldTeacherID = "teacher_id"
	// FieldShowStatus holds the string denoting the show_status field in the database.
	FieldShowStatus = "show_status"
	// FieldReplyShowStatus holds the string denoting the reply_show_status field in the database.
	FieldReplyShowStatus = "reply_show_status"
	// FieldAnswerDesc holds the string denoting the answer_desc field in the database.
	FieldAnswerDesc = "answer_desc"
	// FieldAnswerAt holds the string denoting the answer_at field in the database.
	FieldAnswerAt = "answer_at"
	// EdgeTeacher holds the string denoting the teacher edge name in mutations.
	EdgeTeacher = "teacher"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeAskAnswersAttachments holds the string denoting the ask_answers_attachments edge name in mutations.
	EdgeAskAnswersAttachments = "ask_answers_attachments"
	// Table holds the table name of the useraskanswer in the database.
	Table = "user_ask_answers"
	// TeacherTable is the table the holds the teacher relation/edge.
	TeacherTable = "user_ask_answers"
	// TeacherInverseTable is the table name for the Teacher entity.
	// It exists in this package in order to avoid circular dependency with the "teacher" package.
	TeacherInverseTable = "teachers"
	// TeacherColumn is the table column denoting the teacher relation/edge.
	TeacherColumn = "teacher_id"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "user_ask_answers"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// AskAnswersAttachmentsTable is the table the holds the ask_answers_attachments relation/edge.
	AskAnswersAttachmentsTable = "user_ask_answer_attachments"
	// AskAnswersAttachmentsInverseTable is the table name for the UserAskAnswerAttachment entity.
	// It exists in this package in order to avoid circular dependency with the "useraskanswerattachment" package.
	AskAnswersAttachmentsInverseTable = "user_ask_answer_attachments"
	// AskAnswersAttachmentsColumn is the table column denoting the ask_answers_attachments relation/edge.
	AskAnswersAttachmentsColumn = "ask_id"
)

// Columns holds all SQL columns for useraskanswer fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAskDesc,
	FieldAnswerStatus,
	FieldUserID,
	FieldTeacherID,
	FieldShowStatus,
	FieldReplyShowStatus,
	FieldAnswerDesc,
	FieldAnswerAt,
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
	// DefaultAskDesc holds the default value on creation for the "ask_desc" field.
	DefaultAskDesc string
	// DefaultAnswerStatus holds the default value on creation for the "answer_status" field.
	DefaultAnswerStatus uint8
	// DefaultShowStatus holds the default value on creation for the "show_status" field.
	DefaultShowStatus uint8
	// DefaultReplyShowStatus holds the default value on creation for the "reply_show_status" field.
	DefaultReplyShowStatus uint8
	// DefaultAnswerDesc holds the default value on creation for the "answer_desc" field.
	DefaultAnswerDesc string
)
