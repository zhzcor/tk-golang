// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gserver/internal/store/ent/kccoursesmallcategory"
	"gserver/internal/store/ent/kcsmallcategoryexampaper"
	"gserver/internal/store/ent/tkexampaper"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// KcSmallCategoryExamPaper is the model entity for the KcSmallCategoryExamPaper schema.
type KcSmallCategoryExamPaper struct {
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
	// Type holds the value of the "type" field.
	// 分类，1：试卷，2：作业
	Type uint8 `json:"type"`
	// ExamPaperID holds the value of the "exam_paper_id" field.
	// 试卷id
	ExamPaperID int `json:"exam_paper_id"`
	// SmallCategoryID holds the value of the "small_category_id" field.
	// 课时id
	SmallCategoryID int `json:"small_category_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the KcSmallCategoryExamPaperQuery when eager-loading is set.
	Edges KcSmallCategoryExamPaperEdges `json:"edges"`
}

// KcSmallCategoryExamPaperEdges holds the relations/edges for other nodes in the graph.
type KcSmallCategoryExamPaperEdges struct {
	// ExamPaper holds the value of the exam_paper edge.
	ExamPaper *TkExamPaper `json:"exam_paper,omitempty"`
	// CourseSmallCategory holds the value of the course_small_category edge.
	CourseSmallCategory *KcCourseSmallCategory `json:"course_small_category,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ExamPaperOrErr returns the ExamPaper value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KcSmallCategoryExamPaperEdges) ExamPaperOrErr() (*TkExamPaper, error) {
	if e.loadedTypes[0] {
		if e.ExamPaper == nil {
			// The edge exam_paper was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: tkexampaper.Label}
		}
		return e.ExamPaper, nil
	}
	return nil, &NotLoadedError{edge: "exam_paper"}
}

// CourseSmallCategoryOrErr returns the CourseSmallCategory value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KcSmallCategoryExamPaperEdges) CourseSmallCategoryOrErr() (*KcCourseSmallCategory, error) {
	if e.loadedTypes[1] {
		if e.CourseSmallCategory == nil {
			// The edge course_small_category was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: kccoursesmallcategory.Label}
		}
		return e.CourseSmallCategory, nil
	}
	return nil, &NotLoadedError{edge: "course_small_category"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KcSmallCategoryExamPaper) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case kcsmallcategoryexampaper.FieldID, kcsmallcategoryexampaper.FieldType, kcsmallcategoryexampaper.FieldExamPaperID, kcsmallcategoryexampaper.FieldSmallCategoryID:
			values[i] = new(sql.NullInt64)
		case kcsmallcategoryexampaper.FieldUUID:
			values[i] = new(sql.NullString)
		case kcsmallcategoryexampaper.FieldCreatedAt, kcsmallcategoryexampaper.FieldUpdatedAt, kcsmallcategoryexampaper.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type KcSmallCategoryExamPaper", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KcSmallCategoryExamPaper fields.
func (kscep *KcSmallCategoryExamPaper) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case kcsmallcategoryexampaper.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			kscep.ID = int(value.Int64)
		case kcsmallcategoryexampaper.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				kscep.UUID = value.String
			}
		case kcsmallcategoryexampaper.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				kscep.CreatedAt = new(time.Time)
				*kscep.CreatedAt = value.Time
			}
		case kcsmallcategoryexampaper.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				kscep.UpdatedAt = new(time.Time)
				*kscep.UpdatedAt = value.Time
			}
		case kcsmallcategoryexampaper.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				kscep.DeletedAt = new(time.Time)
				*kscep.DeletedAt = value.Time
			}
		case kcsmallcategoryexampaper.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				kscep.Type = uint8(value.Int64)
			}
		case kcsmallcategoryexampaper.FieldExamPaperID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field exam_paper_id", values[i])
			} else if value.Valid {
				kscep.ExamPaperID = int(value.Int64)
			}
		case kcsmallcategoryexampaper.FieldSmallCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field small_category_id", values[i])
			} else if value.Valid {
				kscep.SmallCategoryID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryExamPaper queries the "exam_paper" edge of the KcSmallCategoryExamPaper entity.
func (kscep *KcSmallCategoryExamPaper) QueryExamPaper() *TkExamPaperQuery {
	return (&KcSmallCategoryExamPaperClient{config: kscep.config}).QueryExamPaper(kscep)
}

// QueryCourseSmallCategory queries the "course_small_category" edge of the KcSmallCategoryExamPaper entity.
func (kscep *KcSmallCategoryExamPaper) QueryCourseSmallCategory() *KcCourseSmallCategoryQuery {
	return (&KcSmallCategoryExamPaperClient{config: kscep.config}).QueryCourseSmallCategory(kscep)
}

// Update returns a builder for updating this KcSmallCategoryExamPaper.
// Note that you need to call KcSmallCategoryExamPaper.Unwrap() before calling this method if this KcSmallCategoryExamPaper
// was returned from a transaction, and the transaction was committed or rolled back.
func (kscep *KcSmallCategoryExamPaper) Update() *KcSmallCategoryExamPaperUpdateOne {
	return (&KcSmallCategoryExamPaperClient{config: kscep.config}).UpdateOne(kscep)
}

// Unwrap unwraps the KcSmallCategoryExamPaper entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kscep *KcSmallCategoryExamPaper) Unwrap() *KcSmallCategoryExamPaper {
	tx, ok := kscep.config.driver.(*txDriver)
	if !ok {
		panic("ent: KcSmallCategoryExamPaper is not a transactional entity")
	}
	kscep.config.driver = tx.drv
	return kscep
}

// String implements the fmt.Stringer.
func (kscep *KcSmallCategoryExamPaper) String() string {
	var builder strings.Builder
	builder.WriteString("KcSmallCategoryExamPaper(")
	builder.WriteString(fmt.Sprintf("id=%v", kscep.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(kscep.UUID)
	if v := kscep.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := kscep.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := kscep.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", kscep.Type))
	builder.WriteString(", exam_paper_id=")
	builder.WriteString(fmt.Sprintf("%v", kscep.ExamPaperID))
	builder.WriteString(", small_category_id=")
	builder.WriteString(fmt.Sprintf("%v", kscep.SmallCategoryID))
	builder.WriteByte(')')
	return builder.String()
}

// KcSmallCategoryExamPapers is a parsable slice of KcSmallCategoryExamPaper.
type KcSmallCategoryExamPapers []*KcSmallCategoryExamPaper

func (kscep KcSmallCategoryExamPapers) config(cfg config) {
	for _i := range kscep {
		kscep[_i].config = cfg
	}
}
