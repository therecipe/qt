package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDomCDATASection struct {
	QDomText
}

type QDomCDATASection_ITF interface {
	QDomText_ITF
	QDomCDATASection_PTR() *QDomCDATASection
}

func PointerFromQDomCDATASection(ptr QDomCDATASection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCDATASection_PTR().Pointer()
	}
	return nil
}

func NewQDomCDATASectionFromPointer(ptr unsafe.Pointer) *QDomCDATASection {
	var n = new(QDomCDATASection)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomCDATASection) QDomCDATASection_PTR() *QDomCDATASection {
	return ptr
}

func NewQDomCDATASection() *QDomCDATASection {
	defer qt.Recovering("QDomCDATASection::QDomCDATASection")

	return NewQDomCDATASectionFromPointer(C.QDomCDATASection_NewQDomCDATASection())
}

func NewQDomCDATASection2(x QDomCDATASection_ITF) *QDomCDATASection {
	defer qt.Recovering("QDomCDATASection::QDomCDATASection")

	return NewQDomCDATASectionFromPointer(C.QDomCDATASection_NewQDomCDATASection2(PointerFromQDomCDATASection(x)))
}

func (ptr *QDomCDATASection) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomCDATASection::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomCDATASection_NodeType(ptr.Pointer()))
	}
	return 0
}
