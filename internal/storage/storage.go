package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/pkgs/log"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/storage/postgres"
)

type StorageI interface {
	GetTodoRepo() postgres.TodoRepoI
}

type storage struct {
	todoRepo postgres.TodoRepoI
}

func NewStorage(db *pgxpool.Pool, log log.Log) StorageI {
	return &storage{
		todoRepo: postgres.NewTodo(db, log),
	}
}

func (s *storage) GetTodoRepo() postgres.TodoRepoI {
	return s.todoRepo
}
