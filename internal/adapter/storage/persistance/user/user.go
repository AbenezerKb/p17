package user

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model/db"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
	"strconv"
)

type userStorge struct {
	db  *pgxpool.Pool
	dbp db.Queries
}

type UserStorage interface {
	AddUser(ctx context.Context, user *dto.User) (*dto.User, error)
	UpdateUser(ctx context.Context, user *dto.User) (*dto.User, error)
	GetUser(ctx context.Context, phone string) (*dto.User, error)
	GetAllUsers(ctx context.Context, params *rest.QueryParams) ([]dto.User, error)
}

func UserStorageInit(utils const_init.Utils) UserStorage {
	return userStorge{
		utils.Conn,
		*db.New(utils.Conn),
	}
}

func (u userStorge) AddUser(ctx context.Context, user *dto.User) (*dto.User, error) {
	us, err := u.dbp.AddUser(ctx, db.AddUserParams{
		FullName: user.FullName,
		Phone:    user.Phone,
		Password: user.Password,
	})

	if err != nil {
		return nil, err
	}

	newUser := dto.User{
		Id:       us.ID.String(),
		FullName: us.FullName,
		Phone:    us.Phone,
	}

	return &newUser, nil
}
func (u userStorge) UpdateUser(ctx context.Context, user *dto.User) (*dto.User, error) {
	us, err := u.dbp.UpdateUser(ctx, db.UpdateUserParams{
		FullName: user.FullName,
	})
	if err != nil {
		return nil, err
	}

	newUser := dto.User{
		Id:       us.ID.String(),
		FullName: us.FullName,
		Phone:    us.Phone,
	}

	return &newUser, nil

}
func (u userStorge) GetUser(ctx context.Context, phone string) (*dto.User, error) {
	us, err := u.dbp.GetUser(ctx, phone)
	if err != nil {
		return nil, err
	}

	newUser := dto.User{
		Id:       us.ID.String(),
		FullName: us.FullName,
		Phone:    us.Phone,
	}

	return &newUser, nil

}
func (u userStorge) GetAllUsers(ctx context.Context, params *rest.QueryParams) ([]dto.User, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)
	tm, err := u.dbp.ListUsers(ctx, db.ListUsersParams{
		resizedPage,
		resizedPerPage,
	})

	if err != nil {
		return nil, err
	}

	var tmp []dto.User

	for _, v := range tm {
		tmp = append(tmp, dto.User{
			Id:       v.ID.String(),
			Phone:    v.Phone,
			FullName: v.FullName,
		})
	}
	return tmp, nil
}
