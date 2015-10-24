package xml

//#include "qxmlnamespacesupport.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QXmlNamespaceSupport struct {
	ptr unsafe.Pointer
}

type QXmlNamespaceSupportITF interface {
	QXmlNamespaceSupportPTR() *QXmlNamespaceSupport
}

func (p *QXmlNamespaceSupport) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlNamespaceSupport) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlNamespaceSupport(ptr QXmlNamespaceSupportITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamespaceSupportPTR().Pointer()
	}
	return nil
}

func QXmlNamespaceSupportFromPointer(ptr unsafe.Pointer) *QXmlNamespaceSupport {
	var n = new(QXmlNamespaceSupport)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlNamespaceSupport) QXmlNamespaceSupportPTR() *QXmlNamespaceSupport {
	return ptr
}

func NewQXmlNamespaceSupport() *QXmlNamespaceSupport {
	return QXmlNamespaceSupportFromPointer(unsafe.Pointer(C.QXmlNamespaceSupport_NewQXmlNamespaceSupport()))
}

func (ptr *QXmlNamespaceSupport) PopContext() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_PopContext(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlNamespaceSupport) Prefix(uri string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlNamespaceSupport_Prefix(C.QtObjectPtr(ptr.Pointer()), C.CString(uri)))
	}
	return ""
}

func (ptr *QXmlNamespaceSupport) Prefixes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QXmlNamespaceSupport_Prefixes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QXmlNamespaceSupport) Prefixes2(uri string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QXmlNamespaceSupport_Prefixes2(C.QtObjectPtr(ptr.Pointer()), C.CString(uri))), "|")
	}
	return make([]string, 0)
}

func (ptr *QXmlNamespaceSupport) PushContext() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_PushContext(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlNamespaceSupport) Reset() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlNamespaceSupport) SetPrefix(pre string, uri string) {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_SetPrefix(C.QtObjectPtr(ptr.Pointer()), C.CString(pre), C.CString(uri))
	}
}

func (ptr *QXmlNamespaceSupport) Uri(prefix string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlNamespaceSupport_Uri(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix)))
	}
	return ""
}

func (ptr *QXmlNamespaceSupport) DestroyQXmlNamespaceSupport() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(C.QtObjectPtr(ptr.Pointer()))
	}
}
