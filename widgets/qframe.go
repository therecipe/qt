package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QFrame struct {
	QWidget
}

type QFrame_ITF interface {
	QWidget_ITF
	QFrame_PTR() *QFrame
}

func PointerFromQFrame(ptr QFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func NewQFrameFromPointer(ptr unsafe.Pointer) *QFrame {
	var n = new(QFrame)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFrame_") {
		n.SetObjectName("QFrame_" + qt.Identifier())
	}
	return n
}

func (ptr *QFrame) QFrame_PTR() *QFrame {
	return ptr
}

//QFrame::Shadow
type QFrame__Shadow int64

const (
	QFrame__Plain  = QFrame__Shadow(0x0010)
	QFrame__Raised = QFrame__Shadow(0x0020)
	QFrame__Sunken = QFrame__Shadow(0x0030)
)

//QFrame::Shape
type QFrame__Shape int64

const (
	QFrame__NoFrame     = QFrame__Shape(0)
	QFrame__Box         = QFrame__Shape(0x0001)
	QFrame__Panel       = QFrame__Shape(0x0002)
	QFrame__WinPanel    = QFrame__Shape(0x0003)
	QFrame__HLine       = QFrame__Shape(0x0004)
	QFrame__VLine       = QFrame__Shape(0x0005)
	QFrame__StyledPanel = QFrame__Shape(0x0006)
)

//QFrame::StyleMask
type QFrame__StyleMask int64

var (
	QFrame__Shadow_Mask = QFrame__StyleMask(0x00f0)
	QFrame__Shape_Mask  = QFrame__StyleMask(0x000f)
)

func (ptr *QFrame) FrameRect() *core.QRect {
	defer qt.Recovering("QFrame::frameRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QFrame_FrameRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFrame) FrameShadow() QFrame__Shadow {
	defer qt.Recovering("QFrame::frameShadow")

	if ptr.Pointer() != nil {
		return QFrame__Shadow(C.QFrame_FrameShadow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) FrameShape() QFrame__Shape {
	defer qt.Recovering("QFrame::frameShape")

	if ptr.Pointer() != nil {
		return QFrame__Shape(C.QFrame_FrameShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) FrameWidth() int {
	defer qt.Recovering("QFrame::frameWidth")

	if ptr.Pointer() != nil {
		return int(C.QFrame_FrameWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) LineWidth() int {
	defer qt.Recovering("QFrame::lineWidth")

	if ptr.Pointer() != nil {
		return int(C.QFrame_LineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) MidLineWidth() int {
	defer qt.Recovering("QFrame::midLineWidth")

	if ptr.Pointer() != nil {
		return int(C.QFrame_MidLineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) SetFrameRect(v core.QRect_ITF) {
	defer qt.Recovering("QFrame::setFrameRect")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameRect(ptr.Pointer(), core.PointerFromQRect(v))
	}
}

func (ptr *QFrame) SetFrameShadow(v QFrame__Shadow) {
	defer qt.Recovering("QFrame::setFrameShadow")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameShadow(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetFrameShape(v QFrame__Shape) {
	defer qt.Recovering("QFrame::setFrameShape")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameShape(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetLineWidth(v int) {
	defer qt.Recovering("QFrame::setLineWidth")

	if ptr.Pointer() != nil {
		C.QFrame_SetLineWidth(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetMidLineWidth(v int) {
	defer qt.Recovering("QFrame::setMidLineWidth")

	if ptr.Pointer() != nil {
		C.QFrame_SetMidLineWidth(ptr.Pointer(), C.int(v))
	}
}

func NewQFrame(parent QWidget_ITF, f core.Qt__WindowType) *QFrame {
	defer qt.Recovering("QFrame::QFrame")

	return NewQFrameFromPointer(C.QFrame_NewQFrame(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QFrame) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QFrame::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QFrame) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QFrame::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQFrameChangeEvent
func callbackQFrameChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQFrameFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QFrame) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QFrame::changeEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QFrame) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QFrame::changeEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QFrame) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QFrame::event")

	if ptr.Pointer() != nil {
		return C.QFrame_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QFrame) FrameStyle() int {
	defer qt.Recovering("QFrame::frameStyle")

	if ptr.Pointer() != nil {
		return int(C.QFrame_FrameStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QFrame) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQFramePaintEvent
func callbackQFramePaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQFrameFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QFrame) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QFrame::paintEvent")

	if ptr.Pointer() != nil {
		C.QFrame_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QFrame) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QFrame::paintEvent")

	if ptr.Pointer() != nil {
		C.QFrame_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QFrame) SetFrameStyle(style int) {
	defer qt.Recovering("QFrame::setFrameStyle")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QFrame) SizeHint() *core.QSize {
	defer qt.Recovering("QFrame::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QFrame_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFrame) DestroyQFrame() {
	defer qt.Recovering("QFrame::~QFrame")

	if ptr.Pointer() != nil {
		C.QFrame_DestroyQFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFrame) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QFrame::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QFrame) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QFrame::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQFrameActionEvent
func callbackQFrameActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QFrame) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QFrame::actionEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QFrame) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QFrame::actionEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QFrame) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QFrame::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QFrame) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QFrame::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQFrameDragEnterEvent
func callbackQFrameDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QFrame) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QFrame::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QFrame_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QFrame) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QFrame::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QFrame_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QFrame) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QFrame) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQFrameDragLeaveEvent
func callbackQFrameDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QFrame) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QFrame_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QFrame) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QFrame_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QFrame) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QFrame::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QFrame) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QFrame::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQFrameDragMoveEvent
func callbackQFrameDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QFrame) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QFrame::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QFrame_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QFrame) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QFrame::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QFrame_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QFrame) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QFrame::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QFrame) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QFrame::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQFrameDropEvent
func callbackQFrameDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QFrame) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QFrame::dropEvent")

	if ptr.Pointer() != nil {
		C.QFrame_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QFrame) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QFrame::dropEvent")

	if ptr.Pointer() != nil {
		C.QFrame_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QFrame) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFrame::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QFrame) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QFrame::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQFrameEnterEvent
func callbackQFrameEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFrame) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFrame::enterEvent")

	if ptr.Pointer() != nil {
		C.QFrame_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFrame) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFrame::enterEvent")

	if ptr.Pointer() != nil {
		C.QFrame_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFrame) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFrame::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QFrame) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QFrame::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQFrameFocusInEvent
func callbackQFrameFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QFrame) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFrame::focusInEvent")

	if ptr.Pointer() != nil {
		C.QFrame_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QFrame) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFrame::focusInEvent")

	if ptr.Pointer() != nil {
		C.QFrame_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QFrame) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFrame::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QFrame) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QFrame::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQFrameFocusOutEvent
func callbackQFrameFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QFrame) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFrame::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QFrame_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QFrame) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QFrame::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QFrame_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QFrame) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QFrame::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QFrame) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QFrame::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQFrameHideEvent
func callbackQFrameHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QFrame) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QFrame::hideEvent")

	if ptr.Pointer() != nil {
		C.QFrame_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QFrame) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QFrame::hideEvent")

	if ptr.Pointer() != nil {
		C.QFrame_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QFrame) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFrame::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QFrame) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QFrame::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQFrameLeaveEvent
func callbackQFrameLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFrame) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFrame::leaveEvent")

	if ptr.Pointer() != nil {
		C.QFrame_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFrame) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFrame::leaveEvent")

	if ptr.Pointer() != nil {
		C.QFrame_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFrame) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QFrame::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QFrame) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QFrame::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQFrameMoveEvent
func callbackQFrameMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QFrame) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QFrame::moveEvent")

	if ptr.Pointer() != nil {
		C.QFrame_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QFrame) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QFrame::moveEvent")

	if ptr.Pointer() != nil {
		C.QFrame_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QFrame) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QFrame::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QFrame) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QFrame::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQFrameSetVisible
func callbackQFrameSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QFrame::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QFrame) SetVisible(visible bool) {
	defer qt.Recovering("QFrame::setVisible")

	if ptr.Pointer() != nil {
		C.QFrame_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFrame) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QFrame::setVisible")

	if ptr.Pointer() != nil {
		C.QFrame_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFrame) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QFrame::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QFrame) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QFrame::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQFrameShowEvent
func callbackQFrameShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QFrame) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QFrame::showEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QFrame) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QFrame::showEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QFrame) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QFrame::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QFrame) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QFrame::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQFrameCloseEvent
func callbackQFrameCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QFrame) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QFrame::closeEvent")

	if ptr.Pointer() != nil {
		C.QFrame_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QFrame) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QFrame::closeEvent")

	if ptr.Pointer() != nil {
		C.QFrame_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QFrame) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QFrame::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QFrame) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QFrame::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQFrameContextMenuEvent
func callbackQFrameContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QFrame) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QFrame::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QFrame) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QFrame::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QFrame) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QFrame::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QFrame) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QFrame::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQFrameInitPainter
func callbackQFrameInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQFrameFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QFrame) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QFrame::initPainter")

	if ptr.Pointer() != nil {
		C.QFrame_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QFrame) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QFrame::initPainter")

	if ptr.Pointer() != nil {
		C.QFrame_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QFrame) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QFrame::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QFrame) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QFrame::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQFrameInputMethodEvent
func callbackQFrameInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QFrame) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QFrame::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QFrame_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QFrame) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QFrame::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QFrame_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QFrame) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFrame::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QFrame) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QFrame::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQFrameKeyPressEvent
func callbackQFrameKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QFrame) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFrame::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QFrame_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QFrame) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFrame::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QFrame_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QFrame) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QFrame) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQFrameKeyReleaseEvent
func callbackQFrameKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QFrame) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFrame_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QFrame) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFrame_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QFrame) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QFrame) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQFrameMouseDoubleClickEvent
func callbackQFrameMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFrame) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QFrame_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFrame) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QFrame_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFrame) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QFrame) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQFrameMouseMoveEvent
func callbackQFrameMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFrame) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QFrame_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFrame) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QFrame_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFrame) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFrame::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QFrame) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QFrame::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQFrameMousePressEvent
func callbackQFrameMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFrame) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFrame::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QFrame_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFrame) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFrame::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QFrame_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFrame) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QFrame) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQFrameMouseReleaseEvent
func callbackQFrameMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QFrame) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFrame_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFrame) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QFrame_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QFrame) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QFrame::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QFrame) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QFrame::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQFrameResizeEvent
func callbackQFrameResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QFrame) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QFrame::resizeEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QFrame) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QFrame::resizeEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QFrame) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QFrame::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QFrame) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QFrame::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQFrameTabletEvent
func callbackQFrameTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QFrame) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QFrame::tabletEvent")

	if ptr.Pointer() != nil {
		C.QFrame_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QFrame) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QFrame::tabletEvent")

	if ptr.Pointer() != nil {
		C.QFrame_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QFrame) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QFrame::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QFrame) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QFrame::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQFrameWheelEvent
func callbackQFrameWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QFrame) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QFrame::wheelEvent")

	if ptr.Pointer() != nil {
		C.QFrame_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QFrame) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QFrame::wheelEvent")

	if ptr.Pointer() != nil {
		C.QFrame_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QFrame) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFrame) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFrameTimerEvent
func callbackQFrameTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QFrame) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QFrame::timerEvent")

	if ptr.Pointer() != nil {
		C.QFrame_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QFrame) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QFrame::timerEvent")

	if ptr.Pointer() != nil {
		C.QFrame_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QFrame) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFrame) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFrameChildEvent
func callbackQFrameChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QFrame) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QFrame::childEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QFrame) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QFrame::childEvent")

	if ptr.Pointer() != nil {
		C.QFrame_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QFrame) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFrame) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFrameCustomEvent
func callbackQFrameCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFrame::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQFrameFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QFrame) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QFrame::customEvent")

	if ptr.Pointer() != nil {
		C.QFrame_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QFrame) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QFrame::customEvent")

	if ptr.Pointer() != nil {
		C.QFrame_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
