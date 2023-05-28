package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (j *jwtUsecase) SetJwtAdmin(g *echo.Group) {

	secret := j.Config.JWTSecret

	// validate jwt token
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))

	g.Use(j.validateJwtAdmin)
}

func (j *jwtUsecase) validateJwtAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user")
		token := user.(*jwt.Token)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["is_admin"] == true {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
