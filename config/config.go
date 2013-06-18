package config

import (
	"encoding/json"
	"fmt"
	"github.com/tabalt/golib/logger"
	"io/ioutil"
)

var configFile string

var Config map[string]interface{}

//从文件读取配置
func File2Config(fileName string) bool {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.Fatal(err.Error())
	}
	result := true
	if err := json.Unmarshal(file, &Config); err != nil {
		result = false
	}
	goConfig, err := Config.(map[string]interface{})
	if err != nil {
		fmt.Println(goConfig)
	}
	return result
}

//将配置写入文件
func Config2File(fileName string) bool {
	configContent, err := json.Marshal(Config)
	if err != nil {
		logger.Fatal(err.Error())
	}
	if err := ioutil.WriteFile(fileName, configContent, 0x777); err != nil {
		logger.Fatal(err.Error())
	}
	return true
}

//从默认配置文件读取配置
func Init() bool {
	configFile = "./config.json"
	return File2Config(configFile)
}

//保存配置到文件
func Save() bool {
	return Config2File(configFile)
}

//获取指定配置项
func Get(key string) interface{} {
	var conf interface{}
	conf, err := Config[key]
	if !err {
		logger.Fatal("配置项“" + key + "”不存在")
	}
	return conf
}

//获取指定全部配置
func GetAll() interface{} {
	return Config
}

//设置指定配置项
func Set(key string, value interface{}) bool {
	Config[key] = value
	_, err := Config[key]
	if !err {
		logger.Fatal("配置项“" + key + "”不存在")
	}
	return false
}

//TODO　Set可选是否保存配置文件
