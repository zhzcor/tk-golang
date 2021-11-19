// Code generated by entc, DO NOT EDIT.

package tkuserwrongquestionrecode

import (
	"time"
	"tkserver/internal/store/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func IDGT(id int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// QuestionID applies equality check predicate on the "question_id" field. It's identical to QuestionIDEQ.
func QuestionID(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuestionID), v))
	})
}

// QuestionBankID applies equality check predicate on the "question_bank_id" field. It's identical to QuestionBankIDEQ.
func QuestionBankID(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuestionBankID), v))
	})
}

// Answer applies equality check predicate on the "answer" field. It's identical to AnswerEQ.
func Answer(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnswer), v))
	})
}

// WrongExamType applies equality check predicate on the "wrong_exam_type" field. It's identical to WrongExamTypeEQ.
func WrongExamType(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWrongExamType), v))
	})
}

// WrongQuestionType applies equality check predicate on the "wrong_question_type" field. It's identical to WrongQuestionTypeEQ.
func WrongQuestionType(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWrongQuestionType), v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func UUIDNotIn(vs ...string) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func UUIDGT(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUUID), v))
	})
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUUID), v))
	})
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUUID), v))
	})
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUUID), v))
	})
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUUID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedAt)))
	})
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedAt)))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...time.Time) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func DeletedAtGT(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// QuestionIDEQ applies the EQ predicate on the "question_id" field.
func QuestionIDEQ(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuestionID), v))
	})
}

// QuestionIDNEQ applies the NEQ predicate on the "question_id" field.
func QuestionIDNEQ(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuestionID), v))
	})
}

// QuestionIDIn applies the In predicate on the "question_id" field.
func QuestionIDIn(vs ...int) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQuestionID), v...))
	})
}

// QuestionIDNotIn applies the NotIn predicate on the "question_id" field.
func QuestionIDNotIn(vs ...int) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQuestionID), v...))
	})
}

// QuestionIDIsNil applies the IsNil predicate on the "question_id" field.
func QuestionIDIsNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldQuestionID)))
	})
}

// QuestionIDNotNil applies the NotNil predicate on the "question_id" field.
func QuestionIDNotNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldQuestionID)))
	})
}

// QuestionBankIDEQ applies the EQ predicate on the "question_bank_id" field.
func QuestionBankIDEQ(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuestionBankID), v))
	})
}

// QuestionBankIDNEQ applies the NEQ predicate on the "question_bank_id" field.
func QuestionBankIDNEQ(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuestionBankID), v))
	})
}

// QuestionBankIDIn applies the In predicate on the "question_bank_id" field.
func QuestionBankIDIn(vs ...int) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQuestionBankID), v...))
	})
}

// QuestionBankIDNotIn applies the NotIn predicate on the "question_bank_id" field.
func QuestionBankIDNotIn(vs ...int) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQuestionBankID), v...))
	})
}

// QuestionBankIDGT applies the GT predicate on the "question_bank_id" field.
func QuestionBankIDGT(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQuestionBankID), v))
	})
}

// QuestionBankIDGTE applies the GTE predicate on the "question_bank_id" field.
func QuestionBankIDGTE(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQuestionBankID), v))
	})
}

// QuestionBankIDLT applies the LT predicate on the "question_bank_id" field.
func QuestionBankIDLT(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQuestionBankID), v))
	})
}

// QuestionBankIDLTE applies the LTE predicate on the "question_bank_id" field.
func QuestionBankIDLTE(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQuestionBankID), v))
	})
}

// QuestionBankIDIsNil applies the IsNil predicate on the "question_bank_id" field.
func QuestionBankIDIsNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldQuestionBankID)))
	})
}

// QuestionBankIDNotNil applies the NotNil predicate on the "question_bank_id" field.
func QuestionBankIDNotNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldQuestionBankID)))
	})
}

// AnswerEQ applies the EQ predicate on the "answer" field.
func AnswerEQ(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnswer), v))
	})
}

// AnswerNEQ applies the NEQ predicate on the "answer" field.
func AnswerNEQ(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAnswer), v))
	})
}

// AnswerIn applies the In predicate on the "answer" field.
func AnswerIn(vs ...string) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAnswer), v...))
	})
}

// AnswerNotIn applies the NotIn predicate on the "answer" field.
func AnswerNotIn(vs ...string) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAnswer), v...))
	})
}

// AnswerGT applies the GT predicate on the "answer" field.
func AnswerGT(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAnswer), v))
	})
}

// AnswerGTE applies the GTE predicate on the "answer" field.
func AnswerGTE(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAnswer), v))
	})
}

// AnswerLT applies the LT predicate on the "answer" field.
func AnswerLT(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAnswer), v))
	})
}

