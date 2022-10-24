package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewAuthorizer(enforcer *casbin.Enforcer, logger *zap.SugaredLogger) gin.HandlerFunc {

	return newAuthorizer(enforcer, logger)
}

func NewErrorHandler() gin.HandlerFunc {
	return ErrorHandling()
}
