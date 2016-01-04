package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QAbstractPlanarVideoBuffer_") {
		n.SetObjectNameAbs("QAbstractPlanarVideoBuffer_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractPlanarVideoBuffer) QAbstractPlanarVideoBuffer_PTR() *QAbstractPlanarVideoBuffer {
	return ptr
}

func (ptr *QAbstractPlanarVideoBuffer) DestroyQAbstractPlanarVideoBuffer() {
	defer qt.Recovering("QAbstractPlanarVideoBuffer::~QAbstractPlanarVideoBuffer")

	if ptr.Pointer() != nil {
		C.QAbstractPlanarVideoBuffer_DestroyQAbstractPlanarVideoBuffer(ptr.Pointer())
	}
}

func (ptr *QAbstractPlanarVideoBuffer) ObjectNameAbs() string {
	defer qt.Recovering("QAbstractPlanarVideoBuffer::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractPlanarVideoBuffer_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractPlanarVideoBuffer) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAbstractPlanarVideoBuffer::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAbstractPlanarVideoBuffer_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAbstractPlanarVideoBuffer) ConnectRelease(f func()) {
	defer qt.Recovering("connect QAbstractPlanarVideoBuffer::release")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "release", f)
	}
}

func (ptr *QAbstractPlanarVideoBuffer) DisconnectRelease() {
	defer qt.Recovering("disconnect QAbstractPlanarVideoBuffer::release")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "release")
	}
}

//export callbackQAbstractPlanarVideoBufferRelease
func callbackQAbstractPlanarVideoBufferRelease(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractPlanarVideoBuffer::release")

	if signal := qt.GetSignal(C.GoString(ptrName), "release"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractPlanarVideoBufferFromPointer(ptr).ReleaseDefault()
	}
}

func (ptr *QAbstractPlanarVideoBuffer) Release() {
	defer qt.Recovering("QAbstractPlanarVideoBuffer::release")

	if ptr.Pointer() != nil {
		C.QAbstractPlanarVideoBuffer_Release(ptr.Pointer())
	}
}

func (ptr *QAbstractPlanarVideoBuffer) ReleaseDefault() {
	defer qt.Recovering("QAbstractPlanarVideoBuffer::release")

	if ptr.Pointer() != nil {
		C.QAbstractPlanarVideoBuffer_ReleaseDefault(ptr.Pointer())
	}
}
