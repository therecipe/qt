package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMacCocoaViewContainer struct {
	QWidget
}

type QMacCocoaViewContainer_ITF interface {
	QWidget_ITF
	QMacCocoaViewContainer_PTR() *QMacCocoaViewContainer
}

func PointerFromQMacCocoaViewContainer(ptr QMacCocoaViewContainer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacCocoaViewContainer_PTR().Pointer()
	}
	return nil
}

func NewQMacCocoaViewContainerFromPointer(ptr unsafe.Pointer) *QMacCocoaViewContainer {
	var n = new(QMacCocoaViewContainer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMacCocoaViewContainer_") {
		n.SetObjectName("QMacCocoaViewContainer_" + qt.Identifier())
	}
	return n
}

func (ptr *QMacCocoaViewContainer) QMacCocoaViewContainer_PTR() *QMacCocoaViewContainer {
	return ptr
}

func (ptr *QMacCocoaViewContainer) DestroyQMacCocoaViewContainer() {
	defer qt.Recovering("QMacCocoaViewContainer::~QMacCocoaViewContainer")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DestroyQMacCocoaViewContainer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacCocoaViewContainer) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQMacCocoaViewContainerActionEvent
func callbackQMacCocoaViewContainerActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::actionEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::actionEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQMacCocoaViewContainerDragEnterEvent
func callbackQMacCocoaViewContainerDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQMacCocoaViewContainerDragLeaveEvent
func callbackQMacCocoaViewContainerDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQMacCocoaViewContainerDragMoveEvent
func callbackQMacCocoaViewContainerDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQMacCocoaViewContainerDropEvent
func callbackQMacCocoaViewContainerDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::dropEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::dropEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQMacCocoaViewContainerEnterEvent
func callbackQMacCocoaViewContainerEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::enterEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::enterEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQMacCocoaViewContainerFocusInEvent
func callbackQMacCocoaViewContainerFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::focusInEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::focusInEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQMacCocoaViewContainerFocusOutEvent
func callbackQMacCocoaViewContainerFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQMacCocoaViewContainerHideEvent
func callbackQMacCocoaViewContainerHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::hideEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::hideEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQMacCocoaViewContainerLeaveEvent
func callbackQMacCocoaViewContainerLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::leaveEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::leaveEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQMacCocoaViewContainerMoveEvent
func callbackQMacCocoaViewContainerMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::moveEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::moveEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQMacCocoaViewContainerPaintEvent
func callbackQMacCocoaViewContainerPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::paintEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::paintEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQMacCocoaViewContainerSetVisible
func callbackQMacCocoaViewContainerSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) SetVisible(visible bool) {
	defer qt.Recovering("QMacCocoaViewContainer::setVisible")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMacCocoaViewContainer) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QMacCocoaViewContainer::setVisible")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQMacCocoaViewContainerShowEvent
func callbackQMacCocoaViewContainerShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::showEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::showEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQMacCocoaViewContainerChangeEvent
func callbackQMacCocoaViewContainerChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::changeEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::changeEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQMacCocoaViewContainerCloseEvent
func callbackQMacCocoaViewContainerCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::closeEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::closeEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQMacCocoaViewContainerContextMenuEvent
func callbackQMacCocoaViewContainerContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQMacCocoaViewContainerInitPainter
func callbackQMacCocoaViewContainerInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QMacCocoaViewContainer) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::initPainter")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QMacCocoaViewContainer) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::initPainter")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQMacCocoaViewContainerInputMethodEvent
func callbackQMacCocoaViewContainerInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQMacCocoaViewContainerKeyPressEvent
func callbackQMacCocoaViewContainerKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQMacCocoaViewContainerKeyReleaseEvent
func callbackQMacCocoaViewContainerKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQMacCocoaViewContainerMouseDoubleClickEvent
func callbackQMacCocoaViewContainerMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQMacCocoaViewContainerMouseMoveEvent
func callbackQMacCocoaViewContainerMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQMacCocoaViewContainerMousePressEvent
func callbackQMacCocoaViewContainerMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQMacCocoaViewContainerMouseReleaseEvent
func callbackQMacCocoaViewContainerMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQMacCocoaViewContainerResizeEvent
func callbackQMacCocoaViewContainerResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::resizeEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::resizeEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQMacCocoaViewContainerTabletEvent
func callbackQMacCocoaViewContainerTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::tabletEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::tabletEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQMacCocoaViewContainerWheelEvent
func callbackQMacCocoaViewContainerWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::wheelEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::wheelEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMacCocoaViewContainerTimerEvent
func callbackQMacCocoaViewContainerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMacCocoaViewContainerChildEvent
func callbackQMacCocoaViewContainerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::childEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::childEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMacCocoaViewContainerCustomEvent
func callbackQMacCocoaViewContainerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacCocoaViewContainer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacCocoaViewContainerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacCocoaViewContainer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::customEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacCocoaViewContainer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacCocoaViewContainer::customEvent")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
