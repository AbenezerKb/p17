package client

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model/db"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
	"sms-gateway/internal/constant/rest/error_types"
	"strconv"
)

type clientStorage struct {
	db  *pgxpool.Pool
	dbp db.Queries
}

type ClientStorge interface {
	AddClient(ctx context.Context, client *dto.Client) (*dto.Client, error)
	ListAllClients(ctx context.Context, params *rest.QueryParams) ([]dto.Client, error)
	GetClient(ctx context.Context, phone string) (*dto.Client, error)
	UpdateClient(ctx context.Context, client *dto.Client) (*dto.Client, error)
}

func ClientStorageInit(utils const_init.Utils) ClientStorge {
	return clientStorage{
		db:  utils.Conn,
		dbp: *db.New(utils.Conn),
	}
}

func (c clientStorage) AddClient(ctx context.Context, client *dto.Client) (*dto.Client, error) {

	cl, err := c.dbp.AddClient(ctx, db.AddClientParams{
		Title:    client.ClientTitle,
		Phone:    client.Phone,
		Email:    client.Email,
		Password: client.Password,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)

	}

	client = &dto.Client{
		Id:          cl.ID.String(),
		ClientTitle: cl.Title,
		Phone:       cl.Phone,
		Email:       cl.Email,
		Password:    cl.Password,
	}

	return client, nil
}

func (c clientStorage) ListAllClients(ctx context.Context, params *rest.QueryParams) ([]dto.Client, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)
	cl, err := c.dbp.ListAllClients(ctx, db.ListAllClientsParams{
		Offset: resizedPage,
		Limit:  resizedPerPage,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	var cli []dto.Client

	for _, v := range cl {
		cli = append(cli, dto.Client{
			Id:          v.ID.String(),
			ClientTitle: v.Title,
			Phone:       v.Phone,
			Email:       v.Email,
			Status:      v.Status,
			Password:    v.Password,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return cli, nil
}

func (c clientStorage) GetClient(ctx context.Context, phone string) (*dto.Client, error) {

	cl, err := c.dbp.GetClient(ctx, phone)

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	client := &dto.Client{
		Id:          cl.ID.String(),
		ClientTitle: cl.Title,
		Phone:       cl.Phone,
		Email:       cl.Email,
		Status:      cl.Status,
		Password:    cl.Password,
		CreatedAt:   cl.CreatedAt,
		UpdatedAt:   cl.UpdatedAt,
	}

	return client, nil
}

func (c clientStorage) UpdateClient(ctx context.Context, client *dto.Client) (*dto.Client, error) {

	updatedCl, err := c.dbp.UpdateClient(ctx, db.UpdateClientParams{
		Phone:    client.Phone,
		Title:    client.ClientTitle,
		Email:    client.Email,
		Password: client.Password,
		Status:   client.Status,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}
	client = &dto.Client{
		Id:          updatedCl.ID.String(),
		ClientTitle: updatedCl.Title,
		Phone:       updatedCl.Phone,
		Email:       updatedCl.Email,
		Status:      updatedCl.Status,
		Password:    updatedCl.Password,
		CreatedAt:   updatedCl.CreatedAt,
		UpdatedAt:   updatedCl.UpdatedAt,
	}

	return client, nil

}
