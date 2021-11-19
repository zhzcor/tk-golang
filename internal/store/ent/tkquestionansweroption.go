// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gserver/internal/store/ent/tkquestion"
	"gserver/internal/store/ent/tkquestionansweroption"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TkQuestionAnswerOption is the model entity for the TkQuestionAnswerOption schema.
type TkQuestionAnswerOption struct {
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
	// OptionName holds the value of the "option_name" field.
	// 选项名称：A,B,C,D等
	OptionName string `json:"option_name"`
	// Content holds the value of the "content" field.
	// 选项内容
	Content string `json:"content"`
	// IsRight holds the value of the "is_right" field.
	// 是否正确 1：否，2：是
	IsRight uint8 `json:"is_right"`
	// QuestionID holds the value of the "question_id" field.
	// 题目id
	QuestionID int `json:"question_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TkQuestionAnswerOptionQuery when eager-loading is set.
	Edges TkQuestionAnswerOptionEdges `json:"edges"`
}

// TkQuestionAnswerOptionEdges holds the relations/edges for other nodes in the graph.
type TkQuestionAnswerOptionEdges struct {
	// Question holds the value of the question edge.
	Question *TkQuestion `json:"question,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// QuestionOrErr returns the Question value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TkQuestionAnswerOptionEdges) QuestionOrErr() (*TkQuestion, error) {
	if e.loadedTypes[0] {
		if e.Question == nil {
			// The edge question was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: tkquestion.Label}
		}
		return e.Question, nil
	}
	return nil, &NotLoadedError{edge: "question"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TkQuestionAnswerOption) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tkquestionansweroption.FieldID, tkquestionansweroption.FieldIsRight, tkquestionansweroption.FieldQuestionID:
			values[i] = new(sql.NullInt64)
		case tkquestionansweroption.FieldUUID, tkquestionansweroption.FieldOptionName, tkquestionansweroption.FieldContent:
			values[i] = new(sql.NullString)
		case tkquestionansweroption.FieldCreatedAt, tkquestionansweroption.FieldUpdatedAt, tkquestionansweroption.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TkQuestionAnswerOption", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TkQuestionAnswerOption fields.
func (tqao *TkQuestionAnswerOption) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tkquestionansweroption.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tqao.ID = int(value.Int64)
		case tkquestionansweroption.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				tqao.UUID = value.String
			}
		case tkquestionansweroption.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tqao.CreatedAt = new(time.Time)
				*tqao.CreatedAt = value.Time
			}
		case tkquestionansweroption.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tqao.UpdatedAt = new(time.Time)
				*tqao.UpdatedAt = value.Time
			}
		case tkquestionansweroption.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				tqao.DeletedAt = new(time.Time)
				*tqao.DeletedAt = value.Time
			}
		case tkquestionansweroption.FieldOptionName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field option_name", values[i])
			} else if value.Valid {
				tqao.OptionName = value.String
			}
		case tkquestionansweroption.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				tqao.Content = value.String
			}
		case tkquestionansweroption.FieldIsRight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_right", values[i])
			} else if value.Valid {
				tqao.IsRight = uint8(value.Int64)
			}
		case tkquestionansweroption.FieldQuestionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field question_id", values[i])
			} else if value.Valid {
				tqao.QuestionID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryQuestion queries the "question" edge of the TkQuestionAnswerOption entity.
func (tqao *TkQuestionAnswerOption) QueryQuestion() *TkQuestionQuery {
	return (&TkQuestionAnswerOptionClient{config: tqao.config}).QueryQuestion(tqao)
}

// Update returns a builder for updating this TkQuestionAnswerOption.
// Note that you need to call TkQuestionAnswerOption.Unwrap() before calling this method if this TkQuestionAnswerOption
// was returned from a transaction, and the transaction was committed or rolled back.
func (tqao *TkQuestionAnswerOption) Update() *TkQuestionAnswerOptionUpdateOne {
	return (&TkQuestionAnswerOptionClient{config: tqao.config}).UpdateOne(tqao)
}

// Unwrap unwraps the TkQuestionAnswerOption entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tqao *TkQuestionAnswerOption) Unwrap() *TkQuestionAnswerOption {
	tx, ok := tqao.config.driver.(*txDriver)
	if !ok {
		panic("ent: TkQuestionAnswerOption is not a transactional entity")
	}
	tqao.config.driver = tx.drv
	return tqao
}

// String implements the fmt.Stringer.
func (tqao *TkQuestionAnswerOption) String() string {
	var builder strings.Builder
	builder.WriteString("TkQuestionAnswerOption(")
	builder.WriteString(fmt.Sprintf("id=%v", tqao.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(tqao.UUID)
	if v := tqao.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := tqao.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := tqao.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", option_name=")
	builder.WriteString(tqao.OptionName)
	builder.WriteString(", content=")
	builder.WriteString(tqao.Content)
	builder.WriteString(", is_right=")
	builder.WriteString(fmt.Sprintf("%v", tqao.IsRight))
	builder.WriteString(", question_id=")
	builder.WriteString(fmt.Sprintf("%v", tqao.QuestionID))
	builder.WriteByte(')')
	return builder.String()
}

// TkQuestionAnswerOptions is a parsable slice of TkQuestionAnswerOption.
type TkQuestionAnswerOptions []*TkQuestionAnswerOption

func (tqao TkQuestionAnswerOptions) config(cfg config) {
	for _i := range tqao {
		tqao[_i].config = cfg
	}
}
