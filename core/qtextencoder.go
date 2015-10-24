package core

//#include "qtextencoder.h"
import "C"
import (
	"unsafe"
)

type QTextEncoder struct {
	ptr unsafe.Pointer
}

type QTextEncoderITF interface {
	QTextEncoderPTR() *QTextEncoder
}

func (p *QTextEncoder) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextEncoder) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextEncoder(ptr QTextEncoderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextEncoderPTR().Pointer()
	}
	return nil
}

func QTextEncoderFromPointer(ptr unsafe.Pointer) *QTextEncoder {
	var n = new(QTextEncoder)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextEncoder) QTextEncoderPTR() *QTextEncoder {
	return ptr
}

func NewQTextEncoder(codec QTextCodecITF) *QTextEncoder {
	return QTextEncoderFromPointer(unsafe.Pointer(C.QTextEncoder_NewQTextEncoder(C.QtObjectPtr(PointerFromQTextCodec(codec)))))
}

func NewQTextEncoder2(codec QTextCodecITF, flags QTextCodec__ConversionFlag) *QTextEncoder {
	return QTextEncoderFromPointer(unsafe.Pointer(C.QTextEncoder_NewQTextEncoder2(C.QtObjectPtr(PointerFromQTextCodec(codec)), C.int(flags))))
}

func (ptr *QTextEncoder) DestroyQTextEncoder() {
	if ptr.Pointer() != nil {
		C.QTextEncoder_DestroyQTextEncoder(C.QtObjectPtr(ptr.Pointer()))
	}
}
