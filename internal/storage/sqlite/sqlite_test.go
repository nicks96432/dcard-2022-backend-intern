package sqlite_test

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"
)

import (
	"dcard-2022-backend-intern/internal/config"
	"dcard-2022-backend-intern/internal/storage"
	"dcard-2022-backend-intern/internal/storage/sqlite"
)

func TestInMemory(t *testing.T) {
	db, err := sqlite.New(config.Default())
	if err != nil {
		t.Fatalf("sqlite.New() failed: %v\n", err)
	}
	ctx := context.Background()
	testOperation(t, db, ctx)
}

func TestNotInMemory(t *testing.T) {
	os.Chdir(t.TempDir())
	db, err := sqlite.New(&config.Config{Port: 8763, InMemory: false})
	if err != nil {
		t.Fatalf("sqlite.New() failed: %v\n", err)
	}
	ctx := context.Background()
	testOperation(t, db, ctx)
}

func testOperation(t *testing.T, db storage.Storage, ctx context.Context) {
	expected := &storage.Entry{Id: 48763, Url: "https://localhost:8763", ExpireAt: time.Now().UTC()}

	if err := db.Add(ctx, expected); err != nil {
		t.Errorf("sqlite.SqliteStorage.Add(): %v\n", err)
	}

	got, err := db.Query(ctx, 48763)
	if err != nil {
		t.Errorf("sqlite.SqliteStorage.Query(): %v\n", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf(
			"sqlite.SqliteStorage.Query(): expect %v, got %v\n",
			expected, got,
		)
	}

	if deleted, err := db.Delete(ctx, 48763); !deleted {
		t.Errorf("sqlite.SqliteStorage.Delete(): expect deleted, got false\n")
	} else if err != nil {
		t.Errorf("sqlite.SqliteStorage.Delete(): expect success, got error %v\n", err)
	}

	if deleted, err := db.Delete(ctx, 48763); deleted {
		t.Errorf("sqlite.SqliteStorage.Delete(): expect not deleted, got true\n")
	} else if err != nil {
		t.Errorf("sqlite.SqliteStorage.Delete(): expect success, got error %v\n", err)
	}
}
