package smarttable_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	dbAdmin "github.com/smart-table/src/domains/admin/infra/pg/codegen"
	dbCustomer "github.com/smart-table/src/domains/customer/infra/pg/codegen"
	"github.com/smart-table/src/utils"

	"github.com/pressly/goose"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/smart-table/src/config"
	"github.com/smart-table/src/dependencies"
	adminDi "github.com/smart-table/src/domains/admin/di"
	customerDi "github.com/smart-table/src/domains/customer/di"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/dig"
)

var testMutex sync.Mutex
var container = dig.New()
var responseRecorder = httptest.NewRecorder()
var ginCtx, _ = gin.CreateTestContext(responseRecorder)
var deps = &dependencies.Dependencies{}

func GetAdminQueries() *dbAdmin.Queries {
	return dbAdmin.New(deps.DBConnPool)
}

func GetCustomerQueries() *dbCustomer.Queries {
	return dbCustomer.New(deps.DBConnPool)
}

func GetCtx() context.Context {
	return ginCtx
}

func GetTestMutex() *sync.Mutex {
	return &testMutex
}

func GetContainer() *dig.Container {
	return container
}

func GetDeps() *dependencies.Dependencies {
	return deps
}

func SetupOnceTest() {
	db, err := sql.Open(
		"pgx",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			deps.Config.Database.Host, deps.Config.Database.Port,
			deps.Config.Database.User, deps.Config.Database.Password, deps.Config.Database.Name))
	if err != nil {
		log.Fatalf("Failed to connect db: %v", err)
	}
	defer db.Close()

	if err = goose.Up(db, "../postgresql/smart_table"); err != nil {
		log.Fatalf("Failed to create migration: %v", err)
	}
}

func CleanTest() {
	db, err := sql.Open(
		"pgx",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			deps.Config.Database.Host, deps.Config.Database.Port,
			deps.Config.Database.User, deps.Config.Database.Password, deps.Config.Database.Name))
	if err != nil {
		log.Fatalf("Failed to connect db: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
        DO $$
        DECLARE
            tbl text;
        BEGIN
            FOR tbl IN 
                SELECT tablename 
                FROM pg_tables 
                WHERE schemaname = 'smart_table_customer'
            LOOP
                EXECUTE 'TRUNCATE TABLE smart_table_customer.' || quote_ident(tbl) || ' CASCADE';
                RAISE NOTICE 'Очищена таблица: smart_table_customer.%', tbl;
            END LOOP;
        END $$;
    `)
	if err != nil {
		log.Fatalf("Failed to truncate smart_table_customer: %v", err)
	}

	_, err = db.Exec(`
        DO $$
        DECLARE
            tbl text;
        BEGIN
            FOR tbl IN 
                SELECT tablename 
                FROM pg_tables 
                WHERE schemaname = 'smart_table_admin'
            LOOP
                EXECUTE 'TRUNCATE TABLE smart_table_admin.' || quote_ident(tbl) || ' CASCADE';
                RAISE NOTICE 'Очищена таблица: smart_table_admin.%', tbl;
            END LOOP;
        END $$;
    `)
	if err != nil {
		log.Fatalf("Failed to truncate smart_table_admin: %v", err)
	}
}

func TestMain(m *testing.M) {
	container = dig.New()
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

	deps = dependencies.InitDependencies(cfg)
	logger := deps.Logger

	SetupOnceTest()

	err = container.Provide(func() *dependencies.Dependencies {
		return deps
	})
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = customerDi.AddDeps(container)
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = adminDi.AddDeps(container)
	if err != nil {
		logger.Fatal(err.Error())
	}

	ginCtx.Set(utils.DiContainerName, container)

	code := m.Run()

	_ = dbContainer.Terminate(ctx)

	os.Exit(code)
}
