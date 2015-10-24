package gui

//#include "qopenglfunctions_1_5.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_5 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_5ITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_1_5PTR() *QOpenGLFunctions_1_5
}

func PointerFromQOpenGLFunctions_1_5(ptr QOpenGLFunctions_1_5ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_5PTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_1_5FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_5 {
	var n = new(QOpenGLFunctions_1_5)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_5) QOpenGLFunctions_1_5PTR() *QOpenGLFunctions_1_5 {
	return ptr
}

func NewQOpenGLFunctions_1_5() *QOpenGLFunctions_1_5 {
	return QOpenGLFunctions_1_5FromPointer(unsafe.Pointer(C.QOpenGLFunctions_1_5_NewQOpenGLFunctions_1_5()))
}

func (ptr *QOpenGLFunctions_1_5) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_1_5_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_1_5) DestroyQOpenGLFunctions_1_5() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_DestroyQOpenGLFunctions_1_5(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlPopClientAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlPopClientAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_5) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_5_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}
