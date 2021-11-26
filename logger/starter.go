package logger

import (
	"github.com/pwh19920920/butterfly/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const defaultLoggerDateFormat string = "2006-01-02 15:04:05.999999999"

type Conf struct {
	Logger config.LoggerConfig `yaml:"logger"`
}

var ConsoleLogger = createConsoleLogger()

// 初始化默认日志
func init() {
	initDefaultLogger()
}

// 初始化日志
func initDefaultLogger() {
	conf := new(Conf)

	// 优先赋予默认值
	viper.SetDefault("logger.dateFormat", defaultLoggerDateFormat)

	// 先设置
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: defaultLoggerDateFormat})

	// 加载配置
	config.LoadConf(&conf)

	// 加载日志
	initLogger(conf.Logger)
}

// CreateConsoleLogger 初始化默认日志
func createConsoleLogger() *log.Logger {
	// 先设置
	defaultLogger := log.New()
	defaultLogger.SetFormatter(&log.JSONFormatter{TimestampFormat: defaultLoggerDateFormat})
	return defaultLogger
}
