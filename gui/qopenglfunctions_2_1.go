package gui

//#include "qopenglfunctions_2_1.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_2_1 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_2_1_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_2_1_PTR() *QOpenGLFunctions_2_1
}

func PointerFromQOpenGLFunctions_2_1(ptr QOpenGLFunctions_2_1_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_2_1_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_2_1FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_2_1 {
	var n = new(QOpenGLFunctions_2_1)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_2_1) QOpenGLFunctions_2_1_PTR() *QOpenGLFunctions_2_1 {
	return ptr
}
