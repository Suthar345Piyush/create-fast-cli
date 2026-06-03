// Package logger provides a wrapper around zap logger for create-fast-cli, generated project get this by the logger template

package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var global *zap.SugaredLogger

// init function initialises the global logger, call from root.go

// verbose = true - it enable debug level output

func Init(verbose bool) error {

	level := zapcore.InfoLevel

	if verbose {
		level = zapcore.DebugLevel
	}

	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderCfg.TimeKey = ""

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stderr),
		level,
	)

	global = zap.New(core).Sugar()

	return nil

}

func must() *zap.SugaredLogger {
	if global == nil {
		panic("logger.Init() has not been called")
	}
	return global
}

// Debug logs a debug message (only visible in --verbose mode).

func Debug(args ...any) {
	must().Debug(args...)
}

// Debugf logs a formatted debug message.

func Debugf(format string, args ...any) {
	must().Debugf(format, args...)
}

// Info logs an informational message.

func Info(args ...any) {
	must().Info(args...)
}

// Infof logs a formatted informational message.

func Infof(format string, args ...any) {
	must().Infof(format, args...)
}

// Warn logs a warning.

func Warn(args ...any) {
	must().Warn(args...)
}

// Warnf logs a formatted warning.

func Warnf(format string, args ...any) {
	must().Warnf(format, args...)
}

// Error logs an error message.
func Error(args ...any) {
	must().Error(args...)
}

// Errorf logs a formatted error message.

func Errorf(format string, args ...any) {
	must().Errorf(format, args...)
}

// Fatal logs a message then calls os.Exit(1).

func Fatal(args ...any) {
	must().Fatal(args...)
}

// Sync flushes any buffered log entries. Call before os.Exit.

func Sync() {
	_ = must().Sync()
}

// Printf satisfies the printf-style interface expected by some third-party libs.

func Printf(format string, args ...any) {
	must().Info(fmt.Sprintf(format, args...))
}
