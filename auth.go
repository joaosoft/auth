package auth

import (
	"sync"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
	migration "github.com/joaosoft/migration/services"
)

type Auth struct {
	config        *AuthConfig
	isLogExternal bool
	pm            *manager.Manager
	logger        logger.ILogger
	mux           sync.Mutex
}

// NewAuth ...
func NewAuth(options ...AuthOption) (*Auth, error) {
	config, simpleConfig, err := NewConfig()

	service := &Auth{
		pm:     manager.NewManager(manager.WithRunInBackground(false)),
		logger: logger.NewLogDefault("auth", logger.WarnLevel),
		config: config.Auth,
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(logger.Instance))
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.Auth != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Auth.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	} else {
		config.Auth = &AuthConfig{
			Host: defaultURL,
		}
	}

	service.Reconfigure(options...)

	// execute migrations
	migrationService, err := migration.NewCmdService(migration.WithCmdConfiguration(service.config.Migration))
	if err != nil {
		return nil, err
	}

	if _, err := migrationService.Execute(migration.OptionUp, 0, migration.ExecutorModeDatabase); err != nil {
		return nil, err
	}

	web := service.pm.NewSimpleWebServer(config.Auth.Host)

	storage, err := newStoragePostgres(config.Auth)
	if err != nil {
		return nil, err
	}

	interactor := newModel(config.Auth, storage)

	controller := newController(config.Auth, interactor)
	controller.registerRoutes(web)

	service.pm.AddWeb("api_web", web)

	return service, nil
}

// Start ...
func (auth *Auth) Start() error {
	return auth.pm.Start()
}

// Stop ...
func (auth *Auth) Stop() error {
	return auth.pm.Stop()
}
