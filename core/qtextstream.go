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

type QTextStreamITF interface {
	QTextStreamPTR() *QTextStream
}

func (p *QTextStream) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextStream) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextStream(ptr QTextStreamITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextStreamPTR().Pointer()
	}
	return nil
}

func QTextStreamFromPointer(ptr unsafe.Pointer) *QTextStream {
	var n = new(QTextStream)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextStream) QTextStreamPTR() *QTextStream {
	return ptr
}

//QTextStream::FieldAlignment
type QTextStream__FieldAlignment int

var (
	QTextStream__AlignLeft            = QTextStream__FieldAlignment(0)
	QTextStream__AlignRight           = QTextStream__FieldAlignment(1)
	QTextStream__AlignCenter          = QTextStream__FieldAlignment(2)
	QTextStream__AlignAccountingStyle = QTextStream__FieldAlignment(3)
)

//QTextStream::NumberFlag
type QTextStream__NumberFlag int

var (
	QTextStream__ShowBase        = QTextStream__NumberFlag(0x1)
	QTextStream__ForcePoint      = QTextStream__NumberFlag(0x2)
	QTextStream__ForceSign       = QTextStream__NumberFlag(0x4)
	QTextStream__UppercaseBase   = QTextStream__NumberFlag(0x8)
	QTextStream__UppercaseDigits = QTextStream__NumberFlag(0x10)
)

//QTextStream::RealNumberNotation
type QTextStream__RealNumberNotation int

var (
	QTextStream__SmartNotation      = QTextStream__RealNumberNotation(0)
	QTextStream__FixedNotation      = QTextStream__RealNumberNotation(1)
	QTextStream__ScientificNotation = QTextStream__RealNumberNotation(2)
)

//QTextStream::Status
type QTextStream__Status int

var (
	QTextStream__Ok              = QTextStream__Status(0)
	QTextStream__ReadPastEnd     = QTextStream__Status(1)
	QTextStream__ReadCorruptData = QTextStream__Status(2)
	QTextStream__WriteFailed     = QTextStream__Status(3)
)

func (ptr *QTextStream) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QTextStream_AtEnd(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextStream) AutoDetectUnicode() bool {
	if ptr.Pointer() != nil {
		return C.QTextStream_AutoDetectUnicode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextStream) Codec() *QTextCodec {
	if ptr.Pointer() != nil {
		return QTextCodecFromPointer(unsafe.Pointer(C.QTextStream_Codec(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextStream) Device() *QIODevice {
	if ptr.Pointer() != nil {
		return QIODeviceFromPointer(unsafe.Pointer(C.QTextStream_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextStream) FieldAlignment() QTextStream__FieldAlignment {
	if ptr.Pointer() != nil {
		return QTextStream__FieldAlignment(C.QTextStream_FieldAlignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextStream) FieldWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QTextStream_FieldWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextStream) Flush() {
	if ptr.Pointer() != nil {
		C.QTextStream_Flush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextStream) GenerateByteOrderMark() bool {
	if ptr.Pointer() != nil {
		return C.QTextStream_GenerateByteOrderMark(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextStream) IntegerBase() int {
	if ptr.Pointer() != nil {
		return int(C.QTextStream_IntegerBase(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextStream) NumberFlags() QTextStream__NumberFlag {
	if ptr.Pointer() != nil {
		return QTextStream__NumberFlag(C.QTextStream_NumberFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextStream) ReadAll() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextStream_ReadAll(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextStream) RealNumberNotation() QTextStream__RealNumberNotation {
	if ptr.Pointer() != nil {
		return QTextStream__RealNumberNotation(C.QTextStream_RealNumberNotation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextStream) RealNumberPrecision() int {
	if ptr.Pointer() != nil {
		return int(C.QTextStream_RealNumberPrecision(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextStream) Reset() {
	if ptr.Pointer() != nil {
		C.QTextStream_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextStream) ResetStatus() {
	if ptr.Pointer() != nil {
		C.QTextStream_ResetStatus(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextStream) SetAutoDetectUnicode(enabled bool) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetAutoDetectUnicode(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTextStream) SetCodec(codec QTextCodecITF) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetCodec(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCodec(codec)))
	}
}

func (ptr *QTextStream) SetCodec2(codecName string) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetCodec2(C.QtObjectPtr(ptr.Pointer()), C.CString(codecName))
	}
}

func (ptr *QTextStream) SetDevice(device QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQIODevice(device)))
	}
}

func (ptr *QTextStream) SetFieldAlignment(mode QTextStream__FieldAlignment) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetFieldAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QTextStream) SetFieldWidth(width int) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetFieldWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width))
	}
}

func (ptr *QTextStream) SetGenerateByteOrderMark(generate bool) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetGenerateByteOrderMark(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(generate)))
	}
}

func (ptr *QTextStream) SetIntegerBase(base int) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetIntegerBase(C.QtObjectPtr(ptr.Pointer()), C.int(base))
	}
}

func (ptr *QTextStream) SetLocale(locale QLocaleITF) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetLocale(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLocale(locale)))
	}
}

func (ptr *QTextStream) SetNumberFlags(flags QTextStream__NumberFlag) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetNumberFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QTextStream) SetPadChar(ch QCharITF) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetPadChar(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(ch)))
	}
}

func (ptr *QTextStream) SetRealNumberNotation(notation QTextStream__RealNumberNotation) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetRealNumberNotation(C.QtObjectPtr(ptr.Pointer()), C.int(notation))
	}
}

func (ptr *QTextStream) SetRealNumberPrecision(precision int) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetRealNumberPrecision(C.QtObjectPtr(ptr.Pointer()), C.int(precision))
	}
}

func (ptr *QTextStream) SetStatus(status QTextStream__Status) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetStatus(C.QtObjectPtr(ptr.Pointer()), C.int(status))
	}
}

func (ptr *QTextStream) SetString(stri string, openMode QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QTextStream_SetString(C.QtObjectPtr(ptr.Pointer()), C.CString(stri), C.int(openMode))
	}
}

func (ptr *QTextStream) SkipWhiteSpace() {
	if ptr.Pointer() != nil {
		C.QTextStream_SkipWhiteSpace(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextStream) Status() QTextStream__Status {
	if ptr.Pointer() != nil {
		return QTextStream__Status(C.QTextStream_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextStream) String() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextStream_String(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextStream) DestroyQTextStream() {
	if ptr.Pointer() != nil {
		C.QTextStream_DestroyQTextStream(C.QtObjectPtr(ptr.Pointer()))
	}
}
