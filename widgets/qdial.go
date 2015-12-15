package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDial struct {
	QAbstractSlider
}

type QDial_ITF interface {
	QAbstractSlider_ITF
	QDial_PTR() *QDial
}

func PointerFromQDial(ptr QDial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDial_PTR().Pointer()
	}
	return nil
}

func NewQDialFromPointer(ptr unsafe.Pointer) *QDial {
	var n = new(QDial)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDial_") {
		n.SetObjectName("QDial_" + qt.Identifier())
	}
	return n
}

func (ptr *QDial) QDial_PTR() *QDial {
	return ptr
}

func (ptr *QDial) NotchSize() int {
	defer qt.Recovering("QDial::notchSize")

	if ptr.Pointer() != nil {
		return int(C.QDial_NotchSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDial) NotchTarget() float64 {
	defer qt.Recovering("QDial::notchTarget")

	if ptr.Pointer() != nil {
		return float64(C.QDial_NotchTarget(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDial) NotchesVisible() bool {
	defer qt.Recovering("QDial::notchesVisible")

	if ptr.Pointer() != nil {
		return C.QDial_NotchesVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDial) SetNotchesVisible(visible bool) {
	defer qt.Recovering("QDial::setNotchesVisible")

	if ptr.Pointer() != nil {
		C.QDial_SetNotchesVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDial) SetWrapping(on bool) {
	defer qt.Recovering("QDial::setWrapping")

	if ptr.Pointer() != nil {
		C.QDial_SetWrapping(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QDial) Wrapping() bool {
	defer qt.Recovering("QDial::wrapping")

	if ptr.Pointer() != nil {
		return C.QDial_Wrapping(ptr.Pointer()) != 0
	}
	return false
}

func NewQDial(parent QWidget_ITF) *QDial {
	defer qt.Recovering("QDial::QDial")

	return NewQDialFromPointer(C.QDial_NewQDial(PointerFromQWidget(parent)))
}

func (ptr *QDial) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDial::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDial) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDial::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDialMouseMoveEvent
func callbackQDialMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::mouseMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QDial) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDial::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDial) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDial::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDialMousePressEvent
func callbackQDialMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QDial) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDial::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDial) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDial::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDialMouseReleaseEvent
func callbackQDialMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QDial) ConnectPaintEvent(f func(pe *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDial::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDial) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDial::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDialPaintEvent
func callbackQDialPaintEvent(ptrName *C.char, pe unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(pe))
		return true
	}
	return false

}

func (ptr *QDial) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDial::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDial) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDial::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDialResizeEvent
func callbackQDialResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QDial) ConnectSliderChange(f func(change QAbstractSlider__SliderChange)) {
	defer qt.Recovering("connect QDial::sliderChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sliderChange", f)
	}
}

func (ptr *QDial) DisconnectSliderChange() {
	defer qt.Recovering("disconnect QDial::sliderChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sliderChange")
	}
}

//export callbackQDialSliderChange
func callbackQDialSliderChange(ptrName *C.char, change C.int) bool {
	defer qt.Recovering("callback QDial::sliderChange")

	var signal = qt.GetSignal(C.GoString(ptrName), "sliderChange")
	if signal != nil {
		defer signal.(func(QAbstractSlider__SliderChange))(QAbstractSlider__SliderChange(change))
		return true
	}
	return false

}

func (ptr *QDial) DestroyQDial() {
	defer qt.Recovering("QDial::~QDial")

	if ptr.Pointer() != nil {
		C.QDial_DestroyQDial(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
