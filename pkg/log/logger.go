package log

import (
	"github.com/rs/zerolog"
	"os"
)

func NewLogger() (*zerolog.Logger, error){
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	// To log a human-friendly, colorized output, use zerolog.ConsoleWriter:
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout})
	return &logger, nil
}
:w
