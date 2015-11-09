package gui

//#include "qopenglfunctions_4_5_compatibility.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_5_Compatibility struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_5_Compatibility_ITF interface {
	QAbstractOpenGLFunctions_ITF
	QOpenGLFunctions_4_5_Compatibility_PTR() *QOpenGLFunctions_4_5_Compatibility
}

func PointerFromQOpenGLFunctions_4_5_Compatibility(ptr QOpenGLFunctions_4_5_Compatibility_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_5_Compatibility_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFunctions_4_5_CompatibilityFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_5_Compatibility {
	var n = new(QOpenGLFunctions_4_5_Compatibility)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_5_Compatibility) QOpenGLFunctions_4_5_Compatibility_PTR() *QOpenGLFunctions_4_5_Compatibility {
	return ptr
}
