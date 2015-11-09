package xml

//#include "qdomdocumentfragment.h"
import "C"
import (
	"unsafe"
)

type QDomDocumentFragment struct {
	QDomNode
}

type QDomDocumentFragment_ITF interface {
	QDomNode_ITF
	QDomDocumentFragment_PTR() *QDomDocumentFragment
}

func PointerFromQDomDocumentFragment(ptr QDomDocumentFragment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentFragment_PTR().Pointer()
	}
	return nil
}

func NewQDomDocumentFragmentFromPointer(ptr unsafe.Pointer) *QDomDocumentFragment {
	var n = new(QDomDocumentFragment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomDocumentFragment) QDomDocumentFragment_PTR() *QDomDocumentFragment {
	return ptr
}

func NewQDomDocumentFragment() *QDomDocumentFragment {
	return NewQDomDocumentFragmentFromPointer(C.QDomDocumentFragment_NewQDomDocumentFragment())
}

func NewQDomDocumentFragment2(x QDomDocumentFragment_ITF) *QDomDocumentFragment {
	return NewQDomDocumentFragmentFromPointer(C.QDomDocumentFragment_NewQDomDocumentFragment2(PointerFromQDomDocumentFragment(x)))
}

func (ptr *QDomDocumentFragment) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocumentFragment_NodeType(ptr.Pointer()))
	}
	return 0
}
