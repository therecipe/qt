package gui

//#include "qopenglfunctions_4_1_core.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_1_Core struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_1_CoreITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_4_1_CorePTR() *QOpenGLFunctions_4_1_Core
}

func PointerFromQOpenGLFunctions_4_1_Core(ptr QOpenGLFunctions_4_1_CoreITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_1_CorePTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_4_1_CoreFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_1_Core {
	var n = new(QOpenGLFunctions_4_1_Core)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_1_Core) QOpenGLFunctions_4_1_CorePTR() *QOpenGLFunctions_4_1_Core {
	return ptr
}

func NewQOpenGLFunctions_4_1_Core() *QOpenGLFunctions_4_1_Core {
	return QOpenGLFunctions_4_1_CoreFromPointer(unsafe.Pointer(C.QOpenGLFunctions_4_1_Core_NewQOpenGLFunctions_4_1_Core()))
}

func (ptr *QOpenGLFunctions_4_1_Core) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_4_1_Core_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_4_1_Core) DestroyQOpenGLFunctions_4_1_Core() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_1_Core_DestroyQOpenGLFunctions_4_1_Core(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_1_Core) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_1_Core_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_1_Core) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_1_Core_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_1_Core) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_1_Core_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_1_Core) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_1_Core_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_1_Core) GlPauseTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_1_Core_GlPauseTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_1_Core) GlReleaseShaderCompiler() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_1_Core_GlReleaseShaderCompiler(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_1_Core) GlResumeTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_1_Core_GlResumeTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}
