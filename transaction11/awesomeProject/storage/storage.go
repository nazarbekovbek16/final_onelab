package storage

import (
	"awesomeProject/config"
	"awesomeProject/storage/postgre"
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Storage struct {
	Pg          *gorm.DB
	User        IUserRepository
	Book        IBookRepository
	Card        ICardRepository
	HistoryBook IHistoryBookRepository
}

func NewStorage(logger *zap.Logger, ctx context.Context, cfg *config.Config) (*Storage, error) {
	var storage Storage

	pgDB, err := postgre.OpenDB(cfg)
	if err != nil {
		logger.Error("Dial error", zap.Error(err))
		return nil, err
	}

	storage.User = postgre.NewUserRepositry(pgDB)
	storage.Book = postgre.NewBookRepository(pgDB)
	storage.Card = postgre.NewCardRepository(pgDB)
	storage.HistoryBook = postgre.NewHistoryBookRepository(pgDB)

	return &storage, nil
}
