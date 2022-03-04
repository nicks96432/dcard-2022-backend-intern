package storage

import (
	"context"
	"time"
)

type Storage interface {
	Add(ctx context.Context, e *Entry) error
	Delete(ctx context.Context, id int64) (bool, error)
	Query(ctx context.Context, id int64) (*Entry, error)
}

type Entry struct {
	Id       int64     `json:"id"`
	Url      string    `json:"url"`
	ExpireAt time.Time `json:"-"`
}
