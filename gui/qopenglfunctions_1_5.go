package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_5 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_5_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_1_5_PTR() *QOpenGLFunctions_1_5
}

func PointerFromQOpenGLFunctions_1_5(ptr QOpenGLFunctions_1_5_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_5_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_1_5FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_5 {
	var n = new(QOpenGLFunctions_1_5)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_5) QOpenGLFunctions_1_5_PTR() *QOpenGLFunctions_1_5 {
	return ptr
}
