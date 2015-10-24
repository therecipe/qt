package dbus

//#include "qdbuspendingcallwatcher.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusPendingCallWatcher struct {
	core.QObject
	QDBusPendingCall
}

type QDBusPendingCallWatcherITF interface {
	core.QObjectITF
	QDBusPendingCallITF
	QDBusPendingCallWatcherPTR() *QDBusPendingCallWatcher
}

func (p *QDBusPendingCallWatcher) Pointer() unsafe.Pointer {
	return p.QObjectPTR().Pointer()
}

func (p *QDBusPendingCallWatcher) SetPointer(ptr unsafe.Pointer) {
	p.QObjectPTR().SetPointer(ptr)
	p.QDBusPendingCallPTR().SetPointer(ptr)
}

func PointerFromQDBusPendingCallWatcher(ptr QDBusPendingCallWatcherITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCallWatcherPTR().Pointer()
	}
	return nil
}

func QDBusPendingCallWatcherFromPointer(ptr unsafe.Pointer) *QDBusPendingCallWatcher {
	var n = new(QDBusPendingCallWatcher)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDBusPendingCallWatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusPendingCallWatcher) QDBusPendingCallWatcherPTR() *QDBusPendingCallWatcher {
	return ptr
}

func (ptr *QDBusPendingCallWatcher) WaitForFinished() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_WaitForFinished(C.QtObjectPtr(ptr.Pointer()))
	}
}

func NewQDBusPendingCallWatcher(call QDBusPendingCallITF, parent core.QObjectITF) *QDBusPendingCallWatcher {
	return QDBusPendingCallWatcherFromPointer(unsafe.Pointer(C.QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(C.QtObjectPtr(PointerFromQDBusPendingCall(call)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QDBusPendingCallWatcher) ConnectFinished(f func(self QDBusPendingCallWatcherITF)) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDBusPendingCallWatcherFinished
func callbackQDBusPendingCallWatcherFinished(ptrName *C.char, self unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func(*QDBusPendingCallWatcher))(QDBusPendingCallWatcherFromPointer(self))
}

func (ptr *QDBusPendingCallWatcher) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_IsFinished(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) DestroyQDBusPendingCallWatcher() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
