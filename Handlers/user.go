package handlers

import (
	db "adminpanel/DB"
	"adminpanel/helper"
	"adminpanel/middleware"
	"adminpanel/model"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Expires", "0")
	ok := middleware.ValidateCookies(c)
	if !ok {
		c.HTML(http.StatusOK, "signup.html", nil)
		return
	}
	c.Redirect(http.StatusNotFound, "/")
}
func SignupPost(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Expires", "0")
	var error model.Invalid
	userName := c.Request.FormValue("Name")
	userEmail := c.Request.FormValue("Email")
	password := c.Request.FormValue("Password")
	confirmPassword := c.Request.FormValue("ConfirmPassword")
	if userName == "" {
		error.NameError = "Name should not be empty"
		c.HTML(http.StatusBadRequest, "signup.html", error)
		return
	}
	if userEmail == "" {
		error.EmailError = "Email should not be empty"
		c.HTML(http.StatusBadRequest, "signup.html", error)
		return
	}
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(userEmail) {
		error.EmailError = "Email not in the correct format"
		c.HTML(http.StatusBadRequest, "signup.html", error)
		return
	}
	if password == "" {
		error.PasswordError = "Password should not empty"
		c.HTML(http.StatusBadRequest, "signup.html", error)
		return
	}
	if password != confirmPassword {
		error.PasswordError = "Password doesnot match"
		c.HTML(http.StatusBadRequest, "signup.html", error)
		return
	}
	var count int
	if err := db.Db.Raw("SELECT COUNT(*) FROM users WHERE email=$1", userEmail).Scan(&count).Error; err != nil {
		fmt.Println(err)
		c.HTML(http.StatusOK, "signup.html", nil)
		return
	}
	if count > 0 {
		error.EmailError = "User already exits"
		c.HTML(http.StatusBadRequest, "signup.html", error)
		return
	}
	err := db.Db.Exec("INSERT INTO users (user_name,email,password) VALUES($1,$2,$3)", userName, userEmail, password).Error
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect(http.StatusFound, "/")
}
func LoginHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Expires", "0")
	ok := middleware.ValidateCookies(c)
	role, _, _ := middleware.FindRole(c)
	if !ok {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		if role == "user" {
			c.Redirect(http.StatusFound, "/home")
			return
		} else if role == "admin" {
			c.Redirect(http.StatusFound, "/admin")
			return
		}
		c.HTML(http.StatusBadRequest, "login.html", nil)
	}
}
func LoginPost(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Expires", "0")
	var error model.Invalid
	Newmail := c.Request.FormValue("Email")
	Newpassword := c.Request.FormValue("Password")
	var compare model.Compare
	if err := db.Db.Raw("SELECT password,role,user_name FROM users WHERE email=$1", Newmail).Scan(&compare).Error; err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "login.html", nil)
		return
	}
	if compare.Password != Newpassword {
		error.PasswordError = "Check password again"
		c.HTML(http.StatusBadRequest, "login.html", error)
		return
	}
	if compare.Role == "user" {
		user := model.User{
			Role:     compare.Role,
			UserName: compare.UserName,
		}
		helper.CreateToken(user, c)
		c.Redirect(http.StatusFound, "/home")
		return
	} else if compare.Role == "admin" {
		user := model.User{
			Role:     compare.Role,
			UserName: compare.UserName,
		}
		helper.CreateToken(user, c)
		c.Redirect(http.StatusFound, "/admin")
		return
	} else {
		error.EmailError = "Role mismatch"
		c.HTML(http.StatusOK, "login.html", error)
		return
	}
}
func HomeHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Expires", "0")
	ok := middleware.ValidateCookies(c)
	role, User, _ := middleware.FindRole(c)
	if !ok {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		if role == "user" {
			c.HTML(http.StatusOK, "home.html", gin.H{"Username": User})
			return
		} else {
			c.Redirect(http.StatusFound, "/")
			return
		}
	}

}
func LogoutHandler(c *gin.Context) {
	middleware.DeleteCookie(c)
	c.Redirect(http.StatusFound, "/")
}
