// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"tkserver/internal/store/ent/kcclass"
	"tkserver/internal/store/ent/kcuserclass"
	"tkserver/internal/store/ent/user"

	"entgo.io/ent/dialect/sql"
)

// KcUserClass is the model entity for the KcUserClass schema.
type KcUserClass struct {
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
	// PeriodType holds the value of the "period_type" field.
	// 1:长期有效，2： 固定周期
	PeriodType uint8 `json:"period_type"`
	// ClosingDate holds the value of the "closing_date" field.
	// 截止日期
	ClosingDate time.Time `json:"closing_date"`
	// StudyRate holds the value of the "study_rate" field.
	// 学习进度
	StudyRate float64 `json:"study_rate"`
	// Remark holds the value of the "remark" field.
	// 备注
	Remark string `json:"remark"`
	// Price holds the value of the "price" field.
	// 购买价格
	Price float64 `json:"price"`
	// UserID holds the value of the "user_id" field.
	// 用户id
	UserID int `json:"user_id"`
	// ClassID holds the value of the "class_id" field.
	// 班级id
	ClassID int `json:"class_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the KcUserClassQuery when eager-loading is set.
	Edges KcUserClassEdges `json:"edges"`
}

// KcUserClassEdges holds the relations/edges for other nodes in the graph.
type KcUserClassEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Class holds the value of the class edge.
	Class *KcClass `json:"class,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KcUserClassEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KcUserClassEdges) ClassOrErr() (*KcClass, error) {
	if e.loadedTypes[1] {
		if e.Class == nil {
			// The edge class was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: kcclass.Label}
		}
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KcUserClass) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case kcuserclass.FieldStudyRate, kcuserclass.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case kcuserclass.FieldID, kcuserclass.FieldPeriodType, kcuserclass.FieldUserID, kcuserclass.FieldClassID:
			values[i] = new(sql.NullInt64)
		case kcuserclass.FieldUUID, kcuserclass.FieldRemark:
			values[i] = new(sql.NullString)
		case kcuserclass.FieldCreatedAt, kcuserclass.FieldUpdatedAt, kcuserclass.FieldDeletedAt, kcuserclass.FieldClosingDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type KcUserClass", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KcUserClass fields.
func (kuc *KcUserClass) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case kcuserclass.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			kuc.ID = int(value.Int64)
		case kcuserclass.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				kuc.UUID = value.String
			}
		case kcuserclass.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				kuc.CreatedAt = new(time.Time)
				*kuc.CreatedAt = value.Time
			}
		case kcuserclass.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				kuc.UpdatedAt = new(time.Time)
				*kuc.UpdatedAt = value.Time
			}
		case kcuserclass.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				kuc.DeletedAt = new(time.Time)
				*kuc.DeletedAt = value.Time
			}
		case kcuserclass.FieldPeriodType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field period_type", values[i])
			} else if value.Valid {
				kuc.PeriodType = uint8(value.Int64)
			}
		case kcuserclass.FieldClosingDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field closing_date", values[i])
			} else if value.Valid {
				kuc.ClosingDate = value.Time
			}
		case kcuserclass.FieldStudyRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field study_rate", values[i])
			} else if value.Valid {
				kuc.StudyRate = value.Float64
			}
		case kcuserclass.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				kuc.Remark = value.String
			}
		case kcuserclass.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				kuc.Price = value.Float64
			}
		case kcuserclass.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				kuc.UserID = int(value.Int64)
			}
		case kcuserclass.FieldClassID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field class_id", values[i])
			} else if value.Valid {
				kuc.ClassID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the KcUserClass entity.
func (kuc *KcUserClass) QueryUser() *UserQuery {
	return (&KcUserClassClient{config: kuc.config}).QueryUser(kuc)
}

// QueryClass queries the "class" edge of the KcUserClass entity.
func (kuc *KcUserClass) QueryClass() *KcClassQuery {
	return (&KcUserClassClient{config: kuc.config}).QueryClass(kuc)
}

// Update returns a builder for updating this KcUserClass.
// Note that you need to call KcUserClass.Unwrap() before calling this method if this KcUserClass
// was returned from a transaction, and the transaction was committed or rolled back.
func (kuc *KcUserClass) Update() *KcUserClassUpdateOne {
	return (&KcUserClassClient{config: kuc.config}).UpdateOne(kuc)
}

// Unwrap unwraps the KcUserClass entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kuc *KcUserClass) Unwrap() *KcUserClass {
	tx, ok := kuc.config.driver.(*txDriver)
	if !ok {
		panic("ent: KcUserClass is not a transactional entity")
	}
	kuc.config.driver = tx.drv
	return kuc
}

// String implements the fmt.Stringer.
func (kuc *KcUserClass) String() string {
	var builder strings.Builder
	builder.WriteString("KcUserClass(")
	builder.WriteString(fmt.Sprintf("id=%v", kuc.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(kuc.UUID)
	if v := kuc.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := kuc.UpdatedAt; v != nil {
		builder.WriteString(", updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := kuc.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", period_type=")
	builder.WriteString(fmt.Sprintf("%v", kuc.PeriodType))
	builder.WriteString(", closing_date=")
	builder.WriteString(kuc.ClosingDate.Format(time.ANSIC))
	builder.WriteString(", study_rate=")
	builder.WriteString(fmt.Sprintf("%v", kuc.StudyRate))
	builder.WriteString(", remark=")
	builder.WriteString(kuc.Remark)
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", kuc.Price))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", kuc.UserID))
	builder.WriteString(", class_id=")
	builder.WriteString(fmt.Sprintf("%v", kuc.ClassID))
	builder.WriteByte(')')
	return builder.String()
}

// KcUserClasses is a parsable slice of KcUserClass.
type KcUserClasses []*KcUserClass

func (kuc KcUserClasses) config(cfg config) {
	for _i := range kuc {
		kuc[_i].config = cfg
	}
}
