// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"tkserver/internal/store/ent/kccoursevideo"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KcCourseVideoCreate is the builder for creating a KcCourseVideo entity.
type KcCourseVideoCreate struct {
	config
	mutation *KcCourseVideoMutation
	hooks    []Hook
}

// SetVideoTitle sets the "video_title" field.
func (kcvc *KcCourseVideoCreate) SetVideoTitle(s string) *KcCourseVideoCreate {
	kcvc.mutation.SetVideoTitle(s)
	return kcvc
}

// SetNillableVideoTitle sets the "video_title" field if the given value is not nil.
func (kcvc *KcCourseVideoCreate) SetNillableVideoTitle(s *string) *KcCourseVideoCreate {
	if s != nil {
		kcvc.SetVideoTitle(*s)
	}
	return kcvc
}

// SetFileType sets the "file_type" field.
func (kcvc *KcCourseVideoCreate) SetFileType(i int8) *KcCourseVideoCreate {
	kcvc.mutation.SetFileType(i)
	return kcvc
}

// SetNillableFileType sets the "file_type" field if the given value is not nil.
func (kcvc *KcCourseVideoCreate) SetNillableFileType(i *int8) *KcCourseVideoCreate {
	if i != nil {
		kcvc.SetFileType(*i)
	}
	return kcvc
}

// SetFilePath sets the "file_path" field.
func (kcvc *KcCourseVideoCreate) SetFilePath(s string) *KcCourseVideoCreate {
	kcvc.mutation.SetFilePath(s)
	return kcvc
}

// SetNillableFilePath sets the "file_path" field if the given value is not nil.
func (kcvc *KcCourseVideoCreate) SetNillableFilePath(s *string) *KcCourseVideoCreate {
	if s != nil {
		kcvc.SetFilePath(*s)
	}
	return kcvc
}

// SetSubtitlePath sets the "subtitle_path" field.
func (kcvc *KcCourseVideoCreate) SetSubtitlePath(s string) *KcCourseVideoCreate {
	kcvc.mutation.SetSubtitlePath(s)
	return kcvc
}

// SetNillableSubtitlePath sets the "subtitle_path" field if the given value is not nil.
func (kcvc *KcCourseVideoCreate) SetNillableSubtitlePath(s *string) *KcCourseVideoCreate {
	if s != nil {
		kcvc.SetSubtitlePath(*s)
	}
	return kcvc
}

// SetVideoTime sets the "video_time" field.
func (kcvc *KcCourseVideoCreate) SetVideoTime(s string) *KcCourseVideoCreate {
	kcvc.mutation.SetVideoTime(s)
	return kcvc
}

// SetNillableVideoTime sets the "video_time" field if the given value is not nil.
func (kcvc *KcCourseVideoCreate) SetNillableVideoTime(s *string) *KcCourseVideoCreate {
	if s != nil {
		kcvc.SetVideoTime(*s)
	}
	return kcvc
}

// Mutation returns the KcCourseVideoMutation object of the builder.
func (kcvc *KcCourseVideoCreate) Mutation() *KcCourseVideoMutation {
	return kcvc.mutation
}

