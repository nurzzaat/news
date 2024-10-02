package news

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

//	@Tags		News
//	@Param		id		path		integer	true	"id"
//	@Success	200		{object}	models.NewsResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/news/{id} [get]
func (nc *NewsController) GetByID(c *gin.Context) {
	newsID, _ := strconv.Atoi(c.Param("id"))

	news, err := nc.NewsRepository.GetByID(c, newsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось получить новости",
		})
		return
	}
	c.JSON(http.StatusOK, news)
}
