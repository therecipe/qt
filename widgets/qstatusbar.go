package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QStatusBar struct {
	QWidget
}

type QStatusBar_ITF interface {
	QWidget_ITF
	QStatusBar_PTR() *QStatusBar
}

func PointerFromQStatusBar(ptr QStatusBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStatusBar_PTR().Pointer()
	}
	return nil
}

func NewQStatusBarFromPointer(ptr unsafe.Pointer) *QStatusBar {
	var n = new(QStatusBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStatusBar_") {
		n.SetObjectName("QStatusBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QStatusBar) QStatusBar_PTR() *QStatusBar {
	return ptr
}

func (ptr *QStatusBar) IsSizeGripEnabled() bool {
	defer qt.Recovering("QStatusBar::isSizeGripEnabled")

	if ptr.Pointer() != nil {
		return C.QStatusBar_IsSizeGripEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStatusBar) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QStatusBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QStatusBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQStatusBarPaintEvent
func callbackQStatusBarPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QStatusBar) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QStatusBar::paintEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QStatusBar) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QStatusBar::paintEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QStatusBar) SetSizeGripEnabled(v bool) {
	defer qt.Recovering("QStatusBar::setSizeGripEnabled")

	if ptr.Pointer() != nil {
		C.QStatusBar_SetSizeGripEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQStatusBar(parent QWidget_ITF) *QStatusBar {
	defer qt.Recovering("QStatusBar::QStatusBar")

	return NewQStatusBarFromPointer(C.QStatusBar_NewQStatusBar(PointerFromQWidget(parent)))
}

func (ptr *QStatusBar) AddPermanentWidget(widget QWidget_ITF, stretch int) {
	defer qt.Recovering("QStatusBar::addPermanentWidget")

	if ptr.Pointer() != nil {
		C.QStatusBar_AddPermanentWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch))
	}
}

func (ptr *QStatusBar) AddWidget(widget QWidget_ITF, stretch int) {
	defer qt.Recovering("QStatusBar::addWidget")

	if ptr.Pointer() != nil {
		C.QStatusBar_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch))
	}
}

func (ptr *QStatusBar) ClearMessage() {
	defer qt.Recovering("QStatusBar::clearMessage")

	if ptr.Pointer() != nil {
		C.QStatusBar_ClearMessage(ptr.Pointer())
	}
}

func (ptr *QStatusBar) CurrentMessage() string {
	defer qt.Recovering("QStatusBar::currentMessage")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStatusBar_CurrentMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStatusBar) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QStatusBar::event")

	if ptr.Pointer() != nil {
		return C.QStatusBar_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QStatusBar) InsertPermanentWidget(index int, widget QWidget_ITF, stretch int) int {
	defer qt.Recovering("QStatusBar::insertPermanentWidget")

	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertPermanentWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) InsertWidget(index int, widget QWidget_ITF, stretch int) int {
	defer qt.Recovering("QStatusBar::insertWidget")

	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) ConnectMessageChanged(f func(message string)) {
	defer qt.Recovering("connect QStatusBar::messageChanged")

	if ptr.Pointer() != nil {
		C.QStatusBar_ConnectMessageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageChanged", f)
	}
}

func (ptr *QStatusBar) DisconnectMessageChanged() {
	defer qt.Recovering("disconnect QStatusBar::messageChanged")

	if ptr.Pointer() != nil {
		C.QStatusBar_DisconnectMessageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageChanged")
	}
}

//export callbackQStatusBarMessageChanged
func callbackQStatusBarMessageChanged(ptr unsafe.Pointer, ptrName *C.char, message *C.char) {
	defer qt.Recovering("callback QStatusBar::messageChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "messageChanged"); signal != nil {
		signal.(func(string))(C.GoString(message))
	}

}

func (ptr *QStatusBar) MessageChanged(message string) {
	defer qt.Recovering("QStatusBar::messageChanged")

	if ptr.Pointer() != nil {
		C.QStatusBar_MessageChanged(ptr.Pointer(), C.CString(message))
	}
}

func (ptr *QStatusBar) RemoveWidget(widget QWidget_ITF) {
	defer qt.Recovering("QStatusBar::removeWidget")

	if ptr.Pointer() != nil {
		C.QStatusBar_RemoveWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStatusBar) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QStatusBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QStatusBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQStatusBarResizeEvent
func callbackQStatusBarResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQStatusBarFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QStatusBar) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QStatusBar::resizeEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QStatusBar) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QStatusBar::resizeEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QStatusBar) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QStatusBar::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QStatusBar::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQStatusBarShowEvent
func callbackQStatusBarShowEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
	} else {
		NewQStatusBarFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(v))
	}
}

func (ptr *QStatusBar) ShowEvent(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QStatusBar::showEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QStatusBar) ShowEventDefault(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QStatusBar::showEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QStatusBar) ShowMessage(message string, timeout int) {
	defer qt.Recovering("QStatusBar::showMessage")

	if ptr.Pointer() != nil {
		C.QStatusBar_ShowMessage(ptr.Pointer(), C.CString(message), C.int(timeout))
	}
}

