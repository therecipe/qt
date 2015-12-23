package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDrag struct {
	core.QObject
}

type QDrag_ITF interface {
	core.QObject_ITF
	QDrag_PTR() *QDrag
}

func PointerFromQDrag(ptr QDrag_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDrag_PTR().Pointer()
	}
	return nil
}

func NewQDragFromPointer(ptr unsafe.Pointer) *QDrag {
	var n = new(QDrag)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDrag_") {
		n.SetObjectName("QDrag_" + qt.Identifier())
	}
	return n
}

func (ptr *QDrag) QDrag_PTR() *QDrag {
	return ptr
}

func NewQDrag(dragSource core.QObject_ITF) *QDrag {
	defer qt.Recovering("QDrag::QDrag")

	return NewQDragFromPointer(C.QDrag_NewQDrag(core.PointerFromQObject(dragSource)))
}

func (ptr *QDrag) ConnectActionChanged(f func(action core.Qt__DropAction)) {
	defer qt.Recovering("connect QDrag::actionChanged")

	if ptr.Pointer() != nil {
		C.QDrag_ConnectActionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actionChanged", f)
	}
}

func (ptr *QDrag) DisconnectActionChanged() {
	defer qt.Recovering("disconnect QDrag::actionChanged")

	if ptr.Pointer() != nil {
		C.QDrag_DisconnectActionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actionChanged")
	}
}

//export callbackQDragActionChanged
func callbackQDragActionChanged(ptrName *C.char, action C.int) {
	defer qt.Recovering("callback QDrag::actionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionChanged"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(action))
	}

}

func (ptr *QDrag) Exec(supportedActions core.Qt__DropAction) core.Qt__DropAction {
	defer qt.Recovering("QDrag::exec")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDrag_Exec(ptr.Pointer(), C.int(supportedActions)))
	}
	return 0
}

func (ptr *QDrag) Exec2(supportedActions core.Qt__DropAction, defaultDropAction core.Qt__DropAction) core.Qt__DropAction {
	defer qt.Recovering("QDrag::exec")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDrag_Exec2(ptr.Pointer(), C.int(supportedActions), C.int(defaultDropAction)))
	}
	return 0
}

func (ptr *QDrag) HotSpot() *core.QPoint {
	defer qt.Recovering("QDrag::hotSpot")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QDrag_HotSpot(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDrag) MimeData() *core.QMimeData {
	defer qt.Recovering("QDrag::mimeData")

	if ptr.Pointer() != nil {
		return core.NewQMimeDataFromPointer(C.QDrag_MimeData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDrag) SetDragCursor(cursor QPixmap_ITF, action core.Qt__DropAction) {
	defer qt.Recovering("QDrag::setDragCursor")

	if ptr.Pointer() != nil {
		C.QDrag_SetDragCursor(ptr.Pointer(), PointerFromQPixmap(cursor), C.int(action))
	}
}

func (ptr *QDrag) SetHotSpot(hotspot core.QPoint_ITF) {
	defer qt.Recovering("QDrag::setHotSpot")

	if ptr.Pointer() != nil {
		C.QDrag_SetHotSpot(ptr.Pointer(), core.PointerFromQPoint(hotspot))
	}
}

func (ptr *QDrag) SetMimeData(data core.QMimeData_ITF) {
	defer qt.Recovering("QDrag::setMimeData")

	if ptr.Pointer() != nil {
		C.QDrag_SetMimeData(ptr.Pointer(), core.PointerFromQMimeData(data))
	}
}

func (ptr *QDrag) SetPixmap(pixmap QPixmap_ITF) {
	defer qt.Recovering("QDrag::setPixmap")

	if ptr.Pointer() != nil {
		C.QDrag_SetPixmap(ptr.Pointer(), PointerFromQPixmap(pixmap))
	}
}

func (ptr *QDrag) Source() *core.QObject {
	defer qt.Recovering("QDrag::source")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QDrag_Source(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDrag) SupportedActions() core.Qt__DropAction {
	defer qt.Recovering("QDrag::supportedActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDrag_SupportedActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDrag) Target() *core.QObject {
	defer qt.Recovering("QDrag::target")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QDrag_Target(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDrag) ConnectTargetChanged(f func(newTarget *core.QObject)) {
	defer qt.Recovering("connect QDrag::targetChanged")

	if ptr.Pointer() != nil {
		C.QDrag_ConnectTargetChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetChanged", f)
	}
}

func (ptr *QDrag) DisconnectTargetChanged() {
	defer qt.Recovering("disconnect QDrag::targetChanged")

	if ptr.Pointer() != nil {
		C.QDrag_DisconnectTargetChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetChanged")
	}
}

//export callbackQDragTargetChanged
func callbackQDragTargetChanged(ptrName *C.char, newTarget unsafe.Pointer) {
	defer qt.Recovering("callback QDrag::targetChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "targetChanged"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(newTarget))
	}

}

func (ptr *QDrag) DestroyQDrag() {
	defer qt.Recovering("QDrag::~QDrag")

	if ptr.Pointer() != nil {
		C.QDrag_DestroyQDrag(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDrag) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDrag::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDrag) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDrag::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDragTimerEvent
func callbackQDragTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDrag::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDrag) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDrag::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDrag) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDrag::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDragChildEvent
func callbackQDragChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDrag::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDrag) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDrag::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDrag) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDrag::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDragCustomEvent
func callbackQDragCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDrag::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
