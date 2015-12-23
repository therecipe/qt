package bluetooth

//#include "bluetooth.h"
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
	for len(n.ObjectName()) < len("QBluetoothTransferReply_") {
		n.SetObjectName("QBluetoothTransferReply_" + qt.Identifier())
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
	defer qt.Recovering("QBluetoothTransferReply::abort")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Abort(ptr.Pointer())
	}
}

func (ptr *QBluetoothTransferReply) ConnectError2(f func(errorType QBluetoothTransferReply__TransferError)) {
	defer qt.Recovering("connect QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQBluetoothTransferReplyError2
func callbackQBluetoothTransferReplyError2(ptrName *C.char, errorType C.int) {
	defer qt.Recovering("callback QBluetoothTransferReply::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothTransferReply__TransferError))(QBluetoothTransferReply__TransferError(errorType))
	}

}

func (ptr *QBluetoothTransferReply) Error() QBluetoothTransferReply__TransferError {
	defer qt.Recovering("QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		return QBluetoothTransferReply__TransferError(C.QBluetoothTransferReply_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothTransferReply) ErrorString() string {
	defer qt.Recovering("QBluetoothTransferReply::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothTransferReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothTransferReply) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	defer qt.Recovering("connect QBluetoothTransferReply::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothTransferReplyFinished
func callbackQBluetoothTransferReplyFinished(ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferReply) IsFinished() bool {
	defer qt.Recovering("QBluetoothTransferReply::isFinished")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) IsRunning() bool {
	defer qt.Recovering("QBluetoothTransferReply::isRunning")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) Manager() *QBluetoothTransferManager {
	defer qt.Recovering("QBluetoothTransferReply::manager")

	if ptr.Pointer() != nil {
		return NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferReply_Manager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothTransferReply) ConnectTransferProgress(f func(bytesTransferred int64, bytesTotal int64)) {
	defer qt.Recovering("connect QBluetoothTransferReply::transferProgress")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectTransferProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "transferProgress", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectTransferProgress() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::transferProgress")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectTransferProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "transferProgress")
	}
}

//export callbackQBluetoothTransferReplyTransferProgress
func callbackQBluetoothTransferReplyTransferProgress(ptrName *C.char, bytesTransferred C.longlong, bytesTotal C.longlong) {
	defer qt.Recovering("callback QBluetoothTransferReply::transferProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "transferProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesTransferred), int64(bytesTotal))
	}

}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReply() {
	defer qt.Recovering("QBluetoothTransferReply::~QBluetoothTransferReply")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DestroyQBluetoothTransferReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferReply) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothTransferReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothTransferReplyTimerEvent
func callbackQBluetoothTransferReplyTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBluetoothTransferReply::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QBluetoothTransferReply) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothTransferReply::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothTransferReplyChildEvent
func callbackQBluetoothTransferReplyChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBluetoothTransferReply::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QBluetoothTransferReply) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothTransferReply::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothTransferReplyCustomEvent
func callbackQBluetoothTransferReplyCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBluetoothTransferReply::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
