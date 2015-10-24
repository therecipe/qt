package multimedia

//#include "qabstractvideobuffer.h"
import "C"
import (
	"unsafe"
)

type QAbstractVideoBuffer struct {
	ptr unsafe.Pointer
}

type QAbstractVideoBufferITF interface {
	QAbstractVideoBufferPTR() *QAbstractVideoBuffer
}

func (p *QAbstractVideoBuffer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractVideoBuffer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractVideoBuffer(ptr QAbstractVideoBufferITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoBufferPTR().Pointer()
	}
	return nil
}

func QAbstractVideoBufferFromPointer(ptr unsafe.Pointer) *QAbstractVideoBuffer {
	var n = new(QAbstractVideoBuffer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractVideoBuffer) QAbstractVideoBufferPTR() *QAbstractVideoBuffer {
	return ptr
}

//QAbstractVideoBuffer::HandleType
type QAbstractVideoBuffer__HandleType int

var (
	QAbstractVideoBuffer__NoHandle         = QAbstractVideoBuffer__HandleType(0)
	QAbstractVideoBuffer__GLTextureHandle  = QAbstractVideoBuffer__HandleType(1)
	QAbstractVideoBuffer__XvShmImageHandle = QAbstractVideoBuffer__HandleType(2)
	QAbstractVideoBuffer__CoreImageHandle  = QAbstractVideoBuffer__HandleType(3)
	QAbstractVideoBuffer__QPixmapHandle    = QAbstractVideoBuffer__HandleType(4)
	QAbstractVideoBuffer__EGLImageHandle   = QAbstractVideoBuffer__HandleType(5)
	QAbstractVideoBuffer__UserHandle       = QAbstractVideoBuffer__HandleType(1000)
)

//QAbstractVideoBuffer::MapMode
type QAbstractVideoBuffer__MapMode int

var (
	QAbstractVideoBuffer__NotMapped = QAbstractVideoBuffer__MapMode(0x00)
	QAbstractVideoBuffer__ReadOnly  = QAbstractVideoBuffer__MapMode(0x01)
	QAbstractVideoBuffer__WriteOnly = QAbstractVideoBuffer__MapMode(0x02)
	QAbstractVideoBuffer__ReadWrite = QAbstractVideoBuffer__MapMode(QAbstractVideoBuffer__ReadOnly | QAbstractVideoBuffer__WriteOnly)
)

func (ptr *QAbstractVideoBuffer) Handle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractVideoBuffer_Handle(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAbstractVideoBuffer) HandleType() QAbstractVideoBuffer__HandleType {
	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__HandleType(C.QAbstractVideoBuffer_HandleType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractVideoBuffer) MapMode() QAbstractVideoBuffer__MapMode {
	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__MapMode(C.QAbstractVideoBuffer_MapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractVideoBuffer) Release() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_Release(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractVideoBuffer) Unmap() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_Unmap(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractVideoBuffer) DestroyQAbstractVideoBuffer() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_DestroyQAbstractVideoBuffer(C.QtObjectPtr(ptr.Pointer()))
	}
}
