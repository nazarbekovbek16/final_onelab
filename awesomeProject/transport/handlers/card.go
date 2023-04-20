package handlers

import (
	"awesomeProject/config"
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/storage"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type CardHandler struct {
	storage     *storage.Storage
	cardManager *service.Service
	config      *config.Config
	log         *zap.Logger
}

func NewCardHandler(logger *zap.Logger, config *config.Config, storage *storage.Storage, cardManager *service.Service) *CardHandler {
	return &CardHandler{log: logger, config: config, storage: storage, cardManager: cardManager}
}

//		Create godoc
//	 	@Summary		Create
//		@Tags			Card
//		@Description	create card
//		@ID				Card
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.Card	true	"card info"
//		@Router			/cards [post]
func (h *CardHandler) Create(c echo.Context) error {
	var card model.Card

	err := readJSON(c.Response(), c.Request(), &card)
	if err != nil {
		return err
	}

	err = h.cardManager.Card.Create(c.Request().Context(), card)
	if err != nil {
		return err
	}

	err = writeJSON(c.Response(), http.StatusAccepted, envelope{"card": card}, nil)
	if err != nil {
		return err
	}

	h.log.Info("card was created")
	return nil
}
