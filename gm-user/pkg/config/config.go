package config

import (
	"github.com/spf13/viper"
	"log"
)

type globalConfig struct {
	//Mysql  *Mysql       `yaml:"mysql"`
	Server *server      `yaml:"server"`
	Redis  *redisConfig `yaml:"redis"`
	Zap    *zap         `yaml:"zap"`
}

var globalConf = new(globalConfig)

func Init() error {
	filePath := "config/app.yaml"
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("load config file failed ", err)
		return err
	}

	if err := viper.Unmarshal(globalConf); err != nil {
		log.Println("unmarshal config file failed ", err)
		return err
	}

	log.Printf("%+v\n", globalConf)

	return nil
}

func ServerConf() *server {
	return globalConf.Server
}

func RedisConf() *redisConfig {
	return globalConf.Redis
}

func ZapConf() *zap {
	return globalConf.Zap
}
