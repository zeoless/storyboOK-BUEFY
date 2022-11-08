// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const test = `-- name: Test :one
select txt from Demo
where txt ~~ '%' || $1 || '%'
`

func (q *Queries) Test(ctx context.Context, val string) (string, error) {
	row := q.db.QueryRow(ctx, test, val)
	var txt string
