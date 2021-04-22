package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigName("config") // 读取yaml配置文件

	//viper.AddConfigPath("/etc/appname/")   //设置配置文件的搜索目录
	//viper.AddConfigPath("$HOME/.appname")  // 设置配置文件的搜索目录

	viper.AddConfigPath(".") // 设置配置文件和可执行二进制文件在用一个目录
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}
}

func main() {
	fmt.Println("获取配置文件的redis", viper.Get(`proxy-groups.Name`))
}
