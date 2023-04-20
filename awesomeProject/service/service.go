package service

import (
	"awesomeProject/storage"
	"fmt"
	"go.uber.org/zap"
)

type Service struct {
	storage            *storage.Storage
	User               IUserService
	Book               IBookService
	Card               ICardService
	HistoryBookService IHistoryBookService
}

func NewService(logger *zap.Logger, storage *storage.Storage) (*Service, error) {
	if storage == nil {
		logger.Error("Storage pointer is empty")
		return nil, fmt.Errorf("storage is empty")
	}

	var service Service

	service.User = NewUserService(storage)
	service.Book = NewBookService(storage)
	service.Card = NewCardService(storage)
	service.HistoryBookService = NewHistoryBookService(storage)

	return &service, nil
}
