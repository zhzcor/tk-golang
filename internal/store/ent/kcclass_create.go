// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tkserver/internal/store/ent/admin"
	"tkserver/internal/store/ent/attachment"
	"tkserver/internal/store/ent/city"
	"tkserver/internal/store/ent/itemcategory"
	"tkserver/internal/store/ent/kcclass"
	"tkserver/internal/store/ent/kcclassteacher"
	"tkserver/internal/store/ent/kccourse"
	"tkserver/internal/store/ent/kcuserclass"
	"tkserver/internal/store/ent/major"
	"tkserver/internal/store/ent/message"
	"tkserver/internal/store/ent/teacher"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KcClassCreate is the builder for creating a KcClass entity.
type KcClassCreate struct {
	config
	mutation *KcClassMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (kcc *KcClassCreate) SetUUID(s string) *KcClassCreate {
	kcc.mutation.SetUUID(s)
	return kcc
}

// SetCreatedAt sets the "created_at" field.
func (kcc *KcClassCreate) SetCreatedAt(t time.Time) *KcClassCreate {
	kcc.mutation.SetCreatedAt(t)
	return kcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableCreatedAt(t *time.Time) *KcClassCreate {
	if t != nil {
		kcc.SetCreatedAt(*t)
	}
	return kcc
}

// SetUpdatedAt sets the "updated_at" field.
func (kcc *KcClassCreate) SetUpdatedAt(t time.Time) *KcClassCreate {
	kcc.mutation.SetUpdatedAt(t)
	return kcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableUpdatedAt(t *time.Time) *KcClassCreate {
	if t != nil {
		kcc.SetUpdatedAt(*t)
	}
	return kcc
}

// SetDeletedAt sets the "deleted_at" field.
func (kcc *KcClassCreate) SetDeletedAt(t time.Time) *KcClassCreate {
	kcc.mutation.SetDeletedAt(t)
	return kcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableDeletedAt(t *time.Time) *KcClassCreate {
	if t != nil {
		kcc.SetDeletedAt(*t)
	}
	return kcc
}

// SetClassTitle sets the "class_title" field.
func (kcc *KcClassCreate) SetClassTitle(s string) *KcClassCreate {
	kcc.mutation.SetClassTitle(s)
	return kcc
}

// SetNillableClassTitle sets the "class_title" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableClassTitle(s *string) *KcClassCreate {
	if s != nil {
		kcc.SetClassTitle(*s)
	}
	return kcc
}

// SetClassCode sets the "class_code" field.
func (kcc *KcClassCreate) SetClassCode(s string) *KcClassCreate {
	kcc.mutation.SetClassCode(s)
	return kcc
}

// SetNillableClassCode sets the "class_code" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableClassCode(s *string) *KcClassCreate {
	if s != nil {
		kcc.SetClassCode(*s)
	}
	return kcc
}

// SetClassDesc sets the "class_desc" field.
func (kcc *KcClassCreate) SetClassDesc(s string) *KcClassCreate {
	kcc.mutation.SetClassDesc(s)
	return kcc
}

// SetNillableClassDesc sets the "class_desc" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableClassDesc(s *string) *KcClassCreate {
	if s != nil {
		kcc.SetClassDesc(*s)
	}
	return kcc
}

// SetIsDisplay sets the "is_display" field.
func (kcc *KcClassCreate) SetIsDisplay(u uint8) *KcClassCreate {
	kcc.mutation.SetIsDisplay(u)
	return kcc
}

// SetNillableIsDisplay sets the "is_display" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableIsDisplay(u *uint8) *KcClassCreate {
	if u != nil {
		kcc.SetIsDisplay(*u)
	}
	return kcc
}

// SetIsBuy sets the "is_buy" field.
func (kcc *KcClassCreate) SetIsBuy(u uint8) *KcClassCreate {
	kcc.mutation.SetIsBuy(u)
	return kcc
}

// SetNillableIsBuy sets the "is_buy" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableIsBuy(u *uint8) *KcClassCreate {
	if u != nil {
		kcc.SetIsBuy(*u)
	}
	return kcc
}

