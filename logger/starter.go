package logger

import (
	"github.com/pwh19920920/butterfly/config"
	"github.com/spf13/viper"
)

const defaultLoggerDateFormat string = "2006-01-02 15:04:05.999999999"

type Conf struct {
	Logger config.LoggerConfig `yaml:"logger"`
}

func init() {
	start(config.GetOptions().ConfigFilePath)
}

// 初始化日志
func start(configFilePath string) {
	conf := new(Conf)

	// 优先赋予默认值
	viper.SetDefault("logger.dateFormat", defaultLoggerDateFormat)

	// 加载配置
	config.LoadConf(&conf, configFilePath)

	// 加载日志
	initLogger(conf.Logger)
}
