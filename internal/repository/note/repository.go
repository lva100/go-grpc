package note

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/lva100/go-grpc/internal/client/db"
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
	db db.Client
}

func NewRepository(db db.Client) repository.NoteRepository {
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
	q := db.Query{
		Name:     "note_repository.Create",
		QueryRaw: query,
	}
	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
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

	q := db.Query{
		Name:     "note_repository.Get",
		QueryRaw: query,
	}

	var note modelRepo.Note

	// err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&note.Id, &note.Info.Title, &note.Info.Content, &note.CreatedAt, &note.UpdatedAt)
	err = r.db.DB().ScanOneContext(ctx, &note, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToNoteFromRepo(&note), nil
}
