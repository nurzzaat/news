package categories

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

// @Tags		News
// @Param		id		path		integer	true	"id"
// @Success	200		{object}	models.CategoryResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/categories/{id} [get]
func (nc *CategoryController) GetByID(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Param("id"))

	category, err := nc.CategoryRepository.GetByID(c, categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось получить категорию",
		})
		return
	}
	c.JSON(http.StatusOK, category)
}
