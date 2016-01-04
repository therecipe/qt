package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMacNativeWidget struct {
	QWidget
}

type QMacNativeWidget_ITF interface {
	QWidget_ITF
	QMacNativeWidget_PTR() *QMacNativeWidget
}

func PointerFromQMacNativeWidget(ptr QMacNativeWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacNativeWidget_PTR().Pointer()
	}
	return nil
}

func NewQMacNativeWidgetFromPointer(ptr unsafe.Pointer) *QMacNativeWidget {
	var n = new(QMacNativeWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMacNativeWidget_") {
		n.SetObjectName("QMacNativeWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QMacNativeWidget) QMacNativeWidget_PTR() *QMacNativeWidget {
	return ptr
}

func (ptr *QMacNativeWidget) Event(ev core.QEvent_ITF) bool {
	defer qt.Recovering("QMacNativeWidget::event")

	if ptr.Pointer() != nil {
		return C.QMacNativeWidget_Event(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QMacNativeWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QMacNativeWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMacNativeWidget_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMacNativeWidget) DestroyQMacNativeWidget() {
	defer qt.Recovering("QMacNativeWidget::~QMacNativeWidget")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DestroyQMacNativeWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacNativeWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQMacNativeWidgetActionEvent
func callbackQMacNativeWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QMacNativeWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQMacNativeWidgetDragEnterEvent
func callbackQMacNativeWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QMacNativeWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQMacNativeWidgetDragLeaveEvent
func callbackQMacNativeWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QMacNativeWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQMacNativeWidgetDragMoveEvent
func callbackQMacNativeWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QMacNativeWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQMacNativeWidgetDropEvent
func callbackQMacNativeWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QMacNativeWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQMacNativeWidgetEnterEvent
func callbackQMacNativeWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacNativeWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQMacNativeWidgetFocusInEvent
func callbackQMacNativeWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMacNativeWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQMacNativeWidgetFocusOutEvent
func callbackQMacNativeWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMacNativeWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQMacNativeWidgetHideEvent
func callbackQMacNativeWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QMacNativeWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQMacNativeWidgetLeaveEvent
func callbackQMacNativeWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacNativeWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQMacNativeWidgetMoveEvent
func callbackQMacNativeWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QMacNativeWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQMacNativeWidgetPaintEvent
func callbackQMacNativeWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QMacNativeWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QMacNativeWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QMacNativeWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQMacNativeWidgetSetVisible
func callbackQMacNativeWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMacNativeWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMacNativeWidget) SetVisible(visible bool) {
	defer qt.Recovering("QMacNativeWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMacNativeWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QMacNativeWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMacNativeWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQMacNativeWidgetShowEvent
func callbackQMacNativeWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QMacNativeWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQMacNativeWidgetChangeEvent
func callbackQMacNativeWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacNativeWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQMacNativeWidgetCloseEvent
func callbackQMacNativeWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QMacNativeWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQMacNativeWidgetContextMenuEvent
func callbackQMacNativeWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QMacNativeWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QMacNativeWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QMacNativeWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQMacNativeWidgetInitPainter
func callbackQMacNativeWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QMacNativeWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QMacNativeWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QMacNativeWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QMacNativeWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QMacNativeWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQMacNativeWidgetInputMethodEvent
func callbackQMacNativeWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QMacNativeWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQMacNativeWidgetKeyPressEvent
func callbackQMacNativeWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMacNativeWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQMacNativeWidgetKeyReleaseEvent
func callbackQMacNativeWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMacNativeWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQMacNativeWidgetMouseDoubleClickEvent
func callbackQMacNativeWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacNativeWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQMacNativeWidgetMouseMoveEvent
func callbackQMacNativeWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacNativeWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQMacNativeWidgetMousePressEvent
func callbackQMacNativeWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacNativeWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQMacNativeWidgetMouseReleaseEvent
func callbackQMacNativeWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacNativeWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQMacNativeWidgetResizeEvent
func callbackQMacNativeWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QMacNativeWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQMacNativeWidgetTabletEvent
func callbackQMacNativeWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QMacNativeWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQMacNativeWidgetWheelEvent
func callbackQMacNativeWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QMacNativeWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMacNativeWidgetTimerEvent
func callbackQMacNativeWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacNativeWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMacNativeWidgetChildEvent
func callbackQMacNativeWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacNativeWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacNativeWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacNativeWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMacNativeWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMacNativeWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMacNativeWidgetCustomEvent
func callbackQMacNativeWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacNativeWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacNativeWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacNativeWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacNativeWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacNativeWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QMacNativeWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
