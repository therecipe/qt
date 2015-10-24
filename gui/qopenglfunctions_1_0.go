package gui

//#include "qopenglfunctions_1_0.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_0 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_0ITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_1_0PTR() *QOpenGLFunctions_1_0
}

func PointerFromQOpenGLFunctions_1_0(ptr QOpenGLFunctions_1_0ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_0PTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_1_0FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_0 {
	var n = new(QOpenGLFunctions_1_0)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_0) QOpenGLFunctions_1_0PTR() *QOpenGLFunctions_1_0 {
	return ptr
}

func NewQOpenGLFunctions_1_0() *QOpenGLFunctions_1_0 {
	return QOpenGLFunctions_1_0FromPointer(unsafe.Pointer(C.QOpenGLFunctions_1_0_NewQOpenGLFunctions_1_0()))
}

func (ptr *QOpenGLFunctions_1_0) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_1_0_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_1_0) DestroyQOpenGLFunctions_1_0() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_DestroyQOpenGLFunctions_1_0(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_0) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_0) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_0) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_0) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_0) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_0) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_0) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_0) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_0) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_0) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_0_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}
