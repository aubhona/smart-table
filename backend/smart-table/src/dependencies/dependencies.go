package dependencies

import (
	"context"
	"fmt"
	"log"

	"github.com/smart-table/src/config"
	"github.com/smart-table/src/logging"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Dependencies struct {
	Config     *config.Config
	DBConnPool *pgxpool.Pool
	Logger     *zap.Logger
}

func InitDependencies() *Dependencies {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger := logging.InitLogger(cfg)
	dbPool := initDBPool(cfg)

	return &Dependencies{
		Config:     cfg,
		DBConnPool: dbPool,
		Logger:     logger,
	}
}

func initDBPool(cfg *config.Config) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Database.Timeout)
	defer cancel()

	dbConnStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name,
	)

	pgConfig, err := pgxpool.ParseConfig(dbConnStr)
	if err != nil {
		logging.GetLogger().Fatal("Failed to parse database config", zap.Error(err))
	}

	pgConfig.MaxConns = cfg.Database.MaxConnections
	pgConfig.MinConns = cfg.Database.MinConnections
	pgConfig.MaxConnLifetime = cfg.Database.MaxConnLifetime
	pgConfig.MaxConnIdleTime = cfg.Database.MaxConnIdleTime

	dbPool, err := pgxpool.NewWithConfig(ctx, pgConfig)
	if err != nil {
		logging.GetLogger().Fatal("Failed to create database connection pool", zap.Error(err))
	}

	err = dbPool.Ping(ctx)
	if err != nil {
		logging.GetLogger().Fatal("Database connection check failed", zap.Error(err))
	}

	logging.GetLogger().Info("Successfully connected to the database")

	return dbPool
}
