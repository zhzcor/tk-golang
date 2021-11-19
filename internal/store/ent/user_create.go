// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/city"
	"gserver/internal/store/ent/itemcategory"
	"gserver/internal/store/ent/kcuserclass"
	"gserver/internal/store/ent/kcusercourse"
	"gserver/internal/store/ent/message"
	"gserver/internal/store/ent/tkuserexamscorerecord"
	"gserver/internal/store/ent/tkuserquestionbankrecord"
	"gserver/internal/store/ent/tkuserquestionrecord"
	"gserver/internal/store/ent/user"
	"gserver/internal/store/ent/useraskanswer"
	"gserver/internal/store/ent/usercourseappraise"
	"gserver/internal/store/ent/userloginlog"
	"gserver/internal/store/ent/videorecord"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (uc *UserCreate) SetUUID(s string) *UserCreate {
	uc.mutation.SetUUID(s)
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetDeletedAt sets the "deleted_at" field.
func (uc *UserCreate) SetDeletedAt(t time.Time) *UserCreate {
	uc.mutation.SetDeletedAt(t)
	return uc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeletedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetDeletedAt(*t)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetSalt sets the "salt" field.
func (uc *UserCreate) SetSalt(s string) *UserCreate {
	uc.mutation.SetSalt(s)
	return uc
}

// SetBossUserID sets the "boss_user_id" field.
func (uc *UserCreate) SetBossUserID(i int) *UserCreate {
	uc.mutation.SetBossUserID(i)
	return uc
}

// SetNillableBossUserID sets the "boss_user_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableBossUserID(i *int) *UserCreate {
	if i != nil {
		uc.SetBossUserID(*i)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetNickname sets the "nickname" field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uc *UserCreate) SetNillableNickname(s *string) *UserCreate {
	if s != nil {
		uc.SetNickname(*s)
	}
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uc *UserCreate) SetNillableUsername(s *string) *UserCreate {
	if s != nil {
		uc.SetUsername(*s)
	}
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(u uint8) *UserCreate {
	uc.mutation.SetStatus(u)
	return uc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(u *uint8) *UserCreate {
	if u != nil {
		uc.SetStatus(*u)
	}
	return uc
}

// SetSex sets the "sex" field.
func (uc *UserCreate) SetSex(u uint8) *UserCreate {
	uc.mutation.SetSex(u)
	return uc
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (uc *UserCreate) SetNillableSex(u *uint8) *UserCreate {
	if u != nil {
		uc.SetSex(*u)
	}
	return uc
}

// SetRegFrom sets the "reg_from" field.
func (uc *UserCreate) SetRegFrom(u uint8) *UserCreate {
	uc.mutation.SetRegFrom(u)
	return uc
}

// SetNillableRegFrom sets the "reg_from" field if the given value is not nil.
func (uc *UserCreate) SetNillableRegFrom(u *uint8) *UserCreate {
	if u != nil {
		uc.SetRegFrom(*u)
	}
	return uc
}

// SetCardType sets the "card_type" field.
func (uc *UserCreate) SetCardType(u uint8) *UserCreate {
	uc.mutation.SetCardType(u)
	return uc
}

// SetNillableCardType sets the "card_type" field if the given value is not nil.
func (uc *UserCreate) SetNillableCardType(u *uint8) *UserCreate {
	if u != nil {
		uc.SetCardType(*u)
	}
	return uc
}

// SetIDCard sets the "id_card" field.
func (uc *UserCreate) SetIDCard(s string) *UserCreate {
	uc.mutation.SetIDCard(s)
	return uc
}

// SetNillableIDCard sets the "id_card" field if the given value is not nil.
func (uc *UserCreate) SetNillableIDCard(s *string) *UserCreate {
	if s != nil {
		uc.SetIDCard(*s)
	}
	return uc
}

// SetFromCityID sets the "from_city_id" field.
func (uc *UserCreate) SetFromCityID(i int) *UserCreate {
	uc.mutation.SetFromCityID(i)
	return uc
}

// SetNillableFromCityID sets the "from_city_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableFromCityID(i *int) *UserCreate {
	if i != nil {
		uc.SetFromCityID(*i)
	}
	return uc
}

// SetFromItemCategoryID sets the "from_item_category_id" field.
func (uc *UserCreate) SetFromItemCategoryID(i int) *UserCreate {
	uc.mutation.SetFromItemCategoryID(i)
	return uc
}

// SetNillableFromItemCategoryID sets the "from_item_category_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableFromItemCategoryID(i *int) *UserCreate {
	if i != nil {
		uc.SetFromItemCategoryID(*i)
	}
	return uc
}

// SetBirthday sets the "birthday" field.
func (uc *UserCreate) SetBirthday(t time.Time) *UserCreate {
	uc.mutation.SetBirthday(t)
	return uc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (uc *UserCreate) SetNillableBirthday(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetBirthday(*t)
	}
	return uc
}

// SetSignRemark sets the "sign_remark" field.
func (uc *UserCreate) SetSignRemark(s string) *UserCreate {
	uc.mutation.SetSignRemark(s)
	return uc
}

// SetNillableSignRemark sets the "sign_remark" field if the given value is not nil.
func (uc *UserCreate) SetNillableSignRemark(s *string) *UserCreate {
	if s != nil {
		uc.SetSignRemark(*s)
	}
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// AddLoginLogIDs adds the "login_log" edge to the UserLoginLog entity by IDs.
func (uc *UserCreate) AddLoginLogIDs(ids ...int) *UserCreate {
	uc.mutation.AddLoginLogIDs(ids...)
	return uc
}

// AddLoginLog adds the "login_log" edges to the UserLoginLog entity.
func (uc *UserCreate) AddLoginLog(u ...*UserLoginLog) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddLoginLogIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (uc *UserCreate) AddMessageIDs(ids ...int) *UserCreate {
	uc.mutation.AddMessageIDs(ids...)
	return uc
}

// AddMessages adds the "messages" edges to the Message entity.
func (uc *UserCreate) AddMessages(m ...*Message) *UserCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddMessageIDs(ids...)
}

// AddUserCourseIDs adds the "user_courses" edge to the KcUserCourse entity by IDs.
func (uc *UserCreate) AddUserCourseIDs(ids ...int) *UserCreate {
	uc.mutation.AddUserCourseIDs(ids...)
	return uc
}

// AddUserCourses adds the "user_courses" edges to the KcUserCourse entity.
func (uc *UserCreate) AddUserCourses(k ...*KcUserCourse) *UserCreate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return uc.AddUserCourseIDs(ids...)
}

// AddUserClassIDs adds the "user_classes" edge to the KcUserClass entity by IDs.
func (uc *UserCreate) AddUserClassIDs(ids ...int) *UserCreate {
	uc.mutation.AddUserClassIDs(ids...)
	return uc
}

// AddUserClasses adds the "user_classes" edges to the KcUserClass entity.
func (uc *UserCreate) AddUserClasses(k ...*KcUserClass) *UserCreate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return uc.AddUserClassIDs(ids...)
}

// AddUserExamsRecordIDs adds the "user_exams_records" edge to the TkUserExamScoreRecord entity by IDs.
func (uc *UserCreate) AddUserExamsRecordIDs(ids ...int) *UserCreate {
	uc.mutation.AddUserExamsRecordIDs(ids...)
	return uc
}

// AddUserExamsRecords adds the "user_exams_records" edges to the TkUserExamScoreRecord entity.
func (uc *UserCreate) AddUserExamsRecords(t ...*TkUserExamScoreRecord) *UserCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddUserExamsRecordIDs(ids...)
}

