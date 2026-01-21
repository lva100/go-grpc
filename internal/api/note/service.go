package note

import (
	"github.com/lva100/go-grpc/internal/service"
	"github.com/lva100/go-grpc/pkg/note_v1"
)

type Implementation struct {
	note_v1.UnimplementedNoteV1Server
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
