package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Role     string `gorm:"not null;default:user"`
	UserName string
	Email    string
	Password string
}
type Invalid struct {
	NameError     string
	EmailError    string
	PasswordError string
	RoleError     string
	CommonError   string
}
type Compare struct {
	Password string
	Role     string
	UserName string
}
type UserDetails struct {
	UserName string
	Email    string
}
