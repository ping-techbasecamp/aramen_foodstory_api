package router

import (
	"aramen-api/cmd/api/handlers"
	"aramen-api/cmd/api/middleware"
	"aramen-api/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterGinRouter(router *gin.Engine, h *handlers.Handler, middleware *middleware.ServiceMiddleware, cfg *config.AppConfig) {
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/auth/otp/:phoneNumber", h.AuthHandler.RequestOtp)
		apiV1.POST("/auth/create", h.AuthHandler.CreateUser)
		apiV1.POST("/auth/login", h.AuthHandler.Login)
		apiV1.GET("/auth/healthcheck", h.AuthHandler.HealthCheck, middleware.Auth())

		apiV1.GET("/branches", h.BranchHandler.GetBranches)
		apiV1.GET("/branches/:branchId", h.BranchHandler.GetBranchDetail)

		apiV1.POST("/scan", h.QrcodeHandler.ScanQr)
	}

	if cfg.Env.EnabledDevRoutes {
		setDevRoute(h, router.Group("/dev"))
	}
}
