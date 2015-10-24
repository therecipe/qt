package svg

//#include "qsvgrenderer.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSvgRenderer struct {
	core.QObject
}

type QSvgRendererITF interface {
	core.QObjectITF
	QSvgRendererPTR() *QSvgRenderer
}

func PointerFromQSvgRenderer(ptr QSvgRendererITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgRendererPTR().Pointer()
	}
	return nil
}

func QSvgRendererFromPointer(ptr unsafe.Pointer) *QSvgRenderer {
	var n = new(QSvgRenderer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSvgRenderer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSvgRenderer) QSvgRendererPTR() *QSvgRenderer {
	return ptr
}

func (ptr *QSvgRenderer) FramesPerSecond() int {
	if ptr.Pointer() != nil {
		return int(C.QSvgRenderer_FramesPerSecond(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSvgRenderer) SetFramesPerSecond(num int) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetFramesPerSecond(C.QtObjectPtr(ptr.Pointer()), C.int(num))
	}
}

func (ptr *QSvgRenderer) SetViewBox(viewbox core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetViewBox(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(viewbox)))
	}
}

func (ptr *QSvgRenderer) SetViewBox2(viewbox core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetViewBox2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(viewbox)))
	}
}

func NewQSvgRenderer(parent core.QObjectITF) *QSvgRenderer {
	return QSvgRendererFromPointer(unsafe.Pointer(C.QSvgRenderer_NewQSvgRenderer(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQSvgRenderer4(contents core.QXmlStreamReaderITF, parent core.QObjectITF) *QSvgRenderer {
	return QSvgRendererFromPointer(unsafe.Pointer(C.QSvgRenderer_NewQSvgRenderer4(C.QtObjectPtr(core.PointerFromQXmlStreamReader(contents)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQSvgRenderer3(contents core.QByteArrayITF, parent core.QObjectITF) *QSvgRenderer {
	return QSvgRendererFromPointer(unsafe.Pointer(C.QSvgRenderer_NewQSvgRenderer3(C.QtObjectPtr(core.PointerFromQByteArray(contents)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQSvgRenderer2(filename string, parent core.QObjectITF) *QSvgRenderer {
	return QSvgRendererFromPointer(unsafe.Pointer(C.QSvgRenderer_NewQSvgRenderer2(C.CString(filename), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QSvgRenderer) Animated() bool {
	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Animated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSvgRenderer) ElementExists(id string) bool {
	if ptr.Pointer() != nil {
		return C.QSvgRenderer_ElementExists(C.QtObjectPtr(ptr.Pointer()), C.CString(id)) != 0
	}
	return false
}

func (ptr *QSvgRenderer) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSvgRenderer_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSvgRenderer) Load3(contents core.QXmlStreamReaderITF) bool {
	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Load3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQXmlStreamReader(contents))) != 0
	}
	return false
}

func (ptr *QSvgRenderer) Load2(contents core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Load2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(contents))) != 0
	}
	return false
}

func (ptr *QSvgRenderer) Load(filename string) bool {
	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Load(C.QtObjectPtr(ptr.Pointer()), C.CString(filename)) != 0
	}
	return false
}

func (ptr *QSvgRenderer) Render(painter gui.QPainterITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)))
	}
}

func (ptr *QSvgRenderer) Render2(painter gui.QPainterITF, bounds core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQRectF(bounds)))
	}
}

func (ptr *QSvgRenderer) Render3(painter gui.QPainterITF, elementId string, bounds core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.CString(elementId), C.QtObjectPtr(core.PointerFromQRectF(bounds)))
	}
}

func (ptr *QSvgRenderer) ConnectRepaintNeeded(f func()) {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_ConnectRepaintNeeded(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "repaintNeeded", f)
	}
}

func (ptr *QSvgRenderer) DisconnectRepaintNeeded() {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_DisconnectRepaintNeeded(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "repaintNeeded")
	}
}

//export callbackQSvgRendererRepaintNeeded
func callbackQSvgRendererRepaintNeeded(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "repaintNeeded").(func())()
}

func (ptr *QSvgRenderer) DestroyQSvgRenderer() {
	if ptr.Pointer() != nil {
		C.QSvgRenderer_DestroyQSvgRenderer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
