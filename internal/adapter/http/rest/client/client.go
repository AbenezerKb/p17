package client

import (
	"context"
	"fmt"
	"github.com/dongri/phonenumber"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/joomcode/errorx"
	"net/http"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
	"sms-gateway/internal/constant/rest/error_types"
	errors "sms-gateway/internal/constant/rest/error_types"
	clientModule "sms-gateway/internal/module/client"
	"time"
)

type clientHandler struct {
	clientModules clientModule.ClientModule
	validate      *validator.Validate
	trans         ut.Translator
}

type ClientHandler interface {
	AddClient(c *gin.Context)
	UpdateClient(c *gin.Context)
	GetAllClients(c *gin.Context)
	GetClient(c *gin.Context)
	ClientLogin(c *gin.Context)
}

func ClientHandlerInit(clientModules clientModule.ClientModule, utils const_init.Utils) ClientHandler {
	return clientHandler{
		clientModules: clientModules,
		validate:      utils.GoValidator,
		trans:         utils.Translator,
	}
}

func (cl clientHandler) AddClient(c *gin.Context) {
	client := &dto.Client{}
	err := c.ShouldBind(client)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	client, err = cl.clientModules.AddClient(ctx, client)
	if err != nil {
		_ = c.Error(err)
		return
	}
	fmt.Println("the client: ", client)

	rest.SuccessResponseJson(c, nil, &client, http.StatusCreated)
}

func (cl clientHandler) UpdateClient(c *gin.Context) {
	clientId := c.Param("id")
	value, ok := c.Get("userID")
	if ok && clientId != value {
		err := errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
		_ = c.Error(err)
	}

	client := &dto.Client{}
	err := c.ShouldBind(client)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}
	client.Id = clientId
	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	client, err = cl.clientModules.UpdateClient(ctx, client)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, &client, http.StatusCreated)
}

func (cl clientHandler) GetAllClients(c *gin.Context) {

	params := rest.QueryParams{
		Page:    c.Query("page"),
		PerPage: c.Query("per_page"),
	}

	ctx := c.Request.Context()

	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	clients, err := cl.clientModules.GetAllClients(ctx, &params)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, &clients, http.StatusOK)
}

func (cl clientHandler) GetClient(c *gin.Context) {

	ctx := c.Request.Context()
	clientId := c.Param("id")
	value, ok := c.Get("userID")
	if ok && clientId != value {
		err := errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
		_ = c.Error(err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	client, err := cl.clientModules.GetClient(ctx, clientId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, &client, http.StatusOK)
}

func (cl clientHandler) ClientLogin(c *gin.Context) {
	clientLogin := &model.ClientLogin{}

	err := c.ShouldBindJSON(clientLogin)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}

	phone := phonenumber.Parse(clientLogin.Phone, "ET")

	if phone == "" {
		_ = c.Error(errorx.IllegalArgument.New("invalid phone"))
		return
	}
	clientLogin.Phone = phone

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	usr, err := cl.clientModules.Login(ctx, clientLogin)
	if err != nil {
		_ = c.Error(err)
		return
	}
	token := `{token:%s}`
	rest.SuccessResponseJson(c, nil, fmt.Sprintf(token, *usr), http.StatusOK)
}
