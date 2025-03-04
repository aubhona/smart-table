package custom

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

const kConfigPath = "configs/config.yaml"
const kEnvPath = "configs/.env"

type Environment string
type LogLevel string
type LogFormat string

const (
	ProductionEnv  Environment = "production"
	DevelopmentEnv Environment = "development"
)

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
	FatalLevel LogLevel = "fatal"
)

const (
	JsonFormat    LogFormat = "json"
	ConsoleFormat LogFormat = "console"
)

func (e *Environment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}
	switch Environment(value) {
	case ProductionEnv, DevelopmentEnv:
		*e = Environment(value)
		return nil
	default:
		return errors.New("unknown value environment: " + value)
	}
}

func (l *LogLevel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}
	switch LogLevel(value) {
	case DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel:
		*l = LogLevel(value)
		return nil
	default:
		return errors.New("unknown value log level: " + value)
	}
}

func (lf *LogFormat) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}
	switch LogFormat(value) {
	case JsonFormat, ConsoleFormat:
		*lf = LogFormat(value)
		return nil
	default:
		return errors.New("unknown value log format: " + value)
	}
}

type Config struct {
	App struct {
		Port string      `yaml:"port"`
		Env  Environment `yaml:"env"`
	} `yaml:"app"`

	Database struct {
		Host                  string        `yaml:"host"`
		Port                  string        `yaml:"port"`
		User                  string        `yaml:"user"`
		Password              string        `yaml:"password"`
		Name                  string        `yaml:"name"`
		Timeout               time.Duration `yaml:"timeout"`
		MaxConnections        int           `yaml:"max_connections"`
		MinConnections        int           `yaml:"min_connections"`
		MaxConnLifetime       time.Duration `yaml:"max_connection_lifetime"`
		MaxConnIdleTime       time.Duration `yaml:"max_connection_idle_time"`
		HealthCheckPeriod     time.Duration `yaml:"health_check_period"`
		MaxConnLifetimeJitter time.Duration `yaml:"max_connection_lifetime_jitter"`
	} `yaml:"database"`

	Logging struct {
		Level                LogLevel  `yaml:"level"`
		Format               LogFormat `yaml:"format"`
		File                 string    `yaml:"file"`
		MaxSize              int       `yaml:"max_size"`
		MaxBackups           int       `yaml:"max_backups"`
		MaxAge               int       `yaml:"max_age"`
		Compress             bool      `yaml:"compress"`
		EnableConsoleAndFile bool      `yaml:"enable_console_and_file"`
	} `yaml:"logging"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(kEnvPath)
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(kConfigPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	cfg.Database.User = replaceEnvVariables(cfg.Database.User)
	cfg.Database.Password = replaceEnvVariables(cfg.Database.Password)

	return &cfg, nil
}

func replaceEnvVariables(input string) string {
	if strings.HasPrefix(input, "${") && strings.HasSuffix(input, "}") {
		envKey := strings.TrimPrefix(strings.TrimSuffix(input, "}"), "${")
		if value, exists := os.LookupEnv(envKey); exists {
			return value
		}
	}
	return input
}
