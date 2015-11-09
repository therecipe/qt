package xml

//#include "qdomtext.h"
import "C"
import (
	"unsafe"
)

type QDomText struct {
	QDomCharacterData
}

type QDomText_ITF interface {
	QDomCharacterData_ITF
	QDomText_PTR() *QDomText
}

func PointerFromQDomText(ptr QDomText_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomText_PTR().Pointer()
	}
	return nil
}

func NewQDomTextFromPointer(ptr unsafe.Pointer) *QDomText {
	var n = new(QDomText)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomText) QDomText_PTR() *QDomText {
	return ptr
}

func NewQDomText() *QDomText {
	return NewQDomTextFromPointer(C.QDomText_NewQDomText())
}

func NewQDomText2(x QDomText_ITF) *QDomText {
	return NewQDomTextFromPointer(C.QDomText_NewQDomText2(PointerFromQDomText(x)))
}

func (ptr *QDomText) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomText_NodeType(ptr.Pointer()))
	}
	return 0
}
