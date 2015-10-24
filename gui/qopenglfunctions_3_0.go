package gui

//#include "qopenglfunctions_3_0.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_3_0 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_3_0ITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_3_0PTR() *QOpenGLFunctions_3_0
}

func PointerFromQOpenGLFunctions_3_0(ptr QOpenGLFunctions_3_0ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_3_0PTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_3_0FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_3_0 {
	var n = new(QOpenGLFunctions_3_0)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_3_0) QOpenGLFunctions_3_0PTR() *QOpenGLFunctions_3_0 {
	return ptr
}

func NewQOpenGLFunctions_3_0() *QOpenGLFunctions_3_0 {
	return QOpenGLFunctions_3_0FromPointer(unsafe.Pointer(C.QOpenGLFunctions_3_0_NewQOpenGLFunctions_3_0()))
}

func (ptr *QOpenGLFunctions_3_0) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_3_0_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_3_0) DestroyQOpenGLFunctions_3_0() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_DestroyQOpenGLFunctions_3_0(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlPopClientAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlPopClientAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_0) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_0_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}
