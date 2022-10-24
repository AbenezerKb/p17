package balance

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/joomcode/errorx"
	"golang.org/x/net/context"
	"net/http"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
	"sms-gateway/internal/constant/rest/error_types"
	errors "sms-gateway/internal/constant/rest/error_types"
	balanceModule "sms-gateway/internal/module/balance"
	messageModule "sms-gateway/internal/module/message"
	"time"
)

type balanceHandler struct {
	balanceModules balanceModule.Module
	validate       *validator.Validate
	trans          ut.Translator
}

type Handler interface {
	SendSMS(c *gin.Context)
	SendBatchSMS(c *gin.Context)
	ReceiveSMS(c *gin.Context)
	GetAllClientMessages(c *gin.Context)
}

var messageSuccessReplay = "Message Sent Successfully"

func HandlerInit(messageModules messageModule.MessageModule, utils const_init.Utils) Handler {
	return balanceHandler{
		messageModules: messageModules,
		validate:       utils.GoValidator,
		trans:          utils.Translator,
	}
}

func (m balanceHandler) SendSMS(c *gin.Context) {

	clientId := c.Param("id")
	value, _ := c.Get("userID")
	if clientId != value {
		err := errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
		_ = c.Error(err)

		return
	}
	message := &model.OutGoingSMS{}
	err := c.ShouldBind(message)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)
	msg := &dto.Message{
		ReceiverPhone: message.To,
		SenderPhone:   c.Param("id"),
		Content:       message.Content,
	}
	response, err := m.messageModules.OutGoingSMS(ctx, msg)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, response, http.StatusCreated)
}

func (m balanceHandler) SendBatchSMS(c *gin.Context) {

	clientId := c.Param("id")
	value, _ := c.Get("userID")
	if clientId != value {
		err := errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
		_ = c.Error(err)

		return
	}
	message := &model.SMS{}
	err := c.ShouldBind(message)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	response, err := m.messageModules.BatchOutGoingSMS(ctx, message)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, response, http.StatusCreated)
}

func (m balanceHandler) ReceiveSMS(c *gin.Context) {
	message := &dto.Message{}
	err := c.ShouldBind(message)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	message, err = m.messageModules.IncomingSMS(ctx, message)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, messageSuccessReplay, http.StatusCreated)

}

func (m balanceHandler) GetAllClientMessages(c *gin.Context) {

	clientId := c.Param("id")
	value, _ := c.Get("userID")

	if clientId != value {
		err := errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
		_ = c.Error(err)
		return
	}

	params := &rest.QueryParams{
		Page:    c.Query("page"),
		PerPage: c.Query("per_page"),
		Filter:  clientId,
	}

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	messages, err := m.messageModules.GetAllClientMessages(ctx, params)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, messages, http.StatusCreated)

}
