package news

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/models"
)

type NewsController struct {
	NewsRepository models.NewsRepository
}

//	@Tags		News
//	@Param		title		formData	string	true	"Title"
//	@Param		content		formData	string	true	"Content"
//	@Param		categoryId	formData	int		true	"CategoryId"
//	@Param		thumbnail	formData	file	true	"thumbnail"
//	@Security	ApiKeyAuth
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/news [post]
func (nc *NewsController) Create(c *gin.Context) {
	roleID := c.GetUint("roleID")
	userID := c.GetUint("userID")

	if roleID != models.ADMIN {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "У вас нет прав для выполнения этой операции",
		})
		return
	}

	newsRequest := models.NewsRequest{}
	newsRequest.Title = c.PostForm("title")
	newsRequest.Content = c.PostForm("content")
	newsRequest.CategoryID, _ = strconv.Atoi(c.PostForm("categoryId"))

	file, err := c.FormFile("thumbnail")
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
	newsRequest.ThumbNail = referenceURL[1:]

	id, err := nc.NewsRepository.Create(c, newsRequest, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось создать новости" ,
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: id})
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
