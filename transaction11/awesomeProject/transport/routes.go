package transport

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"net/http"
)

func (s *Server) routes() {
	v1 := s.HTTP.Group("/api/v1")
	s.HTTP.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	s.HTTP.GET("/live", func(e echo.Context) error {
		return e.NoContent(http.StatusOK)
	})

	v1.POST("/rent", s.h.HistoryBook.Rent)
	v1.POST("/return", s.h.HistoryBook.Return)
	v1.GET("/debtors", s.h.HistoryBook.Debtors)
	v1.GET("/incomes", s.h.HistoryBook.Incomes)
}
