package gui

//#include "qopenglfunctions_4_0_compatibility.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_0_Compatibility struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_0_Compatibility_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_4_0_Compatibility_PTR() *QOpenGLFunctions_4_0_Compatibility
}

func PointerFromQOpenGLFunctions_4_0_Compatibility(ptr QOpenGLFunctions_4_0_Compatibility_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_0_Compatibility_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_4_0_CompatibilityFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_0_Compatibility {
	var n = new(QOpenGLFunctions_4_0_Compatibility)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_0_Compatibility) QOpenGLFunctions_4_0_Compatibility_PTR() *QOpenGLFunctions_4_0_Compatibility {
	return ptr
}
