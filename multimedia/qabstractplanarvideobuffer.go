package multimedia

//#include "qabstractplanarvideobuffer.h"
import "C"
import (
	"unsafe"
)

type QAbstractPlanarVideoBuffer struct {
	QAbstractVideoBuffer
}

type QAbstractPlanarVideoBufferITF interface {
	QAbstractVideoBufferITF
	QAbstractPlanarVideoBufferPTR() *QAbstractPlanarVideoBuffer
}

func PointerFromQAbstractPlanarVideoBuffer(ptr QAbstractPlanarVideoBufferITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractPlanarVideoBufferPTR().Pointer()
	}
	return nil
}

func QAbstractPlanarVideoBufferFromPointer(ptr unsafe.Pointer) *QAbstractPlanarVideoBuffer {
	var n = new(QAbstractPlanarVideoBuffer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractPlanarVideoBuffer) QAbstractPlanarVideoBufferPTR() *QAbstractPlanarVideoBuffer {
	return ptr
}

func (ptr *QAbstractPlanarVideoBuffer) DestroyQAbstractPlanarVideoBuffer() {
	if ptr.Pointer() != nil {
		C.QAbstractPlanarVideoBuffer_DestroyQAbstractPlanarVideoBuffer(C.QtObjectPtr(ptr.Pointer()))
	}
}
