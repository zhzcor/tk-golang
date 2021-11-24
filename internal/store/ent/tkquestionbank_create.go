// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tkserver/internal/store/ent/admin"
	"tkserver/internal/store/ent/itemcategory"
	"tkserver/internal/store/ent/kccourse"
	"tkserver/internal/store/ent/level"
	"tkserver/internal/store/ent/tkchapter"
	"tkserver/internal/store/ent/tkexampaper"
	"tkserver/internal/store/ent/tkexamquestiontype"
	"tkserver/internal/store/ent/tkknowledgepoint"
	"tkserver/internal/store/ent/tkquestion"
	"tkserver/internal/store/ent/tkquestionbank"
	"tkserver/internal/store/ent/tkquestionbankcity"
	"tkserver/internal/store/ent/tkquestionbankmajor"
	"tkserver/internal/store/ent/tkuserquestionbankrecord"
	"tkserver/internal/store/ent/tkuserquestionrecord"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TkQuestionBankCreate is the builder for creating a TkQuestionBank entity.
type TkQuestionBankCreate struct {
	config
	mutation *TkQuestionBankMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (tqbc *TkQuestionBankCreate) SetUUID(s string) *TkQuestionBankCreate {
	tqbc.mutation.SetUUID(s)
	return tqbc
}

// SetCreatedAt sets the "created_at" field.
func (tqbc *TkQuestionBankCreate) SetCreatedAt(t time.Time) *TkQuestionBankCreate {
	tqbc.mutation.SetCreatedAt(t)
	return tqbc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableCreatedAt(t *time.Time) *TkQuestionBankCreate {
	if t != nil {
		tqbc.SetCreatedAt(*t)
	}
	return tqbc
}

// SetUpdatedAt sets the "updated_at" field.
func (tqbc *TkQuestionBankCreate) SetUpdatedAt(t time.Time) *TkQuestionBankCreate {
	tqbc.mutation.SetUpdatedAt(t)
	return tqbc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableUpdatedAt(t *time.Time) *TkQuestionBankCreate {
	if t != nil {
		tqbc.SetUpdatedAt(*t)
	}
	return tqbc
}

// SetDeletedAt sets the "deleted_at" field.
func (tqbc *TkQuestionBankCreate) SetDeletedAt(t time.Time) *TkQuestionBankCreate {
	tqbc.mutation.SetDeletedAt(t)
	return tqbc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableDeletedAt(t *time.Time) *TkQuestionBankCreate {
	if t != nil {
		tqbc.SetDeletedAt(*t)
	}
	return tqbc
}

// SetName sets the "name" field.
func (tqbc *TkQuestionBankCreate) SetName(s string) *TkQuestionBankCreate {
	tqbc.mutation.SetName(s)
	return tqbc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableName(s *string) *TkQuestionBankCreate {
	if s != nil {
		tqbc.SetName(*s)
	}
	return tqbc
}

// SetStatus sets the "status" field.
func (tqbc *TkQuestionBankCreate) SetStatus(u uint8) *TkQuestionBankCreate {
	tqbc.mutation.SetStatus(u)
	return tqbc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableStatus(u *uint8) *TkQuestionBankCreate {
	if u != nil {
		tqbc.SetStatus(*u)
	}
	return tqbc
}

// SetQuestionCount sets the "question_count" field.
func (tqbc *TkQuestionBankCreate) SetQuestionCount(i int) *TkQuestionBankCreate {
	tqbc.mutation.SetQuestionCount(i)
	return tqbc
}

// SetNillableQuestionCount sets the "question_count" field if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableQuestionCount(i *int) *TkQuestionBankCreate {
	if i != nil {
		tqbc.SetQuestionCount(*i)
	}
	return tqbc
}

// SetCreatedAdminID sets the "created_admin_id" field.
func (tqbc *TkQuestionBankCreate) SetCreatedAdminID(i int) *TkQuestionBankCreate {
	tqbc.mutation.SetCreatedAdminID(i)
	return tqbc
}

// SetNillableCreatedAdminID sets the "created_admin_id" field if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableCreatedAdminID(i *int) *TkQuestionBankCreate {
	if i != nil {
		tqbc.SetCreatedAdminID(*i)
	}
	return tqbc
}

// SetItemCategoryID sets the "item_category_id" field.
func (tqbc *TkQuestionBankCreate) SetItemCategoryID(i int) *TkQuestionBankCreate {
	tqbc.mutation.SetItemCategoryID(i)
	return tqbc
}

// SetNillableItemCategoryID sets the "item_category_id" field if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableItemCategoryID(i *int) *TkQuestionBankCreate {
	if i != nil {
		tqbc.SetItemCategoryID(*i)
	}
	return tqbc
}

// SetLevelID sets the "level_id" field.
func (tqbc *TkQuestionBankCreate) SetLevelID(i int) *TkQuestionBankCreate {
	tqbc.mutation.SetLevelID(i)
	return tqbc
}

// SetNillableLevelID sets the "level_id" field if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableLevelID(i *int) *TkQuestionBankCreate {
	if i != nil {
		tqbc.SetLevelID(*i)
	}
	return tqbc
}

// SetSortOrder sets the "sort_order" field.
func (tqbc *TkQuestionBankCreate) SetSortOrder(i int) *TkQuestionBankCreate {
	tqbc.mutation.SetSortOrder(i)
	return tqbc
}

// SetNillableSortOrder sets the "sort_order" field if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableSortOrder(i *int) *TkQuestionBankCreate {
	if i != nil {
		tqbc.SetSortOrder(*i)
	}
	return tqbc
}

// SetItemCategory sets the "item_category" edge to the ItemCategory entity.
func (tqbc *TkQuestionBankCreate) SetItemCategory(i *ItemCategory) *TkQuestionBankCreate {
	return tqbc.SetItemCategoryID(i.ID)
}

// SetLevel sets the "level" edge to the Level entity.
func (tqbc *TkQuestionBankCreate) SetLevel(l *Level) *TkQuestionBankCreate {
	return tqbc.SetLevelID(l.ID)
}

// SetAdminID sets the "admin" edge to the Admin entity by ID.
func (tqbc *TkQuestionBankCreate) SetAdminID(id int) *TkQuestionBankCreate {
	tqbc.mutation.SetAdminID(id)
	return tqbc
}

// SetNillableAdminID sets the "admin" edge to the Admin entity by ID if the given value is not nil.
func (tqbc *TkQuestionBankCreate) SetNillableAdminID(id *int) *TkQuestionBankCreate {
	if id != nil {
		tqbc = tqbc.SetAdminID(*id)
	}
	return tqbc
}

// SetAdmin sets the "admin" edge to the Admin entity.
func (tqbc *TkQuestionBankCreate) SetAdmin(a *Admin) *TkQuestionBankCreate {
	return tqbc.SetAdminID(a.ID)
}

// AddQuestionChapterIDs adds the "question_chapters" edge to the TkChapter entity by IDs.
func (tqbc *TkQuestionBankCreate) AddQuestionChapterIDs(ids ...int) *TkQuestionBankCreate {
	tqbc.mutation.AddQuestionChapterIDs(ids...)
	return tqbc
}

// AddQuestionChapters adds the "question_chapters" edges to the TkChapter entity.
func (tqbc *TkQuestionBankCreate) AddQuestionChapters(t ...*TkChapter) *TkQuestionBankCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tqbc.AddQuestionChapterIDs(ids...)
}

