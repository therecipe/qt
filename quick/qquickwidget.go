package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QQuickWidget struct {
	widgets.QWidget
}

type QQuickWidget_ITF interface {
	widgets.QWidget_ITF
	QQuickWidget_PTR() *QQuickWidget
}

func PointerFromQQuickWidget(ptr QQuickWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWidget_PTR().Pointer()
	}
	return nil
}

func NewQQuickWidgetFromPointer(ptr unsafe.Pointer) *QQuickWidget {
	var n = new(QQuickWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickWidget_") {
		n.SetObjectName("QQuickWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickWidget) QQuickWidget_PTR() *QQuickWidget {
	return ptr
}

//QQuickWidget::ResizeMode
type QQuickWidget__ResizeMode int64

const (
	QQuickWidget__SizeViewToRootObject = QQuickWidget__ResizeMode(0)
	QQuickWidget__SizeRootObjectToView = QQuickWidget__ResizeMode(1)
)

//QQuickWidget::Status
type QQuickWidget__Status int64

const (
	QQuickWidget__Null    = QQuickWidget__Status(0)
	QQuickWidget__Ready   = QQuickWidget__Status(1)
	QQuickWidget__Loading = QQuickWidget__Status(2)
	QQuickWidget__Error   = QQuickWidget__Status(3)
)

func (ptr *QQuickWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QQuickWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQQuickWidgetFocusInEvent
func callbackQQuickWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QQuickWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQQuickWidgetFocusOutEvent
func callbackQQuickWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) ResizeMode() QQuickWidget__ResizeMode {
	defer qt.Recovering("QQuickWidget::resizeMode")

	if ptr.Pointer() != nil {
		return QQuickWidget__ResizeMode(C.QQuickWidget_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWidget) SetResizeMode(v QQuickWidget__ResizeMode) {
	defer qt.Recovering("QQuickWidget::setResizeMode")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetResizeMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QQuickWidget) Status() QQuickWidget__Status {
	defer qt.Recovering("QQuickWidget::status")

	if ptr.Pointer() != nil {
		return QQuickWidget__Status(C.QQuickWidget_Status(ptr.Pointer()))
	}
	return 0
}

func NewQQuickWidget2(engine qml.QQmlEngine_ITF, parent widgets.QWidget_ITF) *QQuickWidget {
	defer qt.Recovering("QQuickWidget::QQuickWidget")

	return NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget2(qml.PointerFromQQmlEngine(engine), widgets.PointerFromQWidget(parent)))
}

func NewQQuickWidget(parent widgets.QWidget_ITF) *QQuickWidget {
	defer qt.Recovering("QQuickWidget::QQuickWidget")

	return NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget(widgets.PointerFromQWidget(parent)))
}

func NewQQuickWidget3(source core.QUrl_ITF, parent widgets.QWidget_ITF) *QQuickWidget {
	defer qt.Recovering("QQuickWidget::QQuickWidget")

	return NewQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget3(core.PointerFromQUrl(source), widgets.PointerFromQWidget(parent)))
}

func (ptr *QQuickWidget) ConnectDragEnterEvent(f func(e *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QQuickWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QQuickWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQQuickWidgetDragEnterEvent
func callbackQQuickWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DragEnterEvent(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QQuickWidget) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QQuickWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QQuickWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQQuickWidgetDragLeaveEvent
func callbackQQuickWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QQuickWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QQuickWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QQuickWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQQuickWidgetDragMoveEvent
func callbackQQuickWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QQuickWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QQuickWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QQuickWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQQuickWidgetDropEvent
func callbackQQuickWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QQuickWidget) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QQuickWidget) Engine() *qml.QQmlEngine {
	defer qt.Recovering("QQuickWidget::engine")

	if ptr.Pointer() != nil {
		return qml.NewQQmlEngineFromPointer(C.QQuickWidget_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWidget::event")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWidget) ConnectHideEvent(f func(v *gui.QHideEvent)) {
	defer qt.Recovering("connect QQuickWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QQuickWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQQuickWidgetHideEvent
func callbackQQuickWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
	} else {
		NewQQuickWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(v))
	}
}

func (ptr *QQuickWidget) HideEvent(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QQuickWidget) HideEventDefault(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QQuickWidget) InitialSize() *core.QSize {
	defer qt.Recovering("QQuickWidget::initialSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWidget_InitialSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QQuickWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQQuickWidgetKeyPressEvent
func callbackQQuickWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWidget) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QQuickWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQQuickWidgetKeyReleaseEvent
func callbackQQuickWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWidget) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QQuickWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQQuickWidgetMouseDoubleClickEvent
func callbackQQuickWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QQuickWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQQuickWidgetMouseMoveEvent
func callbackQQuickWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QQuickWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQQuickWidgetMousePressEvent
func callbackQQuickWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QQuickWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQQuickWidgetMouseReleaseEvent
func callbackQQuickWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) QuickWindow() *QQuickWindow {
	defer qt.Recovering("QQuickWidget::quickWindow")

	if ptr.Pointer() != nil {
		return NewQQuickWindowFromPointer(C.QQuickWidget_QuickWindow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) RootContext() *qml.QQmlContext {
	defer qt.Recovering("QQuickWidget::rootContext")

	if ptr.Pointer() != nil {
		return qml.NewQQmlContextFromPointer(C.QQuickWidget_RootContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) RootObject() *QQuickItem {
	defer qt.Recovering("QQuickWidget::rootObject")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWidget_RootObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	defer qt.Recovering("connect QQuickWidget::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectSceneGraphError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphError", f)
	}
}

func (ptr *QQuickWidget) DisconnectSceneGraphError() {
	defer qt.Recovering("disconnect QQuickWidget::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphError")
	}
}

//export callbackQQuickWidgetSceneGraphError
func callbackQQuickWidgetSceneGraphError(ptr unsafe.Pointer, ptrName *C.char, error C.int, message *C.char) {
	defer qt.Recovering("callback QQuickWidget::sceneGraphError")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
	}

}

func (ptr *QQuickWidget) SceneGraphError(error QQuickWindow__SceneGraphError, message string) {
	defer qt.Recovering("QQuickWidget::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SceneGraphError(ptr.Pointer(), C.int(error), C.CString(message))
	}
}

func (ptr *QQuickWidget) SetClearColor(color gui.QColor_ITF) {
	defer qt.Recovering("QQuickWidget::setClearColor")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetClearColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QQuickWidget) SetFormat(format gui.QSurfaceFormat_ITF) {
	defer qt.Recovering("QQuickWidget::setFormat")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetFormat(ptr.Pointer(), gui.PointerFromQSurfaceFormat(format))
	}
}

func (ptr *QQuickWidget) SetSource(url core.QUrl_ITF) {
	defer qt.Recovering("QQuickWidget::setSource")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQuickWidget) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QQuickWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QQuickWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQQuickWidgetShowEvent
func callbackQQuickWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(v))
	}
}

