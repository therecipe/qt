package gui

//#include "qopenglbuffer.h"
import "C"
import (
	"unsafe"
)

type QOpenGLBuffer struct {
	ptr unsafe.Pointer
}

type QOpenGLBuffer_ITF interface {
	QOpenGLBuffer_PTR() *QOpenGLBuffer
}

func (p *QOpenGLBuffer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLBuffer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLBuffer(ptr QOpenGLBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLBuffer_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLBufferFromPointer(ptr unsafe.Pointer) *QOpenGLBuffer {
	var n = new(QOpenGLBuffer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLBuffer) QOpenGLBuffer_PTR() *QOpenGLBuffer {
	return ptr
}

//QOpenGLBuffer::Access
type QOpenGLBuffer__Access int64

const (
	QOpenGLBuffer__ReadOnly  = QOpenGLBuffer__Access(0x88B8)
	QOpenGLBuffer__WriteOnly = QOpenGLBuffer__Access(0x88B9)
	QOpenGLBuffer__ReadWrite = QOpenGLBuffer__Access(0x88BA)
)

//QOpenGLBuffer::RangeAccessFlag
type QOpenGLBuffer__RangeAccessFlag int64

const (
	QOpenGLBuffer__RangeRead             = QOpenGLBuffer__RangeAccessFlag(0x0001)
	QOpenGLBuffer__RangeWrite            = QOpenGLBuffer__RangeAccessFlag(0x0002)
	QOpenGLBuffer__RangeInvalidate       = QOpenGLBuffer__RangeAccessFlag(0x0004)
	QOpenGLBuffer__RangeInvalidateBuffer = QOpenGLBuffer__RangeAccessFlag(0x0008)
	QOpenGLBuffer__RangeFlushExplicit    = QOpenGLBuffer__RangeAccessFlag(0x0010)
	QOpenGLBuffer__RangeUnsynchronized   = QOpenGLBuffer__RangeAccessFlag(0x0020)
)

//QOpenGLBuffer::Type
type QOpenGLBuffer__Type int64

const (
	QOpenGLBuffer__VertexBuffer      = QOpenGLBuffer__Type(0x8892)
	QOpenGLBuffer__IndexBuffer       = QOpenGLBuffer__Type(0x8893)
	QOpenGLBuffer__PixelPackBuffer   = QOpenGLBuffer__Type(0x88EB)
	QOpenGLBuffer__PixelUnpackBuffer = QOpenGLBuffer__Type(0x88EC)
)

//QOpenGLBuffer::UsagePattern
type QOpenGLBuffer__UsagePattern int64

const (
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
