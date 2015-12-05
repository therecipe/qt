package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_3_Compatibility struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_3_Compatibility_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_4_3_Compatibility_PTR() *QOpenGLFunctions_4_3_Compatibility
}

func PointerFromQOpenGLFunctions_4_3_Compatibility(ptr QOpenGLFunctions_4_3_Compatibility_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_3_Compatibility_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_4_3_CompatibilityFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_3_Compatibility {
	var n = new(QOpenGLFunctions_4_3_Compatibility)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_3_Compatibility) QOpenGLFunctions_4_3_Compatibility_PTR() *QOpenGLFunctions_4_3_Compatibility {
	return ptr
}
