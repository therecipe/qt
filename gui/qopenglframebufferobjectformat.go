package gui

//#include "qopenglframebufferobjectformat.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QOpenGLFramebufferObjectFormat struct {
	ptr unsafe.Pointer
}

type QOpenGLFramebufferObjectFormatITF interface {
	QOpenGLFramebufferObjectFormatPTR() *QOpenGLFramebufferObjectFormat
}

func (p *QOpenGLFramebufferObjectFormat) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLFramebufferObjectFormat) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLFramebufferObjectFormat(ptr QOpenGLFramebufferObjectFormatITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFramebufferObjectFormatPTR().Pointer()
	}
	return nil
}

func QOpenGLFramebufferObjectFormatFromPointer(ptr unsafe.Pointer) *QOpenGLFramebufferObjectFormat {
	var n = new(QOpenGLFramebufferObjectFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFramebufferObjectFormat) QOpenGLFramebufferObjectFormatPTR() *QOpenGLFramebufferObjectFormat {
	return ptr
}

func NewQOpenGLFramebufferObjectFormat() *QOpenGLFramebufferObjectFormat {
	return QOpenGLFramebufferObjectFormatFromPointer(unsafe.Pointer(C.QOpenGLFramebufferObjectFormat_NewQOpenGLFramebufferObjectFormat()))
}

func NewQOpenGLFramebufferObjectFormat2(other QOpenGLFramebufferObjectFormatITF) *QOpenGLFramebufferObjectFormat {
	return QOpenGLFramebufferObjectFormatFromPointer(unsafe.Pointer(C.QOpenGLFramebufferObjectFormat_NewQOpenGLFramebufferObjectFormat2(C.QtObjectPtr(PointerFromQOpenGLFramebufferObjectFormat(other)))))
}

func (ptr *QOpenGLFramebufferObjectFormat) Attachment() QOpenGLFramebufferObject__Attachment {
	if ptr.Pointer() != nil {
		return QOpenGLFramebufferObject__Attachment(C.QOpenGLFramebufferObjectFormat_Attachment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLFramebufferObjectFormat) Mipmap() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFramebufferObjectFormat_Mipmap(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFramebufferObjectFormat) Samples() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLFramebufferObjectFormat_Samples(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLFramebufferObjectFormat) SetAttachment(attachment QOpenGLFramebufferObject__Attachment) {
	if ptr.Pointer() != nil {
		C.QOpenGLFramebufferObjectFormat_SetAttachment(C.QtObjectPtr(ptr.Pointer()), C.int(attachment))
	}
}

func (ptr *QOpenGLFramebufferObjectFormat) SetMipmap(enabled bool) {
	if ptr.Pointer() != nil {
		C.QOpenGLFramebufferObjectFormat_SetMipmap(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QOpenGLFramebufferObjectFormat) SetSamples(samples int) {
	if ptr.Pointer() != nil {
		C.QOpenGLFramebufferObjectFormat_SetSamples(C.QtObjectPtr(ptr.Pointer()), C.int(samples))
	}
}

func (ptr *QOpenGLFramebufferObjectFormat) DestroyQOpenGLFramebufferObjectFormat() {
	if ptr.Pointer() != nil {
		C.QOpenGLFramebufferObjectFormat_DestroyQOpenGLFramebufferObjectFormat(C.QtObjectPtr(ptr.Pointer()))
	}
}
