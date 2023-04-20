package postgre

import (
	"awesomeProject/model"
	"context"
	"gorm.io/gorm"
)

type HistoryBookRepository struct {
	DB *gorm.DB
}

func NewHistoryBookRepository(DB *gorm.DB) *HistoryBookRepository {
	return &HistoryBookRepository{DB: DB}
}

func (r HistoryBookRepository) Create(ctx context.Context, book model.HistoryBook) error {
	err := r.DB.Table("historybook").Create(&book).Error
	return err
}

func (r HistoryBookRepository) List(ctx context.Context) ([]model.HistoryBook, error) {
	//TODO implement me
	panic("implement me")
}

func (r HistoryBookRepository) Update(ctx context.Context, book model.Card) error {
	//TODO implement me
	panic("implement me")
}

func (r HistoryBookRepository) Delete(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}
