package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QFocusFrame struct {
	QWidget
}

type QFocusFrame_ITF interface {
	QWidget_ITF
	QFocusFrame_PTR() *QFocusFrame
}

func PointerFromQFocusFrame(ptr QFocusFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFocusFrame_PTR().Pointer()
	}
	return nil
}

func NewQFocusFrameFromPointer(ptr unsafe.Pointer) *QFocusFrame {
	var n = new(QFocusFrame)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFocusFrame_") {
		n.SetObjectName("QFocusFrame_" + qt.Identifier())
	}
	return n
}

func (ptr *QFocusFrame) QFocusFrame_PTR() *QFocusFrame {
	return ptr
}

func NewQFocusFrame(parent QWidget_ITF) *QFocusFrame {
	defer qt.Recovering("QFocusFrame::QFocusFrame")

	return NewQFocusFrameFromPointer(C.QFocusFrame_NewQFocusFrame(PointerFromQWidget(parent)))
}

func (ptr *QFocusFrame) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QFocusFrame::event")

	if ptr.Pointer() != nil {
		return C.QFocusFrame_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QFocusFrame) EventFilter(o core.QObject_ITF, e core.QEvent_ITF) bool {
	defer qt.Recovering("QFocusFrame::eventFilter")

	if ptr.Pointer() != nil {
		return C.QFocusFrame_EventFilter(ptr.Pointer(), core.PointerFromQObject(o), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QFocusFrame) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QFocusFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QFocusFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQFocusFramePaintEvent
func callbackQFocusFramePaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQFocusFrameFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QFocusFrame) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QFocusFrame::paintEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QFocusFrame) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QFocusFrame::paintEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QFocusFrame) SetWidget(widget QWidget_ITF) {
	defer qt.Recovering("QFocusFrame::setWidget")

	if ptr.Pointer() != nil {
		C.QFocusFrame_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QFocusFrame) Widget() *QWidget {
	defer qt.Recovering("QFocusFrame::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QFocusFrame_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFocusFrame) DestroyQFocusFrame() {
	defer qt.Recovering("QFocusFrame::~QFocusFrame")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DestroyQFocusFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFocusFrame) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QFocusFrame::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QFocusFrame::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQFocusFrameActionEvent
func callbackQFocusFrameActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QFocusFrame::actionEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QFocusFrame) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QFocusFrame::actionEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QFocusFrame::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QFocusFrame::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQFocusFrameDragEnterEvent
func callbackQFocusFrameDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QFocusFrame::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QFocusFrame) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QFocusFrame::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QFocusFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QFocusFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQFocusFrameDragLeaveEvent
func callbackQFocusFrameDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QFocusFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QFocusFrame) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QFocusFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QFocusFrame::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QFocusFrame::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQFocusFrameDragMoveEvent
func callbackQFocusFrameDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QFocusFrame::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QFocusFrame) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QFocusFrame::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QFocusFrame::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QFocusFrame::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQFocusFrameDropEvent
func callbackQFocusFrameDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QFocusFrame::dropEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QFocusFrame) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QFocusFrame::dropEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFocusFrame::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QFocusFrame::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQFocusFrameEnterEvent
func callbackQFocusFrameEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFocusFrame::enterEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFocusFrame) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFocusFrame::enterEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFocusFrame::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QFocusFrame::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQFocusFrameFocusInEvent
func callbackQFocusFrameFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFocusFrame::focusInEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QFocusFrame) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFocusFrame::focusInEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFocusFrame::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QFocusFrame::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQFocusFrameFocusOutEvent
func callbackQFocusFrameFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFocusFrame::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QFocusFrame) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFocusFrame::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QFocusFrame::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QFocusFrame::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQFocusFrameHideEvent
func callbackQFocusFrameHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QFocusFrame::hideEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QFocusFrame) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QFocusFrame::hideEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFocusFrame::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QFocusFrame::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQFocusFrameLeaveEvent
func callbackQFocusFrameLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFocusFrame::leaveEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFocusFrame) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFocusFrame::leaveEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QFocusFrame::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QFocusFrame::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQFocusFrameMoveEvent
func callbackQFocusFrameMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QFocusFrame::moveEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QFocusFrame) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QFocusFrame::moveEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QFocusFrame::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QFocusFrame) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QFocusFrame::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQFocusFrameSetVisible
func callbackQFocusFrameSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QFocusFrame::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QFocusFrame) SetVisible(visible bool) {
	defer qt.Recovering("QFocusFrame::setVisible")

	if ptr.Pointer() != nil {
		C.QFocusFrame_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFocusFrame) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QFocusFrame::setVisible")

	if ptr.Pointer() != nil {
		C.QFocusFrame_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFocusFrame) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QFocusFrame::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QFocusFrame::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQFocusFrameShowEvent
func callbackQFocusFrameShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QFocusFrame::showEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QFocusFrame) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QFocusFrame::showEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFocusFrame::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QFocusFrame::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQFocusFrameChangeEvent
func callbackQFocusFrameChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFocusFrame::changeEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFocusFrame) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFocusFrame::changeEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QFocusFrame::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QFocusFrame::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQFocusFrameCloseEvent
func callbackQFocusFrameCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QFocusFrame::closeEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QFocusFrame) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QFocusFrame::closeEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QFocusFrame::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QFocusFrame::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQFocusFrameContextMenuEvent
func callbackQFocusFrameContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QFocusFrame::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QFocusFrame) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QFocusFrame::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QFocusFrame::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QFocusFrame) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QFocusFrame::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQFocusFrameInitPainter
func callbackQFocusFrameInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQFocusFrameFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QFocusFrame) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QFocusFrame::initPainter")

	if ptr.Pointer() != nil {
		C.QFocusFrame_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QFocusFrame) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QFocusFrame::initPainter")

	if ptr.Pointer() != nil {
		C.QFocusFrame_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QFocusFrame) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QFocusFrame::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QFocusFrame::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQFocusFrameInputMethodEvent
func callbackQFocusFrameInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QFocusFrame::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QFocusFrame) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QFocusFrame::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFocusFrame::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QFocusFrame::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQFocusFrameKeyPressEvent
func callbackQFocusFrameKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFocusFrame::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QFocusFrame) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFocusFrame::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFocusFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QFocusFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQFocusFrameKeyReleaseEvent
func callbackQFocusFrameKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFocusFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QFocusFrame) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFocusFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFocusFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QFocusFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQFocusFrameMouseDoubleClickEvent
func callbackQFocusFrameMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFocusFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFocusFrame) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFocusFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFocusFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QFocusFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQFocusFrameMouseMoveEvent
func callbackQFocusFrameMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFocusFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFocusFrame) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFocusFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFocusFrame::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QFocusFrame::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQFocusFrameMousePressEvent
func callbackQFocusFrameMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFocusFrame::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFocusFrame) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFocusFrame::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFocusFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QFocusFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQFocusFrameMouseReleaseEvent
func callbackQFocusFrameMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFocusFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFocusFrame) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFocusFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QFocusFrame::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QFocusFrame::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQFocusFrameResizeEvent
func callbackQFocusFrameResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QFocusFrame::resizeEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QFocusFrame) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QFocusFrame::resizeEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QFocusFrame::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QFocusFrame::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQFocusFrameTabletEvent
func callbackQFocusFrameTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QFocusFrame::tabletEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QFocusFrame) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QFocusFrame::tabletEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QFocusFrame::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QFocusFrame::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQFocusFrameWheelEvent
func callbackQFocusFrameWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QFocusFrame::wheelEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QFocusFrame) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QFocusFrame::wheelEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QFocusFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFocusFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFocusFrameTimerEvent
func callbackQFocusFrameTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QFocusFrame::timerEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QFocusFrame) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QFocusFrame::timerEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QFocusFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFocusFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFocusFrameChildEvent
func callbackQFocusFrameChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QFocusFrame::childEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QFocusFrame) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QFocusFrame::childEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QFocusFrame) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFocusFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFocusFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFocusFrameCustomEvent
func callbackQFocusFrameCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFocusFrame::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFocusFrameFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFocusFrame) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFocusFrame::customEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFocusFrame) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFocusFrame::customEvent")

	if ptr.Pointer() != nil {
		C.QFocusFrame_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
