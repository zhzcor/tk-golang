// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gserver/internal/store/ent/majordetail"
	"gserver/internal/store/ent/majordetailtag"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// MajorDetailTag is the model entity for the MajorDetailTag schema.
type MajorDetailTag struct {
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
	// 标签名称
	Name string `json:"name"`
	// MajorDetailID holds the value of the "major_detail_id" field.
	// 专业详情id
	MajorDetailID int `json:"major_detail_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MajorDetailTagQuery when eager-loading is set.
	Edges MajorDetailTagEdges `json:"edges"`
}

// MajorDetailTagEdges holds the relations/edges for other nodes in the graph.
type MajorDetailTagEdges struct {
	// MajorDetail holds the value of the major_detail edge.
	MajorDetail *MajorDetail `json:"major_detail,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MajorDetailOrErr returns the MajorDetail value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MajorDetailTagEdges) MajorDetailOrErr() (*MajorDetail, error) {
	if e.loadedTypes[0] {
		if e.MajorDetail == nil {
			// The edge major_detail was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: majordetail.Label}
		}
		return e.MajorDetail, nil
	}
	return nil, &NotLoadedError{edge: "major_detail"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MajorDetailTag) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case majordetailtag.FieldID, majordetailtag.FieldMajorDetailID:
			values[i] = new(sql.NullInt64)
		case majordetailtag.FieldUUID, majordetailtag.FieldName:
			values[i] = new(sql.NullString)
		case majordetailtag.FieldCreatedAt, majordetailtag.FieldUpdatedAt, majordetailtag.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MajorDetailTag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MajorDetailTag fields.
func (mdt *MajorDetailTag) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case majordetailtag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mdt.ID = int(value.Int64)
		case majordetailtag.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				mdt.UUID = value.String
			}
		case majordetailtag.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mdt.CreatedAt = new(time.Time)
				*mdt.CreatedAt = value.Time
			}
		case majordetailtag.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mdt.UpdatedAt = new(time.Time)
				*mdt.UpdatedAt = value.Time
			}
		case majordetailtag.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mdt.DeletedAt = new(time.Time)
				*mdt.DeletedAt = value.Time
			}
		case majordetailtag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mdt.Name = value.String
			}
		case majordetailtag.FieldMajorDetailID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field major_detail_id", values[i])
			} else if value.Valid {
				mdt.MajorDetailID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMajorDetail queries the "major_detail" edge of the MajorDetailTag entity.
func (mdt *MajorDetailTag) QueryMajorDetail() *MajorDetailQuery {
	return (&MajorDetailTagClient{config: mdt.config}).QueryMajorDetail(mdt)
}

// Update returns a builder for updating this MajorDetailTag.
// Note that you need to call MajorDetailTag.Unwrap() before calling this method if this MajorDetailTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (mdt *MajorDetailTag) Update() *MajorDetailTagUpdateOne {
	return (&MajorDetailTagClient{config: mdt.config}).UpdateOne(mdt)
}

// Unwrap unwraps the MajorDetailTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mdt *MajorDetailTag) Unwrap() *MajorDetailTag {
	tx, ok := mdt.config.driver.(*txDriver)
	if !ok {
		panic("ent: MajorDetailTag is not a transactional entity")
	}
	mdt.config.driver = tx.drv
	return mdt
}

// String implements the fmt.Stringer.
func (mdt *MajorDetailTag) String() string {
	var builder strings.Builder
	builder.WriteString("MajorDetailTag(")
	builder.WriteString(fmt.Sprintf("id=%v", mdt.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(mdt.UUID)
	if v := mdt.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := mdt.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := mdt.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(mdt.Name)
	builder.WriteString(", major_detail_id=")
	builder.WriteString(fmt.Sprintf("%v", mdt.MajorDetailID))
	builder.WriteByte(')')
	return builder.String()
}

// MajorDetailTags is a parsable slice of MajorDetailTag.
type MajorDetailTags []*MajorDetailTag

func (mdt MajorDetailTags) config(cfg config) {
	for _i := range mdt {
		mdt[_i].config = cfg
	}
}
