package gui

//#include "qopenglfunctions_1_2.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_2 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_2ITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_1_2PTR() *QOpenGLFunctions_1_2
}

func PointerFromQOpenGLFunctions_1_2(ptr QOpenGLFunctions_1_2ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_2PTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_1_2FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_2 {
	var n = new(QOpenGLFunctions_1_2)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_2) QOpenGLFunctions_1_2PTR() *QOpenGLFunctions_1_2 {
	return ptr
}

func NewQOpenGLFunctions_1_2() *QOpenGLFunctions_1_2 {
	return QOpenGLFunctions_1_2FromPointer(unsafe.Pointer(C.QOpenGLFunctions_1_2_NewQOpenGLFunctions_1_2()))
}

func (ptr *QOpenGLFunctions_1_2) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_1_2_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_1_2) DestroyQOpenGLFunctions_1_2() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_DestroyQOpenGLFunctions_1_2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlPopClientAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlPopClientAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_2) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_2_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}
