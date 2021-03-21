package auth

import (
	"fmt"

	"github.com/joaosoft/dbr"

	"github.com/joaosoft/manager"
	migration "github.com/joaosoft/migration/services"
)

// AppConfig ...
type AppConfig struct {
	Auth *AuthConfig `json:"auth"`
}

// AuthConfig ...
type AuthConfig struct {
	Host              string                     `json:"host"`
	Dbr               *dbr.DbrConfig             `json:"dbr"`
	TokenKey          string                     `json:"token_key"`
	ExpirationMinutes int64                      `json:"expiration_minutes"`
	Migration         *migration.MigrationConfig `json:"migration"`
	Log               struct {
		Level string `json:"level"`
	} `json:"log"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", getEnv()), appConfig)

	return appConfig, simpleConfig, err
}