// AddQuestionBankCourseIDs adds the "question_bank_courses" edge to the KcCourse entity by IDs.
func (tqbc *TkQuestionBankCreate) AddQuestionBankCourseIDs(ids ...int) *TkQuestionBankCreate {
	tqbc.mutation.AddQuestionBankCourseIDs(ids...)
	return tqbc
}

// AddQuestionBankCourses adds the "question_bank_courses" edges to the KcCourse entity.
func (tqbc *TkQuestionBankCreate) AddQuestionBankCourses(k ...*KcCourse) *TkQuestionBankCreate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return tqbc.AddQuestionBankCourseIDs(ids...)
}

// AddQuestionIDs adds the "questions" edge to the TkQuestion entity by IDs.
func (tqbc *TkQuestionBankCreate) AddQuestionIDs(ids ...int) *TkQuestionBankCreate {
	tqbc.mutation.AddQuestionIDs(ids...)
	return tqbc
}

// AddQuestions adds the "questions" edges to the TkQuestion entity.
func (tqbc *TkQuestionBankCreate) AddQuestions(t ...*TkQuestion) *TkQuestionBankCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tqbc.AddQuestionIDs(ids...)
}

// AddExamPaperIDs adds the "exam_papers" edge to the TkExamPaper entity by IDs.
func (tqbc *TkQuestionBankCreate) AddExamPaperIDs(ids ...int) *TkQuestionBankCreate {
	tqbc.mutation.AddExamPaperIDs(ids...)
	return tqbc
}

