package app

import (
	"context"
	"log"

	"github.com/lva100/go-grpc/internal/api/note"
	"github.com/lva100/go-grpc/internal/client/db"
	"github.com/lva100/go-grpc/internal/client/db/pg"
	"github.com/lva100/go-grpc/internal/client/db/transaction"
	"github.com/lva100/go-grpc/internal/closer"
	"github.com/lva100/go-grpc/internal/config"
	"github.com/lva100/go-grpc/internal/config/env"
	"github.com/lva100/go-grpc/internal/repository"
	noteRepository "github.com/lva100/go-grpc/internal/repository/note"
	"github.com/lva100/go-grpc/internal/service"
	noteService "github.com/lva100/go-grpc/internal/service/note"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	noteRepository repository.NoteRepository

	noteService service.NoteService

	noteImpl *note.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err)
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err)
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err)
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err)
		}
		closer.Add(func() error {
			cl.Close()
			return nil
		})
		s.dbClient = cl
	}
	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}
	return s.txManager
}

func (s *serviceProvider) NoteRepository(ctx context.Context) repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepository.NewRepository(s.DBClient(ctx))
	}
	return s.noteRepository
}

func (s *serviceProvider) NoteService(ctx context.Context) service.NoteService {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.NoteRepository(ctx), s.TxManager(ctx))
	}
	return s.noteService
}

func (s *serviceProvider) NoteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.NoteService(ctx))
	}
	return s.noteImpl
}
