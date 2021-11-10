package handler

import (
	"context"
	"g-sig/pkg/domain/application"
	"github.com/gobwas/ws/wsutil"
	"github.com/rs/zerolog"
	"net"
	"time"
)

type WSConnection struct {
	conn             net.Conn
	receiveMessage   chan []byte
	sendingMessage   chan []byte
	signalingUseCase *application.SignalingUseCase
	logger           *zerolog.Logger
	connections      *Connections
	isRegistered     bool
	userID           string
	destination      string
}

func NewWSConnection(conn net.Conn, receiveMessage chan []byte, sendingMessage chan []byte, signalingUseCase *application.SignalingUseCase, connections *Connections, logger *zerolog.Logger) *WSConnection {
	return &WSConnection{
		conn:             conn,
		receiveMessage:   receiveMessage,
		sendingMessage:   sendingMessage,
		signalingUseCase: signalingUseCase,
		connections:      connections,
		logger:           logger,
		isRegistered:     false,
		userID:           "",
		destination:      "",
	}
}

func stopTimer(timer *time.Timer) {
	if !timer.Stop() {
		<-timer.C
	}
}

func (w *WSConnection) selector(ctx context.Context, cancel context.CancelFunc) {
	pingTimer := time.NewTicker(10 * time.Second)
	pongTimer := time.NewTimer(10 * time.Second)
	stopTimer(pongTimer)
	defer func() {
		// TODO　ここ呼ばれない
		stopTimer(pongTimer)
		pingTimer.Stop()
		w.logger.Debug().Caller().Msg("selector is close")
		if err := w.signalingUseCase.Delete(w.userID); err != nil {
			w.logger.Debug().Msg("Delete error")
		}
	}()

	// TODO 疎通確認のping pong
	// Labeled Break
L:
	for {
		select {
		case <-pingTimer.C:
			w.logger.Info().Msg("ping")
			requestMessage, err := w.makePingMessage()
			if err != nil {
				w.logger.Info().Msg("ping make error")
				break L
			}
			if err := w.sendMessage(&w.conn, requestMessage); err != nil {
				break L
			}
			pongTimer.Reset(10 * time.Second)
		case <-pongTimer.C:
			w.logger.Info().Msg("pong is failed")
			break L
		case msg, ok := <-w.receiveMessage:
			w.logger.Info().Msg(string(msg))
			if !ok {
				w.logger.Fatal().Msg("d")
				break L
			}
			w.handleMessage(msg, pongTimer)
		case msg, ok := <-w.sendingMessage:
			w.logger.Info().Caller().Msg(string(msg))
			if !ok {
				w.logger.Fatal().Msg("")
				break L
			}
			// クライアントへメッセージ送信
			if err := w.sendMessage(&w.conn, msg); err != nil {
				break L
			}
		}
	}
	// TODO サーバが死んだことによる
	// TODO: エラーメッセージをクライアント側へ送る
	if err := w.sendMessage(&w.conn, []byte{}); err != nil {
		w.logger.Debug().Msg("")
		return
	}
	cancel()
}

func (w *WSConnection) receiver(ctx context.Context) {
	defer func() {
		w.logger.Debug().Msg("WebSocket is closed")
		w.conn.Close()
	}()
L:
	for {
		if err := w.conn.SetReadDeadline(time.Now().Add(time.Duration(60) * time.Second)); err != nil {
			w.logger.Fatal().Err(err)
			break L
		}
		msg, _, err := wsutil.ReadClientData(w.conn)
		if err != nil {
			w.logger.Fatal().Err(err)
			break L
		}
		w.logger.Info().Caller().Msg(string(msg))
		w.receiveMessage <- msg
	}
	<-ctx.Done()
}
