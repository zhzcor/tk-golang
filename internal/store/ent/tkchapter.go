// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"tkserver/internal/store/ent/tkchapter"
	"tkserver/internal/store/ent/tkquestionbank"

	"entgo.io/ent/dialect/sql"
)

// TkChapter is the model entity for the TkChapter schema.
type TkChapter struct {
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
	// 章名称
	Name string `json:"name"`
	// QuestionBankID holds the value of the "question_bank_id" field.
	// 题库id
	QuestionBankID int `json:"question_bank_id"`
	// QuestionCount holds the value of the "question_count" field.
	// 章题目数量
	QuestionCount int `json:"question_count"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TkChapterQuery when eager-loading is set.
	Edges TkChapterEdges `json:"edges"`
}

// TkChapterEdges holds the relations/edges for other nodes in the graph.
type TkChapterEdges struct {
	// QuestionBank holds the value of the question_bank edge.
	QuestionBank *TkQuestionBank `json:"question_bank,omitempty"`
	// Sections holds the value of the sections edge.
	Sections []*TkSection `json:"sections,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// QuestionBankOrErr returns the QuestionBank value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TkChapterEdges) QuestionBankOrErr() (*TkQuestionBank, error) {
	if e.loadedTypes[0] {
		if e.QuestionBank == nil {
			// The edge question_bank was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: tkquestionbank.Label}
		}
		return e.QuestionBank, nil
	}
	return nil, &NotLoadedError{edge: "question_bank"}
}

// SectionsOrErr returns the Sections value or an error if the edge
// was not loaded in eager-loading.
func (e TkChapterEdges) SectionsOrErr() ([]*TkSection, error) {
	if e.loadedTypes[1] {
		return e.Sections, nil
	}
	return nil, &NotLoadedError{edge: "sections"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TkChapter) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tkchapter.FieldID, tkchapter.FieldQuestionBankID, tkchapter.FieldQuestionCount:
			values[i] = new(sql.NullInt64)
		case tkchapter.FieldUUID, tkchapter.FieldName:
			values[i] = new(sql.NullString)
		case tkchapter.FieldCreatedAt, tkchapter.FieldUpdatedAt, tkchapter.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TkChapter", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TkChapter fields.
func (tc *TkChapter) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tkchapter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tc.ID = int(value.Int64)
		case tkchapter.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				tc.UUID = value.String
			}
		case tkchapter.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tc.CreatedAt = new(time.Time)
				*tc.CreatedAt = value.Time
			}
		case tkchapter.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tc.UpdatedAt = new(time.Time)
				*tc.UpdatedAt = value.Time
			}
		case tkchapter.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				tc.DeletedAt = new(time.Time)
				*tc.DeletedAt = value.Time
			}
		case tkchapter.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tc.Name = value.String
			}
		case tkchapter.FieldQuestionBankID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field question_bank_id", values[i])
			} else if value.Valid {
				tc.QuestionBankID = int(value.Int64)
			}
		case tkchapter.FieldQuestionCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field question_count", values[i])
			} else if value.Valid {
				tc.QuestionCount = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryQuestionBank queries the "question_bank" edge of the TkChapter entity.
func (tc *TkChapter) QueryQuestionBank() *TkQuestionBankQuery {
	return (&TkChapterClient{config: tc.config}).QueryQuestionBank(tc)
}

// QuerySections queries the "sections" edge of the TkChapter entity.
func (tc *TkChapter) QuerySections() *TkSectionQuery {
	return (&TkChapterClient{config: tc.config}).QuerySections(tc)
}

// Update returns a builder for updating this TkChapter.
// Note that you need to call TkChapter.Unwrap() before calling this method if this TkChapter
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *TkChapter) Update() *TkChapterUpdateOne {
	return (&TkChapterClient{config: tc.config}).UpdateOne(tc)
}

// Unwrap unwraps the TkChapter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tc *TkChapter) Unwrap() *TkChapter {
	tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("ent: TkChapter is not a transactional entity")
	}
	tc.config.driver = tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *TkChapter) String() string {
	var builder strings.Builder
	builder.WriteString("TkChapter(")
	builder.WriteString(fmt.Sprintf("id=%v", tc.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(tc.UUID)
	if v := tc.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := tc.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := tc.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(tc.Name)
	builder.WriteString(", question_bank_id=")
	builder.WriteString(fmt.Sprintf("%v", tc.QuestionBankID))
	builder.WriteString(", question_count=")
	builder.WriteString(fmt.Sprintf("%v", tc.QuestionCount))
	builder.WriteByte(')')
	return builder.String()
}

// TkChapters is a parsable slice of TkChapter.
type TkChapters []*TkChapter

func (tc TkChapters) config(cfg config) {
	for _i := range tc {
		tc[_i].config = cfg
	}
}