func (ptr *QStatusBar) DestroyQStatusBar() {
	defer qt.Recovering("QStatusBar::~QStatusBar")

	if ptr.Pointer() != nil {
		C.QStatusBar_DestroyQStatusBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStatusBar) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QStatusBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QStatusBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQStatusBarActionEvent
func callbackQStatusBarActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QStatusBar) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QStatusBar::actionEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QStatusBar) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QStatusBar::actionEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QStatusBar) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QStatusBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QStatusBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQStatusBarDragEnterEvent
func callbackQStatusBarDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QStatusBar) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QStatusBar::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QStatusBar) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QStatusBar::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QStatusBar) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QStatusBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QStatusBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQStatusBarDragLeaveEvent
func callbackQStatusBarDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QStatusBar) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QStatusBar::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QStatusBar) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QStatusBar::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QStatusBar) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QStatusBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QStatusBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQStatusBarDragMoveEvent
func callbackQStatusBarDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QStatusBar) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QStatusBar::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QStatusBar) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QStatusBar::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QStatusBar) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QStatusBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QStatusBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQStatusBarDropEvent
func callbackQStatusBarDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QStatusBar) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QStatusBar::dropEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QStatusBar) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QStatusBar::dropEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QStatusBar) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStatusBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QStatusBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQStatusBarEnterEvent
func callbackQStatusBarEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStatusBar) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QStatusBar::enterEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStatusBar) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QStatusBar::enterEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStatusBar) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QStatusBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QStatusBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQStatusBarFocusInEvent
func callbackQStatusBarFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QStatusBar) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QStatusBar::focusInEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QStatusBar) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QStatusBar::focusInEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QStatusBar) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QStatusBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QStatusBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQStatusBarFocusOutEvent
func callbackQStatusBarFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QStatusBar) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QStatusBar::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QStatusBar) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QStatusBar::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QStatusBar) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QStatusBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QStatusBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQStatusBarHideEvent
func callbackQStatusBarHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QStatusBar) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QStatusBar::hideEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QStatusBar) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QStatusBar::hideEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QStatusBar) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStatusBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QStatusBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQStatusBarLeaveEvent
func callbackQStatusBarLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStatusBar) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QStatusBar::leaveEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStatusBar) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QStatusBar::leaveEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStatusBar) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QStatusBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QStatusBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQStatusBarMoveEvent
func callbackQStatusBarMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QStatusBar) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QStatusBar::moveEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QStatusBar) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QStatusBar::moveEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QStatusBar) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QStatusBar::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QStatusBar) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QStatusBar::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQStatusBarSetVisible
func callbackQStatusBarSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QStatusBar::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QStatusBar) SetVisible(visible bool) {
	defer qt.Recovering("QStatusBar::setVisible")

	if ptr.Pointer() != nil {
		C.QStatusBar_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QStatusBar) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QStatusBar::setVisible")

	if ptr.Pointer() != nil {
		C.QStatusBar_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QStatusBar) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStatusBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QStatusBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQStatusBarChangeEvent
func callbackQStatusBarChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStatusBar) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QStatusBar::changeEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStatusBar) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QStatusBar::changeEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStatusBar) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QStatusBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QStatusBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQStatusBarCloseEvent
func callbackQStatusBarCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QStatusBar) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QStatusBar::closeEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QStatusBar) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QStatusBar::closeEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QStatusBar) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QStatusBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QStatusBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQStatusBarContextMenuEvent
func callbackQStatusBarContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QStatusBar) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QStatusBar::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QStatusBar) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QStatusBar::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QStatusBar) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QStatusBar::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QStatusBar) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QStatusBar::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQStatusBarInitPainter
func callbackQStatusBarInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQStatusBarFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QStatusBar) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QStatusBar::initPainter")

	if ptr.Pointer() != nil {
		C.QStatusBar_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QStatusBar) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QStatusBar::initPainter")

	if ptr.Pointer() != nil {
		C.QStatusBar_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QStatusBar) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QStatusBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QStatusBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQStatusBarInputMethodEvent
func callbackQStatusBarInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QStatusBar) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QStatusBar::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QStatusBar) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QStatusBar::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QStatusBar) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QStatusBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QStatusBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQStatusBarKeyPressEvent
func callbackQStatusBarKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QStatusBar) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QStatusBar::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QStatusBar) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QStatusBar::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QStatusBar) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QStatusBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QStatusBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQStatusBarKeyReleaseEvent
func callbackQStatusBarKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QStatusBar) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QStatusBar::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QStatusBar) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QStatusBar::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QStatusBar) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStatusBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QStatusBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQStatusBarMouseDoubleClickEvent
func callbackQStatusBarMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QStatusBar) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStatusBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStatusBar) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStatusBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStatusBar) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStatusBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QStatusBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQStatusBarMouseMoveEvent
func callbackQStatusBarMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QStatusBar) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStatusBar::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStatusBar) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStatusBar::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStatusBar) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStatusBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QStatusBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQStatusBarMousePressEvent
func callbackQStatusBarMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QStatusBar) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStatusBar::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStatusBar) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStatusBar::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStatusBar) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStatusBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QStatusBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQStatusBarMouseReleaseEvent
func callbackQStatusBarMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QStatusBar) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStatusBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStatusBar) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStatusBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStatusBar) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QStatusBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QStatusBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQStatusBarTabletEvent
func callbackQStatusBarTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QStatusBar) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QStatusBar::tabletEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QStatusBar) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QStatusBar::tabletEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QStatusBar) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QStatusBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QStatusBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQStatusBarWheelEvent
func callbackQStatusBarWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QStatusBar) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QStatusBar::wheelEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QStatusBar) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QStatusBar::wheelEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QStatusBar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QStatusBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QStatusBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQStatusBarTimerEvent
func callbackQStatusBarTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QStatusBar) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QStatusBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QStatusBar) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QStatusBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QStatusBar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QStatusBar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QStatusBar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQStatusBarChildEvent
func callbackQStatusBarChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QStatusBar) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QStatusBar::childEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QStatusBar) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QStatusBar::childEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QStatusBar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStatusBar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QStatusBar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQStatusBarCustomEvent
func callbackQStatusBarCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStatusBar::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQStatusBarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStatusBar) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QStatusBar::customEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStatusBar) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QStatusBar::customEvent")

	if ptr.Pointer() != nil {
		C.QStatusBar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
