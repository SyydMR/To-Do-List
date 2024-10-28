package models

import (
	"github.com/SyydMR/To-Do-List/utils"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	utils.Connect()
	db = utils.GetDB()
    db.AutoMigrate(&User{}, &Item{})
}
