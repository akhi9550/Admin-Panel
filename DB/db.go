package db

import (
	"adminpanel/model"

	"gorm.io/gorm"
)

var Db *gorm.DB
var UserList []model.User
