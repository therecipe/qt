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
