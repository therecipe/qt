package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLVersionProfile struct {
	ptr unsafe.Pointer
}

type QOpenGLVersionProfile_ITF interface {
	QOpenGLVersionProfile_PTR() *QOpenGLVersionProfile
}

func (p *QOpenGLVersionProfile) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLVersionProfile) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLVersionProfile(ptr QOpenGLVersionProfile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLVersionProfile_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLVersionProfileFromPointer(ptr unsafe.Pointer) *QOpenGLVersionProfile {
	var n = new(QOpenGLVersionProfile)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLVersionProfile) QOpenGLVersionProfile_PTR() *QOpenGLVersionProfile {
	return ptr
}
