package user

import (
	"github.com/jwilyandi19/simple-auth-v2/usecase/user"
	"github.com/labstack/echo/v4"
)

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
	return nil
}

func (u *UserHandler) Login(c echo.Context) error {
	return nil
}

func (u *UserHandler) Create(c echo.Context) error {
	return nil
}

func (u *UserHandler) Update(c echo.Context) error {
	return nil
}

func (u *UserHandler) GetAll(c echo.Context) error {
	return nil
}
