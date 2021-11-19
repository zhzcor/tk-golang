// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/teacher"
	"gserver/internal/store/ent/tkexampaper"
	"gserver/internal/store/ent/tksection"
	"gserver/internal/store/ent/tkuserexamscorerecord"
	"gserver/internal/store/ent/tkusersimulationteachermark"
	"gserver/internal/store/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkUserExamScoreRecordCreate is the builder for creating a TkUserExamScoreRecord entity.
type TkUserExamScoreRecordCreate struct {
	config
	mutation *TkUserExamScoreRecordMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetUUID(s string) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetUUID(s)
	return tuesrc
}

// SetCreatedAt sets the "created_at" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetCreatedAt(t time.Time) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetCreatedAt(t)
	return tuesrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableCreatedAt(t *time.Time) *TkUserExamScoreRecordCreate {
	if t != nil {
		tuesrc.SetCreatedAt(*t)
	}
	return tuesrc
}

// SetUpdatedAt sets the "updated_at" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetUpdatedAt(t time.Time) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetUpdatedAt(t)
	return tuesrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableUpdatedAt(t *time.Time) *TkUserExamScoreRecordCreate {
	if t != nil {
		tuesrc.SetUpdatedAt(*t)
	}
	return tuesrc
}

// SetDeletedAt sets the "deleted_at" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetDeletedAt(t time.Time) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetDeletedAt(t)
	return tuesrc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableDeletedAt(t *time.Time) *TkUserExamScoreRecordCreate {
	if t != nil {
		tuesrc.SetDeletedAt(*t)
	}
	return tuesrc
}

// SetSubjectiveQuestionScore sets the "subjective_question_score" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetSubjectiveQuestionScore(u uint8) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetSubjectiveQuestionScore(u)
	return tuesrc
}

// SetNillableSubjectiveQuestionScore sets the "subjective_question_score" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableSubjectiveQuestionScore(u *uint8) *TkUserExamScoreRecordCreate {
	if u != nil {
		tuesrc.SetSubjectiveQuestionScore(*u)
	}
	return tuesrc
}

// SetObjectiveQuestionScore sets the "objective_question_score" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetObjectiveQuestionScore(u uint8) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetObjectiveQuestionScore(u)
	return tuesrc
}

// SetNillableObjectiveQuestionScore sets the "objective_question_score" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableObjectiveQuestionScore(u *uint8) *TkUserExamScoreRecordCreate {
	if u != nil {
		tuesrc.SetObjectiveQuestionScore(*u)
	}
	return tuesrc
}

// SetTotalScore sets the "total_score" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetTotalScore(u uint8) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetTotalScore(u)
	return tuesrc
}

// SetNillableTotalScore sets the "total_score" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableTotalScore(u *uint8) *TkUserExamScoreRecordCreate {
	if u != nil {
		tuesrc.SetTotalScore(*u)
	}
	return tuesrc
}

// SetDuration sets the "duration" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetDuration(i int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetDuration(i)
	return tuesrc
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableDuration(i *int) *TkUserExamScoreRecordCreate {
	if i != nil {
		tuesrc.SetDuration(*i)
	}
	return tuesrc
}

// SetRightCount sets the "right_count" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetRightCount(i int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetRightCount(i)
	return tuesrc
}

// SetNillableRightCount sets the "right_count" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableRightCount(i *int) *TkUserExamScoreRecordCreate {
	if i != nil {
		tuesrc.SetRightCount(*i)
	}
	return tuesrc
}

// SetWrongCount sets the "wrong_count" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetWrongCount(i int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetWrongCount(i)
	return tuesrc
}

// SetNillableWrongCount sets the "wrong_count" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableWrongCount(i *int) *TkUserExamScoreRecordCreate {
	if i != nil {
		tuesrc.SetWrongCount(*i)
	}
	return tuesrc
}

// SetTotalCount sets the "total_count" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetTotalCount(i int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetTotalCount(i)
	return tuesrc
}

// SetNillableTotalCount sets the "total_count" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableTotalCount(i *int) *TkUserExamScoreRecordCreate {
	if i != nil {
		tuesrc.SetTotalCount(*i)
	}
	return tuesrc
}

// SetNoAnswerCount sets the "no_answer_count" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetNoAnswerCount(i int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetNoAnswerCount(i)
	return tuesrc
}

// SetNillableNoAnswerCount sets the "no_answer_count" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableNoAnswerCount(i *int) *TkUserExamScoreRecordCreate {
	if i != nil {
		tuesrc.SetNoAnswerCount(*i)
	}
	return tuesrc
}

// SetRank sets the "rank" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetRank(i int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetRank(i)
	return tuesrc
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableRank(i *int) *TkUserExamScoreRecordCreate {
	if i != nil {
		tuesrc.SetRank(*i)
	}
	return tuesrc
}

