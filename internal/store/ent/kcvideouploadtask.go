// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"tkserver/internal/store/ent/attachment"
	"tkserver/internal/store/ent/kcvideouploadtask"

	"entgo.io/ent/dialect/sql"
)

// KcVideoUploadTask is the model entity for the KcVideoUploadTask schema.
type KcVideoUploadTask struct {
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
	// VideoID holds the value of the "video_id" field.
	// 视频id
	VideoID int `json:"video_id"`
	// AttachmentID holds the value of the "attachment_id" field.
	// 视频oss id
	AttachmentID int `json:"attachment_id"`
	// Type holds the value of the "type" field.
	// 类型 1：点播，2：伪直播 3：回放
	Type uint8 `json:"type"`
	// CourseID holds the value of the "course_id" field.
	// 课程id
	CourseID int `json:"course_id"`
	// Status holds the value of the "status" field.
	// 任务状态:1:未开始，2：上传中，3：上传成功，4：上传失败，5：转码成功，6：转码失败，7：正在生成回放
	Status int `json:"status"`
	// Remark holds the value of the "remark" field.
	// 任务描述
	Remark string `json:"remark"`
	// Title holds the value of the "title" field.
	// 标题
	Title string `json:"title"`
	// VideoName holds the value of the "video_name" field.
	// video名称
	VideoName string `json:"video_name"`
	// TotalSize holds the value of the "total_size" field.
	// 视频大小，单位是字节
	TotalSize int `json:"total_size"`
	// Length holds the value of the "length" field.
	// 视频时长，单位为秒（转码成功回调才有）
	Length int `json:"length"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the KcVideoUploadTaskQuery when eager-loading is set.
	Edges KcVideoUploadTaskEdges `json:"edges"`
}

// KcVideoUploadTaskEdges holds the relations/edges for other nodes in the graph.
type KcVideoUploadTaskEdges struct {
	// Attachment holds the value of the attachment edge.
	Attachment *Attachment `json:"attachment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AttachmentOrErr returns the Attachment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KcVideoUploadTaskEdges) AttachmentOrErr() (*Attachment, error) {
	if e.loadedTypes[0] {
		if e.Attachment == nil {
			// The edge attachment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: attachment.Label}
		}
		return e.Attachment, nil
	}
	return nil, &NotLoadedError{edge: "attachment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KcVideoUploadTask) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case kcvideouploadtask.FieldID, kcvideouploadtask.FieldVideoID, kcvideouploadtask.FieldAttachmentID, kcvideouploadtask.FieldType, kcvideouploadtask.FieldCourseID, kcvideouploadtask.FieldStatus, kcvideouploadtask.FieldTotalSize, kcvideouploadtask.FieldLength:
			values[i] = new(sql.NullInt64)
		case kcvideouploadtask.FieldUUID, kcvideouploadtask.FieldRemark, kcvideouploadtask.FieldTitle, kcvideouploadtask.FieldVideoName:
			values[i] = new(sql.NullString)
		case kcvideouploadtask.FieldCreatedAt, kcvideouploadtask.FieldUpdatedAt, kcvideouploadtask.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type KcVideoUploadTask", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KcVideoUploadTask fields.
func (kvut *KcVideoUploadTask) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case kcvideouploadtask.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			kvut.ID = int(value.Int64)
		case kcvideouploadtask.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				kvut.UUID = value.String
			}
		case kcvideouploadtask.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				kvut.CreatedAt = new(time.Time)
				*kvut.CreatedAt = value.Time
			}
		case kcvideouploadtask.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				kvut.UpdatedAt = new(time.Time)
				*kvut.UpdatedAt = value.Time
			}
		case kcvideouploadtask.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				kvut.DeletedAt = new(time.Time)
				*kvut.DeletedAt = value.Time
			}
		case kcvideouploadtask.FieldVideoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field video_id", values[i])
			} else if value.Valid {
				kvut.VideoID = int(value.Int64)
			}
		case kcvideouploadtask.FieldAttachmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attachment_id", values[i])
			} else if value.Valid {
				kvut.AttachmentID = int(value.Int64)
			}
		case kcvideouploadtask.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				kvut.Type = uint8(value.Int64)
			}
		case kcvideouploadtask.FieldCourseID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field course_id", values[i])
			} else if value.Valid {
				kvut.CourseID = int(value.Int64)
			}
		case kcvideouploadtask.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				kvut.Status = int(value.Int64)
			}
		case kcvideouploadtask.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				kvut.Remark = value.String
			}
		case kcvideouploadtask.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				kvut.Title = value.String
			}
		case kcvideouploadtask.FieldVideoName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_name", values[i])
			} else if value.Valid {
				kvut.VideoName = value.String
			}
		case kcvideouploadtask.FieldTotalSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_size", values[i])
			} else if value.Valid {
				kvut.TotalSize = int(value.Int64)
			}
		case kcvideouploadtask.FieldLength:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field length", values[i])
			} else if value.Valid {
				kvut.Length = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAttachment queries the "attachment" edge of the KcVideoUploadTask entity.
func (kvut *KcVideoUploadTask) QueryAttachment() *AttachmentQuery {
	return (&KcVideoUploadTaskClient{config: kvut.config}).QueryAttachment(kvut)
}

// Update returns a builder for updating this KcVideoUploadTask.
// Note that you need to call KcVideoUploadTask.Unwrap() before calling this method if this KcVideoUploadTask
// was returned from a transaction, and the transaction was committed or rolled back.
func (kvut *KcVideoUploadTask) Update() *KcVideoUploadTaskUpdateOne {
	return (&KcVideoUploadTaskClient{config: kvut.config}).UpdateOne(kvut)
}

// Unwrap unwraps the KcVideoUploadTask entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kvut *KcVideoUploadTask) Unwrap() *KcVideoUploadTask {
	tx, ok := kvut.config.driver.(*txDriver)
	if !ok {
		panic("ent: KcVideoUploadTask is not a transactional entity")
	}
	kvut.config.driver = tx.drv
	return kvut
}

// String implements the fmt.Stringer.
func (kvut *KcVideoUploadTask) String() string {
	var builder strings.Builder
	builder.WriteString("KcVideoUploadTask(")
	builder.WriteString(fmt.Sprintf("id=%v", kvut.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(kvut.UUID)
	if v := kvut.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := kvut.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := kvut.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", video_id=")
	builder.WriteString(fmt.Sprintf("%v", kvut.VideoID))
	builder.WriteString(", attachment_id=")
	builder.WriteString(fmt.Sprintf("%v", kvut.AttachmentID))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", kvut.Type))
	builder.WriteString(", course_id=")
	builder.WriteString(fmt.Sprintf("%v", kvut.CourseID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", kvut.Status))
	builder.WriteString(", remark=")
	builder.WriteString(kvut.Remark)
	builder.WriteString(", title=")
	builder.WriteString(kvut.Title)
	builder.WriteString(", video_name=")
	builder.WriteString(kvut.VideoName)
	builder.WriteString(", total_size=")
	builder.WriteString(fmt.Sprintf("%v", kvut.TotalSize))
	builder.WriteString(", length=")
	builder.WriteString(fmt.Sprintf("%v", kvut.Length))
	builder.WriteByte(')')
	return builder.String()
}

// KcVideoUploadTasks is a parsable slice of KcVideoUploadTask.
type KcVideoUploadTasks []*KcVideoUploadTask

func (kvut KcVideoUploadTasks) config(cfg config) {
	for _i := range kvut {
		kvut[_i].config = cfg
	}
}
