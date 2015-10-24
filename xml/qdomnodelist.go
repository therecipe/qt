package xml

//#include "qdomnodelist.h"
import "C"
import (
	"unsafe"
)

type QDomNodeList struct {
	ptr unsafe.Pointer
}

type QDomNodeListITF interface {
	QDomNodeListPTR() *QDomNodeList
}

func (p *QDomNodeList) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomNodeList) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomNodeList(ptr QDomNodeListITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNodeListPTR().Pointer()
	}
	return nil
}

func QDomNodeListFromPointer(ptr unsafe.Pointer) *QDomNodeList {
	var n = new(QDomNodeList)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomNodeList) QDomNodeListPTR() *QDomNodeList {
	return ptr
}

func NewQDomNodeList() *QDomNodeList {
	return QDomNodeListFromPointer(unsafe.Pointer(C.QDomNodeList_NewQDomNodeList()))
}

func NewQDomNodeList2(n QDomNodeListITF) *QDomNodeList {
	return QDomNodeListFromPointer(unsafe.Pointer(C.QDomNodeList_NewQDomNodeList2(C.QtObjectPtr(PointerFromQDomNodeList(n)))))
}

func (ptr *QDomNodeList) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNodeList) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QDomNodeList_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNodeList) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNodeList) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNodeList) DestroyQDomNodeList() {
	if ptr.Pointer() != nil {
		C.QDomNodeList_DestroyQDomNodeList(C.QtObjectPtr(ptr.Pointer()))
	}
}
