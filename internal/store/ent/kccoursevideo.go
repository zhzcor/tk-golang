// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gserver/internal/store/ent/kccoursevideo"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// KcCourseVideo is the model entity for the KcCourseVideo schema.
type KcCourseVideo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// VideoTitle holds the value of the "video_title" field.
	// 名称
	VideoTitle string `json:"video_title"`
	// FileType holds the value of the "file_type" field.
	// 1:视频，2:音频，后面其他的自己定义拓展
	FileType int8 `json:"file_type"`
	// FilePath holds the value of the "file_path" field.
	// 文件路径
	FilePath string `json:"file_path"`
	// SubtitlePath holds the value of the "subtitle_path" field.
	// 字幕文件路径
	SubtitlePath string `json:"subtitle_path"`
	// VideoTime holds the value of the "video_time" field.
	// 视频时长
	VideoTime string `json:"video_time"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KcCourseVideo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case kccoursevideo.FieldID, kccoursevideo.FieldFileType:
			values[i] = new(sql.NullInt64)
		case kccoursevideo.FieldVideoTitle, kccoursevideo.FieldFilePath, kccoursevideo.FieldSubtitlePath, kccoursevideo.FieldVideoTime:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type KcCourseVideo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KcCourseVideo fields.
func (kcv *KcCourseVideo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case kccoursevideo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			kcv.ID = int(value.Int64)
		case kccoursevideo.FieldVideoTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_title", values[i])
			} else if value.Valid {
				kcv.VideoTitle = value.String
			}
		case kccoursevideo.FieldFileType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field file_type", values[i])
			} else if value.Valid {
				kcv.FileType = int8(value.Int64)
			}
		case kccoursevideo.FieldFilePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_path", values[i])
			} else if value.Valid {
				kcv.FilePath = value.String
			}
		case kccoursevideo.FieldSubtitlePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subtitle_path", values[i])
			} else if value.Valid {
				kcv.SubtitlePath = value.String
			}
		case kccoursevideo.FieldVideoTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_time", values[i])
			} else if value.Valid {
				kcv.VideoTime = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this KcCourseVideo.
// Note that you need to call KcCourseVideo.Unwrap() before calling this method if this KcCourseVideo
// was returned from a transaction, and the transaction was committed or rolled back.
func (kcv *KcCourseVideo) Update() *KcCourseVideoUpdateOne {
	return (&KcCourseVideoClient{config: kcv.config}).UpdateOne(kcv)
}

// Unwrap unwraps the KcCourseVideo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kcv *KcCourseVideo) Unwrap() *KcCourseVideo {
	tx, ok := kcv.config.driver.(*txDriver)
	if !ok {
		panic("ent: KcCourseVideo is not a transactional entity")
	}
	kcv.config.driver = tx.drv
	return kcv
}

// String implements the fmt.Stringer.
func (kcv *KcCourseVideo) String() string {
	var builder strings.Builder
	builder.WriteString("KcCourseVideo(")
	builder.WriteString(fmt.Sprintf("id=%v", kcv.ID))
	builder.WriteString(", video_title=")
	builder.WriteString(kcv.VideoTitle)
	builder.WriteString(", file_type=")
	builder.WriteString(fmt.Sprintf("%v", kcv.FileType))
	builder.WriteString(", file_path=")
	builder.WriteString(kcv.FilePath)
	builder.WriteString(", subtitle_path=")
	builder.WriteString(kcv.SubtitlePath)
	builder.WriteString(", video_time=")
	builder.WriteString(kcv.VideoTime)
	builder.WriteByte(')')
	return builder.String()
}

// KcCourseVideos is a parsable slice of KcCourseVideo.
type KcCourseVideos []*KcCourseVideo

func (kcv KcCourseVideos) config(cfg config) {
	for _i := range kcv {
		kcv[_i].config = cfg
	}
}
