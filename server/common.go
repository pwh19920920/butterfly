package server

import (
	"github.com/pwh19920920/butterfly/config"
	"github.com/spf13/viper"
)

const defaultEngineMode = "debug"
const defaultServerAddr = ":8080"
const defaultServerName = "butterfly"

type Conf struct {
	Server config.ServerConfig `yaml:"server"`
}

var conf *Conf

// 初始化日志, 命令行 > 配置文件 > 默认值
func load() *Conf {
	conf = new(Conf)

	// 优先赋予默认值
	viper.SetDefault("server.engineMode", defaultEngineMode)
	viper.SetDefault("server.serverAddr", defaultServerAddr)
	viper.SetDefault("server.serverName", defaultServerName)

	// 加载配置
	config.LoadConf(&conf, config.GetOptions().ConfigFilePath)
	return conf
}

func getConf() config.ServerConfig {
	if nil == conf {
		conf = load()
	}
	return conf.Server
}
