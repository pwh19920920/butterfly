package server

import (
	"github.com/pwh19920920/butterfly/config"
)

type Conf struct {
	Server config.ServerConfig `yaml:"server"`
}

var conf *Conf

// 初始化日志, 命令行 > 配置文件 > 默认值
func load() *Conf {
	conf = new(Conf)
	config.LoadConf(&conf)
	return conf
}

func GetConf() config.ServerConfig {
	if nil == conf {
		conf = load()
	}
	return conf.Server
}
