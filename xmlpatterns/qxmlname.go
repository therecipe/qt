package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlName struct {
	ptr unsafe.Pointer
}

type QXmlName_ITF interface {
	QXmlName_PTR() *QXmlName
}

func (p *QXmlName) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlName) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlName(ptr QXmlName_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlName_PTR().Pointer()
	}
	return nil
}

func NewQXmlNameFromPointer(ptr unsafe.Pointer) *QXmlName {
	var n = new(QXmlName)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlName) QXmlName_PTR() *QXmlName {
	return ptr
}

func NewQXmlName() *QXmlName {
	defer qt.Recovering("QXmlName::QXmlName")

	return NewQXmlNameFromPointer(C.QXmlName_NewQXmlName())
}

func NewQXmlName2(namePool QXmlNamePool_ITF, localName string, namespaceURI string, prefix string) *QXmlName {
	defer qt.Recovering("QXmlName::QXmlName")

	return NewQXmlNameFromPointer(C.QXmlName_NewQXmlName2(PointerFromQXmlNamePool(namePool), C.CString(localName), C.CString(namespaceURI), C.CString(prefix)))
}

func QXmlName_IsNCName(candidate string) bool {
	defer qt.Recovering("QXmlName::isNCName")

	return C.QXmlName_QXmlName_IsNCName(C.CString(candidate)) != 0
}

func (ptr *QXmlName) IsNull() bool {
	defer qt.Recovering("QXmlName::isNull")

	if ptr.Pointer() != nil {
		return C.QXmlName_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlName) LocalName(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_LocalName(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) NamespaceUri(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::namespaceUri")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_NamespaceUri(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) Prefix(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_Prefix(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) ToClarkName(namePool QXmlNamePool_ITF) string {
	defer qt.Recovering("QXmlName::toClarkName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_ToClarkName(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}
