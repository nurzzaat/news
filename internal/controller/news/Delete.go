package news

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

// @Tags		News
// @Param		id	path	integer	true	"id"
// @Security	ApiKeyAuth
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/news/{id} [delete]
func (nc *NewsController) Delete(c *gin.Context) {
	newsID, _ := strconv.Atoi(c.Param("id"))

	news, _ := nc.NewsRepository.GetByID(c, newsID)
	_ = os.Remove(news.ThumbNail)

	err := nc.NewsRepository.Delete(c, newsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось удалить новости",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: true})
}
