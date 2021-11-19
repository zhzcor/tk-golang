// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gserver/internal/store/ent/attachment"
	"gserver/internal/store/ent/teacher"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Teacher is the model entity for the Teacher schema.
type Teacher struct {
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
	// Name holds the value of the "name" field.
	// 教师名称
	Name string `json:"name"`
	// Sex holds the value of the "sex" field.
	// 性别，1：男，2：女，3：未知
	Sex uint8 `json:"sex"`
	// Email holds the value of the "email" field.
	// 邮箱
	Email string `json:"email"`
	// Phone holds the value of the "phone" field.
	// 手机号
	Phone string `json:"phone"`
	// Nickname holds the value of the "nickname" field.
	// 昵称
	Nickname string `json:"nickname"`
	// SubTitle holds the value of the "sub_title" field.
	// 副标题
	SubTitle string `json:"sub_title"`
	// Detail holds the value of the "detail" field.
	// 详情
	Detail string `json:"detail"`
	// Status holds the value of the "status" field.
	// 是否启用  1：启用，2：未启用
	Status uint8 `json:"status"`
	// TeachingAge holds the value of the "teaching_age" field.
	// 教龄
	TeachingAge uint8 `json:"teaching_age"`
	// AvatarID holds the value of the "avatar_id" field.
	// 教师头像id
	AvatarID int `json:"avatar_id"`
	// SortOrder holds the value of the "sort_order" field.
	// 排序
	SortOrder int `json:"sort_order"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeacherQuery when eager-loading is set.
	Edges TeacherEdges `json:"edges"`
}

// TeacherEdges holds the relations/edges for other nodes in the graph.
type TeacherEdges struct {
	// Majors holds the value of the majors edge.
	Majors []*Major `json:"majors,omitempty"`
	// TeacherTags holds the value of the teacher_tags edge.
	TeacherTags []*TeacherTag `json:"teacher_tags,omitempty"`
	// TeacherClasses holds the value of the teacher_classes edge.
	TeacherClasses []*KcClassTeacher `json:"teacher_classes,omitempty"`
	// KcClassMasters holds the value of the kc_class_masters edge.
	KcClassMasters []*KcClass `json:"kc_class_masters,omitempty"`
	// TeacherCourses holds the value of the teacher_courses edge.
	TeacherCourses []*KcCourseTeacher `json:"teacher_courses,omitempty"`
	// UserExamsTeachers holds the value of the user_exams_teachers edge.
	UserExamsTeachers []*TkUserExamScoreRecord `json:"user_exams_teachers,omitempty"`
	// AskTeachers holds the value of the ask_teachers edge.
	AskTeachers []*UserAskAnswer `json:"ask_teachers,omitempty"`
	// Attachment holds the value of the attachment edge.
	Attachment *Attachment `json:"attachment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// MajorsOrErr returns the Majors value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) MajorsOrErr() ([]*Major, error) {
	if e.loadedTypes[0] {
		return e.Majors, nil
	}
	return nil, &NotLoadedError{edge: "majors"}
}

// TeacherTagsOrErr returns the TeacherTags value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) TeacherTagsOrErr() ([]*TeacherTag, error) {
	if e.loadedTypes[1] {
		return e.TeacherTags, nil
	}
	return nil, &NotLoadedError{edge: "teacher_tags"}
}

// TeacherClassesOrErr returns the TeacherClasses value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) TeacherClassesOrErr() ([]*KcClassTeacher, error) {
	if e.loadedTypes[2] {
		return e.TeacherClasses, nil
	}
	return nil, &NotLoadedError{edge: "teacher_classes"}
}

// KcClassMastersOrErr returns the KcClassMasters value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) KcClassMastersOrErr() ([]*KcClass, error) {
	if e.loadedTypes[3] {
		return e.KcClassMasters, nil
	}
	return nil, &NotLoadedError{edge: "kc_class_masters"}
}

// TeacherCoursesOrErr returns the TeacherCourses value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) TeacherCoursesOrErr() ([]*KcCourseTeacher, error) {
	if e.loadedTypes[4] {
		return e.TeacherCourses, nil
	}
	return nil, &NotLoadedError{edge: "teacher_courses"}
}

// UserExamsTeachersOrErr returns the UserExamsTeachers value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) UserExamsTeachersOrErr() ([]*TkUserExamScoreRecord, error) {
	if e.loadedTypes[5] {
		return e.UserExamsTeachers, nil
	}
	return nil, &NotLoadedError{edge: "user_exams_teachers"}
}

// AskTeachersOrErr returns the AskTeachers value or an error if the edge
// was not loaded in eager-loading.
func (e TeacherEdges) AskTeachersOrErr() ([]*UserAskAnswer, error) {
	if e.loadedTypes[6] {
		return e.AskTeachers, nil
	}
	return nil, &NotLoadedError{edge: "ask_teachers"}
}

