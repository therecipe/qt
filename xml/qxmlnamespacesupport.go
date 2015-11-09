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

type QXmlNamespaceSupport_ITF interface {
	QXmlNamespaceSupport_PTR() *QXmlNamespaceSupport
}

func (p *QXmlNamespaceSupport) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlNamespaceSupport) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlNamespaceSupport(ptr QXmlNamespaceSupport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamespaceSupport_PTR().Pointer()
	}
	return nil
}

func NewQXmlNamespaceSupportFromPointer(ptr unsafe.Pointer) *QXmlNamespaceSupport {
	var n = new(QXmlNamespaceSupport)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlNamespaceSupport) QXmlNamespaceSupport_PTR() *QXmlNamespaceSupport {
	return ptr
}

func NewQXmlNamespaceSupport() *QXmlNamespaceSupport {
	return NewQXmlNamespaceSupportFromPointer(C.QXmlNamespaceSupport_NewQXmlNamespaceSupport())
}

func (ptr *QXmlNamespaceSupport) PopContext() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_PopContext(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) Prefix(uri string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlNamespaceSupport_Prefix(ptr.Pointer(), C.CString(uri)))
	}
	return ""
}

func (ptr *QXmlNamespaceSupport) Prefixes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QXmlNamespaceSupport_Prefixes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QXmlNamespaceSupport) Prefixes2(uri string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QXmlNamespaceSupport_Prefixes2(ptr.Pointer(), C.CString(uri))), "|")
	}
	return make([]string, 0)
}

func (ptr *QXmlNamespaceSupport) PushContext() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_PushContext(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) Reset() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_Reset(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) SetPrefix(pre string, uri string) {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_SetPrefix(ptr.Pointer(), C.CString(pre), C.CString(uri))
	}
}

func (ptr *QXmlNamespaceSupport) Uri(prefix string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlNamespaceSupport_Uri(ptr.Pointer(), C.CString(prefix)))
	}
	return ""
}

func (ptr *QXmlNamespaceSupport) DestroyQXmlNamespaceSupport() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(ptr.Pointer())
	}
}
