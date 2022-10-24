package initiator

import (
	"github.com/gin-gonic/gin"
	handler "sms-gateway/internal/adapter/http/rest/user"
	storage "sms-gateway/internal/adapter/storage/persistance/user"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/glue/user"
	module "sms-gateway/internal/module/user"
)

func UserDomainInit(router *gin.RouterGroup, common const_init.Utils) {

	userStorage := storage.UserStorageInit(common)
	userModule := module.UserInit(userStorage, common)
	userHandler := handler.UserHandlerInit(userModule, common)

	user.UserRouterInit(router, userHandler)
}
