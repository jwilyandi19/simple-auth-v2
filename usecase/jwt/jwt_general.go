package jwt

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jwilyandi19/simple-auth-v2/helper"
	"github.com/labstack/echo/v4"
)

func (j *jwtUsecase) ValidateGeneralJwt(next echo.HandlerFunc) echo.HandlerFunc {
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
				id := c.Param("id")
				mid := claims.Username
				ctx := context.TODO()
				_, err := j.getUserByID(ctx, mid)
				if err != nil {
					return echo.NewHTTPError(http.StatusForbidden, "forbidden")
				}
				if id != mid {
					return echo.NewHTTPError(http.StatusForbidden, "forbidden")
				}
			}
			return next(c)
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
