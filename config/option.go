package config

import (
	"flag"
)

var options *Options

type Options struct {
	ConfigFilePath string
	EngineMode     string
	ServerAddr     string
}

func parseOptions() *Options {
	opts := new(Options)
	flag.StringVar(&opts.ConfigFilePath, "configFilePath", "", "配置文件地址")
	flag.StringVar(&opts.EngineMode, "engineMode", "", "引擎模式：debug，release，test")
	flag.StringVar(&opts.ServerAddr, "serverAddr", "", "启动地址：格式为ip地址:端口, 地址无限制则为:端口")
	flag.Parse()
	return opts
}

// 获取启动命令配置
func GetOptions() *Options {
	if options == nil {
		options = parseOptions()
	}
	return options
}
