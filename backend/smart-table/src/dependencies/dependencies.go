package dependencies

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/smart-table/src/config"
	"github.com/smart-table/src/logging"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Dependencies struct {
	Config     *config.Config
	DBConnPool *pgxpool.Pool
	Logger     *zap.Logger
	S3Client   *minio.Client
}

func InitDependencies(cfg *config.Config) *Dependencies {
	logger := logging.InitLogger(cfg)
	dbPool := initDBPool(cfg)
	s3Client := initS3Client(cfg)

	return &Dependencies{
		Config:     cfg,
		DBConnPool: dbPool,
		Logger:     logger,
		S3Client:   s3Client,
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

func initS3Client(cfg *config.Config) *minio.Client {
	client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.S3.AccessKey, cfg.S3.SecretKey, ""),
		Secure: true,
		Region: cfg.S3.Region,
	})
	if err != nil {
		logging.GetLogger().Fatal("Error creating MinIO client", zap.Error(err))
	}

	return client
}
