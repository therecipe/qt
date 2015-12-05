package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_2_Core struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_2_Core_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_4_2_Core_PTR() *QOpenGLFunctions_4_2_Core
}

func PointerFromQOpenGLFunctions_4_2_Core(ptr QOpenGLFunctions_4_2_Core_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_2_Core_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_4_2_CoreFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_2_Core {
	var n = new(QOpenGLFunctions_4_2_Core)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_2_Core) QOpenGLFunctions_4_2_Core_PTR() *QOpenGLFunctions_4_2_Core {
	return ptr
}
