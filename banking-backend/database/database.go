package database

import (
	"duomly.com/go-bank-backend/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDatabase() {

	database, err := gorm.Open("mysql", "sql12543249:cIc4vD5Bmg@tcp(sql12.freesqldatabase.com:3306)/sql12543249?charset=utf8&parseTime=True")
	helpers.HandleErr(err)
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
