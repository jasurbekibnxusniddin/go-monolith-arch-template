package postgres

import (
	"context"

	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/models"
)

type TodoRepoI interface {
	Create(ctx context.Context, in *models.Todo) (*models.Todo, error)
	GetList(ctx context.Context, limit, page int32) (*models.GetListResp, error)
}
