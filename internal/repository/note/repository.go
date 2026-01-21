package note

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lva100/go-grpc/internal/model"
	"github.com/lva100/go-grpc/internal/repository"
	"github.com/lva100/go-grpc/internal/repository/note/converter"
	modelRepo "github.com/lva100/go-grpc/internal/repository/note/model"
)

const (
	tableName = "note"

	idColumn        = "id"
	titleColumn     = "title"
	contentColumn   = "content"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.NoteRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.NoteInfo) (int64, error) {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn, contentColumn).
		Values(gofakeit.City(), gofakeit.Address().Street).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %s", err)
	}
	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		log.Fatalf("failed to insert notes: %s", err)
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Note, error) {
	builderSelectOne := sq.Select(idColumn, titleColumn, contentColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %s", err)
	}

	var note modelRepo.Note

	err = r.db.QueryRow(ctx, query, args...).Scan(&note.Id, &note.Info.Title, &note.Info.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		log.Printf("failed to selected notes: %s", err)
	}

	return converter.ToNoteFromRepo(&note), nil
}
