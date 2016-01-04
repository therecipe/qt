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
func callbackQQuickViewKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickView) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickView) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQQuickViewKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickView) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickView) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQQuickViewMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickView) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQQuickViewMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickView) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQQuickViewMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickView) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQQuickViewStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQuickView::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQuickView__Status))(QQuickView__Status(status))
	}

}

func (ptr *QQuickView) StatusChanged(status QQuickView__Status) {
	defer qt.Recovering("QQuickView::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickView_StatusChanged(ptr.Pointer(), C.int(status))
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
func callbackQQuickViewExposeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::exposeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposeEvent"); signal != nil {
		signal.(func(*gui.QExposeEvent))(gui.NewQExposeEventFromPointer(v))
	} else {
		NewQQuickViewFromPointer(ptr).ExposeEventDefault(gui.NewQExposeEventFromPointer(v))
	}
}

func (ptr *QQuickView) ExposeEvent(v gui.QExposeEvent_ITF) {
	defer qt.Recovering("QQuickView::exposeEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ExposeEvent(ptr.Pointer(), gui.PointerFromQExposeEvent(v))
	}
}

func (ptr *QQuickView) ExposeEventDefault(v gui.QExposeEvent_ITF) {
	defer qt.Recovering("QQuickView::exposeEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ExposeEventDefault(ptr.Pointer(), gui.PointerFromQExposeEvent(v))
	}
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
func callbackQQuickViewFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickView) FocusInEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickView) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
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
func callbackQQuickViewFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickView) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickView) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
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
func callbackQQuickViewHideEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
	} else {
		NewQQuickViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(v))
	}
}

func (ptr *QQuickView) HideEvent(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickView::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QQuickView) HideEventDefault(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickView::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
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
func callbackQQuickViewMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickView) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQQuickViewResizeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(ev))
	}
}

func (ptr *QQuickView) ResizeEvent(ev gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

func (ptr *QQuickView) ResizeEventDefault(ev gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
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
func callbackQQuickViewShowEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
	} else {
		NewQQuickViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(v))
	}
}

func (ptr *QQuickView) ShowEvent(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickView::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QQuickView) ShowEventDefault(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickView::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
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
func callbackQQuickViewWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickView) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickView) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
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
func callbackQQuickViewMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(ev))
	}
}

func (ptr *QQuickView) MoveEvent(ev gui.QMoveEvent_ITF) {
	defer qt.Recovering("QQuickView::moveEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
}

func (ptr *QQuickView) MoveEventDefault(ev gui.QMoveEvent_ITF) {
	defer qt.Recovering("QQuickView::moveEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
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
func callbackQQuickViewTabletEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(ev))
	}
}

func (ptr *QQuickView) TabletEvent(ev gui.QTabletEvent_ITF) {
	defer qt.Recovering("QQuickView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
}

func (ptr *QQuickView) TabletEventDefault(ev gui.QTabletEvent_ITF) {
	defer qt.Recovering("QQuickView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
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
func callbackQQuickViewTouchEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(ev))
	}
}

func (ptr *QQuickView) TouchEvent(ev gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickView::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
}

func (ptr *QQuickView) TouchEventDefault(ev gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickView::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
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
func callbackQQuickViewTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickView) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickView::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickView) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickView::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQQuickViewChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickView::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickView::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQQuickViewCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickView::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickView::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