// AddUserQuestionBankRecordIDs adds the "user_question_bank_records" edge to the TkUserQuestionBankRecord entity by IDs.
func (uc *UserCreate) AddUserQuestionBankRecordIDs(ids ...int) *UserCreate {
	uc.mutation.AddUserQuestionBankRecordIDs(ids...)
	return uc
}

// AddUserQuestionBankRecords adds the "user_question_bank_records" edges to the TkUserQuestionBankRecord entity.
func (uc *UserCreate) AddUserQuestionBankRecords(t ...*TkUserQuestionBankRecord) *UserCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddUserQuestionBankRecordIDs(ids...)
}

// AddUserQuestionRecordIDs adds the "user_question_records" edge to the TkUserQuestionRecord entity by IDs.
func (uc *UserCreate) AddUserQuestionRecordIDs(ids ...int) *UserCreate {
	uc.mutation.AddUserQuestionRecordIDs(ids...)
	return uc
}

// AddUserQuestionRecords adds the "user_question_records" edges to the TkUserQuestionRecord entity.
func (uc *UserCreate) AddUserQuestionRecords(t ...*TkUserQuestionRecord) *UserCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddUserQuestionRecordIDs(ids...)
}

// AddAskUserIDs adds the "ask_users" edge to the UserAskAnswer entity by IDs.
func (uc *UserCreate) AddAskUserIDs(ids ...int) *UserCreate {
	uc.mutation.AddAskUserIDs(ids...)
	return uc
}

