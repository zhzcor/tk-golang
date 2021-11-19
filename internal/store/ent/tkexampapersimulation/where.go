// Code generated by entc, DO NOT EDIT.

package tkexampapersimulation

import (
	"gserver/internal/store/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// EnableStatus applies equality check predicate on the "enable_status" field. It's identical to EnableStatusEQ.
func EnableStatus(v uint8) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnableStatus), v))
	})
}

// ExamPaperID applies equality check predicate on the "exam_paper_id" field. It's identical to ExamPaperIDEQ.
func ExamPaperID(v int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExamPaperID), v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUUID), v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUUID), v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUUID), v))
	})
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUUID), v))
	})
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUUID), v))
	})
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUUID), v))
	})
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUUID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedAt)))
	})
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedAt)))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAt), v))
	})
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...time.Time) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartAt), v...))
	})
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...time.Time) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartAt), v...))
	})
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAt), v))
	})
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAt), v))
	})
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAt), v))
	})
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAt), v))
	})
}

// StartAtIsNil applies the IsNil predicate on the "start_at" field.
func StartAtIsNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartAt)))
	})
}

// StartAtNotNil applies the NotNil predicate on the "start_at" field.
func StartAtNotNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartAt)))
	})
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndAt), v))
	})
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndAt), v...))
	})
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndAt), v...))
	})
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndAt), v))
	})
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndAt), v))
	})
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndAt), v))
	})
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndAt), v))
	})
}

// EndAtIsNil applies the IsNil predicate on the "end_at" field.
func EndAtIsNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndAt)))
	})
}

// EndAtNotNil applies the NotNil predicate on the "end_at" field.
func EndAtNotNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndAt)))
	})
}

// EnableStatusEQ applies the EQ predicate on the "enable_status" field.
func EnableStatusEQ(v uint8) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnableStatus), v))
	})
}

// EnableStatusNEQ applies the NEQ predicate on the "enable_status" field.
func EnableStatusNEQ(v uint8) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnableStatus), v))
	})
}

// EnableStatusIn applies the In predicate on the "enable_status" field.
func EnableStatusIn(vs ...uint8) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEnableStatus), v...))
	})
}

// EnableStatusNotIn applies the NotIn predicate on the "enable_status" field.
func EnableStatusNotIn(vs ...uint8) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEnableStatus), v...))
	})
}

// EnableStatusGT applies the GT predicate on the "enable_status" field.
func EnableStatusGT(v uint8) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnableStatus), v))
	})
}

// EnableStatusGTE applies the GTE predicate on the "enable_status" field.
func EnableStatusGTE(v uint8) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnableStatus), v))
	})
}

// EnableStatusLT applies the LT predicate on the "enable_status" field.
func EnableStatusLT(v uint8) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnableStatus), v))
	})
}

// EnableStatusLTE applies the LTE predicate on the "enable_status" field.
func EnableStatusLTE(v uint8) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnableStatus), v))
	})
}

// ExamPaperIDEQ applies the EQ predicate on the "exam_paper_id" field.
func ExamPaperIDEQ(v int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExamPaperID), v))
	})
}

// ExamPaperIDNEQ applies the NEQ predicate on the "exam_paper_id" field.
func ExamPaperIDNEQ(v int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExamPaperID), v))
	})
}

// ExamPaperIDIn applies the In predicate on the "exam_paper_id" field.
func ExamPaperIDIn(vs ...int) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExamPaperID), v...))
	})
}

// ExamPaperIDNotIn applies the NotIn predicate on the "exam_paper_id" field.
func ExamPaperIDNotIn(vs ...int) predicate.TkExamPaperSimulation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExamPaperID), v...))
	})
}

// ExamPaperIDGT applies the GT predicate on the "exam_paper_id" field.
func ExamPaperIDGT(v int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExamPaperID), v))
	})
}

// ExamPaperIDGTE applies the GTE predicate on the "exam_paper_id" field.
func ExamPaperIDGTE(v int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExamPaperID), v))
	})
}

// ExamPaperIDLT applies the LT predicate on the "exam_paper_id" field.
func ExamPaperIDLT(v int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExamPaperID), v))
	})
}

// ExamPaperIDLTE applies the LTE predicate on the "exam_paper_id" field.
func ExamPaperIDLTE(v int) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExamPaperID), v))
	})
}

// ExamPaperIDIsNil applies the IsNil predicate on the "exam_paper_id" field.
func ExamPaperIDIsNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExamPaperID)))
	})
}

// ExamPaperIDNotNil applies the NotNil predicate on the "exam_paper_id" field.
func ExamPaperIDNotNil() predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExamPaperID)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TkExamPaperSimulation) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TkExamPaperSimulation) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TkExamPaperSimulation) predicate.TkExamPaperSimulation {
	return predicate.TkExamPaperSimulation(func(s *sql.Selector) {
		p(s.Not())
	})
}
