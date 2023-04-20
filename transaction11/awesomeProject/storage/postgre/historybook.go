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
	err := r.DB.Table("history_books").Create(&book).Error
	return err
}

func (r HistoryBookRepository) List(ctx context.Context) ([]model.HistoryBook, error) {
	//TODO implement me
	panic("implement me")
}
func (r HistoryBookRepository) ListDebtors(ctx context.Context) ([]model.HistoryBook, error) {
	var res []model.HistoryBook
	err := r.DB.Where("is_given = ?", false).Find(&res).Error
	if err != nil {
		return []model.HistoryBook{}, err
	}
	return res, nil
}
func (r HistoryBookRepository) ListIncomes(ctx context.Context) ([]model.Income, error) {
	var res []model.Income
	err := r.DB.Table("history_books").Select("book_id", "SUM(price) as total").Group("book_id").Find(&res).Error
	//err := r.DB.Where("is_given = ?", false).Find(&res).Error
	if err != nil {
		return []model.Income{}, err
	}
	return res, nil
}
func (r HistoryBookRepository) GetByID(ctx context.Context, ID int) (model.HistoryBook, error) {
	var res model.HistoryBook
	err := r.DB.WithContext(ctx).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return model.HistoryBook{}, err
	}
	return res, nil
}

func (r HistoryBookRepository) Update(ctx context.Context, book model.HistoryBook) error {
	err := r.DB.Table("history_books").Updates(book).Error
	return err
}

func (r HistoryBookRepository) Delete(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}
