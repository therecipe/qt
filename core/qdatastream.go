package core

//#include "qdatastream.h"
import "C"
import (
	"unsafe"
)

type QDataStream struct {
	ptr unsafe.Pointer
}

type QDataStreamITF interface {
	QDataStreamPTR() *QDataStream
}

func (p *QDataStream) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDataStream) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDataStream(ptr QDataStreamITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDataStreamPTR().Pointer()
	}
	return nil
}

func QDataStreamFromPointer(ptr unsafe.Pointer) *QDataStream {
	var n = new(QDataStream)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDataStream) QDataStreamPTR() *QDataStream {
	return ptr
}

//QDataStream::ByteOrder
type QDataStream__ByteOrder int

var (
	QDataStream__BigEndian    = QDataStream__ByteOrder(QSysInfo__BigEndian)
	QDataStream__LittleEndian = QDataStream__ByteOrder(QSysInfo__LittleEndian)
)

//QDataStream::FloatingPointPrecision
type QDataStream__FloatingPointPrecision int

var (
	QDataStream__SinglePrecision = QDataStream__FloatingPointPrecision(0)
	QDataStream__DoublePrecision = QDataStream__FloatingPointPrecision(1)
)

//QDataStream::Status
type QDataStream__Status int

var (
	QDataStream__Ok              = QDataStream__Status(0)
	QDataStream__ReadPastEnd     = QDataStream__Status(1)
	QDataStream__ReadCorruptData = QDataStream__Status(2)
	QDataStream__WriteFailed     = QDataStream__Status(3)
)

//QDataStream::Version
type QDataStream__Version int

var (
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

func NewQDataStream3(a QByteArrayITF, mode QIODevice__OpenModeFlag) *QDataStream {
	return QDataStreamFromPointer(unsafe.Pointer(C.QDataStream_NewQDataStream3(C.QtObjectPtr(PointerFromQByteArray(a)), C.int(mode))))
}

func (ptr *QDataStream) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QDataStream_AtEnd(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQDataStream() *QDataStream {
	return QDataStreamFromPointer(unsafe.Pointer(C.QDataStream_NewQDataStream()))
}

func NewQDataStream2(d QIODeviceITF) *QDataStream {
	return QDataStreamFromPointer(unsafe.Pointer(C.QDataStream_NewQDataStream2(C.QtObjectPtr(PointerFromQIODevice(d)))))
}

func NewQDataStream4(a QByteArrayITF) *QDataStream {
	return QDataStreamFromPointer(unsafe.Pointer(C.QDataStream_NewQDataStream4(C.QtObjectPtr(PointerFromQByteArray(a)))))
}

func (ptr *QDataStream) ByteOrder() QDataStream__ByteOrder {
	if ptr.Pointer() != nil {
		return QDataStream__ByteOrder(C.QDataStream_ByteOrder(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDataStream) Device() *QIODevice {
	if ptr.Pointer() != nil {
		return QIODeviceFromPointer(unsafe.Pointer(C.QDataStream_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDataStream) FloatingPointPrecision() QDataStream__FloatingPointPrecision {
	if ptr.Pointer() != nil {
		return QDataStream__FloatingPointPrecision(C.QDataStream_FloatingPointPrecision(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDataStream) ReadRawData(s string, len int) int {
	if ptr.Pointer() != nil {
		return int(C.QDataStream_ReadRawData(C.QtObjectPtr(ptr.Pointer()), C.CString(s), C.int(len)))
	}
	return 0
}

func (ptr *QDataStream) ResetStatus() {
	if ptr.Pointer() != nil {
		C.QDataStream_ResetStatus(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDataStream) SetByteOrder(bo QDataStream__ByteOrder) {
	if ptr.Pointer() != nil {
		C.QDataStream_SetByteOrder(C.QtObjectPtr(ptr.Pointer()), C.int(bo))
	}
}

func (ptr *QDataStream) SetDevice(d QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QDataStream_SetDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQIODevice(d)))
	}
}

func (ptr *QDataStream) SetFloatingPointPrecision(precision QDataStream__FloatingPointPrecision) {
	if ptr.Pointer() != nil {
		C.QDataStream_SetFloatingPointPrecision(C.QtObjectPtr(ptr.Pointer()), C.int(precision))
	}
}

func (ptr *QDataStream) SetStatus(status QDataStream__Status) {
	if ptr.Pointer() != nil {
		C.QDataStream_SetStatus(C.QtObjectPtr(ptr.Pointer()), C.int(status))
	}
}

func (ptr *QDataStream) SetVersion(v int) {
	if ptr.Pointer() != nil {
		C.QDataStream_SetVersion(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QDataStream) SkipRawData(len int) int {
	if ptr.Pointer() != nil {
		return int(C.QDataStream_SkipRawData(C.QtObjectPtr(ptr.Pointer()), C.int(len)))
	}
	return 0
}

func (ptr *QDataStream) Status() QDataStream__Status {
	if ptr.Pointer() != nil {
		return QDataStream__Status(C.QDataStream_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDataStream) Version() int {
	if ptr.Pointer() != nil {
		return int(C.QDataStream_Version(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDataStream) WriteRawData(s string, len int) int {
	if ptr.Pointer() != nil {
		return int(C.QDataStream_WriteRawData(C.QtObjectPtr(ptr.Pointer()), C.CString(s), C.int(len)))
	}
	return 0
}

func (ptr *QDataStream) DestroyQDataStream() {
	if ptr.Pointer() != nil {
		C.QDataStream_DestroyQDataStream(C.QtObjectPtr(ptr.Pointer()))
	}
}
