package xmlpatterns

//#include "qxmlname.h"
import "C"
import (
	"unsafe"
)

type QXmlName struct {
	ptr unsafe.Pointer
}

type QXmlNameITF interface {
	QXmlNamePTR() *QXmlName
}

func (p *QXmlName) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlName) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlName(ptr QXmlNameITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamePTR().Pointer()
	}
	return nil
}

func QXmlNameFromPointer(ptr unsafe.Pointer) *QXmlName {
	var n = new(QXmlName)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlName) QXmlNamePTR() *QXmlName {
	return ptr
}

func NewQXmlName() *QXmlName {
	return QXmlNameFromPointer(unsafe.Pointer(C.QXmlName_NewQXmlName()))
}

func NewQXmlName2(namePool QXmlNamePoolITF, localName string, namespaceURI string, prefix string) *QXmlName {
	return QXmlNameFromPointer(unsafe.Pointer(C.QXmlName_NewQXmlName2(C.QtObjectPtr(PointerFromQXmlNamePool(namePool)), C.CString(localName), C.CString(namespaceURI), C.CString(prefix))))
}

func QXmlName_IsNCName(candidate string) bool {
	return C.QXmlName_QXmlName_IsNCName(C.CString(candidate)) != 0
}

func (ptr *QXmlName) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QXmlName_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlName) LocalName(namePool QXmlNamePoolITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_LocalName(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNamePool(namePool))))
	}
	return ""
}

func (ptr *QXmlName) NamespaceUri(namePool QXmlNamePoolITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_NamespaceUri(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNamePool(namePool))))
	}
	return ""
}

func (ptr *QXmlName) Prefix(namePool QXmlNamePoolITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_Prefix(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNamePool(namePool))))
	}
	return ""
}

func (ptr *QXmlName) ToClarkName(namePool QXmlNamePoolITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlName_ToClarkName(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNamePool(namePool))))
	}
	return ""
}
