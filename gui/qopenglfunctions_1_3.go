package gui

//#include "qopenglfunctions_1_3.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_3 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_3ITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_1_3PTR() *QOpenGLFunctions_1_3
}

func PointerFromQOpenGLFunctions_1_3(ptr QOpenGLFunctions_1_3ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_3PTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_1_3FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_3 {
	var n = new(QOpenGLFunctions_1_3)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_3) QOpenGLFunctions_1_3PTR() *QOpenGLFunctions_1_3 {
	return ptr
}

func NewQOpenGLFunctions_1_3() *QOpenGLFunctions_1_3 {
	return QOpenGLFunctions_1_3FromPointer(unsafe.Pointer(C.QOpenGLFunctions_1_3_NewQOpenGLFunctions_1_3()))
}

func (ptr *QOpenGLFunctions_1_3) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_1_3_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_1_3) DestroyQOpenGLFunctions_1_3() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_DestroyQOpenGLFunctions_1_3(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlPopClientAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlPopClientAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_3) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_3_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}
