// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"tkserver/internal/store/ent/admin"
	"tkserver/internal/store/ent/attachment"
	"tkserver/internal/store/ent/kccoursesmallcategory"

	"entgo.io/ent/dialect/sql"
)

// Attachment is the model entity for the Attachment schema.
type Attachment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at"`
	// Filename holds the value of the "filename" field.
	// 文件名称
	Filename string `json:"filename"`
	// FileSize holds the value of the "file_size" field.
	// 文件大小,字节
	FileSize uint32 `json:"file_size"`
	// MimeType holds the value of the "mime_type" field.
	// 文件类型
	MimeType string `json:"mime_type"`
	// UserID holds the value of the "user_id" field.
	// 上传时的用户ID
	UserID int `json:"user_id"`
	// AdminID holds the value of the "admin_id" field.
	// 如果是后台上传，上传时的管理员ID
	AdminID int `json:"admin_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttachmentQuery when eager-loading is set.
	Edges AttachmentEdges `json:"edges"`
}

// AttachmentEdges holds the relations/edges for other nodes in the graph.
type AttachmentEdges struct {
	// MajorDetailCoverImg holds the value of the major_detail_cover_img edge.
	MajorDetailCoverImg []*MajorDetail `json:"major_detail_cover_img,omitempty"`
	// MajorDetailSubjectImg holds the value of the major_detail_subject_img edge.
	MajorDetailSubjectImg []*MajorDetail `json:"major_detail_subject_img,omitempty"`
	// MajorTeacherAttachment holds the value of the major_teacher_attachment edge.
	MajorTeacherAttachment []*MajorDetail `json:"major_teacher_attachment,omitempty"`
	// MajorServiceAttachment holds the value of the major_service_attachment edge.
	MajorServiceAttachment []*MajorDetail `json:"major_service_attachment,omitempty"`
	// Advertise holds the value of the advertise edge.
	Advertise []*Advertise `json:"advertise,omitempty"`
	// MessageAttachment holds the value of the message_attachment edge.
	MessageAttachment []*Message `json:"message_attachment,omitempty"`
	// SharePosterAttachments holds the value of the share_poster_attachments edge.
	SharePosterAttachments []*SharePoster `json:"share_poster_attachments,omitempty"`
	// TeacherAttachments holds the value of the teacher_attachments edge.
	TeacherAttachments []*Teacher `json:"teacher_attachments,omitempty"`
	// AdminImgID holds the value of the admin_img_id edge.
	AdminImgID *Admin `json:"admin_img_id,omitempty"`
	// SmallCategoryAttachments holds the value of the small_category_attachments edge.
	SmallCategoryAttachments []*KcSmallCategoryAttachment `json:"small_category_attachments,omitempty"`
	// CourseAttachments holds the value of the course_attachments edge.
	CourseAttachments []*KcCourse `json:"course_attachments,omitempty"`
	// ClassCoverAttachments holds the value of the class_cover_attachments edge.
	ClassCoverAttachments []*KcClass `json:"class_cover_attachments,omitempty"`
	// CoursewareAttachment holds the value of the courseware_attachment edge.
	CoursewareAttachment *KcCourseSmallCategory `json:"courseware_attachment,omitempty"`
	// OrderAttachment holds the value of the order_attachment edge.
	OrderAttachment *KcCourseSmallCategory `json:"order_attachment,omitempty"`
	// VideoTaskAttachment holds the value of the video_task_attachment edge.
	VideoTaskAttachment []*KcVideoUploadTask `json:"video_task_attachment,omitempty"`
	// AskAttachments holds the value of the ask_attachments edge.
	AskAttachments []*UserAskAnswerAttachment `json:"ask_attachments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [16]bool
}

// MajorDetailCoverImgOrErr returns the MajorDetailCoverImg value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) MajorDetailCoverImgOrErr() ([]*MajorDetail, error) {
	if e.loadedTypes[0] {
		return e.MajorDetailCoverImg, nil
	}
	return nil, &NotLoadedError{edge: "major_detail_cover_img"}
}

