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

type QPlatformGraphicsBufferITF interface {
	core.QObjectITF
	QPlatformGraphicsBufferPTR() *QPlatformGraphicsBuffer
}

func PointerFromQPlatformGraphicsBuffer(ptr QPlatformGraphicsBufferITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlatformGraphicsBufferPTR().Pointer()
	}
	return nil
}

func QPlatformGraphicsBufferFromPointer(ptr unsafe.Pointer) *QPlatformGraphicsBuffer {
	var n = new(QPlatformGraphicsBuffer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlatformGraphicsBuffer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlatformGraphicsBuffer) QPlatformGraphicsBufferPTR() *QPlatformGraphicsBuffer {
	return ptr
}

//QPlatformGraphicsBuffer::AccessType
type QPlatformGraphicsBuffer__AccessType int

var (
	QPlatformGraphicsBuffer__None          = QPlatformGraphicsBuffer__AccessType(0x00)
	QPlatformGraphicsBuffer__SWReadAccess  = QPlatformGraphicsBuffer__AccessType(0x01)
	QPlatformGraphicsBuffer__SWWriteAccess = QPlatformGraphicsBuffer__AccessType(0x02)
	QPlatformGraphicsBuffer__TextureAccess = QPlatformGraphicsBuffer__AccessType(0x04)
	QPlatformGraphicsBuffer__HWCompositor  = QPlatformGraphicsBuffer__AccessType(0x08)
)

//QPlatformGraphicsBuffer::Origin
type QPlatformGraphicsBuffer__Origin int

var (
	QPlatformGraphicsBuffer__OriginBottomLeft = QPlatformGraphicsBuffer__Origin(0)
	QPlatformGraphicsBuffer__OriginTopLeft    = QPlatformGraphicsBuffer__Origin(1)
)
