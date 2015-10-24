package quick

//#include "qsgengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSGEngine struct {
	core.QObject
}

type QSGEngineITF interface {
	core.QObjectITF
	QSGEnginePTR() *QSGEngine
}

func PointerFromQSGEngine(ptr QSGEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGEnginePTR().Pointer()
	}
	return nil
}

func QSGEngineFromPointer(ptr unsafe.Pointer) *QSGEngine {
	var n = new(QSGEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSGEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSGEngine) QSGEnginePTR() *QSGEngine {
	return ptr
}

//QSGEngine::CreateTextureOption
type QSGEngine__CreateTextureOption int

var (
	QSGEngine__TextureHasAlphaChannel = QSGEngine__CreateTextureOption(0x0001)
	QSGEngine__TextureOwnsGLTexture   = QSGEngine__CreateTextureOption(0x0004)
	QSGEngine__TextureCanUseAtlas     = QSGEngine__CreateTextureOption(0x0008)
)

func NewQSGEngine(parent core.QObjectITF) *QSGEngine {
	return QSGEngineFromPointer(unsafe.Pointer(C.QSGEngine_NewQSGEngine(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QSGEngine) CreateRenderer() *QSGAbstractRenderer {
	if ptr.Pointer() != nil {
		return QSGAbstractRendererFromPointer(unsafe.Pointer(C.QSGEngine_CreateRenderer(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGEngine) CreateTextureFromImage(image gui.QImageITF, options QSGEngine__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		return QSGTextureFromPointer(unsafe.Pointer(C.QSGEngine_CreateTextureFromImage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQImage(image)), C.int(options))))
	}
	return nil
}

func (ptr *QSGEngine) Initialize(context gui.QOpenGLContextITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_Initialize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQOpenGLContext(context)))
	}
}

func (ptr *QSGEngine) Invalidate() {
	if ptr.Pointer() != nil {
		C.QSGEngine_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGEngine) DestroyQSGEngine() {
	if ptr.Pointer() != nil {
		C.QSGEngine_DestroyQSGEngine(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
