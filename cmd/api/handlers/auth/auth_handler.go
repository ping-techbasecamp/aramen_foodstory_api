package authHandler

import (
	authModel "aramen-api/cmd/api/model/auth"
	"aramen-api/pkg/errs"
	authService "aramen-api/srv/services/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	CreateUser(c *gin.Context)
	RequestOtp(c *gin.Context)
	Login(c *gin.Context)
	HealthCheck(c *gin.Context)
}

type AuthHandler struct {
	authService authService.AuthService
}

func NewHandler(authService authService.AuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}

func (h AuthHandler) CreateUser(c *gin.Context) {
	var req authModel.CreateUserModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errs.ErrorStructResponse(err))
		return
	}
	err := h.authService.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.ErrorStructResponse(err))
		return
	}
	c.Status(http.StatusCreated)
}

func (h AuthHandler) RequestOtp(c *gin.Context) {
	phoneNumber := c.Param("phoneNumber")
	fmt.Println(phoneNumber)
	ref, err := h.authService.RequestOtp(phoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.ErrorStructResponse(err))
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"reference": *ref,
	})
}

func (h AuthHandler) Login(c *gin.Context) {
	var req authModel.LoginModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errs.ErrorStructResponse(err))
		return
	}
	signedToken, err := h.authService.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.ErrorStructResponse(err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": *signedToken,
	})
}

func (h AuthHandler) HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}
