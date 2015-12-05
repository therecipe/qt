package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QDBusPendingCallWatcher struct {
	core.QObject
	QDBusPendingCall
}

type QDBusPendingCallWatcher_ITF interface {
	core.QObject_ITF
	QDBusPendingCall_ITF
	QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher
}

func (p *QDBusPendingCallWatcher) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QDBusPendingCallWatcher) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QDBusPendingCall_PTR().SetPointer(ptr)
}

func PointerFromQDBusPendingCallWatcher(ptr QDBusPendingCallWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCallWatcher_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingCallWatcherFromPointer(ptr unsafe.Pointer) *QDBusPendingCallWatcher {
	var n = new(QDBusPendingCallWatcher)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusPendingCallWatcher_") {
		n.SetObjectName("QDBusPendingCallWatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusPendingCallWatcher) QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher {
	return ptr
}

func (ptr *QDBusPendingCallWatcher) WaitForFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusPendingCallWatcher::waitForFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_WaitForFinished(ptr.Pointer())
	}
}

func NewQDBusPendingCallWatcher(call QDBusPendingCall_ITF, parent core.QObject_ITF) *QDBusPendingCallWatcher {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusPendingCallWatcher::QDBusPendingCallWatcher")
		}
	}()

	return NewQDBusPendingCallWatcherFromPointer(C.QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(PointerFromQDBusPendingCall(call), core.PointerFromQObject(parent)))
}

func (ptr *QDBusPendingCallWatcher) ConnectFinished(f func(self *QDBusPendingCallWatcher)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusPendingCallWatcher::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusPendingCallWatcher::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDBusPendingCallWatcherFinished
func callbackQDBusPendingCallWatcherFinished(ptrName *C.char, self unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusPendingCallWatcher::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func(*QDBusPendingCallWatcher))(NewQDBusPendingCallWatcherFromPointer(self))
}

func (ptr *QDBusPendingCallWatcher) IsFinished() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusPendingCallWatcher::isFinished")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) DestroyQDBusPendingCallWatcher() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusPendingCallWatcher::~QDBusPendingCallWatcher")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
