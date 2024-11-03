package auth

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/news/internal/controller/tokenutil"
	"github.com/nurzzaat/news/internal/models"
	"github.com/nurzzaat/news/pkg"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	UserRepository models.UserRepository
	Env            *pkg.Env
}

var (
	verifier = emailverifier.NewVerifier()
)

// @Summary	SignUp
// @Tags		auth
// @Accept		json
// @Produce	json
// @Param		user	body		models.UserRequest	true	"user"
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/register [post]
func (uc *AuthController) Signup(c *gin.Context) {
	var request models.UserRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Неправильный формат данных",
		})
		return
	}

	verifier = verifier.EnableSMTPCheck()
	verifier = verifier.EnableDomainSuggest()

	if request.Email == "" || request.Password == "" || request.Name == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Пустые значения",
		})
		return
	}

	ret, _ := verifier.Verify(request.Email)

	if !ret.Syntax.Valid {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Неверный адрес электронной почты",
		})
		return
	}

	user, _ := uc.UserRepository.GetUserByEmail(c, request.Email)
	if user.ID > 0 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Пользователь существует",
		})
		return
	}
	err := validatePassword(request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Пароль должен содержить не менее 6 символов",
		})
		return
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Пароли не совпадают",
		})
		return
	}
	request.Password = string(encryptedPassword)

	_, err = uc.UserRepository.CreateUser(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось создать пользователя",
		})
		return
	}
	user, err = uc.UserRepository.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Не удалось получить пользователя",
		})
		return
	}
	accessToken, err := tokenutil.CreateAccessToken(&user, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: "Ошибка при создании токена",
		})
		return
	}
	c.JSON(http.StatusOK, models.TokenResponse{Token: accessToken})
}

func isICloudEmail(email string) bool {
	icloudPattern := `@icloud\.com$`
	icloudRegex := regexp.MustCompile(icloudPattern)

	return icloudRegex.MatchString(email)
}

func GenerateRandomPassword(size int) string {
	var alpha = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	password := make([]rune, size)
	for i := 0; i < size; i++ {
		password[i] = alpha[rand.Intn(len(alpha)-1)]
	}
	hashPassword := string(password)
	return hashPassword
}

func validatePassword(password string) error {
	if len(password) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}
	return nil
}
