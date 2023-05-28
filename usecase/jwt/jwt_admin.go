package jwt

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jwilyandi19/simple-auth-v2/helper"
	"github.com/labstack/echo/v4"
)

func (j *jwtUsecase) ValidateAdminJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &helper.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(j.Config.JWTSecret), nil
		})

		if err != nil {
			log.Println("Error parsing: ", err.Error())
			return err
		}

		if claims, ok := token.Claims.(*helper.JwtClaims); ok {
			if claims.IsAdmin {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
