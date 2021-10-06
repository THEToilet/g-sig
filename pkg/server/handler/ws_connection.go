package handler

import (
	"github.com/gobwas/ws/wsutil"
	"github.com/rs/zerolog"
	"net"
	"time"
)

type WSConnection struct {
	conn           net.Conn
	receiveMessage chan []byte
	sendMessage    chan []byte
	logger         *zerolog.Logger
}

func (w *WSConnection) selector() {
	pingTimer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-pingTimer.C:
			pingTimer.Reset(10 * time.Second)
		case msg, ok := <-receiveMessage:
			w.logger.Info().Msg(string(msg))
			if !ok {
				w.logger.Fatal().Msg("d")
			}
		case msg, ok := <-sendMessage:
			w.logger.Info().Msg(string(msg))
			if !ok {
				w.logger.Fatal().Msg("")
			}
			// ここでメッセージを返す
			w.logger.Debug().Msg(string(responseMessage))
			err = wsutil.WriteServerMessage(conn, op, responseMessage)
			if err != nil {
				w.logger.Fatal().Err(err)
			}

		}
	}
	// TODO: エラーメッセージをクライアント側へ送る
	// TODO: 送る
}

func (w *WSConnection) receiver() {
	defer func() { w.conn.Close() }()
	for {
		if err := w.conn.SetReadDeadline(time.Now().Add(time.Duration(60) * time.Second)); err != nil {
			w.logger.Fatal().Err(err)
			return
		}
		msg, op, err := wsutil.ReadClientData(w.conn)
		if err != nil {
			w.logger.Fatal().Err(err)
			return
		}
		w.logger.Info().Msg(string(msg))
		receiveMessage <- msg
	}
}
