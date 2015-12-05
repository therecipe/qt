package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QDataStream struct {
	ptr unsafe.Pointer
}

type QDataStream_ITF interface {
	QDataStream_PTR() *QDataStream
}

func (p *QDataStream) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDataStream) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDataStream(ptr QDataStream_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDataStream_PTR().Pointer()
	}
	return nil
}

func NewQDataStreamFromPointer(ptr unsafe.Pointer) *QDataStream {
	var n = new(QDataStream)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDataStream) QDataStream_PTR() *QDataStream {
	return ptr
}

//QDataStream::ByteOrder
type QDataStream__ByteOrder int64

const (
	QDataStream__BigEndian    = QDataStream__ByteOrder(QSysInfo__BigEndian)
	QDataStream__LittleEndian = QDataStream__ByteOrder(QSysInfo__LittleEndian)
)

//QDataStream::FloatingPointPrecision
type QDataStream__FloatingPointPrecision int64

const (
	QDataStream__SinglePrecision = QDataStream__FloatingPointPrecision(0)
	QDataStream__DoublePrecision = QDataStream__FloatingPointPrecision(1)
)

//QDataStream::Status
type QDataStream__Status int64

const (
	QDataStream__Ok              = QDataStream__Status(0)
	QDataStream__ReadPastEnd     = QDataStream__Status(1)
	QDataStream__ReadCorruptData = QDataStream__Status(2)
	QDataStream__WriteFailed     = QDataStream__Status(3)
)

//QDataStream::Version
type QDataStream__Version int64

const (
	QDataStream__Qt_1_0                    = QDataStream__Version(1)
	QDataStream__Qt_2_0                    = QDataStream__Version(2)
	QDataStream__Qt_2_1                    = QDataStream__Version(3)
	QDataStream__Qt_3_0                    = QDataStream__Version(4)
	QDataStream__Qt_3_1                    = QDataStream__Version(5)
	QDataStream__Qt_3_3                    = QDataStream__Version(6)
	QDataStream__Qt_4_0                    = QDataStream__Version(7)
	QDataStream__Qt_4_1                    = QDataStream__Version(QDataStream__Qt_4_0)
	QDataStream__Qt_4_2                    = QDataStream__Version(8)
	QDataStream__Qt_4_3                    = QDataStream__Version(9)
	QDataStream__Qt_4_4                    = QDataStream__Version(10)
	QDataStream__Qt_4_5                    = QDataStream__Version(11)
	QDataStream__Qt_4_6                    = QDataStream__Version(12)
	QDataStream__Qt_4_7                    = QDataStream__Version(QDataStream__Qt_4_6)
	QDataStream__Qt_4_8                    = QDataStream__Version(QDataStream__Qt_4_7)
	QDataStream__Qt_4_9                    = QDataStream__Version(QDataStream__Qt_4_8)
	QDataStream__Qt_5_0                    = QDataStream__Version(13)
	QDataStream__Qt_5_1                    = QDataStream__Version(14)
	QDataStream__Qt_5_2                    = QDataStream__Version(15)
	QDataStream__Qt_5_3                    = QDataStream__Version(QDataStream__Qt_5_2)
	QDataStream__Qt_5_4                    = QDataStream__Version(16)
	QDataStream__Qt_5_5                    = QDataStream__Version(QDataStream__Qt_5_4)
	QDataStream__Qt_DefaultCompiledVersion = QDataStream__Version(QDataStream__Qt_5_5)
)

func NewQDataStream3(a QByteArray_ITF, mode QIODevice__OpenModeFlag) *QDataStream {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::QDataStream")
		}
	}()

	return NewQDataStreamFromPointer(C.QDataStream_NewQDataStream3(PointerFromQByteArray(a), C.int(mode)))
}

func (ptr *QDataStream) AtEnd() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::atEnd")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDataStream_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func NewQDataStream() *QDataStream {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::QDataStream")
		}
	}()

	return NewQDataStreamFromPointer(C.QDataStream_NewQDataStream())
}

func NewQDataStream2(d QIODevice_ITF) *QDataStream {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::QDataStream")
		}
	}()

	return NewQDataStreamFromPointer(C.QDataStream_NewQDataStream2(PointerFromQIODevice(d)))
}

func NewQDataStream4(a QByteArray_ITF) *QDataStream {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::QDataStream")
		}
	}()

	return NewQDataStreamFromPointer(C.QDataStream_NewQDataStream4(PointerFromQByteArray(a)))
}

func (ptr *QDataStream) ByteOrder() QDataStream__ByteOrder {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::byteOrder")
		}
	}()

	if ptr.Pointer() != nil {
		return QDataStream__ByteOrder(C.QDataStream_ByteOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataStream) Device() *QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::device")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQIODeviceFromPointer(C.QDataStream_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDataStream) FloatingPointPrecision() QDataStream__FloatingPointPrecision {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::floatingPointPrecision")
		}
	}()

	if ptr.Pointer() != nil {
		return QDataStream__FloatingPointPrecision(C.QDataStream_FloatingPointPrecision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataStream) ReadRawData(s string, len int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::readRawData")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDataStream_ReadRawData(ptr.Pointer(), C.CString(s), C.int(len)))
	}
	return 0
}

func (ptr *QDataStream) ResetStatus() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::resetStatus")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataStream_ResetStatus(ptr.Pointer())
	}
}

func (ptr *QDataStream) SetByteOrder(bo QDataStream__ByteOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::setByteOrder")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataStream_SetByteOrder(ptr.Pointer(), C.int(bo))
	}
}

func (ptr *QDataStream) SetDevice(d QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::setDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataStream_SetDevice(ptr.Pointer(), PointerFromQIODevice(d))
	}
}

func (ptr *QDataStream) SetFloatingPointPrecision(precision QDataStream__FloatingPointPrecision) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::setFloatingPointPrecision")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataStream_SetFloatingPointPrecision(ptr.Pointer(), C.int(precision))
	}
}

func (ptr *QDataStream) SetStatus(status QDataStream__Status) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::setStatus")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataStream_SetStatus(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QDataStream) SetVersion(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::setVersion")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataStream_SetVersion(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QDataStream) SkipRawData(len int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::skipRawData")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDataStream_SkipRawData(ptr.Pointer(), C.int(len)))
	}
	return 0
}

func (ptr *QDataStream) Status() QDataStream__Status {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::status")
		}
	}()

	if ptr.Pointer() != nil {
		return QDataStream__Status(C.QDataStream_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataStream) Version() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::version")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDataStream_Version(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataStream) WriteRawData(s string, len int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::writeRawData")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDataStream_WriteRawData(ptr.Pointer(), C.CString(s), C.int(len)))
	}
	return 0
}

func (ptr *QDataStream) DestroyQDataStream() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDataStream::~QDataStream")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDataStream_DestroyQDataStream(ptr.Pointer())
	}
}
