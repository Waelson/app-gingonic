// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package repository

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (sql.Result, error)
	DeleteAuthor(ctx context.Context, id int64) error
	GetAuthor(ctx context.Context, id int64) (Authors, error)
	ListAuthors(ctx context.Context) ([]Authors, error)
}

var _ Querier = (*Queries)(nil)