// SetClassPeriodType sets the "class_period_type" field.
func (kcc *KcClassCreate) SetClassPeriodType(u uint8) *KcClassCreate {
	kcc.mutation.SetClassPeriodType(u)
	return kcc
}

// SetNillableClassPeriodType sets the "class_period_type" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableClassPeriodType(u *uint8) *KcClassCreate {
	if u != nil {
		kcc.SetClassPeriodType(*u)
	}
	return kcc
}

// SetClassStartDate sets the "class_start_date" field.
func (kcc *KcClassCreate) SetClassStartDate(t time.Time) *KcClassCreate {
	kcc.mutation.SetClassStartDate(t)
	return kcc
}

// SetNillableClassStartDate sets the "class_start_date" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableClassStartDate(t *time.Time) *KcClassCreate {
	if t != nil {
		kcc.SetClassStartDate(*t)
	}
	return kcc
}

// SetClassEndDate sets the "class_end_date" field.
func (kcc *KcClassCreate) SetClassEndDate(t time.Time) *KcClassCreate {
	kcc.mutation.SetClassEndDate(t)
	return kcc
}

// SetNillableClassEndDate sets the "class_end_date" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableClassEndDate(t *time.Time) *KcClassCreate {
	if t != nil {
		kcc.SetClassEndDate(*t)
	}
	return kcc
}

// SetClosingDate sets the "closing_date" field.
func (kcc *KcClassCreate) SetClosingDate(t time.Time) *KcClassCreate {
	kcc.mutation.SetClosingDate(t)
	return kcc
}

// SetNillableClosingDate sets the "closing_date" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableClosingDate(t *time.Time) *KcClassCreate {
	if t != nil {
		kcc.SetClosingDate(*t)
	}
	return kcc
}

// SetDaysValidity sets the "days_validity" field.
func (kcc *KcClassCreate) SetDaysValidity(i int) *KcClassCreate {
	kcc.mutation.SetDaysValidity(i)
	return kcc
}

// SetNillableDaysValidity sets the "days_validity" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableDaysValidity(i *int) *KcClassCreate {
	if i != nil {
		kcc.SetDaysValidity(*i)
	}
	return kcc
}

// SetClassHeadMasterID sets the "class_head_master_id" field.
func (kcc *KcClassCreate) SetClassHeadMasterID(i int) *KcClassCreate {
	kcc.mutation.SetClassHeadMasterID(i)
	return kcc
}

// SetNillableClassHeadMasterID sets the "class_head_master_id" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableClassHeadMasterID(i *int) *KcClassCreate {
	if i != nil {
		kcc.SetClassHeadMasterID(*i)
	}
	return kcc
}

// SetClassCoverImgID sets the "class_cover_img_id" field.
func (kcc *KcClassCreate) SetClassCoverImgID(i int) *KcClassCreate {
	kcc.mutation.SetClassCoverImgID(i)
	return kcc
}

// SetNillableClassCoverImgID sets the "class_cover_img_id" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableClassCoverImgID(i *int) *KcClassCreate {
	if i != nil {
		kcc.SetClassCoverImgID(*i)
	}
	return kcc
}

// SetStatus sets the "status" field.
func (kcc *KcClassCreate) SetStatus(u uint8) *KcClassCreate {
	kcc.mutation.SetStatus(u)
	return kcc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableStatus(u *uint8) *KcClassCreate {
	if u != nil {
		kcc.SetStatus(*u)
	}
	return kcc
}

// SetPrice sets the "price" field.
func (kcc *KcClassCreate) SetPrice(f float64) *KcClassCreate {
	kcc.mutation.SetPrice(f)
	return kcc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillablePrice(f *float64) *KcClassCreate {
	if f != nil {
		kcc.SetPrice(*f)
	}
	return kcc
}

// SetStudentCount sets the "student_count" field.
func (kcc *KcClassCreate) SetStudentCount(i int) *KcClassCreate {
	kcc.mutation.SetStudentCount(i)
	return kcc
}

// SetNillableStudentCount sets the "student_count" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableStudentCount(i *int) *KcClassCreate {
	if i != nil {
		kcc.SetStudentCount(*i)
	}
	return kcc
}

// SetCourseCount sets the "course_count" field.
func (kcc *KcClassCreate) SetCourseCount(i int) *KcClassCreate {
	kcc.mutation.SetCourseCount(i)
	return kcc
}

// SetNillableCourseCount sets the "course_count" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableCourseCount(i *int) *KcClassCreate {
	if i != nil {
		kcc.SetCourseCount(*i)
	}
	return kcc
}

