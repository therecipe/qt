package core

//#include "core.h"
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
	for len(n.ObjectNameAbs()) < len("QTextStream_") {
		n.SetObjectNameAbs("QTextStream_" + qt.Identifier())
	}
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
	defer qt.Recovering("QTextStream::atEnd")

	if ptr.Pointer() != nil {
		return C.QTextStream_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextStream) AutoDetectUnicode() bool {
	defer qt.Recovering("QTextStream::autoDetectUnicode")

	if ptr.Pointer() != nil {
		return C.QTextStream_AutoDetectUnicode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextStream) Codec() *QTextCodec {
	defer qt.Recovering("QTextStream::codec")

	if ptr.Pointer() != nil {
		return NewQTextCodecFromPointer(C.QTextStream_Codec(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextStream) Device() *QIODevice {
	defer qt.Recovering("QTextStream::device")

	if ptr.Pointer() != nil {
		return NewQIODeviceFromPointer(C.QTextStream_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextStream) FieldAlignment() QTextStream__FieldAlignment {
	defer qt.Recovering("QTextStream::fieldAlignment")

	if ptr.Pointer() != nil {
		return QTextStream__FieldAlignment(C.QTextStream_FieldAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) FieldWidth() int {
	defer qt.Recovering("QTextStream::fieldWidth")

	if ptr.Pointer() != nil {
		return int(C.QTextStream_FieldWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) Flush() {
	defer qt.Recovering("QTextStream::flush")

	if ptr.Pointer() != nil {
		C.QTextStream_Flush(ptr.Pointer())
	}
}

func (ptr *QTextStream) GenerateByteOrderMark() bool {
	defer qt.Recovering("QTextStream::generateByteOrderMark")

	if ptr.Pointer() != nil {
		return C.QTextStream_GenerateByteOrderMark(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextStream) IntegerBase() int {
	defer qt.Recovering("QTextStream::integerBase")

	if ptr.Pointer() != nil {
		return int(C.QTextStream_IntegerBase(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) NumberFlags() QTextStream__NumberFlag {
	defer qt.Recovering("QTextStream::numberFlags")

	if ptr.Pointer() != nil {
		return QTextStream__NumberFlag(C.QTextStream_NumberFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) Pos() int64 {
	defer qt.Recovering("QTextStream::pos")

	if ptr.Pointer() != nil {
		return int64(C.QTextStream_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) Read(maxlen int64) string {
	defer qt.Recovering("QTextStream::read")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextStream_Read(ptr.Pointer(), C.longlong(maxlen)))
	}
	return ""
}

func (ptr *QTextStream) ReadAll() string {
	defer qt.Recovering("QTextStream::readAll")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextStream_ReadAll(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextStream) ReadLine(maxlen int64) string {
	defer qt.Recovering("QTextStream::readLine")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextStream_ReadLine(ptr.Pointer(), C.longlong(maxlen)))
	}
	return ""
}

func (ptr *QTextStream) ReadLineInto(line string, maxlen int64) bool {
	defer qt.Recovering("QTextStream::readLineInto")

	if ptr.Pointer() != nil {
		return C.QTextStream_ReadLineInto(ptr.Pointer(), C.CString(line), C.longlong(maxlen)) != 0
	}
	return false
}

func (ptr *QTextStream) RealNumberNotation() QTextStream__RealNumberNotation {
	defer qt.Recovering("QTextStream::realNumberNotation")

	if ptr.Pointer() != nil {
		return QTextStream__RealNumberNotation(C.QTextStream_RealNumberNotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) RealNumberPrecision() int {
	defer qt.Recovering("QTextStream::realNumberPrecision")

	if ptr.Pointer() != nil {
		return int(C.QTextStream_RealNumberPrecision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) Reset() {
	defer qt.Recovering("QTextStream::reset")

	if ptr.Pointer() != nil {
		C.QTextStream_Reset(ptr.Pointer())
	}
}

func (ptr *QTextStream) ResetStatus() {
	defer qt.Recovering("QTextStream::resetStatus")

	if ptr.Pointer() != nil {
		C.QTextStream_ResetStatus(ptr.Pointer())
	}
}

func (ptr *QTextStream) Seek(pos int64) bool {
	defer qt.Recovering("QTextStream::seek")

	if ptr.Pointer() != nil {
		return C.QTextStream_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QTextStream) SetAutoDetectUnicode(enabled bool) {
	defer qt.Recovering("QTextStream::setAutoDetectUnicode")

	if ptr.Pointer() != nil {
		C.QTextStream_SetAutoDetectUnicode(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTextStream) SetCodec(codec QTextCodec_ITF) {
	defer qt.Recovering("QTextStream::setCodec")

	if ptr.Pointer() != nil {
		C.QTextStream_SetCodec(ptr.Pointer(), PointerFromQTextCodec(codec))
	}
}

func (ptr *QTextStream) SetCodec2(codecName string) {
	defer qt.Recovering("QTextStream::setCodec")

	if ptr.Pointer() != nil {
		C.QTextStream_SetCodec2(ptr.Pointer(), C.CString(codecName))
	}
}

func (ptr *QTextStream) SetDevice(device QIODevice_ITF) {
	defer qt.Recovering("QTextStream::setDevice")

	if ptr.Pointer() != nil {
		C.QTextStream_SetDevice(ptr.Pointer(), PointerFromQIODevice(device))
	}
}

func (ptr *QTextStream) SetFieldAlignment(mode QTextStream__FieldAlignment) {
	defer qt.Recovering("QTextStream::setFieldAlignment")

	if ptr.Pointer() != nil {
		C.QTextStream_SetFieldAlignment(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QTextStream) SetFieldWidth(width int) {
	defer qt.Recovering("QTextStream::setFieldWidth")

	if ptr.Pointer() != nil {
		C.QTextStream_SetFieldWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QTextStream) SetGenerateByteOrderMark(generate bool) {
	defer qt.Recovering("QTextStream::setGenerateByteOrderMark")

	if ptr.Pointer() != nil {
		C.QTextStream_SetGenerateByteOrderMark(ptr.Pointer(), C.int(qt.GoBoolToInt(generate)))
	}
}

func (ptr *QTextStream) SetIntegerBase(base int) {
	defer qt.Recovering("QTextStream::setIntegerBase")

	if ptr.Pointer() != nil {
		C.QTextStream_SetIntegerBase(ptr.Pointer(), C.int(base))
	}
}

func (ptr *QTextStream) SetLocale(locale QLocale_ITF) {
	defer qt.Recovering("QTextStream::setLocale")

	if ptr.Pointer() != nil {
		C.QTextStream_SetLocale(ptr.Pointer(), PointerFromQLocale(locale))
	}
}

func (ptr *QTextStream) SetNumberFlags(flags QTextStream__NumberFlag) {
	defer qt.Recovering("QTextStream::setNumberFlags")

	if ptr.Pointer() != nil {
		C.QTextStream_SetNumberFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QTextStream) SetPadChar(ch QChar_ITF) {
	defer qt.Recovering("QTextStream::setPadChar")

	if ptr.Pointer() != nil {
		C.QTextStream_SetPadChar(ptr.Pointer(), PointerFromQChar(ch))
	}
}

func (ptr *QTextStream) SetRealNumberNotation(notation QTextStream__RealNumberNotation) {
	defer qt.Recovering("QTextStream::setRealNumberNotation")

	if ptr.Pointer() != nil {
		C.QTextStream_SetRealNumberNotation(ptr.Pointer(), C.int(notation))
	}
}

func (ptr *QTextStream) SetRealNumberPrecision(precision int) {
	defer qt.Recovering("QTextStream::setRealNumberPrecision")

	if ptr.Pointer() != nil {
		C.QTextStream_SetRealNumberPrecision(ptr.Pointer(), C.int(precision))
	}
}

func (ptr *QTextStream) SetStatus(status QTextStream__Status) {
	defer qt.Recovering("QTextStream::setStatus")

	if ptr.Pointer() != nil {
		C.QTextStream_SetStatus(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QTextStream) SetString(stri string, openMode QIODevice__OpenModeFlag) {
	defer qt.Recovering("QTextStream::setString")

	if ptr.Pointer() != nil {
		C.QTextStream_SetString(ptr.Pointer(), C.CString(stri), C.int(openMode))
	}
}

func (ptr *QTextStream) SkipWhiteSpace() {
	defer qt.Recovering("QTextStream::skipWhiteSpace")

	if ptr.Pointer() != nil {
		C.QTextStream_SkipWhiteSpace(ptr.Pointer())
	}
}

func (ptr *QTextStream) Status() QTextStream__Status {
	defer qt.Recovering("QTextStream::status")

	if ptr.Pointer() != nil {
		return QTextStream__Status(C.QTextStream_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) String() string {
	defer qt.Recovering("QTextStream::string")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextStream_String(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextStream) DestroyQTextStream() {
	defer qt.Recovering("QTextStream::~QTextStream")

	if ptr.Pointer() != nil {
		C.QTextStream_DestroyQTextStream(ptr.Pointer())
	}
}

func (ptr *QTextStream) ObjectNameAbs() string {
	defer qt.Recovering("QTextStream::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextStream_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextStream) SetObjectNameAbs(name string) {
	defer qt.Recovering("QTextStream::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QTextStream_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
