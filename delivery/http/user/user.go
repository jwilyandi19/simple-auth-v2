package user

import (
	"context"
	"encoding/json"
	"net/http"

	domain "github.com/jwilyandi19/simple-auth-v2/domain/user"
	"github.com/jwilyandi19/simple-auth-v2/usecase/user"
	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserHandler struct {
	UserUsecase user.UserUsecase
}

func NewUserHandler(e *echo.Echo, uu user.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: uu,
	}

	e.GET("/user/:id", handler.GetByID)
	e.POST("/user/auth", handler.Login)
	e.POST("/user", handler.Create)
	e.PUT("/user", handler.Update)
	e.GET("/user", handler.GetAll)
}

//Route
//GET /user/:id (GetByID) (User logged in only)
//POST /user/auth (Login) (No need to login)
//POST /user (Create) (No need to login)
//PUT /user (Update) (Admin + user logged in only)
//GET /user (GetAll) (Admin only)

func (u *UserHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.TODO()
	}

	res, err := u.UserUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (u *UserHandler) Login(c echo.Context) error {
	return nil
}

func (u *UserHandler) Create(c echo.Context) error {
	var req CreateUserRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.TODO()
	}

	user := domain.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}

	res, err := u.UserUsecase.Create(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (u *UserHandler) Update(c echo.Context) error {
	var req UpdateUserRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.TODO()
	}

	user := domain.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}

	res, err := u.UserUsecase.Update(ctx, req.ID, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (u *UserHandler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.TODO()
	}

	res, err := u.UserUsecase.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)

}
