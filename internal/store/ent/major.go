// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"tkserver/internal/store/ent/major"
	"tkserver/internal/store/ent/majordetail"

	"entgo.io/ent/dialect/sql"
)

// Major is the model entity for the Major schema.
type Major struct {
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
	// 名称
	Name string `json:"name"`
	// Status holds the value of the "status" field.
	// 状态：1、锁定。2:生效
	Status uint8 `json:"status"`
	// Code holds the value of the "code" field.
	// 编码
	Code string `json:"code"`
	// Desc holds the value of the "desc" field.
	// 描述
	Desc string `json:"desc"`
	// SortOrder holds the value of the "sort_order" field.
	// 排序
	SortOrder int `json:"sort_order"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MajorQuery when eager-loading is set.
	Edges MajorEdges `json:"edges"`
}

// MajorEdges holds the relations/edges for other nodes in the graph.
type MajorEdges struct {
	// Teachers holds the value of the teachers edge.
	Teachers []*Teacher `json:"teachers,omitempty"`
	// MajorDetail holds the value of the major_detail edge.
	MajorDetail *MajorDetail `json:"major_detail,omitempty"`
	// KcClasses holds the value of the kc_classes edge.
	KcClasses []*KcClass `json:"kc_classes,omitempty"`
	// Courses holds the value of the courses edge.
	Courses []*KcCourse `json:"courses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TeachersOrErr returns the Teachers value or an error if the edge
// was not loaded in eager-loading.
func (e MajorEdges) TeachersOrErr() ([]*Teacher, error) {
	if e.loadedTypes[0] {
		return e.Teachers, nil
	}
	return nil, &NotLoadedError{edge: "teachers"}
}

// MajorDetailOrErr returns the MajorDetail value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MajorEdges) MajorDetailOrErr() (*MajorDetail, error) {
	if e.loadedTypes[1] {
		if e.MajorDetail == nil {
			// The edge major_detail was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: majordetail.Label}
		}
		return e.MajorDetail, nil
	}
	return nil, &NotLoadedError{edge: "major_detail"}
}

// KcClassesOrErr returns the KcClasses value or an error if the edge
// was not loaded in eager-loading.
func (e MajorEdges) KcClassesOrErr() ([]*KcClass, error) {
	if e.loadedTypes[2] {
		return e.KcClasses, nil
	}
	return nil, &NotLoadedError{edge: "kc_classes"}
}

// CoursesOrErr returns the Courses value or an error if the edge
// was not loaded in eager-loading.
func (e MajorEdges) CoursesOrErr() ([]*KcCourse, error) {
	if e.loadedTypes[3] {
		return e.Courses, nil
	}
	return nil, &NotLoadedError{edge: "courses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Major) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case major.FieldID, major.FieldStatus, major.FieldSortOrder:
			values[i] = new(sql.NullInt64)
		case major.FieldUUID, major.FieldName, major.FieldCode, major.FieldDesc:
			values[i] = new(sql.NullString)
		case major.FieldCreatedAt, major.FieldUpdatedAt, major.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Major", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Major fields.
func (m *Major) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case major.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case major.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				m.UUID = value.String
			}
		case major.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = new(time.Time)
				*m.CreatedAt = value.Time
			}
		case major.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = new(time.Time)
				*m.UpdatedAt = value.Time
			}
		case major.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				m.DeletedAt = new(time.Time)
				*m.DeletedAt = value.Time
			}
		case major.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case major.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = uint8(value.Int64)
			}
		case major.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				m.Code = value.String
			}
		case major.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				m.Desc = value.String
			}
		case major.FieldSortOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort_order", values[i])
			} else if value.Valid {
				m.SortOrder = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTeachers queries the "teachers" edge of the Major entity.
func (m *Major) QueryTeachers() *TeacherQuery {
	return (&MajorClient{config: m.config}).QueryTeachers(m)
}

// QueryMajorDetail queries the "major_detail" edge of the Major entity.
func (m *Major) QueryMajorDetail() *MajorDetailQuery {
	return (&MajorClient{config: m.config}).QueryMajorDetail(m)
}

// QueryKcClasses queries the "kc_classes" edge of the Major entity.
func (m *Major) QueryKcClasses() *KcClassQuery {
	return (&MajorClient{config: m.config}).QueryKcClasses(m)
}

// QueryCourses queries the "courses" edge of the Major entity.
func (m *Major) QueryCourses() *KcCourseQuery {
	return (&MajorClient{config: m.config}).QueryCourses(m)
}

// Update returns a builder for updating this Major.
// Note that you need to call Major.Unwrap() before calling this method if this Major
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Major) Update() *MajorUpdateOne {
	return (&MajorClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Major entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Major) Unwrap() *Major {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Major is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Major) String() string {
	var builder strings.Builder
	builder.WriteString("Major(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(m.UUID)
	if v := m.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := m.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := m.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(m.Name)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", code=")
	builder.WriteString(m.Code)
	builder.WriteString(", desc=")
	builder.WriteString(m.Desc)
	builder.WriteString(", sort_order=")
	builder.WriteString(fmt.Sprintf("%v", m.SortOrder))
	builder.WriteByte(')')
	return builder.String()
}

// Majors is a parsable slice of Major.
type Majors []*Major

func (m Majors) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
