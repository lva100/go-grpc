package note

import (
	"context"
	"log"

	"github.com/lva100/go-grpc/internal/converter"
	"github.com/lva100/go-grpc/pkg/note_v1"
)

func (i *Implementation) Create(ctx context.Context, req *note_v1.CreateRequest) (*note_v1.CreateResponse, error) {
	id, err := i.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}
	log.Printf("inserted note with id: %d", id)

	return &note_v1.CreateResponse{
		Id: id,
	}, nil
}