// MajorDetailSubjectImgOrErr returns the MajorDetailSubjectImg value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) MajorDetailSubjectImgOrErr() ([]*MajorDetail, error) {
	if e.loadedTypes[1] {
		return e.MajorDetailSubjectImg, nil
	}
	return nil, &NotLoadedError{edge: "major_detail_subject_img"}
}

// MajorTeacherAttachmentOrErr returns the MajorTeacherAttachment value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) MajorTeacherAttachmentOrErr() ([]*MajorDetail, error) {
	if e.loadedTypes[2] {
		return e.MajorTeacherAttachment, nil
	}
	return nil, &NotLoadedError{edge: "major_teacher_attachment"}
}

// MajorServiceAttachmentOrErr returns the MajorServiceAttachment value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) MajorServiceAttachmentOrErr() ([]*MajorDetail, error) {
	if e.loadedTypes[3] {
		return e.MajorServiceAttachment, nil
	}
	return nil, &NotLoadedError{edge: "major_service_attachment"}
}

// AdvertiseOrErr returns the Advertise value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) AdvertiseOrErr() ([]*Advertise, error) {
	if e.loadedTypes[4] {
		return e.Advertise, nil
	}
	return nil, &NotLoadedError{edge: "advertise"}
}

// MessageAttachmentOrErr returns the MessageAttachment value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) MessageAttachmentOrErr() ([]*Message, error) {
	if e.loadedTypes[5] {
		return e.MessageAttachment, nil
	}
	return nil, &NotLoadedError{edge: "message_attachment"}
}

// SharePosterAttachmentsOrErr returns the SharePosterAttachments value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) SharePosterAttachmentsOrErr() ([]*SharePoster, error) {
	if e.loadedTypes[6] {
		return e.SharePosterAttachments, nil
	}
	return nil, &NotLoadedError{edge: "share_poster_attachments"}
}

// TeacherAttachmentsOrErr returns the TeacherAttachments value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) TeacherAttachmentsOrErr() ([]*Teacher, error) {
	if e.loadedTypes[7] {
		return e.TeacherAttachments, nil
	}
	return nil, &NotLoadedError{edge: "teacher_attachments"}
}

// AdminImgIDOrErr returns the AdminImgID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttachmentEdges) AdminImgIDOrErr() (*Admin, error) {
	if e.loadedTypes[8] {
		if e.AdminImgID == nil {
			// The edge admin_img_id was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: admin.Label}
		}
		return e.AdminImgID, nil
	}
	return nil, &NotLoadedError{edge: "admin_img_id"}
}

// SmallCategoryAttachmentsOrErr returns the SmallCategoryAttachments value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) SmallCategoryAttachmentsOrErr() ([]*KcSmallCategoryAttachment, error) {
	if e.loadedTypes[9] {
		return e.SmallCategoryAttachments, nil
	}
	return nil, &NotLoadedError{edge: "small_category_attachments"}
}

// CourseAttachmentsOrErr returns the CourseAttachments value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) CourseAttachmentsOrErr() ([]*KcCourse, error) {
	if e.loadedTypes[10] {
		return e.CourseAttachments, nil
	}
	return nil, &NotLoadedError{edge: "course_attachments"}
}

// ClassCoverAttachmentsOrErr returns the ClassCoverAttachments value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) ClassCoverAttachmentsOrErr() ([]*KcClass, error) {
	if e.loadedTypes[11] {
		return e.ClassCoverAttachments, nil
	}
	return nil, &NotLoadedError{edge: "class_cover_attachments"}
}

// CoursewareAttachmentOrErr returns the CoursewareAttachment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttachmentEdges) CoursewareAttachmentOrErr() (*KcCourseSmallCategory, error) {
	if e.loadedTypes[12] {
		if e.CoursewareAttachment == nil {
			// The edge courseware_attachment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: kccoursesmallcategory.Label}
		}
		return e.CoursewareAttachment, nil
	}
	return nil, &NotLoadedError{edge: "courseware_attachment"}
}

