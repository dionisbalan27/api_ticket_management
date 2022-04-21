package middleware

import (
	"net/http"

	"api_ticket/usecase"
	"api_ticket/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth(jwtUsecase usecase.JwtUsecase) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader[7:])
		if err != nil {
			errorRes := utils.ResponseError("Unauthorized", err, 401)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorRes)
			return
		}

		c.Set("user_id", userId)
	}
}
