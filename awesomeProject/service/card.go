package service

import (
	"awesomeProject/model"
	"awesomeProject/storage"
	"context"
)

type CardService struct {
	repo *storage.Storage
}

func NewCardService(repo *storage.Storage) *CardService {
	return &CardService{repo: repo}
}

type ICardService interface {
	Create(ctx context.Context, card model.Card) error
	Update(ctx context.Context, card model.Card) error
	Delete(ctx context.Context, ID int) error
}

func (s CardService) Create(ctx context.Context, card model.Card) error {
	err := s.repo.Card.Create(ctx, card)
	if err != nil {
		return err
	}

	return nil
}

func (s CardService) Update(ctx context.Context, card model.Card) error {
	//TODO implement me
	panic("implement me")
}

func (s CardService) Delete(ctx context.Context, ID int) error {
	return s.repo.Card.Delete(ctx, ID)
}
