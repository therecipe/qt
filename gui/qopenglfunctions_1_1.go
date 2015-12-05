package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_1 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_1_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_1_1_PTR() *QOpenGLFunctions_1_1
}

func PointerFromQOpenGLFunctions_1_1(ptr QOpenGLFunctions_1_1_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_1_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_1_1FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_1 {
	var n = new(QOpenGLFunctions_1_1)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_1) QOpenGLFunctions_1_1_PTR() *QOpenGLFunctions_1_1 {
	return ptr
}
