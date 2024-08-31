package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/models"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/pkgs/log"
)

type todoRepo struct {
	db  *pgxpool.Pool
	log log.Log
}

func NewTodo(db *pgxpool.Pool, log log.Log) TodoRepoI {
	return &todoRepo{db: db, log: log}
}

func (t *todoRepo) Create(ctx context.Context, in *models.Todo) (*models.Todo, error) {

	query := `
		INSERT INTO 
			todos (
				todo_id,
				title,
				created_at
			) VALUES (
				$1, $2, $3
			) 
		RETURNING 
			todo_id, 
			title, 
			created_at`

	row := t.db.QueryRow(ctx, query, in.Id, in.Title, in.CreatedAt)

	var createdTodo models.Todo
	if err := row.Scan(&createdTodo.Id, &createdTodo.Title, &createdTodo.CreatedAt); err != nil {
		return nil, fmt.Errorf("failed to retrieve inserted todo: %w", err)
	}

	return &createdTodo, nil
}

func (t *todoRepo) GetList(ctx context.Context, limit, page int32) (*models.GetListResp, error) {
	return nil, nil
}
