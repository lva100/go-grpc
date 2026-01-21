package note

import (
	"github.com/lva100/go-grpc/internal/repository"
	"github.com/lva100/go-grpc/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
}

func NewService(noteRepository repository.NoteRepository) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
	}
}
