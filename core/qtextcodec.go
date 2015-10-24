package core

//#include "qtextcodec.h"
import "C"
import (
	"unsafe"
)

type QTextCodec struct {
	ptr unsafe.Pointer
}

type QTextCodecITF interface {
	QTextCodecPTR() *QTextCodec
}

func (p *QTextCodec) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextCodec) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextCodec(ptr QTextCodecITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextCodecPTR().Pointer()
	}
	return nil
}

func QTextCodecFromPointer(ptr unsafe.Pointer) *QTextCodec {
	var n = new(QTextCodec)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextCodec) QTextCodecPTR() *QTextCodec {
	return ptr
}

//QTextCodec::ConversionFlag
type QTextCodec__ConversionFlag int

var (
	QTextCodec__DefaultConversion    = QTextCodec__ConversionFlag(0)
	QTextCodec__ConvertInvalidToNull = QTextCodec__ConversionFlag(0x80000000)
	QTextCodec__IgnoreHeader         = QTextCodec__ConversionFlag(0x1)
	QTextCodec__FreeFunction         = QTextCodec__ConversionFlag(0x2)
)

func (ptr *QTextCodec) CanEncode(ch QCharITF) bool {
	if ptr.Pointer() != nil {
		return C.QTextCodec_CanEncode(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(ch))) != 0
	}
	return false
}

func (ptr *QTextCodec) CanEncode2(s string) bool {
	if ptr.Pointer() != nil {
		return C.QTextCodec_CanEncode2(C.QtObjectPtr(ptr.Pointer()), C.CString(s)) != 0
	}
	return false
}

func QTextCodec_CodecForHtml2(ba QByteArrayITF) *QTextCodec {
	return QTextCodecFromPointer(unsafe.Pointer(C.QTextCodec_QTextCodec_CodecForHtml2(C.QtObjectPtr(PointerFromQByteArray(ba)))))
}

func QTextCodec_CodecForHtml(ba QByteArrayITF, defaultCodec QTextCodecITF) *QTextCodec {
	return QTextCodecFromPointer(unsafe.Pointer(C.QTextCodec_QTextCodec_CodecForHtml(C.QtObjectPtr(PointerFromQByteArray(ba)), C.QtObjectPtr(PointerFromQTextCodec(defaultCodec)))))
}

func QTextCodec_CodecForLocale() *QTextCodec {
	return QTextCodecFromPointer(unsafe.Pointer(C.QTextCodec_QTextCodec_CodecForLocale()))
}

func QTextCodec_CodecForMib(mib int) *QTextCodec {
	return QTextCodecFromPointer(unsafe.Pointer(C.QTextCodec_QTextCodec_CodecForMib(C.int(mib))))
}

func QTextCodec_CodecForName(name QByteArrayITF) *QTextCodec {
	return QTextCodecFromPointer(unsafe.Pointer(C.QTextCodec_QTextCodec_CodecForName(C.QtObjectPtr(PointerFromQByteArray(name)))))
}

func QTextCodec_CodecForName2(name string) *QTextCodec {
	return QTextCodecFromPointer(unsafe.Pointer(C.QTextCodec_QTextCodec_CodecForName2(C.CString(name))))
}

func QTextCodec_CodecForUtfText2(ba QByteArrayITF) *QTextCodec {
	return QTextCodecFromPointer(unsafe.Pointer(C.QTextCodec_QTextCodec_CodecForUtfText2(C.QtObjectPtr(PointerFromQByteArray(ba)))))
}

func QTextCodec_CodecForUtfText(ba QByteArrayITF, defaultCodec QTextCodecITF) *QTextCodec {
	return QTextCodecFromPointer(unsafe.Pointer(C.QTextCodec_QTextCodec_CodecForUtfText(C.QtObjectPtr(PointerFromQByteArray(ba)), C.QtObjectPtr(PointerFromQTextCodec(defaultCodec)))))
}

func (ptr *QTextCodec) MakeDecoder(flags QTextCodec__ConversionFlag) *QTextDecoder {
	if ptr.Pointer() != nil {
		return QTextDecoderFromPointer(unsafe.Pointer(C.QTextCodec_MakeDecoder(C.QtObjectPtr(ptr.Pointer()), C.int(flags))))
	}
	return nil
}

func (ptr *QTextCodec) MakeEncoder(flags QTextCodec__ConversionFlag) *QTextEncoder {
	if ptr.Pointer() != nil {
		return QTextEncoderFromPointer(unsafe.Pointer(C.QTextCodec_MakeEncoder(C.QtObjectPtr(ptr.Pointer()), C.int(flags))))
	}
	return nil
}

func (ptr *QTextCodec) MibEnum() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCodec_MibEnum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QTextCodec_SetCodecForLocale(c QTextCodecITF) {
	C.QTextCodec_QTextCodec_SetCodecForLocale(C.QtObjectPtr(PointerFromQTextCodec(c)))
}
