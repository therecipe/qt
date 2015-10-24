package core

//#include "qtextdecoder.h"
import "C"
import (
	"unsafe"
)

type QTextDecoder struct {
	ptr unsafe.Pointer
}

type QTextDecoderITF interface {
	QTextDecoderPTR() *QTextDecoder
}

func (p *QTextDecoder) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextDecoder) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextDecoder(ptr QTextDecoderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextDecoderPTR().Pointer()
	}
	return nil
}

func QTextDecoderFromPointer(ptr unsafe.Pointer) *QTextDecoder {
	var n = new(QTextDecoder)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextDecoder) QTextDecoderPTR() *QTextDecoder {
	return ptr
}

func NewQTextDecoder(codec QTextCodecITF) *QTextDecoder {
	return QTextDecoderFromPointer(unsafe.Pointer(C.QTextDecoder_NewQTextDecoder(C.QtObjectPtr(PointerFromQTextCodec(codec)))))
}

func NewQTextDecoder2(codec QTextCodecITF, flags QTextCodec__ConversionFlag) *QTextDecoder {
	return QTextDecoderFromPointer(unsafe.Pointer(C.QTextDecoder_NewQTextDecoder2(C.QtObjectPtr(PointerFromQTextCodec(codec)), C.int(flags))))
}

func (ptr *QTextDecoder) DestroyQTextDecoder() {
	if ptr.Pointer() != nil {
		C.QTextDecoder_DestroyQTextDecoder(C.QtObjectPtr(ptr.Pointer()))
	}
}
