package xml

//#include "qdomdocumentfragment.h"
import "C"
import (
	"unsafe"
)

type QDomDocumentFragment struct {
	QDomNode
}

type QDomDocumentFragmentITF interface {
	QDomNodeITF
	QDomDocumentFragmentPTR() *QDomDocumentFragment
}

func PointerFromQDomDocumentFragment(ptr QDomDocumentFragmentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentFragmentPTR().Pointer()
	}
	return nil
}

func QDomDocumentFragmentFromPointer(ptr unsafe.Pointer) *QDomDocumentFragment {
	var n = new(QDomDocumentFragment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomDocumentFragment) QDomDocumentFragmentPTR() *QDomDocumentFragment {
	return ptr
}

func NewQDomDocumentFragment() *QDomDocumentFragment {
	return QDomDocumentFragmentFromPointer(unsafe.Pointer(C.QDomDocumentFragment_NewQDomDocumentFragment()))
}

func NewQDomDocumentFragment2(x QDomDocumentFragmentITF) *QDomDocumentFragment {
	return QDomDocumentFragmentFromPointer(unsafe.Pointer(C.QDomDocumentFragment_NewQDomDocumentFragment2(C.QtObjectPtr(PointerFromQDomDocumentFragment(x)))))
}

func (ptr *QDomDocumentFragment) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocumentFragment_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