// OrderAttachmentOrErr returns the OrderAttachment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttachmentEdges) OrderAttachmentOrErr() (*KcCourseSmallCategory, error) {
	if e.loadedTypes[13] {
		if e.OrderAttachment == nil {
			// The edge order_attachment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: kccoursesmallcategory.Label}
		}
		return e.OrderAttachment, nil
	}
	return nil, &NotLoadedError{edge: "order_attachment"}
}

// VideoTaskAttachmentOrErr returns the VideoTaskAttachment value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) VideoTaskAttachmentOrErr() ([]*KcVideoUploadTask, error) {
	if e.loadedTypes[14] {
		return e.VideoTaskAttachment, nil
	}
	return nil, &NotLoadedError{edge: "video_task_attachment"}
}

// AskAttachmentsOrErr returns the AskAttachments value or an error if the edge
// was not loaded in eager-loading.
func (e AttachmentEdges) AskAttachmentsOrErr() ([]*UserAskAnswerAttachment, error) {
	if e.loadedTypes[15] {
		return e.AskAttachments, nil
	}
	return nil, &NotLoadedError{edge: "ask_attachments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attachment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case attachment.FieldID, attachment.FieldFileSize, attachment.FieldUserID, attachment.FieldAdminID:
			values[i] = new(sql.NullInt64)
		case attachment.FieldUUID, attachment.FieldFilename, attachment.FieldMimeType:
			values[i] = new(sql.NullString)
		case attachment.FieldCreatedAt, attachment.FieldUpdatedAt, attachment.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Attachment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attachment fields.
func (a *Attachment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attachment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case attachment.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				a.UUID = value.String
			}
		case attachment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = new(time.Time)
				*a.CreatedAt = value.Time
			}
		case attachment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = new(time.Time)
				*a.UpdatedAt = value.Time
			}
		case attachment.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = new(time.Time)
				*a.DeletedAt = value.Time
			}
		case attachment.FieldFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filename", values[i])
			} else if value.Valid {
				a.Filename = value.String
			}
		case attachment.FieldFileSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field file_size", values[i])
			} else if value.Valid {
				a.FileSize = uint32(value.Int64)
			}
		case attachment.FieldMimeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mime_type", values[i])
			} else if value.Valid {
				a.MimeType = value.String
			}
		case attachment.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				a.UserID = int(value.Int64)
			}
		case attachment.FieldAdminID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field admin_id", values[i])
			} else if value.Valid {
				a.AdminID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMajorDetailCoverImg queries the "major_detail_cover_img" edge of the Attachment entity.
func (a *Attachment) QueryMajorDetailCoverImg() *MajorDetailQuery {
	return (&AttachmentClient{config: a.config}).QueryMajorDetailCoverImg(a)
}

// QueryMajorDetailSubjectImg queries the "major_detail_subject_img" edge of the Attachment entity.
func (a *Attachment) QueryMajorDetailSubjectImg() *MajorDetailQuery {
	return (&AttachmentClient{config: a.config}).QueryMajorDetailSubjectImg(a)
}

// QueryMajorTeacherAttachment queries the "major_teacher_attachment" edge of the Attachment entity.
func (a *Attachment) QueryMajorTeacherAttachment() *MajorDetailQuery {
	return (&AttachmentClient{config: a.config}).QueryMajorTeacherAttachment(a)
}

// QueryMajorServiceAttachment queries the "major_service_attachment" edge of the Attachment entity.
func (a *Attachment) QueryMajorServiceAttachment() *MajorDetailQuery {
	return (&AttachmentClient{config: a.config}).QueryMajorServiceAttachment(a)
}

// QueryAdvertise queries the "advertise" edge of the Attachment entity.
func (a *Attachment) QueryAdvertise() *AdvertiseQuery {
	return (&AttachmentClient{config: a.config}).QueryAdvertise(a)
}

// QueryMessageAttachment queries the "message_attachment" edge of the Attachment entity.
func (a *Attachment) QueryMessageAttachment() *MessageQuery {
	return (&AttachmentClient{config: a.config}).QueryMessageAttachment(a)
}

// QuerySharePosterAttachments queries the "share_poster_attachments" edge of the Attachment entity.
func (a *Attachment) QuerySharePosterAttachments() *SharePosterQuery {
	return (&AttachmentClient{config: a.config}).QuerySharePosterAttachments(a)
}

// QueryTeacherAttachments queries the "teacher_attachments" edge of the Attachment entity.
func (a *Attachment) QueryTeacherAttachments() *TeacherQuery {
	return (&AttachmentClient{config: a.config}).QueryTeacherAttachments(a)
}

// QueryAdminImgID queries the "admin_img_id" edge of the Attachment entity.
func (a *Attachment) QueryAdminImgID() *AdminQuery {
	return (&AttachmentClient{config: a.config}).QueryAdminImgID(a)
}

// QuerySmallCategoryAttachments queries the "small_category_attachments" edge of the Attachment entity.
func (a *Attachment) QuerySmallCategoryAttachments() *KcSmallCategoryAttachmentQuery {
	return (&AttachmentClient{config: a.config}).QuerySmallCategoryAttachments(a)
}

// QueryCourseAttachments queries the "course_attachments" edge of the Attachment entity.
func (a *Attachment) QueryCourseAttachments() *KcCourseQuery {
	return (&AttachmentClient{config: a.config}).QueryCourseAttachments(a)
}

// QueryClassCoverAttachments queries the "class_cover_attachments" edge of the Attachment entity.
func (a *Attachment) QueryClassCoverAttachments() *KcClassQuery {
	return (&AttachmentClient{config: a.config}).QueryClassCoverAttachments(a)
}

// QueryCoursewareAttachment queries the "courseware_attachment" edge of the Attachment entity.
func (a *Attachment) QueryCoursewareAttachment() *KcCourseSmallCategoryQuery {
	return (&AttachmentClient{config: a.config}).QueryCoursewareAttachment(a)
}

// QueryOrderAttachment queries the "order_attachment" edge of the Attachment entity.
func (a *Attachment) QueryOrderAttachment() *KcCourseSmallCategoryQuery {
	return (&AttachmentClient{config: a.config}).QueryOrderAttachment(a)
}

// QueryVideoTaskAttachment queries the "video_task_attachment" edge of the Attachment entity.
func (a *Attachment) QueryVideoTaskAttachment() *KcVideoUploadTaskQuery {
	return (&AttachmentClient{config: a.config}).QueryVideoTaskAttachment(a)
}

// QueryAskAttachments queries the "ask_attachments" edge of the Attachment entity.
func (a *Attachment) QueryAskAttachments() *UserAskAnswerAttachmentQuery {
	return (&AttachmentClient{config: a.config}).QueryAskAttachments(a)
}

// Update returns a builder for updating this Attachment.
// Note that you need to call Attachment.Unwrap() before calling this method if this Attachment
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attachment) Update() *AttachmentUpdateOne {
	return (&AttachmentClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Attachment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attachment) Unwrap() *Attachment {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attachment is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attachment) String() string {
	var builder strings.Builder
	builder.WriteString("Attachment(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(a.UUID)
	if v := a.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := a.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := a.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", filename=")
	builder.WriteString(a.Filename)
	builder.WriteString(", file_size=")
	builder.WriteString(fmt.Sprintf("%v", a.FileSize))
	builder.WriteString(", mime_type=")
	builder.WriteString(a.MimeType)
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteString(", admin_id=")
	builder.WriteString(fmt.Sprintf("%v", a.AdminID))
	builder.WriteByte(')')
	return builder.String()
}

// Attachments is a parsable slice of Attachment.
type Attachments []*Attachment

func (a Attachments) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
