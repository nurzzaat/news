package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/controller/tokenutil"
	models "github.com/nurzzaat/news/internal/models"
)

func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokenutil.ValidateJWT(c, secret)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Result: "Ошибка авторизации",
			})
			c.Abort()
			return
		}
		err = tokenutil.ValidateUserJWT(c, secret)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Result: "Требуется пользователь",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
