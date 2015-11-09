package gui

//#include "qopenglfunctions_1_4.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_4 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_4_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_1_4_PTR() *QOpenGLFunctions_1_4
}

func PointerFromQOpenGLFunctions_1_4(ptr QOpenGLFunctions_1_4_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_4_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_1_4FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_4 {
	var n = new(QOpenGLFunctions_1_4)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_4) QOpenGLFunctions_1_4_PTR() *QOpenGLFunctions_1_4 {
	return ptr
}
