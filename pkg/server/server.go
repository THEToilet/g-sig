package server

import (
	"fmt"
	"g-sig/pkg/domain/application"
	"g-sig/pkg/server/handler"
	"github.com/rs/zerolog"
	"net/http"
)

func NewServer(signalingUseCase *application.SignalingUseCase, logger *zerolog.Logger) *http.Server{
	mux := http.NewServeMux()
	//mux.Handle("/signup", )
	logger.Info().Msg("HELLO")
	signalingHandler := handler.NewSignalingHandler(signalingUseCase, logger)

	mux.HandleFunc("/signaling", signalingHandler.Signaling)
	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request){
		fmt.Fprintf(writer, "HELLOOOOOOO")
		logger.Info().Msg(" / Access is Successful")
	})
	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	return server
}
