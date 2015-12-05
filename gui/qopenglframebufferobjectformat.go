package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFramebufferObjectFormat struct {
	ptr unsafe.Pointer
}

type QOpenGLFramebufferObjectFormat_ITF interface {
	QOpenGLFramebufferObjectFormat_PTR() *QOpenGLFramebufferObjectFormat
}

func (p *QOpenGLFramebufferObjectFormat) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLFramebufferObjectFormat) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLFramebufferObjectFormat(ptr QOpenGLFramebufferObjectFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFramebufferObjectFormat_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFramebufferObjectFormatFromPointer(ptr unsafe.Pointer) *QOpenGLFramebufferObjectFormat {
	var n = new(QOpenGLFramebufferObjectFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFramebufferObjectFormat) QOpenGLFramebufferObjectFormat_PTR() *QOpenGLFramebufferObjectFormat {
	return ptr
}
