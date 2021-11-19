package store

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestOpenDatabase(t *testing.T) {
	ctx := context.Background()
	ctx, err := NewContext(ctx, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Fatal(err)
	}
	c := WithContext(ctx)
	if c == nil {
		t.Fatal(fmt.Errorf("not database opened"))
	}
}

func TestCloseDatabase(t *testing.T) {
	ctx := context.Background()
	ctx, err := NewContext(ctx, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Fatal(err)
	}
	err = CloseDatabase(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMigration(t *testing.T) {
	ctx := context.Background()
	ctx, err := NewContext(ctx, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Fatal(err)
	}
	err = Migration(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
