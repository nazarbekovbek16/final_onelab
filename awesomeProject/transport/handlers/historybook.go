package handlers

import (
	"awesomeProject/config"
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/storage"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
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
//		@Description	make request to another service to create record about renting of book
//		@ID				Rent
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.HistoryBook	true	"history book info"
//		@Router			/rent/:id [post]
func (h *HistoryBookHandler) Rent(c echo.Context) error {
	var rent model.HistoryBook
	id, err := ReadIDParam(c)
	if err != nil {
		h.log.Info("read param error")
		return err
	}

	err = readJSON(c.Response(), c.Request(), &rent)
	if err != nil {
		h.log.Info("read json error")
		return err
	}

	//cookie, err := ReadCookie(c)
	//if err != nil {
	//	h.log.Info("read cookie error")
	//	return err
	//}
	//claims, err := ExtractClaims(h.config.JWTKey, cookie)
	//if err != nil {
	//	h.log.Info("ExtractClaims error")
	//	return err
	//}
	//rent.UserID, err = strconv.Atoi(fmt.Sprint(claims["id"]))
	rent.BookID = id
	if err != nil {
		h.log.Info("user id atoi error")
		return err
	}

	jsonData, err := json.Marshal(rent)
	if err != nil {
		h.log.Info("error marshaling JSON")
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/api/v1/rent", bytes.NewReader(jsonData))
	if err != nil {
		h.log.Info("error creating HTTP request")
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var resID int
	err = json.Unmarshal(body, &resID)
	defer resp.Body.Close()

	h.log.Info("record about renting of book created, id of book: " + strconv.Itoa(resID))

	return nil
}

//		Return godoc
//	 	@Summary		Return
//		@Tags			History of Book
//		@Description	make request to another service to update record when user return book
//		@ID				Return
//		@Accept			json
//		@Produce		json
//		@Param			input	body		integer	true	"history book info"
//		@Router			/rent/return/:id [post]
func (h *HistoryBookHandler) Return(c echo.Context) error {
	id, err := ReadIDParam(c)
	if err != nil {
		h.log.Info("read param error")
		return err
	}

	jsonData, err := json.Marshal(id)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/api/v1/return", bytes.NewReader(jsonData))
	if err != nil {
		h.log.Info("error creating HTTP request")
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var resID int
	err = json.Unmarshal(body, &resID)
	defer resp.Body.Close()
	h.log.Info("Book was returned, id of book: " + strconv.Itoa(resID))
	return nil
}

//		Debtors godoc
//	 	@Summary		Debtors
//		@Tags			History of Book
//		@Description	make request to another service to get records of users with book
//		@ID				Debtors
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.HistoryBook	true	"history book info"
//		@Router			/rent/debtors [get]
func (h *HistoryBookHandler) Debtors(c echo.Context) error {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/api/v1/debtors", nil)
	if err != nil {
		h.log.Info("error creating HTTP request")
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var result []model.HistoryBook
	err = json.Unmarshal(body, &result)
	if err != nil {
		h.log.Info("error unmarshal body")
		return err
	}

	err = writeJSON(c.Response(), http.StatusAccepted, envelope{"debtors": result}, nil)
	if err != nil {
		h.log.Info("error write json")
		return err
	}

	h.log.Info("List of history books successfully came")
	return nil
}

//		Incomes godoc
//	 	@Summary		Incomes
//		@Tags			History of Book
//		@Description	make request to another service to get list of incomes of each books
//		@ID				Income
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.Income	true	"history book info"
//		@Router			/rent/incomes [get]
func (h *HistoryBookHandler) Incomes(c echo.Context) error {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/api/v1/incomes", nil)
	if err != nil {
		h.log.Info("error creating HTTP request")
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var result []model.Income
	err = json.Unmarshal(body, &result)
	if err != nil {
		h.log.Info("error unmarshal body")
		return err
	}
	err = writeJSON(c.Response(), http.StatusAccepted, envelope{"incomes": result}, nil)
	if err != nil {
		h.log.Info("error write json")
		return err
	}

	h.log.Info("Information about incomes of books sent")
	return nil
}
