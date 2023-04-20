package postgre

import (
	"awesomeProject/model"
	"context"
	"gorm.io/gorm"
)

type CardRepository struct {
	DB *gorm.DB
}

func NewCardRepository(DB *gorm.DB) *CardRepository {
	return &CardRepository{DB: DB}
}

func (r CardRepository) Create(ctx context.Context, card model.Card) error {
	err := r.DB.Table("cards").Create(&card).Error
	return err
}

func (r CardRepository) GetByUserID(ctx context.Context, ID int) (model.Card, error) {
	var res model.Card
	err := r.DB.WithContext(ctx).Where("user_id = ?", ID).First(&res).Error
	if err != nil {
		return model.Card{}, err
	}

	return res, nil
}

func (r CardRepository) Update(ctx context.Context, card model.Card) error {
	return r.DB.Where("id = ?", card.ID).Updates(&card).Error
}

func (r CardRepository) Delete(ctx context.Context, ID int) error {
	return r.DB.WithContext(ctx).Delete(&model.Card{}, ID).Error
}
