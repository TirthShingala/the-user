package middleware

import (
	"net/http"
	"strings"

	helper "github.com/TirthShingala/the-user/helpers"
	"github.com/TirthShingala/the-user/responses"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		println(token)
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Success: false,
				Message: "token not provided!",
			})
			ctx.Abort()
			return
		}

		splitToken := strings.Split(token, " ")
		if len(splitToken) != 2 {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Success: false,
				Message: "Bearer token not in proper format!",
			})
			ctx.Abort()
			return
		}

		token = splitToken[1]

		claims, err := helper.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}
