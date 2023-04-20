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

type BookHandler struct {
	storage     *storage.Storage
	bookManager *service.Service
	config      *config.Config
	log         *zap.Logger
}

func NewBookHandler(logger *zap.Logger, config *config.Config, storage *storage.Storage, bookManager *service.Service) *BookHandler {
	return &BookHandler{log: logger, config: config, storage: storage, bookManager: bookManager}
}

//		Create godoc
//	 	@Summary		Create
//		@Tags			Book
//		@Description	create book
//		@ID				BookCreate
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.Book	true	"book info"
//		@Router			/books [post]
func (h *BookHandler) Create(c echo.Context) error {

	var book model.Book

	err := readJSON(c.Response(), c.Request(), &book)
	if err != nil {
		return err
	}

	_, err = h.bookManager.Book.Create(c.Request().Context(), book)
	if err != nil {
		return err
	}

	err = writeJSON(c.Response(), http.StatusAccepted, envelope{"book": book}, nil)
	if err != nil {
		return err
	}
	h.log.Info("book was created")
	return nil
}

//		Get godoc
//	 	@Summary		Get
//		@Tags			BookGet
//		@Description	get book
//		@ID				Book
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.Book	true	"book info"
//		@Router			/books/:is [get]
func (h *BookHandler) Get(c echo.Context) error {
	id, err := ReadIDParam(c)
	if err != nil {
		return err
	}

	book, err := h.bookManager.Book.GetByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	err = writeJSON(c.Response(), http.StatusAccepted, envelope{"book": book}, nil)
	if err != nil {
		return err
	}
	return nil
}
func (h *BookHandler) Update(c echo.Context) error {
	return nil
}

//		Delete godoc
//	 	@Summary		Delete
//		@Tags			Book
//		@Description	Delete book
//		@ID				BookDelete
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.Book	true	"book info"
//		@Router			/books/:id [delete]
func (h *BookHandler) Delete(c echo.Context) error {
	id, err := ReadIDParam(c)
	if err != nil {
		return err
	}

	err = h.bookManager.Book.Delete(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return nil
}
