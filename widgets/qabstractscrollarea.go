package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QAbstractScrollArea struct {
	QFrame
}

type QAbstractScrollArea_ITF interface {
	QFrame_ITF
	QAbstractScrollArea_PTR() *QAbstractScrollArea
}

func PointerFromQAbstractScrollArea(ptr QAbstractScrollArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func NewQAbstractScrollAreaFromPointer(ptr unsafe.Pointer) *QAbstractScrollArea {
	var n = new(QAbstractScrollArea)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractScrollArea_") {
		n.SetObjectName("QAbstractScrollArea_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractScrollArea) QAbstractScrollArea_PTR() *QAbstractScrollArea {
	return ptr
}

//QAbstractScrollArea::SizeAdjustPolicy
type QAbstractScrollArea__SizeAdjustPolicy int64

const (
	QAbstractScrollArea__AdjustIgnored               = QAbstractScrollArea__SizeAdjustPolicy(0)
	QAbstractScrollArea__AdjustToContentsOnFirstShow = QAbstractScrollArea__SizeAdjustPolicy(1)
	QAbstractScrollArea__AdjustToContents            = QAbstractScrollArea__SizeAdjustPolicy(2)
)

func (ptr *QAbstractScrollArea) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQAbstractScrollAreaDragEnterEvent
func callbackQAbstractScrollAreaDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::dragEnterEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "dragEnterEvent")
	if signal != nil {
		defer signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQAbstractScrollAreaDragLeaveEvent
func callbackQAbstractScrollAreaDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::dragLeaveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent")
	if signal != nil {
		defer signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQAbstractScrollAreaDragMoveEvent
func callbackQAbstractScrollAreaDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::dragMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "dragMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQAbstractScrollAreaDropEvent
func callbackQAbstractScrollAreaDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::dropEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "dropEvent")
	if signal != nil {
		defer signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) HorizontalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	defer qt.Recovering("QAbstractScrollArea::horizontalScrollBarPolicy")

	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_HorizontalScrollBarPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractScrollArea) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQAbstractScrollAreaPaintEvent
func callbackQAbstractScrollAreaPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQAbstractScrollAreaResizeEvent
func callbackQAbstractScrollAreaResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBarPolicy(v core.Qt__ScrollBarPolicy) {
	defer qt.Recovering("QAbstractScrollArea::setHorizontalScrollBarPolicy")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBarPolicy(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractScrollArea) SetSizeAdjustPolicy(policy QAbstractScrollArea__SizeAdjustPolicy) {
	defer qt.Recovering("QAbstractScrollArea::setSizeAdjustPolicy")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetSizeAdjustPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBarPolicy(v core.Qt__ScrollBarPolicy) {
	defer qt.Recovering("QAbstractScrollArea::setVerticalScrollBarPolicy")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBarPolicy(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractScrollArea) SizeAdjustPolicy() QAbstractScrollArea__SizeAdjustPolicy {
	defer qt.Recovering("QAbstractScrollArea::sizeAdjustPolicy")

	if ptr.Pointer() != nil {
		return QAbstractScrollArea__SizeAdjustPolicy(C.QAbstractScrollArea_SizeAdjustPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractScrollArea) VerticalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	defer qt.Recovering("QAbstractScrollArea::verticalScrollBarPolicy")

	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_VerticalScrollBarPolicy(ptr.Pointer()))
	}
	return 0
}

func NewQAbstractScrollArea(parent QWidget_ITF) *QAbstractScrollArea {
	defer qt.Recovering("QAbstractScrollArea::QAbstractScrollArea")

	return NewQAbstractScrollAreaFromPointer(C.QAbstractScrollArea_NewQAbstractScrollArea(PointerFromQWidget(parent)))
}

func (ptr *QAbstractScrollArea) AddScrollBarWidget(widget QWidget_ITF, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QAbstractScrollArea::addScrollBarWidget")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_AddScrollBarWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(alignment))
	}
}

func (ptr *QAbstractScrollArea) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQAbstractScrollAreaContextMenuEvent
func callbackQAbstractScrollAreaContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::contextMenuEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "contextMenuEvent")
	if signal != nil {
		defer signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) CornerWidget() *QWidget {
	defer qt.Recovering("QAbstractScrollArea::cornerWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractScrollArea_CornerWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	defer qt.Recovering("QAbstractScrollArea::horizontalScrollBar")

	if ptr.Pointer() != nil {
		return NewQScrollBarFromPointer(C.QAbstractScrollArea_HorizontalScrollBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQAbstractScrollAreaKeyPressEvent
func callbackQAbstractScrollAreaKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQAbstractScrollAreaMouseDoubleClickEvent
func callbackQAbstractScrollAreaMouseDoubleClickEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::mouseDoubleClickEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQAbstractScrollAreaMouseMoveEvent
func callbackQAbstractScrollAreaMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::mouseMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQAbstractScrollAreaMousePressEvent
func callbackQAbstractScrollAreaMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQAbstractScrollAreaMouseReleaseEvent
func callbackQAbstractScrollAreaMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QAbstractScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QAbstractScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQAbstractScrollAreaScrollContentsBy
func callbackQAbstractScrollAreaScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QAbstractScrollArea::scrollContentsBy")

	var signal = qt.GetSignal(C.GoString(ptrName), "scrollContentsBy")
	if signal != nil {
		defer signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) SetCornerWidget(widget QWidget_ITF) {
	defer qt.Recovering("QAbstractScrollArea::setCornerWidget")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetCornerWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBar(scrollBar QScrollBar_ITF) {
	defer qt.Recovering("QAbstractScrollArea::setHorizontalScrollBar")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBar(ptr.Pointer(), PointerFromQScrollBar(scrollBar))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBar(scrollBar QScrollBar_ITF) {
	defer qt.Recovering("QAbstractScrollArea::setVerticalScrollBar")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBar(ptr.Pointer(), PointerFromQScrollBar(scrollBar))
	}
}

func (ptr *QAbstractScrollArea) SetViewport(widget QWidget_ITF) {
	defer qt.Recovering("QAbstractScrollArea::setViewport")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetViewport(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractScrollArea) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QAbstractScrollArea::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QAbstractScrollArea::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQAbstractScrollAreaSetupViewport
func callbackQAbstractScrollAreaSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::setupViewport")

	var signal = qt.GetSignal(C.GoString(ptrName), "setupViewport")
	if signal != nil {
		defer signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	defer qt.Recovering("QAbstractScrollArea::verticalScrollBar")

	if ptr.Pointer() != nil {
		return NewQScrollBarFromPointer(C.QAbstractScrollArea_VerticalScrollBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) Viewport() *QWidget {
	defer qt.Recovering("QAbstractScrollArea::viewport")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractScrollArea_Viewport(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQAbstractScrollAreaWheelEvent
func callbackQAbstractScrollAreaWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::wheelEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "wheelEvent")
	if signal != nil {
		defer signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) DestroyQAbstractScrollArea() {
	defer qt.Recovering("QAbstractScrollArea::~QAbstractScrollArea")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DestroyQAbstractScrollArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
