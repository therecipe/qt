package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSlider struct {
	QAbstractSlider
}

type QSlider_ITF interface {
	QAbstractSlider_ITF
	QSlider_PTR() *QSlider
}

func PointerFromQSlider(ptr QSlider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSlider_PTR().Pointer()
	}
	return nil
}

func NewQSliderFromPointer(ptr unsafe.Pointer) *QSlider {
	var n = new(QSlider)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSlider_") {
		n.SetObjectName("QSlider_" + qt.Identifier())
	}
	return n
}

func (ptr *QSlider) QSlider_PTR() *QSlider {
	return ptr
}

//QSlider::TickPosition
type QSlider__TickPosition int64

const (
	QSlider__NoTicks        = QSlider__TickPosition(0)
	QSlider__TicksAbove     = QSlider__TickPosition(1)
	QSlider__TicksLeft      = QSlider__TickPosition(QSlider__TicksAbove)
	QSlider__TicksBelow     = QSlider__TickPosition(2)
	QSlider__TicksRight     = QSlider__TickPosition(QSlider__TicksBelow)
	QSlider__TicksBothSides = QSlider__TickPosition(3)
)

func (ptr *QSlider) SetTickInterval(ti int) {
	defer qt.Recovering("QSlider::setTickInterval")

	if ptr.Pointer() != nil {
		C.QSlider_SetTickInterval(ptr.Pointer(), C.int(ti))
	}
}

func (ptr *QSlider) SetTickPosition(position QSlider__TickPosition) {
	defer qt.Recovering("QSlider::setTickPosition")

	if ptr.Pointer() != nil {
		C.QSlider_SetTickPosition(ptr.Pointer(), C.int(position))
	}
}

func (ptr *QSlider) TickInterval() int {
	defer qt.Recovering("QSlider::tickInterval")

	if ptr.Pointer() != nil {
		return int(C.QSlider_TickInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSlider) TickPosition() QSlider__TickPosition {
	defer qt.Recovering("QSlider::tickPosition")

	if ptr.Pointer() != nil {
		return QSlider__TickPosition(C.QSlider_TickPosition(ptr.Pointer()))
	}
	return 0
}

func NewQSlider(parent QWidget_ITF) *QSlider {
	defer qt.Recovering("QSlider::QSlider")

	return NewQSliderFromPointer(C.QSlider_NewQSlider(PointerFromQWidget(parent)))
}

func NewQSlider2(orientation core.Qt__Orientation, parent QWidget_ITF) *QSlider {
	defer qt.Recovering("QSlider::QSlider")

	return NewQSliderFromPointer(C.QSlider_NewQSlider2(C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QSlider) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QSlider::event")

	if ptr.Pointer() != nil {
		return C.QSlider_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSlider) ConnectMouseMoveEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSlider::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSlider) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSlider::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSliderMouseMoveEvent
func callbackQSliderMouseMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::mouseMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QSlider) ConnectMousePressEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSlider::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSlider) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSlider::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSliderMousePressEvent
func callbackQSliderMousePressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QSlider) ConnectMouseReleaseEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSlider::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSlider) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSlider::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSliderMouseReleaseEvent
func callbackQSliderMouseReleaseEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QSlider) ConnectPaintEvent(f func(ev *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSlider::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSlider) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSlider::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSliderPaintEvent
func callbackQSliderPaintEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QSlider) DestroyQSlider() {
	defer qt.Recovering("QSlider::~QSlider")

	if ptr.Pointer() != nil {
		C.QSlider_DestroyQSlider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