// AddExamPapers adds the "exam_papers" edges to the TkExamPaper entity.
func (tqbc *TkQuestionBankCreate) AddExamPapers(t ...*TkExamPaper) *TkQuestionBankCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tqbc.AddExamPaperIDs(ids...)
}

// AddExamQuestionTypeIDs adds the "exam_question_types" edge to the TkExamQuestionType entity by IDs.
func (tqbc *TkQuestionBankCreate) AddExamQuestionTypeIDs(ids ...int) *TkQuestionBankCreate {
	tqbc.mutation.AddExamQuestionTypeIDs(ids...)
	return tqbc
}

// AddExamQuestionTypes adds the "exam_question_types" edges to the TkExamQuestionType entity.
func (tqbc *TkQuestionBankCreate) AddExamQuestionTypes(t ...*TkExamQuestionType) *TkQuestionBankCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tqbc.AddExamQuestionTypeIDs(ids...)
}

// AddUserQuestionBankIDs adds the "user_question_bank" edge to the TkUserQuestionBankRecord entity by IDs.
func (tqbc *TkQuestionBankCreate) AddUserQuestionBankIDs(ids ...int) *TkQuestionBankCreate {
	tqbc.mutation.AddUserQuestionBankIDs(ids...)
	return tqbc
}

// AddUserQuestionBank adds the "user_question_bank" edges to the TkUserQuestionBankRecord entity.
func (tqbc *TkQuestionBankCreate) AddUserQuestionBank(t ...*TkUserQuestionBankRecord) *TkQuestionBankCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tqbc.AddUserQuestionBankIDs(ids...)
}

// AddUserBankRecordIDs adds the "user_bank_records" edge to the TkUserQuestionRecord entity by IDs.
func (tqbc *TkQuestionBankCreate) AddUserBankRecordIDs(ids ...int) *TkQuestionBankCreate {
	tqbc.mutation.AddUserBankRecordIDs(ids...)
	return tqbc
}

// AddUserBankRecords adds the "user_bank_records" edges to the TkUserQuestionRecord entity.
func (tqbc *TkQuestionBankCreate) AddUserBankRecords(t ...*TkUserQuestionRecord) *TkQuestionBankCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tqbc.AddUserBankRecordIDs(ids...)
}

// AddKnowledgePointIDs adds the "knowledge_points" edge to the TkKnowledgePoint entity by IDs.
func (tqbc *TkQuestionBankCreate) AddKnowledgePointIDs(ids ...int) *TkQuestionBankCreate {
	tqbc.mutation.AddKnowledgePointIDs(ids...)
	return tqbc
}

