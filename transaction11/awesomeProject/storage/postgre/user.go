package postgre

import (
	"awesomeProject/model"
	"context"
	"gorm.io/gorm"
)

type UserRepositry struct {
	DB *gorm.DB
}

func NewUserRepositry(DB *gorm.DB) *UserRepositry {
	return &UserRepositry{DB: DB}
}
func (r UserRepositry) CreateUser(ctx context.Context, item model.User) (int, error) {
	err := r.DB.Table("users").Create(&item).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

func (r UserRepositry) GetUser(ctx context.Context, ID int) (model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (r UserRepositry) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("email = ?", email).First(&res).Error
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (r UserRepositry) Auth(ctx context.Context, user model.User) error {
	//TODO implement me
	panic("implement me")
}

func (r UserRepositry) DeleteUser(ctx context.Context, ID int) error {
	return r.DB.WithContext(ctx).Delete(&model.User{}, ID).Error
}
