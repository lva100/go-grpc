package note

import (
	"github.com/lva100/go-grpc/internal/client/db"
	"github.com/lva100/go-grpc/internal/repository"
	"github.com/lva100/go-grpc/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
	txManager      db.TxManager
}

func NewService(noteRepository repository.NoteRepository, txManager db.TxManager) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
		txManager:      txManager,
	}
}

func NewMockService(noteRepository repository.NoteRepository) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
	}
}
