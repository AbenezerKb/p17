package message

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model"
	"sms-gateway/internal/constant/model/db"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
	"sms-gateway/internal/constant/rest/error_types"
	"strconv"
)

type messageStorage struct {
	db  *pgxpool.Pool
	dbp db.Queries
}

type MessageStorage interface {
	AddMessage(ctx context.Context, message *dto.Message) (*dto.Message, error)
	BatchOutGoingSMS(ctx context.Context, message *model.SMS) (*dto.Message, error)
	ListAllMessages(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error)
	GetMessagesBySender(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error)
	LastMonthMessagesPriceAndCount(ctx context.Context, clientId string) ([]model.MessageCount, error)
}

func MessageStorageInit(utils const_init.Utils) MessageStorage {
	return messageStorage{
		db:  utils.Conn,
		dbp: *db.New(utils.Conn),
	}
}

func (m messageStorage) AddMessage(ctx context.Context, message *dto.Message) (*dto.Message, error) {
	ms, err := m.dbp.AddMessage(ctx, db.AddMessageParams{
		message.SenderPhone,
		message.Content,
		message.Price,
		message.ReceiverPhone,
		message.MsgType,
		message.Status,
	})

	msg := dto.Message{
		ms.ID.String(),
		ms.ReceiverPhone,
		ms.SenderPhone,
		ms.Content,
		ms.Price,
		ms.Type,
		ms.Status,
		ms.CreatedAt,
	}
	if err != nil {
		return nil, error_types.GetDbError(err)

	}
	return &msg, nil
}

func (m messageStorage) ListAllMessages(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)
	ms, err := m.dbp.ListAllMessages(ctx, db.ListAllMessagesParams{
		resizedPage,
		resizedPerPage,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	var msg []dto.Message

	for _, v := range ms {
		msg = append(msg, dto.Message{
			v.ID.String(),
			v.ReceiverPhone,
			v.SenderPhone,
			v.Content,
			v.Price,
			v.Type,
			v.Status,
			v.CreatedAt,
		})
	}
	return msg, nil
}

func (m messageStorage) GetMessagesBySender(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)
	ms, err := m.dbp.GetMessagesBySender(ctx, db.GetMessagesBySenderParams{
		params.Filter,
		resizedPage,
		resizedPerPage,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	var msg []dto.Message

	for _, v := range ms {
		msg = append(msg, dto.Message{
			v.ID.String(),
			v.ReceiverPhone,
			v.SenderPhone,
			v.Content,
			v.Price,
			v.Type,
			v.Status,
			v.CreatedAt,
		})
	}
	return msg, nil
}

func (m messageStorage) BatchOutGoingSMS(ctx context.Context, message *model.SMS) (*dto.Message, error) {

	return nil, nil
}

const lastMonthMessagesPriceAndCount = `-- name: LastMonthMessagesPriceAndCount :many
SELECT  price, COUNT(id) as COUNT,
        SUM (price) AS sum
FROM public.messages
WHERE sender_phone=$1 AND "created_at" BETWEEN NOW() - INTERVAL '1 MONTH' AND NOW()
GROUP BY  price
`

type LastMonthMessagesPriceAndCountRow struct {
	Price decimal.Decimal `json:"price"`
	Count int64           `json:"count"`
	Sum   int64           `json:"sum"`
}

func (q messageStorage) LastMonthMessagesPriceAndCount(ctx context.Context, senderPhone string) ([]model.MessageCount, error) {
	rows, err := q.db.Query(ctx, lastMonthMessagesPriceAndCount, senderPhone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []model.MessageCount{}
	for rows.Next() {
		var i model.MessageCount
		if err := rows.Scan(&i.Price, &i.Count, &i.Sum); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
