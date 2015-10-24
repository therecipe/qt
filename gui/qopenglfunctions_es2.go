package gui

//#include "qopenglfunctions_es2.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_ES2 struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_ES2ITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_ES2PTR() *QOpenGLFunctions_ES2
}

func PointerFromQOpenGLFunctions_ES2(ptr QOpenGLFunctions_ES2ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_ES2PTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_ES2FromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_ES2 {
	var n = new(QOpenGLFunctions_ES2)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_ES2) QOpenGLFunctions_ES2PTR() *QOpenGLFunctions_ES2 {
	return ptr
}
