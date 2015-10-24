package gui

//#include "qopenglfunctions_1_1.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_1 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_1ITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_1_1PTR() *QOpenGLFunctions_1_1
}

func PointerFromQOpenGLFunctions_1_1(ptr QOpenGLFunctions_1_1ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_1PTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_1_1FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_1 {
	var n = new(QOpenGLFunctions_1_1)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_1) QOpenGLFunctions_1_1PTR() *QOpenGLFunctions_1_1 {
	return ptr
}

func NewQOpenGLFunctions_1_1() *QOpenGLFunctions_1_1 {
	return QOpenGLFunctions_1_1FromPointer(unsafe.Pointer(C.QOpenGLFunctions_1_1_NewQOpenGLFunctions_1_1()))
}

func (ptr *QOpenGLFunctions_1_1) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_1_1_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_1_1) DestroyQOpenGLFunctions_1_1() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_DestroyQOpenGLFunctions_1_1(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlPopClientAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlPopClientAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_1) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_1_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}