// SetCateID sets the "cate_id" field.
func (kcc *KcClassCreate) SetCateID(i int) *KcClassCreate {
	kcc.mutation.SetCateID(i)
	return kcc
}

// SetNillableCateID sets the "cate_id" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableCateID(i *int) *KcClassCreate {
	if i != nil {
		kcc.SetCateID(*i)
	}
	return kcc
}

// SetCityID sets the "city_id" field.
func (kcc *KcClassCreate) SetCityID(i int) *KcClassCreate {
	kcc.mutation.SetCityID(i)
	return kcc
}

// SetNillableCityID sets the "city_id" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableCityID(i *int) *KcClassCreate {
	if i != nil {
		kcc.SetCityID(*i)
	}
	return kcc
}

// SetCreatedAdminID sets the "created_admin_id" field.
func (kcc *KcClassCreate) SetCreatedAdminID(i int) *KcClassCreate {
	kcc.mutation.SetCreatedAdminID(i)
	return kcc
}

// SetNillableCreatedAdminID sets the "created_admin_id" field if the given value is not nil.
func (kcc *KcClassCreate) SetNillableCreatedAdminID(i *int) *KcClassCreate {
	if i != nil {
		kcc.SetCreatedAdminID(*i)
	}
	return kcc
}

// AddMajorIDs adds the "majors" edge to the Major entity by IDs.
func (kcc *KcClassCreate) AddMajorIDs(ids ...int) *KcClassCreate {
	kcc.mutation.AddMajorIDs(ids...)
	return kcc
}

// AddMajors adds the "majors" edges to the Major entity.
func (kcc *KcClassCreate) AddMajors(m ...*Major) *KcClassCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return kcc.AddMajorIDs(ids...)
}

// SetItemID sets the "item" edge to the ItemCategory entity by ID.
func (kcc *KcClassCreate) SetItemID(id int) *KcClassCreate {
	kcc.mutation.SetItemID(id)
	return kcc
}

// SetNillableItemID sets the "item" edge to the ItemCategory entity by ID if the given value is not nil.
func (kcc *KcClassCreate) SetNillableItemID(id *int) *KcClassCreate {
	if id != nil {
		kcc = kcc.SetItemID(*id)
	}
	return kcc
}

// SetItem sets the "item" edge to the ItemCategory entity.
func (kcc *KcClassCreate) SetItem(i *ItemCategory) *KcClassCreate {
	return kcc.SetItemID(i.ID)
}

// SetCity sets the "city" edge to the City entity.
func (kcc *KcClassCreate) SetCity(c *City) *KcClassCreate {
	return kcc.SetCityID(c.ID)
}

// SetAdminID sets the "admin" edge to the Admin entity by ID.
func (kcc *KcClassCreate) SetAdminID(id int) *KcClassCreate {
	kcc.mutation.SetAdminID(id)
	return kcc
}

// SetNillableAdminID sets the "admin" edge to the Admin entity by ID if the given value is not nil.
func (kcc *KcClassCreate) SetNillableAdminID(id *int) *KcClassCreate {
	if id != nil {
		kcc = kcc.SetAdminID(*id)
	}
	return kcc
}

// SetAdmin sets the "admin" edge to the Admin entity.
func (kcc *KcClassCreate) SetAdmin(a *Admin) *KcClassCreate {
	return kcc.SetAdminID(a.ID)
}

// SetAttachmentID sets the "attachment" edge to the Attachment entity by ID.
func (kcc *KcClassCreate) SetAttachmentID(id int) *KcClassCreate {
	kcc.mutation.SetAttachmentID(id)
	return kcc
}

// SetNillableAttachmentID sets the "attachment" edge to the Attachment entity by ID if the given value is not nil.
func (kcc *KcClassCreate) SetNillableAttachmentID(id *int) *KcClassCreate {
	if id != nil {
		kcc = kcc.SetAttachmentID(*id)
	}
	return kcc
}

