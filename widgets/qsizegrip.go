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
func callbackQSizeGripMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSizeGrip::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSizeGrip) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSizeGrip::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQSizeGripMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSizeGrip::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSizeGrip) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSizeGrip::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func NewQSizeGrip(parent QWidget_ITF) *QSizeGrip {
	defer qt.Recovering("QSizeGrip::QSizeGrip")

	return NewQSizeGripFromPointer(C.QSizeGrip_NewQSizeGrip(PointerFromQWidget(parent)))
}

func (ptr *QSizeGrip) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QSizeGrip::event")

	if ptr.Pointer() != nil {
		return C.QSizeGrip_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSizeGrip) EventFilter(o core.QObject_ITF, e core.QEvent_ITF) bool {
	defer qt.Recovering("QSizeGrip::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSizeGrip_EventFilter(ptr.Pointer(), core.PointerFromQObject(o), core.PointerFromQEvent(e)) != 0
	}
	return false
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
func callbackQSizeGripHideEvent(ptr unsafe.Pointer, ptrName *C.char, hideEvent unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(hideEvent))
	} else {
		NewQSizeGripFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(hideEvent))
	}
}

func (ptr *QSizeGrip) HideEvent(hideEvent gui.QHideEvent_ITF) {
	defer qt.Recovering("QSizeGrip::hideEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(hideEvent))
	}
}

func (ptr *QSizeGrip) HideEventDefault(hideEvent gui.QHideEvent_ITF) {
	defer qt.Recovering("QSizeGrip::hideEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(hideEvent))
	}
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
func callbackQSizeGripMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, mouseEvent unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(mouseEvent))
	} else {
		NewQSizeGripFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(mouseEvent))
	}
}

func (ptr *QSizeGrip) MouseReleaseEvent(mouseEvent gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSizeGrip::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(mouseEvent))
	}
}

func (ptr *QSizeGrip) MouseReleaseEventDefault(mouseEvent gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSizeGrip::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(mouseEvent))
	}
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
func callbackQSizeGripMoveEvent(ptr unsafe.Pointer, ptrName *C.char, moveEvent unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(moveEvent))
	} else {
		NewQSizeGripFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(moveEvent))
	}
}

func (ptr *QSizeGrip) MoveEvent(moveEvent gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSizeGrip::moveEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(moveEvent))
	}
}

func (ptr *QSizeGrip) MoveEventDefault(moveEvent gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSizeGrip::moveEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(moveEvent))
	}
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
func callbackQSizeGripPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSizeGrip::paintEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QSizeGrip) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSizeGrip::paintEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
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
func callbackQSizeGripSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSizeGrip::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QSizeGrip) SetVisible(visible bool) {
	defer qt.Recovering("QSizeGrip::setVisible")

	if ptr.Pointer() != nil {
		C.QSizeGrip_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSizeGrip) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QSizeGrip::setVisible")

	if ptr.Pointer() != nil {
		C.QSizeGrip_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQSizeGripShowEvent(ptr unsafe.Pointer, ptrName *C.char, showEvent unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(showEvent))
	} else {
		NewQSizeGripFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(showEvent))
	}
}

func (ptr *QSizeGrip) ShowEvent(showEvent gui.QShowEvent_ITF) {
	defer qt.Recovering("QSizeGrip::showEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(showEvent))
	}
}

