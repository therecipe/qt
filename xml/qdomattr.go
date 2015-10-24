package xml

//#include "qdomattr.h"
import "C"
import (
	"unsafe"
)

type QDomAttr struct {
	QDomNode
}

type QDomAttrITF interface {
	QDomNodeITF
	QDomAttrPTR() *QDomAttr
}

func PointerFromQDomAttr(ptr QDomAttrITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomAttrPTR().Pointer()
	}
	return nil
}

func QDomAttrFromPointer(ptr unsafe.Pointer) *QDomAttr {
	var n = new(QDomAttr)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomAttr) QDomAttrPTR() *QDomAttr {
	return ptr
}

func NewQDomAttr() *QDomAttr {
	return QDomAttrFromPointer(unsafe.Pointer(C.QDomAttr_NewQDomAttr()))
}

func NewQDomAttr2(x QDomAttrITF) *QDomAttr {
	return QDomAttrFromPointer(unsafe.Pointer(C.QDomAttr_NewQDomAttr2(C.QtObjectPtr(PointerFromQDomAttr(x)))))
}

func (ptr *QDomAttr) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomAttr_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDomAttr) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomAttr_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomAttr) SetValue(v string) {
	if ptr.Pointer() != nil {
		C.QDomAttr_SetValue(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QDomAttr) Specified() bool {
	if ptr.Pointer() != nil {
		return C.QDomAttr_Specified(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomAttr) Value() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomAttr_Value(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
