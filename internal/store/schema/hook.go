package schema

import (
	"tkserver/pkg/password"
	strings "tkserver/pkg/strings"

	"context"
	"entgo.io/ent"
	"entgo.io/ent/entc/integration/ent/hook"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func (PasswordMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(mutator ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
				if p, exists := mutation.Field("password"); exists {
					var salt string
					if s, e := mutation.Field("salt"); e {
						salt = s.(string)
					} else {
						salt = strings.Rand(8)
					}
					pwd := p.(string)
					mutation.SetField("salt", salt)
					mutation.SetField("password", password.Gen(pwd+salt))
				}
				return mutator.Mutate(ctx, mutation)
			})
		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}

func (OrderMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(mutator ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
				// 自动创建一个order_id

				if _, exists := mutation.Field("order_id"); !exists {
					mutation.SetField("order_id", mutation.Type()+strings.Rand(6)+strconv.Itoa(int(time.Now().Unix())))
				}
				return mutator.Mutate(ctx, mutation)
			})
		}, ent.OpCreate),
		hook.On(func(mutator ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
				// 自动创建一个uuid
				if v, exists := mutation.Field("version"); !exists {
					if vv, o := v.(int); o {
						mutation.SetField("version", vv+1)
					}

				}
				return mutator.Mutate(ctx, mutation)
			})
		}, ent.OpUpdateOne),
		VersionHook(),
	}
}

func (CommonMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(mutator ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
				// 自动创建一个uuid
				if _, exists := mutation.Field("uuid"); !exists {
					mutation.SetField("uuid", uuid.New().String())
				}
				return mutator.Mutate(ctx, mutation)
			})
		}, ent.OpCreate),
	}
}

// A VersionHook is a dummy example for a hook that validates the "version" field
// is incremented by 1 on each update. Note that this is just a dummy example, and
// it doesn't promise consistency in the database.
func VersionHook() ent.Hook {
	type OldSetVersion interface {
		SetVersion(int)
		Version() (int, bool)
		OldVersion(context.Context) (int, error)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			ver, ok := m.(OldSetVersion)
			if !ok {
				return next.Mutate(ctx, m)
			}
			if m.Op().Is(ent.OpUpdateOne) {
				oldV, err := ver.OldVersion(ctx)
				if err != nil {
					return nil, err
				}
				curV, exists := ver.Version()
				if !exists {
					return nil, fmt.Errorf("version field is required in update mutation")
				}
				if curV != oldV+1 {
					return nil, fmt.Errorf("version field must be incremented by 1")
				}
				// Add an SQL predicate that validates the "version" column is equal
				// to "oldV" (ensure it wasn't changed during the mutation by others).
			}

			return next.Mutate(ctx, m)
		})
	}
}