// SetStatus sets the "status" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetStatus(u uint8) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetStatus(u)
	return tuesrc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableStatus(u *uint8) *TkUserExamScoreRecordCreate {
	if u != nil {
		tuesrc.SetStatus(*u)
	}
	return tuesrc
}

// SetOrderStatus sets the "order_status" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetOrderStatus(u uint8) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetOrderStatus(u)
	return tuesrc
}

// SetNillableOrderStatus sets the "order_status" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableOrderStatus(u *uint8) *TkUserExamScoreRecordCreate {
	if u != nil {
		tuesrc.SetOrderStatus(*u)
	}
	return tuesrc
}

// SetExamPaperID sets the "exam_paper_id" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetExamPaperID(i int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetExamPaperID(i)
	return tuesrc
}

// SetNillableExamPaperID sets the "exam_paper_id" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableExamPaperID(i *int) *TkUserExamScoreRecordCreate {
	if i != nil {
		tuesrc.SetExamPaperID(*i)
	}
	return tuesrc
}

// SetSectionID sets the "section_id" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetSectionID(i int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetSectionID(i)
	return tuesrc
}

// SetNillableSectionID sets the "section_id" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableSectionID(i *int) *TkUserExamScoreRecordCreate {
	if i != nil {
		tuesrc.SetSectionID(*i)
	}
	return tuesrc
}

// SetUserID sets the "user_id" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetUserID(i int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetUserID(i)
	return tuesrc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableUserID(i *int) *TkUserExamScoreRecordCreate {
	if i != nil {
		tuesrc.SetUserID(*i)
	}
	return tuesrc
}

// SetOperationTeacherID sets the "operation_teacher_id" field.
func (tuesrc *TkUserExamScoreRecordCreate) SetOperationTeacherID(i int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetOperationTeacherID(i)
	return tuesrc
}

// SetNillableOperationTeacherID sets the "operation_teacher_id" field if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableOperationTeacherID(i *int) *TkUserExamScoreRecordCreate {
	if i != nil {
		tuesrc.SetOperationTeacherID(*i)
	}
	return tuesrc
}

// SetExamPaper sets the "exam_paper" edge to the TkExamPaper entity.
func (tuesrc *TkUserExamScoreRecordCreate) SetExamPaper(t *TkExamPaper) *TkUserExamScoreRecordCreate {
	return tuesrc.SetExamPaperID(t.ID)
}

// SetSection sets the "section" edge to the TkSection entity.
func (tuesrc *TkUserExamScoreRecordCreate) SetSection(t *TkSection) *TkUserExamScoreRecordCreate {
	return tuesrc.SetSectionID(t.ID)
}

// SetTeacherID sets the "teacher" edge to the Teacher entity by ID.
func (tuesrc *TkUserExamScoreRecordCreate) SetTeacherID(id int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.SetTeacherID(id)
	return tuesrc
}

// SetNillableTeacherID sets the "teacher" edge to the Teacher entity by ID if the given value is not nil.
func (tuesrc *TkUserExamScoreRecordCreate) SetNillableTeacherID(id *int) *TkUserExamScoreRecordCreate {
	if id != nil {
		tuesrc = tuesrc.SetTeacherID(*id)
	}
	return tuesrc
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (tuesrc *TkUserExamScoreRecordCreate) SetTeacher(t *Teacher) *TkUserExamScoreRecordCreate {
	return tuesrc.SetTeacherID(t.ID)
}

// SetUser sets the "user" edge to the User entity.
func (tuesrc *TkUserExamScoreRecordCreate) SetUser(u *User) *TkUserExamScoreRecordCreate {
	return tuesrc.SetUserID(u.ID)
}

// AddUserExamDetailIDs adds the "user_exam_details" edge to the TkUserSimulationTeacherMark entity by IDs.
func (tuesrc *TkUserExamScoreRecordCreate) AddUserExamDetailIDs(ids ...int) *TkUserExamScoreRecordCreate {
	tuesrc.mutation.AddUserExamDetailIDs(ids...)
	return tuesrc
}

// AddUserExamDetails adds the "user_exam_details" edges to the TkUserSimulationTeacherMark entity.
func (tuesrc *TkUserExamScoreRecordCreate) AddUserExamDetails(t ...*TkUserSimulationTeacherMark) *TkUserExamScoreRecordCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuesrc.AddUserExamDetailIDs(ids...)
}

// Mutation returns the TkUserExamScoreRecordMutation object of the builder.
func (tuesrc *TkUserExamScoreRecordCreate) Mutation() *TkUserExamScoreRecordMutation {
	return tuesrc.mutation
}