// SetAttachment sets the "attachment" edge to the Attachment entity.
func (kcc *KcClassCreate) SetAttachment(a *Attachment) *KcClassCreate {
	return kcc.SetAttachmentID(a.ID)
}

// SetMasterTeachersID sets the "master_teachers" edge to the Teacher entity by ID.
func (kcc *KcClassCreate) SetMasterTeachersID(id int) *KcClassCreate {
	kcc.mutation.SetMasterTeachersID(id)
	return kcc
}

// SetNillableMasterTeachersID sets the "master_teachers" edge to the Teacher entity by ID if the given value is not nil.
func (kcc *KcClassCreate) SetNillableMasterTeachersID(id *int) *KcClassCreate {
	if id != nil {
		kcc = kcc.SetMasterTeachersID(*id)
	}
	return kcc
}

// SetMasterTeachers sets the "master_teachers" edge to the Teacher entity.
func (kcc *KcClassCreate) SetMasterTeachers(t *Teacher) *KcClassCreate {
	return kcc.SetMasterTeachersID(t.ID)
}

// AddClassTeacherIDs adds the "class_teachers" edge to the KcClassTeacher entity by IDs.
func (kcc *KcClassCreate) AddClassTeacherIDs(ids ...int) *KcClassCreate {
	kcc.mutation.AddClassTeacherIDs(ids...)
	return kcc
}

// AddClassTeachers adds the "class_teachers" edges to the KcClassTeacher entity.
func (kcc *KcClassCreate) AddClassTeachers(k ...*KcClassTeacher) *KcClassCreate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kcc.AddClassTeacherIDs(ids...)
}

// AddKcClassCourseIDs adds the "kc_class_courses" edge to the KcCourse entity by IDs.
func (kcc *KcClassCreate) AddKcClassCourseIDs(ids ...int) *KcClassCreate {
	kcc.mutation.AddKcClassCourseIDs(ids...)
	return kcc
}

// AddKcClassCourses adds the "kc_class_courses" edges to the KcCourse entity.
func (kcc *KcClassCreate) AddKcClassCourses(k ...*KcCourse) *KcClassCreate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kcc.AddKcClassCourseIDs(ids...)
}

// AddKcUserClassIDs adds the "kc_user_classes" edge to the KcUserClass entity by IDs.
func (kcc *KcClassCreate) AddKcUserClassIDs(ids ...int) *KcClassCreate {
	kcc.mutation.AddKcUserClassIDs(ids...)
	return kcc
}

// AddKcUserClasses adds the "kc_user_classes" edges to the KcUserClass entity.
func (kcc *KcClassCreate) AddKcUserClasses(k ...*KcUserClass) *KcClassCreate {
	ids := make([]int, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kcc.AddKcUserClassIDs(ids...)
}

// AddMessageClassIDs adds the "message_classes" edge to the Message entity by IDs.
func (kcc *KcClassCreate) AddMessageClassIDs(ids ...int) *KcClassCreate {
	kcc.mutation.AddMessageClassIDs(ids...)
	return kcc
}

// AddMessageClasses adds the "message_classes" edges to the Message entity.
func (kcc *KcClassCreate) AddMessageClasses(m ...*Message) *KcClassCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return kcc.AddMessageClassIDs(ids...)
}

// Mutation returns the KcClassMutation object of the builder.
func (kcc *KcClassCreate) Mutation() *KcClassMutation {
	return kcc.mutation
}

