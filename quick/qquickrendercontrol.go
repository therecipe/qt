package quick

//#include "qquickrendercontrol.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickRenderControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickRenderControl) QQuickRenderControl_PTR() *QQuickRenderControl {
	return ptr
}

func NewQQuickRenderControl(parent core.QObject_ITF) *QQuickRenderControl {
	return NewQQuickRenderControlFromPointer(C.QQuickRenderControl_NewQQuickRenderControl(core.PointerFromQObject(parent)))
}

func (ptr *QQuickRenderControl) Initialize(gl gui.QOpenGLContext_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Initialize(ptr.Pointer(), gui.PointerFromQOpenGLContext(gl))
	}
}

func (ptr *QQuickRenderControl) Invalidate() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Invalidate(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) PolishItems() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PolishItems(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) PrepareThread(targetThread core.QThread_ITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PrepareThread(ptr.Pointer(), core.PointerFromQThread(targetThread))
	}
}

func (ptr *QQuickRenderControl) Render() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Render(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) ConnectRenderRequested(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectRenderRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "renderRequested", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectRenderRequested() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectRenderRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "renderRequested")
	}
}

//export callbackQQuickRenderControlRenderRequested
func callbackQQuickRenderControlRenderRequested(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "renderRequested").(func())()
}

func (ptr *QQuickRenderControl) RenderWindow(offset core.QPoint_ITF) *gui.QWindow {
	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QQuickRenderControl_RenderWindow(ptr.Pointer(), core.PointerFromQPoint(offset)))
	}
	return nil
}

func QQuickRenderControl_RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {
	return gui.NewQWindowFromPointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(PointerFromQQuickWindow(win), core.PointerFromQPoint(offset)))
}

func (ptr *QQuickRenderControl) ConnectSceneChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectSceneChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneChanged", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectSceneChanged() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectSceneChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneChanged")
	}
}

//export callbackQQuickRenderControlSceneChanged
func callbackQQuickRenderControlSceneChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneChanged").(func())()
}

func (ptr *QQuickRenderControl) Sync() bool {
	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_Sync(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) DestroyQQuickRenderControl() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DestroyQQuickRenderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
