package message

import (
	"context"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model/db"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
	"sms-gateway/internal/constant/rest/error_types"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type messageStorage struct {
	db  *pgxpool.Pool
	dbp db.Queries
}

type Storage interface {
	AddMessage(ctx context.Context, message *dto.Message) (*dto.Message, error)
	GetMessage(ctx context.Context, id string) (*dto.Message, error)
	BatchOutGoingSMS(ctx context.Context, message []dto.Message) ([]dto.Message, error)
	ListAllMessages(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error)
	GetMessagesBySender(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error)
}

func InitStorage(utils const_init.Utils) Storage {
	return messageStorage{
		db:  utils.Conn,
		dbp: *db.New(utils.Conn),
	}
}

// AddMessage stores single message during message sending
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
		Id:             ms.ID.String(),
		ClientId:       ms.ClientID,
		ReceiverPhone:  ms.ReceiverPhone,
		SenderPhone:    ms.SenderPhone,
		Content:        ms.Content,
		Price:          ms.Price,
		MsgType:        ms.Type,
		Status:         ms.Status,
		DeliveryStatus: ms.DeliveryStatus,
		CreatedAt:      ms.CreatedAt,
	}
	if err != nil {
		return nil, error_types.GetDbError(err)

	}
	return &msg, nil
}

// ListAllMessages lists all messages
func (m messageStorage) ListAllMessages(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)
	msgs, err := m.ListMessages(ctx, ListAllMessagesParams{
		resizedPage,
		resizedPerPage,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	return msgs, nil
}

const listAllMessages = `-- name: ListAllMessages :many
SELECT id, client_id, sender_phone, content, price, receiver_phone, type, status, delivery_status, created_at FROM public.messages
        LIMIT $1
OFFSET $2
`

type ListAllMessagesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (ms messageStorage) ListMessages(ctx context.Context, arg ListAllMessagesParams) ([]dto.Message, error) {
	rows, err := ms.db.Query(ctx, listAllMessages, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []dto.Message{}
	for rows.Next() {
		var i dto.Message
		if err := rows.Scan(
			&i.Id,
			&i.ClientId,
			&i.SenderPhone,
			&i.Content,
			&i.Price,
			&i.ReceiverPhone,
			&i.MsgType,
			&i.Status,
			&i.DeliveryStatus,
			&i.CreatedAt,
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

// MessagesBySender lists all sender messages
func (ms messageStorage) GetMessagesBySender(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)
	msgs, err := ms.MessagesBySender(ctx, GetMessagesBySenderParams{
		params.Filter,
		resizedPage,
		resizedPerPage,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	return msgs, nil

}

const getMessagesBySender = `-- name: GetMessagesBySender :many
SELECT id, client_id, sender_phone, content, price, receiver_phone, type, status, delivery_status, created_at FROM public.messages
WHERE sender_phone=$1
LIMIT $2
    OFFSET $3
`

type GetMessagesBySenderParams struct {
	SenderPhone string `json:"sender_phone"`
	Limit       int32  `json:"limit"`
	Offset      int32  `json:"offset"`
}

func (ms messageStorage) MessagesBySender(ctx context.Context, arg GetMessagesBySenderParams) ([]dto.Message, error) {
	rows, err := ms.db.Query(ctx, getMessagesBySender, arg.SenderPhone, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []dto.Message{}
	for rows.Next() {
		var i dto.Message
		if err := rows.Scan(
			&i.Id,
			&i.ClientId,
			&i.SenderPhone,
			&i.Content,
			&i.Price,
			&i.ReceiverPhone,
			&i.MsgType,
			&i.Status,
			&i.DeliveryStatus,
			&i.CreatedAt,
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

// BatchOutGoingSMS sends messages in batch
func (m messageStorage) BatchOutGoingSMS(ctx context.Context, message []dto.Message) ([]dto.Message, error) {

	return nil, nil
}

func (mm messageStorage) GetMessage(ctx context.Context, id string) (*dto.Message, error) {
	uid, _ := uuid.Parse(id)
	msg, err := mm.dbp.GetMessageById(ctx, uid)

	if err != nil {
		return nil, err
	}

	message := dto.Message{
		Id:             msg.ID.String(),
		ClientId:       msg.ClientID,
		ReceiverPhone:  msg.ReceiverPhone,
		SenderPhone:    msg.SenderPhone,
		Content:        msg.Content,
		Price:          msg.Price,
		MsgType:        msg.Type,
		Status:         msg.Status,
		DeliveryStatus: msg.DeliveryStatus,
		CreatedAt:      msg.CreatedAt,
	}
	return &message, nil

}
