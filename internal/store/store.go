package store

import (
	"bytes"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/pkg/errors"
	"gserver/internal/store/ent"
	_ "gserver/internal/store/ent/runtime"
	"gserver/pkg/log"
	"io/ioutil"
	"runtime/debug"
)

//go:generate rm -rf ent
//go:generate go run entgen/main.go

var (
	ErrUnsupportedDrive = errors.New("unsupported driver")
	ErrDatabaseNotFound = errors.New("database connection not found")
)

const (
	CtxDatabaseKey   string = "CtxDatabaseKey"
	ctxDatabaseTxKey string = "CtxDatabaseTxKey"
)

var migrationDataHook []func(ctx context.Context) error

// 初始化 migrationDataHook 用来做 migration 时初始化数据
func init() {
	migrationDataHook = append(migrationDataHook, func(ctx context.Context) error {

		//  数据初始化
		return nil
	})
}

func NewContext(ctx context.Context, driverName, dataSourceName string) (context.Context, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		dbgDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
			log.Debug("ent", "d", i)
		})
		c := ent.NewClient(ent.Driver(dbgDrv)).Debug()
		return context.WithValue(ctx, CtxDatabaseKey, c), nil
	default:
		return nil, errors.WithMessage(ErrUnsupportedDrive, "driveName:"+driverName)
	}
}

func NewTxContext(ctx context.Context) (context.Context, error) {
	db := WithContext(ctx)
	if db != nil {
		c, err := db.Tx(ctx)
		if err != nil {
			return nil, err
		}
		return context.WithValue(ctx, ctxDatabaseTxKey, c), nil
	}
	return nil, ErrUnsupportedDrive
}

func WithContext(ctx context.Context) *ent.Client {
	// 要是有事务执行，优先用事务
	if c, ok := ctx.Value(ctxDatabaseTxKey).(*ent.Tx); ok {
		return c.Client()
	}
	if c, ok := ctx.Value(CtxDatabaseKey).(*ent.Client); ok {
		return c
	}
	return nil
}

func WithTxContext(ctx context.Context) *ent.Tx {
	if c, ok := ctx.Value(ctxDatabaseTxKey).(*ent.Tx); ok {
		return c
	}
	return nil
}

func CloseDatabase(ctx context.Context) error {
	if c, ok := ctx.Value(CtxDatabaseKey).(*ent.Client); ok {
		return c.Close()
	}
	return nil
}

func Migration(ctx context.Context) error {
	buffer := bytes.NewBuffer([]byte{})
	err := WithContext(ctx).Schema.WriteTo(ctx, buffer)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadAll(buffer)
	if err != nil {
		return err
	}
	log.Debugf("ent \r\n migrationDatabase: \r\n%s", buf)
	err = WithContext(ctx).Schema.Create(ctx)
	if err != nil {
		return err
	}

	err = WithTx(ctx, func(ctx context.Context) error {
		for _, f := range migrationDataHook {
			err := f(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func WithTx(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	ctx, err = NewTxContext(ctx)
	if err != nil {
		return
	}
	tx := WithTxContext(ctx)
	if tx == nil {
		return ErrDatabaseNotFound
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			// handler panic
			err = fmt.Errorf("transaction panic: %v, stack: %s", v, debug.Stack())
			log.Errorf("transaction panic: %v, stack: %s", v, debug.Stack())
			return
		}
	}()
	if err = fn(ctx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return
}
