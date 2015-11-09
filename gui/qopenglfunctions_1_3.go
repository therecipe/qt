package gui

//#include "qopenglfunctions_1_3.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_1_3 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_1_3_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_1_3_PTR() *QOpenGLFunctions_1_3
}

func PointerFromQOpenGLFunctions_1_3(ptr QOpenGLFunctions_1_3_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_1_3_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_1_3FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_1_3 {
	var n = new(QOpenGLFunctions_1_3)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_1_3) QOpenGLFunctions_1_3_PTR() *QOpenGLFunctions_1_3 {
	return ptr
}
