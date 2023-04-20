package handlers

import (
	"awesomeProject/config"
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/storage"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserHandler struct {
	storage     *storage.Storage
	userManager *service.Service
	config      *config.Config
	log         *zap.Logger
}

func NewUserHandler(logger *zap.Logger, config *config.Config, storage *storage.Storage, userManager *service.Service) *UserHandler {
	return &UserHandler{log: logger, config: config, storage: storage, userManager: userManager}
}

//		SignUp godoc
//	 	@Summary		Sign-Up
//		@Tags			user
//		@Description	create account
//		@ID				create-user
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.User	true	"user info"
//		@Router			/users/signup [post]
func (h *UserHandler) SignUp(c echo.Context) error {
	var user model.User

	err := readJSON(c.Response(), c.Request(), &user)
	if err != nil {
		h.log.Info("Read Json Error", zap.Error(err))
		return err
	}

	err = user.SetPassword(user.Password)
	if err != nil {
		h.log.Info("Set Password Error", zap.Error(err))
		return err
	}

	err = h.userManager.User.Create(c.Request().Context(), user)
	if err != nil {
		h.log.Info("Create User Error", zap.Error(err))
		return err
	}

	err = writeJSON(c.Response(), http.StatusAccepted, envelope{"user": user}, nil)
	if err != nil {
		return err
	}

	h.log.Info("User created")

	return c.JSON(http.StatusOK, user.ID)
}

//	 	SignIn godoc
//		@Summary		SignIn
//		@Tags			user
//		@Description	authorization
//		@ID				login
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.User	true	"credentials"
//		@Router			/users/signin [post]
func (h *UserHandler) SignIn(c echo.Context) error {
	var input model.User

	err := readJSON(c.Response(), c.Request(), &input)
	if err != nil {
		h.log.Info("Read Json Error", zap.Error(err))
		return err
	}

	user, err := h.userManager.User.Get(c.Request().Context(), input.Email)
	if err != nil {
		h.log.Info("User Get Error", zap.Error(err))
		return err
	}

	ok, err := user.MatchesPassword(input.Password)
	if !ok || err != nil {
		h.log.Info("Password incorrect or something wrong", zap.Error(err))
		return err
	}

	token, err := user.CreateJWT(h.config.JWTKey)
	if err != nil {
		h.log.Info("Create JWT Error", zap.Error(err))
		return err
	}
	//c.Response().Header().Set("Content-Type", "application/json")
	//c.Response().WriteHeader(http.StatusOK)
	//c.String(http.StatusOK, token)
	err = WriteCookie(c, token)
	if err != nil {
		h.log.Info("Write Cookie Error", zap.Error(err))
		return err
	}

	h.log.Info("User authorized")
	return nil
}

//		Logout godoc
//	 	@Summary		Logout
//		@Tags			user
//		@Description	delete session
//		@ID				delete-session
//		@Accept			json
//		@Produce		json
//		@Param			input	body		model.User	true	"user info"
//		@Router			/users/logout [delete]
func (h *UserHandler) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Expires = time.Now()
	c.SetCookie(cookie)

	h.log.Info("User Logout")
	return c.String(http.StatusOK, "logout")
}
