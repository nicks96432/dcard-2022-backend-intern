package sqlite

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

import (
	"dcard-2022-backend-intern/internal/config"
	"dcard-2022-backend-intern/internal/storage"
)

type SqliteStorage struct{ db *sql.DB }

func New(c *config.Config) (storage.Storage, error) {
	var db *sql.DB
	var err error

	if c.InMemory {
		db, err = sql.Open("sqlite3", "file::memory:?cache=shared")
	} else {
		db, err = sql.Open("sqlite3", "./urls.db")
	}
	if err != nil {
		return nil, err
	}

	createTable := `CREATE TABLE IF NOT EXISTS urls (
		id          BIGINT NOT NULL,
		url         TEXT NOT NULL,
		expireAt    TIMESTAMP NOT NULL,
		PRIMARY KEY (id)
	)`

	if _, err = db.Exec(createTable); err != nil {
		db.Close()
		return nil, err
	}

	return &SqliteStorage{db}, nil
}

func (store *SqliteStorage) Add(ctx context.Context, e *storage.Entry) error {
	query := `INSERT INTO urls (id, url, expireAt) VALUES (?, ?, ?)`
	_, err := store.db.ExecContext(ctx, query, e.Id, e.Url, e.ExpireAt)
	return err
}

func (store *SqliteStorage) Delete(ctx context.Context, id int64) (bool, error) {
	query := `DELETE FROM urls WHERE id = ?`
	result, err := store.db.ExecContext(ctx, query, id)
	if err != nil {
		return false, err
	}

	a, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return !(a == 0), err
}

func (store *SqliteStorage) Query(ctx context.Context, id int64) (*storage.Entry, error) {
	query := `SELECT * FROM urls WHERE id = ?`
	result := store.db.QueryRowContext(ctx, query, id)

	var e *storage.Entry = &storage.Entry{}
	err := result.Scan(&e.Id, &e.Url, &e.ExpireAt)

	return e, err
}