// AddAskUsers adds the "ask_users" edges to the UserAskAnswer entity.
func (uc *UserCreate) AddAskUsers(u ...*UserAskAnswer) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddAskUserIDs(ids...)
}

// AddCourseAppraiseUserIDs adds the "course_appraise_users" edge to the UserCourseAppraise entity by IDs.
func (uc *UserCreate) AddCourseAppraiseUserIDs(ids ...int) *UserCreate {
	uc.mutation.AddCourseAppraiseUserIDs(ids...)
	return uc
}

// AddCourseAppraiseUsers adds the "course_appraise_users" edges to the UserCourseAppraise entity.
func (uc *UserCreate) AddCourseAppraiseUsers(u ...*UserCourseAppraise) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddCourseAppraiseUserIDs(ids...)
}

// AddUserVideoRecordIDs adds the "user_video_record" edge to the VideoRecord entity by IDs.
func (uc *UserCreate) AddUserVideoRecordIDs(ids ...int) *UserCreate {
	uc.mutation.AddUserVideoRecordIDs(ids...)
	return uc
}

// AddUserVideoRecord adds the "user_video_record" edges to the VideoRecord entity.
func (uc *UserCreate) AddUserVideoRecord(v ...*VideoRecord) *UserCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uc.AddUserVideoRecordIDs(ids...)
}

// SetCityID sets the "city" edge to the City entity by ID.
func (uc *UserCreate) SetCityID(id int) *UserCreate {
	uc.mutation.SetCityID(id)
	return uc
}

// SetNillableCityID sets the "city" edge to the City entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableCityID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetCityID(*id)
	}
	return uc
}

// SetCity sets the "city" edge to the City entity.
func (uc *UserCreate) SetCity(c *City) *UserCreate {
	return uc.SetCityID(c.ID)
}

// SetCateID sets the "cate" edge to the ItemCategory entity by ID.
func (uc *UserCreate) SetCateID(id int) *UserCreate {
	uc.mutation.SetCateID(id)
	return uc
}

// SetNillableCateID sets the "cate" edge to the ItemCategory entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableCateID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetCateID(*id)
	}
	return uc
}

