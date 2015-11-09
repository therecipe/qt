package multimedia

//#include "qabstractvideobuffer.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractVideoBuffer struct {
	ptr unsafe.Pointer
}

type QAbstractVideoBuffer_ITF interface {
	QAbstractVideoBuffer_PTR() *QAbstractVideoBuffer
}

func (p *QAbstractVideoBuffer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractVideoBuffer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractVideoBuffer(ptr QAbstractVideoBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoBuffer_PTR().Pointer()
	}
	return nil
}

func NewQAbstractVideoBufferFromPointer(ptr unsafe.Pointer) *QAbstractVideoBuffer {
	var n = new(QAbstractVideoBuffer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractVideoBuffer) QAbstractVideoBuffer_PTR() *QAbstractVideoBuffer {
	return ptr
}

//QAbstractVideoBuffer::HandleType
type QAbstractVideoBuffer__HandleType int64

const (
	QAbstractVideoBuffer__NoHandle         = QAbstractVideoBuffer__HandleType(0)
	QAbstractVideoBuffer__GLTextureHandle  = QAbstractVideoBuffer__HandleType(1)
	QAbstractVideoBuffer__XvShmImageHandle = QAbstractVideoBuffer__HandleType(2)
	QAbstractVideoBuffer__CoreImageHandle  = QAbstractVideoBuffer__HandleType(3)
	QAbstractVideoBuffer__QPixmapHandle    = QAbstractVideoBuffer__HandleType(4)
	QAbstractVideoBuffer__EGLImageHandle   = QAbstractVideoBuffer__HandleType(5)
	QAbstractVideoBuffer__UserHandle       = QAbstractVideoBuffer__HandleType(1000)
)

//QAbstractVideoBuffer::MapMode
type QAbstractVideoBuffer__MapMode int64

const (
	QAbstractVideoBuffer__NotMapped = QAbstractVideoBuffer__MapMode(0x00)
	QAbstractVideoBuffer__ReadOnly  = QAbstractVideoBuffer__MapMode(0x01)
	QAbstractVideoBuffer__WriteOnly = QAbstractVideoBuffer__MapMode(0x02)
	QAbstractVideoBuffer__ReadWrite = QAbstractVideoBuffer__MapMode(QAbstractVideoBuffer__ReadOnly | QAbstractVideoBuffer__WriteOnly)
)

func (ptr *QAbstractVideoBuffer) Handle() *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractVideoBuffer_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractVideoBuffer) HandleType() QAbstractVideoBuffer__HandleType {
	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__HandleType(C.QAbstractVideoBuffer_HandleType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractVideoBuffer) MapMode() QAbstractVideoBuffer__MapMode {
	if ptr.Pointer() != nil {
		return QAbstractVideoBuffer__MapMode(C.QAbstractVideoBuffer_MapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractVideoBuffer) Release() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_Release(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoBuffer) Unmap() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_Unmap(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoBuffer) DestroyQAbstractVideoBuffer() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoBuffer_DestroyQAbstractVideoBuffer(ptr.Pointer())
	}
}
