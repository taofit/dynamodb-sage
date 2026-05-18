package engine

import (
	"os"

	"gopkg.in/yaml.v3"
)

type TableConfig struct {
	Name          string            `yaml:"name"`
	PIIFields     []string          `yaml:"pii_fields"`
	ReadOnly      bool              `yaml:"read_only"`
	EnforceSchema bool              `yaml:"enforce_schema"`
	Columns       map[string]string `yaml:"columns"`
}

type GlobalLimits struct {
	MaxLimit     int32 `yaml:"max_limit"`
	DefaultLimit int32 `yaml:"default_limit"`
}

type AppConfig struct {
	GlobalLimits    GlobalLimits  `yaml:"global_limits"`
	SensitiveFields []string      `yaml:"sensitive_fields"`
	ProtectedTables []string      `yaml:"protected_tables"`
	Tables          []TableConfig `yaml:"tables"`
}

func LoadConfig(filename string) (*AppConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config AppConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	if config.GlobalLimits.MaxLimit <= 0 {
		config.GlobalLimits.MaxLimit = 100
	}

	if config.GlobalLimits.DefaultLimit <= 0 {
		config.GlobalLimits.DefaultLimit = 20
	}

	return &config, nil
}
