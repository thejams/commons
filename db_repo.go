package commons

import (
	"context"
)

type Repository interface {
	// Get gets a specific record from the DB
	Get(string) (interface{}, error)

	// List gets multiple books
	List(context.Context) (interface{}, error)

	// Insert inserts to the DB returning error on conflicts
	Insert(context.Context, []interface{}) error

	// Upsert inserts to the DB validating conflicts on records
	Upsert(context.Context, []interface{}) error
}

type SQLRepository interface {
	Repository

	// BeginTransaction initialize a new psql transaction
	BeginTransaction(context.Context) (interface{}, error)

	// EndTransaction finish a psql transaction
	EndTransaction(context.Context) error
}
