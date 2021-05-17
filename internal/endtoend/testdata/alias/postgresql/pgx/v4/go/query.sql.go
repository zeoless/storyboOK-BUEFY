
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"
)

const aliasBar = `-- name: AliasBar :exec
DELETE FROM bar b
WHERE b.id = $1
`

func (q *Queries) AliasBar(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, aliasBar, id)
	return err
}