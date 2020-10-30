package config

import (
	"github.com/pwh19920920/butterfly/helper"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

const defaultConfigPath string = "src/resources/application.yml"

type LoggerConfig struct {
	Level           string
	LogPath         string
	FileName        string
	DateFormat      string
	CloseConsoleOut bool
}

type ServerConfig struct {
	EngineMode  string
	ServerAddr  string
	ServiceName string
	HtmlGlobs   []string
	Statics     map[string]string
}

func LoadConf(conf interface{}, configFilePath string) {
	if configFilePath == "" {
		configFilePath = defaultConfigPath
	}

	// 检查文件
	exists, err := helper.FileHelper.PathExists(configFilePath)
	if err != nil || !exists {
		logrus.Warn("Loading", reflect.TypeOf(conf), "配置文件不存在")
		return
	}

	// 加载文件
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		logrus.Warn("Loading", reflect.TypeOf(conf), "配置文件加载失败")
		return
	}

	// 序列化
	err = yaml.Unmarshal(configFile, conf)
	if err != nil {
		logrus.Warn("Loading", reflect.TypeOf(conf), "配置文件加载失败")
		return
	}
}
