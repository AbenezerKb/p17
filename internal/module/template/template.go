package template

import (
	"context"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"sms-gateway/internal/adapter/storage/persistance/template"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
)

type templateModule struct {
	templateStorage template.TemplateStorage
	validate        *validator.Validate
	trans           ut.Translator
}

type TemplateModule interface {
	AddTemplate(ctx context.Context, template *dto.Template) (*dto.Template, error)
	UpdateTemplate(ctx context.Context, template *dto.Template) (*dto.Template, error)
	GetAllTemplate(ctx context.Context, params *rest.QueryParams) ([]dto.Template, error)
	GetAllClientTemplate(ctx context.Context, params *rest.QueryParams) ([]dto.Template, error)
	//GetTemplate(ctx context.Context, templateId string) (*dto.Template, error)
}

func TemplateInit(templateStorage template.TemplateStorage, utils const_init.Utils) TemplateModule {
	return templateModule{
		templateStorage: templateStorage,
		validate:        utils.GoValidator,
		trans:           utils.Translator,
	}
}

func (t templateModule) AddTemplate(ctx context.Context, template *dto.Template) (*dto.Template, error) {
	err := template.Validate()
	if err != nil {
		return nil, err
	}

	newTemplate, err := t.templateStorage.AddTemplate(ctx, template)
	if err != nil {
		return nil, err
	}
	return newTemplate, nil
}

func (t templateModule) UpdateTemplate(ctx context.Context, template *dto.Template) (*dto.Template, error) {

	newTemplate, err := t.templateStorage.UpdateTemplate(ctx, template)
	if err != nil {
		return nil, err
	}
	return newTemplate, nil
}

func (t templateModule) GetAllTemplate(ctx context.Context, params *rest.QueryParams) ([]dto.Template, error) {

	clients, err := t.templateStorage.GetAllTemplates(ctx, params)
	if err != nil {
		return nil, err
	}
	return clients, nil

}

func (t templateModule) GetAllClientTemplate(ctx context.Context, params *rest.QueryParams) ([]dto.Template, error) {
	templates, err := t.templateStorage.GetAllClientTemplates(ctx, params)
	if err != nil {
		return nil, err
	}
	return templates, nil

}

//func (t templateModule) GetTemplate(ctx context.Context, templateId string) (*dto.Template, error) {
//	template, err :=t.templateStorage.GetTemplate(ctx,templateId)
//if err != nil {
//return nil, err
//}
//return templates, nil
//}
