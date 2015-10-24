package gui

//#include "qopenglfunctions_4_5_core.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_5_Core struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_5_CoreITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_4_5_CorePTR() *QOpenGLFunctions_4_5_Core
}

func PointerFromQOpenGLFunctions_4_5_Core(ptr QOpenGLFunctions_4_5_CoreITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_5_CorePTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_4_5_CoreFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_5_Core {
	var n = new(QOpenGLFunctions_4_5_Core)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_5_Core) QOpenGLFunctions_4_5_CorePTR() *QOpenGLFunctions_4_5_Core {
	return ptr
}

func NewQOpenGLFunctions_4_5_Core() *QOpenGLFunctions_4_5_Core {
	return QOpenGLFunctions_4_5_CoreFromPointer(unsafe.Pointer(C.QOpenGLFunctions_4_5_Core_NewQOpenGLFunctions_4_5_Core()))
}

func (ptr *QOpenGLFunctions_4_5_Core) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_4_5_Core_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_4_5_Core) DestroyQOpenGLFunctions_4_5_Core() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_5_Core_DestroyQOpenGLFunctions_4_5_Core(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_5_Core) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_5_Core_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_5_Core) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_5_Core_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_5_Core) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_5_Core_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_5_Core) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_5_Core_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_5_Core) GlPauseTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_5_Core_GlPauseTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_5_Core) GlPopDebugGroup() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_5_Core_GlPopDebugGroup(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_5_Core) GlReleaseShaderCompiler() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_5_Core_GlReleaseShaderCompiler(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_5_Core) GlResumeTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_5_Core_GlResumeTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_5_Core) GlTextureBarrier() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_5_Core_GlTextureBarrier(C.QtObjectPtr(ptr.Pointer()))
	}
}
