package logger

import (
	"bufio"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pwh19920920/butterfly/config"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

const defaultLoggerDateFormat string = "2006-01-02 15:04:05.999999999"
const infoLog string = "INFO"
const warnLog string = "WARN"
const errorLog string = "ERROR"

// 初始化日志系统
func initLogger(conf config.LoggerConfig) {
	log.Info("初始化日志服务")

	// 日期格式化
	format := conf.DateFormat
	if format == "" {
		format = defaultLoggerDateFormat
	}

	// 创建扩展，日志切割
	hook := newLfsHook(conf.LogPath, conf.FileName, format)
	if nil != hook {
		log.AddHook(hook)
	}

	// 设置默认日志级别
	level, err := log.ParseLevel(conf.Level)
	if err == nil {
		log.SetLevel(level)
	}

	// 关闭标准输出
	if conf.CloseConsoleOut {
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Error("Open Src File err", err)
		}
		writer := bufio.NewWriter(src)
		log.SetOutput(writer)
	}

	// 设置日志格式
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: format})
}

// 创建日志扩展
func newLfsHook(logDir string, filename string, dateFormat string) log.Hook {
	if logDir == "" {
		logDir = "./logs"
	}

	if filename == "" {
		filename = "server"
	}

	// 创建writer
	infoWriter, err := createWriter(logDir, filename, infoLog)

	if err != nil {
		log.Errorf("create logger infoWriter failure: %v", err)
		return nil
	}

	// 创建writer
	warnWriter, err := createWriter(logDir, filename, warnLog)
	if err != nil {
		log.Errorf("create logger warnWriter failure: %v", err)
		return nil
	}

	// 创建writer
	errorWriter, err := createWriter(logDir, filename, errorLog)
	if err != nil {
		log.Errorf("create logger errorWriter failure: %v", err)
		return nil
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: infoWriter,
		log.InfoLevel:  infoWriter,
		log.TraceLevel: infoWriter,
		log.WarnLevel:  warnWriter,
		log.ErrorLevel: errorWriter,
		log.FatalLevel: errorWriter,
		log.PanicLevel: errorWriter,
	}, &log.JSONFormatter{TimestampFormat: dateFormat})
	return lfsHook
}

// 创建appender
// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次

// WithMaxAge和WithRotationCount二者只能设置一个, WithMaxAge设置文件清理前的最长保存时间, WithRotationCount设置文件清理前最多保存的个数.
// WithMaxAge(time.Hour*24),
// WithRotationCount(maxRemainCnt),
func createWriter(logDir string, filename string, level string) (*rotatelogs.RotateLogs, error) {
	// 创建writer
	infoLogPath := fmt.Sprintf("%s/%s-%s-%s.log", logDir, filename, "%Y%m%d", level)
	infoLogLinkPath := fmt.Sprintf("%s/%s-%s.log", logDir, filename, level)

	option := make([]rotatelogs.Option, 0)
	option = append(option, rotatelogs.WithLinkName(infoLogLinkPath))
	option = append(option, rotatelogs.WithRotationTime(time.Hour*24))

	// 创建
	infoWriter, err := rotatelogs.New(infoLogPath, option...)

	if err != nil {
		return nil, err
	}

	return infoWriter, nil
}
