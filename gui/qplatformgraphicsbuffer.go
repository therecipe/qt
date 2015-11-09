package gui

//#include "qplatformgraphicsbuffer.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPlatformGraphicsBuffer struct {
	core.QObject
}

type QPlatformGraphicsBuffer_ITF interface {
	core.QObject_ITF
	QPlatformGraphicsBuffer_PTR() *QPlatformGraphicsBuffer
}

func PointerFromQPlatformGraphicsBuffer(ptr QPlatformGraphicsBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlatformGraphicsBuffer_PTR().Pointer()
	}
	return nil
}

func NewQPlatformGraphicsBufferFromPointer(ptr unsafe.Pointer) *QPlatformGraphicsBuffer {
	var n = new(QPlatformGraphicsBuffer)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QPlatformGraphicsBuffer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlatformGraphicsBuffer) QPlatformGraphicsBuffer_PTR() *QPlatformGraphicsBuffer {
	return ptr
}

//QPlatformGraphicsBuffer::AccessType
type QPlatformGraphicsBuffer__AccessType int64

const (
	QPlatformGraphicsBuffer__None          = QPlatformGraphicsBuffer__AccessType(0x00)
	QPlatformGraphicsBuffer__SWReadAccess  = QPlatformGraphicsBuffer__AccessType(0x01)
	QPlatformGraphicsBuffer__SWWriteAccess = QPlatformGraphicsBuffer__AccessType(0x02)
	QPlatformGraphicsBuffer__TextureAccess = QPlatformGraphicsBuffer__AccessType(0x04)
	QPlatformGraphicsBuffer__HWCompositor  = QPlatformGraphicsBuffer__AccessType(0x08)
)

//QPlatformGraphicsBuffer::Origin
type QPlatformGraphicsBuffer__Origin int64

const (
	QPlatformGraphicsBuffer__OriginBottomLeft = QPlatformGraphicsBuffer__Origin(0)
	QPlatformGraphicsBuffer__OriginTopLeft    = QPlatformGraphicsBuffer__Origin(1)
)
