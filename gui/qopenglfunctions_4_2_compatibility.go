package gui

//#include "qopenglfunctions_4_2_compatibility.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_2_Compatibility struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_2_CompatibilityITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_4_2_CompatibilityPTR() *QOpenGLFunctions_4_2_Compatibility
}

func PointerFromQOpenGLFunctions_4_2_Compatibility(ptr QOpenGLFunctions_4_2_CompatibilityITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_2_CompatibilityPTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_4_2_CompatibilityFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_2_Compatibility {
	var n = new(QOpenGLFunctions_4_2_Compatibility)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) QOpenGLFunctions_4_2_CompatibilityPTR() *QOpenGLFunctions_4_2_Compatibility {
	return ptr
}

func NewQOpenGLFunctions_4_2_Compatibility() *QOpenGLFunctions_4_2_Compatibility {
	return QOpenGLFunctions_4_2_CompatibilityFromPointer(unsafe.Pointer(C.QOpenGLFunctions_4_2_Compatibility_NewQOpenGLFunctions_4_2_Compatibility()))
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_4_2_Compatibility_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) DestroyQOpenGLFunctions_4_2_Compatibility() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_DestroyQOpenGLFunctions_4_2_Compatibility(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlPauseTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlPauseTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlPopClientAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlPopClientAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlReleaseShaderCompiler() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlReleaseShaderCompiler(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_2_Compatibility) GlResumeTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_2_Compatibility_GlResumeTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}
