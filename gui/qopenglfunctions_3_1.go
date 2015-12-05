package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_3_1 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_3_1_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_3_1_PTR() *QOpenGLFunctions_3_1
}

func PointerFromQOpenGLFunctions_3_1(ptr QOpenGLFunctions_3_1_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_3_1_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_3_1FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_3_1 {
	var n = new(QOpenGLFunctions_3_1)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_3_1) QOpenGLFunctions_3_1_PTR() *QOpenGLFunctions_3_1 {
	return ptr
}
