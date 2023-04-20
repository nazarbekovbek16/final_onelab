package handlers

import (
	"awesomeProject/config"
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/storage"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
)

type HistoryBookHandler struct {
	storage            *storage.Storage
	historyBookManager *service.Service
	config             *config.Config
	log                *zap.Logger
}

func NewHistoryBookHandler(logger *zap.Logger, storage *storage.Storage, historyBookManager *service.Service, config *config.Config) *HistoryBookHandler {
	return &HistoryBookHandler{log: logger, storage: storage, historyBookManager: historyBookManager, config: config}
}

//		Rent godoc
//	 	@Summary		Rent
//		@Tags			History of Book
//		@Description	create record about renting of book
//		@ID				Rent
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.HistoryBook	true	"history book info"
//		@Router			/rent [post]
func (h *HistoryBookHandler) Rent(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		h.log.Info("read body error")
		return err
	}

	var data model.HistoryBook
	err = json.Unmarshal(body, &data)
	if err != nil {
		h.log.Info("unmarshal error")
		return err
	}
	// нужно изучить как использовать анмаршал в echo 
	err = h.historyBookManager.HistoryBookService.Create(c.Request().Context(), data)
	if err != nil {
		h.log.Info("history of book was failed")
		return err
	}
	h.log.Info("history of book was created")
	return c.JSON(http.StatusOK, data.ID)
}

//		Return godoc
//	 	@Summary		Return
//		@Tags			History of Book
//		@Description	update record when user return book
//		@ID				Return
//		@Accept			json
//		@Produce		json
//		@Param			input	body		integer	true	"history book info"
//		@Router			/return [post]
func (h *HistoryBookHandler) Return(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		h.log.Info("read body error")
		return err
	}

	var data int
	err = json.Unmarshal(body, &data)
	if err != nil {
		h.log.Info("unmarshal error")
		return err
	}

	res, err := h.historyBookManager.HistoryBookService.GetByID(c.Request().Context(), data)
	if err != nil {
		h.log.Info("getting of history book was failed")
		return err
	}

	res.IsGiven = true

	err = h.historyBookManager.HistoryBookService.Update(c.Request().Context(), res)
	if err != nil {
		h.log.Info("updating of history book was failed")
		return err
	}
	h.log.Info("book was returned, id:" + strconv.Itoa(data))

	return c.JSON(http.StatusOK, data)
}

//		Debtors godoc
//	 	@Summary		Debtors
//		@Tags			History of Book
//		@Description	get records of users with book
//		@ID				Debtors
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.HistoryBook	true	"history book info"
//		@Router			/debtors [get]
func (h *HistoryBookHandler) Debtors(c echo.Context) error {
	debtors, err := h.historyBookManager.HistoryBookService.ListDebtors(c.Request().Context())
	if err != nil {
		h.log.Info("Service ListDebtors Error")
		return err
	}
	fmt.Println(debtors)
	h.log.Info("Debtors list sent")
	return c.JSON(http.StatusOK, debtors)
}

//		Incomes godoc
//	 	@Summary		Incomes
//		@Tags			History of Book
//		@Description	get list of incomes of each books
//		@ID				Income
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.Income	true	"incomes"
//		@Router			/incomes [get]
func (h *HistoryBookHandler) Incomes(c echo.Context) error {
	incomes, err := h.historyBookManager.HistoryBookService.ListIncomes(c.Request().Context())
	if err != nil {
		h.log.Info("Service List incomes Error")
		return err
	}
	fmt.Println(incomes)
	h.log.Info("Incomes list sent")
	return c.JSON(http.StatusOK, incomes)
}