func (ptr *QQuickWidget) ShowEvent(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QQuickWidget) ShowEventDefault(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QQuickWidget) ConnectStatusChanged(f func(status QQuickWidget__Status)) {
	defer qt.Recovering("connect QQuickWidget::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QQuickWidget) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QQuickWidget::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQQuickWidgetStatusChanged
func callbackQQuickWidgetStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQuickWidget::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQuickWidget__Status))(QQuickWidget__Status(status))
	}

}

func (ptr *QQuickWidget) StatusChanged(status QQuickWidget__Status) {
	defer qt.Recovering("QQuickWidget::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickWidget_StatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QQuickWidget) Source() *core.QUrl {
	defer qt.Recovering("QQuickWidget::source")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQuickWidget_Source(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QQuickWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QQuickWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQQuickWidgetWheelEvent
func callbackQQuickWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QQuickWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QQuickWidget) DestroyQQuickWidget() {
	defer qt.Recovering("QQuickWidget::~QQuickWidget")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DestroyQQuickWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QQuickWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QQuickWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQQuickWidgetActionEvent
func callbackQQuickWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QQuickWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QQuickWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QQuickWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QQuickWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQQuickWidgetEnterEvent
func callbackQQuickWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QQuickWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQQuickWidgetLeaveEvent
func callbackQQuickWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QQuickWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QQuickWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQQuickWidgetMoveEvent
func callbackQQuickWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QQuickWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QQuickWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QQuickWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQQuickWidgetPaintEvent
func callbackQQuickWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QQuickWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QQuickWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QQuickWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QQuickWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QQuickWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQQuickWidgetSetVisible
func callbackQQuickWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QQuickWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QQuickWidget) SetVisible(visible bool) {
	defer qt.Recovering("QQuickWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QQuickWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QQuickWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QQuickWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QQuickWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQQuickWidgetChangeEvent
func callbackQQuickWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QQuickWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QQuickWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQQuickWidgetCloseEvent
func callbackQQuickWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QQuickWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QQuickWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QQuickWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQQuickWidgetContextMenuEvent
func callbackQQuickWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QQuickWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QQuickWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QQuickWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QQuickWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QQuickWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QQuickWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQQuickWidgetInitPainter
func callbackQQuickWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQQuickWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QQuickWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QQuickWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QQuickWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QQuickWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QQuickWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QQuickWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QQuickWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QQuickWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QQuickWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQQuickWidgetInputMethodEvent
func callbackQQuickWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QQuickWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QQuickWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QQuickWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QQuickWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQQuickWidgetResizeEvent
func callbackQQuickWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QQuickWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QQuickWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QQuickWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQQuickWidgetTabletEvent
func callbackQQuickWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QQuickWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QQuickWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QQuickWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickWidgetTimerEvent
func callbackQQuickWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickWidgetChildEvent
func callbackQQuickWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickWidgetCustomEvent
func callbackQQuickWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
