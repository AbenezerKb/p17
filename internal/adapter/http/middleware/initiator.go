package middleware

import (
	clientModule "sms-gateway/internal/module/client"
	messageModule "sms-gateway/internal/module/message"
	templateModule "sms-gateway/internal/module/template"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Modules struct {
	messageModule messageModule.Module
	tempModule    templateModule.TemplateModule
	cltModule     clientModule.Module
}

func NewAuthorizer(enforcer *casbin.Enforcer, logger *zap.SugaredLogger) gin.HandlerFunc {

	return newAuthorizer(enforcer, logger)
}

func NewErrorHandler(msg messageModule.Module) gin.HandlerFunc {
	return ErrorHandling()
}

func GetUserHandler(id string, modules Modules) gin.HandlerFunc {

	return ExtractUserId(id, modules)
}
