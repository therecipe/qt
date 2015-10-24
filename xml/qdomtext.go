package xml

//#include "qdomtext.h"
import "C"
import (
	"unsafe"
)

type QDomText struct {
	QDomCharacterData
}

type QDomTextITF interface {
	QDomCharacterDataITF
	QDomTextPTR() *QDomText
}

func PointerFromQDomText(ptr QDomTextITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomTextPTR().Pointer()
	}
	return nil
}

func QDomTextFromPointer(ptr unsafe.Pointer) *QDomText {
	var n = new(QDomText)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomText) QDomTextPTR() *QDomText {
	return ptr
}

func NewQDomText() *QDomText {
	return QDomTextFromPointer(unsafe.Pointer(C.QDomText_NewQDomText()))
}

func NewQDomText2(x QDomTextITF) *QDomText {
	return QDomTextFromPointer(unsafe.Pointer(C.QDomText_NewQDomText2(C.QtObjectPtr(PointerFromQDomText(x)))))
}

func (ptr *QDomText) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomText_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
