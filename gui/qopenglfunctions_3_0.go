package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_3_0 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_3_0_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_3_0_PTR() *QOpenGLFunctions_3_0
}

func PointerFromQOpenGLFunctions_3_0(ptr QOpenGLFunctions_3_0_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_3_0_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_3_0FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_3_0 {
	var n = new(QOpenGLFunctions_3_0)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_3_0) QOpenGLFunctions_3_0_PTR() *QOpenGLFunctions_3_0 {
	return ptr
}