// SetCate sets the "cate" edge to the ItemCategory entity.
func (uc *UserCreate) SetCate(i *ItemCategory) *UserCreate {
	return uc.SetCateID(i.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.BossUserID(); !ok {
		v := user.DefaultBossUserID
		uc.mutation.SetBossUserID(v)
	}
	if _, ok := uc.mutation.Email(); !ok {
		v := user.DefaultEmail
		uc.mutation.SetEmail(v)
	}
	if _, ok := uc.mutation.Phone(); !ok {
		v := user.DefaultPhone
		uc.mutation.SetPhone(v)
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		v := user.DefaultNickname
		uc.mutation.SetNickname(v)
	}
	if _, ok := uc.mutation.Username(); !ok {
		v := user.DefaultUsername
		uc.mutation.SetUsername(v)
	}
	if _, ok := uc.mutation.Status(); !ok {
		v := user.DefaultStatus
		uc.mutation.SetStatus(v)
	}
	if _, ok := uc.mutation.Sex(); !ok {
		v := user.DefaultSex
		uc.mutation.SetSex(v)
	}
	if _, ok := uc.mutation.RegFrom(); !ok {
		v := user.DefaultRegFrom
		uc.mutation.SetRegFrom(v)
	}
	if _, ok := uc.mutation.CardType(); !ok {
		v := user.DefaultCardType
		uc.mutation.SetCardType(v)
	}
	if _, ok := uc.mutation.IDCard(); !ok {
		v := user.DefaultIDCard
		uc.mutation.SetIDCard(v)
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		v := user.DefaultAvatar
		uc.mutation.SetAvatar(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if _, ok := uc.mutation.Salt(); !ok {
		return &ValidationError{Name: "salt", err: errors.New("ent: missing required field \"salt\"")}
	}
	if _, ok := uc.mutation.BossUserID(); !ok {
		return &ValidationError{Name: "boss_user_id", err: errors.New("ent: missing required field \"boss_user_id\"")}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if _, ok := uc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New("ent: missing required field \"phone\"")}
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New("ent: missing required field \"nickname\"")}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New("ent: missing required field \"username\"")}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := uc.mutation.Sex(); !ok {
		return &ValidationError{Name: "sex", err: errors.New("ent: missing required field \"sex\"")}
	}
	if _, ok := uc.mutation.RegFrom(); !ok {
		return &ValidationError{Name: "reg_from", err: errors.New("ent: missing required field \"reg_from\"")}
	}
	if _, ok := uc.mutation.CardType(); !ok {
		return &ValidationError{Name: "card_type", err: errors.New("ent: missing required field \"card_type\"")}
	}
	if _, ok := uc.mutation.IDCard(); !ok {
		return &ValidationError{Name: "id_card", err: errors.New("ent: missing required field \"id_card\"")}
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New("ent: missing required field \"avatar\"")}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := uc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := uc.mutation.Salt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSalt,
		})
		_node.Salt = value
	}
	if value, ok := uc.mutation.BossUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldBossUserID,
		})
		_node.BossUserID = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := uc.mutation.Nickname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
		_node.Nickname = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: user.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := uc.mutation.Sex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: user.FieldSex,
		})
		_node.Sex = value
	}
	if value, ok := uc.mutation.RegFrom(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: user.FieldRegFrom,
		})
		_node.RegFrom = value
	}
	if value, ok := uc.mutation.CardType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: user.FieldCardType,
		})
		_node.CardType = value
	}
	if value, ok := uc.mutation.IDCard(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldIDCard,
		})
		_node.IDCard = value
	}
	if value, ok := uc.mutation.Birthday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldBirthday,
		})
		_node.Birthday = &value
	}
	if value, ok := uc.mutation.SignRemark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSignRemark,
		})
		_node.SignRemark = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatar,
		})
		_node.Avatar = value
	}
	if nodes := uc.mutation.LoginLogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginLogTable,
			Columns: []string{user.LoginLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userloginlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.MessagesTable,
			Columns: user.MessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserCoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserCoursesTable,
			Columns: []string{user.UserCoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kcusercourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserClassesTable,
			Columns: []string{user.UserClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kcuserclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserExamsRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserExamsRecordsTable,
			Columns: []string{user.UserExamsRecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tkuserexamscorerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserQuestionBankRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserQuestionBankRecordsTable,
			Columns: []string{user.UserQuestionBankRecordsColumn},
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
	if nodes := uc.mutation.UserQuestionRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserQuestionRecordsTable,
			Columns: []string{user.UserQuestionRecordsColumn},
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
	if nodes := uc.mutation.AskUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AskUsersTable,
			Columns: []string{user.AskUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useraskanswer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CourseAppraiseUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CourseAppraiseUsersTable,
			Columns: []string{user.CourseAppraiseUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usercourseappraise.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserVideoRecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserVideoRecordTable,
			Columns: []string{user.UserVideoRecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: videorecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CityTable,
			Columns: []string{user.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FromCityID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CateTable,
			Columns: []string{user.CateColumn},
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
		_node.FromItemCategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
