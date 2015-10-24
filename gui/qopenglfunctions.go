package gui

//#include "qopenglfunctions.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions struct {
	ptr unsafe.Pointer
}

type QOpenGLFunctionsITF interface {
	QOpenGLFunctionsPTR() *QOpenGLFunctions
}

func (p *QOpenGLFunctions) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLFunctions) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLFunctions(ptr QOpenGLFunctionsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctionsPTR().Pointer()
	}
	return nil
}

func QOpenGLFunctionsFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions {
	var n = new(QOpenGLFunctions)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions) QOpenGLFunctionsPTR() *QOpenGLFunctions {
	return ptr
}

//QOpenGLFunctions::OpenGLFeature
type QOpenGLFunctions__OpenGLFeature int

var (
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

func NewQOpenGLFunctions() *QOpenGLFunctions {
	return QOpenGLFunctionsFromPointer(unsafe.Pointer(C.QOpenGLFunctions_NewQOpenGLFunctions()))
}

func NewQOpenGLFunctions2(context QOpenGLContextITF) *QOpenGLFunctions {
	return QOpenGLFunctionsFromPointer(unsafe.Pointer(C.QOpenGLFunctions_NewQOpenGLFunctions2(C.QtObjectPtr(PointerFromQOpenGLContext(context)))))
}

func (ptr *QOpenGLFunctions) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions) GlReleaseShaderCompiler() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_GlReleaseShaderCompiler(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions) HasOpenGLFeature(feature QOpenGLFunctions__OpenGLFeature) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_HasOpenGLFeature(C.QtObjectPtr(ptr.Pointer()), C.int(feature)) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions) InitializeOpenGLFunctions() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions) OpenGLFeatures() QOpenGLFunctions__OpenGLFeature {
	if ptr.Pointer() != nil {
		return QOpenGLFunctions__OpenGLFeature(C.QOpenGLFunctions_OpenGLFeatures(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLFunctions) DestroyQOpenGLFunctions() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_DestroyQOpenGLFunctions(C.QtObjectPtr(ptr.Pointer()))
	}
}
