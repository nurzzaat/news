package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

// @Tags		News
//	@Param		order		query	integer	false	"order by (1-latest,2-older,3-popular)"
// @Success	200		{array}		models.NewsResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/news [get]
func (nc *NewsController) GetAll(c *gin.Context) {
	orderByParameter := c.Query("order")
	news, err := nc.NewsRepository.GetAll(c, orderByParameter)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось получить новости",
		})
		return
	}
	c.JSON(http.StatusOK, news)
}
