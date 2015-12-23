package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"unsafe"
)

type QQuickView struct {
	QQuickWindow
}

type QQuickView_ITF interface {
	QQuickWindow_ITF
	QQuickView_PTR() *QQuickView
}

func PointerFromQQuickView(ptr QQuickView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickView_PTR().Pointer()
	}
	return nil
}

func NewQQuickViewFromPointer(ptr unsafe.Pointer) *QQuickView {
	var n = new(QQuickView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickView_") {
		n.SetObjectName("QQuickView_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickView) QQuickView_PTR() *QQuickView {
	return ptr
}

//QQuickView::ResizeMode
type QQuickView__ResizeMode int64

const (
	QQuickView__SizeViewToRootObject = QQuickView__ResizeMode(0)
	QQuickView__SizeRootObjectToView = QQuickView__ResizeMode(1)
)

//QQuickView::Status
type QQuickView__Status int64

const (
	QQuickView__Null    = QQuickView__Status(0)
	QQuickView__Ready   = QQuickView__Status(1)
	QQuickView__Loading = QQuickView__Status(2)
	QQuickView__Error   = QQuickView__Status(3)
)

func (ptr *QQuickView) ResizeMode() QQuickView__ResizeMode {
	defer qt.Recovering("QQuickView::resizeMode")

	if ptr.Pointer() != nil {
		return QQuickView__ResizeMode(C.QQuickView_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickView) SetResizeMode(v QQuickView__ResizeMode) {
	defer qt.Recovering("QQuickView::setResizeMode")

	if ptr.Pointer() != nil {
		C.QQuickView_SetResizeMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QQuickView) Status() QQuickView__Status {
	defer qt.Recovering("QQuickView::status")

	if ptr.Pointer() != nil {
		return QQuickView__Status(C.QQuickView_Status(ptr.Pointer()))
	}
	return 0
}

func NewQQuickView2(engine qml.QQmlEngine_ITF, parent gui.QWindow_ITF) *QQuickView {
	defer qt.Recovering("QQuickView::QQuickView")

	return NewQQuickViewFromPointer(C.QQuickView_NewQQuickView2(qml.PointerFromQQmlEngine(engine), gui.PointerFromQWindow(parent)))
}

func NewQQuickView(parent gui.QWindow_ITF) *QQuickView {
	defer qt.Recovering("QQuickView::QQuickView")

	return NewQQuickViewFromPointer(C.QQuickView_NewQQuickView(gui.PointerFromQWindow(parent)))
}

func NewQQuickView3(source core.QUrl_ITF, parent gui.QWindow_ITF) *QQuickView {
	defer qt.Recovering("QQuickView::QQuickView")

	return NewQQuickViewFromPointer(C.QQuickView_NewQQuickView3(core.PointerFromQUrl(source), gui.PointerFromQWindow(parent)))
}

func (ptr *QQuickView) Engine() *qml.QQmlEngine {
	defer qt.Recovering("QQuickView::engine")

	if ptr.Pointer() != nil {
		return qml.NewQQmlEngineFromPointer(C.QQuickView_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) InitialSize() *core.QSize {
	defer qt.Recovering("QQuickView::initialSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickView_InitialSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QQuickView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QQuickView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQQuickViewKeyPressEvent
func callbackQQuickViewKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QQuickView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QQuickView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQQuickViewKeyReleaseEvent
func callbackQQuickViewKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QQuickView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQQuickViewMouseMoveEvent
func callbackQQuickViewMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QQuickView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQQuickViewMousePressEvent
func callbackQQuickViewMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QQuickView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQQuickViewMouseReleaseEvent
func callbackQQuickViewMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QQuickView) RootContext() *qml.QQmlContext {
	defer qt.Recovering("QQuickView::rootContext")

	if ptr.Pointer() != nil {
		return qml.NewQQmlContextFromPointer(C.QQuickView_RootContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) RootObject() *QQuickItem {
	defer qt.Recovering("QQuickView::rootObject")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickView_RootObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) SetSource(url core.QUrl_ITF) {
	defer qt.Recovering("QQuickView::setSource")

	if ptr.Pointer() != nil {
		C.QQuickView_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQuickView) Source() *core.QUrl {
	defer qt.Recovering("QQuickView::source")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQuickView_Source(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) ConnectStatusChanged(f func(status QQuickView__Status)) {
	defer qt.Recovering("connect QQuickView::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickView_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QQuickView) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QQuickView::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickView_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQQuickViewStatusChanged
func callbackQQuickViewStatusChanged(ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQuickView::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQuickView__Status))(QQuickView__Status(status))
	}

}

func (ptr *QQuickView) DestroyQQuickView() {
	defer qt.Recovering("QQuickView::~QQuickView")

	if ptr.Pointer() != nil {
		C.QQuickView_DestroyQQuickView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickView) ConnectExposeEvent(f func(v *gui.QExposeEvent)) {
	defer qt.Recovering("connect QQuickView::exposeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "exposeEvent", f)
	}
}

func (ptr *QQuickView) DisconnectExposeEvent() {
	defer qt.Recovering("disconnect QQuickView::exposeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "exposeEvent")
	}
}

//export callbackQQuickViewExposeEvent
func callbackQQuickViewExposeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::exposeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposeEvent"); signal != nil {
		signal.(func(*gui.QExposeEvent))(gui.NewQExposeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectFocusInEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QQuickView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QQuickView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQQuickViewFocusInEvent
func callbackQQuickViewFocusInEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QQuickView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QQuickView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQQuickViewFocusOutEvent
func callbackQQuickViewFocusOutEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectHideEvent(f func(v *gui.QHideEvent)) {
	defer qt.Recovering("connect QQuickView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QQuickView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QQuickView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQQuickViewHideEvent
func callbackQQuickViewHideEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QQuickView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQQuickViewMouseDoubleClickEvent
func callbackQQuickViewMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectResizeEvent(f func(ev *gui.QResizeEvent)) {
	defer qt.Recovering("connect QQuickView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QQuickView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QQuickView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQQuickViewResizeEvent
func callbackQQuickViewResizeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QQuickView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QQuickView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QQuickView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQQuickViewShowEvent
func callbackQQuickViewShowEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QQuickView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QQuickView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QQuickView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQQuickViewWheelEvent
func callbackQQuickViewWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectMoveEvent(f func(ev *gui.QMoveEvent)) {
	defer qt.Recovering("connect QQuickView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QQuickView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQQuickViewMoveEvent
func callbackQQuickViewMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectTabletEvent(f func(ev *gui.QTabletEvent)) {
	defer qt.Recovering("connect QQuickView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QQuickView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QQuickView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQQuickViewTabletEvent
func callbackQQuickViewTabletEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectTouchEvent(f func(ev *gui.QTouchEvent)) {
	defer qt.Recovering("connect QQuickView::touchEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchEvent", f)
	}
}

func (ptr *QQuickView) DisconnectTouchEvent() {
	defer qt.Recovering("disconnect QQuickView::touchEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchEvent")
	}
}

//export callbackQQuickViewTouchEvent
func callbackQQuickViewTouchEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickViewTimerEvent
func callbackQQuickViewTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickViewChildEvent
func callbackQQuickViewChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickViewCustomEvent
func callbackQQuickViewCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
