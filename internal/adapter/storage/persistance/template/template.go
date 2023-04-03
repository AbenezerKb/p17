package template

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

type templateStorage struct {
	db  *pgxpool.Pool
	dbp db.Queries
}

type TemplateStorage interface {
	AddTemplate(ctx context.Context, template *dto.Template) (*dto.Template, error)
	UpdateTemplate(ctx context.Context, template *dto.Template) (*dto.Template, error)
	GetAllTemplates(ctx context.Context, params *rest.QueryParams) ([]dto.Template, error)
	GetAllClientTemplates(ctx context.Context, params *rest.QueryParams) ([]dto.Template, error)
	GetTemplate(ctx context.Context, templateId string) (*dto.Template, error)
}

func TemplateStorageInit(utils const_init.Utils) TemplateStorage {
	return templateStorage{
		db:  utils.Conn,
		dbp: *db.New(utils.Conn),
	}
}

func (t templateStorage) AddTemplate(ctx context.Context, template *dto.Template) (*dto.Template, error) {

	tm, err := t.dbp.AddTemplate(ctx, db.AddTemplateParams{
		Client:   template.Client,
		Template: template.Template,
		Category: template.Category,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)

	}

	tmp := dto.Template{
		Id:       tm.Client,
		Template: tm.Template,
		Category: tm.Category,
	}

	return &tmp, nil
}

func (t templateStorage) UpdateTemplate(ctx context.Context, template *dto.Template) (*dto.Template, error) {
	tm, err := t.dbp.UpdateTemplate(ctx, db.UpdateTemplateParams{
		Client:   template.Client,
		Template: template.Template,
		Category: template.Category,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)

	}

	tmp := dto.Template{
		Id:       tm.Client,
		Template: tm.Template,
		Category: tm.Category,
	}

	return &tmp, nil
}
func (t templateStorage) GetAllTemplates(ctx context.Context, params *rest.QueryParams) ([]dto.Template, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)
	tm, err := t.ListAllTemplates(ctx, ListAllTemplatesParams{
		Limit:  resizedPerPage,
		Offset: resizedPage,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	return tm, nil
}

func (t templateStorage) GetAllClientTemplates(ctx context.Context, params *rest.QueryParams) ([]dto.Template, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)

	cl, err := t.dbp.ListClientTemplates(ctx, db.ListClientTemplatesParams{
		Client: params.Filter,
		Limit:  resizedPage,
		Offset: resizedPerPage,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	var cli []dto.Template

	for _, v := range cl {
		cli = append(cli, dto.Template{
			Id:       v.ID.String(),
			Template: v.Template,
			Client:   v.Client,
			Category: v.Category,
		})
	}
	return cli, nil
}

const listAllTemplates = `-- name: ListAllTemplates :many
SELECT id, template_id, client, template, category, created_at, updated_at FROM templates
LIMIT $1
OFFSET $2
`

type ListAllTemplatesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q templateStorage) ListAllTemplates(ctx context.Context, arg ListAllTemplatesParams) ([]dto.Template, error) {
	rows, err := q.db.Query(ctx, listAllTemplates, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []dto.Template{}
	for rows.Next() {
		var i dto.Template
		if err := rows.Scan(
			&i.Id,
			&i.TemplateID,
			&i.Client,
			&i.Template,
			&i.Category,
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

func (t templateStorage) GetTemplate(ctx context.Context, templateId string) (*dto.Template, error) {
	// TODO: PERSISTANCE IMPLEMENTATION
	return nil, nil
}
