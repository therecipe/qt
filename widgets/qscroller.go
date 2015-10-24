package widgets

//#include "qscroller.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScroller struct {
	core.QObject
}

type QScrollerITF interface {
	core.QObjectITF
	QScrollerPTR() *QScroller
}

func PointerFromQScroller(ptr QScrollerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollerPTR().Pointer()
	}
	return nil
}

func QScrollerFromPointer(ptr unsafe.Pointer) *QScroller {
	var n = new(QScroller)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QScroller_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScroller) QScrollerPTR() *QScroller {
	return ptr
}

//QScroller::Input
type QScroller__Input int

var (
	QScroller__InputPress   = QScroller__Input(1)
	QScroller__InputMove    = QScroller__Input(2)
	QScroller__InputRelease = QScroller__Input(3)
)

//QScroller::ScrollerGestureType
type QScroller__ScrollerGestureType int

var (
	QScroller__TouchGesture             = QScroller__ScrollerGestureType(0)
	QScroller__LeftMouseButtonGesture   = QScroller__ScrollerGestureType(1)
	QScroller__RightMouseButtonGesture  = QScroller__ScrollerGestureType(2)
	QScroller__MiddleMouseButtonGesture = QScroller__ScrollerGestureType(3)
)

//QScroller::State
type QScroller__State int

var (
	QScroller__Inactive  = QScroller__State(0)
	QScroller__Pressed   = QScroller__State(1)
	QScroller__Dragging  = QScroller__State(2)
	QScroller__Scrolling = QScroller__State(3)
)

func (ptr *QScroller) SetScrollerProperties(prop QScrollerPropertiesITF) {
	if ptr.Pointer() != nil {
		C.QScroller_SetScrollerProperties(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScrollerProperties(prop)))
	}
}

func QScroller_GrabGesture(target core.QObjectITF, scrollGestureType QScroller__ScrollerGestureType) core.Qt__GestureType {
	return core.Qt__GestureType(C.QScroller_QScroller_GrabGesture(C.QtObjectPtr(core.PointerFromQObject(target)), C.int(scrollGestureType)))
}

func QScroller_GrabbedGesture(target core.QObjectITF) core.Qt__GestureType {
	return core.Qt__GestureType(C.QScroller_QScroller_GrabbedGesture(C.QtObjectPtr(core.PointerFromQObject(target))))
}

func QScroller_HasScroller(target core.QObjectITF) bool {
	return C.QScroller_QScroller_HasScroller(C.QtObjectPtr(core.PointerFromQObject(target))) != 0
}

func (ptr *QScroller) ResendPrepareEvent() {
	if ptr.Pointer() != nil {
		C.QScroller_ResendPrepareEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QScroller) ScrollTo(pos core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QScroller_ScrollTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(pos)))
	}
}

func (ptr *QScroller) ScrollTo2(pos core.QPointFITF, scrollTime int) {
	if ptr.Pointer() != nil {
		C.QScroller_ScrollTo2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(pos)), C.int(scrollTime))
	}
}

func QScroller_Scroller(target core.QObjectITF) *QScroller {
	return QScrollerFromPointer(unsafe.Pointer(C.QScroller_QScroller_Scroller(C.QtObjectPtr(core.PointerFromQObject(target)))))
}

func QScroller_Scroller2(target core.QObjectITF) *QScroller {
	return QScrollerFromPointer(unsafe.Pointer(C.QScroller_QScroller_Scroller2(C.QtObjectPtr(core.PointerFromQObject(target)))))
}

func (ptr *QScroller) ConnectStateChanged(f func(newState QScroller__State)) {
	if ptr.Pointer() != nil {
		C.QScroller_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QScroller) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QScroller_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQScrollerStateChanged
func callbackQScrollerStateChanged(ptrName *C.char, newState C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QScroller__State))(QScroller__State(newState))
}

func (ptr *QScroller) Stop() {
	if ptr.Pointer() != nil {
		C.QScroller_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QScroller) Target() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QScroller_Target(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func QScroller_UngrabGesture(target core.QObjectITF) {
	C.QScroller_QScroller_UngrabGesture(C.QtObjectPtr(core.PointerFromQObject(target)))
}
