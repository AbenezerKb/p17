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
	ListAllClient(ctx context.Context, params *rest.QueryParams) ([]dto.Client, error)
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

func (c clientStorage) ListAllClient(ctx context.Context, params *rest.QueryParams) ([]dto.Client, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)
	//ListAllClients(ctx context.Context, arg ListAllClientsParams)
	clients, err := c.ListAllClients(ctx, ListAllClientsParams{
		Offset: resizedPage,
		Limit:  resizedPerPage,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	return clients, nil
}

const listAllClients = `-- name: ListAllClients :many
SELECT id, title, phone, email, password, status, created_at, updated_at FROM clients
LIMIT $1
OFFSET $2
`

type ListAllClientsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (c clientStorage) ListAllClients(ctx context.Context, arg ListAllClientsParams) ([]dto.Client, error) {
	rows, err := c.db.Query(ctx, listAllClients, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []dto.Client{}
	for rows.Next() {
		var i dto.Client
		if err := rows.Scan(
			&i.Id,
			&i.ClientTitle,
			&i.Phone,
			&i.Email,
			&i.Password,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (c clientStorage) GetClient(ctx context.Context, email string) (*dto.Client, error) {

	cl, err := c.dbp.GetClient(ctx, email)

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
