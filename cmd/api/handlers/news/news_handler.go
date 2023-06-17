package newsHandler

import newsService "aramen-api/srv/services/news"

type NewsHandlerInterface interface{}

type NewsHandler struct {
	newsService newsService.NewsService
}

func NewHandler(newsService newsService.NewsService) NewsHandler {
	return NewsHandler{newsService: newsService}
}