func (ptr *QSizeGrip) ShowEventDefault(showEvent gui.QShowEvent_ITF) {
	defer qt.Recovering("QSizeGrip::showEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(showEvent))
	}
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

func (ptr *QSizeGrip) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QSizeGrip::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QSizeGrip::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQSizeGripActionEvent
func callbackQSizeGripActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSizeGrip::actionEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSizeGrip) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSizeGrip::actionEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QSizeGrip::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QSizeGrip::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQSizeGripDragEnterEvent
func callbackQSizeGripDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSizeGrip::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSizeGrip) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSizeGrip::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QSizeGrip::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QSizeGrip::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQSizeGripDragLeaveEvent
func callbackQSizeGripDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSizeGrip::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSizeGrip) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSizeGrip::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QSizeGrip::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QSizeGrip::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQSizeGripDragMoveEvent
func callbackQSizeGripDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSizeGrip::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSizeGrip) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSizeGrip::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QSizeGrip::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QSizeGrip::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQSizeGripDropEvent
func callbackQSizeGripDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSizeGrip::dropEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSizeGrip) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSizeGrip::dropEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSizeGrip::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QSizeGrip::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQSizeGripEnterEvent
func callbackQSizeGripEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSizeGrip::enterEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSizeGrip) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSizeGrip::enterEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSizeGrip::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QSizeGrip::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQSizeGripFocusInEvent
func callbackQSizeGripFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSizeGrip::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSizeGrip) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSizeGrip::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSizeGrip::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QSizeGrip::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQSizeGripFocusOutEvent
func callbackQSizeGripFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSizeGrip::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSizeGrip) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSizeGrip::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSizeGrip::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QSizeGrip::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQSizeGripLeaveEvent
func callbackQSizeGripLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSizeGrip::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSizeGrip) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSizeGrip::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSizeGrip::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSizeGrip::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQSizeGripChangeEvent
func callbackQSizeGripChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSizeGrip::changeEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSizeGrip) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSizeGrip::changeEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QSizeGrip::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QSizeGrip::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQSizeGripCloseEvent
func callbackQSizeGripCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSizeGrip::closeEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSizeGrip) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSizeGrip::closeEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QSizeGrip::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QSizeGrip::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQSizeGripContextMenuEvent
func callbackQSizeGripContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSizeGrip::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSizeGrip) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSizeGrip::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSizeGrip::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QSizeGrip) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QSizeGrip::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQSizeGripInitPainter
func callbackQSizeGripInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQSizeGripFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QSizeGrip) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSizeGrip::initPainter")

	if ptr.Pointer() != nil {
		C.QSizeGrip_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSizeGrip) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSizeGrip::initPainter")

	if ptr.Pointer() != nil {
		C.QSizeGrip_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSizeGrip) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QSizeGrip::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QSizeGrip::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQSizeGripInputMethodEvent
func callbackQSizeGripInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSizeGrip::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSizeGrip) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSizeGrip::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSizeGrip::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QSizeGrip::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQSizeGripKeyPressEvent
func callbackQSizeGripKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSizeGrip::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSizeGrip) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSizeGrip::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSizeGrip::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QSizeGrip::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQSizeGripKeyReleaseEvent
func callbackQSizeGripKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSizeGrip::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSizeGrip) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSizeGrip::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSizeGrip::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QSizeGrip::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQSizeGripMouseDoubleClickEvent
func callbackQSizeGripMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSizeGrip::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSizeGrip) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSizeGrip::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSizeGrip::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSizeGrip::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSizeGripResizeEvent
func callbackQSizeGripResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSizeGrip::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSizeGrip) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSizeGrip::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QSizeGrip::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QSizeGrip::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQSizeGripTabletEvent
func callbackQSizeGripTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSizeGrip::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSizeGrip) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSizeGrip::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QSizeGrip::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QSizeGrip::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQSizeGripWheelEvent
func callbackQSizeGripWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSizeGrip::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSizeGrip) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSizeGrip::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSizeGrip::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSizeGrip::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSizeGripTimerEvent
func callbackQSizeGripTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSizeGrip::timerEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSizeGrip) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSizeGrip::timerEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSizeGrip::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSizeGrip::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSizeGripChildEvent
func callbackQSizeGripChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSizeGrip::childEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSizeGrip) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSizeGrip::childEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSizeGrip) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSizeGrip::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSizeGrip::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSizeGripCustomEvent
func callbackQSizeGripCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSizeGrip::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSizeGripFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSizeGrip) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSizeGrip::customEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSizeGrip) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSizeGrip::customEvent")

	if ptr.Pointer() != nil {
		C.QSizeGrip_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
