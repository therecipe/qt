package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSplitterHandle struct {
	QWidget
}

type QSplitterHandle_ITF interface {
	QWidget_ITF
	QSplitterHandle_PTR() *QSplitterHandle
}

func PointerFromQSplitterHandle(ptr QSplitterHandle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplitterHandle_PTR().Pointer()
	}
	return nil
}

func NewQSplitterHandleFromPointer(ptr unsafe.Pointer) *QSplitterHandle {
	var n = new(QSplitterHandle)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSplitterHandle_") {
		n.SetObjectName("QSplitterHandle_" + qt.Identifier())
	}
	return n
}

func (ptr *QSplitterHandle) QSplitterHandle_PTR() *QSplitterHandle {
	return ptr
}

func NewQSplitterHandle(orientation core.Qt__Orientation, parent QSplitter_ITF) *QSplitterHandle {
	defer qt.Recovering("QSplitterHandle::QSplitterHandle")

	return NewQSplitterHandleFromPointer(C.QSplitterHandle_NewQSplitterHandle(C.int(orientation), PointerFromQSplitter(parent)))
}

func (ptr *QSplitterHandle) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QSplitterHandle::event")

	if ptr.Pointer() != nil {
		return C.QSplitterHandle_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSplitterHandle) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitterHandle::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSplitterHandleMouseMoveEvent
func callbackQSplitterHandleMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQSplitterHandleFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QSplitterHandle) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QSplitterHandle) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QSplitterHandle) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitterHandle::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSplitterHandleMousePressEvent
func callbackQSplitterHandleMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQSplitterHandleFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QSplitterHandle) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QSplitterHandle) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QSplitterHandle) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitterHandle::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSplitterHandleMouseReleaseEvent
func callbackQSplitterHandleMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQSplitterHandleFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QSplitterHandle) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QSplitterHandle) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QSplitterHandle) OpaqueResize() bool {
	defer qt.Recovering("QSplitterHandle::opaqueResize")

	if ptr.Pointer() != nil {
		return C.QSplitterHandle_OpaqueResize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSplitterHandle) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QSplitterHandle::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSplitterHandle_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitterHandle) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSplitterHandle::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSplitterHandlePaintEvent
func callbackQSplitterHandlePaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQSplitterHandleFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QSplitterHandle) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::paintEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QSplitterHandle) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::paintEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QSplitterHandle) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSplitterHandle::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSplitterHandleResizeEvent
func callbackQSplitterHandleResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSplitterHandle) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSplitterHandle) SetOrientation(orientation core.Qt__Orientation) {
	defer qt.Recovering("QSplitterHandle::setOrientation")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QSplitterHandle) SizeHint() *core.QSize {
	defer qt.Recovering("QSplitterHandle::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSplitterHandle_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitterHandle) Splitter() *QSplitter {
	defer qt.Recovering("QSplitterHandle::splitter")

	if ptr.Pointer() != nil {
		return NewQSplitterFromPointer(C.QSplitterHandle_Splitter(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitterHandle) DestroyQSplitterHandle() {
	defer qt.Recovering("QSplitterHandle::~QSplitterHandle")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DestroyQSplitterHandle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSplitterHandle) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QSplitterHandle::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQSplitterHandleActionEvent
func callbackQSplitterHandleActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::actionEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSplitterHandle) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::actionEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QSplitterHandle::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQSplitterHandleDragEnterEvent
func callbackQSplitterHandleDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSplitterHandle) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QSplitterHandle::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQSplitterHandleDragLeaveEvent
func callbackQSplitterHandleDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSplitterHandle) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QSplitterHandle::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQSplitterHandleDragMoveEvent
func callbackQSplitterHandleDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSplitterHandle) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QSplitterHandle::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQSplitterHandleDropEvent
func callbackQSplitterHandleDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::dropEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSplitterHandle) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::dropEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplitterHandle::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQSplitterHandleEnterEvent
func callbackQSplitterHandleEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::enterEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitterHandle) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::enterEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSplitterHandle::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQSplitterHandleFocusInEvent
func callbackQSplitterHandleFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplitterHandle) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSplitterHandle::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQSplitterHandleFocusOutEvent
func callbackQSplitterHandleFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplitterHandle) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QSplitterHandle::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQSplitterHandleHideEvent
func callbackQSplitterHandleHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::hideEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSplitterHandle) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::hideEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplitterHandle::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQSplitterHandleLeaveEvent
func callbackQSplitterHandleLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitterHandle) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QSplitterHandle::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQSplitterHandleMoveEvent
func callbackQSplitterHandleMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::moveEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSplitterHandle) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::moveEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QSplitterHandle::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QSplitterHandle) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QSplitterHandle::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQSplitterHandleSetVisible
func callbackQSplitterHandleSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSplitterHandle::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QSplitterHandle) SetVisible(visible bool) {
	defer qt.Recovering("QSplitterHandle::setVisible")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSplitterHandle) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QSplitterHandle::setVisible")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSplitterHandle) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QSplitterHandle::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQSplitterHandleShowEvent
func callbackQSplitterHandleShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::showEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSplitterHandle) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::showEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplitterHandle::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQSplitterHandleChangeEvent
func callbackQSplitterHandleChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::changeEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitterHandle) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::changeEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QSplitterHandle::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQSplitterHandleCloseEvent
func callbackQSplitterHandleCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::closeEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSplitterHandle) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::closeEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QSplitterHandle::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQSplitterHandleContextMenuEvent
func callbackQSplitterHandleContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSplitterHandle) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSplitterHandle::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QSplitterHandle) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QSplitterHandle::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQSplitterHandleInitPainter
func callbackQSplitterHandleInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQSplitterHandleFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QSplitterHandle) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSplitterHandle::initPainter")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSplitterHandle) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSplitterHandle::initPainter")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSplitterHandle) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QSplitterHandle::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQSplitterHandleInputMethodEvent
func callbackQSplitterHandleInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSplitterHandle) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSplitterHandle::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQSplitterHandleKeyPressEvent
func callbackQSplitterHandleKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplitterHandle) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSplitterHandle::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQSplitterHandleKeyReleaseEvent
func callbackQSplitterHandleKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplitterHandle) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitterHandle::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQSplitterHandleMouseDoubleClickEvent
func callbackQSplitterHandleMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplitterHandle) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QSplitterHandle::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQSplitterHandleTabletEvent
func callbackQSplitterHandleTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSplitterHandle) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QSplitterHandle::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQSplitterHandleWheelEvent
func callbackQSplitterHandleWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSplitterHandle) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSplitterHandle::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSplitterHandleTimerEvent
func callbackQSplitterHandleTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::timerEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSplitterHandle) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::timerEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSplitterHandle::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSplitterHandleChildEvent
func callbackQSplitterHandleChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::childEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSplitterHandle) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::childEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSplitterHandle) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSplitterHandle::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSplitterHandleCustomEvent
func callbackQSplitterHandleCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSplitterHandle::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSplitterHandleFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSplitterHandle) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::customEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSplitterHandle) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSplitterHandle::customEvent")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
