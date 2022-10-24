package balance

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
	"sms-gateway/internal/adapter/storage/persistance/balance"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
)

type balanceModule struct {
	balanceStorage balance.Storage
	validate       *validator.Validate
	trans          ut.Translator
}

type Module interface {
	AddBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error)
	ListAllBalance(ctx context.Context, params *rest.QueryParams) ([]dto.Balance, error)
	GetClientBalance(ctx context.Context, phone string) (*dto.Balance, error)
	UpdateClientBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error)
	GetClientInvoice(ctx context.Context, phone string) (*dto.Invoice, error)
}

func ModuleInit(balanceStorage balance.Storage, utils const_init.Utils) Module {
	return balanceModule{
		balanceStorage: balanceStorage,
		validate:       utils.GoValidator,
		trans:          utils.Translator,
	}
}

func (b balanceModule) AddBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error) {

	blc, err := b.balanceStorage.AddBalance(ctx, balance)
	if err != nil {
		return nil, err
	}

	return blc, nil
}

func (b balanceModule) ListAllBalance(ctx context.Context, params *rest.QueryParams) ([]dto.Balance, error) {
	balanceList, err := b.balanceStorage.ListAllBalance(ctx, params)
	if err != nil {
		return nil, err
	}

	return balanceList, nil
}

func (b balanceModule) GetClientBalance(ctx context.Context, phone string) (*dto.Balance, error) {

	blc, err := b.balanceStorage.GetClientBalance(ctx, phone)
	if err != nil {
		return nil, err
	}
	return blc, nil
}

func (b balanceModule) UpdateClientBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error) {
	blc, err := b.balanceStorage.UpdateClientBalance(ctx, balance)
	if err != nil {
		return nil, err
	}
	return blc, nil
}

func (b balanceModule) GetClientInvoice(ctx context.Context, phone string) (*dto.Invoice, error) {
	blc, err := b.balanceStorage.UpdateClientBalance(ctx, balance)
	if err != nil {
		return nil, err
	}
	return blc, nil
}
