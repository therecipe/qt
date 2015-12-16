package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QScrollBar struct {
	QAbstractSlider
}

type QScrollBar_ITF interface {
	QAbstractSlider_ITF
	QScrollBar_PTR() *QScrollBar
}

func PointerFromQScrollBar(ptr QScrollBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollBar_PTR().Pointer()
	}
	return nil
}

func NewQScrollBarFromPointer(ptr unsafe.Pointer) *QScrollBar {
	var n = new(QScrollBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScrollBar_") {
		n.SetObjectName("QScrollBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QScrollBar) QScrollBar_PTR() *QScrollBar {
	return ptr
}

func NewQScrollBar(parent QWidget_ITF) *QScrollBar {
	defer qt.Recovering("QScrollBar::QScrollBar")

	return NewQScrollBarFromPointer(C.QScrollBar_NewQScrollBar(PointerFromQWidget(parent)))
}

func NewQScrollBar2(orientation core.Qt__Orientation, parent QWidget_ITF) *QScrollBar {
	defer qt.Recovering("QScrollBar::QScrollBar")

	return NewQScrollBarFromPointer(C.QScrollBar_NewQScrollBar2(C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QScrollBar) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QScrollBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QScrollBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQScrollBarContextMenuEvent
func callbackQScrollBarContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::contextMenuEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "contextMenuEvent")
	if signal != nil {
		defer signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QScrollBar::event")

	if ptr.Pointer() != nil {
		return C.QScrollBar_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScrollBar) ConnectHideEvent(f func(v *gui.QHideEvent)) {
	defer qt.Recovering("connect QScrollBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QScrollBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQScrollBarHideEvent
func callbackQScrollBarHideEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::hideEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "hideEvent")
	if signal != nil {
		defer signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QScrollBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQScrollBarMouseMoveEvent
func callbackQScrollBarMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::mouseMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QScrollBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQScrollBarMousePressEvent
func callbackQScrollBarMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QScrollBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQScrollBarMouseReleaseEvent
func callbackQScrollBarMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QScrollBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QScrollBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQScrollBarPaintEvent
func callbackQScrollBarPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QScrollBar) SizeHint() *core.QSize {
	defer qt.Recovering("QScrollBar::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QScrollBar_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollBar) ConnectSliderChange(f func(change QAbstractSlider__SliderChange)) {
	defer qt.Recovering("connect QScrollBar::sliderChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sliderChange", f)
	}
}

func (ptr *QScrollBar) DisconnectSliderChange() {
	defer qt.Recovering("disconnect QScrollBar::sliderChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sliderChange")
	}
}

//export callbackQScrollBarSliderChange
func callbackQScrollBarSliderChange(ptrName *C.char, change C.int) bool {
	defer qt.Recovering("callback QScrollBar::sliderChange")

	var signal = qt.GetSignal(C.GoString(ptrName), "sliderChange")
	if signal != nil {
		defer signal.(func(QAbstractSlider__SliderChange))(QAbstractSlider__SliderChange(change))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QScrollBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QScrollBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQScrollBarWheelEvent
func callbackQScrollBarWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::wheelEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "wheelEvent")
	if signal != nil {
		defer signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) DestroyQScrollBar() {
	defer qt.Recovering("QScrollBar::~QScrollBar")

	if ptr.Pointer() != nil {
		C.QScrollBar_DestroyQScrollBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
