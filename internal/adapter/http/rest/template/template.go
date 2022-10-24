package template

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/joomcode/errorx"
	"golang.org/x/net/context"
	"net/http"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
	"sms-gateway/internal/constant/rest/error_types"
	errors "sms-gateway/internal/constant/rest/error_types"
	templateModule "sms-gateway/internal/module/template"
	"time"
)

type templateHandler struct {
	templateModules templateModule.TemplateModule
	validate        *validator.Validate
	trans           ut.Translator
}

type TemplateHandler interface {
	AddTemplate(c *gin.Context)
	UpdateTemplate(c *gin.Context)
	GetAllTemplates(c *gin.Context)
	GetAllClientTemplates(c *gin.Context)
	GetTemplate(c *gin.Context)
}

func TemplateHandlerInit(templateModules templateModule.TemplateModule, utils const_init.Utils) TemplateHandler {
	return templateHandler{
		templateModules: templateModules,
		validate:        utils.GoValidator,
		trans:           utils.Translator,
	}
}

func (t templateHandler) AddTemplate(c *gin.Context) {
	temp := &dto.TemplateMessage{}

	err := c.ShouldBind(temp)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}

	value, ok := c.Get("userID")
	if !ok {
		err := errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
		_ = c.Error(err)
	}

	template := &dto.Template{
		Template: temp.Template,
		Category: temp.Category,
		Client:   fmt.Sprint(value),
	}

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	template, err = t.templateModules.AddTemplate(ctx, template)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, &template, http.StatusCreated)
}

func (t templateHandler) UpdateTemplate(c *gin.Context) {

	value, ok := c.Get("userID")
	if !ok {
		err := errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
		_ = c.Error(err)
	}
	temp := &dto.TemplateMessage{}
	err := c.ShouldBind(temp)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}

	template := &dto.Template{
		Client:   fmt.Sprint(value),
		Template: temp.Template,
		Category: temp.Category,
	}
	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	template, err = t.templateModules.UpdateTemplate(ctx, template)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, &template, http.StatusCreated)

}

func (t templateHandler) GetAllTemplates(c *gin.Context) {

}

func (t templateHandler) GetAllClientTemplates(c *gin.Context) {

	value, ok := c.Get("userID")
	if !ok {
		err := errors.ErrInvalidToken.New(errors.ErrorUnauthorizedError)
		_ = c.Error(err)
	}

	params := &rest.QueryParams{}

	err := c.ShouldBindQuery(params)
	if err != nil {
		_ = c.Error(errorx.IllegalArgument.New(error_types.ErrorInvalidRequestBody))
		return
	}

	params.Filter = fmt.Sprint(value)

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	templates, err := t.templateModules.GetAllClientTemplate(ctx, params)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rest.SuccessResponseJson(c, nil, templates, http.StatusCreated)
}

func (t templateHandler) GetTemplate(c *gin.Context) {
	//
	//ctx := c.Request.Context()
	//templateid := c.Param("templateid")
	//ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*50))
	//defer cancel()
	//
	//c.Request = c.Request.WithContext(ctx)
	//
	//client, err := t.templateModules..GetTemplate(ctx, templateid)
	//if err != nil {
	//	c.Error(err)
	//	return
	//}
	//
	//rest.SuccessResponseJson(c, nil, &client, http.StatusOK)
}
