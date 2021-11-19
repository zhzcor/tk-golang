// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gserver/internal/store/ent/admin"
	"gserver/internal/store/ent/attachment"
	"gserver/internal/store/ent/city"
	"gserver/internal/store/ent/itemcategory"
	"gserver/internal/store/ent/kcclass"
	"gserver/internal/store/ent/teacher"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// KcClass is the model entity for the KcClass schema.
type KcClass struct {
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
	// ClassTitle holds the value of the "class_title" field.
	// 班级标题
	ClassTitle string `json:"class_title"`
	// ClassCode holds the value of the "class_code" field.
	// 班级编号
	ClassCode string `json:"class_code"`
	// ClassDesc holds the value of the "class_desc" field.
	// 班级介绍
	ClassDesc string `json:"class_desc"`
	// IsDisplay holds the value of the "is_display" field.
	// 是否展示1：展示 0：不展示
	IsDisplay uint8 `json:"is_display"`
	// IsBuy holds the value of the "is_buy" field.
	// 是否购买1：可以购买 0：不可以购买
	IsBuy uint8 `json:"is_buy"`
	// ClassPeriodType holds the value of the "class_period_type" field.
	// 1:随到随学，2： 固定周期，3：长期有效
	ClassPeriodType uint8 `json:"class_period_type"`
	// ClassStartDate holds the value of the "class_start_date" field.
	// 开始日期
	ClassStartDate time.Time `json:"class_start_date"`
	// ClassEndDate holds the value of the "class_end_date" field.
	// 结束日期
	ClassEndDate time.Time `json:"class_end_date"`
	// ClosingDate holds the value of the "closing_date" field.
	// 截止日期
	ClosingDate *time.Time `json:"closing_date"`
	// DaysValidity holds the value of the "days_validity" field.
	// 有效期天数
	DaysValidity int `json:"days_validity"`
	// ClassHeadMasterID holds the value of the "class_head_master_id" field.
	// 班主任老师id
	ClassHeadMasterID int `json:"class_head_master_id"`
	// ClassCoverImgID holds the value of the "class_cover_img_id" field.
	// 封面图片id
	ClassCoverImgID int `json:"class_cover_img_id"`
	// Status holds the value of the "status" field.
	// 状态 1：已发布 2：未发布 3：已关闭
	Status uint8 `json:"status"`
	// Price holds the value of the "price" field.
	// 价格
	Price float64 `json:"price"`
	// StudentCount holds the value of the "student_count" field.
	// 班级人数
	StudentCount int `json:"student_count"`
	// CourseCount holds the value of the "course_count" field.
	// 课程数
	CourseCount int `json:"course_count"`
	// CateID holds the value of the "cate_id" field.
	// 项目id
	CateID int `json:"cate_id"`
	// CityID holds the value of the "city_id" field.
	// 地区id
	CityID int `json:"city_id"`
	// CreatedAdminID holds the value of the "created_admin_id" field.
	// 创建人id
	CreatedAdminID int `json:"created_admin_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the KcClassQuery when eager-loading is set.
	Edges KcClassEdges `json:"edges"`
}

// KcClassEdges holds the relations/edges for other nodes in the graph.
type KcClassEdges struct {
	// Majors holds the value of the majors edge.
	Majors []*Major `json:"majors,omitempty"`
	// Item holds the value of the item edge.
	Item *ItemCategory `json:"item,omitempty"`
	// City holds the value of the city edge.
	City *City `json:"city,omitempty"`
	// Admin holds the value of the admin edge.
	Admin *Admin `json:"admin,omitempty"`
	// Attachment holds the value of the attachment edge.
	Attachment *Attachment `json:"attachment,omitempty"`
	// MasterTeachers holds the value of the master_teachers edge.
	MasterTeachers *Teacher `json:"master_teachers,omitempty"`
	// ClassTeachers holds the value of the class_teachers edge.
	ClassTeachers []*KcClassTeacher `json:"class_teachers,omitempty"`
	// KcClassCourses holds the value of the kc_class_courses edge.
	KcClassCourses []*KcCourse `json:"kc_class_courses,omitempty"`
	// KcUserClasses holds the value of the kc_user_classes edge.
	KcUserClasses []*KcUserClass `json:"kc_user_classes,omitempty"`
	// MessageClasses holds the value of the message_classes edge.
	MessageClasses []*Message `json:"message_classes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [10]bool
}