// Save creates the TkUserExamScoreRecord in the database.
func (tuesrc *TkUserExamScoreRecordCreate) Save(ctx context.Context) (*TkUserExamScoreRecord, error) {
	var (
		err  error
		node *TkUserExamScoreRecord
	)
	tuesrc.defaults()
	if len(tuesrc.hooks) == 0 {
		if err = tuesrc.check(); err != nil {
			return nil, err
		}
		node, err = tuesrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkUserExamScoreRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuesrc.check(); err != nil {
				return nil, err
			}
			tuesrc.mutation = mutation
			node, err = tuesrc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuesrc.hooks) - 1; i >= 0; i-- {
			mut = tuesrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuesrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tuesrc *TkUserExamScoreRecordCreate) SaveX(ctx context.Context) *TkUserExamScoreRecord {
	v, err := tuesrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tuesrc *TkUserExamScoreRecordCreate) defaults() {
	if _, ok := tuesrc.mutation.CreatedAt(); !ok {
		v := tkuserexamscorerecord.DefaultCreatedAt()
		tuesrc.mutation.SetCreatedAt(v)
	}
	if _, ok := tuesrc.mutation.UpdatedAt(); !ok {
		v := tkuserexamscorerecord.DefaultUpdatedAt()
		tuesrc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tuesrc.mutation.SubjectiveQuestionScore(); !ok {
		v := tkuserexamscorerecord.DefaultSubjectiveQuestionScore
		tuesrc.mutation.SetSubjectiveQuestionScore(v)
	}
	if _, ok := tuesrc.mutation.ObjectiveQuestionScore(); !ok {
		v := tkuserexamscorerecord.DefaultObjectiveQuestionScore
		tuesrc.mutation.SetObjectiveQuestionScore(v)
	}
	if _, ok := tuesrc.mutation.TotalScore(); !ok {
		v := tkuserexamscorerecord.DefaultTotalScore
		tuesrc.mutation.SetTotalScore(v)
	}
	if _, ok := tuesrc.mutation.Duration(); !ok {
		v := tkuserexamscorerecord.DefaultDuration
		tuesrc.mutation.SetDuration(v)
	}
	if _, ok := tuesrc.mutation.RightCount(); !ok {
		v := tkuserexamscorerecord.DefaultRightCount
		tuesrc.mutation.SetRightCount(v)
	}
	if _, ok := tuesrc.mutation.WrongCount(); !ok {
		v := tkuserexamscorerecord.DefaultWrongCount
		tuesrc.mutation.SetWrongCount(v)
	}
	if _, ok := tuesrc.mutation.TotalCount(); !ok {
		v := tkuserexamscorerecord.DefaultTotalCount
		tuesrc.mutation.SetTotalCount(v)
	}
	if _, ok := tuesrc.mutation.NoAnswerCount(); !ok {
		v := tkuserexamscorerecord.DefaultNoAnswerCount
		tuesrc.mutation.SetNoAnswerCount(v)
	}
	if _, ok := tuesrc.mutation.Rank(); !ok {
		v := tkuserexamscorerecord.DefaultRank
		tuesrc.mutation.SetRank(v)
	}
	if _, ok := tuesrc.mutation.Status(); !ok {
		v := tkuserexamscorerecord.DefaultStatus
		tuesrc.mutation.SetStatus(v)
	}
	if _, ok := tuesrc.mutation.OrderStatus(); !ok {
		v := tkuserexamscorerecord.DefaultOrderStatus
		tuesrc.mutation.SetOrderStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuesrc *TkUserExamScoreRecordCreate) check() error {
	if _, ok := tuesrc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := tuesrc.mutation.SubjectiveQuestionScore(); !ok {
		return &ValidationError{Name: "subjective_question_score", err: errors.New("ent: missing required field \"subjective_question_score\"")}
	}
	if _, ok := tuesrc.mutation.ObjectiveQuestionScore(); !ok {
		return &ValidationError{Name: "objective_question_score", err: errors.New("ent: missing required field \"objective_question_score\"")}
	}
	if _, ok := tuesrc.mutation.TotalScore(); !ok {
		return &ValidationError{Name: "total_score", err: errors.New("ent: missing required field \"total_score\"")}
	}
	if _, ok := tuesrc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New("ent: missing required field \"duration\"")}
	}
	if _, ok := tuesrc.mutation.RightCount(); !ok {
		return &ValidationError{Name: "right_count", err: errors.New("ent: missing required field \"right_count\"")}
	}
	if _, ok := tuesrc.mutation.WrongCount(); !ok {
		return &ValidationError{Name: "wrong_count", err: errors.New("ent: missing required field \"wrong_count\"")}
	}
	if _, ok := tuesrc.mutation.TotalCount(); !ok {
		return &ValidationError{Name: "total_count", err: errors.New("ent: missing required field \"total_count\"")}
	}
	if _, ok := tuesrc.mutation.NoAnswerCount(); !ok {
		return &ValidationError{Name: "no_answer_count", err: errors.New("ent: missing required field \"no_answer_count\"")}
	}
	if _, ok := tuesrc.mutation.Rank(); !ok {
		return &ValidationError{Name: "rank", err: errors.New("ent: missing required field \"rank\"")}
	}
	if _, ok := tuesrc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := tuesrc.mutation.OrderStatus(); !ok {
		return &ValidationError{Name: "order_status", err: errors.New("ent: missing required field \"order_status\"")}
	}
	return nil
}

func (tuesrc *TkUserExamScoreRecordCreate) sqlSave(ctx context.Context) (*TkUserExamScoreRecord, error) {
	_node, _spec := tuesrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tuesrc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tuesrc *TkUserExamScoreRecordCreate) createSpec() (*TkUserExamScoreRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &TkUserExamScoreRecord{config: tuesrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tkuserexamscorerecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkuserexamscorerecord.FieldID,
			},
		}
	)
	if value, ok := tuesrc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkuserexamscorerecord.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := tuesrc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkuserexamscorerecord.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := tuesrc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkuserexamscorerecord.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := tuesrc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkuserexamscorerecord.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := tuesrc.mutation.SubjectiveQuestionScore(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: tkuserexamscorerecord.FieldSubjectiveQuestionScore,
		})
		_node.SubjectiveQuestionScore = value
	}
	if value, ok := tuesrc.mutation.ObjectiveQuestionScore(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: tkuserexamscorerecord.FieldObjectiveQuestionScore,
		})
		_node.ObjectiveQuestionScore = value
	}
	if value, ok := tuesrc.mutation.TotalScore(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: tkuserexamscorerecord.FieldTotalScore,
		})
		_node.TotalScore = value
	}
	if value, ok := tuesrc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkuserexamscorerecord.FieldDuration,
		})
		_node.Duration = value
	}
	if value, ok := tuesrc.mutation.RightCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkuserexamscorerecord.FieldRightCount,
		})
		_node.RightCount = value
	}
	if value, ok := tuesrc.mutation.WrongCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkuserexamscorerecord.FieldWrongCount,
		})
		_node.WrongCount = value
	}
	if value, ok := tuesrc.mutation.TotalCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkuserexamscorerecord.FieldTotalCount,
		})
		_node.TotalCount = value
	}
	if value, ok := tuesrc.mutation.NoAnswerCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkuserexamscorerecord.FieldNoAnswerCount,
		})
		_node.NoAnswerCount = value
	}
	if value, ok := tuesrc.mutation.Rank(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkuserexamscorerecord.FieldRank,
		})
		_node.Rank = value
	}
	if value, ok := tuesrc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: tkuserexamscorerecord.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := tuesrc.mutation.OrderStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: tkuserexamscorerecord.FieldOrderStatus,
		})
		_node.OrderStatus = value
	}
	if nodes := tuesrc.mutation.ExamPaperIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkuserexamscorerecord.ExamPaperTable,
			Columns: []string{tkuserexamscorerecord.ExamPaperColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkexampaper.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ExamPaperID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tuesrc.mutation.SectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkuserexamscorerecord.SectionTable,
			Columns: []string{tkuserexamscorerecord.SectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tksection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SectionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tuesrc.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkuserexamscorerecord.TeacherTable,
			Columns: []string{tkuserexamscorerecord.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OperationTeacherID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tuesrc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkuserexamscorerecord.UserTable,
			Columns: []string{tkuserexamscorerecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tuesrc.mutation.UserExamDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkuserexamscorerecord.UserExamDetailsTable,
			Columns: []string{tkuserexamscorerecord.UserExamDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkusersimulationteachermark.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TkUserExamScoreRecordCreateBulk is the builder for creating many TkUserExamScoreRecord entities in bulk.
type TkUserExamScoreRecordCreateBulk struct {
	config
	builders []*TkUserExamScoreRecordCreate
}

// Save creates the TkUserExamScoreRecord entities in the database.
func (tuesrcb *TkUserExamScoreRecordCreateBulk) Save(ctx context.Context) ([]*TkUserExamScoreRecord, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tuesrcb.builders))
	nodes := make([]*TkUserExamScoreRecord, len(tuesrcb.builders))
	mutators := make([]Mutator, len(tuesrcb.builders))
	for i := range tuesrcb.builders {
		func(i int, root context.Context) {
			builder := tuesrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TkUserExamScoreRecordMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tuesrcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tuesrcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tuesrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tuesrcb *TkUserExamScoreRecordCreateBulk) SaveX(ctx context.Context) []*TkUserExamScoreRecord {
	v, err := tuesrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
