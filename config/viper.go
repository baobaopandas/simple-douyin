package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var CONFIG System

func Viper() {
	//配置文件格式
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yml")
	//读取配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("无法读取配置文件: %w", err)
	}

	err = viper.Unmarshal(&CONFIG)
	if err != nil {
		log.Panic("配置文件读取错误")
	}
	log.Println(CONFIG)
	// 监视配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件被修改：", e.Name)
	})
}
