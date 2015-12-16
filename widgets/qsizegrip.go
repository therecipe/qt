package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSizeGrip struct {
	QWidget
}

type QSizeGrip_ITF interface {
	QWidget_ITF
	QSizeGrip_PTR() *QSizeGrip
}

func PointerFromQSizeGrip(ptr QSizeGrip_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSizeGrip_PTR().Pointer()
	}
	return nil
}

func NewQSizeGripFromPointer(ptr unsafe.Pointer) *QSizeGrip {
	var n = new(QSizeGrip)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSizeGrip_") {
		n.SetObjectName("QSizeGrip_" + qt.Identifier())
	}
	return n
}

func (ptr *QSizeGrip) QSizeGrip_PTR() *QSizeGrip {
	return ptr
}

func (ptr *QSizeGrip) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSizeGrip::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSizeGrip::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSizeGripMouseMoveEvent
func callbackQSizeGripMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::mouseMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSizeGrip::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSizeGrip::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSizeGripMousePressEvent
func callbackQSizeGripMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func NewQSizeGrip(parent QWidget_ITF) *QSizeGrip {
	defer qt.Recovering("QSizeGrip::QSizeGrip")

	return NewQSizeGripFromPointer(C.QSizeGrip_NewQSizeGrip(PointerFromQWidget(parent)))
}

func (ptr *QSizeGrip) ConnectHideEvent(f func(hideEvent *gui.QHideEvent)) {
	defer qt.Recovering("connect QSizeGrip::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QSizeGrip::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQSizeGripHideEvent
func callbackQSizeGripHideEvent(ptrName *C.char, hideEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::hideEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "hideEvent")
	if signal != nil {
		defer signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(hideEvent))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectMouseReleaseEvent(f func(mouseEvent *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSizeGrip::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSizeGrip::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSizeGripMouseReleaseEvent
func callbackQSizeGripMouseReleaseEvent(ptrName *C.char, mouseEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(mouseEvent))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectMoveEvent(f func(moveEvent *gui.QMoveEvent)) {
	defer qt.Recovering("connect QSizeGrip::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QSizeGrip::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQSizeGripMoveEvent
func callbackQSizeGripMoveEvent(ptrName *C.char, moveEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::moveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "moveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(moveEvent))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSizeGrip::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSizeGrip::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSizeGripPaintEvent
func callbackQSizeGripPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QSizeGrip::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QSizeGrip) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QSizeGrip::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQSizeGripSetVisible
func callbackQSizeGripSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSizeGrip::setVisible")

	var signal = qt.GetSignal(C.GoString(ptrName), "setVisible")
	if signal != nil {
		defer signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectShowEvent(f func(showEvent *gui.QShowEvent)) {
	defer qt.Recovering("connect QSizeGrip::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QSizeGrip::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQSizeGripShowEvent
func callbackQSizeGripShowEvent(ptrName *C.char, showEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::showEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "showEvent")
	if signal != nil {
		defer signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(showEvent))
		return true
	}
	return false

}

func (ptr *QSizeGrip) SizeHint() *core.QSize {
	defer qt.Recovering("QSizeGrip::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSizeGrip_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSizeGrip) DestroyQSizeGrip() {
	defer qt.Recovering("QSizeGrip::~QSizeGrip")

	if ptr.Pointer() != nil {
		C.QSizeGrip_DestroyQSizeGrip(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
