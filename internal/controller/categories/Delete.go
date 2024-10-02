package categories

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

// @Tags		News
// @Param		id	path	integer	true	"id"
// @Security	ApiKeyAuth
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/category/{id} [delete]
func (nc *CategoryController) Delete(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Param("id"))
	roleID := c.GetUint("roleID")

	if roleID != models.ADMIN {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "У вас нет прав для выполнения этой операции",
		})
		return
	}

	err := nc.CategoryRepository.Delete(c, categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось удалить категорию",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: true})
}
