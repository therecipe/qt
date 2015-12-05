package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_3_2_Core struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_3_2_Core_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_3_2_Core_PTR() *QOpenGLFunctions_3_2_Core
}

func PointerFromQOpenGLFunctions_3_2_Core(ptr QOpenGLFunctions_3_2_Core_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_3_2_Core_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_3_2_CoreFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_3_2_Core {
	var n = new(QOpenGLFunctions_3_2_Core)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_3_2_Core) QOpenGLFunctions_3_2_Core_PTR() *QOpenGLFunctions_3_2_Core {
	return ptr
}
