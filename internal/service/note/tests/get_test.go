package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/lva100/go-grpc/internal/model"
	"github.com/lva100/go-grpc/internal/repository"
	repoMocks "github.com/lva100/go-grpc/internal/repository/mocks"
	nt "github.com/lva100/go-grpc/internal/service/note"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type noteRepositoryMockFunc func(mc *minimock.Controller) repository.NoteRepository

	type args struct {
		ctx context.Context
		req *model.Note
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		title     = gofakeit.Animal()
		content   = gofakeit.Animal()
		createdAt = gofakeit.Date()
		updatedAt = sql.NullTime{Valid: false}

		repoErr = fmt.Errorf("repo error")

		req = &model.Note{
			Id: id,
			Info: model.NoteInfo{
				Title:   title,
				Content: content,
			},
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
	)
	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name               string
		args               args
		want               *model.Note
		err                error
		noteRepositoryMock noteRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				req: req,
				ctx: ctx,
			},
			want: req,
			err:  nil,
			noteRepositoryMock: func(mc *minimock.Controller) repository.NoteRepository {
				mock := repoMocks.NewNoteRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(req, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				req: req,
				ctx: ctx,
			},
			want: nil,
			err:  repoErr,
			noteRepositoryMock: func(mc *minimock.Controller) repository.NoteRepository {
				mock := repoMocks.NewNoteRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, repoErr)
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			noteRepoMock := tt.noteRepositoryMock(mc)
			service := nt.NewMockService(noteRepoMock)

			newID, err := service.Get(tt.args.ctx, id)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
