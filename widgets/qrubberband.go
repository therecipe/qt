package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QRubberBand struct {
	QWidget
}

type QRubberBand_ITF interface {
	QWidget_ITF
	QRubberBand_PTR() *QRubberBand
}

func PointerFromQRubberBand(ptr QRubberBand_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRubberBand_PTR().Pointer()
	}
	return nil
}

func NewQRubberBandFromPointer(ptr unsafe.Pointer) *QRubberBand {
	var n = new(QRubberBand)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRubberBand_") {
		n.SetObjectName("QRubberBand_" + qt.Identifier())
	}
	return n
}

func (ptr *QRubberBand) QRubberBand_PTR() *QRubberBand {
	return ptr
}

//QRubberBand::Shape
type QRubberBand__Shape int64

const (
	QRubberBand__Line      = QRubberBand__Shape(0)
	QRubberBand__Rectangle = QRubberBand__Shape(1)
)

func (ptr *QRubberBand) SetGeometry(rect core.QRect_ITF) {
	defer qt.Recovering("QRubberBand::setGeometry")

	if ptr.Pointer() != nil {
		C.QRubberBand_SetGeometry(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func NewQRubberBand(s QRubberBand__Shape, p QWidget_ITF) *QRubberBand {
	defer qt.Recovering("QRubberBand::QRubberBand")

	return NewQRubberBandFromPointer(C.QRubberBand_NewQRubberBand(C.int(s), PointerFromQWidget(p)))
}

func (ptr *QRubberBand) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QRubberBand::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QRubberBand::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQRubberBandChangeEvent
func callbackQRubberBandChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRubberBand) Move2(p core.QPoint_ITF) {
	defer qt.Recovering("QRubberBand::move")

	if ptr.Pointer() != nil {
		C.QRubberBand_Move2(ptr.Pointer(), core.PointerFromQPoint(p))
	}
}

func (ptr *QRubberBand) Move(x int, y int) {
	defer qt.Recovering("QRubberBand::move")

	if ptr.Pointer() != nil {
		C.QRubberBand_Move(ptr.Pointer(), C.int(x), C.int(y))
	}
}

func (ptr *QRubberBand) ConnectMoveEvent(f func(v *gui.QMoveEvent)) {
	defer qt.Recovering("connect QRubberBand::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QRubberBand::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQRubberBandMoveEvent
func callbackQRubberBandMoveEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QRubberBand::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QRubberBand::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQRubberBandPaintEvent
func callbackQRubberBandPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QRubberBand) Resize2(size core.QSize_ITF) {
	defer qt.Recovering("QRubberBand::resize")

	if ptr.Pointer() != nil {
		C.QRubberBand_Resize2(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QRubberBand) Resize(width int, height int) {
	defer qt.Recovering("QRubberBand::resize")

	if ptr.Pointer() != nil {
		C.QRubberBand_Resize(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QRubberBand) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QRubberBand::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QRubberBand::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQRubberBandResizeEvent
func callbackQRubberBandResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QRubberBand) SetGeometry2(x int, y int, width int, height int) {
	defer qt.Recovering("QRubberBand::setGeometry")

	if ptr.Pointer() != nil {
		C.QRubberBand_SetGeometry2(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QRubberBand) Shape() QRubberBand__Shape {
	defer qt.Recovering("QRubberBand::shape")

	if ptr.Pointer() != nil {
		return QRubberBand__Shape(C.QRubberBand_Shape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRubberBand) ConnectShowEvent(f func(e *gui.QShowEvent)) {
	defer qt.Recovering("connect QRubberBand::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QRubberBand::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQRubberBandShowEvent
func callbackQRubberBandShowEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRubberBand) DestroyQRubberBand() {
	defer qt.Recovering("QRubberBand::~QRubberBand")

	if ptr.Pointer() != nil {
		C.QRubberBand_DestroyQRubberBand(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRubberBand) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QRubberBand::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QRubberBand::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQRubberBandActionEvent
func callbackQRubberBandActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QRubberBand::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QRubberBand::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQRubberBandDragEnterEvent
func callbackQRubberBandDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QRubberBand::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QRubberBand::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQRubberBandDragLeaveEvent
func callbackQRubberBandDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QRubberBand::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QRubberBand::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQRubberBandDragMoveEvent
func callbackQRubberBandDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QRubberBand::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QRubberBand::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQRubberBandDropEvent
func callbackQRubberBandDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRubberBand::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QRubberBand::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQRubberBandEnterEvent
func callbackQRubberBandEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QRubberBand::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QRubberBand::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQRubberBandFocusInEvent
func callbackQRubberBandFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QRubberBand::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QRubberBand::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQRubberBandFocusOutEvent
func callbackQRubberBandFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QRubberBand::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QRubberBand::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQRubberBandHideEvent
func callbackQRubberBandHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRubberBand::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QRubberBand::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQRubberBandLeaveEvent
func callbackQRubberBandLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QRubberBand::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QRubberBand) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QRubberBand::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQRubberBandSetVisible
func callbackQRubberBandSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QRubberBand::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QRubberBand::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QRubberBand::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQRubberBandCloseEvent
func callbackQRubberBandCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QRubberBand::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QRubberBand::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQRubberBandContextMenuEvent
func callbackQRubberBandContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QRubberBand::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QRubberBand) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QRubberBand::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQRubberBandInitPainter
func callbackQRubberBandInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QRubberBand::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QRubberBand::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQRubberBandInputMethodEvent
func callbackQRubberBandInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QRubberBand::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QRubberBand::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQRubberBandKeyPressEvent
func callbackQRubberBandKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QRubberBand::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QRubberBand::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQRubberBandKeyReleaseEvent
func callbackQRubberBandKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRubberBand::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QRubberBand::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQRubberBandMouseDoubleClickEvent
func callbackQRubberBandMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRubberBand::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QRubberBand::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQRubberBandMouseMoveEvent
func callbackQRubberBandMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRubberBand::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QRubberBand::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQRubberBandMousePressEvent
func callbackQRubberBandMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRubberBand::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QRubberBand::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQRubberBandMouseReleaseEvent
func callbackQRubberBandMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QRubberBand::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QRubberBand::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQRubberBandTabletEvent
func callbackQRubberBandTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QRubberBand::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QRubberBand::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQRubberBandWheelEvent
func callbackQRubberBandWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRubberBand::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRubberBand::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRubberBandTimerEvent
func callbackQRubberBandTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRubberBand::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRubberBand::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRubberBandChildEvent
func callbackQRubberBandChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRubberBand) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRubberBand::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRubberBand) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRubberBand::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRubberBandCustomEvent
func callbackQRubberBandCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRubberBand::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