// Save creates the KcCourseVideo in the database.
func (kcvc *KcCourseVideoCreate) Save(ctx context.Context) (*KcCourseVideo, error) {
	var (
		err  error
		node *KcCourseVideo
	)
	kcvc.defaults()
	if len(kcvc.hooks) == 0 {
		if err = kcvc.check(); err != nil {
			return nil, err
		}
		node, err = kcvc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KcCourseVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kcvc.check(); err != nil {
				return nil, err
			}
			kcvc.mutation = mutation
			node, err = kcvc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kcvc.hooks) - 1; i >= 0; i-- {
			mut = kcvc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kcvc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kcvc *KcCourseVideoCreate) SaveX(ctx context.Context) *KcCourseVideo {
	v, err := kcvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (kcvc *KcCourseVideoCreate) defaults() {
	if _, ok := kcvc.mutation.VideoTitle(); !ok {
		v := kccoursevideo.DefaultVideoTitle
		kcvc.mutation.SetVideoTitle(v)
	}
	if _, ok := kcvc.mutation.FileType(); !ok {
		v := kccoursevideo.DefaultFileType
		kcvc.mutation.SetFileType(v)
	}
	if _, ok := kcvc.mutation.FilePath(); !ok {
		v := kccoursevideo.DefaultFilePath
		kcvc.mutation.SetFilePath(v)
	}
	if _, ok := kcvc.mutation.SubtitlePath(); !ok {
		v := kccoursevideo.DefaultSubtitlePath
		kcvc.mutation.SetSubtitlePath(v)
	}
	if _, ok := kcvc.mutation.VideoTime(); !ok {
		v := kccoursevideo.DefaultVideoTime
		kcvc.mutation.SetVideoTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kcvc *KcCourseVideoCreate) check() error {
	if _, ok := kcvc.mutation.VideoTitle(); !ok {
		return &ValidationError{Name: "video_title", err: errors.New("ent: missing required field \"video_title\"")}
	}
	if _, ok := kcvc.mutation.FileType(); !ok {
		return &ValidationError{Name: "file_type", err: errors.New("ent: missing required field \"file_type\"")}
	}
	if _, ok := kcvc.mutation.FilePath(); !ok {
		return &ValidationError{Name: "file_path", err: errors.New("ent: missing required field \"file_path\"")}
	}
	if _, ok := kcvc.mutation.SubtitlePath(); !ok {
		return &ValidationError{Name: "subtitle_path", err: errors.New("ent: missing required field \"subtitle_path\"")}
	}
	if _, ok := kcvc.mutation.VideoTime(); !ok {
		return &ValidationError{Name: "video_time", err: errors.New("ent: missing required field \"video_time\"")}
	}
	return nil
}

func (kcvc *KcCourseVideoCreate) sqlSave(ctx context.Context) (*KcCourseVideo, error) {
	_node, _spec := kcvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kcvc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (kcvc *KcCourseVideoCreate) createSpec() (*KcCourseVideo, *sqlgraph.CreateSpec) {
	var (
		_node = &KcCourseVideo{config: kcvc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kccoursevideo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kccoursevideo.FieldID,
			},
		}
	)
	if value, ok := kcvc.mutation.VideoTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kccoursevideo.FieldVideoTitle,
		})
		_node.VideoTitle = value
	}
	if value, ok := kcvc.mutation.FileType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: kccoursevideo.FieldFileType,
		})
		_node.FileType = value
	}
	if value, ok := kcvc.mutation.FilePath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kccoursevideo.FieldFilePath,
		})
		_node.FilePath = value
	}
	if value, ok := kcvc.mutation.SubtitlePath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kccoursevideo.FieldSubtitlePath,
		})
		_node.SubtitlePath = value
	}
	if value, ok := kcvc.mutation.VideoTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kccoursevideo.FieldVideoTime,
		})
		_node.VideoTime = value
	}
	return _node, _spec
}

// KcCourseVideoCreateBulk is the builder for creating many KcCourseVideo entities in bulk.
type KcCourseVideoCreateBulk struct {
	config
	builders []*KcCourseVideoCreate
}

// Save creates the KcCourseVideo entities in the database.
func (kcvcb *KcCourseVideoCreateBulk) Save(ctx context.Context) ([]*KcCourseVideo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kcvcb.builders))
	nodes := make([]*KcCourseVideo, len(kcvcb.builders))
	mutators := make([]Mutator, len(kcvcb.builders))
	for i := range kcvcb.builders {
		func(i int, root context.Context) {
			builder := kcvcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KcCourseVideoMutation)
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
					_, err = mutators[i+1].Mutate(root, kcvcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcvcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kcvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcvcb *KcCourseVideoCreateBulk) SaveX(ctx context.Context) []*KcCourseVideo {
	v, err := kcvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
