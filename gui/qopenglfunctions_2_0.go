package gui

//#include "qopenglfunctions_2_0.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_2_0 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_2_0_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_2_0_PTR() *QOpenGLFunctions_2_0
}

func PointerFromQOpenGLFunctions_2_0(ptr QOpenGLFunctions_2_0_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_2_0_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_2_0FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_2_0 {
	var n = new(QOpenGLFunctions_2_0)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_2_0) QOpenGLFunctions_2_0_PTR() *QOpenGLFunctions_2_0 {
	return ptr
}
