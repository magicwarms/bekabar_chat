package user

import (
	"bekabar_chat/apps/user/entity"
	"bekabar_chat/apps/user/model"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// NewService is used to create a single instance of the service
func NewUserService(connection *gorm.DB) *UserService {
	return &UserService{
		DB: connection,
	}
}

func (srv *UserService) StoreUser(user *entity.AddUserRequestDTO) error {

	// Handle errors directly from the transaction
	if err := srv.DB.Transaction(func(tx *gorm.DB) error {

		storeUser := model.UserModel{
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
		}

		return tx.Create(&storeUser).Error
	}); err != nil {
		return err
	}

	return nil
}

func (srv *UserService) FetchAllUser() ([]*model.UserModel, error) {

	var users []*model.UserModel

	results := srv.DB.Select("id", "email", "username", "created_at", "updated_at").Find(&users)

	if results.Error != nil {
		return nil, results.Error
	}

	return users, nil
}
