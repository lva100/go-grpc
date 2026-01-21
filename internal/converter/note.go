package converter

import (
	"github.com/lva100/go-grpc/internal/model"
	"github.com/lva100/go-grpc/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNoteFromService(note *model.Note) *note_v1.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &note_v1.Note{
		Id:        note.Id,
		Info:      ToNoteInfoFromService(&note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToNoteInfoFromService(info *model.NoteInfo) *note_v1.NoteInfo {
	return &note_v1.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}

func ToNoteInfoFromDesc(info *note_v1.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
