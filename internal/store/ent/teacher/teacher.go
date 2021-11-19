// Code generated by entc, DO NOT EDIT.

package teacher

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the teacher type in the database.
	Label = "teacher"
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
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldSubTitle holds the string denoting the sub_title field in the database.
	FieldSubTitle = "sub_title"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTeachingAge holds the string denoting the teaching_age field in the database.
	FieldTeachingAge = "teaching_age"
	// FieldAvatarID holds the string denoting the avatar_id field in the database.
	FieldAvatarID = "avatar_id"
	// FieldSortOrder holds the string denoting the sort_order field in the database.
	FieldSortOrder = "sort_order"
	// EdgeMajors holds the string denoting the majors edge name in mutations.
	EdgeMajors = "majors"
	// EdgeTeacherTags holds the string denoting the teacher_tags edge name in mutations.
	EdgeTeacherTags = "teacher_tags"
	// EdgeTeacherClasses holds the string denoting the teacher_classes edge name in mutations.
	EdgeTeacherClasses = "teacher_classes"
	// EdgeKcClassMasters holds the string denoting the kc_class_masters edge name in mutations.
	EdgeKcClassMasters = "kc_class_masters"
	// EdgeTeacherCourses holds the string denoting the teacher_courses edge name in mutations.
	EdgeTeacherCourses = "teacher_courses"
	// EdgeUserExamsTeachers holds the string denoting the user_exams_teachers edge name in mutations.
	EdgeUserExamsTeachers = "user_exams_teachers"
	// EdgeAskTeachers holds the string denoting the ask_teachers edge name in mutations.
	EdgeAskTeachers = "ask_teachers"
	// EdgeAttachment holds the string denoting the attachment edge name in mutations.
	EdgeAttachment = "attachment"
	// Table holds the table name of the teacher in the database.
	Table = "teachers"
	// MajorsTable is the table the holds the majors relation/edge. The primary key declared below.
	MajorsTable = "teacher_majors"
	// MajorsInverseTable is the table name for the Major entity.
	// It exists in this package in order to avoid circular dependency with the "major" package.
	MajorsInverseTable = "majors"
	// TeacherTagsTable is the table the holds the teacher_tags relation/edge.
	TeacherTagsTable = "teacher_tags"
	// TeacherTagsInverseTable is the table name for the TeacherTag entity.
	// It exists in this package in order to avoid circular dependency with the "teachertag" package.
	TeacherTagsInverseTable = "teacher_tags"
	// TeacherTagsColumn is the table column denoting the teacher_tags relation/edge.
	TeacherTagsColumn = "teacher_id"
	// TeacherClassesTable is the table the holds the teacher_classes relation/edge.
	TeacherClassesTable = "kc_class_teachers"
	// TeacherClassesInverseTable is the table name for the KcClassTeacher entity.
	// It exists in this package in order to avoid circular dependency with the "kcclassteacher" package.
	TeacherClassesInverseTable = "kc_class_teachers"
	// TeacherClassesColumn is the table column denoting the teacher_classes relation/edge.
	TeacherClassesColumn = "teacher_id"
	// KcClassMastersTable is the table the holds the kc_class_masters relation/edge.
	KcClassMastersTable = "kc_classes"
	// KcClassMastersInverseTable is the table name for the KcClass entity.
	// It exists in this package in order to avoid circular dependency with the "kcclass" package.
	KcClassMastersInverseTable = "kc_classes"
	// KcClassMastersColumn is the table column denoting the kc_class_masters relation/edge.
	KcClassMastersColumn = "class_head_master_id"
	// TeacherCoursesTable is the table the holds the teacher_courses relation/edge.
	TeacherCoursesTable = "kc_course_teachers"
	// TeacherCoursesInverseTable is the table name for the KcCourseTeacher entity.
	// It exists in this package in order to avoid circular dependency with the "kccourseteacher" package.
	TeacherCoursesInverseTable = "kc_course_teachers"
	// TeacherCoursesColumn is the table column denoting the teacher_courses relation/edge.
	TeacherCoursesColumn = "teacher_id"
	// UserExamsTeachersTable is the table the holds the user_exams_teachers relation/edge.
	UserExamsTeachersTable = "tk_user_exam_score_records"
	// UserExamsTeachersInverseTable is the table name for the TkUserExamScoreRecord entity.
	// It exists in this package in order to avoid circular dependency with the "tkuserexamscorerecord" package.
	UserExamsTeachersInverseTable = "tk_user_exam_score_records"
	// UserExamsTeachersColumn is the table column denoting the user_exams_teachers relation/edge.
	UserExamsTeachersColumn = "operation_teacher_id"
	// AskTeachersTable is the table the holds the ask_teachers relation/edge.
	AskTeachersTable = "user_ask_answers"
	// AskTeachersInverseTable is the table name for the UserAskAnswer entity.
	// It exists in this package in order to avoid circular dependency with the "useraskanswer" package.
	AskTeachersInverseTable = "user_ask_answers"
	// AskTeachersColumn is the table column denoting the ask_teachers relation/edge.
	AskTeachersColumn = "teacher_id"
	// AttachmentTable is the table the holds the attachment relation/edge.
	AttachmentTable = "teachers"
	// AttachmentInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentInverseTable = "attachments"
	// AttachmentColumn is the table column denoting the attachment relation/edge.
	AttachmentColumn = "avatar_id"
)

// Columns holds all SQL columns for teacher fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldSex,
	FieldEmail,
	FieldPhone,
	FieldNickname,
	FieldSubTitle,
	FieldDetail,
	FieldStatus,
	FieldTeachingAge,
	FieldAvatarID,
	FieldSortOrder,
}

var (
	// MajorsPrimaryKey and MajorsColumn2 are the table columns denoting the
	// primary key for the majors relation (M2M).
	MajorsPrimaryKey = []string{"teacher_id", "major_id"}
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
	// DefaultSex holds the default value on creation for the "sex" field.
	DefaultSex uint8
	// DefaultEmail holds the default value on creation for the "email" field.
	DefaultEmail string
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// DefaultNickname holds the default value on creation for the "nickname" field.
	DefaultNickname string
	// DefaultSubTitle holds the default value on creation for the "sub_title" field.
	DefaultSubTitle string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
	// DefaultTeachingAge holds the default value on creation for the "teaching_age" field.
	DefaultTeachingAge uint8
	// DefaultSortOrder holds the default value on creation for the "sort_order" field.
	DefaultSortOrder int
)
