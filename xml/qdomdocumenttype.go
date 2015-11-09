package xml

//#include "qdomdocumenttype.h"
import "C"
import (
	"unsafe"
)

type QDomDocumentType struct {
	QDomNode
}

type QDomDocumentType_ITF interface {
	QDomNode_ITF
	QDomDocumentType_PTR() *QDomDocumentType
}

func PointerFromQDomDocumentType(ptr QDomDocumentType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentType_PTR().Pointer()
	}
	return nil
}

func NewQDomDocumentTypeFromPointer(ptr unsafe.Pointer) *QDomDocumentType {
	var n = new(QDomDocumentType)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomDocumentType) QDomDocumentType_PTR() *QDomDocumentType {
	return ptr
}

func NewQDomDocumentType() *QDomDocumentType {
	return NewQDomDocumentTypeFromPointer(C.QDomDocumentType_NewQDomDocumentType())
}

func NewQDomDocumentType2(n QDomDocumentType_ITF) *QDomDocumentType {
	return NewQDomDocumentTypeFromPointer(C.QDomDocumentType_NewQDomDocumentType2(PointerFromQDomDocumentType(n)))
}

func (ptr *QDomDocumentType) InternalSubset() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_InternalSubset(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocumentType_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomDocumentType) PublicId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) SystemId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_SystemId(ptr.Pointer()))
	}
	return ""
}
