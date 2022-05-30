
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const getAll = `-- name: GetAll :many
SELECT "CamelCase" FROM users
`

func (q *Queries) GetAll(ctx context.Context) ([]sql.NullString, error) {
	rows, err := q.db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var CamelCase sql.NullString
		if err := rows.Scan(&CamelCase); err != nil {
			return nil, err
		}
		items = append(items, CamelCase)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}