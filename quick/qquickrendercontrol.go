package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QQuickRenderControl struct {
	core.QObject
}

type QQuickRenderControl_ITF interface {
	core.QObject_ITF
	QQuickRenderControl_PTR() *QQuickRenderControl
}

func PointerFromQQuickRenderControl(ptr QQuickRenderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickRenderControl_PTR().Pointer()
	}
	return nil
}

func NewQQuickRenderControlFromPointer(ptr unsafe.Pointer) *QQuickRenderControl {
	var n = new(QQuickRenderControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickRenderControl_") {
		n.SetObjectName("QQuickRenderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickRenderControl) QQuickRenderControl_PTR() *QQuickRenderControl {
	return ptr
}

func NewQQuickRenderControl(parent core.QObject_ITF) *QQuickRenderControl {
	defer qt.Recovering("QQuickRenderControl::QQuickRenderControl")

	return NewQQuickRenderControlFromPointer(C.QQuickRenderControl_NewQQuickRenderControl(core.PointerFromQObject(parent)))
}

func (ptr *QQuickRenderControl) Initialize(gl gui.QOpenGLContext_ITF) {
	defer qt.Recovering("QQuickRenderControl::initialize")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Initialize(ptr.Pointer(), gui.PointerFromQOpenGLContext(gl))
	}
}

func (ptr *QQuickRenderControl) Invalidate() {
	defer qt.Recovering("QQuickRenderControl::invalidate")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Invalidate(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) PolishItems() {
	defer qt.Recovering("QQuickRenderControl::polishItems")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PolishItems(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) PrepareThread(targetThread core.QThread_ITF) {
	defer qt.Recovering("QQuickRenderControl::prepareThread")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PrepareThread(ptr.Pointer(), core.PointerFromQThread(targetThread))
	}
}

func (ptr *QQuickRenderControl) Render() {
	defer qt.Recovering("QQuickRenderControl::render")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Render(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) ConnectRenderRequested(f func()) {
	defer qt.Recovering("connect QQuickRenderControl::renderRequested")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectRenderRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "renderRequested", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectRenderRequested() {
	defer qt.Recovering("disconnect QQuickRenderControl::renderRequested")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectRenderRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "renderRequested")
	}
}

//export callbackQQuickRenderControlRenderRequested
func callbackQQuickRenderControlRenderRequested(ptrName *C.char) {
	defer qt.Recovering("callback QQuickRenderControl::renderRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "renderRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickRenderControl) RenderWindow(offset core.QPoint_ITF) *gui.QWindow {
	defer qt.Recovering("QQuickRenderControl::renderWindow")

	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QQuickRenderControl_RenderWindow(ptr.Pointer(), core.PointerFromQPoint(offset)))
	}
	return nil
}

func QQuickRenderControl_RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {
	defer qt.Recovering("QQuickRenderControl::renderWindowFor")

	return gui.NewQWindowFromPointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(PointerFromQQuickWindow(win), core.PointerFromQPoint(offset)))
}

func (ptr *QQuickRenderControl) ConnectSceneChanged(f func()) {
	defer qt.Recovering("connect QQuickRenderControl::sceneChanged")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectSceneChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneChanged", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectSceneChanged() {
	defer qt.Recovering("disconnect QQuickRenderControl::sceneChanged")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectSceneChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneChanged")
	}
}

//export callbackQQuickRenderControlSceneChanged
func callbackQQuickRenderControlSceneChanged(ptrName *C.char) {
	defer qt.Recovering("callback QQuickRenderControl::sceneChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickRenderControl) Sync() bool {
	defer qt.Recovering("QQuickRenderControl::sync")

	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_Sync(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) DestroyQQuickRenderControl() {
	defer qt.Recovering("QQuickRenderControl::~QQuickRenderControl")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DestroyQQuickRenderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickRenderControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickRenderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickRenderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickRenderControlTimerEvent
func callbackQQuickRenderControlTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickRenderControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickRenderControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickRenderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickRenderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickRenderControlChildEvent
func callbackQQuickRenderControlChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickRenderControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickRenderControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickRenderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickRenderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickRenderControlCustomEvent
func callbackQQuickRenderControlCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickRenderControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