// Save creates the KcClass in the database.
func (kcc *KcClassCreate) Save(ctx context.Context) (*KcClass, error) {
	var (
		err  error
		node *KcClass
	)
	kcc.defaults()
	if len(kcc.hooks) == 0 {
		if err = kcc.check(); err != nil {
			return nil, err
		}
		node, err = kcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KcClassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kcc.check(); err != nil {
				return nil, err
			}
			kcc.mutation = mutation
			node, err = kcc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kcc.hooks) - 1; i >= 0; i-- {
			mut = kcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kcc *KcClassCreate) SaveX(ctx context.Context) *KcClass {
	v, err := kcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (kcc *KcClassCreate) defaults() {
	if _, ok := kcc.mutation.CreatedAt(); !ok {
		v := kcclass.DefaultCreatedAt()
		kcc.mutation.SetCreatedAt(v)
	}
	if _, ok := kcc.mutation.UpdatedAt(); !ok {
		v := kcclass.DefaultUpdatedAt()
		kcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := kcc.mutation.ClassTitle(); !ok {
		v := kcclass.DefaultClassTitle
		kcc.mutation.SetClassTitle(v)
	}
	if _, ok := kcc.mutation.ClassCode(); !ok {
		v := kcclass.DefaultClassCode
		kcc.mutation.SetClassCode(v)
	}
	if _, ok := kcc.mutation.ClassDesc(); !ok {
		v := kcclass.DefaultClassDesc
		kcc.mutation.SetClassDesc(v)
	}
	if _, ok := kcc.mutation.IsDisplay(); !ok {
		v := kcclass.DefaultIsDisplay
		kcc.mutation.SetIsDisplay(v)
	}
	if _, ok := kcc.mutation.IsBuy(); !ok {
		v := kcclass.DefaultIsBuy
		kcc.mutation.SetIsBuy(v)
	}
	if _, ok := kcc.mutation.ClassPeriodType(); !ok {
		v := kcclass.DefaultClassPeriodType
		kcc.mutation.SetClassPeriodType(v)
	}
	if _, ok := kcc.mutation.DaysValidity(); !ok {
		v := kcclass.DefaultDaysValidity
		kcc.mutation.SetDaysValidity(v)
	}
	if _, ok := kcc.mutation.Status(); !ok {
		v := kcclass.DefaultStatus
		kcc.mutation.SetStatus(v)
	}
	if _, ok := kcc.mutation.Price(); !ok {
		v := kcclass.DefaultPrice
		kcc.mutation.SetPrice(v)
	}
	if _, ok := kcc.mutation.StudentCount(); !ok {
		v := kcclass.DefaultStudentCount
		kcc.mutation.SetStudentCount(v)
	}
	if _, ok := kcc.mutation.CourseCount(); !ok {
		v := kcclass.DefaultCourseCount
		kcc.mutation.SetCourseCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kcc *KcClassCreate) check() error {
	if _, ok := kcc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := kcc.mutation.ClassTitle(); !ok {
		return &ValidationError{Name: "class_title", err: errors.New("ent: missing required field \"class_title\"")}
	}
	if _, ok := kcc.mutation.ClassCode(); !ok {
		return &ValidationError{Name: "class_code", err: errors.New("ent: missing required field \"class_code\"")}
	}
	if _, ok := kcc.mutation.ClassDesc(); !ok {
		return &ValidationError{Name: "class_desc", err: errors.New("ent: missing required field \"class_desc\"")}
	}
	if _, ok := kcc.mutation.IsDisplay(); !ok {
		return &ValidationError{Name: "is_display", err: errors.New("ent: missing required field \"is_display\"")}
	}
	if _, ok := kcc.mutation.IsBuy(); !ok {
		return &ValidationError{Name: "is_buy", err: errors.New("ent: missing required field \"is_buy\"")}
	}
	if _, ok := kcc.mutation.ClassPeriodType(); !ok {
		return &ValidationError{Name: "class_period_type", err: errors.New("ent: missing required field \"class_period_type\"")}
	}
	if _, ok := kcc.mutation.DaysValidity(); !ok {
		return &ValidationError{Name: "days_validity", err: errors.New("ent: missing required field \"days_validity\"")}
	}
	if _, ok := kcc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := kcc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if _, ok := kcc.mutation.StudentCount(); !ok {
		return &ValidationError{Name: "student_count", err: errors.New("ent: missing required field \"student_count\"")}
	}
	if _, ok := kcc.mutation.CourseCount(); !ok {
		return &ValidationError{Name: "course_count", err: errors.New("ent: missing required field \"course_count\"")}
	}
	return nil
}

func (kcc *KcClassCreate) sqlSave(ctx context.Context) (*KcClass, error) {
	_node, _spec := kcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kcc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (kcc *KcClassCreate) createSpec() (*KcClass, *sqlgraph.CreateSpec) {
	var (
		_node = &KcClass{config: kcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kcclass.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kcclass.FieldID,
			},
		}
	)
	if value, ok := kcc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcclass.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := kcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcclass.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := kcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcclass.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := kcc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcclass.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := kcc.mutation.ClassTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcclass.FieldClassTitle,
		})
		_node.ClassTitle = value
	}
	if value, ok := kcc.mutation.ClassCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcclass.FieldClassCode,
		})
		_node.ClassCode = value
	}
	if value, ok := kcc.mutation.ClassDesc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kcclass.FieldClassDesc,
		})
		_node.ClassDesc = value
	}
	if value, ok := kcc.mutation.IsDisplay(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcclass.FieldIsDisplay,
		})
		_node.IsDisplay = value
	}
	if value, ok := kcc.mutation.IsBuy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcclass.FieldIsBuy,
		})
		_node.IsBuy = value
	}
	if value, ok := kcc.mutation.ClassPeriodType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcclass.FieldClassPeriodType,
		})
		_node.ClassPeriodType = value
	}
	if value, ok := kcc.mutation.ClassStartDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcclass.FieldClassStartDate,
		})
		_node.ClassStartDate = value
	}
	if value, ok := kcc.mutation.ClassEndDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcclass.FieldClassEndDate,
		})
		_node.ClassEndDate = value
	}
	if value, ok := kcc.mutation.ClosingDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kcclass.FieldClosingDate,
		})
		_node.ClosingDate = &value
	}
	if value, ok := kcc.mutation.DaysValidity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: kcclass.FieldDaysValidity,
		})
		_node.DaysValidity = value
	}
	if value, ok := kcc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: kcclass.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := kcc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: kcclass.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := kcc.mutation.StudentCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: kcclass.FieldStudentCount,
		})
		_node.StudentCount = value
	}
	if value, ok := kcc.mutation.CourseCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: kcclass.FieldCourseCount,
		})
		_node.CourseCount = value
	}
	if nodes := kcc.mutation.MajorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   kcclass.MajorsTable,
			Columns: kcclass.MajorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: major.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kcc.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   kcclass.ItemTable,
			Columns: []string{kcclass.ItemColumn},
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
		_node.CateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kcc.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcclass.CityTable,
			Columns: []string{kcclass.CityColumn},
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
		_node.CityID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kcc.mutation.AdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcclass.AdminTable,
			Columns: []string{kcclass.AdminColumn},
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
	if nodes := kcc.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcclass.AttachmentTable,
			Columns: []string{kcclass.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ClassCoverImgID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kcc.mutation.MasterTeachersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   kcclass.MasterTeachersTable,
			Columns: []string{kcclass.MasterTeachersColumn},
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
		_node.ClassHeadMasterID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kcc.mutation.ClassTeachersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kcclass.ClassTeachersTable,
			Columns: []string{kcclass.ClassTeachersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: kcclassteacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kcc.mutation.KcClassCoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   kcclass.KcClassCoursesTable,
			Columns: kcclass.KcClassCoursesPrimaryKey,
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
	if nodes := kcc.mutation.KcUserClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kcclass.KcUserClassesTable,
			Columns: []string{kcclass.KcUserClassesColumn},
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
	if nodes := kcc.mutation.MessageClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kcclass.MessageClassesTable,
			Columns: []string{kcclass.MessageClassesColumn},
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
	return _node, _spec
}

// KcClassCreateBulk is the builder for creating many KcClass entities in bulk.
type KcClassCreateBulk struct {
	config
	builders []*KcClassCreate
}

// Save creates the KcClass entities in the database.
func (kccb *KcClassCreateBulk) Save(ctx context.Context) ([]*KcClass, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kccb.builders))
	nodes := make([]*KcClass, len(kccb.builders))
	mutators := make([]Mutator, len(kccb.builders))
	for i := range kccb.builders {
		func(i int, root context.Context) {
			builder := kccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KcClassMutation)
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
					_, err = mutators[i+1].Mutate(root, kccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kccb *KcClassCreateBulk) SaveX(ctx context.Context) []*KcClass {
	v, err := kccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
