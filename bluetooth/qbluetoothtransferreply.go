package bluetooth

//#include "qbluetoothtransferreply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBluetoothTransferReply struct {
	core.QObject
}

type QBluetoothTransferReply_ITF interface {
	core.QObject_ITF
	QBluetoothTransferReply_PTR() *QBluetoothTransferReply
}

func PointerFromQBluetoothTransferReply(ptr QBluetoothTransferReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferReply_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferReplyFromPointer(ptr unsafe.Pointer) *QBluetoothTransferReply {
	var n = new(QBluetoothTransferReply)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QBluetoothTransferReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBluetoothTransferReply) QBluetoothTransferReply_PTR() *QBluetoothTransferReply {
	return ptr
}

//QBluetoothTransferReply::TransferError
type QBluetoothTransferReply__TransferError int64

const (
	QBluetoothTransferReply__NoError                   = QBluetoothTransferReply__TransferError(0)
	QBluetoothTransferReply__UnknownError              = QBluetoothTransferReply__TransferError(1)
	QBluetoothTransferReply__FileNotFoundError         = QBluetoothTransferReply__TransferError(2)
	QBluetoothTransferReply__HostNotFoundError         = QBluetoothTransferReply__TransferError(3)
	QBluetoothTransferReply__UserCanceledTransferError = QBluetoothTransferReply__TransferError(4)
	QBluetoothTransferReply__IODeviceNotReadableError  = QBluetoothTransferReply__TransferError(5)
	QBluetoothTransferReply__ResourceBusyError         = QBluetoothTransferReply__TransferError(6)
	QBluetoothTransferReply__SessionError              = QBluetoothTransferReply__TransferError(7)
)

func (ptr *QBluetoothTransferReply) Abort() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Abort(ptr.Pointer())
	}
}

func (ptr *QBluetoothTransferReply) Error() QBluetoothTransferReply__TransferError {
	if ptr.Pointer() != nil {
		return QBluetoothTransferReply__TransferError(C.QBluetoothTransferReply_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothTransferReply) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothTransferReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothTransferReply) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothTransferReplyFinished
func callbackQBluetoothTransferReplyFinished(ptrName *C.char, reply unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
}

func (ptr *QBluetoothTransferReply) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) Manager() *QBluetoothTransferManager {
	if ptr.Pointer() != nil {
		return NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferReply_Manager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReply() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DestroyQBluetoothTransferReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
