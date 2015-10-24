package xml

//#include "qdomdocumenttype.h"
import "C"
import (
	"unsafe"
)

type QDomDocumentType struct {
	QDomNode
}

type QDomDocumentTypeITF interface {
	QDomNodeITF
	QDomDocumentTypePTR() *QDomDocumentType
}

func PointerFromQDomDocumentType(ptr QDomDocumentTypeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentTypePTR().Pointer()
	}
	return nil
}

func QDomDocumentTypeFromPointer(ptr unsafe.Pointer) *QDomDocumentType {
	var n = new(QDomDocumentType)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomDocumentType) QDomDocumentTypePTR() *QDomDocumentType {
	return ptr
}

func NewQDomDocumentType() *QDomDocumentType {
	return QDomDocumentTypeFromPointer(unsafe.Pointer(C.QDomDocumentType_NewQDomDocumentType()))
}

func NewQDomDocumentType2(n QDomDocumentTypeITF) *QDomDocumentType {
	return QDomDocumentTypeFromPointer(unsafe.Pointer(C.QDomDocumentType_NewQDomDocumentType2(C.QtObjectPtr(PointerFromQDomDocumentType(n)))))
}

func (ptr *QDomDocumentType) InternalSubset() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_InternalSubset(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDomDocumentType) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDomDocumentType) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocumentType_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomDocumentType) PublicId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_PublicId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDomDocumentType) SystemId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_SystemId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
