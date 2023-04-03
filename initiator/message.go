package initiator

//
//import (
//	"github.com/gin-gonic/gin"
//	handler "sms-gateway/internal/adapter/http/rest/message"
//	storage "sms-gateway/internal/adapter/storage/persistance/message"
//	const_init "sms-gateway/internal/constant/init"
//	module "sms-gateway/internal/module/message"
//	"sms-gateway/internal/glue/message"
//)
//
//func MessageDomainInit(router *gin.RouterGroup, common const_init.Utils) {
//
//	messageStorage := storage.MessageStorageInit(common)
//	messageModule := module.TemplateInit(messageStorage, common)
//	messageHandler := handler.MessageHandlerInit(messageModule, common)
//
//	message(router, messageHandler)
//}
