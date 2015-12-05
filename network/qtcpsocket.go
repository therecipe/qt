package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QTcpSocket struct {
	QAbstractSocket
}

type QTcpSocket_ITF interface {
	QAbstractSocket_ITF
	QTcpSocket_PTR() *QTcpSocket
}

func PointerFromQTcpSocket(ptr QTcpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func NewQTcpSocketFromPointer(ptr unsafe.Pointer) *QTcpSocket {
	var n = new(QTcpSocket)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTcpSocket_") {
		n.SetObjectName("QTcpSocket_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTcpSocket) QTcpSocket_PTR() *QTcpSocket {
	return ptr
}

func NewQTcpSocket(parent core.QObject_ITF) *QTcpSocket {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTcpSocket::QTcpSocket")
		}
	}()

	return NewQTcpSocketFromPointer(C.QTcpSocket_NewQTcpSocket(core.PointerFromQObject(parent)))
}

func (ptr *QTcpSocket) DestroyQTcpSocket() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTcpSocket::~QTcpSocket")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTcpSocket_DestroyQTcpSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
