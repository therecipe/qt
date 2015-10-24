package xml

//#include "qxmlparseexception.h"
import "C"
import (
	"unsafe"
)

type QXmlParseException struct {
	ptr unsafe.Pointer
}

type QXmlParseExceptionITF interface {
	QXmlParseExceptionPTR() *QXmlParseException
}

func (p *QXmlParseException) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlParseException) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlParseException(ptr QXmlParseExceptionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlParseExceptionPTR().Pointer()
	}
	return nil
}

func QXmlParseExceptionFromPointer(ptr unsafe.Pointer) *QXmlParseException {
	var n = new(QXmlParseException)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlParseException) QXmlParseExceptionPTR() *QXmlParseException {
	return ptr
}

func NewQXmlParseException(name string, c int, l int, p string, s string) *QXmlParseException {
	return QXmlParseExceptionFromPointer(unsafe.Pointer(C.QXmlParseException_NewQXmlParseException(C.CString(name), C.int(c), C.int(l), C.CString(p), C.CString(s))))
}

func NewQXmlParseException2(other QXmlParseExceptionITF) *QXmlParseException {
	return QXmlParseExceptionFromPointer(unsafe.Pointer(C.QXmlParseException_NewQXmlParseException2(C.QtObjectPtr(PointerFromQXmlParseException(other)))))
}

func (ptr *QXmlParseException) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QXmlParseException_ColumnNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlParseException) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QXmlParseException_LineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlParseException) Message() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_Message(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlParseException) PublicId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_PublicId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlParseException) SystemId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_SystemId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlParseException) DestroyQXmlParseException() {
	if ptr.Pointer() != nil {
		C.QXmlParseException_DestroyQXmlParseException(C.QtObjectPtr(ptr.Pointer()))
	}
}
