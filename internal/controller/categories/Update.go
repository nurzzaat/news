package categories

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

// @Tags		News
// @Param		id		path	int				true	"id"
// @Param		name	body	models.CategoryRequest	true	"category"
// @Security	ApiKeyAuth
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/category/{id} [put]
func (nc *CategoryController) Update(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Param("id"))

	categoryRequest := models.CategoryRequest{}
	if err := c.ShouldBind(&categoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Неправильный формат данных",
		})
		return
	}
	categoryRequest.ID = categoryID
	category, _ := nc.CategoryRepository.GetByID(c, categoryID)
	if categoryRequest.Name == "" {
		categoryRequest.Name = category.Name
	}

	err := nc.CategoryRepository.Edit(c, categoryRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось обновить категорию",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: true})
}
