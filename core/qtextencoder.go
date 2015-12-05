package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QTextEncoder struct {
	ptr unsafe.Pointer
}

type QTextEncoder_ITF interface {
	QTextEncoder_PTR() *QTextEncoder
}

func (p *QTextEncoder) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextEncoder) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextEncoder(ptr QTextEncoder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextEncoder_PTR().Pointer()
	}
	return nil
}

func NewQTextEncoderFromPointer(ptr unsafe.Pointer) *QTextEncoder {
	var n = new(QTextEncoder)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextEncoder) QTextEncoder_PTR() *QTextEncoder {
	return ptr
}

func NewQTextEncoder(codec QTextCodec_ITF) *QTextEncoder {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEncoder::QTextEncoder")
		}
	}()

	return NewQTextEncoderFromPointer(C.QTextEncoder_NewQTextEncoder(PointerFromQTextCodec(codec)))
}

func NewQTextEncoder2(codec QTextCodec_ITF, flags QTextCodec__ConversionFlag) *QTextEncoder {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEncoder::QTextEncoder")
		}
	}()

	return NewQTextEncoderFromPointer(C.QTextEncoder_NewQTextEncoder2(PointerFromQTextCodec(codec), C.int(flags)))
}

func (ptr *QTextEncoder) FromUnicode2(uc QChar_ITF, len int) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEncoder::fromUnicode")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QTextEncoder_FromUnicode2(ptr.Pointer(), PointerFromQChar(uc), C.int(len)))
	}
	return nil
}

func (ptr *QTextEncoder) FromUnicode(str string) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEncoder::fromUnicode")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QTextEncoder_FromUnicode(ptr.Pointer(), C.CString(str)))
	}
	return nil
}

func (ptr *QTextEncoder) DestroyQTextEncoder() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextEncoder::~QTextEncoder")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextEncoder_DestroyQTextEncoder(ptr.Pointer())
	}
}
