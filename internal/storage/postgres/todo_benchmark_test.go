package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/models"
)

func BenchmarkCreateTodo_Success(b *testing.B) {
	// Context with timeout to manage long-running operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < b.N; i++ {
		// Mock data
		todo := &models.Todo{
			Id:        uuid.NewString(),
			Title:     "Benchmark Todo",
			CreatedAt: time.Now().UTC(),
		}

		// Call method
		_, err := todoRepo.Create(ctx, todo)
		if err != nil {
			b.Fatalf("Failed to create todo: %v", err)
		}
	}
}
