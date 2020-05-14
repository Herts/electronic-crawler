package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

func InitDB() {
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", viper.GetString("db.username"),
		viper.GetString("db.password"), viper.GetString("db.url"),
		viper.GetString("db.port"), viper.GetString("db.database"))
	var err error
	db, err = gorm.Open("mysql", dbLink)
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(viper.GetBool("db.debug"))
	db.AutoMigrate(&Page{})
	db.AutoMigrate(&Task{})
}

var db *gorm.DB
