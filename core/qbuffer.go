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

type QBufferITF interface {
	QIODeviceITF
	QBufferPTR() *QBuffer
}

func PointerFromQBuffer(ptr QBufferITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBufferPTR().Pointer()
	}
	return nil
}

func QBufferFromPointer(ptr unsafe.Pointer) *QBuffer {
	var n = new(QBuffer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QBuffer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBuffer) QBufferPTR() *QBuffer {
	return ptr
}

func NewQBuffer2(byteArray QByteArrayITF, parent QObjectITF) *QBuffer {
	return QBufferFromPointer(unsafe.Pointer(C.QBuffer_NewQBuffer2(C.QtObjectPtr(PointerFromQByteArray(byteArray)), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQBuffer(parent QObjectITF) *QBuffer {
	return QBufferFromPointer(unsafe.Pointer(C.QBuffer_NewQBuffer(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QBuffer) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QBuffer_AtEnd(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBuffer) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QBuffer_CanReadLine(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBuffer) Close() {
	if ptr.Pointer() != nil {
		C.QBuffer_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBuffer) Open(flags QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QBuffer_Open(C.QtObjectPtr(ptr.Pointer()), C.int(flags)) != 0
	}
	return false
}

func (ptr *QBuffer) SetBuffer(byteArray QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QBuffer_SetBuffer(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(byteArray)))
	}
}

func (ptr *QBuffer) SetData(data QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QBuffer_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(data)))
	}
}

func (ptr *QBuffer) SetData2(data string, size int) {
	if ptr.Pointer() != nil {
		C.QBuffer_SetData2(C.QtObjectPtr(ptr.Pointer()), C.CString(data), C.int(size))
	}
}

func (ptr *QBuffer) DestroyQBuffer() {
	if ptr.Pointer() != nil {
		C.QBuffer_DestroyQBuffer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
