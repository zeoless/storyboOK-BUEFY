//go:build examples
// +build examples

package booktest

import (
	"context"
	"testing"
	"time"

	"github.com/kyleconroy/sqlc/internal/sqltest"
)

func TestBooks(t *testing.T) {
	db, cleanup := sqltest.PostgreSQL(t, []string{"schema.sql"})
	defer cleanup()

	ctx := context.Background()
	dq := New(db)

	// create an author
	a, err := dq.CreateAuthor(ctx, "Unknown Master")
	if err != nil {
		t.Fatal(err)
	}

	// create transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	tq := dq.WithTx(tx)

	// save first book
	now := time.Now()
	_, err = tq.CreateBook(ctx, CreateBookParams{
		AuthorID:  a.AuthorID,
		Isbn:      "1",
		Title:     "my book title",
		BookType:  BookTypeFICTION,
		Year:      2016,
		Available: now,
		Tags:      []string{},
	})
	if err != nil {
		t.Fatal(err)
	}

	// save second book
	b1, err := tq.CreateBook(ctx, CreateBookParams{
		AuthorID:  a.AuthorID,
		Isbn:      "2",
		Title:     "the second book",
		BookType:  BookTypeFICTION,
		Year:      2016,
		Available: now,
		Tags:      []string{"cool", "unique"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// update the title and tags
	err = tq.UpdateBook(ctx, UpdateBookParams{
		BookID: b1.BookID,
		Title:  "changed second title",
		Tags:   []string{"cool", "disastor"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// save third book
	_, err = tq.CreateBook(ctx, CreateBookParams{
		AuthorID:  a.AuthorID,
		Isbn:      "3",
		Title:     "the third book",
		BookType:  BookTypeFICTION,
		Year:      2001,
		Available: now,
		Tags:      []string{"cool"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// save fourth book
