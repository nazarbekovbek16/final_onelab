package service

import (
	"awesomeProject/model"
	"awesomeProject/storage"
	"context"
)

type BookService struct {
	repo *storage.Storage
}

func NewBookService(repo *storage.Storage) *BookService {
	return &BookService{repo: repo}
}

type IBookService interface {
	Create(ctx context.Context, book model.Book) (uint, error)
	List(ctx context.Context) ([]model.Book, error)
	GetByID(ctx context.Context, ID int) (model.Book, error)
	Update(ctx context.Context, book model.Book) error
	Delete(ctx context.Context, ID int) error
}

func (s BookService) Create(ctx context.Context, book model.Book) (uint, error) {
	err := s.repo.Book.Create(ctx, book)
	if err != nil {
		return 0, err
	}
	return book.ID, nil
}

func (s BookService) List(ctx context.Context) ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (s BookService) GetByID(ctx context.Context, ID int) (model.Book, error) {
	var book model.Book
	book, err := s.repo.Book.GetByID(ctx, ID)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s BookService) Update(ctx context.Context, book model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (s BookService) Delete(ctx context.Context, ID int) error {
	return s.repo.Book.Delete(ctx, ID)
}
