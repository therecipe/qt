package multimedia

//#include "multimedia.h"
import "C"
import (
	"log"
	"unsafe"
)

type QAbstractPlanarVideoBuffer struct {
	QAbstractVideoBuffer
}

type QAbstractPlanarVideoBuffer_ITF interface {
	QAbstractVideoBuffer_ITF
	QAbstractPlanarVideoBuffer_PTR() *QAbstractPlanarVideoBuffer
}

func PointerFromQAbstractPlanarVideoBuffer(ptr QAbstractPlanarVideoBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractPlanarVideoBuffer_PTR().Pointer()
	}
	return nil
}

func NewQAbstractPlanarVideoBufferFromPointer(ptr unsafe.Pointer) *QAbstractPlanarVideoBuffer {
	var n = new(QAbstractPlanarVideoBuffer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractPlanarVideoBuffer) QAbstractPlanarVideoBuffer_PTR() *QAbstractPlanarVideoBuffer {
	return ptr
}

func (ptr *QAbstractPlanarVideoBuffer) DestroyQAbstractPlanarVideoBuffer() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractPlanarVideoBuffer::~QAbstractPlanarVideoBuffer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractPlanarVideoBuffer_DestroyQAbstractPlanarVideoBuffer(ptr.Pointer())
	}
}
