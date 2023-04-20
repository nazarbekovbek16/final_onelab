package handlers

import (
	"awesomeProject/config"
	"awesomeProject/service"
	"awesomeProject/storage"
	"awesomeProject/transport/middleware"
	"go.uber.org/zap"
)

type Handlers struct {
	HistoryBook *HistoryBookHandler
	mid         *middleware.JWTAuth
	log         *zap.Logger
}

func NewHandlers(l *zap.Logger, config *config.Config, storage *storage.Storage, service *service.Service, auth *middleware.JWTAuth) *Handlers {
	return &Handlers{
		HistoryBook: NewHistoryBookHandler(l, storage, service, config),
		mid:         auth,
		log:         l,
	}
}
