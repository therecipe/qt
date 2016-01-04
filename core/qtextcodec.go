package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextCodec struct {
	ptr unsafe.Pointer
}

type QTextCodec_ITF interface {
	QTextCodec_PTR() *QTextCodec
}

func (p *QTextCodec) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextCodec) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextCodec(ptr QTextCodec_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextCodec_PTR().Pointer()
	}
	return nil
}

func NewQTextCodecFromPointer(ptr unsafe.Pointer) *QTextCodec {
	var n = new(QTextCodec)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QTextCodec_") {
		n.SetObjectNameAbs("QTextCodec_" + qt.Identifier())
	}
	return n
}

func (ptr *QTextCodec) QTextCodec_PTR() *QTextCodec {
	return ptr
}

//QTextCodec::ConversionFlag
type QTextCodec__ConversionFlag int64

const (
	QTextCodec__DefaultConversion    = QTextCodec__ConversionFlag(0)
	QTextCodec__ConvertInvalidToNull = QTextCodec__ConversionFlag(0x80000000)
	QTextCodec__IgnoreHeader         = QTextCodec__ConversionFlag(0x1)
	QTextCodec__FreeFunction         = QTextCodec__ConversionFlag(0x2)
)

func (ptr *QTextCodec) CanEncode(ch QChar_ITF) bool {
	defer qt.Recovering("QTextCodec::canEncode")

	if ptr.Pointer() != nil {
		return C.QTextCodec_CanEncode(ptr.Pointer(), PointerFromQChar(ch)) != 0
	}
	return false
}

func (ptr *QTextCodec) CanEncode2(s string) bool {
	defer qt.Recovering("QTextCodec::canEncode")

	if ptr.Pointer() != nil {
		return C.QTextCodec_CanEncode2(ptr.Pointer(), C.CString(s)) != 0
	}
	return false
}

func QTextCodec_CodecForHtml2(ba QByteArray_ITF) *QTextCodec {
	defer qt.Recovering("QTextCodec::codecForHtml")

	return NewQTextCodecFromPointer(C.QTextCodec_QTextCodec_CodecForHtml2(PointerFromQByteArray(ba)))
}

func QTextCodec_CodecForHtml(ba QByteArray_ITF, defaultCodec QTextCodec_ITF) *QTextCodec {
	defer qt.Recovering("QTextCodec::codecForHtml")

	return NewQTextCodecFromPointer(C.QTextCodec_QTextCodec_CodecForHtml(PointerFromQByteArray(ba), PointerFromQTextCodec(defaultCodec)))
}

func QTextCodec_CodecForLocale() *QTextCodec {
	defer qt.Recovering("QTextCodec::codecForLocale")

	return NewQTextCodecFromPointer(C.QTextCodec_QTextCodec_CodecForLocale())
}

func QTextCodec_CodecForMib(mib int) *QTextCodec {
	defer qt.Recovering("QTextCodec::codecForMib")

	return NewQTextCodecFromPointer(C.QTextCodec_QTextCodec_CodecForMib(C.int(mib)))
}

func QTextCodec_CodecForName(name QByteArray_ITF) *QTextCodec {
	defer qt.Recovering("QTextCodec::codecForName")

	return NewQTextCodecFromPointer(C.QTextCodec_QTextCodec_CodecForName(PointerFromQByteArray(name)))
}

func QTextCodec_CodecForName2(name string) *QTextCodec {
	defer qt.Recovering("QTextCodec::codecForName")

	return NewQTextCodecFromPointer(C.QTextCodec_QTextCodec_CodecForName2(C.CString(name)))
}

func QTextCodec_CodecForUtfText2(ba QByteArray_ITF) *QTextCodec {
	defer qt.Recovering("QTextCodec::codecForUtfText")

	return NewQTextCodecFromPointer(C.QTextCodec_QTextCodec_CodecForUtfText2(PointerFromQByteArray(ba)))
}

func QTextCodec_CodecForUtfText(ba QByteArray_ITF, defaultCodec QTextCodec_ITF) *QTextCodec {
	defer qt.Recovering("QTextCodec::codecForUtfText")

	return NewQTextCodecFromPointer(C.QTextCodec_QTextCodec_CodecForUtfText(PointerFromQByteArray(ba), PointerFromQTextCodec(defaultCodec)))
}

func (ptr *QTextCodec) FromUnicode(str string) *QByteArray {
	defer qt.Recovering("QTextCodec::fromUnicode")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QTextCodec_FromUnicode(ptr.Pointer(), C.CString(str)))
	}
	return nil
}

func (ptr *QTextCodec) MakeDecoder(flags QTextCodec__ConversionFlag) *QTextDecoder {
	defer qt.Recovering("QTextCodec::makeDecoder")

	if ptr.Pointer() != nil {
		return NewQTextDecoderFromPointer(C.QTextCodec_MakeDecoder(ptr.Pointer(), C.int(flags)))
	}
	return nil
}

func (ptr *QTextCodec) MakeEncoder(flags QTextCodec__ConversionFlag) *QTextEncoder {
	defer qt.Recovering("QTextCodec::makeEncoder")

	if ptr.Pointer() != nil {
		return NewQTextEncoderFromPointer(C.QTextCodec_MakeEncoder(ptr.Pointer(), C.int(flags)))
	}
	return nil
}

func (ptr *QTextCodec) MibEnum() int {
	defer qt.Recovering("QTextCodec::mibEnum")

	if ptr.Pointer() != nil {
		return int(C.QTextCodec_MibEnum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCodec) Name() *QByteArray {
	defer qt.Recovering("QTextCodec::name")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QTextCodec_Name(ptr.Pointer()))
	}
	return nil
}

func QTextCodec_SetCodecForLocale(c QTextCodec_ITF) {
	defer qt.Recovering("QTextCodec::setCodecForLocale")

	C.QTextCodec_QTextCodec_SetCodecForLocale(PointerFromQTextCodec(c))
}

func (ptr *QTextCodec) DestroyQTextCodec() {
	defer qt.Recovering("QTextCodec::~QTextCodec")

	if ptr.Pointer() != nil {
		C.QTextCodec_DestroyQTextCodec(ptr.Pointer())
	}
}

func (ptr *QTextCodec) ObjectNameAbs() string {
	defer qt.Recovering("QTextCodec::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextCodec_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextCodec) SetObjectNameAbs(name string) {
	defer qt.Recovering("QTextCodec::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QTextCodec_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
