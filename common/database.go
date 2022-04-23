package common

import (
	"fmt"

	"web.com/ginGormJwt/model"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

//连接数据库
func InitDB() *gorm.DB {

	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	/*driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "test"
	username := "root"
	password := "123456"
	charset := "utf8"*/
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	//根据结构体快速创建表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

//获取DB
func GetDB() *gorm.DB {
	return DB
}
