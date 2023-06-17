package di

import (
	"aramen-api/pkg/config"
	"aramen-api/srv/repositories"
	"aramen-api/srv/services"
	authService "aramen-api/srv/services/auth"
	branchService "aramen-api/srv/services/branch"
	newsService "aramen-api/srv/services/news"
	qrcodeService "aramen-api/srv/services/qrcode"
)

func newServices(cfg *config.AppConfig, repository *repositories.Repositories) *services.Service {
	return &services.Service{
		AuthService:   authService.NewService(cfg, repository.AuthRepo),
		BranchService: branchService.NewService(cfg, repository.BranchRepo),
		NewsService:   newsService.NewService(cfg, repository.NewsRepo),
		QrcodeService: qrcodeService.NewQrcodeService(repository.QrcodeRepo),
	}
}