// AddKnowledgePoints adds the "knowledge_points" edges to the TkKnowledgePoint entity.
func (tqbc *TkQuestionBankCreate) AddKnowledgePoints(t ...*TkKnowledgePoint) *TkQuestionBankCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tqbc.AddKnowledgePointIDs(ids...)
}

// AddCityQuestionBankIDs adds the "city_question_banks" edge to the TkQuestionBankCity entity by IDs.
func (tqbc *TkQuestionBankCreate) AddCityQuestionBankIDs(ids ...int) *TkQuestionBankCreate {
	tqbc.mutation.AddCityQuestionBankIDs(ids...)
	return tqbc
}

// AddCityQuestionBanks adds the "city_question_banks" edges to the TkQuestionBankCity entity.
func (tqbc *TkQuestionBankCreate) AddCityQuestionBanks(t ...*TkQuestionBankCity) *TkQuestionBankCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tqbc.AddCityQuestionBankIDs(ids...)
}

// AddMajorQuestionBankIDs adds the "major_question_banks" edge to the TkQuestionBankMajor entity by IDs.
func (tqbc *TkQuestionBankCreate) AddMajorQuestionBankIDs(ids ...int) *TkQuestionBankCreate {
	tqbc.mutation.AddMajorQuestionBankIDs(ids...)
	return tqbc
}

// AddMajorQuestionBanks adds the "major_question_banks" edges to the TkQuestionBankMajor entity.
func (tqbc *TkQuestionBankCreate) AddMajorQuestionBanks(t ...*TkQuestionBankMajor) *TkQuestionBankCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tqbc.AddMajorQuestionBankIDs(ids...)
}

// Mutation returns the TkQuestionBankMutation object of the builder.
func (tqbc *TkQuestionBankCreate) Mutation() *TkQuestionBankMutation {
	return tqbc.mutation
}

