package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jwilyandi19/simple-auth-v2/helper"
)

func AuthorizationMiddleware(ctx *gin.Context) {
	s := ctx.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := helper.ValidateToken(token); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
}
