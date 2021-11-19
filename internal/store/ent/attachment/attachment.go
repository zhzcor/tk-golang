// Code generated by entc, DO NOT EDIT.

package attachment

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the attachment type in the database.
	Label = "attachment"
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
	// FieldFilename holds the string denoting the filename field in the database.
	FieldFilename = "filename"
	// FieldFileSize holds the string denoting the file_size field in the database.
	FieldFileSize = "file_size"
	// FieldMimeType holds the string denoting the mime_type field in the database.
	FieldMimeType = "mime_type"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAdminID holds the string denoting the admin_id field in the database.
	FieldAdminID = "admin_id"
	// EdgeMajorDetailCoverImg holds the string denoting the major_detail_cover_img edge name in mutations.
	EdgeMajorDetailCoverImg = "major_detail_cover_img"
	// EdgeMajorDetailSubjectImg holds the string denoting the major_detail_subject_img edge name in mutations.
	EdgeMajorDetailSubjectImg = "major_detail_subject_img"
	// EdgeMajorTeacherAttachment holds the string denoting the major_teacher_attachment edge name in mutations.
	EdgeMajorTeacherAttachment = "major_teacher_attachment"
	// EdgeMajorServiceAttachment holds the string denoting the major_service_attachment edge name in mutations.
	EdgeMajorServiceAttachment = "major_service_attachment"
	// EdgeAdvertise holds the string denoting the advertise edge name in mutations.
	EdgeAdvertise = "advertise"
	// EdgeMessageAttachment holds the string denoting the message_attachment edge name in mutations.
	EdgeMessageAttachment = "message_attachment"
	// EdgeSharePosterAttachments holds the string denoting the share_poster_attachments edge name in mutations.
	EdgeSharePosterAttachments = "share_poster_attachments"
	// EdgeTeacherAttachments holds the string denoting the teacher_attachments edge name in mutations.
	EdgeTeacherAttachments = "teacher_attachments"
	// EdgeAdminImgID holds the string denoting the admin_img_id edge name in mutations.
	EdgeAdminImgID = "admin_img_id"
	// EdgeSmallCategoryAttachments holds the string denoting the small_category_attachments edge name in mutations.
	EdgeSmallCategoryAttachments = "small_category_attachments"
	// EdgeCourseAttachments holds the string denoting the course_attachments edge name in mutations.
	EdgeCourseAttachments = "course_attachments"
	// EdgeClassCoverAttachments holds the string denoting the class_cover_attachments edge name in mutations.
	EdgeClassCoverAttachments = "class_cover_attachments"
	// EdgeCoursewareAttachment holds the string denoting the courseware_attachment edge name in mutations.
	EdgeCoursewareAttachment = "courseware_attachment"
	// EdgeOrderAttachment holds the string denoting the order_attachment edge name in mutations.
	EdgeOrderAttachment = "order_attachment"
	// EdgeVideoTaskAttachment holds the string denoting the video_task_attachment edge name in mutations.
	EdgeVideoTaskAttachment = "video_task_attachment"
	// EdgeAskAttachments holds the string denoting the ask_attachments edge name in mutations.
	EdgeAskAttachments = "ask_attachments"
	// EdgeGroupCard holds the string denoting the group_card edge name in mutations.
	EdgeGroupCard = "group_card"
	// Table holds the table name of the attachment in the database.
	Table = "attachments"
	// MajorDetailCoverImgTable is the table the holds the major_detail_cover_img relation/edge.
	MajorDetailCoverImgTable = "major_details"
	// MajorDetailCoverImgInverseTable is the table name for the MajorDetail entity.
	// It exists in this package in order to avoid circular dependency with the "majordetail" package.
	MajorDetailCoverImgInverseTable = "major_details"
	// MajorDetailCoverImgColumn is the table column denoting the major_detail_cover_img relation/edge.
	MajorDetailCoverImgColumn = "cover_img_id"
	// MajorDetailSubjectImgTable is the table the holds the major_detail_subject_img relation/edge.
	MajorDetailSubjectImgTable = "major_details"
	// MajorDetailSubjectImgInverseTable is the table name for the MajorDetail entity.
	// It exists in this package in order to avoid circular dependency with the "majordetail" package.
	MajorDetailSubjectImgInverseTable = "major_details"
	// MajorDetailSubjectImgColumn is the table column denoting the major_detail_subject_img relation/edge.
	MajorDetailSubjectImgColumn = "subject_img_id"
	// MajorTeacherAttachmentTable is the table the holds the major_teacher_attachment relation/edge. The primary key declared below.
	MajorTeacherAttachmentTable = "attachment_major_teacher_attachment"
	// MajorTeacherAttachmentInverseTable is the table name for the MajorDetail entity.
	// It exists in this package in order to avoid circular dependency with the "majordetail" package.
	MajorTeacherAttachmentInverseTable = "major_details"
	// MajorServiceAttachmentTable is the table the holds the major_service_attachment relation/edge. The primary key declared below.
	MajorServiceAttachmentTable = "attachment_major_service_attachment"
	// MajorServiceAttachmentInverseTable is the table name for the MajorDetail entity.
	// It exists in this package in order to avoid circular dependency with the "majordetail" package.
	MajorServiceAttachmentInverseTable = "major_details"
	// AdvertiseTable is the table the holds the advertise relation/edge.
	AdvertiseTable = "advertises"
	// AdvertiseInverseTable is the table name for the Advertise entity.
	// It exists in this package in order to avoid circular dependency with the "advertise" package.
	AdvertiseInverseTable = "advertises"
	// AdvertiseColumn is the table column denoting the advertise relation/edge.
	AdvertiseColumn = "attachment_id"
	// MessageAttachmentTable is the table the holds the message_attachment relation/edge.
	MessageAttachmentTable = "messages"
	// MessageAttachmentInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessageAttachmentInverseTable = "messages"
	// MessageAttachmentColumn is the table column denoting the message_attachment relation/edge.
	MessageAttachmentColumn = "attachment_id"
	// SharePosterAttachmentsTable is the table the holds the share_poster_attachments relation/edge.
	SharePosterAttachmentsTable = "share_posters"
	// SharePosterAttachmentsInverseTable is the table name for the SharePoster entity.
	// It exists in this package in order to avoid circular dependency with the "shareposter" package.
	SharePosterAttachmentsInverseTable = "share_posters"
	// SharePosterAttachmentsColumn is the table column denoting the share_poster_attachments relation/edge.
	SharePosterAttachmentsColumn = "share_poster_img_id"
	// TeacherAttachmentsTable is the table the holds the teacher_attachments relation/edge.
	TeacherAttachmentsTable = "teachers"
	// TeacherAttachmentsInverseTable is the table name for the Teacher entity.
	// It exists in this package in order to avoid circular dependency with the "teacher" package.
	TeacherAttachmentsInverseTable = "teachers"
	// TeacherAttachmentsColumn is the table column denoting the teacher_attachments relation/edge.
	TeacherAttachmentsColumn = "avatar_id"
	// AdminImgIDTable is the table the holds the admin_img_id relation/edge.
	AdminImgIDTable = "admins"
	// AdminImgIDInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	AdminImgIDInverseTable = "admins"
	// AdminImgIDColumn is the table column denoting the admin_img_id relation/edge.
	AdminImgIDColumn = "admin_avatar_id"
	// SmallCategoryAttachmentsTable is the table the holds the small_category_attachments relation/edge.
	SmallCategoryAttachmentsTable = "kc_small_category_attachments"
	// SmallCategoryAttachmentsInverseTable is the table name for the KcSmallCategoryAttachment entity.
	// It exists in this package in order to avoid circular dependency with the "kcsmallcategoryattachment" package.
	SmallCategoryAttachmentsInverseTable = "kc_small_category_attachments"
	// SmallCategoryAttachmentsColumn is the table column denoting the small_category_attachments relation/edge.
	SmallCategoryAttachmentsColumn = "attachment_id"
	// CourseAttachmentsTable is the table the holds the course_attachments relation/edge.
	CourseAttachmentsTable = "kc_courses"
	// CourseAttachmentsInverseTable is the table name for the KcCourse entity.
	// It exists in this package in order to avoid circular dependency with the "kccourse" package.
	CourseAttachmentsInverseTable = "kc_courses"
	// CourseAttachmentsColumn is the table column denoting the course_attachments relation/edge.
	CourseAttachmentsColumn = "course_cover_img_id"
	// ClassCoverAttachmentsTable is the table the holds the class_cover_attachments relation/edge.
	ClassCoverAttachmentsTable = "kc_classes"
	// ClassCoverAttachmentsInverseTable is the table name for the KcClass entity.
	// It exists in this package in order to avoid circular dependency with the "kcclass" package.
	ClassCoverAttachmentsInverseTable = "kc_classes"
	// ClassCoverAttachmentsColumn is the table column denoting the class_cover_attachments relation/edge.
	ClassCoverAttachmentsColumn = "class_cover_img_id"
	// CoursewareAttachmentTable is the table the holds the courseware_attachment relation/edge.
	CoursewareAttachmentTable = "kc_course_small_categories"
	// CoursewareAttachmentInverseTable is the table name for the KcCourseSmallCategory entity.
	// It exists in this package in order to avoid circular dependency with the "kccoursesmallcategory" package.
	CoursewareAttachmentInverseTable = "kc_course_small_categories"
	// CoursewareAttachmentColumn is the table column denoting the courseware_attachment relation/edge.
	CoursewareAttachmentColumn = "courseware_attach_id"
	// OrderAttachmentTable is the table the holds the order_attachment relation/edge.
	OrderAttachmentTable = "kc_course_small_categories"
	// OrderAttachmentInverseTable is the table name for the KcCourseSmallCategory entity.
	// It exists in this package in order to avoid circular dependency with the "kccoursesmallcategory" package.
	OrderAttachmentInverseTable = "kc_course_small_categories"
	// OrderAttachmentColumn is the table column denoting the order_attachment relation/edge.
	OrderAttachmentColumn = "order_video_attach_id"
	// VideoTaskAttachmentTable is the table the holds the video_task_attachment relation/edge.
	VideoTaskAttachmentTable = "kc_video_upload_tasks"
	// VideoTaskAttachmentInverseTable is the table name for the KcVideoUploadTask entity.
	// It exists in this package in order to avoid circular dependency with the "kcvideouploadtask" package.
	VideoTaskAttachmentInverseTable = "kc_video_upload_tasks"
	// VideoTaskAttachmentColumn is the table column denoting the video_task_attachment relation/edge.
	VideoTaskAttachmentColumn = "attachment_id"
	// AskAttachmentsTable is the table the holds the ask_attachments relation/edge.
	AskAttachmentsTable = "user_ask_answer_attachments"
	// AskAttachmentsInverseTable is the table name for the UserAskAnswerAttachment entity.
	// It exists in this package in order to avoid circular dependency with the "useraskanswerattachment" package.
	AskAttachmentsInverseTable = "user_ask_answer_attachments"
	// AskAttachmentsColumn is the table column denoting the ask_attachments relation/edge.
	AskAttachmentsColumn = "attachment_id"
	// GroupCardTable is the table the holds the group_card relation/edge.
	GroupCardTable = "group_cards"
	// GroupCardInverseTable is the table name for the GroupCard entity.
	// It exists in this package in order to avoid circular dependency with the "groupcard" package.
	GroupCardInverseTable = "group_cards"
	// GroupCardColumn is the table column denoting the group_card relation/edge.
	GroupCardColumn = "attachment_id"
)

// Columns holds all SQL columns for attachment fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldFilename,
	FieldFileSize,
	FieldMimeType,
	FieldUserID,
	FieldAdminID,
}

var (
	// MajorTeacherAttachmentPrimaryKey and MajorTeacherAttachmentColumn2 are the table columns denoting the
	// primary key for the major_teacher_attachment relation (M2M).
	MajorTeacherAttachmentPrimaryKey = []string{"attachment_id", "major_detail_id"}
	// MajorServiceAttachmentPrimaryKey and MajorServiceAttachmentColumn2 are the table columns denoting the
	// primary key for the major_service_attachment relation (M2M).
	MajorServiceAttachmentPrimaryKey = []string{"attachment_id", "major_detail_id"}
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
	// DefaultFilename holds the default value on creation for the "filename" field.
	DefaultFilename string
	// DefaultFileSize holds the default value on creation for the "file_size" field.
	DefaultFileSize uint32
	// DefaultMimeType holds the default value on creation for the "mime_type" field.
	DefaultMimeType string
)
