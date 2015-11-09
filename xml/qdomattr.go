package xml

//#include "qdomattr.h"
import "C"
import (
	"unsafe"
)

type QDomAttr struct {
	QDomNode
}

type QDomAttr_ITF interface {
	QDomNode_ITF
	QDomAttr_PTR() *QDomAttr
}

func PointerFromQDomAttr(ptr QDomAttr_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomAttr_PTR().Pointer()
	}
	return nil
}

func NewQDomAttrFromPointer(ptr unsafe.Pointer) *QDomAttr {
	var n = new(QDomAttr)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomAttr) QDomAttr_PTR() *QDomAttr {
	return ptr
}

func NewQDomAttr() *QDomAttr {
	return NewQDomAttrFromPointer(C.QDomAttr_NewQDomAttr())
}

func NewQDomAttr2(x QDomAttr_ITF) *QDomAttr {
	return NewQDomAttrFromPointer(C.QDomAttr_NewQDomAttr2(PointerFromQDomAttr(x)))
}

func (ptr *QDomAttr) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomAttr_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomAttr) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomAttr_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomAttr) SetValue(v string) {
	if ptr.Pointer() != nil {
		C.QDomAttr_SetValue(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QDomAttr) Specified() bool {
	if ptr.Pointer() != nil {
		return C.QDomAttr_Specified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomAttr) Value() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomAttr_Value(ptr.Pointer()))
	}
	return ""
}
