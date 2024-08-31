package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/config"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/pkgs/db"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/pkgs/log"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

type ServiceI interface {
	GetTodoService() *todoService
	CleanUP()
}

type services struct {
	pgxPool     *pgxpool.Pool
	todoService *todoService
}

func NewService(cfg config.Config, log log.Log) ServiceI {

	pgxPool, err := db.ConnDB(cfg.PgConfig)
	if err != nil {
		log.Error("error on connecting with database!", logger.Error(err))
		return nil
	}
	//defer pgxPool.Close()

	storage := storage.NewStorage(pgxPool, log)

	services := services{
		pgxPool:     pgxPool,
		todoService: NewTodoService(storage, log),
	}

	return &services
}

func (s *services) GetTodoService() *todoService {

	if s != nil {
		return s.todoService
	}

	return &todoService{}
}

func (s *services) CleanUP() {
	s.pgxPool.Close()
	s = nil
}
