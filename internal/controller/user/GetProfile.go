package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

type UserController struct {
	UserRepository models.UserRepository
}

// @Tags		User
// @Accept		json
// @Produce	json
// @Security	ApiKeyAuth
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/user/profile [get]
func (sc *UserController) GetProfile(c *gin.Context) {
	userID := c.GetUint("userID")

	profile, err := sc.UserRepository.GetProfile(c, int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось получить пользователя",
		})
		return
	}
	user := models.UserResponse{
		ID:    profile.ID,
		Name:  profile.Name,
		Email: profile.Email,
	}
	c.JSON(http.StatusOK, user)
}
