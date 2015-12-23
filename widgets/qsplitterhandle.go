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
func callbackQSplitterHandleMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQSplitterHandleMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQSplitterHandleMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQSplitterHandlePaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

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
func callbackQSplitterHandleResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSplitterHandle::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQSplitterHandleShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQSplitterHandleInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSplitterHandleCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
