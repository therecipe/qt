package gui

//#include "qtextlength.h"
import "C"
import (
	"unsafe"
)

type QTextLength struct {
	ptr unsafe.Pointer
}

type QTextLength_ITF interface {
	QTextLength_PTR() *QTextLength
}

func (p *QTextLength) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextLength) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextLength(ptr QTextLength_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextLength_PTR().Pointer()
	}
	return nil
}

func NewQTextLengthFromPointer(ptr unsafe.Pointer) *QTextLength {
	var n = new(QTextLength)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextLength) QTextLength_PTR() *QTextLength {
	return ptr
}

//QTextLength::Type
type QTextLength__Type int64

const (
	QTextLength__VariableLength   = QTextLength__Type(0)
	QTextLength__FixedLength      = QTextLength__Type(1)
	QTextLength__PercentageLength = QTextLength__Type(2)
)

func NewQTextLength() *QTextLength {
	return NewQTextLengthFromPointer(C.QTextLength_NewQTextLength())
}

func NewQTextLength2(ty QTextLength__Type, value float64) *QTextLength {
	return NewQTextLengthFromPointer(C.QTextLength_NewQTextLength2(C.int(ty), C.double(value)))
}

func (ptr *QTextLength) RawValue() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextLength_RawValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLength) Type() QTextLength__Type {
	if ptr.Pointer() != nil {
		return QTextLength__Type(C.QTextLength_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLength) Value(maximumLength float64) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextLength_Value(ptr.Pointer(), C.double(maximumLength)))
	}
	return 0
}
