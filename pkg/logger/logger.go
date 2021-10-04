package logger

import (
	"fmt"
	"g-sig/pkg/config"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

func NewLogger(config *config.Config) (*zerolog.Logger, error) {
	// Debugモードでなければ自動的にInfoモードになる
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if isDebug(config.LogInfo.Level) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger := zerolog.New(customFormat()).With().Timestamp().Logger()
	return &logger, nil
}

// Logの出力形式を調整
func customFormat() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
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
