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
func callbackQQuickWidgetFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetDragEnterEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
		return true
	}
	return false

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
func callbackQQuickWidgetDragLeaveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
		return true
	}
	return false

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
func callbackQQuickWidgetDragMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
		return true
	}
	return false

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
func callbackQQuickWidgetDropEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QQuickWidget) Engine() *qml.QQmlEngine {
	defer qt.Recovering("QQuickWidget::engine")

	if ptr.Pointer() != nil {
		return qml.NewQQmlEngineFromPointer(C.QQuickWidget_Engine(ptr.Pointer()))
	}
	return nil
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
func callbackQQuickWidgetHideEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
		return true
	}
	return false

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
func callbackQQuickWidgetKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

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
func callbackQQuickWidgetKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

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
func callbackQQuickWidgetMouseDoubleClickEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQQuickWidgetMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQQuickWidgetMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQQuickWidgetMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQQuickWidgetSceneGraphError(ptrName *C.char, error C.int, message *C.char) {
	defer qt.Recovering("callback QQuickWidget::sceneGraphError")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
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
func callbackQQuickWidgetShowEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
		return true
	}
	return false

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
func callbackQQuickWidgetStatusChanged(ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQuickWidget::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQuickWidget__Status))(QQuickWidget__Status(status))
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
func callbackQQuickWidgetWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

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
func callbackQQuickWidgetActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QQuickWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQQuickWidgetChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQQuickWidgetInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickWidgetCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
