package user

import (
	"context"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"sms-gateway/internal/adapter/storage/persistance/user"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
)

type userModule struct {
	userStorage user.UserStorage
	validate    *validator.Validate
	trans       ut.Translator
}

type UserModule interface {
	AddUser(ctx context.Context, user *dto.User) (*dto.User, error)
	GetUser(ctx context.Context, userId string) (*dto.User, error)
	UpdateUser(ctx context.Context, user *dto.User) (*dto.User, error)
	GetAllUsers(ctx context.Context, params *rest.QueryParams) ([]dto.User, error)
}

func UserInit(userStorage user.UserStorage, utils const_init.Utils) UserModule {
	return userModule{
		userStorage: userStorage,
		validate:    utils.GoValidator,
		trans:       utils.Translator,
	}
}

func (u userModule) AddUser(ctx context.Context, user *dto.User) (*dto.User, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	newUser, err := u.userStorage.AddUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u userModule) GetUser(ctx context.Context, userId string) (*dto.User, error) {
	user_, err := u.userStorage.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user_, nil
}

func (u userModule) UpdateUser(ctx context.Context, user *dto.User) (*dto.User, error) {
	updatedUser, err := u.userStorage.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
func (u userModule) GetAllUsers(ctx context.Context, params *rest.QueryParams) ([]dto.User, error) {

	clients, err := u.userStorage.GetAllUsers(ctx, params)
	if err != nil {
		return nil, err
	}
	return clients, nil
}
