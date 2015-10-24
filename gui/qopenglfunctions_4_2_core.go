package gui

//#include "qopenglfunctions_4_2_core.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_2_Core struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_2_CoreITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_4_2_CorePTR() *QOpenGLFunctions_4_2_Core
}

func PointerFromQOpenGLFunctions_4_2_Core(ptr QOpenGLFunctions_4_2_CoreITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_2_CorePTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_4_2_CoreFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_2_Core {
	var n = new(QOpenGLFunctions_4_2_Core)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_2_Core) QOpenGLFunctions_4_2_CorePTR() *QOpenGLFunctions_4_2_Core {
	return ptr
}

func NewQOpenGLFunctions_4_2_Core() *QOpenGLFunctions_4_2_Core {
	return QOpenGLFunctions_4_2_CoreFromPointer(unsafe.Pointer(C.QOpenGLFunctions_4_2_Core_NewQOpenGLFunctions_4_2_Core()))
}

func (ptr *QOpenGLFunctions_4_2_Core) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_4_2_Core_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_4_2_Core) DestroyQOpenGLFunctions_4_2_Core() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Core_DestroyQOpenGLFunctions_4_2_Core(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Core) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Core_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Core) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Core_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Core) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Core_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Core) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Core_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Core) GlPauseTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Core_GlPauseTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Core) GlReleaseShaderCompiler() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Core_GlReleaseShaderCompiler(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Core) GlResumeTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Core_GlResumeTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}
