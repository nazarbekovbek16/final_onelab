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

	users := v1.Group("/users")
	users.POST("/signin", s.h.User.SignIn)
	users.POST("/signup", s.h.User.SignUp)
	users.DELETE("/logout", s.h.User.Logout)

	books := v1.Group("/books")
	books.GET("/:id", s.h.Book.Get)
	books.POST("/", s.h.Book.Create)
	books.DELETE("/:id", s.h.Book.Delete)

	cards := v1.Group("/cards")
	cards.POST("/", s.h.Card.Create)

	rent := v1.Group("/rent")
	rent.POST("/:id", s.h.HistoryBook.Rent)
	rent.POST("/return/:id", s.h.HistoryBook.Return)
	rent.GET("/debtors", s.h.HistoryBook.Debtors)
	rent.GET("/incomes", s.h.HistoryBook.Incomes)

}
