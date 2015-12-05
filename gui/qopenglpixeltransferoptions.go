package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLPixelTransferOptions struct {
	ptr unsafe.Pointer
}

type QOpenGLPixelTransferOptions_ITF interface {
	QOpenGLPixelTransferOptions_PTR() *QOpenGLPixelTransferOptions
}

func (p *QOpenGLPixelTransferOptions) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLPixelTransferOptions) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLPixelTransferOptions(ptr QOpenGLPixelTransferOptions_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLPixelTransferOptions_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLPixelTransferOptionsFromPointer(ptr unsafe.Pointer) *QOpenGLPixelTransferOptions {
	var n = new(QOpenGLPixelTransferOptions)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLPixelTransferOptions) QOpenGLPixelTransferOptions_PTR() *QOpenGLPixelTransferOptions {
	return ptr
}
