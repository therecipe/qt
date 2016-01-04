package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QStackedWidget struct {
	QFrame
}

type QStackedWidget_ITF interface {
	QFrame_ITF
	QStackedWidget_PTR() *QStackedWidget
}

func PointerFromQStackedWidget(ptr QStackedWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStackedWidget_PTR().Pointer()
	}
	return nil
}

func NewQStackedWidgetFromPointer(ptr unsafe.Pointer) *QStackedWidget {
	var n = new(QStackedWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStackedWidget_") {
		n.SetObjectName("QStackedWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QStackedWidget) QStackedWidget_PTR() *QStackedWidget {
	return ptr
}

func (ptr *QStackedWidget) Count() int {
	defer qt.Recovering("QStackedWidget::count")

	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedWidget) CurrentIndex() int {
	defer qt.Recovering("QStackedWidget::currentIndex")

	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedWidget) SetCurrentIndex(index int) {
	defer qt.Recovering("QStackedWidget::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QStackedWidget_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QStackedWidget) SetCurrentWidget(widget QWidget_ITF) {
	defer qt.Recovering("QStackedWidget::setCurrentWidget")

	if ptr.Pointer() != nil {
		C.QStackedWidget_SetCurrentWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func NewQStackedWidget(parent QWidget_ITF) *QStackedWidget {
	defer qt.Recovering("QStackedWidget::QStackedWidget")

	return NewQStackedWidgetFromPointer(C.QStackedWidget_NewQStackedWidget(PointerFromQWidget(parent)))
}

func (ptr *QStackedWidget) AddWidget(widget QWidget_ITF) int {
	defer qt.Recovering("QStackedWidget::addWidget")

	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_AddWidget(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedWidget) ConnectCurrentChanged(f func(index int)) {
	defer qt.Recovering("connect QStackedWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QStackedWidget) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QStackedWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQStackedWidgetCurrentChanged
func callbackQStackedWidgetCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QStackedWidget::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QStackedWidget) CurrentChanged(index int) {
	defer qt.Recovering("QStackedWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QStackedWidget_CurrentChanged(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QStackedWidget) CurrentWidget() *QWidget {
	defer qt.Recovering("QStackedWidget::currentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedWidget_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStackedWidget) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QStackedWidget::event")

	if ptr.Pointer() != nil {
		return C.QStackedWidget_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QStackedWidget) IndexOf(widget QWidget_ITF) int {
	defer qt.Recovering("QStackedWidget::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_IndexOf(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedWidget) InsertWidget(index int, widget QWidget_ITF) int {
	defer qt.Recovering("QStackedWidget::insertWidget")

	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedWidget) RemoveWidget(widget QWidget_ITF) {
	defer qt.Recovering("QStackedWidget::removeWidget")

	if ptr.Pointer() != nil {
		C.QStackedWidget_RemoveWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStackedWidget) Widget(index int) *QWidget {
	defer qt.Recovering("QStackedWidget::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedWidget_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedWidget) ConnectWidgetRemoved(f func(index int)) {
	defer qt.Recovering("connect QStackedWidget::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ConnectWidgetRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "widgetRemoved", f)
	}
}

func (ptr *QStackedWidget) DisconnectWidgetRemoved() {
	defer qt.Recovering("disconnect QStackedWidget::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DisconnectWidgetRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "widgetRemoved")
	}
}

//export callbackQStackedWidgetWidgetRemoved
func callbackQStackedWidgetWidgetRemoved(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QStackedWidget::widgetRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "widgetRemoved"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QStackedWidget) WidgetRemoved(index int) {
	defer qt.Recovering("QStackedWidget::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QStackedWidget_WidgetRemoved(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QStackedWidget) DestroyQStackedWidget() {
	defer qt.Recovering("QStackedWidget::~QStackedWidget")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DestroyQStackedWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStackedWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QStackedWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QStackedWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQStackedWidgetChangeEvent
func callbackQStackedWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQStackedWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QStackedWidget) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QStackedWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QStackedWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QStackedWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QStackedWidget) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QStackedWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QStackedWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQStackedWidgetPaintEvent
func callbackQStackedWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQStackedWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QStackedWidget) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QStackedWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QStackedWidget) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QStackedWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QStackedWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QStackedWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QStackedWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQStackedWidgetActionEvent
func callbackQStackedWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QStackedWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QStackedWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QStackedWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QStackedWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QStackedWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQStackedWidgetDragEnterEvent
func callbackQStackedWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QStackedWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QStackedWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QStackedWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QStackedWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QStackedWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQStackedWidgetDragLeaveEvent
func callbackQStackedWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QStackedWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QStackedWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QStackedWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QStackedWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QStackedWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQStackedWidgetDragMoveEvent
func callbackQStackedWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QStackedWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QStackedWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QStackedWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QStackedWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QStackedWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQStackedWidgetDropEvent
func callbackQStackedWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QStackedWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QStackedWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QStackedWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStackedWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QStackedWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQStackedWidgetEnterEvent
func callbackQStackedWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QStackedWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStackedWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QStackedWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QStackedWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QStackedWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQStackedWidgetFocusInEvent
func callbackQStackedWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QStackedWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QStackedWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QStackedWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QStackedWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QStackedWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQStackedWidgetFocusOutEvent
func callbackQStackedWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QStackedWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QStackedWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QStackedWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QStackedWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QStackedWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQStackedWidgetHideEvent
func callbackQStackedWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QStackedWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QStackedWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QStackedWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStackedWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QStackedWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQStackedWidgetLeaveEvent
func callbackQStackedWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QStackedWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStackedWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QStackedWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QStackedWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QStackedWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQStackedWidgetMoveEvent
func callbackQStackedWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QStackedWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QStackedWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QStackedWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QStackedWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QStackedWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QStackedWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQStackedWidgetSetVisible
func callbackQStackedWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QStackedWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QStackedWidget) SetVisible(visible bool) {
	defer qt.Recovering("QStackedWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QStackedWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QStackedWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QStackedWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QStackedWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QStackedWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QStackedWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QStackedWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQStackedWidgetShowEvent
func callbackQStackedWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QStackedWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QStackedWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QStackedWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QStackedWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QStackedWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQStackedWidgetCloseEvent
func callbackQStackedWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QStackedWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QStackedWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QStackedWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QStackedWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QStackedWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQStackedWidgetContextMenuEvent
func callbackQStackedWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QStackedWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QStackedWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QStackedWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QStackedWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QStackedWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QStackedWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQStackedWidgetInitPainter
func callbackQStackedWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQStackedWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QStackedWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QStackedWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QStackedWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QStackedWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QStackedWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QStackedWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QStackedWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QStackedWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QStackedWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQStackedWidgetInputMethodEvent
func callbackQStackedWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QStackedWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QStackedWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QStackedWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QStackedWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QStackedWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQStackedWidgetKeyPressEvent
func callbackQStackedWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QStackedWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QStackedWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QStackedWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QStackedWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QStackedWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQStackedWidgetKeyReleaseEvent
func callbackQStackedWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QStackedWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QStackedWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QStackedWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStackedWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QStackedWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQStackedWidgetMouseDoubleClickEvent
func callbackQStackedWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStackedWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStackedWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStackedWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStackedWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QStackedWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQStackedWidgetMouseMoveEvent
func callbackQStackedWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStackedWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStackedWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStackedWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStackedWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QStackedWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQStackedWidgetMousePressEvent
func callbackQStackedWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStackedWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStackedWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStackedWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStackedWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QStackedWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQStackedWidgetMouseReleaseEvent
func callbackQStackedWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStackedWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStackedWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QStackedWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QStackedWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QStackedWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQStackedWidgetResizeEvent
func callbackQStackedWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QStackedWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QStackedWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QStackedWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QStackedWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QStackedWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQStackedWidgetTabletEvent
func callbackQStackedWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QStackedWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QStackedWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QStackedWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QStackedWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QStackedWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQStackedWidgetWheelEvent
func callbackQStackedWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QStackedWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QStackedWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QStackedWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QStackedWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QStackedWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQStackedWidgetTimerEvent
func callbackQStackedWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QStackedWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QStackedWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QStackedWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QStackedWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QStackedWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQStackedWidgetChildEvent
func callbackQStackedWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QStackedWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QStackedWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QStackedWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QStackedWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStackedWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QStackedWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QStackedWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQStackedWidgetCustomEvent
func callbackQStackedWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStackedWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQStackedWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStackedWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QStackedWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStackedWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QStackedWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QStackedWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
