package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/lva100/go-grpc/internal/api/note"
	"github.com/lva100/go-grpc/internal/model"
	"github.com/lva100/go-grpc/internal/service"
	serviceMocks "github.com/lva100/go-grpc/internal/service/mocks"
	"github.com/lva100/go-grpc/pkg/note_v1"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type noteServiceMockFunc func(mc *minimock.Controller) service.NoteService

	type args struct {
		ctx context.Context
		req *note_v1.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = gofakeit.Int64()
		title   = gofakeit.Animal()
		content = gofakeit.Animal()

		serviceErr = fmt.Errorf("service error")

		req = &note_v1.CreateRequest{
			Info: &note_v1.NoteInfo{
				Title:   title,
				Content: content,
			},
		}
		info = &model.NoteInfo{
			Title:   title,
			Content: content,
		}

		res = &note_v1.CreateResponse{
			Id: id,
		}
	)
	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *note_v1.CreateResponse
		err             error
		noteServiceMock noteServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				req: req,
				ctx: ctx,
			},
			want: res,
			err:  nil,
			noteServiceMock: func(mc *minimock.Controller) service.NoteService {
				mock := serviceMocks.NewNoteServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
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
			err:  serviceErr,
			noteServiceMock: func(mc *minimock.Controller) service.NoteService {
				mock := serviceMocks.NewNoteServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(0, serviceErr)
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			noteServiceMock := tt.noteServiceMock(mc)
			api := note.NewImplementation(noteServiceMock)

			resHandler, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}
}
