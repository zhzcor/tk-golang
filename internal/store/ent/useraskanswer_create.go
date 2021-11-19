// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gserver/internal/store/ent/teacher"
	"gserver/internal/store/ent/user"
	"gserver/internal/store/ent/useraskanswer"
	"gserver/internal/store/ent/useraskanswerattachment"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserAskAnswerCreate is the builder for creating a UserAskAnswer entity.
type UserAskAnswerCreate struct {
	config
	mutation *UserAskAnswerMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (uaac *UserAskAnswerCreate) SetUUID(s string) *UserAskAnswerCreate {
	uaac.mutation.SetUUID(s)
	return uaac
}

// SetCreatedAt sets the "created_at" field.
func (uaac *UserAskAnswerCreate) SetCreatedAt(t time.Time) *UserAskAnswerCreate {
	uaac.mutation.SetCreatedAt(t)
	return uaac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableCreatedAt(t *time.Time) *UserAskAnswerCreate {
	if t != nil {
		uaac.SetCreatedAt(*t)
	}
	return uaac
}

// SetUpdatedAt sets the "updated_at" field.
func (uaac *UserAskAnswerCreate) SetUpdatedAt(t time.Time) *UserAskAnswerCreate {
	uaac.mutation.SetUpdatedAt(t)
	return uaac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableUpdatedAt(t *time.Time) *UserAskAnswerCreate {
	if t != nil {
		uaac.SetUpdatedAt(*t)
	}
	return uaac
}

// SetDeletedAt sets the "deleted_at" field.
func (uaac *UserAskAnswerCreate) SetDeletedAt(t time.Time) *UserAskAnswerCreate {
	uaac.mutation.SetDeletedAt(t)
	return uaac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableDeletedAt(t *time.Time) *UserAskAnswerCreate {
	if t != nil {
		uaac.SetDeletedAt(*t)
	}
	return uaac
}

// SetAskDesc sets the "ask_desc" field.
func (uaac *UserAskAnswerCreate) SetAskDesc(s string) *UserAskAnswerCreate {
	uaac.mutation.SetAskDesc(s)
	return uaac
}

// SetNillableAskDesc sets the "ask_desc" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableAskDesc(s *string) *UserAskAnswerCreate {
	if s != nil {
		uaac.SetAskDesc(*s)
	}
	return uaac
}

// SetAnswerStatus sets the "answer_status" field.
func (uaac *UserAskAnswerCreate) SetAnswerStatus(u uint8) *UserAskAnswerCreate {
	uaac.mutation.SetAnswerStatus(u)
	return uaac
}

// SetNillableAnswerStatus sets the "answer_status" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableAnswerStatus(u *uint8) *UserAskAnswerCreate {
	if u != nil {
		uaac.SetAnswerStatus(*u)
	}
	return uaac
}

// SetUserID sets the "user_id" field.
func (uaac *UserAskAnswerCreate) SetUserID(i int) *UserAskAnswerCreate {
	uaac.mutation.SetUserID(i)
	return uaac
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableUserID(i *int) *UserAskAnswerCreate {
	if i != nil {
		uaac.SetUserID(*i)
	}
	return uaac
}

// SetTeacherID sets the "teacher_id" field.
func (uaac *UserAskAnswerCreate) SetTeacherID(i int) *UserAskAnswerCreate {
	uaac.mutation.SetTeacherID(i)
	return uaac
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableTeacherID(i *int) *UserAskAnswerCreate {
	if i != nil {
		uaac.SetTeacherID(*i)
	}
	return uaac
}

// SetShowStatus sets the "show_status" field.
func (uaac *UserAskAnswerCreate) SetShowStatus(u uint8) *UserAskAnswerCreate {
	uaac.mutation.SetShowStatus(u)
	return uaac
}

// SetNillableShowStatus sets the "show_status" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableShowStatus(u *uint8) *UserAskAnswerCreate {
	if u != nil {
		uaac.SetShowStatus(*u)
	}
	return uaac
}

// SetReplyShowStatus sets the "reply_show_status" field.
func (uaac *UserAskAnswerCreate) SetReplyShowStatus(u uint8) *UserAskAnswerCreate {
	uaac.mutation.SetReplyShowStatus(u)
	return uaac
}

// SetNillableReplyShowStatus sets the "reply_show_status" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableReplyShowStatus(u *uint8) *UserAskAnswerCreate {
	if u != nil {
		uaac.SetReplyShowStatus(*u)
	}
	return uaac
}

// SetAnswerDesc sets the "answer_desc" field.
func (uaac *UserAskAnswerCreate) SetAnswerDesc(s string) *UserAskAnswerCreate {
	uaac.mutation.SetAnswerDesc(s)
	return uaac
}

// SetNillableAnswerDesc sets the "answer_desc" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableAnswerDesc(s *string) *UserAskAnswerCreate {
	if s != nil {
		uaac.SetAnswerDesc(*s)
	}
	return uaac
}

// SetAnswerAt sets the "answer_at" field.
func (uaac *UserAskAnswerCreate) SetAnswerAt(t time.Time) *UserAskAnswerCreate {
	uaac.mutation.SetAnswerAt(t)
	return uaac
}

// SetNillableAnswerAt sets the "answer_at" field if the given value is not nil.
func (uaac *UserAskAnswerCreate) SetNillableAnswerAt(t *time.Time) *UserAskAnswerCreate {
	if t != nil {
		uaac.SetAnswerAt(*t)
	}
	return uaac
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (uaac *UserAskAnswerCreate) SetTeacher(t *Teacher) *UserAskAnswerCreate {
	return uaac.SetTeacherID(t.ID)
}

// SetUser sets the "user" edge to the User entity.
func (uaac *UserAskAnswerCreate) SetUser(u *User) *UserAskAnswerCreate {
	return uaac.SetUserID(u.ID)
}

// AddAskAnswersAttachmentIDs adds the "ask_answers_attachments" edge to the UserAskAnswerAttachment entity by IDs.
func (uaac *UserAskAnswerCreate) AddAskAnswersAttachmentIDs(ids ...int) *UserAskAnswerCreate {
	uaac.mutation.AddAskAnswersAttachmentIDs(ids...)
	return uaac
}

// AddAskAnswersAttachments adds the "ask_answers_attachments" edges to the UserAskAnswerAttachment entity.
func (uaac *UserAskAnswerCreate) AddAskAnswersAttachments(u ...*UserAskAnswerAttachment) *UserAskAnswerCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uaac.AddAskAnswersAttachmentIDs(ids...)
}

// Mutation returns the UserAskAnswerMutation object of the builder.
func (uaac *UserAskAnswerCreate) Mutation() *UserAskAnswerMutation {
	return uaac.mutation
}

// Save creates the UserAskAnswer in the database.
func (uaac *UserAskAnswerCreate) Save(ctx context.Context) (*UserAskAnswer, error) {
	var (
		err  error
		node *UserAskAnswer
	)
	uaac.defaults()
	if len(uaac.hooks) == 0 {
		if err = uaac.check(); err != nil {
			return nil, err
		}
		node, err = uaac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserAskAnswerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uaac.check(); err != nil {
				return nil, err
			}
			uaac.mutation = mutation
			node, err = uaac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uaac.hooks) - 1; i >= 0; i-- {
			mut = uaac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uaac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uaac *UserAskAnswerCreate) SaveX(ctx context.Context) *UserAskAnswer {
	v, err := uaac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uaac *UserAskAnswerCreate) defaults() {
	if _, ok := uaac.mutation.CreatedAt(); !ok {
		v := useraskanswer.DefaultCreatedAt()
		uaac.mutation.SetCreatedAt(v)
	}
	if _, ok := uaac.mutation.UpdatedAt(); !ok {
		v := useraskanswer.DefaultUpdatedAt()
		uaac.mutation.SetUpdatedAt(v)
	}
	if _, ok := uaac.mutation.AskDesc(); !ok {
		v := useraskanswer.DefaultAskDesc
		uaac.mutation.SetAskDesc(v)
	}
	if _, ok := uaac.mutation.AnswerStatus(); !ok {
		v := useraskanswer.DefaultAnswerStatus
		uaac.mutation.SetAnswerStatus(v)
	}
	if _, ok := uaac.mutation.ShowStatus(); !ok {
		v := useraskanswer.DefaultShowStatus
		uaac.mutation.SetShowStatus(v)
	}
	if _, ok := uaac.mutation.ReplyShowStatus(); !ok {
		v := useraskanswer.DefaultReplyShowStatus
		uaac.mutation.SetReplyShowStatus(v)
	}
	if _, ok := uaac.mutation.AnswerDesc(); !ok {
		v := useraskanswer.DefaultAnswerDesc
		uaac.mutation.SetAnswerDesc(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uaac *UserAskAnswerCreate) check() error {
	if _, ok := uaac.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := uaac.mutation.AskDesc(); !ok {
		return &ValidationError{Name: "ask_desc", err: errors.New("ent: missing required field \"ask_desc\"")}
	}
	if _, ok := uaac.mutation.AnswerStatus(); !ok {
		return &ValidationError{Name: "answer_status", err: errors.New("ent: missing required field \"answer_status\"")}
	}
	if _, ok := uaac.mutation.ShowStatus(); !ok {
		return &ValidationError{Name: "show_status", err: errors.New("ent: missing required field \"show_status\"")}
	}
	if _, ok := uaac.mutation.ReplyShowStatus(); !ok {
		return &ValidationError{Name: "reply_show_status", err: errors.New("ent: missing required field \"reply_show_status\"")}
	}
	if _, ok := uaac.mutation.AnswerDesc(); !ok {
		return &ValidationError{Name: "answer_desc", err: errors.New("ent: missing required field \"answer_desc\"")}
	}
	return nil
}

func (uaac *UserAskAnswerCreate) sqlSave(ctx context.Context) (*UserAskAnswer, error) {
	_node, _spec := uaac.createSpec()
	if err := sqlgraph.CreateNode(ctx, uaac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uaac *UserAskAnswerCreate) createSpec() (*UserAskAnswer, *sqlgraph.CreateSpec) {
	var (
		_node = &UserAskAnswer{config: uaac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: useraskanswer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: useraskanswer.FieldID,
			},
		}
	)
	if value, ok := uaac.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useraskanswer.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := uaac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useraskanswer.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := uaac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useraskanswer.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := uaac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useraskanswer.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := uaac.mutation.AskDesc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useraskanswer.FieldAskDesc,
		})
		_node.AskDesc = value
	}
	if value, ok := uaac.mutation.AnswerStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: useraskanswer.FieldAnswerStatus,
		})
		_node.AnswerStatus = value
	}
	if value, ok := uaac.mutation.ShowStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: useraskanswer.FieldShowStatus,
		})
		_node.ShowStatus = value
	}
	if value, ok := uaac.mutation.ReplyShowStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: useraskanswer.FieldReplyShowStatus,
		})
		_node.ReplyShowStatus = value
	}
	if value, ok := uaac.mutation.AnswerDesc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useraskanswer.FieldAnswerDesc,
		})
		_node.AnswerDesc = value
	}
	if value, ok := uaac.mutation.AnswerAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useraskanswer.FieldAnswerAt,
		})
		_node.AnswerAt = value
	}
	if nodes := uaac.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useraskanswer.TeacherTable,
			Columns: []string{useraskanswer.TeacherColumn},
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
		_node.TeacherID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uaac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useraskanswer.UserTable,
			Columns: []string{useraskanswer.UserColumn},
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
	if nodes := uaac.mutation.AskAnswersAttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   useraskanswer.AskAnswersAttachmentsTable,
			Columns: []string{useraskanswer.AskAnswersAttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useraskanswerattachment.FieldID,
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

// UserAskAnswerCreateBulk is the builder for creating many UserAskAnswer entities in bulk.
type UserAskAnswerCreateBulk struct {
	config
	builders []*UserAskAnswerCreate
}

// Save creates the UserAskAnswer entities in the database.
func (uaacb *UserAskAnswerCreateBulk) Save(ctx context.Context) ([]*UserAskAnswer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uaacb.builders))
	nodes := make([]*UserAskAnswer, len(uaacb.builders))
	mutators := make([]Mutator, len(uaacb.builders))
	for i := range uaacb.builders {
		func(i int, root context.Context) {
			builder := uaacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserAskAnswerMutation)
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
					_, err = mutators[i+1].Mutate(root, uaacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uaacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uaacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uaacb *UserAskAnswerCreateBulk) SaveX(ctx context.Context) []*UserAskAnswer {
	v, err := uaacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
