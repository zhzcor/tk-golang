// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gserver/internal/store/ent/activitytype"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ActivityType is the model entity for the ActivityType schema.
type ActivityType struct {
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
	// 活动类型名称
	Name string `json:"name"`
	// Status holds the value of the "status" field.
	// 状态：1、锁定 2:启用
	Status uint8 `json:"status"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActivityTypeQuery when eager-loading is set.
	Edges ActivityTypeEdges `json:"edges"`
}

// ActivityTypeEdges holds the relations/edges for other nodes in the graph.
type ActivityTypeEdges struct {
	// Activities holds the value of the activities edge.
	Activities []*Activity `json:"activities,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ActivitiesOrErr returns the Activities value or an error if the edge
// was not loaded in eager-loading.
func (e ActivityTypeEdges) ActivitiesOrErr() ([]*Activity, error) {
	if e.loadedTypes[0] {
		return e.Activities, nil
	}
	return nil, &NotLoadedError{edge: "activities"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ActivityType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case activitytype.FieldID, activitytype.FieldStatus:
			values[i] = new(sql.NullInt64)
		case activitytype.FieldUUID, activitytype.FieldName:
			values[i] = new(sql.NullString)
		case activitytype.FieldCreatedAt, activitytype.FieldUpdatedAt, activitytype.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ActivityType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ActivityType fields.
func (at *ActivityType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case activitytype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			at.ID = int(value.Int64)
		case activitytype.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				at.UUID = value.String
			}
		case activitytype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				at.CreatedAt = new(time.Time)
				*at.CreatedAt = value.Time
			}
		case activitytype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				at.UpdatedAt = new(time.Time)
				*at.UpdatedAt = value.Time
			}
		case activitytype.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				at.DeletedAt = new(time.Time)
				*at.DeletedAt = value.Time
			}
		case activitytype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				at.Name = value.String
			}
		case activitytype.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				at.Status = uint8(value.Int64)
			}
		}
	}
	return nil
}

// QueryActivities queries the "activities" edge of the ActivityType entity.
func (at *ActivityType) QueryActivities() *ActivityQuery {
	return (&ActivityTypeClient{config: at.config}).QueryActivities(at)
}

// Update returns a builder for updating this ActivityType.
// Note that you need to call ActivityType.Unwrap() before calling this method if this ActivityType
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *ActivityType) Update() *ActivityTypeUpdateOne {
	return (&ActivityTypeClient{config: at.config}).UpdateOne(at)
}

// Unwrap unwraps the ActivityType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *ActivityType) Unwrap() *ActivityType {
	tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("ent: ActivityType is not a transactional entity")
	}
	at.config.driver = tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *ActivityType) String() string {
	var builder strings.Builder
	builder.WriteString("ActivityType(")
	builder.WriteString(fmt.Sprintf("id=%v", at.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(at.UUID)
	if v := at.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := at.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := at.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(at.Name)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", at.Status))
	builder.WriteByte(')')
	return builder.String()
}

// ActivityTypes is a parsable slice of ActivityType.
type ActivityTypes []*ActivityType

func (at ActivityTypes) config(cfg config) {
	for _i := range at {
		at[_i].config = cfg
	}
}
