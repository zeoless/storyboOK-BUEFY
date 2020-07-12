// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const fullJoin = `-- name: FullJoin :many
SELECT f.id, f.bar_id, b.id
FROM foo f
FULL OUTER JOIN bar b ON b.id = f.bar_id
WHERE f.id = $1
`

type FullJoinRow struct {
	ID    sql.NullInt32
	BarID sql.NullInt32
	ID_2  sql.NullInt32
}

func (q *Queries) FullJoin(ctx context.Context, id int32) ([]FullJoinRow, error) {
	rows, err := q.db.QueryContext(ctx, fullJoin, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FullJoinRow