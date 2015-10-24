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

type QQuickRenderControlITF interface {
	core.QObjectITF
	QQuickRenderControlPTR() *QQuickRenderControl
}

func PointerFromQQuickRenderControl(ptr QQuickRenderControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickRenderControlPTR().Pointer()
	}
	return nil
}

func QQuickRenderControlFromPointer(ptr unsafe.Pointer) *QQuickRenderControl {
	var n = new(QQuickRenderControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickRenderControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickRenderControl) QQuickRenderControlPTR() *QQuickRenderControl {
	return ptr
}

func NewQQuickRenderControl(parent core.QObjectITF) *QQuickRenderControl {
	return QQuickRenderControlFromPointer(unsafe.Pointer(C.QQuickRenderControl_NewQQuickRenderControl(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QQuickRenderControl) Initialize(gl gui.QOpenGLContextITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Initialize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQOpenGLContext(gl)))
	}
}

func (ptr *QQuickRenderControl) Invalidate() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickRenderControl) PolishItems() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PolishItems(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickRenderControl) PrepareThread(targetThread core.QThreadITF) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PrepareThread(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQThread(targetThread)))
	}
}

func (ptr *QQuickRenderControl) Render() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Render(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickRenderControl) ConnectRenderRequested(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectRenderRequested(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "renderRequested", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectRenderRequested() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectRenderRequested(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "renderRequested")
	}
}

//export callbackQQuickRenderControlRenderRequested
func callbackQQuickRenderControlRenderRequested(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "renderRequested").(func())()
}

func (ptr *QQuickRenderControl) RenderWindow(offset core.QPointITF) *gui.QWindow {
	if ptr.Pointer() != nil {
		return gui.QWindowFromPointer(unsafe.Pointer(C.QQuickRenderControl_RenderWindow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(offset)))))
	}
	return nil
}

func QQuickRenderControl_RenderWindowFor(win QQuickWindowITF, offset core.QPointITF) *gui.QWindow {
	return gui.QWindowFromPointer(unsafe.Pointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(C.QtObjectPtr(PointerFromQQuickWindow(win)), C.QtObjectPtr(core.PointerFromQPoint(offset)))))
}

func (ptr *QQuickRenderControl) ConnectSceneChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectSceneChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sceneChanged", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectSceneChanged() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectSceneChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sceneChanged")
	}
}

//export callbackQQuickRenderControlSceneChanged
func callbackQQuickRenderControlSceneChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneChanged").(func())()
}

func (ptr *QQuickRenderControl) Sync() bool {
	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_Sync(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) DestroyQQuickRenderControl() {
	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DestroyQQuickRenderControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