// Save creates the TkQuestionBank in the database.
func (tqbc *TkQuestionBankCreate) Save(ctx context.Context) (*TkQuestionBank, error) {
	var (
		err  error
		node *TkQuestionBank
	)
	tqbc.defaults()
	if len(tqbc.hooks) == 0 {
		if err = tqbc.check(); err != nil {
			return nil, err
		}
		node, err = tqbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TkQuestionBankMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tqbc.check(); err != nil {
				return nil, err
			}
			tqbc.mutation = mutation
			node, err = tqbc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tqbc.hooks) - 1; i >= 0; i-- {
			mut = tqbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tqbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tqbc *TkQuestionBankCreate) SaveX(ctx context.Context) *TkQuestionBank {
	v, err := tqbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tqbc *TkQuestionBankCreate) defaults() {
	if _, ok := tqbc.mutation.CreatedAt(); !ok {
		v := tkquestionbank.DefaultCreatedAt()
		tqbc.mutation.SetCreatedAt(v)
	}
	if _, ok := tqbc.mutation.UpdatedAt(); !ok {
		v := tkquestionbank.DefaultUpdatedAt()
		tqbc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tqbc.mutation.Name(); !ok {
		v := tkquestionbank.DefaultName
		tqbc.mutation.SetName(v)
	}
	if _, ok := tqbc.mutation.Status(); !ok {
		v := tkquestionbank.DefaultStatus
		tqbc.mutation.SetStatus(v)
	}
	if _, ok := tqbc.mutation.QuestionCount(); !ok {
		v := tkquestionbank.DefaultQuestionCount
		tqbc.mutation.SetQuestionCount(v)
	}
	if _, ok := tqbc.mutation.SortOrder(); !ok {
		v := tkquestionbank.DefaultSortOrder
		tqbc.mutation.SetSortOrder(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tqbc *TkQuestionBankCreate) check() error {
	if _, ok := tqbc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := tqbc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := tqbc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := tqbc.mutation.QuestionCount(); !ok {
		return &ValidationError{Name: "question_count", err: errors.New("ent: missing required field \"question_count\"")}
	}
	if _, ok := tqbc.mutation.SortOrder(); !ok {
		return &ValidationError{Name: "sort_order", err: errors.New("ent: missing required field \"sort_order\"")}
	}
	return nil
}

func (tqbc *TkQuestionBankCreate) sqlSave(ctx context.Context) (*TkQuestionBank, error) {
	_node, _spec := tqbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tqbc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tqbc *TkQuestionBankCreate) createSpec() (*TkQuestionBank, *sqlgraph.CreateSpec) {
	var (
		_node = &TkQuestionBank{config: tqbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tkquestionbank.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tkquestionbank.FieldID,
			},
		}
	)
	if value, ok := tqbc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkquestionbank.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := tqbc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkquestionbank.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := tqbc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkquestionbank.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := tqbc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tkquestionbank.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := tqbc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tkquestionbank.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tqbc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: tkquestionbank.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := tqbc.mutation.QuestionCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkquestionbank.FieldQuestionCount,
		})
		_node.QuestionCount = value
	}
	if value, ok := tqbc.mutation.SortOrder(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tkquestionbank.FieldSortOrder,
		})
		_node.SortOrder = value
	}
	if nodes := tqbc.mutation.ItemCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbank.ItemCategoryTable,
			Columns: []string{tkquestionbank.ItemCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: itemcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ItemCategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.LevelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbank.LevelTable,
			Columns: []string{tkquestionbank.LevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: level.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.LevelID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.AdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tkquestionbank.AdminTable,
			Columns: []string{tkquestionbank.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreatedAdminID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.QuestionChaptersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkquestionbank.QuestionChaptersTable,
			Columns: []string{tkquestionbank.QuestionChaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkchapter.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.QuestionBankCoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkquestionbank.QuestionBankCoursesTable,
			Columns: []string{tkquestionbank.QuestionBankCoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kccourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkquestionbank.QuestionsTable,
			Columns: []string{tkquestionbank.QuestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.ExamPapersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkquestionbank.ExamPapersTable,
			Columns: []string{tkquestionbank.ExamPapersColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.ExamQuestionTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkquestionbank.ExamQuestionTypesTable,
			Columns: []string{tkquestionbank.ExamQuestionTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkexamquestiontype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.UserQuestionBankIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkquestionbank.UserQuestionBankTable,
			Columns: []string{tkquestionbank.UserQuestionBankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkuserquestionbankrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.UserBankRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkquestionbank.UserBankRecordsTable,
			Columns: []string{tkquestionbank.UserBankRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkuserquestionrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.KnowledgePointsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkquestionbank.KnowledgePointsTable,
			Columns: []string{tkquestionbank.KnowledgePointsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkknowledgepoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.CityQuestionBanksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkquestionbank.CityQuestionBanksTable,
			Columns: []string{tkquestionbank.CityQuestionBanksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestionbankcity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tqbc.mutation.MajorQuestionBanksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tkquestionbank.MajorQuestionBanksTable,
			Columns: []string{tkquestionbank.MajorQuestionBanksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkquestionbankmajor.FieldID,
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

// TkQuestionBankCreateBulk is the builder for creating many TkQuestionBank entities in bulk.
type TkQuestionBankCreateBulk struct {
	config
	builders []*TkQuestionBankCreate
}

// Save creates the TkQuestionBank entities in the database.
func (tqbcb *TkQuestionBankCreateBulk) Save(ctx context.Context) ([]*TkQuestionBank, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tqbcb.builders))
	nodes := make([]*TkQuestionBank, len(tqbcb.builders))
	mutators := make([]Mutator, len(tqbcb.builders))
	for i := range tqbcb.builders {
		func(i int, root context.Context) {
			builder := tqbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TkQuestionBankMutation)
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
					_, err = mutators[i+1].Mutate(root, tqbcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tqbcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tqbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tqbcb *TkQuestionBankCreateBulk) SaveX(ctx context.Context) []*TkQuestionBank {
	v, err := tqbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
