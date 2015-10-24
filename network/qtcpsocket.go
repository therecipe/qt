package network

//#include "qtcpsocket.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTcpSocket struct {
	QAbstractSocket
}

type QTcpSocketITF interface {
	QAbstractSocketITF
	QTcpSocketPTR() *QTcpSocket
}

func PointerFromQTcpSocket(ptr QTcpSocketITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocketPTR().Pointer()
	}
	return nil
}

func QTcpSocketFromPointer(ptr unsafe.Pointer) *QTcpSocket {
	var n = new(QTcpSocket)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTcpSocket_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTcpSocket) QTcpSocketPTR() *QTcpSocket {
	return ptr
}

func NewQTcpSocket(parent core.QObjectITF) *QTcpSocket {
	return QTcpSocketFromPointer(unsafe.Pointer(C.QTcpSocket_NewQTcpSocket(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QTcpSocket) DestroyQTcpSocket() {
	if ptr.Pointer() != nil {
		C.QTcpSocket_DestroyQTcpSocket(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