// MajorsOrErr returns the Majors value or an error if the edge
// was not loaded in eager-loading.
func (e KcClassEdges) MajorsOrErr() ([]*Major, error) {
	if e.loadedTypes[0] {
		return e.Majors, nil
	}
	return nil, &NotLoadedError{edge: "majors"}
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KcClassEdges) ItemOrErr() (*ItemCategory, error) {
	if e.loadedTypes[1] {
		if e.Item == nil {
			// The edge item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: itemcategory.Label}
		}
		return e.Item, nil
	}
	return nil, &NotLoadedError{edge: "item"}
}

// CityOrErr returns the City value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KcClassEdges) CityOrErr() (*City, error) {
	if e.loadedTypes[2] {
		if e.City == nil {
			// The edge city was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: city.Label}
		}
		return e.City, nil
	}
	return nil, &NotLoadedError{edge: "city"}
}

// AdminOrErr returns the Admin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KcClassEdges) AdminOrErr() (*Admin, error) {
	if e.loadedTypes[3] {
		if e.Admin == nil {
			// The edge admin was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: admin.Label}
		}
		return e.Admin, nil
	}
	return nil, &NotLoadedError{edge: "admin"}
}

// AttachmentOrErr returns the Attachment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KcClassEdges) AttachmentOrErr() (*Attachment, error) {
	if e.loadedTypes[4] {
		if e.Attachment == nil {
			// The edge attachment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: attachment.Label}
		}
		return e.Attachment, nil
	}
	return nil, &NotLoadedError{edge: "attachment"}
}

// MasterTeachersOrErr returns the MasterTeachers value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KcClassEdges) MasterTeachersOrErr() (*Teacher, error) {
	if e.loadedTypes[5] {
		if e.MasterTeachers == nil {
			// The edge master_teachers was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: teacher.Label}
		}
		return e.MasterTeachers, nil
	}
	return nil, &NotLoadedError{edge: "master_teachers"}
}

// ClassTeachersOrErr returns the ClassTeachers value or an error if the edge
// was not loaded in eager-loading.
func (e KcClassEdges) ClassTeachersOrErr() ([]*KcClassTeacher, error) {
	if e.loadedTypes[6] {
		return e.ClassTeachers, nil
	}
	return nil, &NotLoadedError{edge: "class_teachers"}
}

// KcClassCoursesOrErr returns the KcClassCourses value or an error if the edge
// was not loaded in eager-loading.
func (e KcClassEdges) KcClassCoursesOrErr() ([]*KcCourse, error) {
	if e.loadedTypes[7] {
		return e.KcClassCourses, nil
	}
	return nil, &NotLoadedError{edge: "kc_class_courses"}
}

// KcUserClassesOrErr returns the KcUserClasses value or an error if the edge
// was not loaded in eager-loading.
func (e KcClassEdges) KcUserClassesOrErr() ([]*KcUserClass, error) {
	if e.loadedTypes[8] {
		return e.KcUserClasses, nil
	}
	return nil, &NotLoadedError{edge: "kc_user_classes"}
}

