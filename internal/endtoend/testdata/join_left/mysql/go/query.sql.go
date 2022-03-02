// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
	"time"
)

const allAuthors = `-- name: AllAuthors :many
SELECT  a.id, a.name, a.parent_id, p.id, p.name, p.parent_id
FROM    authors a
        LEFT JOIN authors p
            ON a.parent_id = p.id
`

type AllAuthorsRow struct {
	ID         int32
	Name       string
	ParentID   sql.NullInt32
	ID_2       sql.NullInt32
	Name_2     sql.NullString
	ParentID_2 sql.NullInt32
}

func (q *Queries) AllAuthors(ctx context.Context) ([]AllAuthorsRow, error) {
	rows, err := q.db.QueryContext(ctx, allAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllAuthorsRow
	for rows.Next() {
		var i AllAuthorsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ParentID,
			&i.ID_2,
			&i.Name_2,
			&i.ParentID_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const allAuthorsAliases = `-- name: AllAuthorsAliases :many
SELECT  a.id, a.name, a.parent_id, p.id, p.name, p.parent_id
FROM    authors a
        LEFT JOIN authors p
            ON a.parent_id = p.id
`

type AllAuthorsAliasesRow struct {
	ID         int32
	Name       string
	ParentID   sql.NullInt32
	ID_2       sql.NullInt32
	Name_2     sql.NullString
	ParentID_2 sql.NullInt32
}

func (q *Queries) AllAuthorsAliases(ctx context.Context) ([]AllAuthorsAliasesRow, error) {
	rows, err := q.db.QueryContext(ctx, allAuthorsAliases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllAuthorsAliasesRow
	for rows.Next() {
		var i AllAuthorsAliasesRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ParentID,
			&i.ID_2,
			&i.Name_2,
			&i.ParentID_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const allAuthorsAliases2 = `-- name: AllAuthorsAliases2 :many
SELECT  a.id, a.name, a.parent_id, p.id, p.name, p.parent_id
FROM    authors a
        LEFT JOIN authors p
            ON a.parent_id = p.id
`

type AllAuthorsAliases2Row struct {
	ID         int32
	Name       string
	ParentID   sql.NullInt32
	ID_2       sql.NullInt32
	Name_2     sql.NullString
	ParentID_2 sql.NullInt32
}

func (q *Queries) AllAuthorsAliases2(ctx context.Context) ([]AllAuthorsAliases2Row, error) {
	rows, err := q.db.QueryContext(ctx, allA