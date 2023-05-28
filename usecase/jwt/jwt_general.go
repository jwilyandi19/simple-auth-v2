package jwt

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (j *jwtUsecase) SetJwtGeneral(g *echo.Group) {
	secret := j.Config.JWTSecret

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))

	g.Use(j.ValidateGeneralJwt)
}

func (j *jwtUsecase) ValidateGeneralJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["is_admin"] == true {
				return next(c)
			} else {
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
			}
			return next(c)
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
