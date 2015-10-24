package xml

//#include "qdomcdatasection.h"
import "C"
import (
	"unsafe"
)

type QDomCDATASection struct {
	QDomText
}

type QDomCDATASectionITF interface {
	QDomTextITF
	QDomCDATASectionPTR() *QDomCDATASection
}

func PointerFromQDomCDATASection(ptr QDomCDATASectionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCDATASectionPTR().Pointer()
	}
	return nil
}

func QDomCDATASectionFromPointer(ptr unsafe.Pointer) *QDomCDATASection {
	var n = new(QDomCDATASection)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomCDATASection) QDomCDATASectionPTR() *QDomCDATASection {
	return ptr
}

func NewQDomCDATASection() *QDomCDATASection {
	return QDomCDATASectionFromPointer(unsafe.Pointer(C.QDomCDATASection_NewQDomCDATASection()))
}

func NewQDomCDATASection2(x QDomCDATASectionITF) *QDomCDATASection {
	return QDomCDATASectionFromPointer(unsafe.Pointer(C.QDomCDATASection_NewQDomCDATASection2(C.QtObjectPtr(PointerFromQDomCDATASection(x)))))
}

func (ptr *QDomCDATASection) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomCDATASection_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
