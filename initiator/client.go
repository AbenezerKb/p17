package initiator

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	client2 "sms-gateway/internal/adapter/http/client"
	clhandler "sms-gateway/internal/adapter/http/rest/client"
	clstorage "sms-gateway/internal/adapter/storage/persistance/client"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/glue/client"
	clmodule "sms-gateway/internal/module/client"

	mhandler "sms-gateway/internal/adapter/http/rest/message"
	mstorage "sms-gateway/internal/adapter/storage/persistance/message"
	mmodule "sms-gateway/internal/module/message"

	thandler "sms-gateway/internal/adapter/http/rest/template"
	tstorage "sms-gateway/internal/adapter/storage/persistance/template"
	tmodule "sms-gateway/internal/module/template"
)

func ClientDomainInit(casbin *casbin.Enforcer, router *gin.RouterGroup, common const_init.Utils, log *zap.SugaredLogger) {
	cli := client2.ClientInit(common)

	//client
	clientStorage := clstorage.InitStorage(common)
	clientModule := clmodule.InitModule(clientStorage, common)
	clientHandler := clhandler.InitHandler(clientModule, common)

	//message
	messageStorage := mstorage.InitStorage(common)
	messageModule := mmodule.InitModule(cli, messageStorage, common)
	messageHandler := mhandler.InitHandler(messageModule, common)

	//template
	templaeStorage := tstorage.TemplateStorageInit(common)
	templaeModule := tmodule.TemplateInit(templaeStorage, common)
	templaeHandler := thandler.TemplateHandlerInit(templaeModule, common)

	client.ClientRouterInit(casbin, router, clientHandler, templaeHandler, messageHandler, log)
}
