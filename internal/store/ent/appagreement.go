// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"gserver/internal/store/ent/appagreement"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// AppAgreement is the model entity for the AppAgreement schema.
type AppAgreement struct {
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
	// 协议名称
	Name string `json:"name"`
	// Type holds the value of the "type" field.
	// 协议类型：1：服务协议，2：隐私政策，3：ios充值协议，4：账户注销协议，5：App温馨提示，6：优惠卷使用规则，7：关于我们，8：加入我们，9：版权声明，10：联系我们，11：常见问题
	Type uint8 `json:"type"`
	// Detail holds the value of the "detail" field.
	// 协议详情
	Detail string `json:"detail"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppAgreement) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appagreement.FieldID, appagreement.FieldType:
			values[i] = new(sql.NullInt64)
		case appagreement.FieldUUID, appagreement.FieldName, appagreement.FieldDetail:
			values[i] = new(sql.NullString)
		case appagreement.FieldCreatedAt, appagreement.FieldUpdatedAt, appagreement.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppAgreement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppAgreement fields.
func (aa *AppAgreement) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appagreement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aa.ID = int(value.Int64)
		case appagreement.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				aa.UUID = value.String
			}
		case appagreement.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				aa.CreatedAt = new(time.Time)
				*aa.CreatedAt = value.Time
			}
		case appagreement.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				aa.UpdatedAt = new(time.Time)
				*aa.UpdatedAt = value.Time
			}
		case appagreement.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				aa.DeletedAt = new(time.Time)
				*aa.DeletedAt = value.Time
			}
		case appagreement.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				aa.Name = value.String
			}
		case appagreement.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				aa.Type = uint8(value.Int64)
			}
		case appagreement.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				aa.Detail = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppAgreement.
// Note that you need to call AppAgreement.Unwrap() before calling this method if this AppAgreement
// was returned from a transaction, and the transaction was committed or rolled back.
func (aa *AppAgreement) Update() *AppAgreementUpdateOne {
	return (&AppAgreementClient{config: aa.config}).UpdateOne(aa)
}

// Unwrap unwraps the AppAgreement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aa *AppAgreement) Unwrap() *AppAgreement {
	tx, ok := aa.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppAgreement is not a transactional entity")
	}
	aa.config.driver = tx.drv
	return aa
}

// String implements the fmt.Stringer.
func (aa *AppAgreement) String() string {
	var builder strings.Builder
	builder.WriteString("AppAgreement(")
	builder.WriteString(fmt.Sprintf("id=%v", aa.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(aa.UUID)
	if v := aa.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := aa.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := aa.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(aa.Name)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", aa.Type))
	builder.WriteString(", detail=")
	builder.WriteString(aa.Detail)
	builder.WriteByte(')')
	return builder.String()
}

// AppAgreements is a parsable slice of AppAgreement.
type AppAgreements []*AppAgreement

func (aa AppAgreements) config(cfg config) {
	for _i := range aa {
		aa[_i].config = cfg
	}
}
