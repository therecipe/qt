package core

//#include "qjsonparseerror.h"
import "C"
import (
	"unsafe"
)

type QJsonParseError struct {
	ptr unsafe.Pointer
}

type QJsonParseError_ITF interface {
	QJsonParseError_PTR() *QJsonParseError
}

func (p *QJsonParseError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJsonParseError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJsonParseError(ptr QJsonParseError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonParseError_PTR().Pointer()
	}
	return nil
}

func NewQJsonParseErrorFromPointer(ptr unsafe.Pointer) *QJsonParseError {
	var n = new(QJsonParseError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJsonParseError) QJsonParseError_PTR() *QJsonParseError {
	return ptr
}

//QJsonParseError::ParseError
type QJsonParseError__ParseError int64

const (
	QJsonParseError__NoError               = QJsonParseError__ParseError(0)
	QJsonParseError__UnterminatedObject    = QJsonParseError__ParseError(1)
	QJsonParseError__MissingNameSeparator  = QJsonParseError__ParseError(2)
	QJsonParseError__UnterminatedArray     = QJsonParseError__ParseError(3)
	QJsonParseError__MissingValueSeparator = QJsonParseError__ParseError(4)
	QJsonParseError__IllegalValue          = QJsonParseError__ParseError(5)
	QJsonParseError__TerminationByNumber   = QJsonParseError__ParseError(6)
	QJsonParseError__IllegalNumber         = QJsonParseError__ParseError(7)
	QJsonParseError__IllegalEscapeSequence = QJsonParseError__ParseError(8)
	QJsonParseError__IllegalUTF8String     = QJsonParseError__ParseError(9)
	QJsonParseError__UnterminatedString    = QJsonParseError__ParseError(10)
	QJsonParseError__MissingObject         = QJsonParseError__ParseError(11)
	QJsonParseError__DeepNesting           = QJsonParseError__ParseError(12)
	QJsonParseError__DocumentTooLarge      = QJsonParseError__ParseError(13)
	QJsonParseError__GarbageAtEnd          = QJsonParseError__ParseError(14)
)

func (ptr *QJsonParseError) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QJsonParseError_ErrorString(ptr.Pointer()))
	}
	return ""
}
