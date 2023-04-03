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
	messageStorage message.Storage
	validate       *validator.Validate
	trans          ut.Translator
}

type Module interface {
	OutGoingSMS(ctx context.Context, message *dto.Message) (*dto.Message, error)
	BatchOutGoingSMS(ctx context.Context, message *model.SMS) (*client.Response, error)
	IncomingSMS(ctx context.Context, message *dto.Message) (*dto.Message, error)
	GetAllClientMessages(ctx context.Context, params *rest.QueryParams) ([]dto.Message, error)
	GetMessage(ctx context.Context, id string) (*dto.Message, error)
}

func InitModule(jasmin client.JasminClient, messageStorage message.Storage, utils const_init.Utils) Module {
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

func (m messageModule) BatchOutGoingSMS(ctx context.Context, sms *model.SMS) (*client.Response, error) {

	msg := model.SMS{
		// To:message[0]
	}
	//TODO sending to jasmin
	resp, err := m.jasmin.BatchOutGoingSMS(ctx, &msg)

	if err != nil {
		return nil, err
	}

	//TODO save bulk sms to persistance
	// _, err = m.messageStorage.BatchOutGoingSMS(ctx, sms)

	// if err != nil {
	// 	return nil, err
	// }

	return resp, nil
}

func (m messageModule) GetMessage(ctx context.Context, id string) (*dto.Message, error) {
	// TODO: MODULE IMPLEMENTATION
	return nil, nil
}
