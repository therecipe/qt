package gui

//#include "qopenglfunctions_3_1.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_3_1 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_3_1ITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_3_1PTR() *QOpenGLFunctions_3_1
}

func PointerFromQOpenGLFunctions_3_1(ptr QOpenGLFunctions_3_1ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_3_1PTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_3_1FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_3_1 {
	var n = new(QOpenGLFunctions_3_1)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_3_1) QOpenGLFunctions_3_1PTR() *QOpenGLFunctions_3_1 {
	return ptr
}

func NewQOpenGLFunctions_3_1() *QOpenGLFunctions_3_1 {
	return QOpenGLFunctions_3_1FromPointer(unsafe.Pointer(C.QOpenGLFunctions_3_1_NewQOpenGLFunctions_3_1()))
}

func (ptr *QOpenGLFunctions_3_1) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_3_1_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_3_1) DestroyQOpenGLFunctions_3_1() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_1_DestroyQOpenGLFunctions_3_1(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_1) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_1_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_1) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_1_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_1) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_1_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_1) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_1_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}
