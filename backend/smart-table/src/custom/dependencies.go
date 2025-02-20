package custom

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Dependencies struct {
	Config     *Config
	DbConnPool *pgxpool.Pool
	Logger     *zap.Logger
}

func InitDependencies() *Dependencies {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger := InitLogger(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Database.Timeout)
	defer cancel()

	dbConnStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name,
	)

	config, err := pgxpool.ParseConfig(dbConnStr)
	if err != nil {
		logger.Fatal("Failed to parse database config", zap.Error(err))
	}

	config.MaxConns = int32(cfg.Database.MaxConnections)
	config.MinConns = int32(cfg.Database.MinConnections)
	config.MaxConnLifetime = cfg.Database.MaxConnLifetime
	config.MaxConnIdleTime = cfg.Database.MaxConnIdleTime

	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logger.Fatal("Failed to create database connection pool", zap.Error(err))
	}

	err = dbPool.Ping(ctx)
	if err != nil {
		logger.Fatal("Database connection check failed", zap.Error(err))
	}

	logger.Info("Successfully connected to the database")

	return &Dependencies{
		Config:     cfg,
		DbConnPool: dbPool,
		Logger:     logger,
	}
}
