package middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateCookies(c *gin.Context) bool {
	cookie, _ := c.Cookie("cookie")
	if cookie == "" {
		fmt.Println("Cookie not found")
		return false
	} else {
		fmt.Println("Cookie ", cookie)
		return true
	}
}
func FindRole(c *gin.Context) (string, string, error) {
	cookie, _ := c.Cookie("cookie")
	if cookie == "" {
		return "", "", fmt.Errorf("cookie not found")
	}
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:%v", token.Header["alg"])
		}
		return []byte(os.Getenv("key")), nil
	})
	if err != nil {
		return "", "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", fmt.Errorf("invalid token")
	}
	role := claims["role"].(string)
	user := claims["user"].(string)
	return role, user, nil
}
func DeleteCookie(c *gin.Context) {
	c.SetCookie("cookie", "", 0, "", "", false, true)
	fmt.Println("cookie deleted")
}
