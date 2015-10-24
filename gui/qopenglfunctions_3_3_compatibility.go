package gui

//#include "qopenglfunctions_3_3_compatibility.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_3_3_Compatibility struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_3_3_CompatibilityITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_3_3_CompatibilityPTR() *QOpenGLFunctions_3_3_Compatibility
}

func PointerFromQOpenGLFunctions_3_3_Compatibility(ptr QOpenGLFunctions_3_3_CompatibilityITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_3_3_CompatibilityPTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_3_3_CompatibilityFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_3_3_Compatibility {
	var n = new(QOpenGLFunctions_3_3_Compatibility)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) QOpenGLFunctions_3_3_CompatibilityPTR() *QOpenGLFunctions_3_3_Compatibility {
	return ptr
}

func NewQOpenGLFunctions_3_3_Compatibility() *QOpenGLFunctions_3_3_Compatibility {
	return QOpenGLFunctions_3_3_CompatibilityFromPointer(unsafe.Pointer(C.QOpenGLFunctions_3_3_Compatibility_NewQOpenGLFunctions_3_3_Compatibility()))
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_3_3_Compatibility_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) DestroyQOpenGLFunctions_3_3_Compatibility() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_DestroyQOpenGLFunctions_3_3_Compatibility(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlPopClientAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlPopClientAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Compatibility_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}
