package config

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/joho/godotenv"
)

const kConfigPath = "configs/config.yaml"
const kEnvPath = "configs/.env"

type Config struct {
	App struct {
		Port  int         `mapstructure:"port"`
		Env   Environment `mapstructure:"env"`
		Admin struct {
			Jwt struct {
				SecretKey  string        `mapstructure:"secret_key"`
				Expiration time.Duration `mapstructure:"expiration"`
			} `mapstructure:"jwt"`
		} `mapstructure:"admin"`
		Cors struct {
			AllowOrigins     []string `mapstructure:"allow_origins"`
			AllowMethods     []string `mapstructure:"allow_methods"`
			AllowHeaders     []string `mapstructure:"allow_headers"`
			AllowCredentials bool     `mapstructure:"allow_credentials"`
		}
	} `mapstructure:"app"`

	Database struct {
		Host                  string        `mapstructure:"host"`
		Port                  string        `mapstructure:"port"`
		User                  string        `mapstructure:"user"`
		Password              string        `mapstructure:"password"`
		Name                  string        `mapstructure:"name"`
		Timeout               time.Duration `mapstructure:"timeout"`
		MaxConnections        int32         `mapstructure:"max_connections"`
		MinConnections        int32         `mapstructure:"min_connections"`
		MaxConnLifetime       time.Duration `mapstructure:"max_connection_lifetime"`
		MaxConnIdleTime       time.Duration `mapstructure:"max_connection_idle_time"`
		HealthCheckPeriod     time.Duration `mapstructure:"health_check_period"`
		MaxConnLifetimeJitter time.Duration `mapstructure:"max_connection_lifetime_jitter"`
	} `mapstructure:"database"`

	Logging struct {
		Level                LogLevel  `mapstructure:"level"`
		Format               LogFormat `mapstructure:"format"`
		File                 string    `mapstructure:"file"`
		MaxSize              int       `mapstructure:"max_size"`
		MaxBackups           int       `mapstructure:"max_backups"`
		MaxAge               int       `mapstructure:"max_age"`
		Compress             bool      `mapstructure:"compress"`
		EnableConsoleAndFile bool      `mapstructure:"enable_console_and_file"`
		Server               struct {
			RequestSymLimit  int `mapstructure:"request_sym_limit"`
			ResponseSymLimit int `mapstructure:"response_sym_limit"`
		} `mapstructure:"server"`
	} `mapstructure:"logging"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(kEnvPath); err != nil {
		log.Printf("⚠️ Could not load .env file: %v", err)
	}

	viper.SetConfigFile(kConfigPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("❌ Failed to load service configuration: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("❌ Failed to parse service configuration: %v", err)
	}

	log.Println("✅ Service configuration loaded successfully")

	return &cfg, nil
}
