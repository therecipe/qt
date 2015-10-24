package gui

//#include "qopenglfunctions_4_3_core.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_3_Core struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_3_CoreITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_4_3_CorePTR() *QOpenGLFunctions_4_3_Core
}

func PointerFromQOpenGLFunctions_4_3_Core(ptr QOpenGLFunctions_4_3_CoreITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_3_CorePTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_4_3_CoreFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_3_Core {
	var n = new(QOpenGLFunctions_4_3_Core)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_3_Core) QOpenGLFunctions_4_3_CorePTR() *QOpenGLFunctions_4_3_Core {
	return ptr
}

func NewQOpenGLFunctions_4_3_Core() *QOpenGLFunctions_4_3_Core {
	return QOpenGLFunctions_4_3_CoreFromPointer(unsafe.Pointer(C.QOpenGLFunctions_4_3_Core_NewQOpenGLFunctions_4_3_Core()))
}

func (ptr *QOpenGLFunctions_4_3_Core) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_4_3_Core_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_4_3_Core) DestroyQOpenGLFunctions_4_3_Core() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_3_Core_DestroyQOpenGLFunctions_4_3_Core(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_3_Core) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_3_Core_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_3_Core) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_3_Core_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_3_Core) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_3_Core_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_3_Core) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_3_Core_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_3_Core) GlPauseTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_3_Core_GlPauseTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_3_Core) GlReleaseShaderCompiler() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_3_Core_GlReleaseShaderCompiler(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_3_Core) GlResumeTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_3_Core_GlResumeTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}
