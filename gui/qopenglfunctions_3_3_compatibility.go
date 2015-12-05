package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_3_3_Compatibility struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_3_3_Compatibility_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_3_3_Compatibility_PTR() *QOpenGLFunctions_3_3_Compatibility
}

func PointerFromQOpenGLFunctions_3_3_Compatibility(ptr QOpenGLFunctions_3_3_Compatibility_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_3_3_Compatibility_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_3_3_CompatibilityFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_3_3_Compatibility {
	var n = new(QOpenGLFunctions_3_3_Compatibility)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_3_3_Compatibility) QOpenGLFunctions_3_3_Compatibility_PTR() *QOpenGLFunctions_3_3_Compatibility {
	return ptr
}
