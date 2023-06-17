package middleware

import (
	"aramen-api/pkg/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	IDKey = "id"
)

type JWTClaim struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type ServiceMiddleware struct {
	cfg *config.AppConfig
}

func NewService(cfg *config.AppConfig) *ServiceMiddleware {
	return &ServiceMiddleware{
		cfg: cfg,
	}
}

func (s *ServiceMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		if bearer == "" {
			c.JSON(http.StatusUnauthorized, unauthorizedResponse(fmt.Errorf("Invalid authorization"), "getAuthorization"))
			c.Abort()
			return
		}

		auth := strings.Split(bearer, " ")
		if len(auth) != 2 || auth[0] != "Bearer" || auth[1] == "" {
			c.JSON(http.StatusUnauthorized, unauthorizedResponse(fmt.Errorf("Invalid token : %v", bearer), "getAuthorization"))
			c.Abort()
			return
		}

		var jwtClaim JWTClaim
		token, err := jwt.ParseWithClaims(auth[1], &jwtClaim, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.cfg.JWT.SignedKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, unauthorizedResponse(err, "ParseWithClaims"))
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, unauthorizedResponse(fmt.Errorf("JWT token is invalid"), "invalidToken"))
			c.Abort()
			return
		}
		id := jwtClaim.ID
		c.Set("id", id)
		c.Next()
	}
}
