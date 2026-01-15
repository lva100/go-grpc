package main

import (
	"context"
	"database/sql"
	"log"
	"time"

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

	res, err := con.Exec(ctx, "INSERT INTO note (title, body) VALUES ($1, $2)", gofakeit.City(), gofakeit.Address().Street)
	if err != nil {
		log.Fatalf("failed to insert notes: %s", err)
	}

	log.Printf("Inserted %d rows", res.RowsAffected())

	rows, err := con.Query(ctx, "SELECT id, title, body, created_at, updated_at FROM note")
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
}

func handleNullTime(tm sql.NullTime) string {
	if tm.Valid {
		return tm.Time.Format("02-01-2006")
	} else {
		return "-"
	}
}
