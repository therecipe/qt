package multimediawidgets

//#include "multimediawidgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QCameraViewfinder struct {
	QVideoWidget
}

type QCameraViewfinder_ITF interface {
	QVideoWidget_ITF
	QCameraViewfinder_PTR() *QCameraViewfinder
}

func PointerFromQCameraViewfinder(ptr QCameraViewfinder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinder_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderFromPointer(ptr unsafe.Pointer) *QCameraViewfinder {
	var n = new(QCameraViewfinder)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraViewfinder_") {
		n.SetObjectName("QCameraViewfinder_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraViewfinder) QCameraViewfinder_PTR() *QCameraViewfinder {
	return ptr
}

func NewQCameraViewfinder(parent widgets.QWidget_ITF) *QCameraViewfinder {
	defer qt.Recovering("QCameraViewfinder::QCameraViewfinder")

	return NewQCameraViewfinderFromPointer(C.QCameraViewfinder_NewQCameraViewfinder(widgets.PointerFromQWidget(parent)))
}

func (ptr *QCameraViewfinder) MediaObject() *multimedia.QMediaObject {
	defer qt.Recovering("QCameraViewfinder::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QCameraViewfinder_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCameraViewfinder) DestroyQCameraViewfinder() {
	defer qt.Recovering("QCameraViewfinder::~QCameraViewfinder")

	if ptr.Pointer() != nil {
		C.QCameraViewfinder_DestroyQCameraViewfinder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraViewfinder) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQCameraViewfinderHideEvent
func callbackQCameraViewfinderHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQCameraViewfinderMoveEvent
func callbackQCameraViewfinderMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQCameraViewfinderPaintEvent
func callbackQCameraViewfinderPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQCameraViewfinderResizeEvent
func callbackQCameraViewfinderResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQCameraViewfinderShowEvent
func callbackQCameraViewfinderShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQCameraViewfinderActionEvent
func callbackQCameraViewfinderActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQCameraViewfinderDragEnterEvent
func callbackQCameraViewfinderDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQCameraViewfinderDragLeaveEvent
func callbackQCameraViewfinderDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQCameraViewfinderDragMoveEvent
func callbackQCameraViewfinderDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQCameraViewfinderDropEvent
func callbackQCameraViewfinderDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQCameraViewfinderEnterEvent
func callbackQCameraViewfinderEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQCameraViewfinderFocusInEvent
func callbackQCameraViewfinderFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQCameraViewfinderFocusOutEvent
func callbackQCameraViewfinderFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQCameraViewfinderLeaveEvent
func callbackQCameraViewfinderLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QCameraViewfinder::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QCameraViewfinder::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQCameraViewfinderSetVisible
func callbackQCameraViewfinderSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QCameraViewfinder::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQCameraViewfinderChangeEvent
func callbackQCameraViewfinderChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQCameraViewfinderCloseEvent
func callbackQCameraViewfinderCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQCameraViewfinderContextMenuEvent
func callbackQCameraViewfinderContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QCameraViewfinder::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QCameraViewfinder::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQCameraViewfinderInitPainter
func callbackQCameraViewfinderInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQCameraViewfinderInputMethodEvent
func callbackQCameraViewfinderInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQCameraViewfinderKeyPressEvent
func callbackQCameraViewfinderKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQCameraViewfinderKeyReleaseEvent
func callbackQCameraViewfinderKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQCameraViewfinderMouseDoubleClickEvent
func callbackQCameraViewfinderMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQCameraViewfinderMouseMoveEvent
func callbackQCameraViewfinderMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQCameraViewfinderMousePressEvent
func callbackQCameraViewfinderMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQCameraViewfinderMouseReleaseEvent
func callbackQCameraViewfinderMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQCameraViewfinderTabletEvent
func callbackQCameraViewfinderTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQCameraViewfinderWheelEvent
func callbackQCameraViewfinderWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraViewfinderTimerEvent
func callbackQCameraViewfinderTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraViewfinderChildEvent
func callbackQCameraViewfinderChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraViewfinder) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinder::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraViewfinder) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraViewfinder::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraViewfinderCustomEvent
func callbackQCameraViewfinderCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraViewfinder::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
