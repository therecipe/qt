package core

//#include "qtextstream.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextStream struct {
	ptr unsafe.Pointer
}

type QTextStream_ITF interface {
	QTextStream_PTR() *QTextStream
}

func (p *QTextStream) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextStream) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextStream(ptr QTextStream_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextStream_PTR().Pointer()
	}
	return nil
}

func NewQTextStreamFromPointer(ptr unsafe.Pointer) *QTextStream {
	var n = new(QTextStream)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextStream) QTextStream_PTR() *QTextStream {
	return ptr
}

//QTextStream::FieldAlignment
type QTextStream__FieldAlignment int64

const (
	QTextStream__AlignLeft            = QTextStream__FieldAlignment(0)
	QTextStream__AlignRight           = QTextStream__FieldAlignment(1)
	QTextStream__AlignCenter          = QTextStream__FieldAlignment(2)
	QTextStream__AlignAccountingStyle = QTextStream__FieldAlignment(3)
)

//QTextStream::NumberFlag
type QTextStream__NumberFlag int64

const (
	QTextStream__ShowBase        = QTextStream__NumberFlag(0x1)
	QTextStream__ForcePoint      = QTextStream__NumberFlag(0x2)
	QTextStream__ForceSign       = QTextStream__NumberFlag(0x4)
	QTextStream__UppercaseBase   = QTextStream__NumberFlag(0x8)
	QTextStream__UppercaseDigits = QTextStream__NumberFlag(0x10)
)

//QTextStream::RealNumberNotation
type QTextStream__RealNumberNotation int64

const (
	QTextStream__SmartNotation      = QTextStream__RealNumberNotation(0)
	QTextStream__FixedNotation      = QTextStream__RealNumberNotation(1)
	QTextStream__ScientificNotation = QTextStream__RealNumberNotation(2)
)

//QTextStream::Status
type QTextStream__Status int64

const (
	QTextStream__Ok              = QTextStream__Status(0)
	QTextStream__ReadPastEnd     = QTextStream__Status(1)
	QTextStream__ReadCorruptData = QTextStream__Status(2)
	QTextStream__WriteFailed     = QTextStream__Status(3)
)

func (ptr *QTextStream) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QTextStream_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextStream) AutoDetectUnicode() bool {
	if ptr.Pointer() != nil {
		return C.QTextStream_AutoDetectUnicode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextStream) Codec() *QTextCodec {
	if ptr.Pointer() != nil {
		return NewQTextCodecFromPointer(C.QTextStream_Codec(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextStream) Device() *QIODevice {
	if ptr.Pointer() != nil {
		return NewQIODeviceFromPointer(C.QTextStream_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextStream) FieldAlignment() QTextStream__FieldAlignment {
	if ptr.Pointer() != nil {
		return QTextStream__FieldAlignment(C.QTextStream_FieldAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) FieldWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QTextStream_FieldWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) Flush() {
	if ptr.Pointer() != nil {
		C.QTextStream_Flush(ptr.Pointer())
	}
}

func (ptr *QTextStream) GenerateByteOrderMark() bool {
	if ptr.Pointer() != nil {
		return C.QTextStream_GenerateByteOrderMark(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextStream) IntegerBase() int {
	if ptr.Pointer() != nil {
		return int(C.QTextStream_IntegerBase(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) NumberFlags() QTextStream__NumberFlag {
	if ptr.Pointer() != nil {
		return QTextStream__NumberFlag(C.QTextStream_NumberFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) ReadAll() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextStream_ReadAll(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextStream) RealNumberNotation() QTextStream__RealNumberNotation {
	if ptr.Pointer() != nil {
		return QTextStream__RealNumberNotation(C.QTextStream_RealNumberNotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) RealNumberPrecision() int {
	if ptr.Pointer() != nil {
		return int(C.QTextStream_RealNumberPrecision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) Reset() {
	if ptr.Pointer() != nil {
		C.QTextStream_Reset(ptr.Pointer())
	}
}

func (ptr *QTextStream) ResetStatus() {
	if ptr.Pointer() != nil {
		C.QTextStream_ResetStatus(ptr.Pointer())
	}
}

func (ptr *QTextStream) SetAutoDetectUnicode(enabled bool) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetAutoDetectUnicode(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTextStream) SetCodec(codec QTextCodec_ITF) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetCodec(ptr.Pointer(), PointerFromQTextCodec(codec))
	}
}

func (ptr *QTextStream) SetCodec2(codecName string) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetCodec2(ptr.Pointer(), C.CString(codecName))
	}
}

func (ptr *QTextStream) SetDevice(device QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetDevice(ptr.Pointer(), PointerFromQIODevice(device))
	}
}

func (ptr *QTextStream) SetFieldAlignment(mode QTextStream__FieldAlignment) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetFieldAlignment(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QTextStream) SetFieldWidth(width int) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetFieldWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QTextStream) SetGenerateByteOrderMark(generate bool) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetGenerateByteOrderMark(ptr.Pointer(), C.int(qt.GoBoolToInt(generate)))
	}
}

func (ptr *QTextStream) SetIntegerBase(base int) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetIntegerBase(ptr.Pointer(), C.int(base))
	}
}

func (ptr *QTextStream) SetLocale(locale QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetLocale(ptr.Pointer(), PointerFromQLocale(locale))
	}
}

func (ptr *QTextStream) SetNumberFlags(flags QTextStream__NumberFlag) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetNumberFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QTextStream) SetPadChar(ch QChar_ITF) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetPadChar(ptr.Pointer(), PointerFromQChar(ch))
	}
}

func (ptr *QTextStream) SetRealNumberNotation(notation QTextStream__RealNumberNotation) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetRealNumberNotation(ptr.Pointer(), C.int(notation))
	}
}

func (ptr *QTextStream) SetRealNumberPrecision(precision int) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetRealNumberPrecision(ptr.Pointer(), C.int(precision))
	}
}

func (ptr *QTextStream) SetStatus(status QTextStream__Status) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetStatus(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QTextStream) SetString(stri string, openMode QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetString(ptr.Pointer(), C.CString(stri), C.int(openMode))
	}
}

func (ptr *QTextStream) SkipWhiteSpace() {
	if ptr.Pointer() != nil {
		C.QTextStream_SkipWhiteSpace(ptr.Pointer())
	}
}

func (ptr *QTextStream) Status() QTextStream__Status {
	if ptr.Pointer() != nil {
		return QTextStream__Status(C.QTextStream_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) String() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextStream_String(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextStream) DestroyQTextStream() {
	if ptr.Pointer() != nil {
		C.QTextStream_DestroyQTextStream(ptr.Pointer())
	}
}
