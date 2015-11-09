package gui

//#include "qopenglfunctions_1_0.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_0 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_0_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_1_0_PTR() *QOpenGLFunctions_1_0
}

func PointerFromQOpenGLFunctions_1_0(ptr QOpenGLFunctions_1_0_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_0_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_1_0FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_0 {
	var n = new(QOpenGLFunctions_1_0)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_0) QOpenGLFunctions_1_0_PTR() *QOpenGLFunctions_1_0 {
	return ptr
}
