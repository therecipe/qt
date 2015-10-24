package webchannel

//#include "qwebchannelabstracttransport.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWebChannelAbstractTransport struct {
	core.QObject
}

type QWebChannelAbstractTransportITF interface {
	core.QObjectITF
	QWebChannelAbstractTransportPTR() *QWebChannelAbstractTransport
}

func PointerFromQWebChannelAbstractTransport(ptr QWebChannelAbstractTransportITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannelAbstractTransportPTR().Pointer()
	}
	return nil
}

func QWebChannelAbstractTransportFromPointer(ptr unsafe.Pointer) *QWebChannelAbstractTransport {
	var n = new(QWebChannelAbstractTransport)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWebChannelAbstractTransport_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWebChannelAbstractTransport) QWebChannelAbstractTransportPTR() *QWebChannelAbstractTransport {
	return ptr
}

func (ptr *QWebChannelAbstractTransport) SendMessage(message core.QJsonObjectITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_SendMessage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQJsonObject(message)))
	}
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransport() {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
