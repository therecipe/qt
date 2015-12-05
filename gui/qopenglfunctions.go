package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions struct {
	ptr unsafe.Pointer
}

type QOpenGLFunctions_ITF interface {
	QOpenGLFunctions_PTR() *QOpenGLFunctions
}

func (p *QOpenGLFunctions) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLFunctions) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLFunctions(ptr QOpenGLFunctions_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctionsFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions {
	var n = new(QOpenGLFunctions)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions) QOpenGLFunctions_PTR() *QOpenGLFunctions {
	return ptr
}

//QOpenGLFunctions::OpenGLFeature
type QOpenGLFunctions__OpenGLFeature int64

const (
	QOpenGLFunctions__Multitexture          = QOpenGLFunctions__OpenGLFeature(0x0001)
	QOpenGLFunctions__Shaders               = QOpenGLFunctions__OpenGLFeature(0x0002)
	QOpenGLFunctions__Buffers               = QOpenGLFunctions__OpenGLFeature(0x0004)
	QOpenGLFunctions__Framebuffers          = QOpenGLFunctions__OpenGLFeature(0x0008)
	QOpenGLFunctions__BlendColor            = QOpenGLFunctions__OpenGLFeature(0x0010)
	QOpenGLFunctions__BlendEquation         = QOpenGLFunctions__OpenGLFeature(0x0020)
	QOpenGLFunctions__BlendEquationSeparate = QOpenGLFunctions__OpenGLFeature(0x0040)
	QOpenGLFunctions__BlendFuncSeparate     = QOpenGLFunctions__OpenGLFeature(0x0080)
	QOpenGLFunctions__BlendSubtract         = QOpenGLFunctions__OpenGLFeature(0x0100)
	QOpenGLFunctions__CompressedTextures    = QOpenGLFunctions__OpenGLFeature(0x0200)
	QOpenGLFunctions__Multisample           = QOpenGLFunctions__OpenGLFeature(0x0400)
	QOpenGLFunctions__StencilSeparate       = QOpenGLFunctions__OpenGLFeature(0x0800)
	QOpenGLFunctions__NPOTTextures          = QOpenGLFunctions__OpenGLFeature(0x1000)
	QOpenGLFunctions__NPOTTextureRepeat     = QOpenGLFunctions__OpenGLFeature(0x2000)
	QOpenGLFunctions__FixedFunctionPipeline = QOpenGLFunctions__OpenGLFeature(0x4000)
	QOpenGLFunctions__TextureRGFormats      = QOpenGLFunctions__OpenGLFeature(0x8000)
)
