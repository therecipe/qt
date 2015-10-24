package gui

//#include "qopenglframebufferobject.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLFramebufferObject struct {
	ptr unsafe.Pointer
}

type QOpenGLFramebufferObjectITF interface {
	QOpenGLFramebufferObjectPTR() *QOpenGLFramebufferObject
}

func (p *QOpenGLFramebufferObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLFramebufferObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLFramebufferObject(ptr QOpenGLFramebufferObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFramebufferObjectPTR().Pointer()
	}
	return nil
}

func QOpenGLFramebufferObjectFromPointer(ptr unsafe.Pointer) *QOpenGLFramebufferObject {
	var n = new(QOpenGLFramebufferObject)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFramebufferObject) QOpenGLFramebufferObjectPTR() *QOpenGLFramebufferObject {
	return ptr
}

//QOpenGLFramebufferObject::Attachment
type QOpenGLFramebufferObject__Attachment int

var (
	QOpenGLFramebufferObject__NoAttachment         = QOpenGLFramebufferObject__Attachment(0)
	QOpenGLFramebufferObject__CombinedDepthStencil = QOpenGLFramebufferObject__Attachment(1)
	QOpenGLFramebufferObject__Depth                = QOpenGLFramebufferObject__Attachment(2)
)

func (ptr *QOpenGLFramebufferObject) Bind() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFramebufferObject_Bind(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QOpenGLFramebufferObject_BindDefault() bool {
	return C.QOpenGLFramebufferObject_QOpenGLFramebufferObject_BindDefault() != 0
}

func QOpenGLFramebufferObject_HasOpenGLFramebufferBlit() bool {
	return C.QOpenGLFramebufferObject_QOpenGLFramebufferObject_HasOpenGLFramebufferBlit() != 0
}

func QOpenGLFramebufferObject_HasOpenGLFramebufferObjects() bool {
	return C.QOpenGLFramebufferObject_QOpenGLFramebufferObject_HasOpenGLFramebufferObjects() != 0
}

func (ptr *QOpenGLFramebufferObject) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFramebufferObject_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFramebufferObject) Release() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFramebufferObject_Release(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFramebufferObject) DestroyQOpenGLFramebufferObject() {
	if ptr.Pointer() != nil {
		C.QOpenGLFramebufferObject_DestroyQOpenGLFramebufferObject(C.QtObjectPtr(ptr.Pointer()))
	}
}

func NewQOpenGLFramebufferObject3(size core.QSizeITF, format QOpenGLFramebufferObjectFormatITF) *QOpenGLFramebufferObject {
	return QOpenGLFramebufferObjectFromPointer(unsafe.Pointer(C.QOpenGLFramebufferObject_NewQOpenGLFramebufferObject3(C.QtObjectPtr(core.PointerFromQSize(size)), C.QtObjectPtr(PointerFromQOpenGLFramebufferObjectFormat(format)))))
}

func NewQOpenGLFramebufferObject4(width int, height int, format QOpenGLFramebufferObjectFormatITF) *QOpenGLFramebufferObject {
	return QOpenGLFramebufferObjectFromPointer(unsafe.Pointer(C.QOpenGLFramebufferObject_NewQOpenGLFramebufferObject4(C.int(width), C.int(height), C.QtObjectPtr(PointerFromQOpenGLFramebufferObjectFormat(format)))))
}

func (ptr *QOpenGLFramebufferObject) Attachment() QOpenGLFramebufferObject__Attachment {
	if ptr.Pointer() != nil {
		return QOpenGLFramebufferObject__Attachment(C.QOpenGLFramebufferObject_Attachment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLFramebufferObject) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLFramebufferObject_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLFramebufferObject) IsBound() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFramebufferObject_IsBound(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFramebufferObject) SetAttachment(attachment QOpenGLFramebufferObject__Attachment) {
	if ptr.Pointer() != nil {
		C.QOpenGLFramebufferObject_SetAttachment(C.QtObjectPtr(ptr.Pointer()), C.int(attachment))
	}
}

func (ptr *QOpenGLFramebufferObject) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLFramebufferObject_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
