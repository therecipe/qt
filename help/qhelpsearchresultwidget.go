package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QHelpSearchResultWidget struct {
	widgets.QWidget
}

type QHelpSearchResultWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget
}

func PointerFromQHelpSearchResultWidget(ptr QHelpSearchResultWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchResultWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchResultWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchResultWidget {
	var n = new(QHelpSearchResultWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpSearchResultWidget_") {
		n.SetObjectName("QHelpSearchResultWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpSearchResultWidget) QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget {
	return ptr
}

func (ptr *QHelpSearchResultWidget) LinkAt(point core.QPoint_ITF) *core.QUrl {
	defer qt.Recovering("QHelpSearchResultWidget::linkAt")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QHelpSearchResultWidget_LinkAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) ConnectRequestShowLink(f func(link *core.QUrl)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectRequestShowLink(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestShowLink", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRequestShowLink() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectRequestShowLink(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestShowLink")
	}
}

//export callbackQHelpSearchResultWidgetRequestShowLink
func callbackQHelpSearchResultWidgetRequestShowLink(ptr unsafe.Pointer, ptrName *C.char, link unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::requestShowLink")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestShowLink"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpSearchResultWidget) RequestShowLink(link core.QUrl_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RequestShowLink(ptr.Pointer(), core.PointerFromQUrl(link))
	}
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidget() {
	defer qt.Recovering("QHelpSearchResultWidget::~QHelpSearchResultWidget")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQHelpSearchResultWidgetActionEvent
func callbackQHelpSearchResultWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQHelpSearchResultWidgetDragEnterEvent
func callbackQHelpSearchResultWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQHelpSearchResultWidgetDragLeaveEvent
func callbackQHelpSearchResultWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQHelpSearchResultWidgetDragMoveEvent
func callbackQHelpSearchResultWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQHelpSearchResultWidgetDropEvent
func callbackQHelpSearchResultWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQHelpSearchResultWidgetEnterEvent
func callbackQHelpSearchResultWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQHelpSearchResultWidgetFocusInEvent
func callbackQHelpSearchResultWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQHelpSearchResultWidgetFocusOutEvent
func callbackQHelpSearchResultWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQHelpSearchResultWidgetHideEvent
func callbackQHelpSearchResultWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQHelpSearchResultWidgetLeaveEvent
func callbackQHelpSearchResultWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQHelpSearchResultWidgetMoveEvent
func callbackQHelpSearchResultWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQHelpSearchResultWidgetPaintEvent
func callbackQHelpSearchResultWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQHelpSearchResultWidgetSetVisible
func callbackQHelpSearchResultWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QHelpSearchResultWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QHelpSearchResultWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpSearchResultWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQHelpSearchResultWidgetShowEvent
func callbackQHelpSearchResultWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQHelpSearchResultWidgetCloseEvent
func callbackQHelpSearchResultWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQHelpSearchResultWidgetContextMenuEvent
func callbackQHelpSearchResultWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQHelpSearchResultWidgetInitPainter
func callbackQHelpSearchResultWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpSearchResultWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpSearchResultWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQHelpSearchResultWidgetInputMethodEvent
func callbackQHelpSearchResultWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQHelpSearchResultWidgetKeyPressEvent
func callbackQHelpSearchResultWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQHelpSearchResultWidgetKeyReleaseEvent
func callbackQHelpSearchResultWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQHelpSearchResultWidgetMouseDoubleClickEvent
func callbackQHelpSearchResultWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQHelpSearchResultWidgetMouseMoveEvent
func callbackQHelpSearchResultWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQHelpSearchResultWidgetMousePressEvent
func callbackQHelpSearchResultWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQHelpSearchResultWidgetMouseReleaseEvent
func callbackQHelpSearchResultWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQHelpSearchResultWidgetResizeEvent
func callbackQHelpSearchResultWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQHelpSearchResultWidgetTabletEvent
func callbackQHelpSearchResultWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQHelpSearchResultWidgetWheelEvent
func callbackQHelpSearchResultWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpSearchResultWidgetTimerEvent
func callbackQHelpSearchResultWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpSearchResultWidgetChildEvent
func callbackQHelpSearchResultWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpSearchResultWidgetCustomEvent
func callbackQHelpSearchResultWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
