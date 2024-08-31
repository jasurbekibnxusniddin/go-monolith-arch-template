package postgres_test

import (
	"os"
	"testing"

	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/config"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/pkgs/db"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/pkgs/log"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/storage/postgres"
	"github.com/saidamir98/udevs_pkg/logger"
)

var (
	todoRepo postgres.TodoRepoI
)

func TestMain(m *testing.M) {
	cfg := config.Load()
	log := log.NewLogger(cfg.GeneralConfig)

	var err error
	mockDb, err := db.ConnDB(cfg.PgConfig)
	if err != nil {
		log.Error("Failed to connect to database.", logger.Error(err))
		os.Exit(1)
	}

	todoRepo = postgres.NewTodo(mockDb, log)
	exitCode := m.Run()

	// Cleanup
	defer mockDb.Close()

	os.Exit(exitCode)
}
