package gui

//#include "qopenglbuffer.h"
import "C"
import (
	"unsafe"
)

type QOpenGLBuffer struct {
	ptr unsafe.Pointer
}

type QOpenGLBufferITF interface {
	QOpenGLBufferPTR() *QOpenGLBuffer
}

func (p *QOpenGLBuffer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLBuffer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLBuffer(ptr QOpenGLBufferITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLBufferPTR().Pointer()
	}
	return nil
}

func QOpenGLBufferFromPointer(ptr unsafe.Pointer) *QOpenGLBuffer {
	var n = new(QOpenGLBuffer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLBuffer) QOpenGLBufferPTR() *QOpenGLBuffer {
	return ptr
}

//QOpenGLBuffer::Access
type QOpenGLBuffer__Access int

var (
	QOpenGLBuffer__ReadOnly  = QOpenGLBuffer__Access(0x88B8)
	QOpenGLBuffer__WriteOnly = QOpenGLBuffer__Access(0x88B9)
	QOpenGLBuffer__ReadWrite = QOpenGLBuffer__Access(0x88BA)
)

//QOpenGLBuffer::RangeAccessFlag
type QOpenGLBuffer__RangeAccessFlag int

var (
	QOpenGLBuffer__RangeRead             = QOpenGLBuffer__RangeAccessFlag(0x0001)
	QOpenGLBuffer__RangeWrite            = QOpenGLBuffer__RangeAccessFlag(0x0002)
	QOpenGLBuffer__RangeInvalidate       = QOpenGLBuffer__RangeAccessFlag(0x0004)
	QOpenGLBuffer__RangeInvalidateBuffer = QOpenGLBuffer__RangeAccessFlag(0x0008)
	QOpenGLBuffer__RangeFlushExplicit    = QOpenGLBuffer__RangeAccessFlag(0x0010)
	QOpenGLBuffer__RangeUnsynchronized   = QOpenGLBuffer__RangeAccessFlag(0x0020)
)

//QOpenGLBuffer::Type
type QOpenGLBuffer__Type int

var (
	QOpenGLBuffer__VertexBuffer      = QOpenGLBuffer__Type(0x8892)
	QOpenGLBuffer__IndexBuffer       = QOpenGLBuffer__Type(0x8893)
	QOpenGLBuffer__PixelPackBuffer   = QOpenGLBuffer__Type(0x88EB)
	QOpenGLBuffer__PixelUnpackBuffer = QOpenGLBuffer__Type(0x88EC)
)

//QOpenGLBuffer::UsagePattern
type QOpenGLBuffer__UsagePattern int

var (
	QOpenGLBuffer__StreamDraw  = QOpenGLBuffer__UsagePattern(0x88E0)
	QOpenGLBuffer__StreamRead  = QOpenGLBuffer__UsagePattern(0x88E1)
	QOpenGLBuffer__StreamCopy  = QOpenGLBuffer__UsagePattern(0x88E2)
	QOpenGLBuffer__StaticDraw  = QOpenGLBuffer__UsagePattern(0x88E4)
	QOpenGLBuffer__StaticRead  = QOpenGLBuffer__UsagePattern(0x88E5)
	QOpenGLBuffer__StaticCopy  = QOpenGLBuffer__UsagePattern(0x88E6)
	QOpenGLBuffer__DynamicDraw = QOpenGLBuffer__UsagePattern(0x88E8)
	QOpenGLBuffer__DynamicRead = QOpenGLBuffer__UsagePattern(0x88E9)
	QOpenGLBuffer__DynamicCopy = QOpenGLBuffer__UsagePattern(0x88EA)
)

func NewQOpenGLBuffer() *QOpenGLBuffer {
	return QOpenGLBufferFromPointer(unsafe.Pointer(C.QOpenGLBuffer_NewQOpenGLBuffer()))
}

func NewQOpenGLBuffer2(ty QOpenGLBuffer__Type) *QOpenGLBuffer {
	return QOpenGLBufferFromPointer(unsafe.Pointer(C.QOpenGLBuffer_NewQOpenGLBuffer2(C.int(ty))))
}

func NewQOpenGLBuffer3(other QOpenGLBufferITF) *QOpenGLBuffer {
	return QOpenGLBufferFromPointer(unsafe.Pointer(C.QOpenGLBuffer_NewQOpenGLBuffer3(C.QtObjectPtr(PointerFromQOpenGLBuffer(other)))))
}

func (ptr *QOpenGLBuffer) Allocate2(count int) {
	if ptr.Pointer() != nil {
		C.QOpenGLBuffer_Allocate2(C.QtObjectPtr(ptr.Pointer()), C.int(count))
	}
}

func (ptr *QOpenGLBuffer) Bind() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLBuffer_Bind(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLBuffer) Create() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLBuffer_Create(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLBuffer) Destroy() {
	if ptr.Pointer() != nil {
		C.QOpenGLBuffer_Destroy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLBuffer) IsCreated() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLBuffer_IsCreated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLBuffer) Map(access QOpenGLBuffer__Access) {
	if ptr.Pointer() != nil {
		C.QOpenGLBuffer_Map(C.QtObjectPtr(ptr.Pointer()), C.int(access))
	}
}

func (ptr *QOpenGLBuffer) MapRange(offset int, count int, access QOpenGLBuffer__RangeAccessFlag) {
	if ptr.Pointer() != nil {
		C.QOpenGLBuffer_MapRange(C.QtObjectPtr(ptr.Pointer()), C.int(offset), C.int(count), C.int(access))
	}
}

func (ptr *QOpenGLBuffer) Release() {
	if ptr.Pointer() != nil {
		C.QOpenGLBuffer_Release(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QOpenGLBuffer_Release2(ty QOpenGLBuffer__Type) {
	C.QOpenGLBuffer_QOpenGLBuffer_Release2(C.int(ty))
}

func (ptr *QOpenGLBuffer) SetUsagePattern(value QOpenGLBuffer__UsagePattern) {
	if ptr.Pointer() != nil {
		C.QOpenGLBuffer_SetUsagePattern(C.QtObjectPtr(ptr.Pointer()), C.int(value))
	}
}

func (ptr *QOpenGLBuffer) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLBuffer_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLBuffer) Type() QOpenGLBuffer__Type {
	if ptr.Pointer() != nil {
		return QOpenGLBuffer__Type(C.QOpenGLBuffer_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLBuffer) Unmap() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLBuffer_Unmap(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLBuffer) UsagePattern() QOpenGLBuffer__UsagePattern {
	if ptr.Pointer() != nil {
		return QOpenGLBuffer__UsagePattern(C.QOpenGLBuffer_UsagePattern(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLBuffer) DestroyQOpenGLBuffer() {
	if ptr.Pointer() != nil {
		C.QOpenGLBuffer_DestroyQOpenGLBuffer(C.QtObjectPtr(ptr.Pointer()))
	}
}
