package gui

//#include "qopenglfunctions_1_4.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_4 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_4ITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_1_4PTR() *QOpenGLFunctions_1_4
}

func PointerFromQOpenGLFunctions_1_4(ptr QOpenGLFunctions_1_4ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_4PTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_1_4FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_4 {
	var n = new(QOpenGLFunctions_1_4)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_4) QOpenGLFunctions_1_4PTR() *QOpenGLFunctions_1_4 {
	return ptr
}

func NewQOpenGLFunctions_1_4() *QOpenGLFunctions_1_4 {
	return QOpenGLFunctions_1_4FromPointer(unsafe.Pointer(C.QOpenGLFunctions_1_4_NewQOpenGLFunctions_1_4()))
}

func (ptr *QOpenGLFunctions_1_4) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_1_4_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_1_4) DestroyQOpenGLFunctions_1_4() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_DestroyQOpenGLFunctions_1_4(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlEnd() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlEndList() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlEndList(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlInitNames() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlInitNames(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlLoadIdentity() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlLoadIdentity(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlPopAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlPopAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlPopClientAttrib() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlPopClientAttrib(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlPopMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlPopMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlPopName() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlPopName(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_1_4) GlPushMatrix() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_1_4_GlPushMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}
