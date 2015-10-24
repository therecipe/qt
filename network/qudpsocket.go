package network

//#include "qudpsocket.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QUdpSocket struct {
	QAbstractSocket
}

type QUdpSocketITF interface {
	QAbstractSocketITF
	QUdpSocketPTR() *QUdpSocket
}

func PointerFromQUdpSocket(ptr QUdpSocketITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUdpSocketPTR().Pointer()
	}
	return nil
}

func QUdpSocketFromPointer(ptr unsafe.Pointer) *QUdpSocket {
	var n = new(QUdpSocket)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QUdpSocket_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QUdpSocket) QUdpSocketPTR() *QUdpSocket {
	return ptr
}

func NewQUdpSocket(parent core.QObjectITF) *QUdpSocket {
	return QUdpSocketFromPointer(unsafe.Pointer(C.QUdpSocket_NewQUdpSocket(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QUdpSocket) HasPendingDatagrams() bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_HasPendingDatagrams(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup(groupAddress QHostAddressITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHostAddress(groupAddress))) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup2(groupAddress QHostAddressITF, iface QNetworkInterfaceITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHostAddress(groupAddress)), C.QtObjectPtr(PointerFromQNetworkInterface(iface))) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup(groupAddress QHostAddressITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHostAddress(groupAddress))) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup2(groupAddress QHostAddressITF, iface QNetworkInterfaceITF) bool {
	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHostAddress(groupAddress)), C.QtObjectPtr(PointerFromQNetworkInterface(iface))) != 0
	}
	return false
}

func (ptr *QUdpSocket) SetMulticastInterface(iface QNetworkInterfaceITF) {
	if ptr.Pointer() != nil {
		C.QUdpSocket_SetMulticastInterface(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkInterface(iface)))
	}
}

func (ptr *QUdpSocket) DestroyQUdpSocket() {
	if ptr.Pointer() != nil {
		C.QUdpSocket_DestroyQUdpSocket(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
