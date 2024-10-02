package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

// @Tags		News
// @Success	200		{array}		models.CategoryResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/categories [get]
func (nc *CategoryController) GetAll(c *gin.Context) {
	categories, err := nc.CategoryRepository.GetAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось получить категории",
		})
		return
	}
	c.JSON(http.StatusOK, categories)
}
