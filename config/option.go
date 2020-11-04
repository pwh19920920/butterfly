package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var options *Options

const defaultConfigPath string = "src/resources/application.yml"

type Options struct {
	ConfigFilePath string
}

func parseOptions() *Options {
	opts := new(Options)
	pflag.StringVar(&opts.ConfigFilePath, "configFilePath", defaultConfigPath, "配置文件地址")
	pflag.String("server.engineMode", "", "引擎模式：debug，release，test")
	pflag.String("server.serverAddr", "", "启动地址：格式为ip地址:端口, 地址无限制则为:端口")
	pflag.Parse()

	// 绑定到viper上
	_ = viper.BindPFlags(pflag.CommandLine)
	return opts
}

// 获取启动命令配置
func GetOptions() *Options {
	if options == nil {
		options = parseOptions()
	}
	return options
}
