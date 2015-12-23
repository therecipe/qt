package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQuickItemGrabResult struct {
	core.QObject
}

type QQuickItemGrabResult_ITF interface {
	core.QObject_ITF
	QQuickItemGrabResult_PTR() *QQuickItemGrabResult
}

func PointerFromQQuickItemGrabResult(ptr QQuickItemGrabResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItemGrabResult_PTR().Pointer()
	}
	return nil
}

func NewQQuickItemGrabResultFromPointer(ptr unsafe.Pointer) *QQuickItemGrabResult {
	var n = new(QQuickItemGrabResult)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickItemGrabResult_") {
		n.SetObjectName("QQuickItemGrabResult_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickItemGrabResult) QQuickItemGrabResult_PTR() *QQuickItemGrabResult {
	return ptr
}

func (ptr *QQuickItemGrabResult) Url() *core.QUrl {
	defer qt.Recovering("QQuickItemGrabResult::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQuickItemGrabResult_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItemGrabResult) ConnectReady(f func()) {
	defer qt.Recovering("connect QQuickItemGrabResult::ready")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "ready", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectReady() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::ready")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "ready")
	}
}

//export callbackQQuickItemGrabResultReady
func callbackQQuickItemGrabResultReady(ptrName *C.char) {
	defer qt.Recovering("callback QQuickItemGrabResult::ready")

	if signal := qt.GetSignal(C.GoString(ptrName), "ready"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickItemGrabResult) SaveToFile(fileName string) bool {
	defer qt.Recovering("QQuickItemGrabResult::saveToFile")

	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_SaveToFile(ptr.Pointer(), C.CString(fileName)) != 0
	}
	return false
}

func (ptr *QQuickItemGrabResult) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickItemGrabResult::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickItemGrabResultTimerEvent
func callbackQQuickItemGrabResultTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItemGrabResult::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItemGrabResult) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickItemGrabResult::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickItemGrabResultChildEvent
func callbackQQuickItemGrabResultChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItemGrabResult::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickItemGrabResult) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickItemGrabResult::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickItemGrabResultCustomEvent
func callbackQQuickItemGrabResultCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickItemGrabResult::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
