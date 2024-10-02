package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/controller/tokenutil"
	"github.com/nurzzaat/news/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// @Summary	SignIn
// @Tags		auth
// @Accept		json
// @Produce	json
// @Param		input	body		models.LoginRequest	true	"login"
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/login [post]
func (lc *AuthController) Signin(c *gin.Context) {
	var loginRequest models.LoginRequest

	err := c.ShouldBind(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Неправильный формат данных",
		})
		return
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Пустые значения",
		})
		return
	}

	user, err := lc.UserRepository.GetUserByEmail(c, loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось получить пользователя",
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Пароли не совпадают",
		})
		return
	}
	accessToken, err := tokenutil.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Ошибка при создании токена",
		})
		return
	}
	c.JSON(http.StatusOK, models.TokenResponse{Token: accessToken})
}
