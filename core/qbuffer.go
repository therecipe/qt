package core

//#include "qbuffer.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QBuffer struct {
	QIODevice
}

type QBuffer_ITF interface {
	QIODevice_ITF
	QBuffer_PTR() *QBuffer
}

func PointerFromQBuffer(ptr QBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBuffer_PTR().Pointer()
	}
	return nil
}

func NewQBufferFromPointer(ptr unsafe.Pointer) *QBuffer {
	var n = new(QBuffer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QBuffer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBuffer) QBuffer_PTR() *QBuffer {
	return ptr
}

func NewQBuffer2(byteArray QByteArray_ITF, parent QObject_ITF) *QBuffer {
	return NewQBufferFromPointer(C.QBuffer_NewQBuffer2(PointerFromQByteArray(byteArray), PointerFromQObject(parent)))
}

func NewQBuffer(parent QObject_ITF) *QBuffer {
	return NewQBufferFromPointer(C.QBuffer_NewQBuffer(PointerFromQObject(parent)))
}

func (ptr *QBuffer) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QBuffer_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBuffer) Buffer() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Buffer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) Buffer2() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Buffer2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QBuffer_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBuffer) Close() {
	if ptr.Pointer() != nil {
		C.QBuffer_Close(ptr.Pointer())
	}
}

func (ptr *QBuffer) Data() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) Open(flags QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QBuffer_Open(ptr.Pointer(), C.int(flags)) != 0
	}
	return false
}

func (ptr *QBuffer) SetBuffer(byteArray QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBuffer_SetBuffer(ptr.Pointer(), PointerFromQByteArray(byteArray))
	}
}

func (ptr *QBuffer) SetData(data QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBuffer_SetData(ptr.Pointer(), PointerFromQByteArray(data))
	}
}

func (ptr *QBuffer) SetData2(data string, size int) {
	if ptr.Pointer() != nil {
		C.QBuffer_SetData2(ptr.Pointer(), C.CString(data), C.int(size))
	}
}

func (ptr *QBuffer) DestroyQBuffer() {
	if ptr.Pointer() != nil {
		C.QBuffer_DestroyQBuffer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
