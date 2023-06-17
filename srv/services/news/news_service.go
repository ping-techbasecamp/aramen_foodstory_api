package newsService

import (
	"aramen-api/pkg/config"
	newsRepo "aramen-api/srv/repositories/news"
)

type NewsServiceInterface interface {
}

type NewsService struct {
	cfg      *config.AppConfig
	newsRepo newsRepo.NewsRepo
}

func NewService(cfg *config.AppConfig, newsRepo newsRepo.NewsRepo) NewsService {
	return NewsService{cfg: cfg, newsRepo: newsRepo}
}
