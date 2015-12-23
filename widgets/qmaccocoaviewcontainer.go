package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMacCocoaViewContainer struct {
	QWidget
}

type QMacCocoaViewContainer_ITF interface {
	QWidget_ITF
	QMacCocoaViewContainer_PTR() *QMacCocoaViewContainer
}

func PointerFromQMacCocoaViewContainer(ptr QMacCocoaViewContainer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacCocoaViewContainer_PTR().Pointer()
	}
	return nil
}

func NewQMacCocoaViewContainerFromPointer(ptr unsafe.Pointer) *QMacCocoaViewContainer {
	var n = new(QMacCocoaViewContainer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMacCocoaViewContainer_") {
		n.SetObjectName("QMacCocoaViewContainer_" + qt.Identifier())
	}
	return n
}

func (ptr *QMacCocoaViewContainer) QMacCocoaViewContainer_PTR() *QMacCocoaViewContainer {
	return ptr
}

func (ptr *QMacCocoaViewContainer) DestroyQMacCocoaViewContainer() {
	defer qt.Recovering("QMacCocoaViewContainer::~QMacCocoaViewContainer")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DestroyQMacCocoaViewContainer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacCocoaViewContainer) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQMacCocoaViewContainerActionEvent
func callbackQMacCocoaViewContainerActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQMacCocoaViewContainerDragEnterEvent
func callbackQMacCocoaViewContainerDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQMacCocoaViewContainerDragLeaveEvent
func callbackQMacCocoaViewContainerDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQMacCocoaViewContainerDragMoveEvent
func callbackQMacCocoaViewContainerDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQMacCocoaViewContainerDropEvent
func callbackQMacCocoaViewContainerDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQMacCocoaViewContainerEnterEvent
func callbackQMacCocoaViewContainerEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQMacCocoaViewContainerFocusOutEvent
func callbackQMacCocoaViewContainerFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQMacCocoaViewContainerHideEvent
func callbackQMacCocoaViewContainerHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQMacCocoaViewContainerLeaveEvent
func callbackQMacCocoaViewContainerLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQMacCocoaViewContainerMoveEvent
func callbackQMacCocoaViewContainerMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQMacCocoaViewContainerPaintEvent
func callbackQMacCocoaViewContainerPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQMacCocoaViewContainerSetVisible
func callbackQMacCocoaViewContainerSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQMacCocoaViewContainerShowEvent
func callbackQMacCocoaViewContainerShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQMacCocoaViewContainerCloseEvent
func callbackQMacCocoaViewContainerCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQMacCocoaViewContainerContextMenuEvent
func callbackQMacCocoaViewContainerContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQMacCocoaViewContainerInitPainter
func callbackQMacCocoaViewContainerInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQMacCocoaViewContainerInputMethodEvent
func callbackQMacCocoaViewContainerInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQMacCocoaViewContainerKeyPressEvent
func callbackQMacCocoaViewContainerKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQMacCocoaViewContainerKeyReleaseEvent
func callbackQMacCocoaViewContainerKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQMacCocoaViewContainerMouseDoubleClickEvent
func callbackQMacCocoaViewContainerMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQMacCocoaViewContainerMouseMoveEvent
func callbackQMacCocoaViewContainerMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQMacCocoaViewContainerMousePressEvent
func callbackQMacCocoaViewContainerMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQMacCocoaViewContainerMouseReleaseEvent
func callbackQMacCocoaViewContainerMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQMacCocoaViewContainerResizeEvent
func callbackQMacCocoaViewContainerResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQMacCocoaViewContainerTabletEvent
func callbackQMacCocoaViewContainerTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQMacCocoaViewContainerWheelEvent
func callbackQMacCocoaViewContainerWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMacCocoaViewContainerTimerEvent
func callbackQMacCocoaViewContainerTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMacCocoaViewContainerChildEvent
func callbackQMacCocoaViewContainerChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMacCocoaViewContainer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacCocoaViewContainer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMacCocoaViewContainer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMacCocoaViewContainer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMacCocoaViewContainerCustomEvent
func callbackQMacCocoaViewContainerCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMacCocoaViewContainer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
