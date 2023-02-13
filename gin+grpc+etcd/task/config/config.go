package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	path, _ := os.Getwd()                    // 获取的是cmd目录
	workDir := path[:len(path)-4]            // 改为上一级目录
	viper.SetConfigName("config")            // 配置文件的名称
	viper.SetConfigType("yml")               // 配置文件的类型
	viper.AddConfigPath(workDir + "/config") // 配置文件的路径
	err := viper.ReadInConfig()              // 读取
	if err != nil {
		return
	}
}
