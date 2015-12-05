package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QScroller struct {
	core.QObject
}

type QScroller_ITF interface {
	core.QObject_ITF
	QScroller_PTR() *QScroller
}

func PointerFromQScroller(ptr QScroller_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScroller_PTR().Pointer()
	}
	return nil
}

func NewQScrollerFromPointer(ptr unsafe.Pointer) *QScroller {
	var n = new(QScroller)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScroller_") {
		n.SetObjectName("QScroller_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScroller) QScroller_PTR() *QScroller {
	return ptr
}

//QScroller::Input
type QScroller__Input int64

const (
	QScroller__InputPress   = QScroller__Input(1)
	QScroller__InputMove    = QScroller__Input(2)
	QScroller__InputRelease = QScroller__Input(3)
)

//QScroller::ScrollerGestureType
type QScroller__ScrollerGestureType int64

const (
	QScroller__TouchGesture             = QScroller__ScrollerGestureType(0)
	QScroller__LeftMouseButtonGesture   = QScroller__ScrollerGestureType(1)
	QScroller__RightMouseButtonGesture  = QScroller__ScrollerGestureType(2)
	QScroller__MiddleMouseButtonGesture = QScroller__ScrollerGestureType(3)
)

//QScroller::State
type QScroller__State int64

const (
	QScroller__Inactive  = QScroller__State(0)
	QScroller__Pressed   = QScroller__State(1)
	QScroller__Dragging  = QScroller__State(2)
	QScroller__Scrolling = QScroller__State(3)
)

func (ptr *QScroller) SetScrollerProperties(prop QScrollerProperties_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::setScrollerProperties")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_SetScrollerProperties(ptr.Pointer(), PointerFromQScrollerProperties(prop))
	}
}

func (ptr *QScroller) EnsureVisible(rect core.QRectF_ITF, xmargin float64, ymargin float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::ensureVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_EnsureVisible(ptr.Pointer(), core.PointerFromQRectF(rect), C.double(xmargin), C.double(ymargin))
	}
}

func (ptr *QScroller) EnsureVisible2(rect core.QRectF_ITF, xmargin float64, ymargin float64, scrollTime int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::ensureVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_EnsureVisible2(ptr.Pointer(), core.PointerFromQRectF(rect), C.double(xmargin), C.double(ymargin), C.int(scrollTime))
	}
}

func QScroller_GrabGesture(target core.QObject_ITF, scrollGestureType QScroller__ScrollerGestureType) core.Qt__GestureType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::grabGesture")
		}
	}()

	return core.Qt__GestureType(C.QScroller_QScroller_GrabGesture(core.PointerFromQObject(target), C.int(scrollGestureType)))
}

func QScroller_GrabbedGesture(target core.QObject_ITF) core.Qt__GestureType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::grabbedGesture")
		}
	}()

	return core.Qt__GestureType(C.QScroller_QScroller_GrabbedGesture(core.PointerFromQObject(target)))
}

func QScroller_HasScroller(target core.QObject_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::hasScroller")
		}
	}()

	return C.QScroller_QScroller_HasScroller(core.PointerFromQObject(target)) != 0
}

func (ptr *QScroller) ResendPrepareEvent() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::resendPrepareEvent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_ResendPrepareEvent(ptr.Pointer())
	}
}

func (ptr *QScroller) ScrollTo(pos core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::scrollTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_ScrollTo(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func (ptr *QScroller) ScrollTo2(pos core.QPointF_ITF, scrollTime int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::scrollTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_ScrollTo2(ptr.Pointer(), core.PointerFromQPointF(pos), C.int(scrollTime))
	}
}

func QScroller_Scroller(target core.QObject_ITF) *QScroller {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::scroller")
		}
	}()

	return NewQScrollerFromPointer(C.QScroller_QScroller_Scroller(core.PointerFromQObject(target)))
}

func QScroller_Scroller2(target core.QObject_ITF) *QScroller {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::scroller")
		}
	}()

	return NewQScrollerFromPointer(C.QScroller_QScroller_Scroller2(core.PointerFromQObject(target)))
}

func (ptr *QScroller) SetSnapPositionsX2(first float64, interval float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::setSnapPositionsX")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_SetSnapPositionsX2(ptr.Pointer(), C.double(first), C.double(interval))
	}
}

func (ptr *QScroller) SetSnapPositionsY2(first float64, interval float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::setSnapPositionsY")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_SetSnapPositionsY2(ptr.Pointer(), C.double(first), C.double(interval))
	}
}

func (ptr *QScroller) ConnectStateChanged(f func(newState QScroller__State)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QScroller) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQScrollerStateChanged
func callbackQScrollerStateChanged(ptrName *C.char, newState C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QScroller__State))(QScroller__State(newState))
}

func (ptr *QScroller) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScroller_Stop(ptr.Pointer())
	}
}

func (ptr *QScroller) Target() *core.QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::target")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QScroller_Target(ptr.Pointer()))
	}
	return nil
}

func QScroller_UngrabGesture(target core.QObject_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScroller::ungrabGesture")
		}
	}()

	C.QScroller_QScroller_UngrabGesture(core.PointerFromQObject(target))
}
