package gui

//#include "qopenglfunctions_4_0_compatibility.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_0_Compatibility struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_0_CompatibilityITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_4_0_CompatibilityPTR() *QOpenGLFunctions_4_0_Compatibility
}

func PointerFromQOpenGLFunctions_4_0_Compatibility(ptr QOpenGLFunctions_4_0_CompatibilityITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_0_CompatibilityPTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_4_0_CompatibilityFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_0_Compatibility {
	var n = new(QOpenGLFunctions_4_0_Compatibility)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) QOpenGLFunctions_4_0_CompatibilityPTR() *QOpenGLFunctions_4_0_Compatibility {
	return ptr
}

func NewQOpenGLFunctions_4_0_Compatibility() *QOpenGLFunctions_4_0_Compatibility {
	return QOpenGLFunctions_4_0_CompatibilityFromPointer(unsafe.Pointer(C.QOpenGLFunctions_4_0_Compatibility_NewQOpenGLFunctions_4_0_Compatibility()))
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_4_0_Compatibility_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) DestroyQOpenGLFunctions_4_0_Compatibility() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_DestroyQOpenGLFunctions_4_0_Compatibility(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlPauseTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlPauseTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlPopClientAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlPopClientAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) GlResumeTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_0_Compatibility_GlResumeTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}
