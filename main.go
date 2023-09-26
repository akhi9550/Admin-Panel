package main

import (
	db "adminpanel/DB"
	handler "adminpanel/Handlers"
	"adminpanel/model"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	r := gin.Default()
	db.Db, err = gorm.Open(postgres.Open(os.Getenv("DBS")), &gorm.Config{})
	if err != nil {
		fmt.Println("Database not loaded")
	}
	db.Db.AutoMigrate(&model.User{})
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")
	r.GET("/signup", handler.SignupHandler)
	r.POST("/signup", handler.SignupPost)
	r.GET("/", handler.LoginHandler)
	r.POST("/", handler.LoginPost)
	r.GET("/home", handler.HomeHandler)
	r.GET("/logout", handler.LogoutHandler)

	r.GET("/admin", handler.AdminHome)
	r.POST("/admin", handler.AdminAddUser)
	r.GET("/adminupdate", handler.AdminUpdate)
	r.POST("adminupdatepost", handler.AdminUpdatePost)
	r.GET("/admindelete", handler.AdminDelete)
	r.GET("/logoutadmin", handler.LogoutadminHandler)
	r.Run(":8000")
}
