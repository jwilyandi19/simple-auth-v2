package jwt

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//SetJwtUser Set Only JWT for For User
func (j *jwtUsecase) SetJwtUser(g *echo.Group) {

	secret := j.Config.JWTSecret
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))

	g.Use(j.validateJwtUser)
}

func (j *jwtUsecase) validateJwtUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			mid, ok := claims["jti"].(string)
			if !ok {
				return echo.NewHTTPError(http.StatusForbidden, "something wrong with your token id")
			}

			ctx := context.TODO()
			user, err := j.getUserByID(ctx, mid)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, "forbidden")
			}

			c.Set("user", user)

			return next(c)
		}

		return echo.NewHTTPError(http.StatusForbidden, "invalid token")
	}
}
