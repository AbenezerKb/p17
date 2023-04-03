package client

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sms-gateway/internal/adapter/http/middleware"
	"sms-gateway/internal/adapter/http/rest/client"
	"sms-gateway/internal/adapter/http/rest/message"
	"sms-gateway/internal/adapter/http/rest/template"
)

func ClientRouterInit(enforcer *casbin.Enforcer, router *gin.RouterGroup, clientHandler client.Handler, templateHandler template.TemplateHandler, messageHandler message.Handler, log *zap.SugaredLogger) gin.IRoutes {

	router.POST("/clients", clientHandler.AddClient)
	router.POST("/clients/login", clientHandler.ClientLogin)
	router.GET("/clients", clientHandler.GetAllClients)
	router.PATCH("/clients/:id", middleware.NewAuthorizer(enforcer, log), clientHandler.UpdateClient)
	router.GET("/clients/:id", middleware.NewAuthorizer(enforcer, log), clientHandler.GetClient)
	router.GET("/clients/:id/messages", middleware.NewAuthorizer(enforcer, log), messageHandler.GetAllClientMessages)
	router.POST("/clients/:id/messages", middleware.NewAuthorizer(enforcer, log), messageHandler.SendSMS)

	router.POST("/clients/:id/incomingmessage", middleware.NewAuthorizer(enforcer, log), messageHandler.ReceiveSMS)
	router.POST("/clients/:id/templates", middleware.NewAuthorizer(enforcer, log), templateHandler.AddTemplate)
	router.GET("/clients/:id/templates", middleware.NewAuthorizer(enforcer, log), templateHandler.GetAllClientTemplates)
	router.PATCH("/clients/:id/templates/:id", middleware.NewAuthorizer(enforcer, log), templateHandler.UpdateTemplate)

	return router
}
