package gui

//#include "qopenglfunctions_2_0.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_2_0 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_2_0ITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_2_0PTR() *QOpenGLFunctions_2_0
}

func PointerFromQOpenGLFunctions_2_0(ptr QOpenGLFunctions_2_0ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_2_0PTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_2_0FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_2_0 {
	var n = new(QOpenGLFunctions_2_0)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_2_0) QOpenGLFunctions_2_0PTR() *QOpenGLFunctions_2_0 {
	return ptr
}

func NewQOpenGLFunctions_2_0() *QOpenGLFunctions_2_0 {
	return QOpenGLFunctions_2_0FromPointer(unsafe.Pointer(C.QOpenGLFunctions_2_0_NewQOpenGLFunctions_2_0()))
}

func (ptr *QOpenGLFunctions_2_0) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_2_0_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_2_0) DestroyQOpenGLFunctions_2_0() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_DestroyQOpenGLFunctions_2_0(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlPopClientAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlPopClientAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_2_0) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_2_0_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}