// MessageClassesOrErr returns the MessageClasses value or an error if the edge
// was not loaded in eager-loading.
func (e KcClassEdges) MessageClassesOrErr() ([]*Message, error) {
	if e.loadedTypes[9] {
		return e.MessageClasses, nil
	}
	return nil, &NotLoadedError{edge: "message_classes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KcClass) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case kcclass.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case kcclass.FieldID, kcclass.FieldIsDisplay, kcclass.FieldIsBuy, kcclass.FieldClassPeriodType, kcclass.FieldDaysValidity, kcclass.FieldClassHeadMasterID, kcclass.FieldClassCoverImgID, kcclass.FieldStatus, kcclass.FieldStudentCount, kcclass.FieldCourseCount, kcclass.FieldCateID, kcclass.FieldCityID, kcclass.FieldCreatedAdminID:
			values[i] = new(sql.NullInt64)
		case kcclass.FieldUUID, kcclass.FieldClassTitle, kcclass.FieldClassCode, kcclass.FieldClassDesc:
			values[i] = new(sql.NullString)
		case kcclass.FieldCreatedAt, kcclass.FieldUpdatedAt, kcclass.FieldDeletedAt, kcclass.FieldClassStartDate, kcclass.FieldClassEndDate, kcclass.FieldClosingDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type KcClass", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KcClass fields.
func (kc *KcClass) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case kcclass.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			kc.ID = int(value.Int64)
		case kcclass.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				kc.UUID = value.String
			}
		case kcclass.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				kc.CreatedAt = new(time.Time)
				*kc.CreatedAt = value.Time
			}
		case kcclass.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				kc.UpdatedAt = new(time.Time)
				*kc.UpdatedAt = value.Time
			}
		case kcclass.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				kc.DeletedAt = new(time.Time)
				*kc.DeletedAt = value.Time
			}
		case kcclass.FieldClassTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class_title", values[i])
			} else if value.Valid {
				kc.ClassTitle = value.String
			}
		case kcclass.FieldClassCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class_code", values[i])
			} else if value.Valid {
				kc.ClassCode = value.String
			}
		case kcclass.FieldClassDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class_desc", values[i])
			} else if value.Valid {
				kc.ClassDesc = value.String
			}
		case kcclass.FieldIsDisplay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_display", values[i])
			} else if value.Valid {
				kc.IsDisplay = uint8(value.Int64)
			}
		case kcclass.FieldIsBuy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_buy", values[i])
			} else if value.Valid {
				kc.IsBuy = uint8(value.Int64)
			}
		case kcclass.FieldClassPeriodType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field class_period_type", values[i])
			} else if value.Valid {
				kc.ClassPeriodType = uint8(value.Int64)
			}
		case kcclass.FieldClassStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field class_start_date", values[i])
			} else if value.Valid {
				kc.ClassStartDate = value.Time
			}
		case kcclass.FieldClassEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field class_end_date", values[i])
			} else if value.Valid {
				kc.ClassEndDate = value.Time
			}
		case kcclass.FieldClosingDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field closing_date", values[i])
			} else if value.Valid {
				kc.ClosingDate = new(time.Time)
				*kc.ClosingDate = value.Time
			}
		case kcclass.FieldDaysValidity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field days_validity", values[i])
			} else if value.Valid {
				kc.DaysValidity = int(value.Int64)
			}
		case kcclass.FieldClassHeadMasterID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field class_head_master_id", values[i])
			} else if value.Valid {
				kc.ClassHeadMasterID = int(value.Int64)
			}
		case kcclass.FieldClassCoverImgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field class_cover_img_id", values[i])
			} else if value.Valid {
				kc.ClassCoverImgID = int(value.Int64)
			}
		case kcclass.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				kc.Status = uint8(value.Int64)
			}
		case kcclass.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				kc.Price = value.Float64
			}
		case kcclass.FieldStudentCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field student_count", values[i])
			} else if value.Valid {
				kc.StudentCount = int(value.Int64)
			}
		case kcclass.FieldCourseCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field course_count", values[i])
			} else if value.Valid {
				kc.CourseCount = int(value.Int64)
			}
		case kcclass.FieldCateID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cate_id", values[i])
			} else if value.Valid {
				kc.CateID = int(value.Int64)
			}
		case kcclass.FieldCityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field city_id", values[i])
			} else if value.Valid {
				kc.CityID = int(value.Int64)
			}
		case kcclass.FieldCreatedAdminID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_admin_id", values[i])
			} else if value.Valid {
				kc.CreatedAdminID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMajors queries the "majors" edge of the KcClass entity.
func (kc *KcClass) QueryMajors() *MajorQuery {
	return (&KcClassClient{config: kc.config}).QueryMajors(kc)
}

// QueryItem queries the "item" edge of the KcClass entity.
func (kc *KcClass) QueryItem() *ItemCategoryQuery {
	return (&KcClassClient{config: kc.config}).QueryItem(kc)
}

// QueryCity queries the "city" edge of the KcClass entity.
func (kc *KcClass) QueryCity() *CityQuery {
	return (&KcClassClient{config: kc.config}).QueryCity(kc)
}

// QueryAdmin queries the "admin" edge of the KcClass entity.
func (kc *KcClass) QueryAdmin() *AdminQuery {
	return (&KcClassClient{config: kc.config}).QueryAdmin(kc)
}

// QueryAttachment queries the "attachment" edge of the KcClass entity.
func (kc *KcClass) QueryAttachment() *AttachmentQuery {
	return (&KcClassClient{config: kc.config}).QueryAttachment(kc)
}

// QueryMasterTeachers queries the "master_teachers" edge of the KcClass entity.
func (kc *KcClass) QueryMasterTeachers() *TeacherQuery {
	return (&KcClassClient{config: kc.config}).QueryMasterTeachers(kc)
}

// QueryClassTeachers queries the "class_teachers" edge of the KcClass entity.
func (kc *KcClass) QueryClassTeachers() *KcClassTeacherQuery {
	return (&KcClassClient{config: kc.config}).QueryClassTeachers(kc)
}

// QueryKcClassCourses queries the "kc_class_courses" edge of the KcClass entity.
func (kc *KcClass) QueryKcClassCourses() *KcCourseQuery {
	return (&KcClassClient{config: kc.config}).QueryKcClassCourses(kc)
}

// QueryKcUserClasses queries the "kc_user_classes" edge of the KcClass entity.
func (kc *KcClass) QueryKcUserClasses() *KcUserClassQuery {
	return (&KcClassClient{config: kc.config}).QueryKcUserClasses(kc)
}

// QueryMessageClasses queries the "message_classes" edge of the KcClass entity.
func (kc *KcClass) QueryMessageClasses() *MessageQuery {
	return (&KcClassClient{config: kc.config}).QueryMessageClasses(kc)
}

// Update returns a builder for updating this KcClass.
// Note that you need to call KcClass.Unwrap() before calling this method if this KcClass
// was returned from a transaction, and the transaction was committed or rolled back.
func (kc *KcClass) Update() *KcClassUpdateOne {
	return (&KcClassClient{config: kc.config}).UpdateOne(kc)
}

// Unwrap unwraps the KcClass entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kc *KcClass) Unwrap() *KcClass {
	tx, ok := kc.config.driver.(*txDriver)
	if !ok {
		panic("ent: KcClass is not a transactional entity")
	}
	kc.config.driver = tx.drv
	return kc
}

