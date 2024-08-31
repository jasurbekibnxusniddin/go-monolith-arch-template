package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo_Success(t *testing.T) {
	// Context with timeout to manage long-running operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Mock data
	todo := &models.Todo{
		Id:        uuid.NewString(),
		Title:     "Test Todo",
		CreatedAt: time.Now().UTC(),
	}

	// Call method
	createdTodo, err := todoRepo.Create(ctx, todo)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdTodo)
	assert.Equal(t, todo.Id, createdTodo.Id)
	assert.Equal(t, todo.Title, createdTodo.Title)
	assert.WithinDuration(t, todo.CreatedAt, createdTodo.CreatedAt, time.Second)
}

func TestCreateTodo_Error(t *testing.T) {
	// Context with timeout to manage long-running operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Mock data with missing title to trigger an error
	todo := &models.Todo{
		Id:        "",
		CreatedAt: time.Now(),
	}

	// Call method
	createdTodo, err := todoRepo.Create(ctx, todo)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, createdTodo)
}
