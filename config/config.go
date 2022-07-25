package config

import (
	"errors"
	"fmt"
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
	EngineMode     string            `yaml:"engineMode"`
	ServerAddr     string            `yaml:"serverAddr"`
	ServerName     string            `yaml:"serverName"`
	HtmlGlobs      []string          `yaml:"htmlGlobs"`
	Statics        map[string]string `yaml:"statics"`
	MethodOverride bool              `yaml:"methodOverride"`
}

var initFlag = false

const defaultEngineMode = "debug"
const defaultServerAddr = ":8080"
const defaultServerName = "butterfly"

func LoadConf(conf interface{}) {
	if initFlag {
		// 反序列化，配置文件加载不到也要反序列化，可能会有默认配置
		err := viper.Unmarshal(conf)
		if err != nil {
			logrus.Panic("Loading", reflect.TypeOf(conf), "配置文件序列化失败")
		}
		return
	}

	// 配置读取
	configPath, configName, configType, err := splitViperConfig(GetOptions().ConfigFilePath)
	if err != nil {
		logrus.Warn("Loading", reflect.TypeOf(conf), "配置文件不正确")
		return
	}

	// 优先赋予默认值
	viper.SetDefault("server.engineMode", defaultEngineMode)
	viper.SetDefault("server.serverAddr", defaultServerAddr)
	viper.SetDefault("server.serverName", defaultServerName)

	// 设置viper配置
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	// 加载配置
	err = viper.ReadInConfig()
	logrus.Infof("系统启动, 开始加载配置文件: %s/%s.%s", configPath, configName, configType)
	if err != nil {
		logrus.Warnf("系统启动, 加载配置文件: %s/%s.%s 失败", configPath, configName, configType)
	}

	// 设置viper配置
	envConf := fmt.Sprintf("%s-%s", configName, viper.Get("server.engineMode"))
	viper.SetConfigName(envConf)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	err = viper.MergeInConfig()
	logrus.Infof("系统启动, 开始加载配置文件: %s/%s.%s", configPath, envConf, configType)
	if err != nil {
		logrus.Warnf("系统启动, 加载配置文件: %s/%s.%s 失败, 不影响系统正常启动", configPath, envConf, configType)
	}

	// 设置标记位置
	initFlag = true

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
