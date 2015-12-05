package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QBluetoothTransferManager struct {
	core.QObject
}

type QBluetoothTransferManager_ITF interface {
	core.QObject_ITF
	QBluetoothTransferManager_PTR() *QBluetoothTransferManager
}

func PointerFromQBluetoothTransferManager(ptr QBluetoothTransferManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferManager_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferManagerFromPointer(ptr unsafe.Pointer) *QBluetoothTransferManager {
	var n = new(QBluetoothTransferManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothTransferManager_") {
		n.SetObjectName("QBluetoothTransferManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBluetoothTransferManager) QBluetoothTransferManager_PTR() *QBluetoothTransferManager {
	return ptr
}

func (ptr *QBluetoothTransferManager) Put(request QBluetoothTransferRequest_ITF, data core.QIODevice_ITF) *QBluetoothTransferReply {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferManager::put")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBluetoothTransferReplyFromPointer(C.QBluetoothTransferManager_Put(ptr.Pointer(), PointerFromQBluetoothTransferRequest(request), core.PointerFromQIODevice(data)))
	}
	return nil
}

func NewQBluetoothTransferManager(parent core.QObject_ITF) *QBluetoothTransferManager {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferManager::QBluetoothTransferManager")
		}
	}()

	return NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferManager_NewQBluetoothTransferManager(core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothTransferManager) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferManager::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferManager::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothTransferManagerFinished
func callbackQBluetoothTransferManagerFinished(ptrName *C.char, reply unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferManager::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManager() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferManager::~QBluetoothTransferManager")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DestroyQBluetoothTransferManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
