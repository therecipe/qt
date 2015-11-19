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

type QSGEngine_ITF interface {
	core.QObject_ITF
	QSGEngine_PTR() *QSGEngine
}

func PointerFromQSGEngine(ptr QSGEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGEngine_PTR().Pointer()
	}
	return nil
}

func NewQSGEngineFromPointer(ptr unsafe.Pointer) *QSGEngine {
	var n = new(QSGEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSGEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSGEngine) QSGEngine_PTR() *QSGEngine {
	return ptr
}

//QSGEngine::CreateTextureOption
type QSGEngine__CreateTextureOption int64

const (
	QSGEngine__TextureHasAlphaChannel = QSGEngine__CreateTextureOption(0x0001)
	QSGEngine__TextureOwnsGLTexture   = QSGEngine__CreateTextureOption(0x0004)
	QSGEngine__TextureCanUseAtlas     = QSGEngine__CreateTextureOption(0x0008)
)

func NewQSGEngine(parent core.QObject_ITF) *QSGEngine {
	return NewQSGEngineFromPointer(C.QSGEngine_NewQSGEngine(core.PointerFromQObject(parent)))
}

func (ptr *QSGEngine) CreateRenderer() *QSGAbstractRenderer {
	if ptr.Pointer() != nil {
		return NewQSGAbstractRendererFromPointer(C.QSGEngine_CreateRenderer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGEngine) CreateTextureFromImage(image gui.QImage_ITF, options QSGEngine__CreateTextureOption) *QSGTexture {
	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGEngine_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.int(options)))
	}
	return nil
}

func (ptr *QSGEngine) Initialize(context gui.QOpenGLContext_ITF) {
	if ptr.Pointer() != nil {
		C.QSGEngine_Initialize(ptr.Pointer(), gui.PointerFromQOpenGLContext(context))
	}
}

func (ptr *QSGEngine) Invalidate() {
	if ptr.Pointer() != nil {
		C.QSGEngine_Invalidate(ptr.Pointer())
	}
}

func (ptr *QSGEngine) DestroyQSGEngine() {
	if ptr.Pointer() != nil {
		C.QSGEngine_DestroyQSGEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
