package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDesktopWidget struct {
	QWidget
}

type QDesktopWidget_ITF interface {
	QWidget_ITF
	QDesktopWidget_PTR() *QDesktopWidget
}

func PointerFromQDesktopWidget(ptr QDesktopWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesktopWidget_PTR().Pointer()
	}
	return nil
}

func NewQDesktopWidgetFromPointer(ptr unsafe.Pointer) *QDesktopWidget {
	var n = new(QDesktopWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDesktopWidget_") {
		n.SetObjectName("QDesktopWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QDesktopWidget) QDesktopWidget_PTR() *QDesktopWidget {
	return ptr
}

func (ptr *QDesktopWidget) AvailableGeometry2(widget QWidget_ITF) *core.QRect {
	defer qt.Recovering("QDesktopWidget::availableGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_AvailableGeometry2(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QDesktopWidget) AvailableGeometry(screen int) *core.QRect {
	defer qt.Recovering("QDesktopWidget::availableGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_AvailableGeometry(ptr.Pointer(), C.int(screen)))
	}
	return nil
}

func (ptr *QDesktopWidget) IsVirtualDesktop() bool {
	defer qt.Recovering("QDesktopWidget::isVirtualDesktop")

	if ptr.Pointer() != nil {
		return C.QDesktopWidget_IsVirtualDesktop(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesktopWidget) PrimaryScreen() int {
	defer qt.Recovering("QDesktopWidget::primaryScreen")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_PrimaryScreen(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDesktopWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDesktopWidgetResizeEvent
func callbackQDesktopWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDesktopWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDesktopWidget) Screen(screen int) *QWidget {
	defer qt.Recovering("QDesktopWidget::screen")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDesktopWidget_Screen(ptr.Pointer(), C.int(screen)))
	}
	return nil
}

func (ptr *QDesktopWidget) ScreenGeometry2(widget QWidget_ITF) *core.QRect {
	defer qt.Recovering("QDesktopWidget::screenGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_ScreenGeometry2(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QDesktopWidget) ScreenGeometry(screen int) *core.QRect {
	defer qt.Recovering("QDesktopWidget::screenGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_ScreenGeometry(ptr.Pointer(), C.int(screen)))
	}
	return nil
}

func (ptr *QDesktopWidget) ScreenNumber2(point core.QPoint_ITF) int {
	defer qt.Recovering("QDesktopWidget::screenNumber")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber2(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return 0
}

func (ptr *QDesktopWidget) ScreenNumber(widget QWidget_ITF) int {
	defer qt.Recovering("QDesktopWidget::screenNumber")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QDesktopWidget) AvailableGeometry3(p core.QPoint_ITF) *core.QRect {
	defer qt.Recovering("QDesktopWidget::availableGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_AvailableGeometry3(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QDesktopWidget) ConnectResized(f func(screen int)) {
	defer qt.Recovering("connect QDesktopWidget::resized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "resized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectResized() {
	defer qt.Recovering("disconnect QDesktopWidget::resized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "resized")
	}
}

//export callbackQDesktopWidgetResized
func callbackQDesktopWidgetResized(ptr unsafe.Pointer, ptrName *C.char, screen C.int) {
	defer qt.Recovering("callback QDesktopWidget::resized")

	if signal := qt.GetSignal(C.GoString(ptrName), "resized"); signal != nil {
		signal.(func(int))(int(screen))
	}

}

func (ptr *QDesktopWidget) Resized(screen int) {
	defer qt.Recovering("QDesktopWidget::resized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_Resized(ptr.Pointer(), C.int(screen))
	}
}

func (ptr *QDesktopWidget) ScreenCount() int {
	defer qt.Recovering("QDesktopWidget::screenCount")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectScreenCountChanged(f func(newCount int)) {
	defer qt.Recovering("connect QDesktopWidget::screenCountChanged")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectScreenCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenCountChanged", f)
	}
}

func (ptr *QDesktopWidget) DisconnectScreenCountChanged() {
	defer qt.Recovering("disconnect QDesktopWidget::screenCountChanged")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectScreenCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenCountChanged")
	}
}

//export callbackQDesktopWidgetScreenCountChanged
func callbackQDesktopWidgetScreenCountChanged(ptr unsafe.Pointer, ptrName *C.char, newCount C.int) {
	defer qt.Recovering("callback QDesktopWidget::screenCountChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "screenCountChanged"); signal != nil {
		signal.(func(int))(int(newCount))
	}

}

func (ptr *QDesktopWidget) ScreenCountChanged(newCount int) {
	defer qt.Recovering("QDesktopWidget::screenCountChanged")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ScreenCountChanged(ptr.Pointer(), C.int(newCount))
	}
}

func (ptr *QDesktopWidget) ScreenGeometry3(p core.QPoint_ITF) *core.QRect {
	defer qt.Recovering("QDesktopWidget::screenGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_ScreenGeometry3(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QDesktopWidget) ConnectWorkAreaResized(f func(screen int)) {
	defer qt.Recovering("connect QDesktopWidget::workAreaResized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectWorkAreaResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "workAreaResized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectWorkAreaResized() {
	defer qt.Recovering("disconnect QDesktopWidget::workAreaResized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectWorkAreaResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "workAreaResized")
	}
}

//export callbackQDesktopWidgetWorkAreaResized
func callbackQDesktopWidgetWorkAreaResized(ptr unsafe.Pointer, ptrName *C.char, screen C.int) {
	defer qt.Recovering("callback QDesktopWidget::workAreaResized")

	if signal := qt.GetSignal(C.GoString(ptrName), "workAreaResized"); signal != nil {
		signal.(func(int))(int(screen))
	}

}

func (ptr *QDesktopWidget) WorkAreaResized(screen int) {
	defer qt.Recovering("QDesktopWidget::workAreaResized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_WorkAreaResized(ptr.Pointer(), C.int(screen))
	}
}

func (ptr *QDesktopWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDesktopWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQDesktopWidgetActionEvent
func callbackQDesktopWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDesktopWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDesktopWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQDesktopWidgetDragEnterEvent
func callbackQDesktopWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDesktopWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDesktopWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQDesktopWidgetDragLeaveEvent
func callbackQDesktopWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDesktopWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDesktopWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQDesktopWidgetDragMoveEvent
func callbackQDesktopWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDesktopWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDesktopWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQDesktopWidgetDropEvent
func callbackQDesktopWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDesktopWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesktopWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQDesktopWidgetEnterEvent
func callbackQDesktopWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesktopWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesktopWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQDesktopWidgetFocusInEvent
func callbackQDesktopWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesktopWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesktopWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQDesktopWidgetFocusOutEvent
func callbackQDesktopWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesktopWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDesktopWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQDesktopWidgetHideEvent
func callbackQDesktopWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDesktopWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesktopWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQDesktopWidgetLeaveEvent
func callbackQDesktopWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesktopWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDesktopWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQDesktopWidgetMoveEvent
func callbackQDesktopWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDesktopWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDesktopWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDesktopWidgetPaintEvent
func callbackQDesktopWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDesktopWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDesktopWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QDesktopWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDesktopWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQDesktopWidgetSetVisible
func callbackQDesktopWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDesktopWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDesktopWidget) SetVisible(visible bool) {
	defer qt.Recovering("QDesktopWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDesktopWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDesktopWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDesktopWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDesktopWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQDesktopWidgetShowEvent
func callbackQDesktopWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDesktopWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesktopWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQDesktopWidgetChangeEvent
func callbackQDesktopWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesktopWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDesktopWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQDesktopWidgetCloseEvent
func callbackQDesktopWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDesktopWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDesktopWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQDesktopWidgetContextMenuEvent
func callbackQDesktopWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDesktopWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QDesktopWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QDesktopWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QDesktopWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQDesktopWidgetInitPainter
func callbackQDesktopWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQDesktopWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QDesktopWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDesktopWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDesktopWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDesktopWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDesktopWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDesktopWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQDesktopWidgetInputMethodEvent
func callbackQDesktopWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDesktopWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesktopWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQDesktopWidgetKeyPressEvent
func callbackQDesktopWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesktopWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesktopWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQDesktopWidgetKeyReleaseEvent
func callbackQDesktopWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesktopWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesktopWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQDesktopWidgetMouseDoubleClickEvent
func callbackQDesktopWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesktopWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesktopWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDesktopWidgetMouseMoveEvent
func callbackQDesktopWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesktopWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesktopWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDesktopWidgetMousePressEvent
func callbackQDesktopWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesktopWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesktopWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDesktopWidgetMouseReleaseEvent
func callbackQDesktopWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesktopWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDesktopWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQDesktopWidgetTabletEvent
func callbackQDesktopWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDesktopWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDesktopWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQDesktopWidgetWheelEvent
func callbackQDesktopWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDesktopWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDesktopWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDesktopWidgetTimerEvent
func callbackQDesktopWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDesktopWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDesktopWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDesktopWidgetChildEvent
func callbackQDesktopWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDesktopWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDesktopWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesktopWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDesktopWidgetCustomEvent
func callbackQDesktopWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesktopWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesktopWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesktopWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesktopWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesktopWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