// String implements the fmt.Stringer.
func (kc *KcClass) String() string {
	var builder strings.Builder
	builder.WriteString("KcClass(")
	builder.WriteString(fmt.Sprintf("id=%v", kc.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(kc.UUID)
	if v := kc.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := kc.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := kc.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", class_title=")
	builder.WriteString(kc.ClassTitle)
	builder.WriteString(", class_code=")
	builder.WriteString(kc.ClassCode)
	builder.WriteString(", class_desc=")
	builder.WriteString(kc.ClassDesc)
	builder.WriteString(", is_display=")
	builder.WriteString(fmt.Sprintf("%v", kc.IsDisplay))
	builder.WriteString(", is_buy=")
	builder.WriteString(fmt.Sprintf("%v", kc.IsBuy))
	builder.WriteString(", class_period_type=")
	builder.WriteString(fmt.Sprintf("%v", kc.ClassPeriodType))
	builder.WriteString(", class_start_date=")
	builder.WriteString(kc.ClassStartDate.Format(time.ANSIC))
	builder.WriteString(", class_end_date=")
	builder.WriteString(kc.ClassEndDate.Format(time.ANSIC))
	if v := kc.ClosingDate; v != nil {
		builder.WriteString(", closing_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", days_validity=")
	builder.WriteString(fmt.Sprintf("%v", kc.DaysValidity))
	builder.WriteString(", class_head_master_id=")
	builder.WriteString(fmt.Sprintf("%v", kc.ClassHeadMasterID))
	builder.WriteString(", class_cover_img_id=")
	builder.WriteString(fmt.Sprintf("%v", kc.ClassCoverImgID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", kc.Status))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", kc.Price))
	builder.WriteString(", student_count=")
	builder.WriteString(fmt.Sprintf("%v", kc.StudentCount))
	builder.WriteString(", course_count=")
	builder.WriteString(fmt.Sprintf("%v", kc.CourseCount))
	builder.WriteString(", cate_id=")
	builder.WriteString(fmt.Sprintf("%v", kc.CateID))
	builder.WriteString(", city_id=")
	builder.WriteString(fmt.Sprintf("%v", kc.CityID))
	builder.WriteString(", created_admin_id=")
	builder.WriteString(fmt.Sprintf("%v", kc.CreatedAdminID))
	builder.WriteByte(')')
	return builder.String()
}

// KcClasses is a parsable slice of KcClass.
type KcClasses []*KcClass

func (kc KcClasses) config(cfg config) {
	for _i := range kc {
		kc[_i].config = cfg
	}
}
