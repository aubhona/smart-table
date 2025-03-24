package config

import "errors"

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
	JSONFormat    LogFormat = "json"
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
	case JSONFormat, ConsoleFormat:
		*lf = LogFormat(value)
		return nil
	default:
		return errors.New("unknown value log format: " + value)
	}
}
