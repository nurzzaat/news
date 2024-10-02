package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

type CategoryController struct {
	CategoryRepository models.CategoryRepository
}

// @Tags		News
// @Param		name	body	models.CategoryRequest	true	"category"
// @Security	ApiKeyAuth
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/category [post]
func (nc *CategoryController) Create(c *gin.Context) {
	roleID := c.GetUint("roleID")

	if roleID != models.ADMIN {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "У вас нет прав для выполнения этой операции",
		})
		return
	}

	category := models.CategoryRequest{}
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Неправильный формат данных",
		})
		return
	}
	id, err := nc.CategoryRepository.Create(c, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось создать категорию",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: id})

}