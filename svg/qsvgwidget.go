package svg

//#include "svg.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QSvgWidget struct {
	widgets.QWidget
}

type QSvgWidget_ITF interface {
	widgets.QWidget_ITF
	QSvgWidget_PTR() *QSvgWidget
}

func PointerFromQSvgWidget(ptr QSvgWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgWidget_PTR().Pointer()
	}
	return nil
}

func NewQSvgWidgetFromPointer(ptr unsafe.Pointer) *QSvgWidget {
	var n = new(QSvgWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSvgWidget_") {
		n.SetObjectName("QSvgWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QSvgWidget) QSvgWidget_PTR() *QSvgWidget {
	return ptr
}

func NewQSvgWidget(parent widgets.QWidget_ITF) *QSvgWidget {
	defer qt.Recovering("QSvgWidget::QSvgWidget")

	return NewQSvgWidgetFromPointer(C.QSvgWidget_NewQSvgWidget(widgets.PointerFromQWidget(parent)))
}

func NewQSvgWidget2(file string, parent widgets.QWidget_ITF) *QSvgWidget {
	defer qt.Recovering("QSvgWidget::QSvgWidget")

	return NewQSvgWidgetFromPointer(C.QSvgWidget_NewQSvgWidget2(C.CString(file), widgets.PointerFromQWidget(parent)))
}

func (ptr *QSvgWidget) Load2(contents core.QByteArray_ITF) {
	defer qt.Recovering("QSvgWidget::load")

	if ptr.Pointer() != nil {
		C.QSvgWidget_Load2(ptr.Pointer(), core.PointerFromQByteArray(contents))
	}
}

func (ptr *QSvgWidget) Load(file string) {
	defer qt.Recovering("QSvgWidget::load")

	if ptr.Pointer() != nil {
		C.QSvgWidget_Load(ptr.Pointer(), C.CString(file))
	}
}

func (ptr *QSvgWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSvgWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSvgWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSvgWidgetPaintEvent
func callbackQSvgWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSvgWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QSvgWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSvgWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QSvgWidget) Renderer() *QSvgRenderer {
	defer qt.Recovering("QSvgWidget::renderer")

	if ptr.Pointer() != nil {
		return NewQSvgRendererFromPointer(C.QSvgWidget_Renderer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QSvgWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSvgWidget_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgWidget) DestroyQSvgWidget() {
	defer qt.Recovering("QSvgWidget::~QSvgWidget")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DestroyQSvgWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSvgWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQSvgWidgetActionEvent
func callbackQSvgWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSvgWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQSvgWidgetDragEnterEvent
func callbackQSvgWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSvgWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQSvgWidgetDragLeaveEvent
func callbackQSvgWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSvgWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQSvgWidgetDragMoveEvent
func callbackQSvgWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSvgWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQSvgWidgetDropEvent
func callbackQSvgWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSvgWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQSvgWidgetEnterEvent
func callbackQSvgWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQSvgWidgetFocusInEvent
func callbackQSvgWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSvgWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQSvgWidgetFocusOutEvent
func callbackQSvgWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSvgWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQSvgWidgetHideEvent
func callbackQSvgWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSvgWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQSvgWidgetLeaveEvent
func callbackQSvgWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQSvgWidgetMoveEvent
func callbackQSvgWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSvgWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QSvgWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QSvgWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QSvgWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQSvgWidgetSetVisible
func callbackQSvgWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSvgWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QSvgWidget) SetVisible(visible bool) {
	defer qt.Recovering("QSvgWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSvgWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QSvgWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSvgWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QSvgWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QSvgWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQSvgWidgetShowEvent
func callbackQSvgWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSvgWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSvgWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSvgWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQSvgWidgetChangeEvent
func callbackQSvgWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQSvgWidgetCloseEvent
func callbackQSvgWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSvgWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQSvgWidgetContextMenuEvent
func callbackQSvgWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSvgWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSvgWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QSvgWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QSvgWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQSvgWidgetInitPainter
func callbackQSvgWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQSvgWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QSvgWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSvgWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QSvgWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSvgWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSvgWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QSvgWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSvgWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQSvgWidgetInputMethodEvent
func callbackQSvgWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSvgWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQSvgWidgetKeyPressEvent
func callbackQSvgWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSvgWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQSvgWidgetKeyReleaseEvent
func callbackQSvgWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSvgWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQSvgWidgetMouseDoubleClickEvent
func callbackQSvgWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSvgWidgetMouseMoveEvent
func callbackQSvgWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSvgWidgetMousePressEvent
func callbackQSvgWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSvgWidgetMouseReleaseEvent
func callbackQSvgWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSvgWidgetResizeEvent
func callbackQSvgWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSvgWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQSvgWidgetTabletEvent
func callbackQSvgWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSvgWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQSvgWidgetWheelEvent
func callbackQSvgWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSvgWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSvgWidgetTimerEvent
func callbackQSvgWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSvgWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSvgWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSvgWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSvgWidgetChildEvent
func callbackQSvgWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSvgWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSvgWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSvgWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSvgWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSvgWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSvgWidgetCustomEvent
func callbackQSvgWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
