package repository

import (
	"context"

	"github.com/lva100/go-grpc/pkg/note_v1"
)

type NoteRepository interface {
	Create(ctx context.Context, info *note_v1.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*note_v1.Note, error)
	// Get1(ctx context.Context, id int64) (string, error)
}
