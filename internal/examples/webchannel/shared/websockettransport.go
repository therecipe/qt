package shared

import (
	"fmt"

	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/webchannel"
	"github.com/StarAurryon/qt/websockets"
)

type WebSocketTransport struct {
	webchannel.QWebChannelAbstractTransport

	_ func(message string) `slot:"textMessageReceived"`

	//TODO: custom constructors

	m_socket *websockets.QWebSocket
}

func (t *WebSocketTransport) Init(socket *websockets.QWebSocket) {
	t.m_socket = socket

	socket.ConnectTextMessageReceived(t.textMessageReceived)
	socket.ConnectDisconnected(t.DeleteLater)

	t.ConnectSendMessage(t.sendMessage)
}

func (t *WebSocketTransport) sendMessage(message *core.QJsonObject) {
	doc := core.NewQJsonDocument2(message)
	t.m_socket.SendTextMessage(doc.ToJson2(core.QJsonDocument__Compact).ConstData())
}

func (t *WebSocketTransport) textMessageReceived(messageData string) {
	var error *core.QJsonParseError
	message := core.QJsonDocument_FromJson(core.NewQByteArray2(messageData, -1), error)

	if error.Error() != 0 {
		fmt.Printf("Failed to parse text message as JSON object: %v\nError is: %v\n", messageData, error.ErrorString())
		return
	} else if !message.IsObject() {
		fmt.Printf("Received JSON message that is not an object: %v\n", messageData)
		return
	}
	t.MessageReceived(message.Object(), t)
}
