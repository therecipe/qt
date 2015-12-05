package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	for len(n.ObjectName()) < len("QBuffer_") {
		n.SetObjectName("QBuffer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBuffer) QBuffer_PTR() *QBuffer {
	return ptr
}

func NewQBuffer2(byteArray QByteArray_ITF, parent QObject_ITF) *QBuffer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::QBuffer")
		}
	}()

	return NewQBufferFromPointer(C.QBuffer_NewQBuffer2(PointerFromQByteArray(byteArray), PointerFromQObject(parent)))
}

func NewQBuffer(parent QObject_ITF) *QBuffer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::QBuffer")
		}
	}()

	return NewQBufferFromPointer(C.QBuffer_NewQBuffer(PointerFromQObject(parent)))
}

func (ptr *QBuffer) AtEnd() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::atEnd")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBuffer_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBuffer) Buffer() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::buffer")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Buffer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) Buffer2() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::buffer")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Buffer2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) CanReadLine() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::canReadLine")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBuffer_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBuffer) Close() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::close")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBuffer_Close(ptr.Pointer())
	}
}

func (ptr *QBuffer) Data() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::data")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) Open(flags QIODevice__OpenModeFlag) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::open")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBuffer_Open(ptr.Pointer(), C.int(flags)) != 0
	}
	return false
}

func (ptr *QBuffer) SetBuffer(byteArray QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::setBuffer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBuffer_SetBuffer(ptr.Pointer(), PointerFromQByteArray(byteArray))
	}
}

func (ptr *QBuffer) SetData(data QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::setData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBuffer_SetData(ptr.Pointer(), PointerFromQByteArray(data))
	}
}

func (ptr *QBuffer) SetData2(data string, size int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::setData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBuffer_SetData2(ptr.Pointer(), C.CString(data), C.int(size))
	}
}

func (ptr *QBuffer) DestroyQBuffer() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBuffer::~QBuffer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBuffer_DestroyQBuffer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
