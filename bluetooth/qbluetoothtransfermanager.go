package bluetooth

//#include "qbluetoothtransfermanager.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBluetoothTransferManager struct {
	core.QObject
}

type QBluetoothTransferManagerITF interface {
	core.QObjectITF
	QBluetoothTransferManagerPTR() *QBluetoothTransferManager
}

func PointerFromQBluetoothTransferManager(ptr QBluetoothTransferManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferManagerPTR().Pointer()
	}
	return nil
}

func QBluetoothTransferManagerFromPointer(ptr unsafe.Pointer) *QBluetoothTransferManager {
	var n = new(QBluetoothTransferManager)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QBluetoothTransferManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBluetoothTransferManager) QBluetoothTransferManagerPTR() *QBluetoothTransferManager {
	return ptr
}

func (ptr *QBluetoothTransferManager) Put(request QBluetoothTransferRequestITF, data core.QIODeviceITF) *QBluetoothTransferReply {
	if ptr.Pointer() != nil {
		return QBluetoothTransferReplyFromPointer(unsafe.Pointer(C.QBluetoothTransferManager_Put(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBluetoothTransferRequest(request)), C.QtObjectPtr(core.PointerFromQIODevice(data)))))
	}
	return nil
}

func NewQBluetoothTransferManager(parent core.QObjectITF) *QBluetoothTransferManager {
	return QBluetoothTransferManagerFromPointer(unsafe.Pointer(C.QBluetoothTransferManager_NewQBluetoothTransferManager(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QBluetoothTransferManager) ConnectFinished(f func(reply QBluetoothTransferReplyITF)) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothTransferManagerFinished
func callbackQBluetoothTransferManagerFinished(ptrName *C.char, reply unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func(*QBluetoothTransferReply))(QBluetoothTransferReplyFromPointer(reply))
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManager() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DestroyQBluetoothTransferManager(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
