package gui

//#include "qopenglfunctions_3_3_core.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_3_3_Core struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_3_3_CoreITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_3_3_CorePTR() *QOpenGLFunctions_3_3_Core
}

func PointerFromQOpenGLFunctions_3_3_Core(ptr QOpenGLFunctions_3_3_CoreITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_3_3_CorePTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_3_3_CoreFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_3_3_Core {
	var n = new(QOpenGLFunctions_3_3_Core)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_3_3_Core) QOpenGLFunctions_3_3_CorePTR() *QOpenGLFunctions_3_3_Core {
	return ptr
}

func NewQOpenGLFunctions_3_3_Core() *QOpenGLFunctions_3_3_Core {
	return QOpenGLFunctions_3_3_CoreFromPointer(unsafe.Pointer(C.QOpenGLFunctions_3_3_Core_NewQOpenGLFunctions_3_3_Core()))
}

func (ptr *QOpenGLFunctions_3_3_Core) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_3_3_Core_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_3_3_Core) DestroyQOpenGLFunctions_3_3_Core() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Core_DestroyQOpenGLFunctions_3_3_Core(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Core) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Core_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Core) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Core_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Core) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Core_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_3_3_Core) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_3_3_Core_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}
