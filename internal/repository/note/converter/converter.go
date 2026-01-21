package converter

import (
	"github.com/lva100/go-grpc/internal/model"
	modelRepo "github.com/lva100/go-grpc/internal/repository/note/model"
)

func ToNoteFromRepo(note *modelRepo.Note) *model.Note {
	return &model.Note{
		Id:        note.Id,
		Info:      ToNoteInfoFromRepo(note.Info),
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

func ToNoteInfoFromRepo(info modelRepo.NoteInfo) model.NoteInfo {
	return model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
