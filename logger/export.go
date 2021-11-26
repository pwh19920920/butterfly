package logger

import (
	"context"
	"github.com/pwh19920920/butterfly/common"
	"github.com/sirupsen/logrus"
)

func Trace(ctx context.Context, args ...interface{}) {
	entry(ctx).Trace(args...)
}

// TraceFormat 格式化输出
func TraceFormat(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Tracef(format, args...)
}

func TraceEntry(ctx context.Context, entry *logrus.Entry, args ...interface{}) {
	entryForEntry(ctx, entry).Trace(args...)
}

func TraceEntryFormat(ctx context.Context, entry *logrus.Entry, format string, args ...interface{}) {
	entryForEntry(ctx, entry).Tracef(format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	entry(ctx).Info(args...)
}

// InfoFormat 格式化输出
func InfoFormat(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Infof(format, args...)
}

func InfoEntry(ctx context.Context, entry *logrus.Entry, args ...interface{}) {
	entryForEntry(ctx, entry).Info(args...)
}

func InfoEntryFormat(ctx context.Context, entry *logrus.Entry, format string, args ...interface{}) {
	entryForEntry(ctx, entry).Infof(format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	entry(ctx).Error(args...)
}

// ErrorFormat 格式化输出
func ErrorFormat(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Errorf(format, args...)
}

func ErrorEntry(ctx context.Context, entry *logrus.Entry, args ...interface{}) {
	entryForEntry(ctx, entry).Error(args...)
}

func ErrorEntryFormat(ctx context.Context, entry *logrus.Entry, format string, args ...interface{}) {
	entryForEntry(ctx, entry).Errorf(format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	entry(ctx).Warn(args...)
}

// WarnFormat 格式化输出
func WarnFormat(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Warnf(format, args...)
}

func WarnEntry(ctx context.Context, entry *logrus.Entry, args ...interface{}) {
	entryForEntry(ctx, entry).Warn(args...)
}

func WarnEntryFormat(ctx context.Context, entry *logrus.Entry, format string, args ...interface{}) {
	entryForEntry(ctx, entry).Warnf(format, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	entry(ctx).Fatal(args...)
}

// FatalFormat 格式化输出
func FatalFormat(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Fatalf(format, args...)
}

func FatalEntry(ctx context.Context, entry *logrus.Entry, args ...interface{}) {
	entryForEntry(ctx, entry).Fatal(args...)
}

func FatalEntryFormat(ctx context.Context, entry *logrus.Entry, format string, args ...interface{}) {
	entryForEntry(ctx, entry).Fatalf(format, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	entry(ctx).Debug(args...)
}

// DebugFormat 格式化输出
func DebugFormat(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Debugf(format, args...)
}

func DebugEntry(ctx context.Context, entry *logrus.Entry, args ...interface{}) {
	entryForEntry(ctx, entry).Debug(args...)
}

func DebugEntryFormat(ctx context.Context, entry *logrus.Entry, format string, args ...interface{}) {
	entryForEntry(ctx, entry).Debugf(format, args...)
}

func Panic(ctx context.Context, args ...interface{}) {
	entry(ctx).Panic(args...)
}

// PanicFormat 格式化输出
func PanicFormat(ctx context.Context, format string, args ...interface{}) {
	entry(ctx).Panicf(format, args...)
}

func PanicEntry(ctx context.Context, entry *logrus.Entry, args ...interface{}) {
	entryForEntry(ctx, entry).Panic(args...)
}

func PanicEntryFormat(ctx context.Context, entry *logrus.Entry, format string, args ...interface{}) {
	entryForEntry(ctx, entry).Panicf(format, args...)
}

func entry(ctx context.Context) *logrus.Entry {
	traceId, _ := ctx.Value(common.TraceIdHeaderKey).(string)
	spanId, _ := ctx.Value(common.TraceSpanHeaderKey).(string)
	return logrus.WithField(common.TraceIdLogKey, traceId).WithField(common.TraceSpanLogKey, spanId)
}

func entryForEntry(ctx context.Context, entry *logrus.Entry) *logrus.Entry {
	traceId, _ := ctx.Value(common.TraceIdHeaderKey).(string)
	spanId, _ := ctx.Value(common.TraceSpanHeaderKey).(string)
	return entry.WithField(common.TraceIdLogKey, traceId).WithField(common.TraceSpanLogKey, spanId)
}
