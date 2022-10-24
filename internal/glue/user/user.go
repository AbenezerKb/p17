package user

import (
	"github.com/gin-gonic/gin"
	"sms-gateway/internal/adapter/http/rest/user"
)

func UserRouterInit(router *gin.RouterGroup, usersHandler user.UserHandler) gin.IRoutes {

	router.POST("/users", usersHandler.AddUser)
	router.GET("/users", usersHandler.GetAllUsers)
	router.PATCH("/users", usersHandler.UpdateUser)
	router.GET("/users/:id", usersHandler.GetUser)

	return router
}
