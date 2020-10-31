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
	Level           string `yaml:"level"`
	LogPath         string `yaml:"logPath"`
	FileName        string `yaml:"fileName"`
	DateFormat      string `yaml:"dateFormat"`
	CloseConsoleOut bool   `yaml:"closeConsoleOut"`
}

type ServerConfig struct {
	EngineMode  string            `yaml:"engineMode"`
	ServerAddr  string            `yaml:"serverAddr"`
	ServiceName string            `yaml:"serviceName"`
	HtmlGlobs   []string          `yaml:"htmlGlobs"`
	Statics     map[string]string `yaml:"statics"`
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