// AttachmentOrErr returns the Attachment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeacherEdges) AttachmentOrErr() (*Attachment, error) {
	if e.loadedTypes[7] {
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
func (*Teacher) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case teacher.FieldID, teacher.FieldSex, teacher.FieldStatus, teacher.FieldTeachingAge, teacher.FieldAvatarID, teacher.FieldSortOrder:
			values[i] = new(sql.NullInt64)
		case teacher.FieldUUID, teacher.FieldName, teacher.FieldEmail, teacher.FieldPhone, teacher.FieldNickname, teacher.FieldSubTitle, teacher.FieldDetail:
			values[i] = new(sql.NullString)
		case teacher.FieldCreatedAt, teacher.FieldUpdatedAt, teacher.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Teacher", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Teacher fields.
func (t *Teacher) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case teacher.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case teacher.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				t.UUID = value.String
			}
		case teacher.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = new(time.Time)
				*t.CreatedAt = value.Time
			}
		case teacher.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = new(time.Time)
				*t.UpdatedAt = value.Time
			}
		case teacher.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = new(time.Time)
				*t.DeletedAt = value.Time
			}
		case teacher.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case teacher.FieldSex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sex", values[i])
			} else if value.Valid {
				t.Sex = uint8(value.Int64)
			}
		case teacher.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				t.Email = value.String
			}
		case teacher.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				t.Phone = value.String
			}
		case teacher.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				t.Nickname = value.String
			}
		case teacher.FieldSubTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sub_title", values[i])
			} else if value.Valid {
				t.SubTitle = value.String
			}
		case teacher.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				t.Detail = value.String
			}
		case teacher.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = uint8(value.Int64)
			}
		case teacher.FieldTeachingAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field teaching_age", values[i])
			} else if value.Valid {
				t.TeachingAge = uint8(value.Int64)
			}
		case teacher.FieldAvatarID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_id", values[i])
			} else if value.Valid {
				t.AvatarID = int(value.Int64)
			}
		case teacher.FieldSortOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort_order", values[i])
			} else if value.Valid {
				t.SortOrder = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMajors queries the "majors" edge of the Teacher entity.
func (t *Teacher) QueryMajors() *MajorQuery {
	return (&TeacherClient{config: t.config}).QueryMajors(t)
}

// QueryTeacherTags queries the "teacher_tags" edge of the Teacher entity.
func (t *Teacher) QueryTeacherTags() *TeacherTagQuery {
	return (&TeacherClient{config: t.config}).QueryTeacherTags(t)
}

// QueryTeacherClasses queries the "teacher_classes" edge of the Teacher entity.
func (t *Teacher) QueryTeacherClasses() *KcClassTeacherQuery {
	return (&TeacherClient{config: t.config}).QueryTeacherClasses(t)
}

// QueryKcClassMasters queries the "kc_class_masters" edge of the Teacher entity.
func (t *Teacher) QueryKcClassMasters() *KcClassQuery {
	return (&TeacherClient{config: t.config}).QueryKcClassMasters(t)
}

// QueryTeacherCourses queries the "teacher_courses" edge of the Teacher entity.
func (t *Teacher) QueryTeacherCourses() *KcCourseTeacherQuery {
	return (&TeacherClient{config: t.config}).QueryTeacherCourses(t)
}

// QueryUserExamsTeachers queries the "user_exams_teachers" edge of the Teacher entity.
func (t *Teacher) QueryUserExamsTeachers() *TkUserExamScoreRecordQuery {
	return (&TeacherClient{config: t.config}).QueryUserExamsTeachers(t)
}

// QueryAskTeachers queries the "ask_teachers" edge of the Teacher entity.
func (t *Teacher) QueryAskTeachers() *UserAskAnswerQuery {
	return (&TeacherClient{config: t.config}).QueryAskTeachers(t)
}

// QueryAttachment queries the "attachment" edge of the Teacher entity.
func (t *Teacher) QueryAttachment() *AttachmentQuery {
	return (&TeacherClient{config: t.config}).QueryAttachment(t)
}

// Update returns a builder for updating this Teacher.
// Note that you need to call Teacher.Unwrap() before calling this method if this Teacher
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Teacher) Update() *TeacherUpdateOne {
	return (&TeacherClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Teacher entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Teacher) Unwrap() *Teacher {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Teacher is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Teacher) String() string {
	var builder strings.Builder
	builder.WriteString("Teacher(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(t.UUID)
	if v := t.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := t.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := t.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", sex=")
	builder.WriteString(fmt.Sprintf("%v", t.Sex))
	builder.WriteString(", email=")
	builder.WriteString(t.Email)
	builder.WriteString(", phone=")
	builder.WriteString(t.Phone)
	builder.WriteString(", nickname=")
	builder.WriteString(t.Nickname)
	builder.WriteString(", sub_title=")
	builder.WriteString(t.SubTitle)
	builder.WriteString(", detail=")
	builder.WriteString(t.Detail)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", teaching_age=")
	builder.WriteString(fmt.Sprintf("%v", t.TeachingAge))
	builder.WriteString(", avatar_id=")
	builder.WriteString(fmt.Sprintf("%v", t.AvatarID))
	builder.WriteString(", sort_order=")
	builder.WriteString(fmt.Sprintf("%v", t.SortOrder))
	builder.WriteByte(')')
	return builder.String()
}

// Teachers is a parsable slice of Teacher.
type Teachers []*Teacher

func (t Teachers) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
