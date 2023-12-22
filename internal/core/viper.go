package core

import (
	"fmt"
	"github.com/spf13/viper"
)

var DefaultConfig = "./configs/conf.yaml"

func InitConfig(configFile string) (err error) {
	if configFile == "" {
		configFile = DefaultConfig
	}

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Sprintf("viper read config error:%v", err))
		return
	}

	if err = viper.Unmarshal(&GlobalConfig); err != nil {
		fmt.Println("config unmarshal error:", err)
		return
	}
	return nil
}
