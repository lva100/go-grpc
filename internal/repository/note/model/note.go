package model

import (
	"database/sql"
	"time"
)

type Note struct {
	Id        int64        `db:"id"`
	Info      NoteInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type NoteInfo struct {
	Title   string `db:"title"`
	Content string `db:"content"`
}
