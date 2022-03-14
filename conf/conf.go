package conf

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//定义配置文件结构体
//有什么方式向spring一样的注解的呢???
type Configuration struct {
	DB struct {
		Driver string `json:"driver"`
		Addr   string `json:"addr"`
	} `json:"db"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`
	Admin struct {
		Name     string `json:"name"`
		Mail     string `json:"mail"`
		Password string `json:"password"`
	} `json:"admin"`
	Address   string `json:"address"`
	Lang      string `json:"lang"`
	SecretKey string `json:"secretkey"`
}

var conf *Configuration

//解析配置文件
//微服务是否有配置中心???
func Config() *Configuration {
	if conf != nil {
		return conf
	}
	var err error
	viper.SetConfigName("config")
	viper.AddConfigPath("/Users/chenjia/Documents/code/go_code/my-guora/conf")
	viper.SetConfigType("yaml")
	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("config file error %s\n", err)
		os.Exit(1)
	}
	if err = viper.Unmarshal(&conf); err != nil {
		fmt.Printf("config file error %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Configuration.conf", conf)
	return conf
}
