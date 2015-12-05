package webchannel

//#include "webchannel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QWebChannelAbstractTransport_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWebChannelAbstractTransport) QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport {
	return ptr
}

func (ptr *QWebChannelAbstractTransport) ConnectMessageReceived(f func(message *core.QJsonObject, transport *QWebChannelAbstractTransport)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannelAbstractTransport::messageReceived")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ConnectMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageReceived", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectMessageReceived() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannelAbstractTransport::messageReceived")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DisconnectMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageReceived")
	}
}

//export callbackQWebChannelAbstractTransportMessageReceived
func callbackQWebChannelAbstractTransportMessageReceived(ptrName *C.char, message unsafe.Pointer, transport unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannelAbstractTransport::messageReceived")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "messageReceived").(func(*core.QJsonObject, *QWebChannelAbstractTransport))(core.NewQJsonObjectFromPointer(message), NewQWebChannelAbstractTransportFromPointer(transport))
}

func (ptr *QWebChannelAbstractTransport) SendMessage(message core.QJsonObject_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannelAbstractTransport::sendMessage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_SendMessage(ptr.Pointer(), core.PointerFromQJsonObject(message))
	}
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransport() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannelAbstractTransport::~QWebChannelAbstractTransport")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
