package quick

//#include "qsgabstractrenderer.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSGAbstractRenderer struct {
	core.QObject
}

type QSGAbstractRendererITF interface {
	core.QObjectITF
	QSGAbstractRendererPTR() *QSGAbstractRenderer
}

func PointerFromQSGAbstractRenderer(ptr QSGAbstractRendererITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGAbstractRendererPTR().Pointer()
	}
	return nil
}

func QSGAbstractRendererFromPointer(ptr unsafe.Pointer) *QSGAbstractRenderer {
	var n = new(QSGAbstractRenderer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSGAbstractRenderer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSGAbstractRenderer) QSGAbstractRendererPTR() *QSGAbstractRenderer {
	return ptr
}

//QSGAbstractRenderer::ClearModeBit
type QSGAbstractRenderer__ClearModeBit int

var (
	QSGAbstractRenderer__ClearColorBuffer   = QSGAbstractRenderer__ClearModeBit(0x0001)
	QSGAbstractRenderer__ClearDepthBuffer   = QSGAbstractRenderer__ClearModeBit(0x0002)
	QSGAbstractRenderer__ClearStencilBuffer = QSGAbstractRenderer__ClearModeBit(0x0004)
)

func (ptr *QSGAbstractRenderer) ClearMode() QSGAbstractRenderer__ClearModeBit {
	if ptr.Pointer() != nil {
		return QSGAbstractRenderer__ClearModeBit(C.QSGAbstractRenderer_ClearMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGAbstractRenderer) ConnectSceneGraphChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ConnectSceneGraphChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphChanged", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectSceneGraphChanged() {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DisconnectSceneGraphChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphChanged")
	}
}

//export callbackQSGAbstractRendererSceneGraphChanged
func callbackQSGAbstractRendererSceneGraphChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sceneGraphChanged").(func())()
}

func (ptr *QSGAbstractRenderer) SetClearColor(color gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetClearColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(color)))
	}
}

func (ptr *QSGAbstractRenderer) SetClearMode(mode QSGAbstractRenderer__ClearModeBit) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetClearMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QSGAbstractRenderer) SetDeviceRect(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetDeviceRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QSGAbstractRenderer) SetDeviceRect2(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetDeviceRect2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QSGAbstractRenderer) SetProjectionMatrix(matrix gui.QMatrix4x4ITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetProjectionMatrix(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQMatrix4x4(matrix)))
	}
}

func (ptr *QSGAbstractRenderer) SetProjectionMatrixToRect(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetProjectionMatrixToRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QSGAbstractRenderer) SetViewportRect(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetViewportRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QSGAbstractRenderer) SetViewportRect2(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetViewportRect2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}
