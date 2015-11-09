package gui

//#include "qopenglfunctions_3_3_core.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_3_3_Core struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_3_3_Core_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_3_3_Core_PTR() *QOpenGLFunctions_3_3_Core
}

func PointerFromQOpenGLFunctions_3_3_Core(ptr QOpenGLFunctions_3_3_Core_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_3_3_Core_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_3_3_CoreFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_3_3_Core {
	var n = new(QOpenGLFunctions_3_3_Core)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_3_3_Core) QOpenGLFunctions_3_3_Core_PTR() *QOpenGLFunctions_3_3_Core {
	return ptr
}
