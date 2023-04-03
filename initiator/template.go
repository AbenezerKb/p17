package initiator

//
//import (
//	"github.com/gin-gonic/gin"
//	handler "sms-gateway/internal/adapter/http/rest/template"
//	storage "sms-gateway/internal/adapter/storage/persistance/template"
//	const_init "sms-gateway/internal/constant/init"
//	"sms-gateway/internal/glue/template"
//	module "sms-gateway/internal/module/template"
//)
//
//func TemplateDomainInit(router *gin.RouterGroup, common const_init.Utils) {
//
//	templateStorage := storage.TemplateStorageInit(common)
//	templateModule := module.TemplateInit(templateStorage, common)
//	templateHandler := handler.TemplateHandlerInit(templateModule, common)
//
//	template.TemplateRouterInit(router, templateHandler)
//}
