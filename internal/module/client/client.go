package client

import (
	"context"
	"github.com/dongri/phonenumber"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/joomcode/errorx"
	"sms-gateway/internal/adapter/storage/persistance/client"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
)

type clientModule struct {
	clientStorage client.ClientStorge
	validate      *validator.Validate
	trans         ut.Translator
}

type Module interface {
	AddClient(ctx context.Context, client *dto.Client) (*dto.Client, error)
	UpdateClient(ctx context.Context, client *dto.Client) (*dto.Client, error)
	GetAllClients(ctx context.Context, params *rest.QueryParams) ([]dto.Client, error)
	GetClient(ctx context.Context, clientId string) (*dto.Client, error)
	Login(ctx context.Context, clientLogin *model.ClientLogin) (*string, error)
}

func InitModule(clientStorage client.ClientStorge, utils const_init.Utils) Module {
	return clientModule{
		clientStorage: clientStorage,
		validate:      utils.GoValidator,
		trans:         utils.Translator,
	}
}

func (c clientModule) AddClient(ctx context.Context, client *dto.Client) (*dto.Client, error) {

	err := client.Validate()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(client.Phone); i++ {
		client.Phone[i] = phonenumber.Parse(client.Phone[i], "ET")
		if client.Phone[i] == "" {
			return nil, errorx.IllegalArgument.New("incorrect phone number format")
		}
	}

	newClient, err := c.clientStorage.AddClient(ctx, client)
	if err != nil {
		return nil, err
	}

	return newClient, nil
}

func (c clientModule) UpdateClient(ctx context.Context, client *dto.Client) (*dto.Client, error) {
	updatedClient, err := c.clientStorage.UpdateClient(ctx, client)
	if err != nil {
		return nil, err
	}
	return updatedClient, nil
}

func (c clientModule) GetAllClients(ctx context.Context, params *rest.QueryParams) ([]dto.Client, error) {

	clients, err := c.clientStorage.ListAllClient(ctx, params)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (c clientModule) GetClient(ctx context.Context, phone string) (*dto.Client, error) {
	clt, err := c.clientStorage.GetClient(ctx, phone)
	if err != nil {
		return nil, err
	}
	return clt, nil
}

func (c clientModule) Login(ctx context.Context, clientLogin *model.ClientLogin) (*string, error) {
	var customClaims model.CustomClaims

	token, err := customClaims.GenerateToken(clientLogin.Email)
	if err != nil {

		//TODO ERROR HANDLING
		return nil, err
	}

	return token, nil
}
