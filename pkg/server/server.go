package server

import (
	"fmt"
	"g-sig/pkg/domain/application"
	"g-sig/pkg/server/handler"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

func NewServer(signalingUseCase *application.SignalingUseCase, logger *zerolog.Logger) *http.Server {
	mux := http.NewServeMux()
	connections := handler.NewConnections(logger)
	signalingHandler := handler.NewSignalingHandler(signalingUseCase, connections, logger)

	mux.HandleFunc("/signaling", signalingHandler.Signaling)
	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "HELLOOOOOOO")
		logger.Info().Msg(" / Access is Successful")
	})
	mux.HandleFunc("/stun", func(writer http.ResponseWriter, r *http.Request) {
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		fmt.Fprintf(writer, "stun:")
		logger.Info().Msg(r.RemoteAddr)
		fmt.Fprintf(writer, r.RemoteAddr)
		logger.Info().Msg(" /stun Access is Successful")
	})
	server := &http.Server{
		// NOTE: ここ変えるならクライアントも変えなければならない
		// NOTE: 127.0.0.1 では繋がらないが、localhostは繋がる
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server
}
