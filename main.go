package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"web.com/ginGormJwt/common"
)

func main() {

	//db := InitDB()

	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) //默认8080
}

func InitConfig() {
	//获取当前的工作目录
	workDir, _ := os.Getwd()
	//设置要读取的文件名
	viper.SetConfigName("application")
	//设置读取文件的类型
	viper.SetConfigType("yml")
	//设置文件的路径
	viper.AddConfigPath(workDir + "/config")

	err := viper.ReadInConfig()
	if err != nil {

	}

}
