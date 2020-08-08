package shared

import (
	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/websockets"
)

type WebSocketClientWrapper struct {
	core.QObject

	_ func(client *WebSocketTransport) `signal:"clientConnected"`

	_ func() `slot:"handleNewConnection"`

	//TODO: custom constructors

	m_server *websockets.QWebSocketServer
}

func (w *WebSocketClientWrapper) Init(server *websockets.QWebSocketServer) {
	w.m_server = server

	server.ConnectNewConnection(w.handleNewConnection)
}

func (w *WebSocketClientWrapper) handleNewConnection() {
	var t = NewWebSocketTransport(nil)
	t.Init(w.m_server.NextPendingConnection())
	w.ClientConnected(t)
}
