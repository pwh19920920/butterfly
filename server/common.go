package server

import "github.com/pwh19920920/butterfly/config"

const defaultEngineMode = "debug"
const defaultServerAddr = ":8080"

type Conf struct {
	Server config.ServerConfig `yaml:"server"`
}

var conf *Conf

// 初始化日志, 命令行 > 配置文件 > 默认值
func load() *Conf {
	conf = new(Conf)
	config.LoadConf(&conf, config.GetOptions().ConfigFilePath)

	// 命令行优先级大于配置文件
	if config.GetOptions().EngineMode != "" {
		conf.Server.EngineMode = config.GetOptions().EngineMode
	}

	if config.GetOptions().ServerAddr != "" {
		conf.Server.ServerAddr = config.GetOptions().ServerAddr
	}

	// 空值判断
	if conf.Server.EngineMode == "" {
		conf.Server.EngineMode = defaultEngineMode
	}

	if conf.Server.ServerAddr == "" {
		conf.Server.ServerAddr = defaultServerAddr
	}

	return conf
}

func getConf() config.ServerConfig {
	if nil == conf {
		conf = load()
	}
	return conf.Server
}