// AnswerLTE applies the LTE predicate on the "answer" field.
func AnswerLTE(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAnswer), v))
	})
}

// AnswerContains applies the Contains predicate on the "answer" field.
func AnswerContains(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAnswer), v))
	})
}

// AnswerHasPrefix applies the HasPrefix predicate on the "answer" field.
func AnswerHasPrefix(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAnswer), v))
	})
}

// AnswerHasSuffix applies the HasSuffix predicate on the "answer" field.
func AnswerHasSuffix(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAnswer), v))
	})
}

// AnswerIsNil applies the IsNil predicate on the "answer" field.
func AnswerIsNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAnswer)))
	})
}

// AnswerNotNil applies the NotNil predicate on the "answer" field.
func AnswerNotNil() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAnswer)))
	})
}

// AnswerEqualFold applies the EqualFold predicate on the "answer" field.
func AnswerEqualFold(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAnswer), v))
	})
}

// AnswerContainsFold applies the ContainsFold predicate on the "answer" field.
func AnswerContainsFold(v string) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAnswer), v))
	})
}

// WrongExamTypeEQ applies the EQ predicate on the "wrong_exam_type" field.
func WrongExamTypeEQ(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWrongExamType), v))
	})
}

// WrongExamTypeNEQ applies the NEQ predicate on the "wrong_exam_type" field.
func WrongExamTypeNEQ(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWrongExamType), v))
	})
}

// WrongExamTypeIn applies the In predicate on the "wrong_exam_type" field.
func WrongExamTypeIn(vs ...int) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWrongExamType), v...))
	})
}

// WrongExamTypeNotIn applies the NotIn predicate on the "wrong_exam_type" field.
func WrongExamTypeNotIn(vs ...int) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWrongExamType), v...))
	})
}

// WrongExamTypeGT applies the GT predicate on the "wrong_exam_type" field.
func WrongExamTypeGT(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWrongExamType), v))
	})
}

// WrongExamTypeGTE applies the GTE predicate on the "wrong_exam_type" field.
func WrongExamTypeGTE(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWrongExamType), v))
	})
}

// WrongExamTypeLT applies the LT predicate on the "wrong_exam_type" field.
func WrongExamTypeLT(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWrongExamType), v))
	})
}

// WrongExamTypeLTE applies the LTE predicate on the "wrong_exam_type" field.
func WrongExamTypeLTE(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWrongExamType), v))
	})
}

// WrongQuestionTypeEQ applies the EQ predicate on the "wrong_question_type" field.
func WrongQuestionTypeEQ(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWrongQuestionType), v))
	})
}

// WrongQuestionTypeNEQ applies the NEQ predicate on the "wrong_question_type" field.
func WrongQuestionTypeNEQ(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWrongQuestionType), v))
	})
}

// WrongQuestionTypeIn applies the In predicate on the "wrong_question_type" field.
func WrongQuestionTypeIn(vs ...int) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWrongQuestionType), v...))
	})
}

// WrongQuestionTypeNotIn applies the NotIn predicate on the "wrong_question_type" field.
func WrongQuestionTypeNotIn(vs ...int) predicate.TkUserWrongQuestionRecode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWrongQuestionType), v...))
	})
}

// WrongQuestionTypeGT applies the GT predicate on the "wrong_question_type" field.
func WrongQuestionTypeGT(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWrongQuestionType), v))
	})
}

// WrongQuestionTypeGTE applies the GTE predicate on the "wrong_question_type" field.
func WrongQuestionTypeGTE(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWrongQuestionType), v))
	})
}

// WrongQuestionTypeLT applies the LT predicate on the "wrong_question_type" field.
func WrongQuestionTypeLT(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWrongQuestionType), v))
	})
}

// WrongQuestionTypeLTE applies the LTE predicate on the "wrong_question_type" field.
func WrongQuestionTypeLTE(v int) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWrongQuestionType), v))
	})
}

// HasQuestionWrong applies the HasEdge predicate on the "question_wrong" edge.
func HasQuestionWrong() predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuestionWrongTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, QuestionWrongTable, QuestionWrongColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionWrongWith applies the HasEdge predicate on the "question_wrong" edge with a given conditions (other predicates).
func HasQuestionWrongWith(preds ...predicate.TkQuestion) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuestionWrongInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, QuestionWrongTable, QuestionWrongColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TkUserWrongQuestionRecode) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TkUserWrongQuestionRecode) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
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
func Not(p predicate.TkUserWrongQuestionRecode) predicate.TkUserWrongQuestionRecode {
	return predicate.TkUserWrongQuestionRecode(func(s *sql.Selector) {
		p(s.Not())
	})
}
