
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: batch.go

package batch

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

const booksByYear = `-- name: BooksByYear :batchmany
SELECT book_id, author_id, isbn, book_type, title, year, available, tags FROM books
WHERE year = $1
`

type BooksByYearBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

func (q *Queries) BooksByYear(ctx context.Context, year []int32) *BooksByYearBatchResults {
	batch := &pgx.Batch{}
	for _, a := range year {
		vals := []interface{}{
			a,
		}
		batch.Queue(booksByYear, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &BooksByYearBatchResults{br, len(year), false}
}

func (b *BooksByYearBatchResults) Query(f func(int, []Book, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		var items []Book
		if b.closed {
			if f != nil {
				f(t, items, errors.New("batch already closed"))
			}
			continue
		}
		err := func() error {
			rows, err := b.br.Query()
			defer rows.Close()
			if err != nil {
				return err
			}
			for rows.Next() {
				var i Book
				if err := rows.Scan(
					&i.BookID,
					&i.AuthorID,
					&i.Isbn,
					&i.BookType,
					&i.Title,
					&i.Year,
					&i.Available,
					&i.Tags,
				); err != nil {
					return err
				}
				items = append(items, i)
			}
			return rows.Err()
		}()
		if f != nil {
			f(t, items, err)
		}
	}
}

func (b *BooksByYearBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const createBook = `-- name: CreateBook :batchone
INSERT INTO books (
    author_id,
    isbn,
    book_type,
    title,
    year,
    available,
    tags
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING book_id, author_id, isbn, book_type, title, year, available, tags
`

type CreateBookBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type CreateBookParams struct {
	AuthorID  int32     `json:"author_id"`
	Isbn      string    `json:"isbn"`
	BookType  BookType  `json:"book_type"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Available time.Time `json:"available"`
	Tags      []string  `json:"tags"`
}

func (q *Queries) CreateBook(ctx context.Context, arg []CreateBookParams) *CreateBookBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.AuthorID,
			a.Isbn,
			a.BookType,
			a.Title,
			a.Year,
			a.Available,
			a.Tags,
		}
		batch.Queue(createBook, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &CreateBookBatchResults{br, len(arg), false}
}

func (b *CreateBookBatchResults) QueryRow(f func(int, Book, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		var i Book
		if b.closed {
			if f != nil {
				f(t, i, errors.New("batch already closed"))
			}
			continue
		}
		row := b.br.QueryRow()
		err := row.Scan(
			&i.BookID,
			&i.AuthorID,
			&i.Isbn,
			&i.BookType,
			&i.Title,
			&i.Year,
			&i.Available,
			&i.Tags,
		)
		if f != nil {
			f(t, i, err)
		}
	}
}

func (b *CreateBookBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const deleteBook = `-- name: DeleteBook :batchexec
DELETE FROM books
WHERE book_id = $1
`

type DeleteBookBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

func (q *Queries) DeleteBook(ctx context.Context, bookID []int32) *DeleteBookBatchResults {
	batch := &pgx.Batch{}
	for _, a := range bookID {
		vals := []interface{}{
			a,
		}
		batch.Queue(deleteBook, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &DeleteBookBatchResults{br, len(bookID), false}
}

func (b *DeleteBookBatchResults) Exec(f func(int, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		if b.closed {
			if f != nil {
				f(t, errors.New("batch already closed"))
			}
			continue
		}
		_, err := b.br.Exec()
		if f != nil {
			f(t, err)
		}
	}
}

func (b *DeleteBookBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const deleteBookNamedFunc = `-- name: DeleteBookNamedFunc :batchexec
DELETE FROM books
WHERE book_id = $1
`

type DeleteBookNamedFuncBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

func (q *Queries) DeleteBookNamedFunc(ctx context.Context, bookID []int32) *DeleteBookNamedFuncBatchResults {
	batch := &pgx.Batch{}
	for _, a := range bookID {
		vals := []interface{}{
			a,
		}
		batch.Queue(deleteBookNamedFunc, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &DeleteBookNamedFuncBatchResults{br, len(bookID), false}
}

func (b *DeleteBookNamedFuncBatchResults) Exec(f func(int, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		if b.closed {
			if f != nil {
				f(t, errors.New("batch already closed"))
			}
			continue
		}
		_, err := b.br.Exec()
		if f != nil {
			f(t, err)
		}
	}
}

func (b *DeleteBookNamedFuncBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const deleteBookNamedSign = `-- name: DeleteBookNamedSign :batchexec
DELETE FROM books
WHERE book_id = $1
`

type DeleteBookNamedSignBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

func (q *Queries) DeleteBookNamedSign(ctx context.Context, bookID []int32) *DeleteBookNamedSignBatchResults {
	batch := &pgx.Batch{}
	for _, a := range bookID {
		vals := []interface{}{
			a,
		}
		batch.Queue(deleteBookNamedSign, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &DeleteBookNamedSignBatchResults{br, len(bookID), false}
}

func (b *DeleteBookNamedSignBatchResults) Exec(f func(int, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		if b.closed {
			if f != nil {
				f(t, errors.New("batch already closed"))
			}
			continue
		}
		_, err := b.br.Exec()
		if f != nil {
			f(t, err)
		}
	}
}

func (b *DeleteBookNamedSignBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const getBiography = `-- name: GetBiography :batchone
SELECT biography FROM authors
WHERE author_id = $1
`

type GetBiographyBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

func (q *Queries) GetBiography(ctx context.Context, authorID []int32) *GetBiographyBatchResults {
	batch := &pgx.Batch{}
	for _, a := range authorID {
		vals := []interface{}{
			a,
		}
		batch.Queue(getBiography, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &GetBiographyBatchResults{br, len(authorID), false}
}

func (b *GetBiographyBatchResults) QueryRow(f func(int, pgtype.JSONB, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		var biography pgtype.JSONB
		if b.closed {
			if f != nil {
				f(t, biography, errors.New("batch already closed"))
			}
			continue
		}
		row := b.br.QueryRow()
		err := row.Scan(&biography)
		if f != nil {
			f(t, biography, err)
		}
	}
}

func (b *GetBiographyBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}

const updateBook = `-- name: UpdateBook :batchexec
UPDATE books
SET title = $1, tags = $2
WHERE book_id = $3
`

type UpdateBookBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type UpdateBookParams struct {
	Title  string   `json:"title"`
	Tags   []string `json:"tags"`
	BookID int32    `json:"book_id"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg []UpdateBookParams) *UpdateBookBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.Title,
			a.Tags,
			a.BookID,
		}
		batch.Queue(updateBook, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &UpdateBookBatchResults{br, len(arg), false}
}

func (b *UpdateBookBatchResults) Exec(f func(int, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		if b.closed {
			if f != nil {
				f(t, errors.New("batch already closed"))
			}
			continue
		}
		_, err := b.br.Exec()
		if f != nil {
			f(t, err)
		}
	}
}

func (b *UpdateBookBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}