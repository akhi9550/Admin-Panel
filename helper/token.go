package helper

import (
	"adminpanel/model"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(user model.User, c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": user.Role,
		"user": user.UserName,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("KEY")))
	if err == nil {
		fmt.Println("Token created")
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("cookie", tokenString, 3600, "", "", false, true)
}
