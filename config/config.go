package config

import (
	"errors"
	"github.com/pwh19920920/butterfly/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

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
	// 配置读取
	configPath, configName, configType, err := splitViperConfig(configFilePath)
	if err != nil {
		logrus.Warn("Loading", reflect.TypeOf(conf), "配置文件不正确")
		return
	}

	// 设置viper配置
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	// 加载配置
	err = viper.ReadInConfig()
	if err != nil {
		logrus.Warn("Loading", reflect.TypeOf(conf), "配置文件加载失败")
	}

	// 反序列化，配置文件加载不到也要反序列化，可能会有默认配置
	err = viper.Unmarshal(conf)
	if err != nil {
		logrus.Warn("Loading", reflect.TypeOf(conf), "配置文件序列化失败")
		return
	}
}

// 切割文件
func splitViperConfig(configFilePath string) (configPath string, configName string, configType string, e error) {
	lastIndexDirectory := strings.LastIndex(configFilePath, "/")
	lastIndexDot := strings.LastIndex(configFilePath, ".")
	if lastIndexDot == -1 {
		return "", "", "", errors.New("fileType is error")
	}

	configPath = "."
	if lastIndexDirectory != -1 {
		configPath = helper.StringHelper.SubString(configFilePath, 0, lastIndexDirectory)
	}

	configName = helper.StringHelper.SubString(configFilePath, lastIndexDirectory+1, lastIndexDot)
	configType = helper.StringHelper.SubString(configFilePath, lastIndexDot+1, len(configFilePath))
	return configPath, configName, configType, nil
}
