package service

import (
	"awesomeProject/model"
	"awesomeProject/storage"
	"context"
)

type HistoryBookService struct {
	repo *storage.Storage
}

func NewHistoryBookService(repo *storage.Storage) *HistoryBookService {
	return &HistoryBookService{repo: repo}
}

type IHistoryBookService interface {
	Create(ctx context.Context, book model.HistoryBook) error
	List(ctx context.Context) ([]model.HistoryBook, error)
	Update(ctx context.Context, book model.Card) error
	Delete(ctx context.Context, ID string) error
}

func (s HistoryBookService) Create(ctx context.Context, book model.HistoryBook) error {
	// 1 day costs 50 tenge
	sum := float64(book.Duration * 50)

	card, err := s.repo.Card.GetByUserID(ctx, book.UserID)
	if err != nil {
		return err
	}

	if card.Money-sum >= 0 {
		book.IsPaid = true
		card.Money = card.Money - sum
		s.repo.Card.Update(ctx, card)
	}

	return s.repo.HistoryBook.Create(ctx, book)
}

func (s HistoryBookService) List(ctx context.Context) ([]model.HistoryBook, error) {
	//TODO implement me
	panic("implement me")
}

func (s HistoryBookService) Update(ctx context.Context, book model.Card) error {
	//TODO implement me
	panic("implement me")
}

func (s HistoryBookService) Delete(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}
