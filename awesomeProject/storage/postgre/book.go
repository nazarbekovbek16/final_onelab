package postgre

import (
	"awesomeProject/model"
	"context"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) *BookRepository {
	return &BookRepository{DB: DB}
}

func (r *BookRepository) Create(ctx context.Context, book model.Book) error {
	err := r.DB.Table("books").Create(&book).Error
	return err
}

func (r *BookRepository) List(ctx context.Context) ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *BookRepository) GetByID(ctx context.Context, ID int) (model.Book, error) {
	var res model.Book
	err := r.DB.WithContext(ctx).Where("id = ?", ID).First(&res).Error
	if err != nil {
		return model.Book{}, err
	}
	return res, nil
}

func (r *BookRepository) Update(ctx context.Context, book model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r *BookRepository) Delete(ctx context.Context, ID int) error {
	return r.DB.WithContext(ctx).Delete(&model.Book{}, ID).Error
}
