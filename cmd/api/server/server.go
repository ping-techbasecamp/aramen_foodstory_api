package server

import (
	"aramen-api/cmd/api/handlers"
	"aramen-api/cmd/api/middleware"
	"aramen-api/cmd/api/router"
	"aramen-api/pkg/config"
	"aramen-api/srv/logs"
	"fmt"

	srvgin "aramen-api/srv/gin"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srvGin *srvgin.Gin
	config *config.AppConfig
}

func (s *Server) IsHTTPRunning() bool {
	return s.srvGin.IsRunning()
}

func (s *Server) HTTPPort() int {
	return s.srvGin.Port()
}

func (s *Server) Run() {
	go (func() {
		addr := fmt.Sprintf("%s:%s", s.config.Env.Host, s.config.Env.Port)
		s.srvGin.Run(addr)
	})()
}

func (s *Server) Stop() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error().Msgf("found error during stop, do not raise panic %+v", r)
		}
	}()
	s.srvGin.ShutdownGracefully()
}

func NewServer(cfg *config.AppConfig, handler *handlers.Handler, middleware *middleware.ServiceMiddleware) (*Server, func()) {
	srvGin := srvgin.NewServer()

	if cfg.Env.GinMode != "" {
		gin.SetMode(cfg.Env.GinMode)
	}

	router.RegisterGinRouter(srvGin.Engine, handler, middleware, cfg)

	s := Server{
		srvGin: srvGin,
		config: cfg,
	}

	return &s, s.Stop
}
