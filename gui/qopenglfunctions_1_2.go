package gui

//#include "qopenglfunctions_1_2.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_2 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_2_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_1_2_PTR() *QOpenGLFunctions_1_2
}

func PointerFromQOpenGLFunctions_1_2(ptr QOpenGLFunctions_1_2_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_2_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_1_2FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_2 {
	var n = new(QOpenGLFunctions_1_2)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_2) QOpenGLFunctions_1_2_PTR() *QOpenGLFunctions_1_2 {
	return ptr
}
