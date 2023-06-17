package di

import (
	"aramen-api/cmd/api/middleware"
	"aramen-api/cmd/api/server"
	"aramen-api/pkg/config"
	"aramen-api/pkg/db"
)

func InitializeApplication() (*server.Server, func(), error) {
	AppConfig := config.NewAppConfig()

	db, cleanup1, err := db.NewMySQLConnection(AppConfig)
	if err != nil {
		return nil, nil, err
	}

	serverServer, cleanup2 := server.NewServer(
		AppConfig,
		newHandlers(newServices(AppConfig, newRepositories(db))),
		middleware.NewService(AppConfig),
	)

	return serverServer, func() {
		cleanup1()
		cleanup2()
	}, nil
}
