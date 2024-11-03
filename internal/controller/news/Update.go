package news

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

// @Tags		News
// @Param		id		path		int		true	"id"
// @Param		title		formData	string	false	"Title"
// @Param		content		formData	string	false	"Content"
// @Param		categoryId	formData	int		false	"CategoryId"
// @Param		thumbnail		formData	file	false	"thumbnail"
// @Security	ApiKeyAuth
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/news/{id} [put]
func (nc *NewsController) Update(c *gin.Context) {
	newsRequest := models.NewsRequest{}
	newsRequest.ID, _ = strconv.Atoi(c.Param("id"))
	newsRequest.Title = c.PostForm("title")
	newsRequest.Content = c.PostForm("content")
	newsRequest.CategoryID, _ = strconv.Atoi(c.PostForm("categoryId"))

	var tempFilePath string
	file, err := c.FormFile("thumbnail")
	if err != nil {
		tempFilePath = ""
	} else {
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Result: "Не удалось открыть файл",
			})
			return
		}

		filename := GetMD5Hash(file.Filename)
		referenceURL := "./news_images/" + filename

		if err := c.SaveUploadedFile(file, referenceURL); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Result: "Не удалось сохранить файл",
			})
			return
		}
		tempFilePath = referenceURL[1:]
	}
	newsRequest.ThumbNail = tempFilePath

	news, _ := nc.NewsRepository.GetByID(c, newsRequest.ID)
	if newsRequest.Title == "" {
		newsRequest.Title = news.Title
	}
	if newsRequest.Content == "" {
		newsRequest.Content = news.Content
	}
	if newsRequest.CategoryID == 0 {
		newsRequest.CategoryID = news.Category.ID
	}
	if newsRequest.ThumbNail == "" {
		newsRequest.ThumbNail = news.ThumbNail
	}

	err = nc.NewsRepository.Edit(c, newsRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось обновить новости",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: true})
}
