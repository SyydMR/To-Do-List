package utils


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	db *gorm.DB
)

func Connect() error {
	dbDriver := "mysql"
	dbName := "itemsDB"
	dbUser := "syydmr"
	dbPassword := "183461"
	dbTcp := "@tcp(127.0.0.1:3306)/"
	gormDb, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+dbTcp+dbName+
		"?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("gorm Db connection ", err)
		return err
	}
	db = gormDb
	return nil
}

func GetDB() *gorm.DB {
	return db
}