package message

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
	"sms-gateway/internal/adapter/http/client"
	"sms-gateway/internal/adapter/storage/persistance/message"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
	"sms-gateway/internal/constant/sms_type"
)

type messageModule struct {
	jasmin         client.JasminClient
	messageStorage message.MessageStorage
	validate       *validator.Validate
	trans          ut.Translator
}

type MessageModule interface {
	OutGoingSMS(ctx context.Context, message *dto.Message) (*dto.Message, error)
	BatchOutGoingSMS(ctx context.Context, message *model.SMS) (*client.Response, error)
	IncomingSMS(ctx context.Context, message *dto.Message) (*dto.Message, error)
	GetAllClientMessages(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error)
}

func MessageInit(jasmin client.JasminClient, messageStorage message.MessageStorage, utils const_init.Utils) MessageModule {
	return messageModule{
		jasmin:         jasmin,
		messageStorage: messageStorage,
		validate:       utils.GoValidator,
		trans:          utils.Translator,
	}
}

func (m messageModule) OutGoingSMS(ctx context.Context, message *dto.Message) (*dto.Message, error) {

	err := message.MsgType.Scan(sms_type.OutGoing)
	if err != nil {
		return nil, err
	}

	sms := &model.OutGoingSMS{
		To:      message.ReceiverPhone,
		Content: message.Content,
	}
	_, err = m.jasmin.OutGoingSMS(ctx, sms)

	if err != nil {
		return nil, err
	}
	message, err = m.messageStorage.AddMessage(ctx, message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (m messageModule) IncomingSMS(ctx context.Context, message *dto.Message) (*dto.Message, error) {

	err := message.MsgType.Scan(sms_type.Incoming)
	if err != nil {
		return nil, err
	}

	message, err = m.messageStorage.AddMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (m messageModule) GetAllClientMessages(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error) {

	messages, err := m.messageStorage.GetMessagesBySender(ctx, params)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (m messageModule) BatchOutGoingSMS(ctx context.Context, message *model.SMS) (*client.Response, error) {

	resp, err := m.jasmin.BatchOutGoingSMS(ctx, message)

	if err != nil {
		return nil, err
	}

	_, err = m.messageStorage.BatchOutGoingSMS(ctx, message)

	if err != nil {
		return nil, err
	}

	return resp, nil
}