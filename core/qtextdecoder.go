package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextDecoder struct {
	ptr unsafe.Pointer
}

type QTextDecoder_ITF interface {
	QTextDecoder_PTR() *QTextDecoder
}

func (p *QTextDecoder) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextDecoder) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextDecoder(ptr QTextDecoder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextDecoder_PTR().Pointer()
	}
	return nil
}

func NewQTextDecoderFromPointer(ptr unsafe.Pointer) *QTextDecoder {
	var n = new(QTextDecoder)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextDecoder) QTextDecoder_PTR() *QTextDecoder {
	return ptr
}

func NewQTextDecoder(codec QTextCodec_ITF) *QTextDecoder {
	defer qt.Recovering("QTextDecoder::QTextDecoder")

	return NewQTextDecoderFromPointer(C.QTextDecoder_NewQTextDecoder(PointerFromQTextCodec(codec)))
}

func NewQTextDecoder2(codec QTextCodec_ITF, flags QTextCodec__ConversionFlag) *QTextDecoder {
	defer qt.Recovering("QTextDecoder::QTextDecoder")

	return NewQTextDecoderFromPointer(C.QTextDecoder_NewQTextDecoder2(PointerFromQTextCodec(codec), C.int(flags)))
}

func (ptr *QTextDecoder) DestroyQTextDecoder() {
	defer qt.Recovering("QTextDecoder::~QTextDecoder")

	if ptr.Pointer() != nil {
		C.QTextDecoder_DestroyQTextDecoder(ptr.Pointer())
	}
}
