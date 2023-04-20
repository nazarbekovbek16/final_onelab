package storage

import (
	"awesomeProject/model"
	"context"
)

type IUserRepository interface {
	GetUser(ctx context.Context, ID int) (model.User, error)
	GetByEmail(ctx context.Context, username string) (model.User, error)
	Auth(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, ID int) error
	CreateUser(ctx context.Context, item model.User) (int, error)
}

type IBookRepository interface {
	Create(ctx context.Context, book model.Book) error
	List(ctx context.Context) ([]model.Book, error)
	GetByID(ctx context.Context, ID int) (model.Book, error)
	Update(ctx context.Context, book model.Book) error
	Delete(ctx context.Context, ID int) error
}

type ICardRepository interface {
	Create(ctx context.Context, card model.Card) error
	Update(ctx context.Context, card model.Card) error
	GetByUserID(ctx context.Context, ID int) (model.Card, error)
	Delete(ctx context.Context, ID int) error
}

type IHistoryBookRepository interface {
	Create(ctx context.Context, book model.HistoryBook) error
	List(ctx context.Context) ([]model.HistoryBook, error)
	Update(ctx context.Context, book model.Card) error
	Delete(ctx context.Context, ID string) error
}
