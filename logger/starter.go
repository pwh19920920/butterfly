package logger

import (
	"github.com/pwh19920920/butterfly/config"
)

type Conf struct {
	Logger config.LoggerConfig `yaml:"logger"`
}

func init() {
	start(config.GetOptions().ConfigFilePath)
}

// 初始化日志
func start(configFilePath string) {
	conf := new(Conf)
	config.LoadConf(&conf, configFilePath)

	// 加载日志
	initLogger(conf.Logger)
}
