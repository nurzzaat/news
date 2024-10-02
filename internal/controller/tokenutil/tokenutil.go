package tokenutil

import (
	"fmt"

	models "github.com/nurzzaat/news/internal/models"

	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	claims := &models.JwtClaims{
		ID:     user.ID,
		RoleID: user.RoleID,
	}
	fmt.Println("claim", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func ValidateJWT(c *gin.Context, secret string) error {
	token, err := getToken(c, secret)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func ValidateUserJWT(c *gin.Context, secret string) error {
	token, err := getToken(c, secret)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRoleID := uint(claims["role"].(float64))
	userID := uint(claims["id"].(float64))
	if ok && token.Valid {
		c.Set("userID", userID)
		c.Set("roleID", userRoleID)
		return nil
	}
	return errors.New("invalidц token provided")
}

func getToken(c *gin.Context, secret string) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	if tokenString == "" {
		return nil, errors.New("invalid token provided")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
