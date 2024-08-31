package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/models"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/pkgs/log"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

type todoService struct {
	storage storage.StorageI
	log     log.Log
}

func NewTodoService(storage storage.StorageI, log log.Log) *todoService {
	return &todoService{storage: storage, log: log}
}

func (t *todoService) Create(ctx context.Context, in *models.CreateTodoReq) (*models.Todo, error) {

	todo, err := t.storage.GetTodoRepo().Create(ctx, &models.Todo{
		Id:        uuid.New().String(),
		Title:     in.Title,
		CreatedAt: time.Now(),
	})
	if err != nil {
		t.log.Error("erron on creating todo.", logger.Error(err))
		return nil, err
	}

	return todo, nil
}
