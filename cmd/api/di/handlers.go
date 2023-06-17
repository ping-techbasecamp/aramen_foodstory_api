package di

import (
	"aramen-api/cmd/api/handlers"
	authHandler "aramen-api/cmd/api/handlers/auth"
	branchHandler "aramen-api/cmd/api/handlers/branch"
	newsHandler "aramen-api/cmd/api/handlers/news"
	qrcodeHandler "aramen-api/cmd/api/handlers/qrcode"
	"aramen-api/srv/services"
)

func newHandlers(service *services.Service) *handlers.Handler {
	return &handlers.Handler{
		AuthHandler:   authHandler.NewHandler(service.AuthService),
		BranchHandler: branchHandler.NewHandler(service.BranchService),
		NewsHandler:   newsHandler.NewHandler(service.NewsService),
		QrcodeHandler: qrcodeHandler.NewQrcodeHandler(service.QrcodeService),
	}
}
