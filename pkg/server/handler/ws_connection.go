package handler

import (
	"context"
	"encoding/json"
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
	isRegistered     bool
}

func NewWSConnection(conn net.Conn, receiveMessage chan []byte, sendingMessage chan []byte, signalingUseCase *application.SignalingUseCase, logger *zerolog.Logger) *WSConnection {
	return &WSConnection{
		conn:             conn,
		receiveMessage:   receiveMessage,
		sendingMessage:   sendingMessage,
		signalingUseCase: signalingUseCase,
		logger:           logger,
		isRegistered:     false,
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
		w.logger.Debug().Caller().Msg("ddddddd")
		// TODO userをデリートする
	}()

	// TODO 疎通確認のping pong
	// Labeled Break
L:
	for {
		select {
		case <-pingTimer.C:
			w.logger.Info().Msg("ping")
			// send ping message
			/* TODO: 要リファクタリング */
			status := w.makeStatusMessage("ping", "200", "please send pong")
			responseMessage, err := json.Marshal(status)
			if err != nil {
				w.logger.Fatal().Err(err)
			}
			if err := w.sendMessage(responseMessage); err != nil {
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
			// TODO バグがある
			/* 2021-10-06T20:54:16+09:00 | INFO  | {"code":"400","message":"Invalid Message","type":"undefined"}
			2021-10-06T20:54:16+09:00 | DEBUG | Unmarshall message
			2021-10-06T20:54:16+09:00 | DEBUG | Invalid Message rawMessage=eyJjb2RlIjoiNDAwIiwibWVzc2FnZSI6IkludmFsaWQgTWVzc2FnZSIsInR5cGUiOiJ1bmRlZmluZWQifQ== */

			//w.handleMessage(msg, pongTimer)
		case msg, ok := <-w.sendingMessage:
			w.logger.Info().Msg(string(msg))
			if !ok {
				w.logger.Fatal().Msg("")
				break L
			}
			// クライアントへメッセージ送信
			if err := w.sendMessage(msg); err != nil {
				break L
			}
		}
	}
	// TODO サーバが死んだことによる
	// TODO: エラーメッセージをクライアント側へ送る
	if err := w.sendMessage([]byte{}); err != nil {
		w.logger.Debug().Msg("")
		return
	}
	cancel()
}

func (w *WSConnection) receiver(ctx context.Context) {
	defer func() {
		w.logger.Debug().Msg("ws id close")
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
		w.logger.Info().Msg(string(msg))
		w.receiveMessage <- msg
	}
	<-ctx.Done()
}
