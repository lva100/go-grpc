package note

import (
	"context"

	"github.com/lva100/go-grpc/internal/converter"
	"github.com/lva100/go-grpc/pkg/note_v1"
)

func (i *Implementation) Get(ctx context.Context, req *note_v1.GetRequest) (*note_v1.GetResponse, error) {
	noteObj, err := i.noteService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &note_v1.GetResponse{
		Note: converter.ToNoteFromService(noteObj),
	}, nil
}
