package webchannel

//#include "webchannel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWebChannelAbstractTransport struct {
	core.QObject
}

type QWebChannelAbstractTransport_ITF interface {
	core.QObject_ITF
	QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport
}

func PointerFromQWebChannelAbstractTransport(ptr QWebChannelAbstractTransport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannelAbstractTransport_PTR().Pointer()
	}
	return nil
}

func NewQWebChannelAbstractTransportFromPointer(ptr unsafe.Pointer) *QWebChannelAbstractTransport {
	var n = new(QWebChannelAbstractTransport)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QWebChannelAbstractTransport_") {
		n.SetObjectName("QWebChannelAbstractTransport_" + qt.Identifier())
	}
	return n
}

func (ptr *QWebChannelAbstractTransport) QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport {
	return ptr
}

func (ptr *QWebChannelAbstractTransport) ConnectMessageReceived(f func(message *core.QJsonObject, transport *QWebChannelAbstractTransport)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ConnectMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageReceived", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectMessageReceived() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DisconnectMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageReceived")
	}
}

//export callbackQWebChannelAbstractTransportMessageReceived
func callbackQWebChannelAbstractTransportMessageReceived(ptrName *C.char, message unsafe.Pointer, transport unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::messageReceived")

	var signal = qt.GetSignal(C.GoString(ptrName), "messageReceived")
	if signal != nil {
		signal.(func(*core.QJsonObject, *QWebChannelAbstractTransport))(core.NewQJsonObjectFromPointer(message), NewQWebChannelAbstractTransportFromPointer(transport))
	}

}

func (ptr *QWebChannelAbstractTransport) SendMessage(message core.QJsonObject_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::sendMessage")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_SendMessage(ptr.Pointer(), core.PointerFromQJsonObject(message))
	}
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransport() {
	defer qt.Recovering("QWebChannelAbstractTransport::~QWebChannelAbstractTransport")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
