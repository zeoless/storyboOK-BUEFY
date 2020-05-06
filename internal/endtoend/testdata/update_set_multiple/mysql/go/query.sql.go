
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"
)

const updateSetMultiple = `-- name: UpdateSetMultiple :exec
UPDATE foo SET name = ?, slug = ?
`

type UpdateSetMultipleParams struct {
	Name string
	Slug string
}

func (q *Queries) UpdateSetMultiple(ctx context.Context, arg UpdateSetMultipleParams) error {
	_, err := q.db.ExecContext(ctx, updateSetMultiple, arg.Name, arg.Slug)
	return err
}