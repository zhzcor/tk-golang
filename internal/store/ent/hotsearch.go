// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"tkserver/internal/store/ent/hotsearch"

	"entgo.io/ent/dialect/sql"
)

// HotSearch is the model entity for the HotSearch schema.
type HotSearch struct {
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
	// SearchCount holds the value of the "search_count" field.
	// 搜索次数
	SearchCount int `json:"search_count"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HotSearch) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case hotsearch.FieldID, hotsearch.FieldStatus, hotsearch.FieldSearchCount:
			values[i] = new(sql.NullInt64)
		case hotsearch.FieldUUID, hotsearch.FieldName:
			values[i] = new(sql.NullString)
		case hotsearch.FieldCreatedAt, hotsearch.FieldUpdatedAt, hotsearch.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type HotSearch", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HotSearch fields.
func (hs *HotSearch) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hotsearch.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			hs.ID = int(value.Int64)
		case hotsearch.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				hs.UUID = value.String
			}
		case hotsearch.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				hs.CreatedAt = new(time.Time)
				*hs.CreatedAt = value.Time
			}
		case hotsearch.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				hs.UpdatedAt = new(time.Time)
				*hs.UpdatedAt = value.Time
			}
		case hotsearch.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				hs.DeletedAt = new(time.Time)
				*hs.DeletedAt = value.Time
			}
		case hotsearch.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				hs.Name = value.String
			}
		case hotsearch.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				hs.Status = uint8(value.Int64)
			}
		case hotsearch.FieldSearchCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field search_count", values[i])
			} else if value.Valid {
				hs.SearchCount = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this HotSearch.
// Note that you need to call HotSearch.Unwrap() before calling this method if this HotSearch
// was returned from a transaction, and the transaction was committed or rolled back.
func (hs *HotSearch) Update() *HotSearchUpdateOne {
	return (&HotSearchClient{config: hs.config}).UpdateOne(hs)
}

// Unwrap unwraps the HotSearch entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hs *HotSearch) Unwrap() *HotSearch {
	tx, ok := hs.config.driver.(*txDriver)
	if !ok {
		panic("ent: HotSearch is not a transactional entity")
	}
	hs.config.driver = tx.drv
	return hs
}

// String implements the fmt.Stringer.
func (hs *HotSearch) String() string {
	var builder strings.Builder
	builder.WriteString("HotSearch(")
	builder.WriteString(fmt.Sprintf("id=%v", hs.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(hs.UUID)
	if v := hs.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := hs.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := hs.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(hs.Name)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", hs.Status))
	builder.WriteString(", search_count=")
	builder.WriteString(fmt.Sprintf("%v", hs.SearchCount))
	builder.WriteByte(')')
	return builder.String()
}

// HotSearches is a parsable slice of HotSearch.
type HotSearches []*HotSearch

func (hs HotSearches) config(cfg config) {
	for _i := range hs {
		hs[_i].config = cfg
	}
}
