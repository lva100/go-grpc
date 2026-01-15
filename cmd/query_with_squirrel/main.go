package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v5"
)

const (
	dbDSN = "host=localhost port=5432 dbname=PG_TEST user=admin password=123456"
)

func main() {
	ctx := context.Background()

	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	defer func() {
		err := con.Close(ctx)
		if err != nil {
			log.Fatalf("failed to close conn to db: %s", err)
		}
	}()

	builderInsert := sq.Insert("note").
		PlaceholderFormat(sq.Dollar).
		Columns("title", "body").
		Values(gofakeit.City(), gofakeit.Address().Street).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %s", err)
	}
	var noteID int
	err = con.QueryRow(ctx, query, args...).Scan(&noteID)
	if err != nil {
		log.Fatalf("failed to insert notes: %s", err)
	}

	log.Printf("Inserted note with id: %d", noteID)

	builderSelect := sq.Select("id", "title", "body", "created_at", "updated_at").
		From("note").
		PlaceholderFormat(sq.Dollar).
		OrderBy("id ASC").
		Limit(10)

	query, args, err = builderSelect.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %s", err)
	}
	rows, err := con.Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to select notes: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title, body string
		var created_at time.Time
		var updated_at sql.NullTime

		err = rows.Scan(&id, &title, &body, &created_at, &updated_at)
		if err != nil {
			log.Fatalf("failed to scan note: %s", err)
		}
		log.Printf("id: %d, title: %s, body: %s, created_at: %s, updated_at: %s", id, title, body, created_at.Format("02-01-2006"), handleNullTime(updated_at))
	}

	builderUpdate := sq.Update("note").
		PlaceholderFormat(sq.Dollar).
		Set("title", gofakeit.City()).
		Set("body", gofakeit.Address().Street).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": noteID})
	query, args, err = builderUpdate.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %s", err)
	}
	res, err := con.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to insert notes: %s", err)
	}
	log.Printf("Updated %d rows", res.RowsAffected())
}

func handleNullTime(tm sql.NullTime) string {
	if tm.Valid {
		return tm.Time.Format("02-01-2006")
	} else {
		return "-/-/-"
	}
}
