package tests //nolint

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/pressly/goose"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/smart-table/src/config"
	"github.com/smart-table/src/dependencies"
	"github.com/smart-table/src/domains/customer/di"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/dig"
)

var container *dig.Container

func TestMain(m *testing.M) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "test_db",
			"POSTGRES_USER":     "test_user",
			"POSTGRES_PASSWORD": "test_password",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	dbContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal("Failed to launch db container", err)
	}
	defer dbContainer.Terminate(ctx) //nolint

	host, err := dbContainer.Host(ctx)
	if err != nil {
		log.Fatal("Failed to get container host:", err)
	}

	port, err := dbContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatal("Failed to get container port:", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	cfg.Database.Host = host
	cfg.Database.Port = port.Port()

	db, err := sql.Open(
		"pgx",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name))
	if err != nil {
		log.Fatalf("Failed to connect db: %v", err)
	}

	if err = goose.Up(db, "../postgresql/smart_table"); err != nil {
		log.Fatalf("Failed to create migration: %v", err)
	}

	deps := dependencies.InitDependencies(cfg)
	logger := deps.Logger

	container, err = di.BuildContainer(deps)
	if err != nil {
		logger.Fatal(err.Error())
	}

	code := m.Run()

	os.Exit(code)
}
