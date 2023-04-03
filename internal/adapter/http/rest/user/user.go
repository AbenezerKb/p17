package user

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/joomcode/errorx"
	"golang.org/x/net/context"
	"net/http"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
	"sms-gateway/internal/constant/rest/error_types"
	userModule "sms-gateway/internal/module/user"
	"time"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userModules userModule.UserModule
	validate    *validator.Validate
	trans       ut.Translator
}

type UserHandler interface {
	AddUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUserByParam(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

func UserHandlerInit(userModules userModule.UserModule, utils const_init.Utils) UserHandler {
	return userHandler{
		userModules: userModules,
		validate:    utils.GoValidator,
		trans:       utils.Translator,
	}
}

func (u userHandler) AddUser(c *gin.Context) {
	user := &dto.User{}
	err := c.ShouldBind(user)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	user, err = u.userModules.AddUser(ctx, user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, &user, http.StatusCreated)
}
func (u userHandler) GetUser(c *gin.Context) {

	ctx := c.Request.Context()
	userid := c.Param("userid")
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	user, err := u.userModules.GetUser(ctx, userid)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, &user, http.StatusOK)
}
func (u userHandler) UpdateUser(c *gin.Context) {
	user := &dto.User{}
	err := c.ShouldBind(user)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	user, err = u.userModules.UpdateUser(ctx, user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, &user, http.StatusCreated)
}
func (u userHandler) GetAllUsers(c *gin.Context) {
	params := rest.QueryParams{
		Page:    c.Query("page"),
		PerPage: c.Query("per_page"),
	}

	ctx := c.Request.Context()

	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	users, err := u.userModules.GetAllUsers(ctx, &params)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, &users, http.StatusOK)
}

func (uh userHandler) GetUserByParam(c *gin.Context) {

}
