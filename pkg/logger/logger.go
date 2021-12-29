package logger

import (
	"fmt"
	"g-sig/pkg/config"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
	"time"
)

func NewLogger(config *config.Config) (*zerolog.Logger, error) {

	writer := &lumberjack.Logger{
		Filename: "log/" + time.Now().Format("2006-01-02-1504") + ".log",
		// NOTE: Mb
		MaxSize: 0,
		// NOTE: 日
		MaxAge:     1,
		MaxBackups: 0,
		LocalTime:  false,
		Compress:   true,
	}

	// Debugモードでなければ自動的にInfoモードになる
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if isDebug(config.LogInfo.Level) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// NOTE: ログの出力を標準出力とファイルにする
	writers := io.MultiWriter(customFormat(), writer)
	logger := zerolog.New(writers).With().Timestamp().Logger()
	zerolog.TimeFieldFormat = "2006/01/02 15:04:05.000"
	return &logger, nil
}

// Logの出力形式を調整
func customFormat() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006/01/02 15:04:05.000"}
	output.FormatLevel = func(i interface{}) string {
		// 左詰め
		return strings.ToUpper(fmt.Sprintf("| %-5s |", i))
	}
	return output
}

func isDebug(logLevel string) bool {
	if logLevel == "DEBUG" {
		return true
	} else {
		return false
	}
}
