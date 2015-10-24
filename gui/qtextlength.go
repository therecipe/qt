package gui

//#include "qtextlength.h"
import "C"
import (
	"unsafe"
)

type QTextLength struct {
	ptr unsafe.Pointer
}

type QTextLengthITF interface {
	QTextLengthPTR() *QTextLength
}

func (p *QTextLength) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextLength) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextLength(ptr QTextLengthITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextLengthPTR().Pointer()
	}
	return nil
}

func QTextLengthFromPointer(ptr unsafe.Pointer) *QTextLength {
	var n = new(QTextLength)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextLength) QTextLengthPTR() *QTextLength {
	return ptr
}

//QTextLength::Type
type QTextLength__Type int

var (
	QTextLength__VariableLength   = QTextLength__Type(0)
	QTextLength__FixedLength      = QTextLength__Type(1)
	QTextLength__PercentageLength = QTextLength__Type(2)
)

func NewQTextLength() *QTextLength {
	return QTextLengthFromPointer(unsafe.Pointer(C.QTextLength_NewQTextLength()))
}

func (ptr *QTextLength) Type() QTextLength__Type {
	if ptr.Pointer() != nil {
		return QTextLength__Type(C.QTextLength_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
