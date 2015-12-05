package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFramebufferObject struct {
	ptr unsafe.Pointer
}

type QOpenGLFramebufferObject_ITF interface {
	QOpenGLFramebufferObject_PTR() *QOpenGLFramebufferObject
}

func (p *QOpenGLFramebufferObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLFramebufferObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLFramebufferObject(ptr QOpenGLFramebufferObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFramebufferObject_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLFramebufferObjectFromPointer(ptr unsafe.Pointer) *QOpenGLFramebufferObject {
	var n = new(QOpenGLFramebufferObject)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFramebufferObject) QOpenGLFramebufferObject_PTR() *QOpenGLFramebufferObject {
	return ptr
}

//QOpenGLFramebufferObject::Attachment
type QOpenGLFramebufferObject__Attachment int64

const (
	QOpenGLFramebufferObject__NoAttachment         = QOpenGLFramebufferObject__Attachment(0)
	QOpenGLFramebufferObject__CombinedDepthStencil = QOpenGLFramebufferObject__Attachment(1)
	QOpenGLFramebufferObject__Depth                = QOpenGLFramebufferObject__Attachment(2)
)
