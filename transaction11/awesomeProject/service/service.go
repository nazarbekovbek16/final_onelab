package service

import (
	"awesomeProject/storage"
	"fmt"
	"go.uber.org/zap"
)

type Service struct {
	storage            *storage.Storage
	HistoryBookService IHistoryBookService
}

func NewService(logger *zap.Logger, storage *storage.Storage) (*Service, error) {
	if storage == nil {
		logger.Error("Storage pointer is empty")
		return nil, fmt.Errorf("storage is empty")
	}

	var service Service

	service.HistoryBookService = NewHistoryBookService(storage)

	return &service, nil
}
